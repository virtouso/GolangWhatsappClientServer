package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/virtouso/WhatsappClientServer/ClientApp/controller"
	MiddleWare "github.com/virtouso/WhatsappClientServer/ClientApp/middleware"
)

func mapUrls() {
	mapAdminEndpoints()
	mapUserEndpoints()

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	router.Use(MiddleWare.LoggingMiddleware())
}

func mapAdminEndpoints() {

	router.POST("/admin/sendMessage", MiddleWare.AuthorizeAdminSimpleSecret, controller.SendMessageToAllUsers)

}

func mapUserEndpoints() {
	router.GET("/ping", controller.Ping)
	router.GET("/api/v1/user/subscribe", controller.SubscribeWithUsername)

}
