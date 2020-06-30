package service

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

var xPosition float64
var yPosition float64

//calculating direction coordinates using a right triangle
func angleToXY(angle float64) (float64, float64) {
	var x float64
	var y float64

	angle = angle * 0.0174533 //convert to radian

	x = (2 * ballRadius) * math.Cos(angle)
	y = (2 * ballRadius) * math.Sin(angle)

	return x, y
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

func ballStartPositionGeneration(dc *gg.Context) {
	rand.Seed(time.Now().UnixNano())
	xPosition = (minX + borderWidth) + rand.Float64()*((maxX+borderWidth)-(minX+borderWidth))
	yPosition = minY + rand.Float64()*(maxY-minY)
	ball(dc, xPosition, yPosition, red, green, blue)
	dc.SavePNG("images/out.png") //start position preview
}

func ballMove(dc *gg.Context, angle float64) {
	//var angle float64
	var powerX float64
	var powerY float64

	ballStartPositionGeneration(dc)
	/*
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
	*/

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
