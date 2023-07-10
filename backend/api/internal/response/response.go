package response

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (b *Body) Error(code int, msg string) {
	b.Code = code
	b.Message = msg
	b.Data = nil
}
