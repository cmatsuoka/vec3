package vec3

import (
	"fmt"
	"math"
)

type Vector struct {
	x, y, z float64
}

var (
	V0 = Vector{0.0, 0.0, 0.0}
	Vi = Vector{1.0, 0.0, 0.0}
	Vj = Vector{0.0, 1.0, 0.0}
	Vk = Vector{0.0, 0.0, 1.0}
)

func New(x, y, z float64) Vector {
	return Vector{x, y, z}
}

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

// Add returns the resultant of adding vectors v and u
func (v Vector) Add(u Vector) Vector {
	return Vector{v.x + u.x, v.y + u.y, v.z + u.z}
}

// Sub returns the resultant of subtracting vectors v and u
func (v Vector) Sub(u Vector) Vector {
	return Vector{v.x - u.x, v.y - u.y, v.z - u.z}
}

// Neg reverses the direction of vector v.
func (v Vector) Neg() Vector {
	return Vector{-v.x, -v.y, -v.z}
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
