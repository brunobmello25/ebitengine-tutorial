package main

import "math"

type Camera struct {
	X, Y float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{x, y}
}

func (c *Camera) FollowTarget(targetX, targetY, screenWidth, screenHeight float64) {
	c.X = -targetX + screenWidth/2
	c.Y = -targetY + screenHeight/2
}

func (c *Camera) Constrain(tilemapWidth, tilemapHeight, screenWidth, screenHeight float64) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, screenWidth-float64(tilemapWidth))
	c.Y = math.Max(c.Y, screenHeight-float64(tilemapHeight))
}
