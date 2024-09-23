package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/virtouso/WhatsappClientServer/ClientApp/basic"
	"github.com/virtouso/WhatsappClientServer/ClientApp/model/dto/req"
	"github.com/virtouso/WhatsappClientServer/ClientApp/service"
	"net/http"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func SubscribeWithUsername(c *gin.Context) {

	var request req.SubscribeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.MetaResult[string]{
			Result:       "",
			ErrorMessage: "invalid input",
			ResponseCode: 400,
		})
		return
	}

	result := service.Subscribe(request)
	c.JSON(result.ResponseCode, result)

}
