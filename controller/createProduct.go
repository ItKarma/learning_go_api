package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePrudctController(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"OK": "Ok",
	})

}
