package main

import (
	"image"
	"time"

	vidio "github.com/AlexEidt/Vidio"
	"github.com/nsf/termbox-go"
)


func main(){
	frameTime := time.Second/24

	err := termbox.Init() 
	if err != nil{
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	//fmt.Println(termbox.Output256)
	//ColorMode256 = ColorMode(256)
	//termbox.Close()//------------------------------------------
	video, err := vidio.NewVideo("bad-apple.mp4")
	
	img := image.NewRGBA(image.Rect(0, 0, video.Width(), video.Height()))
	video.SetFrameBuffer(img.Pix)

	frame := 0
	for video.Read() {
		frameTimeStart := time.Now()

		//f, _ := os.Create(fmt.Sprintf("%d.jpg", frame))
		//jpeg.Encode(f, img, nil)
		//f.Close()
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++{
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x,y).RGBA()
				//fmt.Println(math.Sqrt(math.Pow(0.299*rN, 2) + math.Pow(0.587*gN,2) + math.Pow(0.114*bN, 2)))
				//termbox.Close()
				//fmt.Println(rN, gN, bN)
				termbox.SetBg(x*2, y, termbox.Attribute(termColor(uint16(r/256), uint16(g/256), uint16(b/256))))
				termbox.SetBg(x*2-1, y, termbox.Attribute(termColor(uint16(r/256), uint16(g/256), uint16(b/256))))
				// if math.Sqrt(math.Pow(0.299*rN, 2) + math.Pow(0.587*gN,2) + math.Pow(0.114*bN, 2)) > 255{
				// 	//fmt.Println(math.Sqrt(math.Pow(0.299*rN, 2) + math.Pow(0.587*gN,2) + math.Pow(0.114*bN, 2)))
				// 	termbox.SetBg(x*2-1, y, termbox.ColorWhite)
				// 	termbox.SetBg(x*2 , y, termbox.ColorWhite)
				// 	//fmt.Println(x, y)
				// } else {
				// 	termbox.SetBg(x*2-1, y, termbox.ColorBlack)
				// 	termbox.SetBg(x*2, y, termbox.ColorBlack)
				// }
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
