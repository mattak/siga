package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"testing"
)

func TestMatrixInnerProduct(t *testing.T) {
	matrix := dataframe.Matrix{{1,2,3}, {10, 10, 10}}
	product := matrix.InnerProduct()
	if len(product) != 3 {
		log.Fatal("matrix product length is not expected")
	}
	if product[0] != 10 {
		log.Fatal("product[0] is not expected")
	}
	if product[1] != 20 {
		log.Fatal("product[1] is not expected")
	}
	if product[2] != 30 {
		log.Fatal("product[2] is not expected")
	}
}
