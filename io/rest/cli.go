package rest

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/sys"
	"go.uber.org/zap"
	"resty.dev/v3"
	"time"
)

type Rest struct {
	baseUrl     string
	timeSetter  func(cli *resty.Client)
	requestMid  []func(*resty.Client, *resty.Request) *errors.Error
	responseMid []func(*resty.Client, *resty.Response) *errors.Error
}

func ZeroRest() *Rest {
	return &Rest{
		baseUrl:     "",
		timeSetter:  nil,
		requestMid:  nil,
		responseMid: nil,
	}
}

func NewRest() *Rest {
	return ZeroRest().
		WithRequestMid(RequestLogger).
		WithResponseMid(ResponseLogger, GuardRespMid)
}

func (r *Rest) SetBaseURL(baseUrl string) *Rest {
	r.baseUrl = baseUrl
	return r
}

func (r *Rest) WithTimeSetter(timeSetter func(cli *resty.Client)) *Rest {
	r.timeSetter = timeSetter
	return r
}

func (r *Rest) WithRequestMid(mid ...func(*resty.Client, *resty.Request) *errors.Error) *Rest {
	r.requestMid = mid
	return r
}

func (r *Rest) WithResponseMid(mid ...func(*resty.Client, *resty.Response) *errors.Error) *Rest {
	r.responseMid = mid
	return r
}

func (r *Rest) NewClient() *resty.Client {
	cli := resty.New()
	cli.SetBaseURL(r.baseUrl)
	if r.timeSetter != nil {
		r.timeSetter(cli)
	} else {
		cli.SetRetryWaitTime(2 * time.Second).
			SetRetryMaxWaitTime(10 * time.Second).
			SetTimeout(60 * time.Second)
	}
	if len(r.requestMid) > 0 {
		for _, reqMid := range r.requestMid {
			cli.AddRequestMiddleware(func(cli *resty.Client, req *resty.Request) error {
				err := reqMid(cli, req)
				if err != nil {
					return err.Native()
				}
				return nil
			})
		}
	}
	if len(r.responseMid) > 0 {
		for _, respMid := range r.responseMid {
			cli.AddResponseMiddleware(func(cli *resty.Client, resp *resty.Response) error {
				err := respMid(cli, resp)
				if err != nil {
					return err.Native()
				}
				return nil
			})
		}
	}
	if sys.RunMode.IsLocal() {
		cli.SetLogger(&CliLogger{})
		cli.EnableTrace()
	}
	return cli
}

func Call[REQ any, RESP any](rest *Rest, path string, req *Request[REQ], priKey []byte) *Response[RESP] {
	err := req.Sign(priKey)
	if err != nil {
		return FailResponse[RESP](req.ID, err)
	}
	bodyDataBytes, err := req.Marshal()
	if err != nil {
		return FailResponse[RESP](req.ID, err)
	}
	cli := rest.NewClient()
	var resp Response[RESP]
	_, nErr := cli.R().SetBody(bodyDataBytes).SetResult(&resp).Post(path)
	if nErr != nil {
		gLogger.Error("resty.call err", zap.Error(nErr))
		return FailResponse[RESP](req.ID, errors.System("rest error:"+nErr.Error(), nErr))
	}
	return &resp
}
