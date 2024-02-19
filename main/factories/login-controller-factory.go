package factories

import (
	"os"

	"github.com/RafaelCava/chat-auth-go/data/usecases"
	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/RafaelCava/chat-auth-go/infra/db/postgres/repositories"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
)

func NewLoginControllerFactory() protocols.Controller {
	secretKey := os.Getenv("SECRET_TOKEN")
	issuerKey := os.Getenv("ISS_TOKEN")
	findUserByEmailRepository := repositories.NewFindUserByEmailPostgresRepository(db_postgres_con)
	hasherCompare := adapters.NewHasherAdapter()
	encrypter := adapters.NewEncrypterAdapter(secretKey, issuerKey)
	authUseCase := usecases.NewDbLogin(findUserByEmailRepository, hasherCompare, encrypter)
	loginController := controllers.NewLoginController(authUseCase)
	return loginController
}
