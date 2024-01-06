package main

import (
	"image"
	"math"
	"time"

	vidio "github.com/AlexEidt/Vidio"
	"github.com/nsf/termbox-go"
)

type color struct {
	r uint16 
	g uint16
	b uint16
}

func main(){
	frameTime := time.Second/24

	err := termbox.Init() 
	if err != nil{
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	video, err := vidio.NewVideo("badApple.mp4")
	if err != nil{
		panic(err)
	}
	img := image.NewRGBA(image.Rect(0, 0, video.Width(), video.Height()))
	video.SetFrameBuffer(img.Pix)

	frame := 0

	blockSizeX := 17/16//scaleSize * float64(img.Bounds().Max.X)
	blockSizeY := 11/5//scaleSize * float64(img.Bounds().Max.Y)
	for video.Read() {
		frameTimeStart := time.Now()
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y+=int(blockSizeY) {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x+=int(blockSizeX) {

				colorList := []color{}
				for yS := 0; yS < int(blockSizeY); yS++{
					for xS := 0; xS < int(blockSizeX); xS++{
						r, g, b, _ := img.At(xS+x, yS+y).RGBA()
						color := color{uint16(r)/255, uint16(g)/255, uint16(b)/255}
						colorList = append(colorList, color)
					}
				}
				colorAverage := averageColorList(colorList)
				r := colorAverage.r
				g := colorAverage.g
				b := colorAverage.b

				if y/blockSizeY % 2 != 0{
					termbox.SetCell(x/int(blockSizeX), y/int(blockSizeY)-int(math.Floor(float64(y/int(blockSizeY)*1/2))), '▀',termbox.Attribute(termColor(uint16(r), uint16(g), uint16(b))), termbox.ColorBlue)
				} else {
					termbox.SetBg(x/int(blockSizeX), int(float32(y/int(blockSizeY))*float32(0.5)), termbox.Attribute(termColor(uint16(r), uint16(g), uint16(b))))
				}
			}
		}
		termbox.Flush()
		
		frame++
		// if frame == (24 * 20){
		// 	termbox.Close()
		// }
		elapsedTime := time.Since(frameTimeStart)
		if (elapsedTime < frameTime){
			time.Sleep((frameTime - elapsedTime))
		}

	}
	
}

// func unit8(g uint32) {
// 	panic("unimplemented")
// }

func termColor(r, g, b uint16) uint16 {
	rterm := (((r * 5) + 127) / 255) * 36
	gterm := (((g * 5) + 127) / 255) * 6
	bterm := (((b * 5) + 127) / 255)

	return rterm + gterm + bterm + 16 + 1 // termbox default color offset
}

func averageColorList(colorList []color) color {
	var r, g, b uint16
	for _, i := range colorList{
		r += i.r
		g += i.g
		b += i.b
	}
	color := color{r/uint16(len(colorList)), g/uint16(len(colorList)), b/uint16(len(colorList))}
	return color
}
