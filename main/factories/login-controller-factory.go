package factories

import (
	"os"

	"github.com/RafaelCava/chat-auth-go/data/usecases"
	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/RafaelCava/chat-auth-go/infra/db/postgres/repositories"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
)

// AuthLogin     godoc
// @Summary      Authenticate a user
// @Description  Authenticate a user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        AuthRequest  body usecases.AuthRequest false "Data to authenticate one user"
// @Success      201  {object}  models.TokenPair "TokenPair"
// @Failure      400  {object}  protocols.HttpResponse "Bad Request"
// @Failure      404  {object}  protocols.HttpResponse "Not Found"
// @Failure      500  {object}  protocols.HttpResponse "Internal Server Error"
// @Router       /auth/login [post]
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
