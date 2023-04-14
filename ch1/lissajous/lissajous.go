// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 调色板
// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}

// 调色板中的颜色索引
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.Create("a.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous: %v\n", err)
		return
	}
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // 动画帧的数量
		delay   = 8     // 两帧之间的延迟 以 10ms为一个1单位
	)

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // 创建GIF对象，指定帧数
	phase := 0.0                        // phase difference
	// 生成动画帧
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 生成2*size+1长宽的二维矩阵
		img := image.NewPaletted(rect, palette)      // 创建rect相同大小的调色板
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
