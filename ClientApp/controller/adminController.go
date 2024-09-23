package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/virtouso/WhatsappClientServer/ClientApp/basic"
	"github.com/virtouso/WhatsappClientServer/ClientApp/model/dto/req"
	"github.com/virtouso/WhatsappClientServer/ClientApp/service"
	"net/http"
)

func SendMessageToAllUsers(c *gin.Context) {
	var request req.MessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, basic.MetaResult[string]{
			Result:       "",
			ErrorMessage: "invalid input",
			ResponseCode: 400,
		})
		return
	}
	result := service.SendMessageToAllUsers(request.Data)
	c.JSON(result.ResponseCode, result)

}
