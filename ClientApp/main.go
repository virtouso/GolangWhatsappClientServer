package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
	"image/color"

	"log"
	"os"

	"gioui.org/widget"
	"gioui.org/widget/material"
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

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Check if the button was clicked
			if button.Clicked(gtx) {
				fmt.Println("Button clicked!")
			}
			if exitButton.Clicked(gtx) {
				fmt.Println("exit Button clicked!")
			}
			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceStart,
				Alignment: 2,
			}.Layout(gtx,

				// Image (or colored box) above the button
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						size := gtx.Constraints.Max
						size.Y = 100
						size.X = 100
						rect := image.Rect(0, 0, size.X, size.Y)
						paint.FillShape(gtx.Ops, color.NRGBA{R: 0, G: 100, B: 200, A: 255}, clip.Rect(rect).Op())
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
