package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

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
	pixelBuffer, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, false, 8, WIDTH, HEIGHT)
	if err != nil {
		panic(err)
	}
	image, err := gtk.ImageNewFromPixbuf(pixelBuffer)
	if err != nil {
		panic(err)
	}
	win.Add(image)
	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}
