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

// NewService - initialising service
func NewService(repo *database.DB) *Service {
	return &Service{
		Repo: repo,
	}
}

// CreatePost - create post service
func (serv *Service) CreatePost(c *gin.Context) {
	var post entity.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, constants.INVALID_DATA)
		return
	}

	post.ID = uuid.New().String()
	post.CreatedAt = datetime.GetCurrentTimestampAsInt()
	err := post.Validate(constants.CREATE)
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

// GetAllPosts- get all posts service
func (serv *Service) GetAllPosts(c *gin.Context) {
	posts, err := serv.Repo.GetPosts("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)

}

// GetPostByID - get posts by id service
func (serv *Service) GetPostByID(c *gin.Context) {
	postId := c.Param("post_id")
	if postId == "" {
		c.JSON(http.StatusBadRequest, constants.INVALID_DATA)
		return
	}

	result, err := serv.Repo.GetPosts(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, constants.RECORD_NOT_FOUND)
		return
	}

	c.JSON(http.StatusOK, result)

}

// UpdatePostByID - update posts by id service
func (serv *Service) UpdatePostByID(c *gin.Context) {

	var post entity.Post
	postId := c.Param("post_id")

	if postId == "" {
		c.JSON(http.StatusBadRequest, constants.INVALID_DATA)
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, constants.INVALID_DATA)
		return
	}

	if postId != post.ID {
		c.JSON(http.StatusBadRequest, constants.INVALID_DATA)
		return
	}

	post.UpdatedAt = datetime.GetCurrentTimestampAsInt()
	err := post.Validate(constants.UPDATE)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedPost, err := serv.Repo.UpdatePostByID(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedPost)

}

// DeletePostByID - delete posts by id service
func (serv *Service) DeletePostByID(c *gin.Context) {
	postId := c.Param("post_id")

	if postId == "" {
		c.JSON(http.StatusBadRequest, constants.INVALID_DATA)
		return
	}

	deletedPost, err := serv.Repo.DeletePostByID(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.INVALID_DATA)
		return
	}
	c.JSON(http.StatusOK, deletedPost)

}
