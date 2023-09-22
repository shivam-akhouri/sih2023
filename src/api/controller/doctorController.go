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
