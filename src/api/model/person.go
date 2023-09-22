package model

type Person struct {
	id            int
	name          string
	mobile        uint
	aadhar        int
	location      map[string]float32
	appointements []map[string]string
}
