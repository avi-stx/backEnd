package utils

import (
	"Backend_side/config"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {

	router := gin.Default()
	router.Use(CORSMiddleware())
	defineRoutes(router)
	return router
}

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

// defineRoutes set the routes for the engine
func defineRoutes(router *gin.Engine) {
	//Upload
	router.POST("/files", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		converted := io.Reader(file)
		filename := header.Filename
		if err != nil {
			log.Fatal(err)
		}

		isFileUploaded := SaveFileHandler(&converted, filename)
		if isFileUploaded {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "upload succeeded"})
		} else {
			c.String(http.StatusBadRequest, "internal server error")
		}
	})

	// Download
	router.GET("/files/:name", func(c *gin.Context) {
		fileName := c.Param("name")
		pathName := config.AppConfig.TargetFolder + fileName
		c.File(pathName)
	})

	// get list of all files
	router.GET("/files", func(c *gin.Context) {
		filesList := ReadFiles()
		jsonData, err := json.Marshal(filesList)
		if err != nil {
			log.Println(err)
		}
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	// delete a file
	router.DELETE("/files/:name", func(c *gin.Context) {
		fileName := c.Param("name")
		isRemoved := RemoveFile(fileName)
		if isRemoved {
			c.String(http.StatusOK, "deleted %s\n", fileName)
		}
	})

}
