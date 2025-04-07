package rest

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"resty.dev/v3"
)

var gLogger = logger.GetLogger("rest")

type CliLogger struct {
}

func (c *CliLogger) Errorf(format string, v ...any) {
	gLogger.Error(fmt.Sprintf(format, v...))
}

func (c *CliLogger) Warnf(format string, v ...any) {
	gLogger.Warn(fmt.Sprintf(format, v...))
}

func (c *CliLogger) Debugf(format string, v ...any) {
	gLogger.Debug(fmt.Sprintf(format, v...))
}

func RequestLogger(_ *resty.Client, req *resty.Request) *errors.Error {
	if gLogger.Level() <= zapcore.InfoLevel {
		gLogger.Info("->"+req.URL,
			zap.String("URL", req.URL),
			zap.String("method", req.Method),
			zap.Any("header", req.Header),
		)
	}
	return nil
}

func ResponseLogger(_ *resty.Client, resp *resty.Response) *errors.Error {
	if gLogger.Level() <= zapcore.InfoLevel {
		gLogger.Info("<-"+resp.Request.URL,
			zap.String("URL", resp.Request.URL),
			zap.String("method", resp.Request.Method),
			zap.Bool("success", resp.IsSuccess()),
			zap.String("status", resp.Status()),
		)
	}
	return nil
}
