package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"v0/src/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"google.golang.org/api/iterator"
)

func GetNearbyHospitals(c *gin.Context) {
	client := utils.GetClient()
	ctx := context.Background()
	hospitals := client.Collection("hospital").Documents(ctx)
	results := []map[string]interface{}{}
	defer hospitals.Stop()
	for {
		doc, err := hospitals.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		results = append(results, doc.Data())
	}
	c.JSON(http.StatusOK, map[string]any{
		"hospital": results,
	})
}

func getNearbyHospitals(latitude float64, longitude float64) []string {
	client := utils.GetClient()
	ctx := context.Background()
	hospitals := client.Collection("hospital").Documents(ctx)
	results := []string{}
	defer hospitals.Stop()
	for {
		doc, err := hospitals.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		results = append(results, doc.Ref.ID)
	}
	return results
}

func GetHospitalDoctors(c *gin.Context) {
	ctx := context.Background()
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]any)
	json.Unmarshal(body, &jsonmap)
	client := utils.GetClient()
	doc := client.Collection("hospital").Doc(jsonmap["Id"].(string))
	data, _ := doc.Get(ctx)
	doctorIds := data.Data()["Doctors"].([]interface{})
	s := make([]string, len(doctorIds))
	for i, v := range doctorIds {
		s[i] = fmt.Sprint(v)
	}
	// fmt.Println(doctorIds)
	ids := client.Collection("doctor").DocumentRefs(ctx)
	doctors := []map[string]interface{}{}
	for {
		id, err := ids.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		if slices.Contains(s, id.ID) {
			data, _ := client.Collection("doctor").Doc(id.ID).Get(ctx)
			fmt.Println(data.Data())
			doctors = append(doctors, data.Data())
		}
	}
	c.JSON(http.StatusOK, map[string]any{
		"doctors": doctors,
	})
}

func CreateHospital(c *gin.Context) {
	ctx := context.Background()
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]interface{})
	json.Unmarshal(body, &jsonmap)
	client := utils.GetClient()
	client.Collection("hospital").Doc(jsonmap["Id"].(string)).Set(ctx, jsonmap)
	c.JSON(http.StatusOK, map[string]string{
		"Status": "Success",
	})
}
