package factories

import (
	"github.com/RafaelCava/chat-auth-go/data/usecases"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
)

func NewCreateUserControllerFactory() protocols.Controller {
	usecase := usecases.NewDbCreateUserUseCase()
	controller := controllers.NewCreateUserController(usecase)
	return controller
}
