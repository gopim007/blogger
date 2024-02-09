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
	//Getting database details from env file
	host := os.Getenv("POSTGRESDB_HOST")
	dbname := os.Getenv("POSTGRESDB_DBNAME")
	user := os.Getenv("POSTGRESDB_USER")
	password := os.Getenv("POSTGRESDB_PASSWORD")
	port := os.Getenv("POSTGRESDB_PORT")

	//framing connection string
	connectionString := "host=" + host + " user=" + user + " password=" + password +
		" dbname=" + dbname + " port=" + port

	//setting up new database connection
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	//defer closing database connection
	defer conn.CloseConnection()

	//initialising services
	service := handlers.NewService(conn)

	//creating router
	router := gin.Default()
	postsRouter := router.Group("/post")

	//end points
	postsRouter.POST("/", service.CreatePost)
	postsRouter.GET("/", service.GetAllPosts)
	postsRouter.GET("/:post_id", service.GetPostByID)
	postsRouter.PUT("/:post_id", service.UpdatePostByID)
	postsRouter.DELETE("/:post_id", service.DeletePostByID)

	router.Run()

}
