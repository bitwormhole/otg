package dto

import "github.com/bitwormhole/nasyun/app/data/dxo"

// Example ... VO
type Example struct {
	ID dxo.ExampleID `json:"id"`
	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
