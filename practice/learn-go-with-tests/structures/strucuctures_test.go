package structures

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expect := 40.0

	if result != expect {
		t.Errorf("result %.2f, expect %.2f ", result, expect)
	}
}

func TestArea(t *testing.T) {
	testArea := []struct {
		form   Form
		expect float64
	}{
		{form: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{form: Circle{Ray: 10}, expect: 314.1592653589793},
		{form: Triangle{Base: 12, Width: 6}, expect: 36.0},
	}

	for _, tt := range testArea {
		result := tt.form.Area()
		if result != tt.expect {
			t.Errorf("%#v result %.2f, expect %.2f", tt.form, result, tt.expect)
		}
	}
}

// func TestArea(t *testing.T) {

// 	verifyArea := func(t *testing.T, form Form, expect float64) {
// 		t.Helper()
// 		result := form.Area()

// 		if result != expect {
// 			t.Errorf("result %.2f, expect %.2f", result, expect)
// 		}
// 	}
// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		verifyArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		verifyArea(t, circle, 314.1592653589793)
// 	})
// }
