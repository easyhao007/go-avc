package avc

import (
	"errors"
	"fmt"
	"go-avc/bitbuffer"
)

func Introduce() {
	fmt.Println("解析AVC的头")
}

const (
	NalUintTypeUnuse = iota
	NalUintTypeNoSliceNoIdr
	NalUnitTypeSliceA
	NalUnitTypeSliceB
	NalUnitTypeSliceC
	NalUnitTypeIDR
	NalUnitTypeSEI
	NalUintSPS
	NalUnitPPS
	NalUnitAUD
	NalUnitEndSeq
	NalUnitEndStream
	NalUnitFillerData
	NalUnitSPSExt
	NalUnitReserv14
	NalUnitReserv15
	NalUnitReserv16
	NalUnitReserv17
	NalUnitReserv18
	NalUnitSliceLayerWithoutPartitioning
)

type AvcHeader struct {
	buf              bitbuffer.BitBuffer
	ForBiddenZeroBit uint8 //1b
	NalRefIdc        uint8 //2b
	NalUnitType      uint8 //5b
}

func (head *AvcHeader) demux(buf []uint8) (err error) {
	//设置缓存
	//解析数据
	head.ForBiddenZeroBit = buf[0] & 0x80		// 0x80 = 10000000
	if head.ForBiddenZeroBit != 0 {
		err = errors.New("head.ForBiddenZeroBit 不等于0 , buff")
	}

	head.NalRefIdc = buf[0] & 0x60
	head.NalUnitType = buf[0] & 0x1F

	switch head.NalUnitType {
	case NalUintTypeUnuse:
		fmt.Println("未指定")
		break
	case NalUintTypeNoSliceNoIdr:
		fmt.Println("一个非IDR图像的编码条带")
		break
	case NalUnitTypeSliceA:
		fmt.Println("编码条带数据分割块A")
		break
	case NalUnitTypeSliceB:
		fmt.Println("编码条带数据分割块B")
		break
	case NalUnitTypeSliceC:
		fmt.Println("编码条带数据分割块C")
		break
	case NalUnitTypeIDR:
		fmt.Println("IDR图像的编码条带")
		break
	case NalUnitTypeSEI:
		fmt.Println("辅助增强信息 (SEI)")
		break
	case NalUintSPS:
		fmt.Println("序列参数集")
		break
	case NalUnitPPS:
		fmt.Println("图像参数集")
		break
	case NalUnitAUD:
		fmt.Println("访问单元分隔符")
		break
	case NalUnitEndSeq:
		fmt.Println("序列结尾")
		break
	case NalUnitEndStream:
		fmt.Println("流结尾")
		break
	case NalUnitFillerData:
		fmt.Println("填充数据")
		break
	case NalUnitSPSExt:
		fmt.Println("序列参数集扩展")
		break
	default:
		fmt.Println("无效的类型")
		break
	}
	return nil
}

func NewAvcHeader() (head *AvcHeader) {
	head = new(AvcHeader)
	return head
}
