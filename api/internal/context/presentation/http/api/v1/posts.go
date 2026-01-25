package v1

import (
	"net/http"
	"strconv"

	"app/internal/context/application/usecase/posts"

	"github.com/gin-gonic/gin"
)

type PostsController struct {
	uc *posts.GetListUseCase
}

func NewPostsController(uc *posts.GetListUseCase) *PostsController {
	return &PostsController{
		uc: uc,
	}
}

// ListPosts godoc
//
//	@Summary		List posts
//	@Description	Get a list of posts with pagination
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Page number"	default(1)
//	@Param			limit	query		int	false	"Number of posts per page"	default(9)
//	@Success		200		{object}	posts.ListResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/posts [get]
func (controller *PostsController) ListPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "9"))
	if err != nil || limit < 1 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	response, err := controller.uc.Execute(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list posts"})
		return
	}

	c.JSON(http.StatusOK, response)
}
