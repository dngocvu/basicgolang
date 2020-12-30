package triangle

import "math"

func AreaTri(adge1, adge2, angle float64 ) float64{
	parimeterTriangle := (adge1*adge2*math.Sin(angle))/2
	return parimeterTriangle
}
func ParimeterTri(adge1, edge2, edge3 float64) float64 {
	parimeterTriangle := adge1 + edge2 + edge3 
	return parimeterTriangle
}