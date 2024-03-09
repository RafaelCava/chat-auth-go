// main/main.go - Camada Principal
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RafaelCava/chat-auth-go/main/factories"
	"github.com/RafaelCava/chat-auth-go/main/routes"
)

//	@title			Chat Auth Api
//	@version		0.0.1
//	@description	API para gerenciamento de chats, grupos e autenticação de usuários

//	@contact.name	Rafael Cavalcante
//	@contact.url	https://github.com/RafaelCava
//	@contact.email	rafael.software-developer@outlook.com

//	@host		localhost:3000
//	@BasePath	/api
// 	@securityDefinitions.apikey Bearer
// 	@in header
// 	@name Authorization
//  @description Authorization header using the Bearer scheme.

func main() {
	go func() {
		errPostgresCon := factories.NewDatabasePostgresOpenConnection()
		if errPostgresCon != nil {
			panic("Falha ao conectar ao banco de dados - Postgres")
		}
		errMongoCon := factories.NewDatabaseMongoOpenConnection()
		if errMongoCon != nil {
			panic("Falha ao conectar ao banco de dados - MongoDB")
		}
		// errRedisCon := factories.NewDatabaseRedisOpenConnection()
		// if errRedisCon != nil {
		// 	panic("Falha ao conectar ao banco de dados - Redis")
		// }
		if err := routes.Run(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Stopping server...")

	if err := routes.ShutDown(ctx); err != nil {
		panic(err)
	}
	if err := factories.NewCloseDatabasePostgresConnection(); err != nil {
		panic(err)
	}
	if err := factories.NewCloseDatabaseMongoConnection(); err != nil {
		panic(err)
	}
	// if err := factories.NewCloseDatabaseRedisConnection(); err != nil {
	// 	panic(err)
	// }

	fmt.Println("Server stopped successfully!")
}
