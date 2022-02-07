package stream

import (
	"context"
	"fmt"
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"
	gosoKinesis "github.com/justtrackio/gosoline/pkg/cloud/aws/kinesis"
	"github.com/justtrackio/gosoline/pkg/exec"
	"github.com/justtrackio/gosoline/pkg/log"
)

type KinesisOutputSettings struct {
	cfg.AppId
	StreamName string
}

type kinesisOutput struct {
	recordWriter gosoKinesis.RecordWriter
}

func NewKinesisOutput(ctx context.Context, config cfg.Config, logger log.Logger, settings *KinesisOutputSettings) (Output, error) {
	var err error
	var recordWriter gosoKinesis.RecordWriter

	settings.PadFromConfig(config)
	fullStreamName := fmt.Sprintf("%s-%s-%s-%s-%s", settings.Project, settings.Environment, settings.Family, settings.Application, settings.StreamName)
	backoffSettings := exec.ReadBackoffSettings(config)
	backoffSettings.InitialInterval = time.Second

	recordWriterSettings := &gosoKinesis.RecordWriterSettings{
		StreamName: fullStreamName,
		Backoff:    backoffSettings,
	}

	if recordWriter, err = gosoKinesis.NewRecordWriter(ctx, config, logger, recordWriterSettings); err != nil {
		return nil, fmt.Errorf("can not create record writer for stream %s: %w", fullStreamName, err)
	}

	return NewKinesisOutputWithInterfaces(recordWriter), nil
}

func NewKinesisOutputWithInterfaces(recordWriter gosoKinesis.RecordWriter) Output {
	return &kinesisOutput{
		recordWriter: recordWriter,
	}
}

func (o *kinesisOutput) WriteOne(ctx context.Context, record WritableMessage) error {
	return o.Write(ctx, []WritableMessage{record})
}

func (o *kinesisOutput) Write(ctx context.Context, batch []WritableMessage) error {
	var err error
	var record []byte
	var records [][]byte

	for _, msg := range batch {
		if record, err = msg.MarshalToBytes(); err != nil {
			return fmt.Errorf("can not marshal message to bytes: %w", err)
		}

		records = append(records, record)
	}

	return o.recordWriter.PutRecords(ctx, records)
}
