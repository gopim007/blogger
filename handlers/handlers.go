package handlers

import (
	"blogger/constants"
	"blogger/database"
	"blogger/datetime"
	"blogger/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	Repo *database.DB
}

func NewService(repo *database.DB) *Service {
	return &Service{
		Repo: repo,
	}
}

func (serv *Service) CreatePost(c *gin.Context) {
	var post entity.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, constants.INVALID_DATA)
		return
	}

	post.ID = uuid.New().String()
	post.CreatedAt = datetime.GetCurrentTimestampAsInt()
	err := post.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := serv.Repo.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)

}

func (serv *Service) GetAllPosts(c *gin.Context) {
	posts, err := serv.Repo.GetPosts("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)

}
