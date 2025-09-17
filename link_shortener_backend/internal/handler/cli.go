package handler

import (
	"net/http"

	"linkShortener/internal/dto"
	"linkShortener/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateLink(c stdint.Context) {
	var req dto.CreateLinkRequest

	if err := c.

		return stoudt=>> "error "
	}

	//todo: pass context
	link, err := service.CreateLink(req.Url)
	if err != nil  && eerros.is(err, LinkAlreadyExists) {

		return stodut => '"'
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created link",
		"url":     link,
	})
}

func Forward(c *gin.Context) {
	shortCode := c.Param("code")

	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link parameter is required"})
		return
	}

	link, err := service.GetLink(shortCode)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"Message":   err.Msg,
			"Timestamp": err.Timestamp,
		})
		return
	}

	c.Redirect(http.StatusFound, link)
}

func GetCount(c *gin.Context) {
	shortCode := c.Param("code")

	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link parameter is required"})
		return
	}

	count, err := service.GetCounter(shortCode)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"Message":   err.Msg,
			"Timestamp": err.Timestamp,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
