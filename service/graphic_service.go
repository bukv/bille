package service

import "github.com/fogleman/gg"

func ImageСreationProcess(angle float64) {
	dc := gg.NewContext(1000, 1000)

	table(dc)

	ballMove(dc, angle)

	dc.SavePNG("images/out.png")
}

func hole(dc *gg.Context, xPos float64, yPos float64, red int, green int, blue int) {
	dc.DrawCircle(xPos, yPos, 2*ballRadius)
	dc.SetRGB255(red, green, blue)
	dc.Fill()
}

func table(dc *gg.Context) {
	dc.DrawRectangle(0, minY-50, maxX, heightTable)
	dc.SetRGB255(102, 51, 0)
	dc.Fill()
	dc.DrawRectangle(50, minY, maxX-2*borderWidth, heightTable-2*borderWidth)
	dc.SetRGB255(0, 153, 51)
	dc.Fill()
	hole(dc, borderWidth, minY, 0, 0, 0)
	hole(dc, maxX/2, minY, 0, 0, 0)
	hole(dc, maxX-borderWidth, minY, 0, 0, 0)
	hole(dc, borderWidth, maxY, 0, 0, 0)
	hole(dc, maxX/2, maxY, 0, 0, 0)
	hole(dc, maxX-borderWidth, maxY, 0, 0, 0)
}

func ball(dc *gg.Context, xPos float64, yPos float64, red int, green int, blue int) {
	dc.DrawCircle(xPos, yPos, ballRadius)
	dc.SetRGB255(red, green, blue)
	dc.Fill()
}

func ballColor() {
	red = red - (255 / power)
	if red < 0 {
		red = 0
	}
	blue = blue + (255 / power)
	if blue > 255 {
		blue = 255
	}
}
