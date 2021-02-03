package sigil

import (
	"errors"
	"image"
	"math"
	"math/big"
	"reflect"

	"github.com/fogleman/gg"
)

const maxSlice = 1000
const moreSlice = 250

type coord struct {
	x, y float64
}

func PisanoPeriod(m uint) ([]uint64, error) {
	currSlice := moreSlice
	f := make([]big.Int, currSlice, maxSlice)
	p := make([]uint64, currSlice, maxSlice)
	cycle := 0
	for i := 0; i < maxSlice; i++ {
		if i == currSlice {
			currSlice += moreSlice
			f = f[:currSlice]
			p = p[:currSlice]
		}
		// fibonnaci sequence
		if i == 0 || i == 1 {

			f[i] = *big.NewInt(int64(i))
		} else {
			//f[i] = f[i-1] + f[i-2]
			f[i] = *big.NewInt(int64(0)).Add(&f[i-1], &f[i-2])
		}
		//p[i] = uint(f[i] % uint64(m))
		p[i] = big.NewInt(int64(0)).Mod(&f[i], big.NewInt(int64(m))).Uint64()
		// requirements for beginning of pisano period
		if i > 3 && p[i] == 1 && p[i-1] == 1 && p[i-2] == 0 {
			cycle = i - 2
		}
		if i%2 == 0 && cycle == i/2 && cycle != 0 {
			/*
				fmt.Printf("%v\n", f[i])
				fmt.Printf("%v\n", big.NewInt(int64(0)).Mod(&f[i], big.NewInt(int64(m))).Uint64())
				fmt.Printf("First Cycle: %v\n", p[0:cycle])
				fmt.Printf("Second Cycle: %v\n", p[cycle:i])
			*/
			if reflect.DeepEqual(p[0:cycle], p[cycle:i]) {
				return p[0:cycle], nil
			} else {
				cycle = 0
			}
		}
	}
	return nil, errors.New("No period")
}

func convertPeriodToXy(m uint, p []uint64, cx, cy, r float64) ([]coord, error) {
	coords := make([]coord, len(p))
	for i, v := range p {
		a := (2 * math.Pi) / float64(m)
		x := cx + (r * math.Cos(float64(a*float64(v))))
		y := cy + (r * math.Sin(float64(a*float64(v))))
		coords[i].x, coords[i].y = x, y
	}
	return coords, nil
}

func MakeSigil(modulo uint, width, height int, radius, lineWidth float64) (image.Image, error) {

	circleCenterW := float64(width / 2)
	circleCenterH := float64(height / 2)

	pp, err := PisanoPeriod(modulo)
	if err != nil {
		return nil, err
	}
	coords, _ := convertPeriodToXy(modulo, pp, circleCenterW, circleCenterH, radius)
	//fmt.Printf("%v", coords)
	dc := gg.NewContext(width, height)
	dc.DrawCircle(circleCenterW, circleCenterH, radius+lineWidth)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(lineWidth)
	dc.Fill()
	dc.DrawCircle(circleCenterW, circleCenterH, radius)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.SetRGB(.5, 0, 1)
	dc.RotateAbout(math.Pi/2, circleCenterW, circleCenterH)
	for i := range coords {
		if i == len(coords)-1 {
			break
		}
		dc.DrawLine(coords[i].x, coords[i].y, coords[i+1].x, coords[i+1].y)
		dc.Stroke()
	}
	return dc.Image(), nil
}
