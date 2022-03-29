package utils

import (
	"quadEquation/data"
)

type Msg struct {
	Msg string `json:"msg"`
}

func CalcNumRoots() {
	for _, v := range data.DB {
		// ax2 = 0
		if v.B == 0 && v.C == 0 {
			v.Nroots = 1
			v.Calculated = true
			return
		}
		// ax2 + bx = 0
		if v.A != 0 && v.B != 0 && v.C == 0 {
			v.Nroots = 2
			v.Calculated = true
			return
		}
		// ax2 + c = 0
		// solution depends from sign of a & c
		if v.B == 0 && v.A != 0 && v.C != 0 {
			if (v.A > 0 && v.C > 0) || (v.A < 0 && v.C < 0) {
				v.Nroots = 0
				v.Calculated = true
			} else {
				v.Nroots = 2
				v.Calculated = true
			}
			return
		}
		// bx + c = 0
		if v.A == 0 && v.B != 0 && v.C != 0 {
			v.Nroots = 1
			v.Calculated = true
			return
		}
		// not equation c = 0
		if v.A == 0 && v.B == 0 && v.C != 0 {
			v.Nroots = 0
			v.Calculated = true
			return
		}
		// full quadractic equation
		D := v.B*v.B - 4*v.A*v.C
		if D > 0 {
			v.Nroots = 2
			v.Calculated = true
		} else if D == 0 {
			v.Nroots = 1
			v.Calculated = true
		} else {
			v.Nroots = 0
			v.Calculated = true
		}
	}
	return
}
