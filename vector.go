package vec3

import (
	"fmt"
	"math"
)

type Vector struct {
	x, y, z float64
}

var (
	Vi = Vector{1.0, 0.0, 0.0}
	Vj = Vector{0.0, 1.0, 0.0}
	Vk = Vector{0.0, 0.0, 1.0}
)

// String formats vector v as a string.
func (v Vector) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.x, v.y, v.z)
}

func equalT(a, b, e float64) bool {
	return a-b <= e && b-a <= e
}

// Equal checks if vector v and u are equal, with tolerance e.
func (v Vector) Equal(u Vector, e float64) bool {
	return equalT(v.x, u.x, e) && equalT(v.y, u.y, e) && equalT(v.z, u.z, e)
}

// Modulus computes the magnitude of vector v.
func (v Vector) Modulus() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// Unit returns an unit vector with the same direction as vector v.
func (v Vector) Unit() Vector {
	m := v.Modulus()
	return Vector{v.x / m, v.y / m, v.z / m}
}

// Dot computes the dot product between vectors v and u.
func (v Vector) Dot(u Vector) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

// Cross computes the cross product between vectors v and u.
func (v Vector) Cross(u Vector) Vector {
	return Vector{
		v.y*u.z - v.z*u.y,
		v.z*u.x - v.x*u.z,
		v.x*u.y - v.y*u.x,
	}
}

// Angle returns the angle, in radians, between vectors v and u.
func (v Vector) Angle(u Vector) float64 {
	vu := v.Unit()
	return math.Acos(vu.Dot(u.Unit()))
}
