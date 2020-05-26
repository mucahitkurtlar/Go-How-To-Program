package main

type rectangle struct {
	width, height float64
	x, y          float64
}

func checkCollision(r1, r2 rectangle) bool {
	if r1.x-r1.width/2 < r2.x+r2.width/2 && r1.x+r1.width/2 > r2.x-r2.width/2 && r1.y-r1.height/2 < r2.y+r2.height/2 && r1.y+r1.height/2 > r2.y-r2.height/2 {
		return true
	}
	return false
}

//r1.x < r2.x+r2.width && r1.x+r1.width > r2.x && r1.y < r2.y+r2.height && r1.y+r1.height > r2.y
