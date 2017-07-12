package avc

import (
	//"encoding/hex"
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

	head , tail , startLen := 0 , 0 , 0
	index := 1
	for i:= 0 ; i < len(buf) ; i++{
		if buf[i] == 0 && buf[i+1] == 0 && buf[i+2] == 1{
			//find 0x000001
			if head == 0{
				head = i + 3
				i = i + 3
			}else {
				tail = i
				i = i + 3
				startLen = 3
			}
		}else if buf[i] == 0 && buf[i+1] == 0 && buf[i+2] == 0 && buf[i+3] == 1{
			//find 0x00000001
			if head == 0{
				head = i + 4
				i = i + 4
			}else {
				tail = i
				i = i + 4
				startLen = 4
			}
		}
		if head != 0 && tail != 0{
			fmt.Printf("index:%d" , index)
			//fmt.Println(hex.Dump(buf[head:tail]))
			err := nal.header.demux(buf[head:tail])
			if err != nil{
				fmt.Println("demux the nal faild")
			}
			head = tail + startLen
			tail = 0
			index++
		}
	}


	return nil
}
