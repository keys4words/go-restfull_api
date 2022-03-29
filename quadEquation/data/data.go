package data

var DB []*QuadraticEquation

type QuadraticEquation struct {
	A          int  `json:"a"`
	B          int  `json:"b"`
	C          int  `json:"c"`
	Nroots     int  `json:"nroots"`
	Calculated bool `json:"0"`
}
