package entity

// Colour represents a set of bytes to show colour on a display in an RGB format.
type Colour struct {
	Red, Green, Blue byte
}

// Position is a representation of the location of a 2D game object in world
// space.
type Position struct {
	X, Y float32
}

// Velocity is a 2D representation of movement for an object in the game world.
type Velocity struct {
	X, Y float32
}

// Dimensions is a representation of the 2D geomtry that makes up an object in
// the game world.
type Dimensions struct {
	Width, Height, Radius float32
}
