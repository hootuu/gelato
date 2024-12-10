package errors

import "fmt"

type Template struct {
	Code  string `bson:"code" json:"code"`
	Label string `bson:"label" json:"label"`
}

func TPL(code string, label string) *Template {
	return &Template{
		Code:  code,
		Label: label,
	}
}

func (t *Template) Error(args ...any) *Error {
	if len(args) == 0 {
		return Of(TypeBiz, t.Code, t.Label)
	}
	format, ok := args[0].(string)
	if !ok {
		return Of(TypeBiz, t.Code, t.Label)
	}
	return Of(TypeBiz, t.Code, fmt.Sprintf(t.Label+": "+format, args[1:]...))
}
