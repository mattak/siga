package dataframe

import "math"

type AverageVariance struct {
	Total    float64 `json:"total"`
	Count    int     `json:"count"`
	Average  float64 `json:"average"`
	Variance float64 `json:"variance"`
}

func (av *AverageVariance) Add(value float64) {
	av.Count++
	av.Total += value
}

func (av *AverageVariance) FinalizeAverageCalculation() {
	if av.Count > 0 {
		av.Average = av.Total / float64(av.Count)
	} else {
		av.Average = 0
	}
}

func (av *AverageVariance) FinalizeVarianceCalculation() {
	if av.Count > 0 {
		av.Variance = math.Sqrt(av.Variance / float64(av.Count))
	} else {
		av.Variance = 0
	}
}
