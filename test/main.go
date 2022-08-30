package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		var data map[string]string
		if err := c.BindJSON(&data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	})
	router.Run(":8080")
}
