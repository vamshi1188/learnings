package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/vamshi/simulation/internal/render"
	"github.com/vamshi/simulation/internal/sim"
)

type options struct {
	massSolar          float64
	radiusMultiples    float64
	radialVelocityFrac float64
	tangentialFactor   float64
	stepSizeSeconds    float64
	steps              int
	horizonFactor      float64
	outputPath         string
	ascii              bool
	renderPath         string
	renderWidth        int
	renderHeight       int
	renderTilt         float64
	renderExposure     float64
	renderStepSize     float64
	renderMaxSteps     int
	renderFarPlane     float64
	renderHorizon      float64
	renderCameraDist   float64
	renderCameraHeight float64
	renderCameraAz     float64
}

func main() {
	opts := parseFlags()
	if err := run(opts); err != nil {
		log.Fatal(err)
	}
}

func parseFlags() options {
	opts := options{}
	flag.Float64Var(&opts.massSolar, "mass", 10, "Black hole mass in solar masses")
	flag.Float64Var(&opts.radiusMultiples, "radius", 10, "Initial radius expressed in Schwarzschild radii")
	flag.Float64Var(&opts.radialVelocityFrac, "vr", 0, "Initial radial velocity as a fraction of the speed of light")
	flag.Float64Var(&opts.tangentialFactor, "vt-factor", 0.95, "Tangential velocity as a fraction of the circular orbit speed")
	flag.Float64Var(&opts.stepSizeSeconds, "step", 0, "Integrator step size in seconds (0 = adaptive default)")
	flag.IntVar(&opts.steps, "steps", 8000, "Number of integration steps")
	flag.Float64Var(&opts.horizonFactor, "horizon-factor", 1.01, "Multiple of the Schwarzschild radius that counts as capture")
	flag.StringVar(&opts.outputPath, "output", "", "Optional CSV output path for trajectory samples")
	flag.BoolVar(&opts.ascii, "ascii", false, "Render an ASCII snapshot of the trajectory")
	flag.StringVar(&opts.renderPath, "render", "", "Optional PNG output path for an Interstellar-style visualization")
	flag.IntVar(&opts.renderWidth, "render-width", 1280, "Render width in pixels")
	flag.IntVar(&opts.renderHeight, "render-height", 720, "Render height in pixels")
	flag.Float64Var(&opts.renderTilt, "render-tilt", 20, "Accretion disk tilt angle in degrees")
	flag.Float64Var(&opts.renderExposure, "render-exposure", 1.8, "Tone-mapping exposure for render")
	flag.Float64Var(&opts.renderStepSize, "render-step", 0.02, "Integrator step size in Schwarzschild radii for render")
	flag.IntVar(&opts.renderMaxSteps, "render-max-steps", 2800, "Maximum steps for render integrator")
	flag.Float64Var(&opts.renderFarPlane, "render-far", 90, "Distance in Schwarzschild radii at which rays escape to background")
	flag.Float64Var(&opts.renderHorizon, "render-horizon", 1.03, "Multiple of Schwarzschild radius treated as capture in render")
	flag.Float64Var(&opts.renderCameraDist, "render-camera-dist", 22, "Camera distance from the black hole in Schwarzschild radii")
	flag.Float64Var(&opts.renderCameraHeight, "render-camera-height", 6, "Camera height above the disk plane in Schwarzschild radii")
	flag.Float64Var(&opts.renderCameraAz, "render-camera-az", 18, "Camera azimuth angle in degrees")
	flag.Parse()
	return opts
}

func run(opts options) error {
	if opts.radiusMultiples <= 1 {
		return fmt.Errorf("initial radius must exceed 1 Schwarzschild radius")
	}
	if opts.steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}

	bh := sim.NewBlackHole(opts.massSolar)
	initialRadius := opts.radiusMultiples * bh.SchwarzschildRadius

	stepSize := opts.stepSizeSeconds
	if stepSize <= 0 {
		stepSize = 0.05 * bh.SchwarzschildRadius / sim.SpeedOfLight
	}

	circOmega, err := sim.CircularAngularVelocity(bh, initialRadius)
	if err != nil {
		return fmt.Errorf("compute circular angular velocity: %w", err)
	}

	angularVelocity := circOmega * opts.tangentialFactor
	radialVelocity := opts.radialVelocityFrac * sim.SpeedOfLight

	state := sim.ParticleState{
		Radius:          initialRadius,
		Theta:           0,
		RadialVelocity:  radialVelocity,
		AngularVelocity: angularVelocity,
	}

	cfg := sim.Config{
		StepSizeSeconds: stepSize,
		Steps:           opts.steps,
		HorizonFactor:   opts.horizonFactor,
	}

	result, err := sim.Run(bh, state, cfg)
	if err != nil {
		return fmt.Errorf("run simulation: %w", err)
	}

	samples := result.Samples
	if len(samples) == 0 {
		return fmt.Errorf("simulation returned no samples")
	}

	if opts.outputPath != "" {
		if err := writeCSV(opts.outputPath, samples, bh.SchwarzschildRadius); err != nil {
			return err
		}
		fmt.Printf("trajectory written to %s\n", opts.outputPath)
	}

	final := samples[len(samples)-1]
	fmt.Printf("steps executed: %d\n", len(samples)-1)
	fmt.Printf("duration: %.3fs\n", final.Time)
	fmt.Printf("final radius: %.3f Rs\n", final.State.Radius/bh.SchwarzschildRadius)
	fmt.Printf("captured: %t\n", result.Captured)

	if opts.ascii {
		fmt.Println()
		fmt.Println(renderASCII(samples, bh.SchwarzschildRadius))
	}

	if opts.renderPath != "" {
		if err := renderImage(opts); err != nil {
			return fmt.Errorf("render image: %w", err)
		}
	}

	return nil
}

