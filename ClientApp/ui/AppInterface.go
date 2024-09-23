package ui

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/virtouso/WhatsappClientServer/ClientApp/whatsapp"
	"time"

	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
	_ "image/jpeg" // Support for JPEG images
	"log"
	"os"
)

func RunUi2() {
	go func() {
		window := new(app.Window)
		err := run2(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var img image.Image

func run2(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	var messageButton widget.Clickable
	var readQrButton widget.Clickable
	var exitButton widget.Clickable
	var textBox widget.Editor
	var img image.Image

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if messageButton.Clicked(gtx) {
				fmt.Println("message Button clicked!")

			}

			if readQrButton.Clicked(gtx) {
				fmt.Println("qr Button clicked!")
				//img = DisplayImage(img, "d://x.jpg")
				go whatsapp.Init()
				//img = DisplayImage(img, "d://x.jpg")

				ticker := time.NewTicker(1 * time.Second)
				//	defer ticker.Stop()

				go Display(ticker)
			}
			if exitButton.Clicked(gtx) {
				fmt.Println("exit Button clicked!")
				os.Exit(0)
			}
			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceStart,
				Alignment: 2,
			}.Layout(gtx,

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						size := image.Point{X: 200, Y: 200} // Box size (100x100)

						if img != nil {

							imgBounds := img.Bounds().Size()
							scaleX := float32(size.X) / float32(imgBounds.X)
							scaleY := float32(size.Y) / float32(imgBounds.Y)
							scaleOp := f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(scaleX, scaleY))
							op.Affine(scaleOp).Add(gtx.Ops)
							imgOp := paint.NewImageOp(img)
							imgOp.Add(gtx.Ops)
							paint.PaintOp{}.Add(gtx.Ops)
						} else {

							rect := image.Rect(0, 0, size.X, size.Y)
							paint.FillShape(gtx.Ops, color.NRGBA{R: 0, G: 100, B: 200, A: 255}, clip.Rect(rect).Op())
						}

						return layout.Dimensions{Size: size}
					},
				),

				layout.Rigid(
					layout.Spacer{Height: unit.Dp(10)}.Layout,
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						editor := material.Editor(theme, &textBox, "Enter text...")
						return editor.Layout(gtx)
					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &readQrButton, "Read Whatsapp Qr")
						return btn.Layout(gtx)
					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &messageButton, "SendMessage")
						return btn.Layout(gtx)
					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &exitButton, "Exit")
						return btn.Layout(gtx)
					},
				),

				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func Display(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			img = DisplayImage(img, "d://x.jpg")

		}
	}
}

func DisplayImage(img image.Image, dir string) image.Image {
	f, err := os.Open(dir)
	if err == nil {
		defer f.Close()
		img, _, _ = image.Decode(f)
	}
	return img
}
