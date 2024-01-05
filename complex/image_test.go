package complex

import (
	"fmt"
	"image"
	"testing"
)

// image 包定义了 Image 接口
//
//	type Image interface {
//	   ColorModel() color.Model
//	   Bounds() Rectangle
//	   At(x, y int) color.Color
//	}
func TestImage(t *testing.T) {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
