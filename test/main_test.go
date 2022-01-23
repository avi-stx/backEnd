package test

import (
	"Backend_side/config"
	"Backend_side/utils"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func init() {
	err := config.ReadConfig("../config/test_server.json")
	if err != nil {
		panic(err)
	}

}

func setup() (*gin.Engine, *httptest.ResponseRecorder) {

	router := utils.InitServer()
	w := httptest.NewRecorder()

	return router, w
}
