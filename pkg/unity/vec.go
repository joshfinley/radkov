package unity

import (
	"encoding/binary"
	"math"
)

type Vec2 struct {
	X float32
	Y float32
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type Vec4 struct {
	X float32
	Y float32
	Z float32
	D float32
}

type Matrix3x4 struct {
	A Vec4
	B Vec4
	C Vec4
}

type Matrix4x4 struct {
	A Vec4
	B Vec4
	C Vec4
	D Vec4
}

// Get Vec2 from array of bytes
func UnmarshalVec2(v []byte) Vec2 {
	x := v[0:4]
	y := v[4:8]

	fx := math.Float32frombits(
		binary.LittleEndian.Uint32(x))
	fy := math.Float32frombits(
		binary.LittleEndian.Uint32(y))

	return Vec2{
		X: fx,
		Y: fy,
	}
}

// Get array of bytes from Vec2
func MarshalVec2(v Vec2) []byte {
	var raw []byte
	binary.LittleEndian.PutUint32(
		raw[0:4], math.Float32bits(v.X))
	binary.LittleEndian.PutUint32(
		raw[4:8], math.Float32bits(v.Y))

	return raw
}

// Get Vec3 from array of bytes
func UnmarshalVec3(v [12]byte) Vec3 {
	x := v[0:4]
	y := v[4:8]
	z := v[8:12]

	fx := math.Float32frombits(
		binary.LittleEndian.Uint32(x))
	fy := math.Float32frombits(
		binary.LittleEndian.Uint32(y))
	fz := math.Float32frombits(
		binary.LittleEndian.Uint32(z))

	return Vec3{
		X: fx,
		Y: fy,
		Z: fz,
	}
}

// Get array of bytes from Vec3
func MarshalVec3(v Vec3) [12]byte {
	var raw [12]byte
	binary.LittleEndian.PutUint32(
		raw[0:4], math.Float32bits(v.X))
	binary.LittleEndian.PutUint32(
		raw[4:8], math.Float32bits(v.Y))
	binary.LittleEndian.PutUint32(
		raw[8:12], math.Float32bits(v.Z))

	return raw
}
