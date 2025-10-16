# Black Hole Trajectory Simulation

A Go-based toy simulator that integrates test-particle trajectories around a non-rotating (Schwarzschild) black hole. The dynamics use a Paczyński–Wiita pseudo-Newtonian potential and a fourth-order Runge–Kutta integrator to illustrate stable orbits, inspirals, and captures.

## Quick start

Run the CLI directly from the repository:

```bash
go run ./cmd/simulate --mass 10 --radius 12 --steps 6000 --ascii
```

The command prints a short numerical summary and, with `--ascii`, renders a coarse trajectory plot. Use `--output` to dump all samples to CSV:

```bash
go run ./cmd/simulate --mass 4e6 --radius 20 --vt-factor 0.98 --steps 8000 --output data/orbit.csv
```

Generate an Interstellar-inspired black hole render by providing a PNG path:

```bash
go run ./cmd/simulate --mass 6.5e9 --radius 25 --steps 1 --render renders/gargantua.png --render-width 1920 --render-height 1080 --render-tilt 25
```

The renderer integrates photon trajectories through a pseudo-relativistic potential to approximate gravitational lensing, Doppler beaming, and gravitational redshift. Increase `--render-max-steps` for sharper limb detail or tweak `--render-exposure` to rebalance brightness.

## Key flags

| Flag | Description | Default |
| --- | --- | --- |
| `--mass` | Black hole mass in solar masses. | `10` |
| `--radius` | Initial radius in Schwarzschild radii. Must be > 1. | `10` |
| `--vr` | Initial radial velocity as a fraction of *c*. Negative values point inward. | `0` |
| `--vt-factor` | Tangential speed factor relative to the circular orbit speed at `radius`. | `0.95` |
| `--step` | Integration step size in seconds. `0` selects an automatic value tied to the light-crossing time. | `0` |
| `--steps` | Number of integration steps. | `8000` |
| `--horizon-factor` | Multiple of the Schwarzschild radius that counts as capture. | `1.01` |
| `--output` | Optional CSV destination for trajectory samples. | *(empty)* |
| `--ascii` | Render a coarse ASCII map of the path. | `false` |
| `--render` | PNG destination for photorealistic render. | *(empty)* |
| `--render-width` / `--render-height` | Output resolution in pixels. | `1280` / `720` |
| `--render-tilt` | Disk inclination angle for the render (degrees). | `20` |
| `--render-exposure` | Exposure multiplier for tone mapping. | `1.8` |
| `--render-step` / `--render-max-steps` | Step size (in Schwarzschild radii) and iteration cap for ray integration. | `0.02` / `2800` |
| `--render-far` | Distance (in Schwarzschild radii) where rays sample the star background. | `90` |
| `--render-horizon` | Multiple of Schwarzschild radius treated as capture in render. | `1.03` |
| `--render-camera-dist` / `--render-camera-height` / `--render-camera-az` | Camera placement controls in Schwarzschild radii and degrees. | `22` / `6` / `18` |

Each CSV row contains time, polar coordinates, velocities, and Cartesian projections so you can plot the trajectory in Python, gnuplot, or any spreadsheet.

## Output example

```
steps executed: 5999
duration: 0.003s
final radius: 11.996 Rs
captured: false

                             ooooooooooooo                             
                         oooo             oooo                         
                      ooo                   ooo                        
                    ooo                       ooo                      
                  ooo                           ooo                    
                 oo                               oo                   
                oo                                 oo                  
               oo                                   oo                 
              oo                                     oo                
              o                                       o                
             o                    *****                o               
            oo                 ***     ***             oo              
            oo               ***         ***           oo              
           oo               **             **          oo              
           o               **               **         o               
           o               *                 *         o               
           o               *      X          *         o               
           o               *                 *         o               
           oo              **               **         oo              
            oo              **             **          oo              
            oo               ***         ***           oo              
             o                 ***     ***             o               
              o                    *****                o              
              oo                                     oo                
               oo                                   oo                 
                oo                                 oo                  
                 oo                               oo                   
                  ooo                           ooo                    
                    ooo                       ooo                      
                      ooo                   ooo                        
                         oooo             oooo                         
                             ooooooooooooo                             
```

`X` marks the black hole and `o` approximates the event horizon. `*` tracks the particle trajectory.

## Testing

Unit tests verify the orbital solver, energy conservation, and capture detection. Run them with:

```bash
go test ./...
```

The repository has no external dependencies beyond the Go toolchain.
