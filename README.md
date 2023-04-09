# refraction_model
A refraction model of a single ray on a multiple horizontal earth layers. This was my Seismic Wave class assignment back then during college and I rewrote this because I'm bored.

## Limitations:
- Assumes the earth layers are horizontal
- Attenuation is not taken into consideration
- May not be able to calculate large numbers

## Usage
1. `go get module github.com/wuflenso/refraction_model`
2. `import "github.com/wuflenso/refraction_model"` to your Go code
2. Implement the methods

## Examples
See [examples](/examples/trace_ray_refraction.go). You can also check [refraction_plot repository](https://github.com/wuflenso/refraction_plot) to see the example when we plot the results to a graph.
