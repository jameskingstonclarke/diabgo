package util


type Vec3f struct {
	X, Y, Z float64
}

func (vec Vec3f) add(other Vec3f) (Vec3f){
	return Vec3f{vec.X + other.X, vec.Y + other.Y, vec.Z + other.Z}
}

func (vec Vec3f) sub(other Vec3f) (Vec3f){
	return Vec3f{vec.X - other.X, vec.Y - other.Y, vec.Z - other.Z}
}

func (vec Vec3f) mul(other Vec3f) (Vec3f){
	return Vec3f{vec.X * other.X, vec.Y * other.Y, vec.Z * other.Z}
}

func (vec Vec3f) div(other Vec3f) (Vec3f){
	return Vec3f{vec.X / other.X, vec.Y / other.Y, vec.Z / other.Z}
}