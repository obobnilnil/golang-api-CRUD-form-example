package handler

import (
	"mars_git/model"
	"mars_git/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	LeaderCreateFormHandler(c *gin.Context)
	SubmitFormHandler(c *gin.Context)
}

type handlerAdapter struct {
	s service.ServicePort
}

func NewHanerhandlerAdapter(s service.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) LeaderCreateFormHandler(c *gin.Context) {
	var leaderCreateForm model.LeaderCreateForm
	if err := c.ShouldBindJSON(&leaderCreateForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	err := h.s.LeaderCreateFormService(leaderCreateForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "OK", "message": "Form created successfully."})
}

func (h *handlerAdapter) SubmitFormHandler(c *gin.Context) {
	// fmt.Println(c)
	var submitDetail model.SubmitForm
	if err := c.ShouldBindJSON(&submitDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	err := h.s.SubmitFormService(submitDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Form submitted successfully"})
}
