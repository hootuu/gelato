package collection

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
