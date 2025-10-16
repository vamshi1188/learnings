package sim

import (
	"errors"
	"math"
)

// BlackHole represents a simplified Schwarzschild black hole characterized by its mass.
type BlackHole struct {
	MassKilograms       float64
	GravitationalParam  float64
	SchwarzschildRadius float64
}

const (
	gravitationalConstant = 6.67430e-11
	speedOfLight          = 299792458.0
	solarMassKilograms    = 1.98847e30
)

// NewBlackHole constructs a BlackHole from a mass expressed in solar masses.
func NewBlackHole(massSolar float64) BlackHole {
	massKg := massSolar * solarMassKilograms
	mu := gravitationalConstant * massKg
	rs := 2 * mu / (speedOfLight * speedOfLight)
	return BlackHole{
		MassKilograms:       massKg,
		GravitationalParam:  mu,
		SchwarzschildRadius: rs,
	}
}

// ParticleState captures the position and velocity of a test particle in polar coordinates (r, θ).
type ParticleState struct {
	Radius          float64 // meters
	Theta           float64 // radians
	RadialVelocity  float64 // meters/second
	AngularVelocity float64 // radians/second
}

// Sample stores a snapshot of the particle's state at a specific simulation time.
type Sample struct {
	Time  float64
	State ParticleState
}

// Config holds the user-configurable parameters for the simulation.
type Config struct {
	StepSizeSeconds float64
	Steps           int
	HorizonFactor   float64
}

// Result collects the simulated states for analysis or visualization.
type Result struct {
	Samples      []Sample
	Captured     bool
	CaptureIndex int
}

// Run executes the simulation for the given black hole, initial state, and configuration.
func Run(bh BlackHole, initial ParticleState, cfg Config) (Result, error) {
	if cfg.StepSizeSeconds <= 0 {
		return Result{}, errors.New("step size must be positive")
	}
	if cfg.Steps <= 0 {
		return Result{}, errors.New("steps must be positive")
	}
	if initial.Radius <= bh.SchwarzschildRadius {
		return Result{}, errors.New("initial radius must exceed the event horizon")
	}

	horizonRadius := bh.SchwarzschildRadius
	if cfg.HorizonFactor > 0 {
		horizonRadius = bh.SchwarzschildRadius * cfg.HorizonFactor
	}

	samples := make([]Sample, 0, cfg.Steps+1)
	vec := fromState(initial)
	samples = append(samples, Sample{Time: 0, State: initial})

	stepFn := func(v stateVector) stateVector {
		return derivatives(bh, v)
	}

	captured := false
	captureIndex := -1
	dt := cfg.StepSizeSeconds

	for i := 1; i <= cfg.Steps; i++ {
		vec = rk4Step(vec, dt, stepFn)
		state := toState(vec)

		if math.IsNaN(state.Radius) || math.IsInf(state.Radius, 0) {
			return Result{}, errors.New("numerical instability encountered (radius)")
		}
		if math.IsNaN(state.AngularVelocity) || math.IsInf(state.AngularVelocity, 0) {
			return Result{}, errors.New("numerical instability encountered (angular velocity)")
		}

		time := float64(i) * dt
		samples = append(samples, Sample{Time: time, State: state})

		if state.Radius <= horizonRadius {
			captured = true
			captureIndex = i
			break
		}
	}

	return Result{
		Samples:      samples,
		Captured:     captured,
		CaptureIndex: captureIndex,
	}, nil
}

// CircularAngularVelocity returns the angular velocity required for a circular orbit at the specified radius.
func CircularAngularVelocity(bh BlackHole, radius float64) (float64, error) {
	if radius <= bh.SchwarzschildRadius {
		return 0, errors.New("radius must exceed the event horizon")
	}
	denom := radius - bh.SchwarzschildRadius
	if denom <= 0 {
		return 0, errors.New("radius too close to event horizon for circular orbit")
	}
	omegaSquared := bh.GravitationalParam / (radius * denom * denom)
	if omegaSquared <= 0 {
		return 0, errors.New("non-positive angular velocity squared")
	}
	return math.Sqrt(omegaSquared), nil
}

// SpecificEnergy returns the specific mechanical energy (per unit mass) of the given state.
func SpecificEnergy(bh BlackHole, state ParticleState) (float64, error) {
	if state.Radius <= bh.SchwarzschildRadius {
		return 0, errors.New("energy undefined inside the event horizon")
	}
	denom := state.Radius - bh.SchwarzschildRadius
	if denom <= 0 {
		return 0, errors.New("radius too close to event horizon")
	}
	kinetic := 0.5 * (state.RadialVelocity*state.RadialVelocity + (state.Radius*state.AngularVelocity)*(state.Radius*state.AngularVelocity))
	potential := -bh.GravitationalParam / denom
	return kinetic + potential, nil
}

type stateVector struct {
	Radius          float64
	Theta           float64
	RadialVelocity  float64
	AngularVelocity float64
}

type derivativeFunc func(stateVector) stateVector

func rk4Step(current stateVector, dt float64, f derivativeFunc) stateVector {
	k1 := f(current)
	k2 := f(addStateVectors(current, scaleStateVector(k1, dt*0.5)))
	k3 := f(addStateVectors(current, scaleStateVector(k2, dt*0.5)))
	k4 := f(addStateVectors(current, scaleStateVector(k3, dt)))

	increment := stateVector{
		Radius:          (k1.Radius + 2*k2.Radius + 2*k3.Radius + k4.Radius) * dt / 6,
		Theta:           (k1.Theta + 2*k2.Theta + 2*k3.Theta + k4.Theta) * dt / 6,
		RadialVelocity:  (k1.RadialVelocity + 2*k2.RadialVelocity + 2*k3.RadialVelocity + k4.RadialVelocity) * dt / 6,
		AngularVelocity: (k1.AngularVelocity + 2*k2.AngularVelocity + 2*k3.AngularVelocity + k4.AngularVelocity) * dt / 6,
	}

	return addStateVectors(current, increment)
}

func addStateVectors(a, b stateVector) stateVector {
	return stateVector{
		Radius:          a.Radius + b.Radius,
		Theta:           a.Theta + b.Theta,
		RadialVelocity:  a.RadialVelocity + b.RadialVelocity,
		AngularVelocity: a.AngularVelocity + b.AngularVelocity,
	}
}

func scaleStateVector(v stateVector, factor float64) stateVector {
	return stateVector{
		Radius:          v.Radius * factor,
		Theta:           v.Theta * factor,
		RadialVelocity:  v.RadialVelocity * factor,
		AngularVelocity: v.AngularVelocity * factor,
	}
}

func fromState(s ParticleState) stateVector {
	return stateVector(s)
}

func toState(v stateVector) ParticleState {
	return ParticleState(v)
}

func derivatives(bh BlackHole, v stateVector) stateVector {
	r := v.Radius
	if r <= 0 {
		return stateVector{}
	}

	denom := r - bh.SchwarzschildRadius
	minDenom := 1e-9 * math.Max(bh.SchwarzschildRadius, 1)
	if denom < minDenom {
		denom = minDenom
	}
	grav := bh.GravitationalParam / (denom * denom)

	radialAccel := r*v.AngularVelocity*v.AngularVelocity - grav
	angularAccel := -2 * v.RadialVelocity * v.AngularVelocity / r

	return stateVector{
		Radius:          v.RadialVelocity,
		Theta:           v.AngularVelocity,
		RadialVelocity:  radialAccel,
		AngularVelocity: angularAccel,
	}
}
