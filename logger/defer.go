package logger

import (
	"go.uber.org/zap"
	"time"
)

func Elapse(fun string, logger *zap.Logger, fix ...func() []zap.Field) func() {
	start := time.Now().UnixMilli()
	var prefixFields []zap.Field
	if len(fix) > 0 {
		prefixFields = fix[0]()
	}
	if len(prefixFields) > 0 {
		logger.Info("==>"+fun, prefixFields...)
	} else {
		logger.Info("==>" + fun)
	}

	return func() {
		elapse := time.Now().UnixMilli() - start
		suffixFields := []zap.Field{zap.Int64("_elapse", elapse)}
		if len(fix) > 1 {
			fs := fix[1]()
			if len(fs) > 1 {
				suffixFields = append(suffixFields, fs...)
			}
		}
		if len(suffixFields) > 0 {
			logger.Info("<=="+fun, suffixFields...)
		} else {
			logger.Info("<==" + fun)
		}
	}
}
