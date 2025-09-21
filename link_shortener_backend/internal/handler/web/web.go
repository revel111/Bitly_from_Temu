package web

import (
	"errors"
	"fmt"
	countererror "linkShortener/internal/domain/counter"
	domain "linkShortener/internal/domain/link"
	"linkShortener/internal/usecase/counter"
	"linkShortener/internal/usecase/link"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LinkController struct {
	getCounterUseCase counter.GetCounterUseCase
	createLinkUseCase link.CreateLinkUseCase
	getLinkUseCase    link.GetLinkUseCase
}

func NewLinkController(createLinkUseCase link.CreateLinkUseCase, getLinkUseCase link.GetLinkUseCase, getCounterUseCase counter.GetCounterUseCase) *LinkController {
	return &LinkController{
		createLinkUseCase: createLinkUseCase,
		getLinkUseCase:    getLinkUseCase,
		getCounterUseCase: getCounterUseCase,
	}
}

func (lc *LinkController) CreateLink(c *gin.Context) {
	var req domain.CreateLinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	link, err := lc.createLinkUseCase.Execute(req.Url)
	if err != nil {
		status := http.StatusInternalServerError
		msg := err.Error()

		if errors.Is(err, domain.ErrInvalidURL) {
			status = http.StatusBadRequest
			msg = domain.ErrInvalidURL.Error()
		} else if errors.Is(err, domain.ErrExists) {
			status = http.StatusConflict
			msg = domain.ErrExists.Error()
		} else if errors.Is(err, domain.ErrServiceUnavailable) {
			status = http.StatusServiceUnavailable
			msg = domain.ErrServiceUnavailable.Error()
		}

		response := gin.H{
			"Message":   msg,
			"Timestamp": time.Now(),
		}
		if link != "" {
			response["url"] = link
		}
		c.JSON(status, response)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created link",
		"url":     fmt.Sprintf("/api/v1/links/%s", link),
	})
}

func (lc *LinkController) Forward(c *gin.Context) {
	shortCode := c.Param("code")

	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link parameter is required"})
		return
	}

	link, err := lc.getLinkUseCase.Execute(shortCode)
	if err != nil {
		status := http.StatusInternalServerError
		msg := err.Error()

		if errors.Is(err, domain.ErrNotFound) {
			status = http.StatusNotFound
			msg = domain.ErrNotFound.Error()
		} else if errors.Is(err, domain.ErrServiceUnavailable) {
			status = http.StatusServiceUnavailable
			msg = domain.ErrServiceUnavailable.Error()
		}

		c.JSON(status, gin.H{
			"Message":   msg,
			"Timestamp": time.Now(),
		})
		return
	}

	c.Redirect(http.StatusFound, link)
}

func (lc *LinkController) GetCount(c *gin.Context) {
	shortCode := c.Param("code")

	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link parameter is required"})
		return
	}

	count, err := lc.getCounterUseCase.Execute(shortCode)
	if err != nil {
		status := http.StatusInternalServerError
		msg := err.Error()

		if errors.Is(err, countererror.ErrNotFound) {
			status = http.StatusNotFound
			msg = countererror.ErrNotFound.Error()
		}

		c.JSON(status, gin.H{
			"Message":   msg,
			"Timestamp": time.Now(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
