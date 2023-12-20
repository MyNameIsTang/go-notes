package demo

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func InitValuable() {
	var o valuable = stockPosition{"gGoo", 743.12, 32}
	showValue(o)
	o = car{"as", "su", 2900}
	showValue(o)
}
