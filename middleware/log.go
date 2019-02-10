package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogRequest(c *gin.Context){
	logrus.WithField("URL:", c.Request.URL).Info("Receive new request")
}
