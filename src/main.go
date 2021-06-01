package main

import (
	gin "github.com/gin-gonic/gin"
	"log"
	"post-service/infrastructure/cassandra_config"
	"post-service/interactor"
)

func main() {
	cassandraSession, err := cassandra_config.NewCassandraSession()
	if err != nil {
		log.Println(err)
	}

	i := interactor.NewInteractor(cassandraSession)

	handler := i.NewAppHandler()

	g := gin.Default()
	g.GET("/ping", handler.AddPost)

	g.Run("localhost:8083")
}
