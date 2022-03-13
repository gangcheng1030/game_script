package main

func main() {
	//for {
	//	x, y := robotgo.GetMousePos()
	//	fmt.Println("pos: ", x, y)
	//	robotgo.Sleep(3)
	//}
	//color := robotgo.GetPixelColor(100, 200)
	//fmt.Println("color---- ", color)
	////
	//sx, sy := robotgo.GetScreenSize()
	//fmt.Println("get screen size: ", sx, sy)
	//
	//robotgo.Sleep(10)
	// 这个capture screen不能用
	//bit := robotgo.CaptureScreen(0, 0, 1360, 768)
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

	// capture screen
	//robotgo.Sleep(10)
	//img, err := screenshot.Capture(0, 0, 1360, 768)
	//if err != nil {
	//	panic(err)
	//}
	//fileName := fmt.Sprintf("%d_%dx%d.png", time.Now().Unix(), 1360, 768)
	//file, _ := os.Create(fileName)
	//defer file.Close()
	//png.Encode(file, img)
}
