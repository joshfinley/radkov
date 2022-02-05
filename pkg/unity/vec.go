package unity

type Vec2 struct {
	X float64
	Y float64
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Vec4 struct {
	X float64
	Y float64
	Z float64
	D float64
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
