package appenv

import "go.uber.org/zap"

func panicDuringInit(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
	panic(msg)
}
