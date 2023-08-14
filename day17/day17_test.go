package day17

import (
	"testing"
)

func TestDay17(t *testing.T) {
	t.Skip()
	target := parseTarget("target area: x=20..30, y=-10..-5")
	launchX, iterations := findValidX(target)

	t.Logf("Iterations: %v\n", iterations)

	if launchX != 6 {
		t.Errorf("Expected launchX of 6, got %v\n", launchX)
	}

	launchY := findValidY(target, 50)

	if launchY != 9 {
		t.Errorf("Expected launchY of 9, got %v\n", launchY)
	}

}

func TestBallStepX(t *testing.T) {
	b := step(ball{dx: 5})
	if b.x != 5 {
		t.Error("Ball X")
	}

	if b.dx != 4 {
		t.Error("Ball dx")
	}

}
func TestBallStepY(t *testing.T) {
	b := step(ball{dy: 5})
	if b.y != 5 {
		t.Error("Ball Y")
	}

	if b.dy != 4 {
		t.Error("Ball dY")
	}

}

func TestSim(t *testing.T) {
	t.Skip()
	target := parseTarget("target area: x=20..30, y=-10..-5")

	// b := ball{dx: 6, dy: 9}
	b := ball{dx: 20, dy: 85}

	for i := 0; i < 200; i++ {
		b = step(b)
		t.Logf("Step %v inzone: %v ball: %+v\n", b.iter, ballInZone(target, b), b)
	}

	// t.Fail()

	if !ballInZone(target, b) {
		t.Fail()
	}
}

func TestHitCount(t *testing.T) {
	target := parseTarget("target area: x=20..30, y=-10..-5")

	hits := countHits(target)

	if hits != 112 {
		t.Errorf("Expected 112 hits but got %v\n", hits)
	}
}
