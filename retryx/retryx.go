package retryx

import (
	"github.com/avast/retry-go"
	"github.com/hootuu/gelato/configure"
	"github.com/hootuu/gelato/logger"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

func Universal(call func() error) {
	err := retry.Do(func() error {
		return call()
	},
		retry.Attempts(cast.ToUint(configure.GetInt("retry.universal.attempts", 3))),
		retry.Delay(configure.GetDuration("retry.universal.delay", 600*time.Millisecond)),
	)
	if err != nil {
		logger.Error.Error("retryx.must.do err", zap.Error(err))
	}
}
