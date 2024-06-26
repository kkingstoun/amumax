package engine

import (
	"bytes"
	"image"
	"image/jpeg"
	"math"
	"net/http"
	"sync"

	"github.com/MathieuMoalic/amumax/cuda"
	"github.com/MathieuMoalic/amumax/data"
	"github.com/MathieuMoalic/amumax/draw"
)

type Render struct {
	Mutex        sync.Mutex
	quant        Quantity
	comp         string
	layer, scale int
	rescaleBuf   *data.Slice // GPU
	imgBuf       *data.Slice // CPU
	Img          *image.RGBA
}

const (
	maxScale   = 32  // maximum zoom-out setting
	maxImgSize = 512 // maximum render image size
)

// Render image of quantity.
func (g *guistate) ServeRender(w http.ResponseWriter, r *http.Request) {
	g.Render.Mutex.Lock()
	defer g.Render.Mutex.Unlock()

	g.Render.Render()
	jpeg.Encode(w, g.Render.Img, &jpeg.Options{Quality: 100})
}

func (g *guistate) GetRenderedImg(quant, comp string, zlice int) *bytes.Buffer {
	g.Render.quant = g.Quants[quant]
	g.Render.comp = comp
	g.Render.layer = zlice
	g.Render.Mutex.Lock()
	defer g.Render.Mutex.Unlock()

	g.Render.Render()
	buff := bytes.NewBuffer([]byte{})
	jpeg.Encode(buff, g.Render.Img, &jpeg.Options{Quality: 100})
	return buff

}

// rescale and download quantity, save in rescaleBuf
func (ren *Render) download() {
	InjectAndWait(func() {
		if ren.quant == nil { // not yet set, default = m
			ren.quant = &M
		}
		quant := ren.quant
		size := MeshOf(quant).Size()

		// don't slice out of bounds
		renderLayer := ren.layer
		if renderLayer >= size[Z] {
			renderLayer = size[Z] - 1
		}
		if renderLayer < 0 {
			renderLayer = 0
		}

		// scaling sanity check
		if ren.scale < 1 {
			ren.scale = 1
		}
		if ren.scale > maxScale {
			ren.scale = maxScale
		}
		// Don't render too large images or we choke
		for size[X]/ren.scale > maxImgSize {
			ren.scale++
		}
		for size[Y]/ren.scale > maxImgSize {
			ren.scale++
		}

		for i := range size {
			size[i] /= ren.scale
			if size[i] == 0 {
				size[i] = 1
			}
		}
		size[Z] = 1 // selects one layer

		// make sure buffers are there
		if ren.imgBuf.Size() != size {
			ren.imgBuf = data.NewSlice(3, size) // always 3-comp, may be re-used
		}
		buf := ValueOf(quant)
		defer cuda.Recycle(buf)
		if !buf.GPUAccess() {
			ren.imgBuf = Download(quant) // fallback (no zoom)
			return
		}
		// make sure buffers are there (in CUDA context)
		if ren.rescaleBuf.Size() != size {
			ren.rescaleBuf.Free()
			ren.rescaleBuf = cuda.NewSlice(1, size)
		}
		for c := 0; c < quant.NComp(); c++ {
			cuda.Resize(ren.rescaleBuf, buf.Comp(c), renderLayer)
			data.Copy(ren.imgBuf.Comp(c), ren.rescaleBuf)
		}
	})
}

var arrowSize = 16

func (ren *Render) Render() {
	ren.download()
	// imgBuf always has 3 components, we may need just one...
	imgBuf := ren.imgBuf
	comp := ren.comp
	quant := ren.quant
	if comp == "All" {
		Normalize(imgBuf)
	}
	if comp != "All" && quant.NComp() > 1 { // ... if one has been selected by gui
		imgBuf = imgBuf.Comp(compstr[comp])
	}
	if quant.NComp() == 1 { // ...or if the original data only had one (!)
		imgBuf = imgBuf.Comp(0)
	}
	if ren.Img == nil {
		ren.Img = new(image.RGBA)
	}
	draw.On(ren.Img, imgBuf, "auto", "auto", arrowSize)
}

var compstr = map[string]int{"x": 0, "y": 1, "z": 2}

func Normalize(f *data.Slice) {
	a := f.Vectors()
	maxnorm := 0.
	for i := range a[0] {
		for j := range a[0][i] {
			for k := range a[0][i][j] {

				x, y, z := a[0][i][j][k], a[1][i][j][k], a[2][i][j][k]
				norm := math.Sqrt(float64(x*x + y*y + z*z))
				if norm > maxnorm {
					maxnorm = norm
				}

			}
		}
	}
	factor := float32(1 / maxnorm)

	for i := range a[0] {
		for j := range a[0][i] {
			for k := range a[0][i][j] {
				a[0][i][j][k] *= factor
				a[1][i][j][k] *= factor
				a[2][i][j][k] *= factor

			}
		}
	}
}
