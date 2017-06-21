package avc

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

type AvcNalUnit struct {
	header *AvcHeader
}

func NewAvcNalUnit() (nal *AvcNalUnit) {
	nal = new(AvcNalUnit)
	nal.header = NewAvcHeader()
	return nal
}

//查找0x000001
func (nal *AvcNalUnit) FindStartCode2(buf []byte) bool {
	if buf[0] == 0 && buf[1] == 0 && buf[2] == 1 {
		return true
	} else {
		return false
	}
}

//查找0x00000001
func (nal *AvcNalUnit) FindStartCode3(buf []byte) bool {
	if buf[0] == 0 && buf[1] == 0 && buf[2] == 0 && buf[3] == 1 {
		return true
	} else {
		return false
	}
}

//开始分析
func (nal *AvcNalUnit) StartAnalyze(filename string) (err error) {
	//读取文件
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(hex.Dump(buf))

	fileLen := len(buf)
	index := 0

	findStartCode := false

	for {
		if nal.FindStartCode2(buf[index : index+3]) {
			findStartCode = true
			index += 3
		} else {
			if nal.FindStartCode3(buf[index : index+4]) {
				findStartCode = true
				index += 4
			}
		}

		if findStartCode {
			err = nal.header.demux(buf[index:])
			if err != nil {
				fmt.Println("demux header error , error info is ", err.Error())
				break
			}
			index++
		}
		index++
		if index >= fileLen {
			break
		}
	}
	return nil
}
