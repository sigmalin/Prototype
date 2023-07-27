package modify

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

	"api/test/modify/bank"
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

// @Summary Modify user's bank data
// @Tags test
// @version 1.0
// @produce application/json
// @Param userID formData string true "userID"
// @Param Coin formData string true "Coin"
// @Param Faith formData string true "Faith"
// @Param Gems formData string true "Gems"
// @Param Treasures formData string true "Treasures"
// @Success 200 {object} response.Body{data=bank.Result} "Success"
// @Router /test/modify/bank [post]
func (h *handler) bank(c *gin.Context) {

	id := c.PostForm("userID")
	coin := c.PostForm("Coin")
	faith := c.PostForm("Faith")
	gems := c.PostForm("Gems")
	treasure := c.PostForm("Treasures")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DATABASE_TIMEOUT)
	defer cancel()

	resp := h.newResponse()

	args := bank.NewArguments(
		h.mainDB,
		ctx,
		id,
		h.parseValue(coin), h.parseValue(faith),
		h.parseValue(gems), h.parseValue(treasure))

	bank.Handle(args, resp)

	h.send(c, resp)
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