func renderImage(opts options) error {
	cfg := render.Config{
		Width:              opts.renderWidth,
		Height:             opts.renderHeight,
		DiskTiltDegrees:    opts.renderTilt,
		Exposure:           opts.renderExposure,
		StepSize:           opts.renderStepSize,
		MaxSteps:           opts.renderMaxSteps,
		FarPlane:           opts.renderFarPlane,
		HorizonFactor:      opts.renderHorizon,
		CameraDistance:     opts.renderCameraDist,
		CameraHeight:       opts.renderCameraHeight,
		CameraAzimuthDeg:   opts.renderCameraAz,
		BlackHoleMassSolar: opts.massSolar,
	}

	img, err := render.Render(cfg)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(opts.renderPath), 0o755); err != nil {
		return fmt.Errorf("create render directory: %w", err)
	}

	file, err := os.Create(opts.renderPath)
	if err != nil {
		return fmt.Errorf("create render file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("encode render: %w", err)
	}

	fmt.Printf("render saved to %s (%dx%d)\n", opts.renderPath, cfg.Width, cfg.Height)
	return nil
}

func writeCSV(path string, samples []sim.Sample, rs float64) error {
	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create output directory: %w", err)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create csv: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"time_s", "radius_m", "radius_rs", "theta_rad", "radial_velocity_m_s", "angular_velocity_rad_s", "x_m", "y_m"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	for _, sample := range samples {
		state := sample.State
		x := state.Radius * math.Cos(state.Theta)
		y := state.Radius * math.Sin(state.Theta)
		record := []string{
			fmt.Sprintf("%.6f", sample.Time),
			fmt.Sprintf("%.6f", state.Radius),
			fmt.Sprintf("%.6f", state.Radius/rs),
			fmt.Sprintf("%.6f", state.Theta),
			fmt.Sprintf("%.6f", state.RadialVelocity),
			fmt.Sprintf("%.6f", state.AngularVelocity),
			fmt.Sprintf("%.6f", x),
			fmt.Sprintf("%.6f", y),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("write record: %w", err)
		}
	}

	return writer.Error()
}

func renderASCII(samples []sim.Sample, rs float64) string {
	const (
		width  = 61
		height = 31
	)

	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	var maxRadius float64
	for _, sample := range samples {
		if sample.State.Radius > maxRadius {
			maxRadius = sample.State.Radius
		}
	}
	if maxRadius == 0 {
		maxRadius = rs
	}

	scale := float64(width-1) / (2 * maxRadius)
	cx := (width - 1) / 2
	cy := (height - 1) / 2

	horizonRadius := rs * scale
	drawCircle(grid, cx, cy, horizonRadius, 'o')

	for _, sample := range samples {
		x := sample.State.Radius * math.Cos(sample.State.Theta)
		y := sample.State.Radius * math.Sin(sample.State.Theta)
		gx := int(math.Round(scale*x)) + cx
		gy := cy - int(math.Round(scale*y))
		if gx >= 0 && gx < width && gy >= 0 && gy < height {
			grid[gy][gx] = '*'
		}
	}

	grid[cy][cx] = 'X'

	lines := make([]byte, 0, height*(width+1))
	for _, row := range grid {
		lines = append(lines, []byte(string(row))...)
		lines = append(lines, '\n')
	}

	return string(lines)
}

func drawCircle(grid [][]rune, cx, cy int, radius float64, char rune) {
	if radius <= 0 {
		return
	}
	width := len(grid[0])
	height := len(grid)
	for angle := 0.0; angle < 2*math.Pi; angle += 0.03 {
		x := int(math.Round(math.Cos(angle)*radius)) + cx
		y := cy - int(math.Round(math.Sin(angle)*radius))
		if x >= 0 && x < width && y >= 0 && y < height {
			if grid[y][x] == ' ' {
				grid[y][x] = char
			}
		}
	}
}
