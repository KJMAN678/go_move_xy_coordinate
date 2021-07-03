package main

import (
	"fmt"
	"math"
)

func main() {

	var angle float64
	angle = 90
	var radian float64

	var x, y float64

	for i := 0; i < 8; i++ {
		radian = angle * math.Pi / 180
		fmt.Println("そのまま 角度:", angle, " Cos:", math.Cos(radian), " Sin:", math.Sin(radian))
		fmt.Println("切り捨て 角度:", angle, " Cos:", math.Floor(math.Cos(radian)), " Sin:", math.Floor(math.Sin(radian)))
		fmt.Println("切り上げ 角度:", angle, " Cos:", math.Ceil(math.Cos(radian)), " Sin:", math.Ceil(math.Sin(radian)))
		fmt.Println("四捨五入 角度:", angle, " Cos:", math.Round(math.Cos(radian)), " Sin:", math.Round(math.Sin(radian)))

		x = x + math.Cos(radian)
		y = y + math.Sin(radian)
		angle = angle + 90.0
	}

	fmt.Println("そのまま")
	x, y = 5, -5
	angle = 90
	for i := 0; i < 8; i++ {
		radian = angle * math.Pi / 180
		x = x + math.Cos(radian)*10
		y = y + math.Sin(radian)*10
		fmt.Println("角度:", angle, x, y)
		angle = angle + 90.0
	}

	fmt.Println("切り捨て")
	x, y = 5, -5
	angle = 90
	for i := 0; i < 8; i++ {
		radian = angle * math.Pi / 180
		x = x + math.Floor(math.Cos(radian))*10
		y = y + math.Floor(math.Sin(radian))*10
		fmt.Println("角度:", angle, x, y)
		angle = angle + 90.0
	}

	fmt.Println("切り上げ")
	x, y = 5, -5
	angle = 90
	for i := 0; i < 8; i++ {
		radian = angle * math.Pi / 180
		x = x + math.Ceil(math.Cos(radian))*10
		y = y + math.Ceil(math.Sin(radian))*10
		fmt.Println("角度:", angle, x, y)
		fmt.Println("角度:", angle, x, y)
		angle = angle + 90.0
	}

	fmt.Println("四捨五入")
	x, y = 5, -5
	angle = 90
	for i := 0; i < 8; i++ {
		radian = angle * math.Pi / 180
		x = x + math.Round(math.Cos(radian))*10
		y = y + math.Round(math.Sin(radian))*10
		fmt.Println("角度:", angle, x, y)
		angle = angle + 90.0
	}

}
