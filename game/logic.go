package game

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func BallMove(dc *gg.Context) {
	var powerX float64
	var powerY float64

	xPosition, yPosition := startPositionOfTheBall(dc)

	angle := dialogCLI()

	powerX, powerY = angleToXY(angle)

	for i := 0; i <= power*10 && !winningState; i++ {
		xPosition, yPosition = moveOneStep(xPosition, yPosition, powerX, powerY)

		// Color changing and drawing a point
		if i%pointStep == 0 {
			pointColor()
			point(dc, xPosition, yPosition, red, green, blue)
		}

		targetHitCheck(xPosition, yPosition, dc)
		wallHitCheck(xPosition, yPosition, dc)
		if i == power*10 {
			ball(dc, xPosition, yPosition, 255, 255, 255)
		}
	}

}

// Calculating direction coordinates, using a right triangle
func angleToXY(angle float64) (float64, float64) {
	var x float64
	var y float64

	angle = angle * 0.0174533 // Convert to radian

	x = math.Cos(angle)
	y = math.Sin(angle)
	return x, y
}

func win(dc *gg.Context, holeNum int) {
	dc.DrawString("VICTORY", 500, 500)
	dc.SetRGB255(0, 0, 0)
	dc.Fill()
	winningState = true
	switch holeNum {
	case 1:
		hole(dc, borderWidth, minY, 255, 128, 0)
		break
	case 2:
		hole(dc, (maxX+borderWidth)/2, minY, 255, 128, 0)
		break
	case 3:
		hole(dc, maxX, minY, 255, 128, 0)
		break
	case 4:
		hole(dc, borderWidth, maxY, 255, 128, 0)
		break
	case 5:
		hole(dc, (maxX+borderWidth)/2, maxY, 255, 128, 0)
		break
	case 6:
		hole(dc, maxX, maxY, 255, 128, 0)
		break
	}
	fmt.Println("VICTORY")
}

func dialogCLI() float64 {
	var angle float64
	j := true
	for j {
		fmt.Println("Enter impact force from 1 to 255:")
		fmt.Scan(&power)
		if power < 1 || power > maxPower {
			fmt.Println("Wrong data. Try again.")
		} else {
			j = false
		}
	}

	k := true
	for k {
		fmt.Println("Enter direction of impact (from 0 to 360 degrees):")
		fmt.Scan(&angle)
		if angle < 0 || angle > maxAngle {
			fmt.Println("Wrong data. Try again.")
		} else {
			k = false
		}
	}
	return angle
}

func startPositionOfTheBall(dc *gg.Context) (float64, float64) {
	var xPosition float64
	var yPosition float64
	rand.Seed(time.Now().UnixNano())
	xPosition = (minX + borderWidth) + rand.Float64()*((maxX+borderWidth)-(minX+borderWidth))
	yPosition = minY + rand.Float64()*(maxY-minY)
	ball(dc, xPosition, yPosition, 255, 255, 255)
	dc.SavePNG("out/out.png") // Start position preview
	return xPosition, yPosition
}

func moveOneStep(xPosition float64, yPosition float64, powerX float64, powerY float64) (float64, float64) {
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
	return xPosition, yPosition
}

func targetHitCheck(xPosition float64, yPosition float64, dc *gg.Context) {
	switch {
	case xPosition < minX+2*ballRadius && yPosition < minY+2*ballRadius:
		win(dc, 1)
		break

	case xPosition > (maxX+borderWidth)/2-2*ballRadius && xPosition < (maxX+borderWidth)/2+2*ballRadius && yPosition < minY+2*ballRadius:
		win(dc, 2)
		break

	case xPosition > maxX-2*ballRadius && yPosition < minY+2*ballRadius:
		win(dc, 3)
		break

	case xPosition < minX+2*ballRadius && yPosition > maxY-2*ballRadius:
		win(dc, 4)
		break

	case xPosition > (maxX+borderWidth)/2-2*ballRadius && xPosition < (maxX+borderWidth)/2+2*ballRadius && yPosition > maxY-2*ballRadius:
		win(dc, 5)
		break

	case xPosition > maxX-2*ballRadius && yPosition > maxY-2*ballRadius:
		win(dc, 6)
		break
	}
}

func wallHitCheck(xPosition float64, yPosition float64, dc *gg.Context) {
	if xPosition > maxX-ballRadius || xPosition < minX+ballRadius {
		directionX = directionX * -1
		ball(dc, xPosition, yPosition, 255, 255, 0)
	}

	if yPosition > maxY-ballRadius || yPosition < minY+ballRadius {
		directionY = directionY * -1
		ball(dc, xPosition, yPosition, 255, 255, 0)
	}
}
