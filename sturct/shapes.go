package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}


type Circle struct {
	Radius float64
}

// Go 語言沒有像一些其他語言那樣直接支持傳統的函數重載（function overloading）。在 Go 中，每個函數必須具有唯一的簽名（包括函數名、參數類型和返回類型）。
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64{
	return 0
}
