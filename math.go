package pgl

func DEG_TO_RAD(x float32) float32 {
	return (x) * PI_DIV_180
}

func RAD_TO_DEG(x float32) float32 {
	return (x) * INV_PI_DIV_180
}
