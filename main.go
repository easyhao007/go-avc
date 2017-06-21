package main

import (
	"fmt"
	"go-avc/avc"
)

func main() {
	fmt.Println("h264结构分析学习:")

	nal := avc.NewAvcNalUnit()
	nal.StartAnalyze("D:\\video\\aFanDa.h264")
}
