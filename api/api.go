package api

import (
	"booking/api/handler"
	"booking/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handler.Handler, cfg config.Config) {
	r.Use(customCORSMiddleware())

	// Auth Endpoints
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	//Role Endpoints
	r.POST("/createrole", h.CreateRole)
	r.PUT("/updaterole", h.UpdateRole)
	r.GET("/roles", h.GetRolesList)
	r.GET("/role/:id", h.GetRolesByIDHandler)
	r.DELETE("/deleterole/:id", h.DeleteRole)

	// Users Endpoints
	r.POST("/createuser", h.CreateUser)
	r.PUT("/updateuser", h.UpdateUser)
	r.GET("/users", h.GetUsersList)
	r.GET("/user/:id", h.GetUsersByIDHandler)
	r.DELETE("/deleteuser/:id", h.DeleteUser)

	// Doctors Endpoints
	r.POST("/createdoctor", h.CreateDoctor)
	r.PUT("/updatedoctor", h.UpdateDoctor)
	r.GET("/doctors", h.GetDoctorsList)
	r.GET("/doctor/:id", h.GetDoctorsByIDHandler)
	r.DELETE("/deletedoctor/:id", h.DeleteDoctor)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSF-TOKEN, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
