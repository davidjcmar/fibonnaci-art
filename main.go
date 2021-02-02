package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"

	"github.com/fogleman/gg"
)

/*
Working for small modulo, but failing for larger modulo.
The p[] slice is getting off somewhere, maybe they should all be uint64?
That doesn't seem right though since f[]%m is bounded by m...
*/

const maxSlice = 1000
const moreSlice = 250

type coord struct {
	x, y float64
}

func generateFibonacci(l int) []uint64 {
	f := make([]uint64, 0, maxSlice)
	f = f[0:l]
	f[0] = 0
	f[1] = 1
	for i := 2; i < l; i++ {
		f[i] = f[i-1] + f[i-2]
		fmt.Printf("%d\n", f[i])
	}
	return f
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
func main() {
	modulo := uint(10)
	width := 1000
	height := 1000
	circleCenterW := float64(width / 2)
	circleCenterH := float64(height / 2)
	radius := float64(400)
	lineWdith := float64(5)

	pp, ppErr := PisanoPeriod(modulo)
	if ppErr != nil {
		fmt.Printf("Handle error ya bish")
	}
	coords, _ := convertPeriodToXy(modulo, pp, circleCenterW, circleCenterH, radius)
	fmt.Printf("%v", coords)
	dc := gg.NewContext(width, height)
	dc.DrawCircle(circleCenterW, circleCenterH, radius)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(5)
	dc.Fill()
	dc.DrawCircle(circleCenterW, circleCenterH, radius-lineWdith)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	for i, v := range(coords) {
		dc.DrawLine
	}
	dc.SavePNG("out.png")
}
