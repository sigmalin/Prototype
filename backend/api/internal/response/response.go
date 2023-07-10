package response

type Body struct {
	Code    int         `json:"code" example:"0"`
	Message string      `json:"message" example:""`
	Data    interface{} `json:"data"`
}

func (b *Body) Error(code int, msg string) {
	b.Code = code
	b.Message = msg
	b.Data = nil
}
