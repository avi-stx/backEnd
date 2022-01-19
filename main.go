package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

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
		// err := os.Remove("testFile.txt")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		c.String(http.StatusOK, "deleted %s\n", fileName)

	})

}

func initServer() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	defineRoutes(router)
	router.Run("localhost:8080")

}

func main() {
	initServer()
}
