package errors

import (
	"errors"
	"fmt"
)

type Type = int32

const (
	TypeCoding Type = -999 // code error
	TypeSystem Type = -444 // system error
	TypeBiz    Type = -777 // biz error
)

type Error struct {
	Type    Type   `bson:"type" json:"type"`
	Code    string `bson:"code"  json:"code"`
	Message string `bson:"message" json:"message"`

	Err error `bson:"-" json:"-"` // Native error
}

func Of(t Type, code string, message string, nativeErr ...error) *Error {
	err := &Error{
		Type:    t,
		Code:    code,
		Message: message,
	}
	if len(nativeErr) > 0 && nativeErr[0] != nil {
		err.Err = nativeErr[0]
	}
	return err
}

func Assert(expect string, actual string, nativeErr ...error) *Error {
	return Of(TypeCoding,
		"assert",
		fmt.Sprintf("assert failed, expect: %s, but %s", expect, actual),
		nativeErr...)
}

func System(message string, nativeErr ...error) *Error {
	return Of(TypeSystem,
		"system",
		message,
		nativeErr...)
}

func Verify(message string, nativeErr ...error) *Error {

	return Of(TypeBiz, "verify", message, nativeErr...)
}

func (e *Error) Error() string {
	if e.Err != nil && !errors.Is(e, e.Err) {
		return fmt.Sprintf("[%s] %s. ******** %s ********", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *Error) GetType() Type {
	return e.Type
}

func (e *Error) GetCode() string {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) Native() error {
	if e.Err == nil {
		return errors.New(fmt.Sprintf("[%s] %s", e.Code, e.Message))
	}
	return e.Err
}
