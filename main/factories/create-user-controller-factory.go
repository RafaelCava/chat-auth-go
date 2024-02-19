package factories

import (
	"github.com/RafaelCava/chat-auth-go/data/usecases"
	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/RafaelCava/chat-auth-go/infra/db/postgres/repositories"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
)

func NewCreateUserControllerFactory() protocols.Controller {
	validateIfUserExistsRepository := repositories.NewValidateIfUserExistsPostgresRepository(db_postgres_con)
	hasher := adapters.NewHasherAdapter()
	createUserRepository := repositories.NewCreateUserPostgresRepository(db_postgres_con)
	usecase := usecases.NewDbCreateUserUseCase(validateIfUserExistsRepository, hasher, createUserRepository)
	controller := controllers.NewCreateUserController(usecase)
	return controller
}
