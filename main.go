package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Cell struct {
	points int
}

var BYTES_PER_PIXEL int = 3

func toBytes(cells *[]Cell, width, height int) []byte {
	out := make([]byte, len(*cells)*BYTES_PER_PIXEL)
	for i := 0; i < len(out); i += 3 {
		for j := 0; j < BYTES_PER_PIXEL; j++ {
			index := i + j
			if j == 0 {
				out[index] = 255
			} else {
				out[index] = 0
			}
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
	cells := make([]Cell, WIDTH*HEIGHT)
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
