package model

type Attendance struct {
	Id       int
	DateTime string
	Location map[string]float64
}
