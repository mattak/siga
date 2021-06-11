package toolkit

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadDataFrameByStdinTsv() *DataFrame {
	df := &DataFrame{}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	text := string(bytes)
	text = strings.TrimRight(text, "\n")
	lines := strings.Split(text, "\n")
	if len(lines) < 1 {
		log.Fatal(err)
	}

	df.Headers = strings.Split(lines[0], "\t")
	if len(lines) < 2 {
		return df
	}

	df.Labels = make([]string, len(lines)-1)
	df.Data = make([][]float64, len(lines)-1)
	for i := 1; i < len(lines); i++ {
		cells := strings.Split(lines[i], "\t")
		df.Labels[i-1] = cells[0]
		df.Data[i-1] = make([]float64, len(cells)-1)

		for j := 1; j < len(cells); j++ {
			df.Data[i-1][j-1] = ParseFloat64(cells[j])
		}
	}
	return df
}
