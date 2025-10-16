package render

import (
	"errors"
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
)

// Config controls the physically-inspired renderer.
type Config struct {
	Width              int
	Height             int
	SamplesPerPixel    int
	FOVDegrees         float64
	BlackHoleMassSolar float64
	DiskInnerRadius    float64
	DiskOuterRadius    float64
	DiskTiltDegrees    float64
	DiskThickness      float64
	Exposure           float64
	StepSize           float64
	MaxSteps           int
	FarPlane           float64
	HorizonFactor      float64
	CameraDistance     float64
	CameraHeight       float64
	CameraAzimuthDeg   float64
	Parallelism        int
}

// Render produces an RGBA image approximating an Interstellar-style black hole.
func Render(cfg Config) (image.Image, error) {
	if cfg.Width <= 0 || cfg.Height <= 0 {
		return nil, errors.New("image dimensions must be positive")
	}

	if cfg.SamplesPerPixel <= 0 {
		cfg.SamplesPerPixel = 1
	}
	if cfg.FOVDegrees <= 0 {
		cfg.FOVDegrees = 55
	}
	if cfg.DiskInnerRadius <= 0 {
		cfg.DiskInnerRadius = 1.3
	}
	if cfg.DiskOuterRadius <= 0 {
		cfg.DiskOuterRadius = 25
	}
	if cfg.DiskThickness <= 0 {
		cfg.DiskThickness = 0.08
	}
	if cfg.Exposure <= 0 {
		cfg.Exposure = 1.6
	}
	if cfg.StepSize <= 0 {
		cfg.StepSize = 0.02
	}
	if cfg.MaxSteps <= 0 {
		cfg.MaxSteps = 2200
	}
	if cfg.FarPlane <= 0 {
		cfg.FarPlane = 80
	}
	if cfg.HorizonFactor <= 0 {
		cfg.HorizonFactor = 1.02
	}
	if cfg.CameraDistance <= 0 {
		cfg.CameraDistance = 18
	}
	if cfg.CameraHeight == 0 {
		cfg.CameraHeight = 5
	}
	if cfg.Parallelism <= 0 {
		cfg.Parallelism = runtime.NumCPU()
	}

	eye := defaultCameraEye(cfg)
	target := Vec3{0, 0, 0}
	up := Vec3{0, 1, 0}
	forward := target.Sub(eye).Normalize()
	right := forward.Cross(up).Normalize()
	upVec := right.Cross(forward).Normalize()

	tilt := cfg.DiskTiltDegrees
	if tilt == 0 {
		tilt = 18
	}
	diskNormal := rotateAroundAxis(Vec3{0, 1, 0}, right, tilt*math.Pi/180)
	diskBasisU := diskNormal.Cross(Vec3{0, 0, 1})
	if diskBasisU.Length() < 1e-5 {
		diskBasisU = diskNormal.Cross(Vec3{1, 0, 0})
	}
	diskBasisU = diskBasisU.Normalize()

	img := image.NewRGBA(image.Rect(0, 0, cfg.Width, cfg.Height))

	aspect := float64(cfg.Width) / float64(cfg.Height)
	scale := math.Tan(cfg.FOVDegrees * math.Pi / 180 / 2)

	tasks := make(chan int, cfg.Height)
	var wg sync.WaitGroup
	for i := 0; i < cfg.Parallelism; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range tasks {
				for x := 0; x < cfg.Width; x++ {
					dir := cameraRayDirection(x, y, cfg.Width, cfg.Height, aspect, scale, forward, right, upVec)
					col := traceRay(cfg, eye, dir, diskNormal, diskBasisU)
					img.Set(x, y, toRGBA(col, cfg.Exposure))
				}
			}
		}()
	}

	for y := 0; y < cfg.Height; y++ {
		tasks <- y
	}
	close(tasks)
	wg.Wait()

	return img, nil
}

func defaultCameraEye(cfg Config) Vec3 {
	az := cfg.CameraAzimuthDeg * math.Pi / 180
	if cfg.CameraAzimuthDeg == 0 {
		az = 15 * math.Pi / 180
	}
	x := cfg.CameraDistance * math.Sin(az)
	z := -cfg.CameraDistance * math.Cos(az)
	return Vec3{x, cfg.CameraHeight, z}
}

