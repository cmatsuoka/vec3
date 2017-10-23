package vec3

import (
	"math"
	"testing"
)

func TestModulus(t *testing.T) {

	v := Vector{10, 0, 0}
	if s := v.Modulus(); s != 10 {
		t.Errorf("modulus error, expected %f, got %f", 10.0, s)
	}

	v = Vector{0, 20, 0}
	if s := v.Modulus(); s != 20 {
		t.Errorf("modulus error, expected %f, got %f", 20.0, s)
	}

	v = Vector{0, 0, 30}
	if s := v.Modulus(); s != 30 {
		t.Errorf("modulus error, expected %f, got %f", 30.0, s)
	}

	v = Vector{3, 4, 12}
	if s := v.Modulus(); s != 13 {
		t.Errorf("modulus error, expected %f, got %f", 13.0, s)
	}
}

func TestUnit(t *testing.T) {

	v := Vector{10, 0, 0}
	if vu := v.Unit(); !vu.Equal(Vi, 10e-12) {
		t.Errorf("unit error, expected %s, got %s", Vi, vu)
	}

	v = Vector{0, 20, 0}
	if vu := v.Unit(); !vu.Equal(Vj, 10e-12) {
		t.Errorf("unit error, expected %s, got %s", Vj, vu)
	}

	v = Vector{0, 0, 30}
	if vu := v.Unit(); !vu.Equal(Vk, 10e-12) {
		t.Errorf("unit error, expected %s, got %s", Vk, vu)
	}

	v = Vector{10, 20, 30}
	if vu := v.Unit(); !vu.Equal(Vector{0.267261, 0.534522, 0.801784}, 10e-6) {
		t.Errorf("unit error, got %s", vu)
	}

}

func TestDot(t *testing.T) {

	v := Vector{1, 2, 3}
	u := Vector{7, 8, 9}

	if d := v.Dot(u); d != 50 {
		t.Errorf("dot product error, expected %f, got %f", 60, d)
	}

}

func TestCross(t *testing.T) {

	v := Vector{2, 0, 0}
	u := Vector{0, 3, 0}

	if d := v.Cross(u); !d.Equal(Vector{0, 0, 6}, 10e-12) {
		t.Errorf("cross product error, got %s", d)
	}

	v = Vector{0, 0, 4}
	u = Vector{0, 3, 0}

	if d := v.Cross(u); !d.Equal(Vector{-12, 0, 0}, 10e-12) {
		t.Errorf("cross product error, got %s", d)
	}

	v = Vector{1, 2, -3}
	u = Vector{4, 5, 6}

	if d := v.Cross(u); !d.Equal(Vector{27, -18, -3}, 10e-12) {
		t.Errorf("cross product error, got %s", d)
	}

}

func TestAngle(t *testing.T) {

	v := Vector{2, 0, 0}
	u := Vector{0, 3, 0}

	if a := v.Angle(u); a != math.Pi/2 {
		t.Errorf("angle error, expected %f, got %f", math.Pi/2, a)
	}

	v = Vector{0, 0, 4}
	u = Vector{0, 3, 0}

	if a := v.Angle(u); a != math.Pi/2 {
		t.Errorf("angle error, expected %f, got %f", math.Pi/2, a)
	}

	v = Vector{1, 2, -3}
	u = Vector{4, 5, 6}

	if a := v.Angle(u); a-1.692929 > 10e-6 {
		t.Errorf("cross product error, expected %f, got %f", 1.692929, a)
	}

}
