package main

import (
	v1 "jwk-token/handler/v1"
	"jwk-token/initializer"
	"jwk-token/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
}
func main() {
	// fmt.Println("helo")
	r := gin.Default()
	r.POST("/signup", v1.SignUp)
	r.POST("/login", v1.Login)
	r.GET("/validate", middleware.Authorize, v1.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
