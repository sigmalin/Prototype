package query

import (
	"config"
	"context"
	"net/http"
	"response"
	"response/code"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	db "connect/mongo"

	"api/test/query/allusers"
)

type handler struct {
	mainDB *mongo.Database
}

func newHandler() *handler {
	return &handler{
		mainDB: db.GetDB(config.DATABASE_TABLE),
	}
}

func (h *handler) newResponse() *response.Body {
	return &response.Body{Code: code.SUCCESS, Message: ""}
}

func (h *handler) send(c *gin.Context, resp *response.Body) {
	if resp.Code == code.SUCCESS {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, resp)
	}
}

// @Summary Query All users
// @Tags test
// @version 1.0
// @produce application/json
// @Param start formData string true "Query start"
// @Param count formData string true "Query count"
// @Success 200 {object} response.Body{data=[]allusers.Result} "Success"
// @Router /test/query/allusers [post]
func (h *handler) allUsers(c *gin.Context) {

	start := c.PostForm("start")
	count := c.PostForm("count")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DATABASE_TIMEOUT)
	defer cancel()

	res := h.newResponse()

	args := allusers.NewArguments(
		h.mainDB,
		ctx,
		h.parseValue(start),
		h.parseValue(count))

	allusers.Handle(args, res)

	h.send(c, res)
}

func (h *handler) parseValue(data string) int64 {
	value, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		value = 0
	}

	if value < 0 {
		value = 0
	}

	return value
}
