package vo

type Data struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewData() *Data {
	data := &Data{
		Data: make(map[string]interface{}),
	}

	return data
}
