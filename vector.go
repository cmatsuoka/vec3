package vec3

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

var (
	V0 = Vector{0.0, 0.0, 0.0}
	Vi = Vector{1.0, 0.0, 0.0}
	Vj = Vector{0.0, 1.0, 0.0}
	Vk = Vector{0.0, 0.0, 1.0}
)

// String formats vector v as a string.
func (v Vector) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func equalT(a, b, e float64) bool {
	return a-b <= e && b-a <= e
}

// Equal checks if vector v and u are equal, with tolerance e.
func (v Vector) Equal(u Vector, e float64) bool {
	return equalT(v.X, u.X, e) && equalT(v.Y, u.Y, e) && equalT(v.Z, u.Z, e)
}

// Add returns the resultant of adding vectors v and u
func (v Vector) Add(u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// Sub returns the resultant of subtracting vectors v and u
func (v Vector) Sub(u Vector) Vector {
	return Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// Neg reverses the direction of vector v.
func (v Vector) Neg() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}

// Modulus computes the magnitude of vector v.
func (v Vector) Modulus() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Unit returns an unit vector with the same direction as vector v.
func (v Vector) Unit() Vector {
	m := v.Modulus()
	return Vector{v.X / m, v.Y / m, v.Z / m}
}

// Dot computes the dot product between vectors v and u.
func (v Vector) Dot(u Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Cross computes the cross product between vectors v and u.
func (v Vector) Cross(u Vector) Vector {
	return Vector{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

// Angle returns the angle, in radians, between vectors v and u.
func (v Vector) Angle(u Vector) float64 {
	vu := v.Unit()
	return math.Acos(vu.Dot(u.Unit()))
}
