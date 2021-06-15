package dataframe

type DataFrame struct {
	Headers []string    `json:"headers"`
	Labels  []string    `json:"labels"`
	Data    [][]float64 `json:"data"`
}
