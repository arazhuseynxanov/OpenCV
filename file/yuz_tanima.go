package file

import (
	"gocv.io/x/gocv"
	"image/color"
)

func main() {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	blueColor := color.RGBA{R: 0, G: 0, B: 255, A: 255} // Tam opak mavi renk

	if !classifier.Load("/Users/araz/Desktop/openCVGolang/haarcascade_frontalface_default.xml") {
		println("Error reading cascade file")
		return
	}

	webcam, _ := gocv.VideoCaptureDevice(0)
	defer webcam.Close()

	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		webcam.Read(&img)
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			gocv.Rectangle(&img, r, blueColor, 3)
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
