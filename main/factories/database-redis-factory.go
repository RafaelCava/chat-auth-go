package factories

import (
	"os"

	"github.com/go-redis/redis"
)

var db_redis_con *redis.Client

func NewDatabaseRedisOpenConnection() error {
	redisHost := os.Getenv("DB_RD_HOST")
	redisPass := os.Getenv("DB_RD_PASSWORD")
	redisPort := os.Getenv("DB_RD_PORT")
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	db_redis_con = client
	return nil
}

func NewCloseDatabaseRedisConnection() error {
	return db_redis_con.Close()
}
