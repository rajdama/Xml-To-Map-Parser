package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", health)         // /ping endpint provide the info about the health of server
	router.POST("/xml/yaml", xmlToYaml) // /xml/yaml endpoint convert the provided xml into yaml
	router.GET("/xml/isValid", isXml)   //  /xml/isValid used to check if a given xml data is valid or not

	return router
}
