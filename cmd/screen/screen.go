package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		x, y := robotgo.GetMousePos()
		fmt.Println("pos: ", x, y)
		robotgo.Sleep(3)
	}
	// qq
	// 121 96; 1198 703  窗口
	// 586 453;  650 463  进本
	// 633 425, 688 437  组队
	// sogou
	//593, 445; 658 454   进本
	//643, 417; 692, 430  组队
	// t360
	// 131 103; 1201 706  窗口
	// 640 430; 690 441  组队
	// 594 458; 656 468  进本
	// chrome
	// 128 102; 1197 705 窗口
	// 636 430; 690 443  组队
	// 588 457; 653 467  进本
	//color := robotgo.GetPixelColor(100, 200)
	//fmt.Println("color---- ", color)
	////
	//sx, sy := robotgo.GetScreenSize()
	//fmt.Println("get screen size: ", sx, sy)
	//
	//bit := robotgo.CaptureScreen(10, 10, 30, 30)
	//defer robotgo.FreeBitmap(bit)
	//
	//img := robotgo.ToImage(bit)
	//imgo.Save("test.png", img)
	//
	//num := robotgo.DisplaysNum()
	//for i := 0; i < num; i++ {
	//	robotgo.DisplayID = i
	//	img1 := robotgo.CaptureImg()
	//	path1 := "save_" + strconv.Itoa(i)
	//	robotgo.Save(img1, path1+".png")
	//	robotgo.SaveJpeg(img1, path1+".jpeg", 50)
	//
	//	img2 := robotgo.CaptureImg(10, 10, 20, 20)
	//	robotgo.Save(img2, "test_"+strconv.Itoa(i)+".png")
	//}
}
