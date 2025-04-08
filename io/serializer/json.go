package serializer

import (
	"encoding/json"
	"github.com/hootuu/gelato/errors"
)

func JsonTo(obj interface{}) (string, *errors.Error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", errors.Verify("invalid to json", err)
	}
	return string(data), nil
}

func JsonMustTo(obj interface{}) string {
	str, _ := JsonTo(obj)
	return str
}

func JsonMustToBytes(obj interface{}) []byte {
	data, _ := json.Marshal(obj)
	return data
}

func JsonOf[T any](str string) (*T, *errors.Error) {
	var obj T
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		return nil, errors.Verify("invalid from json str", err)
	}
	return &obj, nil
}

func JsonMustOf[T any](str string) *T {
	obj, _ := JsonOf[T](str)
	return obj
}

func JsonMustOfBytes[T any](data []byte) *T {
	var obj T
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil
	}
	return &obj
}

func JsonFrom[T any](obj *T, str string) *errors.Error {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		return errors.Verify("invalid from json str", err)
	}
	return nil
}

func JsonMustFrom[T any](obj *T, str string) {
	_ = JsonFrom[T](obj, str)
}
