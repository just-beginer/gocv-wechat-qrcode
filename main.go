package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func main() {
	qrcode()
}

// 识别二维码
func qrcode() {
	mat := gocv.IMRead("./images/qrcode.png", gocv.IMReadColor)
	mats := make([]gocv.Mat, 0)
	defer mat.Close()

	type args struct {
		img   gocv.Mat
		point *[]gocv.Mat
	}

	tests := []struct {
		name string
		args args
	}{
		{"TestDetectAndDecode", args{img: mat, point: &mats}},
	}

	path := "./opencv_3rdparty"
	for _, tt := range tests {
		wq := contrib.NewWeChatQRCode(path+"/detect.prototxt", path+"/detect.caffemodel",
			path+"/sr.prototxt", path+"/sr.caffemodel")

		// 检测和识别二维码
		got := wq.DetectAndDecode(tt.args.img, tt.args.point)
		for _, qrcode := range got {
			// 打印二维码内容
			fmt.Println(qrcode)
		}

		// 确定二维码位置  每个二维码返回4个点确定位置
		for _, point := range *(tt.args.point) {
			for _, index := range [4][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}} {
				// 划线  
				gocv.Line(&mat, image.Point{int(point.GetFloatAt(index[0], 0)), int(point.GetFloatAt(index[0], 1))},
					image.Point{int(point.GetFloatAt(index[1], 0)), int(point.GetFloatAt(index[1], 1))}, color.RGBA{255, 0, 0, 255}, 5)
			}
		}
	}

	// 保存图片
	gocv.IMWrite("./output.png", mat)
}
