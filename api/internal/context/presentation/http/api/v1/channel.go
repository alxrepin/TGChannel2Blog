package v1

import (
	"net/http"

	"app/internal/context/application/usecase/channel"

	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	uc *channel.GetChannelUseCase
}

func NewChannelController(uc *channel.GetChannelUseCase) *ChannelController {
	return &ChannelController{
		uc: uc,
	}
}

// GetChannel godoc
//
//	@Summary		Get channel information
//	@Description	Get information about the channel
//	@Tags			channel
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	channel.ChannelResponse
//	@Failure		500	{object}	map[string]string
//	@Router			/channel [get]
func (controller *ChannelController) GetChannel(c *gin.Context) {
	response, err := controller.uc.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get channel"})
		return
	}

	c.JSON(http.StatusOK, response)
}
