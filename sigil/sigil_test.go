package sigil

import (
	"image"
	"reflect"
	"testing"

	"github.com/fogleman/gg"
)

func TestPisanoPeriod(t *testing.T) {
	t.Run("Modulo 2", func(t *testing.T) {
		got, _ := PisanoPeriod(2)
		want := []uint64{0, 1, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("PisanoPeriod(2) = %q, want %q", got, want)
		}
	})

	t.Run("Modulo 10", func(t *testing.T) {
		got, _ := PisanoPeriod(10)
		want := []uint64{0, 1, 1, 2, 3, 5, 8, 3, 1, 4, 5, 9, 4, 3, 7, 0, 7, 7, 4, 1, 5, 6, 1, 7, 8, 5, 3, 8, 1, 9, 0, 9, 9, 8, 7, 5, 2, 7, 9, 6, 5, 1, 6, 7, 3, 0, 3, 3, 6, 9, 5, 4, 9, 3, 2, 5, 7, 2, 9, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("PisanoPeriod(10) = %q, want %q", got, want)
		}
	})
}
/*
func TestMakeSigil(t *testing.T) {
	t.Run("Radius 400, Mod 4, LineWidth 3, 1000 x 1000", func(t *testing.T) {
		var got, want image.Image
		got, _ = MakeSigil(4, 1000, 1000, 400, 3)
		want, _ = gg.LoadPNG("./radius400mod4lw3_1000x1000.png")
		if got != want {
			t.Errorf("MakeSigil() = %q\nwant = %q", got, want)
		}
	})
}
*/