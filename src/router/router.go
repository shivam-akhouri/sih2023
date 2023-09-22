package router

import (
	"net/http"
	"v0/src/api/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run()
}

func NewRouter() *gin.Engine {
	router := gin.New()
	// TODO: Export grouping login here
	hospitalRoutes := router.Group("/hospital")
	{
		hospitalRoutes.POST("/doctors", controller.GetHospitalDoctors)
		hospitalRoutes.POST("/nearby", controller.GetNearbyHospitals)
		hospitalRoutes.POST("/createHospital", controller.CreateHospital)
	}
	patientRoutes := router.Group("/patient")
	{
		patientRoutes.POST("/createAccount", controller.CreatePatients)
		patientRoutes.POST("/login", controller.Login)
		patientRoutes.POST("/deleteAccount", controller.DeletePatient)
		patientRoutes.POST("/aiAppointement", controller.AiAppointement)
		patientRoutes.POST("/bookAppointement", controller.BookAppointement)
	}
	doctorRoutes := router.Group("/doctor")
	{
		doctorRoutes.GET("/createAccount", controller.CreateDoctor)
		doctorRoutes.GET("/markAttendance", controller.MarkAttendance)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"Status": "Success",
		})
	})
	return router
}
