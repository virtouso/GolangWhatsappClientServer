package main

import (
	"github.com/virtouso/WhatsappClientServer/ClientApp/app"
	"github.com/virtouso/WhatsappClientServer/ClientApp/repository"
	"github.com/virtouso/WhatsappClientServer/ClientApp/whatsapp"
)

func main() {
	go whatsapp.Init()
	repository.UserRepo = repository.PostgresUserRepository{}
	repository.UserRepo.Init()
	app.StartApplication()

}
