package dataframe

import (
	"github.com/mattak/siga/pkg/util"
	"math/rand"
	"strings"
)

type RwalkValue struct {
	Weight float64
	Range  RangeValue
}

type RwalkSetting struct {
	StartValue    float64
	Probabilities []float64
	Ranges        []RangeValue
}

func CreateRwalkValue(text string, delimiter string) RwalkValue {
	values := strings.Split(text, delimiter)

	return RwalkValue{
		Weight: util.ParseFloat64(values[0]),
		Range: RangeValue{
			From: util.ParseFloat64(values[1]),
			To:   util.ParseFloat64(values[2]),
		},
	}
}

func CreateRwalkSetting(startValue float64, values []RwalkValue) RwalkSetting {
	setting := RwalkSetting{
		StartValue:    startValue,
		Probabilities: make([]float64, len(values)),
		Ranges:        make([]RangeValue, len(values)),
	}

	totalWeights := 0.0
	for _, value := range values {
		totalWeights += value.Weight
	}
	setting.Ranges[0] = values[0].Range
	setting.Probabilities[0] = values[0].Weight / totalWeights

	for i := 1; i < len(values); i++ {
		setting.Ranges[i] = values[i].Range
		setting.Probabilities[i] = values[i].Weight/totalWeights + setting.Probabilities[i-1]
	}

	return setting
}

func (rs RwalkSetting) GetNextRange() RangeValue {
	r := rand.Float64()
	for i := 0; i < len(rs.Probabilities); i++ {
		if r <= rs.Probabilities[i] {
			return rs.Ranges[i]
		}
	}
	return rs.Ranges[len(rs.Probabilities)-1]
}
