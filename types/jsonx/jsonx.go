package jsonx

import (
	"encoding/json"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/logger"
	"go.uber.org/zap"
)

func JSON2Bytes(obj interface{}) ([]byte, *errors.Error) {
	if obj == nil {

	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.Verify("JSON2Bytes: invalid to json", err)
	}
	return data, nil
}

func MustJSON2Bytes(obj interface{}) []byte {
	bData, err := JSON2Bytes(obj)
	if err != nil {
		logger.Error.Error("MustJSON2Bytes: Err", zap.Error(err))
		return nil
	}
	return bData
}

func JSON2String(obj interface{}) (string, *errors.Error) {
	bData, err := JSON2Bytes(obj)
	if err != nil {
		return "", err
	}
	return string(bData), nil
}

func MustJSON2String(obj interface{}) string {
	str, err := JSON2String(obj)
	if err != nil {
		return ""
	}
	return str
}

func JSONByBytes[T any](bData []byte) (*T, *errors.Error) {
	var obj T
	err := json.Unmarshal(bData, &obj)
	if err != nil {
		return nil, errors.Verify("invalid from json bytes", err)
	}
	return &obj, nil
}

func MustJSONByBytes[T any](bData []byte) *T {
	obj, err := JSONByBytes[T](bData)
	if err != nil {
		logger.Error.Error("MustJSONByBytes: err", zap.Error(err))
		return nil
	}
	return obj
}

func JSONByString[T any](str string) (*T, *errors.Error) {
	return JSONByBytes[T]([]byte(str))
}

func MustJSONByString[T any](str string) *T {
	obj, err := JSONByString[T](str)
	if err != nil {
		return nil
	}
	return obj
}

func JSONOfBytes[T any](ptr *T, bData []byte) *errors.Error {
	err := json.Unmarshal(bData, ptr)
	if err != nil {
		return errors.Verify("JSONOfBytes: err", err)
	}
	return nil
}

func MustJSONOfBytes[T any](ptr *T, bData []byte) {
	err := JSONOfBytes[T](ptr, bData)
	if err != nil {
		logger.Error.Error("MustJSONOfBytes: err", zap.Error(err))
	}
}

func JSONOfString[T any](ptr *T, str string) *errors.Error {
	return JSONOfBytes[T](ptr, []byte(str))
}

func MustJSONOfString[T any](ptr *T, str string) {
	JSONOfString[T](ptr, str)
}
