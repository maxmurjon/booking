package api

import (
	"booking/api/handler"
	"booking/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "booking/api/docs"
)

func SetUpAPI(r *gin.Engine, h handler.Handler, cfg config.Config) {
	r.Use(customCORSMiddleware())
	r.Static("/swagger", "./docs")


	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	protected := r.Group("/")
	protected.Use(h.AuthMiddleware())
	{
		// Users
		protected.POST("/createuser", h.CreateUser)
		protected.PUT("/updateuser", h.UpdateUser)
		protected.GET("/users", h.GetUsersList)
		protected.GET("/user/:id", h.GetUsersByIDHandler)
		protected.DELETE("/deleteuser/:id", h.DeleteUser)

		// Appointments
		protected.POST("/createappointment", h.CreateAppointment)
		protected.PUT("/updateappointment", h.UpdateAppointment)
		protected.GET("/appointments", h.GetAppointmentsList)
		protected.GET("/appointment/:id", h.GetAppointmentsByIDHandler)
		protected.DELETE("/deleteappointment/:id", h.DeleteAppointment)
	}

	admin := r.Group("/")
	admin.Use(h.AuthMiddleware(), h.RoleMiddleware("admin", "doctor"))
	{
		// Roles
		admin.POST("/createrole", h.CreateRole)
		admin.PUT("/updaterole", h.UpdateRole)
		admin.GET("/roles", h.GetRolesList)
		admin.GET("/role/:id", h.GetRolesByIDHandler)
		admin.DELETE("/deleterole/:id", h.DeleteRole)

		// Doctors
		admin.POST("/createdoctor", h.CreateDoctor)
		admin.PUT("/updatedoctor", h.UpdateDoctor)
		admin.GET("/doctors", h.GetDoctorsList)
		admin.GET("/doctor/:id", h.GetDoctorsByIDHandler)
		admin.DELETE("/deletedoctor/:id", h.DeleteDoctor)
	}

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
