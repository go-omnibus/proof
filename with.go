package proof

import "go.uber.org/zap"

func With(k string, v interface{}) zap.Field {
	return zap.Any(k, v)
}

func WithByteString(k string, v []byte) zap.Field {
	return zap.ByteString(k, v)
}

func WithError(err error) zap.Field {
	return zap.NamedError("error", err)
}
