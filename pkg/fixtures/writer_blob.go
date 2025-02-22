package fixtures

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/justtrackio/gosoline/pkg/blob"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
)

type BlobFixturesSettings struct {
	BasePath   string
	ConfigName string
}

type blobFixtureWriter struct {
	logger      log.Logger
	batchRunner blob.BatchRunner
	purger      *blobPurger
	store       blob.Store
	basePath    string
}

func BlobFixtureWriterFactory(settings *BlobFixturesSettings) FixtureWriterFactory {
	return func(ctx context.Context, config cfg.Config, logger log.Logger) (FixtureWriter, error) {
		basePath, err := filepath.Abs(settings.BasePath)
		if err != nil {
			return nil, err
		}

		settings.BasePath = basePath

		store, err := blob.NewStore(ctx, config, logger, settings.ConfigName)
		if err != nil {
			return nil, fmt.Errorf("can not create blob store: %w", err)
		}

		br, err := blob.NewBatchRunner(ctx, config, logger)
		if err != nil {
			return nil, fmt.Errorf("can not create blob batch runner: %w", err)
		}

		purger, err := newBlobPurger(ctx, config, logger, settings)
		if err != nil {
			return nil, fmt.Errorf("can not create blob purger: %w", err)
		}

		return NewBlobFixtureWriterWithInterfaces(logger, br, purger, store, settings.BasePath), nil
	}
}

func NewBlobFixtureWriterWithInterfaces(logger log.Logger, batchRunner blob.BatchRunner, purger *blobPurger, store blob.Store, basePath string) FixtureWriter {
	return &blobFixtureWriter{
		logger:      logger,
		batchRunner: batchRunner,
		purger:      purger,
		store:       store,
		basePath:    basePath,
	}
}

func (s *blobFixtureWriter) Purge(ctx context.Context) error {
	return s.purger.purge(ctx)
}

func (s *blobFixtureWriter) Write(ctx context.Context, _ *FixtureSet) error {
	if err := s.store.CreateBucket(ctx); err != nil {
		return fmt.Errorf("can not create bucket: %w", err)
	}

	var files []string
	err := filepath.Walk(s.basePath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	var batch blob.Batch
	for _, file := range files {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		object := blob.Object{
			Key:  aws.String(strings.Replace(file, s.basePath, "", 1)),
			Body: blob.StreamBytes(body),
		}

		batch = append(batch, &object)
	}

	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		err = s.batchRunner.Run(ctx)
	}(ctx)
	defer cancel()

	if err := s.store.Write(batch); err != nil {
		return fmt.Errorf("can not write fixtes: %w", err)
	}

	s.logger.Info("loaded %d files", len(files))

	return err
}
