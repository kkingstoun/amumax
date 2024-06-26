package engine

import (
	"fmt"

	"github.com/MathieuMoalic/amumax/cuda"
	"github.com/MathieuMoalic/amumax/data"
	"github.com/MathieuMoalic/amumax/util"
)

func sOneRegion(q Quantity, r int) *sOneReg {
	util.Argument(q.NComp() == 1)
	return &sOneReg{oneReg{q, r}}
}

func vOneRegion(q Quantity, r int) *vOneReg {
	util.Argument(q.NComp() == 3)
	return &vOneReg{oneReg{q, r}}
}

type sOneReg struct{ oneReg }

func (q *sOneReg) Average() float64 { return q.average()[0] }

type vOneReg struct{ oneReg }

func (q *vOneReg) Average() data.Vector { return unslice(q.average()) }

// represents a new quantity equal to q in the given region, 0 outside.
type oneReg struct {
	parent Quantity
	region int
}

func inRegion(q Quantity, region int) Quantity {
	return &oneReg{q, region}
}

func (q *oneReg) NComp() int             { return q.parent.NComp() }
func (q *oneReg) Name() string           { return fmt.Sprint(NameOf(q.parent), ".region", q.region) }
func (q *oneReg) Unit() string           { return UnitOf(q.parent) }
func (q *oneReg) Mesh() *data.Mesh       { return MeshOf(q.parent) }
func (q *oneReg) EvalTo(dst *data.Slice) { EvalTo(q, dst) }

// returns a new slice equal to q in the given region, 0 outside.
func (q *oneReg) Slice() (*data.Slice, bool) {
	src := ValueOf(q.parent)
	defer cuda.Recycle(src)
	out := cuda.Buffer(q.NComp(), q.Mesh().Size())
	cuda.RegionSelect(out, src, Regions.Gpu(), byte(q.region))
	return out, true
}

func (q *oneReg) average() []float64 {
	slice, r := q.Slice()
	if r {
		defer cuda.Recycle(slice)
	}
	avg := sAverageUniverse(slice)
	sDiv(avg, Regions.volume(q.region))
	return avg
}

func (q *oneReg) Average() []float64 { return q.average() }

// slice division
func sDiv(v []float64, x float64) {
	for i := range v {
		v[i] /= x
	}
}
