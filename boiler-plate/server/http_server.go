package server

import (
	"example/boiler-plate/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type HttpServer struct {
	Router *gin.Engine
}

func Init(db *gorm.DB) (*HttpServer, error) {
	server := &HttpServer{}
	router := gin.Default()
	server.Router = router
	redisClient := InitRedisCache()

	// create postgres db connection database.DB
	dbStore := database.NewDb(db)
	setUpRoutes(router, dbStore, redisClient)
	return server, nil
}

func (s *HttpServer) Start() error {
	s.Router.Run(":8081")
	return nil
}

func InitRedisCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
