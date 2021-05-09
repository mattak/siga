package cmd

func BoolToFloat64(result bool) float64 {
	if result {
		return 1.0
	} else {
		return 0.0
	}
}
