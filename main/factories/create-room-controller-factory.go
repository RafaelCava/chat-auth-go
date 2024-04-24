package factories

import (
	"github.com/RafaelCava/chat-auth-go/data/usecases"
	"github.com/RafaelCava/chat-auth-go/infra/db/postgres/repositories"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
)

// CreateRoom godoc
// @Summary      Create a new room
// @Description  Create a new room
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param        CreateRoomParams  body usecases.CreateRoomParams false "Data to create one room"
// @Success      201  {object}  models.Room "Room"
// @Failure      400  {object}  protocols.HttpResponse "Bad Request"
// @Failure      404  {object}  protocols.HttpResponse "Not Found"
// @Failure      500  {object}  protocols.HttpResponse "Internal Server Error"
// @Router       /rooms [post]
func NewCreateRoomControllerFactory() protocols.Controller {
	CreateRoomPostgresRepository := repositories.NewCreateRoomPostgresRepository(db_postgres_con)
	usecase := usecases.NewDbCreateRoomUseCase(CreateRoomPostgresRepository)
	controller := controllers.NewCreateRoomController(usecase)
	return controller
}
