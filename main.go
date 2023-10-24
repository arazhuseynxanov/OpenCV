package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {
	blueColor := color.RGBA{R: 0, G: 0, B: 255, A: 255}  
	greenColor := color.RGBA{R: 0, G: 255, B: 0, A: 255} 
	webcam, _ := gocv.OpenVideoCapture(0)
	defer webcam.Close()

	window := gocv.NewWindow("Tap Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	mask := gocv.NewMat()
	defer mask.Close()

	hsv := gocv.NewMat()
	defer hsv.Close()

	for {
		webcam.Read(&img)
		if img.Empty() {
			continue
		}

		gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)
		gocv.InRangeWithScalar(hsv, gocv.NewScalar(0, 20, 70, 0), gocv.NewScalar(20, 255, 255, 0), &mask)

		contours := gocv.FindContours(mask, gocv.RetrievalExternal, gocv.ChainApproxSimple)
		for i := 0; i < contours.Size(); i++ {
			contour := contours.At(i)
			area := gocv.ContourArea(contour)
			if area > 3000 {
				rect := gocv.BoundingRect(contour)
				gocv.Rectangle(&img, rect, greenColor, 2)

				center := image.Point{
					X: rect.Min.X + rect.Dx()/2,
					Y: rect.Min.Y + rect.Dy()/2,
				}
				gocv.Circle(&img, center, 4, blueColor, 4)

			}
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
