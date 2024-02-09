package main

import (
	"blogger/database"
	"blogger/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load environment variables")
	}
}

func main() {
	host := os.Getenv("POSTGRESDB_HOST")
	dbname := os.Getenv("POSTGRESDB_DBNAME")
	user := os.Getenv("POSTGRESDB_USER")
	password := os.Getenv("POSTGRESDB_PASSWORD")
	port := os.Getenv("POSTGRESDB_PORT")

	connectionString := "host=" + host + " user=" + user + " password=" + password +
		" dbname=" + dbname + " port=" + port

	conn, err := database.NewConnection(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	//defer closing database connection
	defer conn.CloseConnection()

	service := handlers.NewService(conn)

	router := gin.Default()
	postsRouter := router.Group("/post")

	postsRouter.POST("/", service.CreatePost)
	postsRouter.GET("/", service.GetAllPosts)
	postsRouter.GET("/:post_id", service.GetPostByID)

	router.Run()

}
