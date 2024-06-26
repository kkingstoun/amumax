package engine

import (
	"github.com/MathieuMoalic/amumax/cuda"
	"github.com/MathieuMoalic/amumax/data"
)

var (
	TotalShift, TotalYShift                    float64                        // accumulated window shift (X and Y) in meter
	ShiftMagL, ShiftMagR, ShiftMagU, ShiftMagD data.Vector                    // when shifting m, put these value at the left/right edge.
	ShiftM, ShiftGeom, ShiftRegions            bool        = true, true, true // should shift act on magnetization, geometry, regions?
)

func init() {
	DeclFunc("Shift", Shift, "Shifts the simulation by +1/-1 cells along X")
	DeclVar("ShiftMagL", &ShiftMagL, "Upon shift, insert this magnetization from the left")
	DeclVar("ShiftMagR", &ShiftMagR, "Upon shift, insert this magnetization from the right")
	DeclVar("ShiftMagU", &ShiftMagU, "Upon shift, insert this magnetization from the top")
	DeclVar("ShiftMagD", &ShiftMagD, "Upon shift, insert this magnetization from the bottom")
	DeclVar("ShiftM", &ShiftM, "Whether Shift() acts on magnetization")
	DeclVar("ShiftGeom", &ShiftGeom, "Whether Shift() acts on geometry")
	DeclVar("ShiftRegions", &ShiftRegions, "Whether Shift() acts on regions")
	DeclVar("TotalShift", &TotalShift, "Amount by which the simulation has been shifted (m).")
}

// position of the window lab frame
func GetShiftPos() float64  { return -TotalShift }
func GetShiftYPos() float64 { return -TotalYShift }

// shift the simulation window over dx cells in X direction
func Shift(dx int) {
	TotalShift += float64(dx) * GetMesh().CellSize()[X] // needed to re-init geom, regions
	if ShiftM {
		shiftMag(M.Buffer(), dx) // TODO: M.shift?
	}
	if ShiftRegions {
		Regions.shift(dx)
	}
	if ShiftGeom {
		geometry.shift(dx)
	}
	M.normalize()
}

func shiftMag(m *data.Slice, dx int) {
	m2 := cuda.Buffer(1, m.Size())
	defer cuda.Recycle(m2)
	for c := 0; c < m.NComp(); c++ {
		comp := m.Comp(c)
		cuda.ShiftX(m2, comp, dx, float32(ShiftMagL[c]), float32(ShiftMagR[c]))
		data.Copy(comp, m2) // str0 ?
	}
}

// shift the simulation window over dy cells in Y direction
func YShift(dy int) {
	TotalYShift += float64(dy) * GetMesh().CellSize()[Y] // needed to re-init geom, regions
	if ShiftM {
		shiftMagY(M.Buffer(), dy)
	}
	if ShiftRegions {
		Regions.shiftY(dy)
	}
	if ShiftGeom {
		geometry.shiftY(dy)
	}
	M.normalize()
}

func shiftMagY(m *data.Slice, dy int) {
	m2 := cuda.Buffer(1, m.Size())
	defer cuda.Recycle(m2)
	for c := 0; c < m.NComp(); c++ {
		comp := m.Comp(c)
		cuda.ShiftY(m2, comp, dy, float32(ShiftMagU[c]), float32(ShiftMagD[c]))
		data.Copy(comp, m2) // str0 ?
	}
}
