package handler

import (
	"linkShortener/internal/dto"
	"linkShortener/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	var req dto.CreateLinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	link, err := service.CreateLink(req.Url)
	if err != nil {
		response := gin.H{
			"Message":   err.Msg,
			"Timestamp": err.Timestamp,
		}
		if link != "" {
			response["url"] = link
		}
		c.JSON(err.Code, response)
		return
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
