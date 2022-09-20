package proof

import (
	"github.com/luci/go-render/render"
	"go.uber.org/zap"
)

func With(k string, v interface{}) zap.Field {
	return zap.Any(k, v)
}

func WithJSONByte(k string, v []byte) zap.Field {
	return zap.ByteString(k, v)
}

func WithError(err error) zap.Field {
	return zap.NamedError("error", err)
}

func WithStruct(v interface{}) zap.Field {
	return zap.String("s", render.Render(v))
}

func Render(k string, v interface{}) zap.Field {
	return zap.String(k, render.Render(v))
}
