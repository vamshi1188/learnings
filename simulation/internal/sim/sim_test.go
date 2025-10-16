package sim

import (
	"math"
	"testing"
)

func TestCircularOrbitStability(t *testing.T) {
	bh := NewBlackHole(10)
	radius := 10 * bh.SchwarzschildRadius

	omega, err := CircularAngularVelocity(bh, radius)
	if err != nil {
		t.Fatalf("circular angular velocity: %v", err)
	}

	state := ParticleState{
		Radius:          radius,
		Theta:           0,
		RadialVelocity:  0,
		AngularVelocity: omega,
	}

	cfg := Config{
		StepSizeSeconds: 0.02 * bh.SchwarzschildRadius / SpeedOfLight,
		Steps:           6000,
		HorizonFactor:   1.01,
	}

	result, err := Run(bh, state, cfg)
	if err != nil {
		t.Fatalf("run simulation: %v", err)
	}
	if result.Captured {
		t.Fatalf("particle should not be captured in near-circular orbit")
	}

	finalRadius := result.Samples[len(result.Samples)-1].State.Radius
	fracDiff := math.Abs(finalRadius-radius) / radius
	if fracDiff > 0.02 {
		t.Fatalf("radius drifted too much: got %.4f, want %.4f (diff %.3f)", finalRadius, radius, fracDiff)
	}
}

func TestSpecificEnergyApproximatelyConserved(t *testing.T) {
	bh := NewBlackHole(5)
	radius := 8 * bh.SchwarzschildRadius

	omega, err := CircularAngularVelocity(bh, radius)
	if err != nil {
		t.Fatalf("circular angular velocity: %v", err)
	}

	state := ParticleState{
		Radius:          radius,
		Theta:           0,
		RadialVelocity:  0,
		AngularVelocity: omega * 0.98,
	}

	cfg := Config{
		StepSizeSeconds: 0.015 * bh.SchwarzschildRadius / SpeedOfLight,
		Steps:           8000,
		HorizonFactor:   1.01,
	}

	result, err := Run(bh, state, cfg)
	if err != nil {
		t.Fatalf("run simulation: %v", err)
	}

	initialEnergy, err := SpecificEnergy(bh, state)
	if err != nil {
		t.Fatalf("initial energy: %v", err)
	}
	finalEnergy, err := SpecificEnergy(bh, result.Samples[len(result.Samples)-1].State)
	if err != nil {
		t.Fatalf("final energy: %v", err)
	}

	if initialEnergy == 0 {
		t.Fatalf("unexpected zero initial energy")
	}

	fracDiff := math.Abs(finalEnergy-initialEnergy) / math.Abs(initialEnergy)
	if fracDiff > 0.05 {
		t.Fatalf("energy drift %.3f exceeds tolerance", fracDiff)
	}
}

func TestCaptureDetection(t *testing.T) {
	bh := NewBlackHole(3)
	radius := 4 * bh.SchwarzschildRadius

	omega, err := CircularAngularVelocity(bh, radius)
	if err != nil {
		t.Fatalf("circular angular velocity: %v", err)
	}

	state := ParticleState{
		Radius:          radius,
		Theta:           0,
		RadialVelocity:  -0.05 * SpeedOfLight,
		AngularVelocity: omega * 0.3,
	}

	cfg := Config{
		StepSizeSeconds: 0.01 * bh.SchwarzschildRadius / SpeedOfLight,
		Steps:           20000,
		HorizonFactor:   1.01,
	}

	result, err := Run(bh, state, cfg)
	if err != nil {
		t.Fatalf("run simulation: %v", err)
	}

	if !result.Captured {
		t.Fatalf("expected capture but particle escaped")
	}
	if result.CaptureIndex <= 0 {
		t.Fatalf("invalid capture index: %d", result.CaptureIndex)
	}
}
