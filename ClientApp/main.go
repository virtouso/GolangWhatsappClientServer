package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"

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

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	var button widget.Clickable
	var exitButton widget.Clickable
	var textBox widget.Editor
	var img image.Image

	f, err := os.Open("D://x.jpg")
	if err == nil {
		defer f.Close()
		img, _, _ = image.Decode(f)
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if button.Clicked(gtx) {
				fmt.Println("Button clicked!")
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
							// Get the image size
							imgBounds := img.Bounds().Size()

							// Calculate scaling factors
							scaleX := float32(size.X) / float32(imgBounds.X)
							scaleY := float32(size.Y) / float32(imgBounds.Y)

							// Create an Affine2D transformation for scaling
							scaleOp := f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(scaleX, scaleY))

							// Apply the scaling transformation using op.Affine
							op.Affine(scaleOp).Add(gtx.Ops)

							// Render the scaled image
							imgOp := paint.NewImageOp(img)
							imgOp.Add(gtx.Ops)
							paint.PaintOp{}.Add(gtx.Ops)
						} else {
							// Fallback: Draw a blue box if the image couldn't be loaded
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

				// Button
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &button, "SendMessage")
						return btn.Layout(gtx)
					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &exitButton, "Exit")
						return btn.Layout(gtx)
					},
				),

				// Spacer below the button
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			e.Frame(gtx.Ops)
		}
	}
}
