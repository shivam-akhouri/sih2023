package model

import (
	"github.com/gin-gonic/gin"
)

type Doctor struct {
	Id             int
	Name           string
	Specialization []string
	Designation    string
	CabinLocation  map[string]float64
	CabinNumber    string
	Face           string
	Dutyhour       map[string]string
	Slots          map[string]string
}

func (d *Doctor) GetDetails() gin.H {
	return gin.H{
		"id":             d.Id,
		"name":           d.Name,
		"specialization": d.Specialization,
		"designation":    d.Designation,
		"cabinLocation":  d.CabinLocation,
		"cabinNumber":    d.CabinNumber,
		"face":           d.Face,
		"dutyHour":       d.Dutyhour,
		"slots":          d.Slots,
	}
}
