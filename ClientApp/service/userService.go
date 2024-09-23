package service

import (
	"github.com/virtouso/WhatsappClientServer/ClientApp/basic"
	"github.com/virtouso/WhatsappClientServer/ClientApp/model/domain"
	"github.com/virtouso/WhatsappClientServer/ClientApp/model/dto/req"
	"github.com/virtouso/WhatsappClientServer/ClientApp/repository"
)

func Subscribe(req req.SubscribeRequest) basic.MetaResult[string] {
	_, _ = repository.UserRepo.Create(&domain.User{

		Name:      req.UserId,
		AccountId: req.AccountId,
	})

	return basic.MetaResult[string]{
		Result:       "ok!",
		ErrorMessage: "",
		ResponseCode: 200,
	}

}
