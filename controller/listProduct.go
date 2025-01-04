package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPrudctController(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"OK": "Ok",
	})

}
