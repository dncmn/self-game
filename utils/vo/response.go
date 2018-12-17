package vo

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewData() *Data {
	data := &Data{
		Data: make(map[string]interface{}),
	}

	return data
}
