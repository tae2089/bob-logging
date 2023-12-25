package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type EscapeSeqJSONEncoder struct {
	zapcore.Encoder
}

func (enc *EscapeSeqJSONEncoder) Clone() zapcore.Encoder {
	return enc // TODO: change me
}

func (enc *EscapeSeqJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// call EncodeEntry on the embedded interface to get the
	// original output
	b, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	newb := buffer.NewPool().Get()

	// then manipulate that output into what you need it to be
	newb.Write(bytes.Replace(b.Bytes(), []byte("\\n"), []byte("\n"), -1))
	newb.Write(bytes.Replace(newb.Bytes(), []byte("\\t"), []byte("\t"), -1))
	return newb, nil
}

type CustomWriter struct{}

func (e CustomWriter) Write(p []byte) (int, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, p, "", "    "); err != nil {
		return 0, err
	}
	n, err := os.Stdout.Write(prettyJSON.Bytes())
	if err != nil {
		return n, err
	}
	if n != len(prettyJSON.Bytes()) {
		return n, io.ErrShortWrite
	}
	return len(p), nil
}
