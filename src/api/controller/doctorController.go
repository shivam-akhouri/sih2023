package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"v0/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateDoctor(c *gin.Context) {
	ctx := context.Background()
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]interface{})
	json.Unmarshal(body, &jsonmap)
	client := utils.GetClient()
	client.Collection("doctor").Doc(jsonmap["Id"].(string)).Set(ctx, jsonmap)
	c.JSON(http.StatusOK, "done")
}

func MarkAttendance(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	jsonmap := make(map[string]interface{})
	json.Unmarshal(body, &jsonmap)
	if jsonmap["payload"].(string) == "602188916031" || jsonmap["payload"] == "197205543529" || jsonmap["payload"] == "2011897164215" {
		c.JSON(http.StatusOK, 1)
	} else {
		c.JSON(http.StatusOK, 0)
	}
}
