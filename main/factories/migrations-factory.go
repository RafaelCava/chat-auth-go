package factories

import "github.com/RafaelCava/chat-auth-go/domain/models"

func newMigratePostgresModels() {
	// Migrar modelos para o banco de dados
	db_postgres_con.AutoMigrate(&models.User{})
}
