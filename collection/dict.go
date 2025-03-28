package collection

import "github.com/hootuu/gelato/strs"

type Dict struct {
	m map[string]interface{}
}

func NewDict(m map[string]interface{}) *Dict {
	if len(m) == 0 {
		m = make(map[string]interface{})
	}
	return &Dict{
		m: m,
	}
}

func (d *Dict) GetString(key string) string {
	val, ok := d.m[key]
	if !ok {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}

func (d *Dict) GetUint64(key string) uint64 {
	str := d.GetString(key)
	return strs.ToUint64(str)
}
