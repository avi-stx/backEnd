package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defineRoutes(router *gin.Engine) {

	//Upload
	router.PUT("/files", func(c *gin.Context) {
		c.String(http.StatusOK, "file uploaded\n")
	})

	// Download
	router.GET("/files/:name", func(c *gin.Context) {
		fileName := c.Param("name")
		c.String(http.StatusOK, "downloading %s\n", fileName)
	})

	// get list of all files
	router.GET("/files", func(c *gin.Context) {
		c.String(http.StatusOK, "got all files \n")

	})

	// delete a file
	router.DELETE("/files/:name", func(c *gin.Context) {
		fileName := c.Param("name")

		c.String(http.StatusOK, "deleted %s\n", fileName)

	})

}

func initServer() {

	router := gin.Default()

	defineRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func main() {
	initServer()
}