func cameraRayDirection(x, y, width, height int, aspect, scale float64, forward, right, up Vec3) Vec3 {
	u := (float64(x) + 0.5) / float64(width)
	v := (float64(y) + 0.5) / float64(height)
	px := (2*u - 1) * aspect * scale
	py := (1 - 2*v) * scale
	dir := right.Scale(px).Add(up.Scale(py)).Add(forward)
	return dir.Normalize()
}

func traceRay(cfg Config, origin, direction Vec3, diskNormal, diskU Vec3) Vec3 {
	pos := origin
	dir := direction

	for i := 0; i < cfg.MaxSteps; i++ {
		r := pos.Length()
		if r <= cfg.HorizonFactor {
			return Vec3{0, 0, 0}
		}
		if r >= cfg.FarPlane {
			return backgroundColor(dir)
		}

		dist := diskNormal.Dot(pos)
		if math.Abs(dist) <= cfg.DiskThickness {
			col, ok := diskShade(cfg, pos, dir, diskNormal, diskU)
			if ok {
				return col
			}
		}

		acc := gravitationalAcceleration(pos)
		dir = dir.Add(acc.Scale(cfg.StepSize)).Normalize()
		pos = pos.Add(dir.Scale(cfg.StepSize))
	}

	return Vec3{0, 0, 0}
}

func diskShade(cfg Config, pos, dir, normal, diskU Vec3) (Vec3, bool) {
	planePoint := pos.Sub(normal.Scale(normal.Dot(pos)))
	radius := planePoint.Length()
	if radius < cfg.DiskInnerRadius || radius > cfg.DiskOuterRadius {
		return Vec3{}, false
	}

	radiusT := clamp((radius-cfg.DiskInnerRadius)/(cfg.DiskOuterRadius-cfg.DiskInnerRadius), 0, 1)
	inner := Vec3{1.8, 1.0, 0.7}
	outer := Vec3{0.2, 0.1, 0.08}
	base := inner.Scale(1 - radiusT).Add(outer.Scale(radiusT))

	tangent := normal.Cross(planePoint).Normalize()
	if tangent.Length() == 0 {
		tangent = diskU
	}
	beta := math.Sqrt(clamp(1/(2*radius), 0, 0.7*0.7))
	beta = clamp(beta, 0, 0.65)
	gamma := 1 / math.Sqrt(1-beta*beta)
	cosTheta := tangent.Dot(dir.Scale(-1))
	doppler := 1 / (gamma * (1 - beta*cosTheta))
	doppler = clamp(doppler, 0.2, 4)

	redshift := math.Sqrt(math.Max(1-1/radius, 0.05))
	intensity := doppler * doppler * redshift

	col := base.Scale(intensity)
	return col, true
}

func gravitationalAcceleration(pos Vec3) Vec3 {
	r := pos.Length()
	if r == 0 {
		return Vec3{}
	}
	rs := 1.0
	denom := math.Max(r-rs, 1e-3)
	strength := -0.5 / (denom * denom * r)
	return pos.Normalize().Scale(strength)
}

func backgroundColor(dir Vec3) Vec3 {
	d := dir.Normalize()
	bright := math.Pow(math.Max(0, d.Y), 0.8)

	noise := hashNoise(d)
	stars := clamp(noise*noise*3, 0, 1)

	base := Vec3{0.02, 0.03, 0.05}
	nebula := Vec3{0.4, 0.2, 0.6}.Scale(bright)
	starColor := Vec3{1, 0.95, 0.85}.Scale(stars)

	return base.Add(nebula).Add(starColor)
}

func hashNoise(v Vec3) float64 {
	s := math.Sin(v.X*12.9898 + v.Y*78.233 + v.Z*37.719)
	return fract(s * 43758.5453)
}

func rotateAroundAxis(vec, axis Vec3, angle float64) Vec3 {
	a := axis.Normalize()
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return vec.Scale(cos).
		Add(a.Cross(vec).Scale(sin)).
		Add(a.Scale(a.Dot(vec) * (1 - cos)))
}

func toRGBA(col Vec3, exposure float64) color.RGBA {
	tone := func(c float64) uint8 {
		mapped := 1 - math.Exp(-c*exposure)
		mapped = clamp(mapped, 0, 1)
		return uint8(mapped * 255)
	}
	return color.RGBA{R: tone(col.X), G: tone(col.Y), B: tone(col.Z), A: 255}
}

func fract(x float64) float64 {
	_, frac := math.Modf(x)
	if frac < 0 {
		frac += 1
	}
	return frac
}
