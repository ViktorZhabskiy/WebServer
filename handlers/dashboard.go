package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DashboardHandler struct {

}

func NewDashboard() func(*gin.Context){
	h := DashboardHandler{}
	return h.handle
}

func(d *DashboardHandler) handle(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"data": "some data"})
}