package emailctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/email"
)

func SendTestEmail(c *gin.Context) {
	email.SendTestEmail()

	c.JSON(http.StatusOK, gin.H{
		"message": "sent",
	})
}
