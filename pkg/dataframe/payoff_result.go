package dataframe

import (
	"fmt"
	"strings"
)

type PayoffResult struct {
	PayoffRatio   float64 `json:"payoff_ratio"`
	ProfitAverage float64 `json:"profit_average"`
	LossAverage   float64 `json:"loss_average"`
	ProfitCount   int     `json:"profit_count"`
	LossCount     int     `json:"loss_count"`
}

func (pr *PayoffResult) ToVector() Vector {
	return Vector{
		pr.PayoffRatio,
		pr.ProfitAverage,
		pr.LossAverage,
		float64(pr.ProfitCount),
		float64(pr.LossCount),
	}
}

func (pr *PayoffResult) ToHeader() []string {
	return []string{"payoff", "profit_average", "loss_average", "profit_count", "loss_count"}
}

func (pr *PayoffResult) PrintTsvHeader() {
	fmt.Println(strings.Join(pr.ToHeader(), "\t"))
}

func (pr *PayoffResult) PrintTsvData(precise bool) {
	pr.ToVector().PrintTsv(precise)
}
