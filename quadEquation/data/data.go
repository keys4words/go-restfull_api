package data

var DB []*QuadraticEquation

type QuadraticEquation struct {
	A          int  `json:"A"`
	B          int  `json:"B"`
	C          int  `json:"C"`
	Nroots     int  `json:"Nroots"`
	Calculated bool `json:"0"`
}
