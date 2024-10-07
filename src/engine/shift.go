package engine

import (
	"github.com/MathieuMoalic/amumax/src/cuda"
	"github.com/MathieuMoalic/amumax/src/data"
)

var (
	totalShift, totalYShift                    float64                        // accumulated window shift (X and Y) in meter
	shiftMagL, shiftMagR, shiftMagU, shiftMagD data.Vector                    // when shifting m, put these value at the left/right edge.
	shiftM, shiftGeom, shiftRegions            bool        = true, true, true // should shift act on magnetization, geometry, regions?
)

// position of the window lab frame
func getShiftPos() float64  { return -totalShift }
func getShiftYPos() float64 { return -totalYShift }

// shift the simulation window over dx cells in X direction
func shift(dx int) {
	totalShift += float64(dx) * getMesh().CellSize()[X] // needed to re-init geom, regions
	if shiftM {
		shiftMag(normMag.Buffer(), dx) // TODO: M.shift?
	}
	if shiftRegions {
		Regions.shift(dx)
	}
	if shiftGeom {
		Geometry.shift(dx)
	}
	normMag.normalize()
}

func shiftMag(m *data.Slice, dx int) {
	m2 := cuda.Buffer(1, m.Size())
	defer cuda.Recycle(m2)
	for c := 0; c < m.NComp(); c++ {
		comp := m.Comp(c)
		cuda.ShiftX(m2, comp, dx, float32(shiftMagL[c]), float32(shiftMagR[c]))
		data.Copy(comp, m2) // str0 ?
	}
}

// shift the simulation window over dy cells in Y direction
func yShift(dy int) {
	totalYShift += float64(dy) * getMesh().CellSize()[Y] // needed to re-init geom, regions
	if shiftM {
		shiftMagY(normMag.Buffer(), dy)
	}
	if shiftRegions {
		Regions.shiftY(dy)
	}
	if shiftGeom {
		Geometry.shiftY(dy)
	}
	normMag.normalize()
}

func shiftMagY(m *data.Slice, dy int) {
	m2 := cuda.Buffer(1, m.Size())
	defer cuda.Recycle(m2)
	for c := 0; c < m.NComp(); c++ {
		comp := m.Comp(c)
		cuda.ShiftY(m2, comp, dy, float32(shiftMagU[c]), float32(shiftMagD[c]))
		data.Copy(comp, m2) // str0 ?
	}
}
