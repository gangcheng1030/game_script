package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
	"time"
)

func main() {
	opencv()
}

func opencv() {
	time.Sleep(5 * time.Second)
	img := robotgo.CaptureImg(578, 414, 240, 47)
	err := robotgo.Save(img, "img.png")
	if err != nil {
		fmt.Println(err)
	}
	img1 := robotgo.CaptureImg(393, 468, 35, 35)
	err = robotgo.Save(img1, "chuhuoyuanzhi.png")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Print("gcv find image: ")
	//gcv.FindImg(img1, img)
	//fmt.Println()
	//
	res := gcv.FindAllImg(img1, img)
	fmt.Println(res)
	//x, y := res[0].TopLeft.X, res[0].TopLeft.Y
	//robotgo.Move(x, y-rand.Intn(5))
	//robotgo.MilliSleep(100)
	//robotgo.Click()
	//
	//res = gcv.FindAll(img1, img) // use find template and sift
	//fmt.Println("find all: ", res)
	//res1 := gcv.Find(img1, img)
	//fmt.Println("find: ", res1)
	//
	//img2, _, _ := robotgo.DecodeImg("test_001.png")
	//x, y = gcv.FindX(img2, img)
	//fmt.Println(x, y)
}
