package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// 改成10.25, 10.26和41.02會有錯，改成.20f ，got 41.01999999999999602096 want 41.02000000000000312639
		// 建議如果有小數的比較，可以引入一個小的數epsilon，並針對兩數差取絕對值
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got :=  circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

type Shape interface {
	Area() float64
}

//增加接口的使用
func TestAreaInterface(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

// table driven tests 表格驅動測試。
func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{{shape: Rectangle{Width: 12, Height: 6}, want: 72.0}, {Circle{10}, 314.1592653589793}, {Triangle{12, 6}, 36.0}}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}

