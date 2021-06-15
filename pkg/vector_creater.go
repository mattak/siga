package pkg

import (
	"math"
	"math/rand"
	"time"
)

func CreateVector(size int) Vector {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.NaN()
	}
	return data
}

func CreateVectorWithValue(size int, value float64) Vector {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = value
	}
	return data
}

func CreateVectorWithRandom(size int, fromValue float64, toValue float64) Vector {
	rand.Seed(time.Now().UTC().UnixNano())
	if toValue < fromValue {
		tmp := fromValue
		fromValue = toValue
		toValue = tmp
	}
	span := toValue - fromValue

	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Float64()*span + fromValue
	}
	return data
}

func CreateVectorWithRandomWalk(size int, setting RwalkSetting) Vector {
	rand.Seed(time.Now().UTC().UnixNano())
	data := make([]float64, size)

	// u: 15:0.05,0.15
	// d: 15:0,0.05
	// l: 35:-0.05,0
	// r: 35:-0.15,-0.05
	data[0] = setting.StartValue
	for i := 1; i < size; i++ {
		ratio := setting.GetNextRange().CreateRand()
		data[i] = (ratio + 1) * data[i-1]
	}

	return data
}
