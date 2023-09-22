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
		hospitalRoutes.GET("/doctors", controller.GetHospitalDoctors)
		hospitalRoutes.GET("/nearby", controller.GetNearbyHospitals)
		hospitalRoutes.GET("/createHospital", controller.CreateHospital)
	}
	patientRoutes := router.Group("/patient")
	{
		patientRoutes.GET("/createAccount", controller.CreatePatients)
		patientRoutes.GET("/login", controller.Login)
		patientRoutes.GET("/deleteAccount", controller.DeletePatient)
		patientRoutes.GET("/aiAppointement", controller.AiAppointement)
		patientRoutes.GET("/bookAppointement", controller.BookAppointement)
	}
	doctorRoutes := router.Group("/doctor")
	{
		doctorRoutes.GET("/createAccount", controller.CreateDoctor)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"Status": "Success",
		})
	})
	return router
}
