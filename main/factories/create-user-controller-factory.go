package factories

import (
	"github.com/RafaelCava/chat-auth-go/data/usecases"
	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/RafaelCava/chat-auth-go/infra/db/postgres/repositories"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
)

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        CreateUserParams  body usecases.CreateUserParams false "Data to create one user"
// @Success      201  {object}  models.User "User"
// @Failure      400  {object}  protocols.HttpResponse "Bad Request"
// @Failure      404  {object}  protocols.HttpResponse "Not Found"
// @Failure      500  {object}  protocols.HttpResponse "Internal Server Error"
// @Router       /users [post]
func NewCreateUserControllerFactory() protocols.Controller {
	validateIfUserExistsRepository := repositories.NewValidateIfUserExistsPostgresRepository(db_postgres_con)
	hasher := adapters.NewHasherAdapter()
	createUserRepository := repositories.NewCreateUserPostgresRepository(db_postgres_con)
	usecase := usecases.NewDbCreateUserUseCase(validateIfUserExistsRepository, hasher, createUserRepository)
	controller := controllers.NewCreateUserController(usecase)
	return controller
}
