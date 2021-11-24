package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func writeCircle(buffer *[]uint, width, height int, foreground, background uint, radius int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := width - x*2 - 1
			dy := height - y*2 - 1
			r := radius * radius
			if (dx*dx)+(dy*dy) <= (r) {
				(*buffer)[y*width+x] = foreground
			} else {
				(*buffer)[y*width+x] = background
			}
		}
	}
}

var BYTES_PER_PIXEL int = 3

func toBytes(cells *[]uint, width, height int) []byte {
	var out []byte
	for i := 0; i < len((*cells)); i++ {
		pixel := (*cells)[i]
		for j := 0; j < BYTES_PER_PIXEL; j++ {
			shiftValue := 0
			// TODO: ref this
			if j == 0 {
				shiftValue = 16
			} else if j == 1 {
				shiftValue = 8
			} else {
				shiftValue = 0
			}
			out = append(out, byte(pixel>>uint(shiftValue)&0xff))
		}
	}
	return out
}

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Sandpile")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	WIDTH := 800
	HEIGHT := 600
	var FOREGROUND uint = 0xff0000
	var BACKGROUND uint = 0x0000ff
	cells := make([]uint, WIDTH*HEIGHT)
	writeCircle(&cells, WIDTH, HEIGHT, FOREGROUND, BACKGROUND, HEIGHT)
	buff, err := gdk.PixbufNewFromData(toBytes(&cells, WIDTH, HEIGHT), gdk.COLORSPACE_RGB, false, 8, WIDTH, HEIGHT, BYTES_PER_PIXEL*WIDTH)
	if err != nil {
		panic(err)
	}
	image, err := gtk.ImageNewFromPixbuf(buff)
	if err != nil {
		panic(err)
	}
	win.Add(image)
	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}
