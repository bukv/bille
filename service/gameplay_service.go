package bille

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const maxX = 1000
const minX = 0
const maxY = 800
const minY = 300
const borderWidth = 50
const heightTable = 600
const maxPower = 50
const maxAngle = 360

var ballRadius float64 = 20

var power int

var red int = 255
var green int = 0
var blue int = 0

var directionX int = 1
var directionY int = 1

var winningState bool = false

func hole(dc *gg.Context, xPos float64, yPos float64, red int, green int, blue int) {
	dc.DrawCircle(xPos, yPos, 2*ballRadius)
	dc.SetRGB255(red, green, blue)
	dc.Fill()
}

//calculating direction coordinates using a right triangle
func angleToXY(angle float64) (float64, float64) {
	var x float64
	var y float64

	angle = angle * 0.0174533 //convert to radian

	x = (2 * ballRadius) * math.Cos(angle)
	y = (2 * ballRadius) * math.Sin(angle)

	return x, y
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

func win(dc *gg.Context, holePos int) {
	dc.DrawString("VICTORY", 500, 500)
	dc.SetRGB255(0, 0, 0)
	dc.Fill()
	winningState = true
	switch holePos {
	case 1:
		hole(dc, borderWidth, minY, 255, 128, 0)
		break
	case 2:
		hole(dc, maxX/2, minY, 255, 128, 0)
		break
	case 3:
		hole(dc, maxX-borderWidth, minY, 255, 128, 0)
		break
	case 4:
		hole(dc, borderWidth, maxY, 255, 128, 0)
		break
	case 5:
		hole(dc, maxX/2, maxY, 255, 128, 0)
		break
	case 6:
		hole(dc, maxX-borderWidth, maxY, 255, 128, 0)
		break
	}
	fmt.Println("VICTORY")
}

func ballMove(dc *gg.Context) {
	var xPosition float64
	var yPosition float64
	var angle float64
	var powerX float64
	var powerY float64

	//ball start position generation
	rand.Seed(time.Now().UnixNano())
	xPosition = (minX + borderWidth) + rand.Float64()*((maxX+borderWidth)-(minX+borderWidth))
	yPosition = minY + rand.Float64()*(maxY-minY)
	ball(dc, xPosition, yPosition, red, green, blue)
	dc.SavePNG("out.png") //start position preview

	j := 0
	for j < 1 {
		fmt.Println("Enter impact force from 0 to 50:")
		fmt.Scan(&power)
		if power < 0 || power > maxPower {
			fmt.Println("Wrong data. Try again.")
		} else {
			j = 1
		}
	}

	k := 0
	for k < 1 {
		fmt.Println("Enter direction of impact (from 0 to 360 degrees):")
		fmt.Scan(&angle)
		if angle < 0 || angle > maxAngle {
			fmt.Println("Wrong data. Try again.")
		} else {
			k = 1
		}
	}

	powerX, powerY = angleToXY(angle)

	for i := 0; i < power && !winningState; i++ {
		//next position
		switch directionX {
		case 1:
			xPosition = xPosition + powerX
			break
		case -1:
			xPosition = xPosition - powerX
			break
		}

		switch directionY {
		case 1:
			yPosition = yPosition + powerY
			break
		case -1:
			yPosition = yPosition - powerY
			break
		}

		//color changing and drawing a ball
		ballColor()
		ball(dc, xPosition, yPosition, red, green, blue)

		//check falling into the hole
		switch {
		case xPosition < borderWidth+2*ballRadius && yPosition < minY+2*ballRadius:
			win(dc, 1)
			break

		case xPosition > maxX/2-2*ballRadius && xPosition < maxX/2+2*ballRadius && yPosition < minY+2*ballRadius:
			win(dc, 2)
			break

		case xPosition > maxX-borderWidth-2*ballRadius && yPosition < minY+2*ballRadius:
			win(dc, 3)
			break

		case xPosition < borderWidth+2*ballRadius && yPosition > maxY-2*ballRadius:
			win(dc, 4)
			break

		case xPosition > maxX/2-2*ballRadius && xPosition < maxX/2+2*ballRadius && yPosition > maxY-2*ballRadius:
			win(dc, 5)
			break

		case xPosition > maxX-borderWidth-2*ballRadius && yPosition > maxY-2*ballRadius:
			win(dc, 6)
			break
		}

		//wall hit check
		if xPosition > maxX-borderWidth-ballRadius || xPosition < minX+borderWidth+ballRadius {
			directionX = directionX * -1
			ball(dc, xPosition, yPosition, 255, 255, 0)
		}

		if yPosition > maxY-ballRadius || yPosition < minY+ballRadius {
			directionY = directionY * -1
			ball(dc, xPosition, yPosition, 255, 255, 0)
		}
	}

}

func main() {
	dc := gg.NewContext(1000, 1000)

	table(dc)

	ballMove(dc)

	dc.SavePNG("out.png")
}
