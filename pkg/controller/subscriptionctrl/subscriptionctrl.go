package subscriptionctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCheckoutURL(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"data": "Hi"})
}
