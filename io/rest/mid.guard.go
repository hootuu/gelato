package rest

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"resty.dev/v3"
)

func GuardRespMid(_ *resty.Client, restyResp *resty.Response) *errors.Error {
	if restyResp == nil {
		return errors.E("901", "response is null")
	}
	if !restyResp.IsSuccess() {
		return errors.E(
			fmt.Sprintf("%d", restyResp.StatusCode()),
			"response is invalid: %s", restyResp.Status(),
		)
	}
	return nil
}
