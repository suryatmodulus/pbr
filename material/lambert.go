package material

import (
	"math"
	"math/rand"

	"github.com/hunterloftis/pbr/geom"
	"github.com/hunterloftis/pbr/rgb"
)

// TODO: Oren-Nayar for roughness
type Lambert struct {
	Color      rgb.Energy
	Multiplier float64
}

func (l Lambert) Sample(wo geom.Direction, rnd *rand.Rand) (geom.Direction, float64) {
	wi := geom.Up.RandHemiCos(rnd)
	return wi, l.PDF(wi, wo)
}

func (l Lambert) PDF(wi, wo geom.Direction) float64 {
	return wi.Dot(geom.Up) * math.Pi
}

func (l Lambert) Eval(wi, wo geom.Direction) rgb.Energy {
	return l.Color.Scaled(l.Multiplier)
}
