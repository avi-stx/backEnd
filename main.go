package main

import (
	"encoding/json"
	"log"
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
	router.POST("/files", func(c *gin.Context) {
		c.String(http.StatusOK, "file uploaded\n")
	})

	// Download
	router.GET("/files/:name", func(c *gin.Context) {
		fileName := c.Param("name")
		pathName := getRelativePath() + DIR_NAME + fileName
		c.File(pathName)
	})

	// get list of all files
	router.GET("/files", func(c *gin.Context) {
		filesList := readFiles()
		jsonData, err := json.Marshal(filesList)
		if err != nil {
			log.Println(err)
		}
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	// delete a file
	router.DELETE("/files/:name", func(c *gin.Context) {
		fileName := c.Param("name")
		isRemoved := removeFile(fileName)
		if isRemoved {
			c.String(http.StatusOK, "deleted %s\n", fileName)
		}
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
