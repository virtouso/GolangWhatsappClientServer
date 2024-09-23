package service

import (
	"github.com/virtouso/WhatsappClientServer/ClientApp/basic"
	"github.com/virtouso/WhatsappClientServer/ClientApp/repository"
	"github.com/virtouso/WhatsappClientServer/ClientApp/whatsapp"
)

func SendMessageToAllUsers(message string) basic.MetaResult[string] {
	users := repository.UserRepo.ReadAll()

	for _, user := range users {
		whatsapp.SendMessage(user.AccountId+"@s.whatsapp.net", message)
	}
	return basic.MetaResult[string]{
		Result:       "ok!",
		ErrorMessage: "",
		ResponseCode: 200,
	}
}
