package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

import (
	"log"
	"net/http"
	"time"
)

// 调色板，可以用来改变背景和线条的颜色
var palette = []color.Color{ color.RGBA{255,255,255,math.MaxUint8},
							 color.Black,
							 color.RGBA{0,0,128,math.MaxUint8},
							 color.RGBA{255,0,255,math.MaxUint8},
							 color.RGBA{23,133,124,math.MaxUint8},
							 color.RGBA{134,23,0,math.MaxUint8},
							 color.RGBA{1,124,43,math.MaxUint8}}


const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	// 设立时间种子，生成随机数的
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	// >out.gif 用来输出文件
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的振荡器个数
		res     = 0.001 // 角度分辨率
		size    = 150   // 控制图像画布的大小
		nframes = 128    // 动画中的帧数
		delay   = 8    // 以10ms为单位的帧问题延迟
	)
	freq := rand.Float64() * 3.0 // y 振荡器的频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 不同的相位
	for i,c := 0,1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(c))
		}
		c ++ //修改颜色选择
		if c >= len(palette) {
			c = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 忽略编码错误
}