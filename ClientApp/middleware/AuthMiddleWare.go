package MiddleWare

import (
	"github.com/gin-gonic/gin"
	"github.com/virtouso/WhatsappClientServer/ClientApp/basic"
	"github.com/virtouso/WhatsappClientServer/ClientApp/shared"
	"os"
)

func AuthorizeAdminSimpleSecret(c *gin.Context) {
	token := c.GetHeader(shared.AdminSecretKey)

	if len(token) <= 0 || token != os.Getenv(shared.AdminSecretKey) {
		c.JSON(403, basic.MetaResult[string]{
			Result:       "",
			ErrorMessage: "invalid admin secret",
			ResponseCode: 403,
		})
		c.Abort()
		return
	}

}
