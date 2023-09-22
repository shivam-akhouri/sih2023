package model

import (
	"math"

	"github.com/gin-gonic/gin"
)

type Hospital struct {
	Id            int
	Name          string
	Location      map[string]float64
	Doctors       []int
	Address       string
	Ambulance     int
	BedsAvailable uint
	EmergencyBeds uint
}

func (h *Hospital) GetDetails() gin.H {
	return gin.H{
		"id":            h.Id,
		"name":          h.Name,
		"location":      h.Location,
		"doctors":       h.Doctors,
		"address":       h.Address,
		"ambulance":     h.Ambulance,
		"bedsAvailable": h.BedsAvailable,
		"emergencyBeds": h.EmergencyBeds,
	}
}

func (h *Hospital) calcDistance(latitude float64, longitude float64) float64 {
	return math.Sqrt(math.Pow(h.Location["latitude"]-latitude, 2) + math.Pow(h.Location["longitude"]-longitude, 2))
}
