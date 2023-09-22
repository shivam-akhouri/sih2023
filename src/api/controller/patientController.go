package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"v0/src/utils"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func CreatePatients(c *gin.Context) {
	ctx := context.Background()
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]interface{})
	json.Unmarshal(body, &jsonmap)
	client := utils.GetClient()
	client.Collection("patient").Add(ctx, jsonmap)
	c.JSON(http.StatusOK, "done")
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "DONE")
}

func DeletePatient(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println(body)
	c.JSON(http.StatusOK, "Done")
}

func getTop3Doctors(hospitals []string) []map[string]interface{} {
	ctx := context.Background()
	client := utils.GetClient()
	doctors := []map[string]interface{}{}
	for i := 0; i < len(hospitals); i++ {
		doc, _ := client.Collection("hospital").Doc(hospitals[i]).Get(ctx)
		results := doc.Data()["Doctors"].([]interface{})
		s := make([]string, len(results))
		for j, v := range results {
			s[j] = fmt.Sprint(v)
		}
		for j, _ := range s {
			data, _ := client.Collection("doctor").Doc(s[j]).Get(ctx)
			doctors = append(doctors, data.Data())
		}
	}
	return doctors
}

func AiAppointement(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]any)
	json.Unmarshal(body, &jsonmap)
	latitude := jsonmap["Latitude"].(float64)
	longitude := jsonmap["Longitude"].(float64)
	// specialization := jsonmap["Specialization"].(string)
	hospitals := getNearbyHospitals(latitude, longitude)
	doctors := getTop3Doctors(hospitals)
	c.JSON(http.StatusOK, doctors)
}

func BookAppointement(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]any)
	json.Unmarshal(body, &jsonmap)
	ctx := context.Background()
	client := utils.GetClient()
	doc := client.Collection("doctor").Doc(jsonmap["Id"].(string))
	data, _ := doc.Get(ctx)
	appointements := data.Data()["Slots"].([]any)
	new_appointements := append(appointements, map[string]interface{}{
		"StartTime": jsonmap["StartTime"],
		"EndTime":   jsonmap["EndTime"],
		"Date":      jsonmap["Date"],
	})
	doc.Update(ctx, []firestore.Update{
		{Path: "Slots", Value: new_appointements},
	})
}
