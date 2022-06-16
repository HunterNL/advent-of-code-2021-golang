package day17

func stepX(x, dx int) (int, int) {
	if dx > 0 {
		return x + dx, dx - 1
	} else {
		return x, 0
	}
}

func stepY(y, dy int) (int, int) {
	return y + dy, dy - 1
}

func xInZone(t target, x int) bool {
	return x >= t.left && x <= t.right
}
func YInZone(t target, y int) bool {
	return y >= t.bottom && y <= t.top
}

func isXHit(launchX int, t target) (bool, int) {
	b := ball{dx: launchX}
	for b.x <= t.right {
		// fmt.Printf("Ball status in X Sim: %+v\n", b)
		if xInZone(t, b.x) {
			return true, b.iter
		}
		if b.dx == 0 {
			return xInZone(t, b.x), 1000
		}
		b = step(b)
	}

	return false, b.iter
}

func isYHit(launchY int, t target, iterations int) bool {
	b := ball{dy: launchY}
	for i := 0; i < iterations; i++ {
		b = step(b)
		// fmt.Printf("Ball status in Y Sim: %+v\n", b)
	}

	return YInZone(t, b.y)
}

type ball struct {
	x, y, dx, dy, iter int
}

func step(b ball) ball {
	newBall := ball{}
	newBall.x, newBall.dx = stepX(b.x, b.dx)
	newBall.y, newBall.dy = stepY(b.y, b.dy)
	newBall.iter = b.iter + 1
	return newBall
}

func ballInZone(t target, b ball) bool {
	return xInZone(t, b.x) && YInZone(t, b.y)
}

func iterTillXHit(t target, launchX, launchY, iterLimit int) bool {
	b := ball{dx: launchX, dy: launchY}

	for i := 0; i < iterLimit; i++ {
		b = step(b)
		if ballInZone(t, b) {
			return true
		}

		// fmt.Printf("X Hit[%v,%v]\n", b.x, b.y)
	}

	return false
}

func iterTillYHit(t target, launchY, iterLimit int) bool {
	b := ball{y: launchY}

	for i := 0; i < iterLimit; i++ {
		b = step(b)
		if YInZone(t, b.y) {
			return true
		}

		// fmt.Printf("Y Hit [%v,%v]\n", b.x, b.y)
	}

	return false
}

func iterTillHit(t target, launchX, launchY, iterLimit int) bool {
	b := ball{dy: launchY, dx: launchX}

	for i := 0; i < iterLimit; i++ {
		b = step(b)
		if ballInZone(t, b) {
			return true
		}
		if b.x > t.right {
			return false
		}
		if b.y < t.bottom {
			return false
		}

		// fmt.Printf("Y Hit [%v,%v]\n", b.x, b.y)
	}

	return false
}
