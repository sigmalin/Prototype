package modify

import (
	"config"
	"context"
	"net/http"
	"response"
	"response/code"
	"strconv"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"

	"api/test/modify/bank"
)

type handler struct {
	mainDB *sql.DB
}

func newHandler() *handler {
	return &handler{
		mainDB: db.GetDB(config.SQL_DATABASE),
	}
}

func (h *handler) newResponse() *response.Body {
	return &response.Body{Code: code.SUCCESS, Message: ""}
}

func (h *handler) send(c *gin.Context, res *response.Body) {
	if res.Code == code.SUCCESS {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, res)
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
// @Success 200 {object} response.Body{data=bankData.Content} "Success"
// @Router /test/modify/bank [post]
func (h *handler) bank(c *gin.Context) {

	id := c.PostForm("userID")
	coin := c.PostForm("Coin")
	faith := c.PostForm("Faith")
	gems := c.PostForm("Gems")
	treasure := c.PostForm("Treasures")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	res := h.newResponse()

	args := bank.NewArguments(
		h.mainDB,
		ctx,
		id,
		h.parseValue(coin), h.parseValue(faith),
		h.parseValue(gems), h.parseValue(treasure))

	bank.Handle(args, res)

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
