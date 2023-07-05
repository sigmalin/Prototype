package custom

import (
	"net/http"
)

func NewResponse(w http.ResponseWriter, r *http.Request) CustomResponse {
	return CustomResponse{writer: w, request: r, data: dict{}}
}
