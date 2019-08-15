package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
)

type UserScore struct {
	Uid    int
	Mid    int
	Rating int
}

type MovieScore struct {
	Mid    int
	Rating int
}

type UidScore struct {
	Uid    int
	Rating int
}

type FidScore struct {
	Fid    int
	Rating float64
}

type FidScoreArray []FidScore

func (f FidScoreArray) Len() int {
	return len(f)
}
func (f FidScoreArray) Less(i, j int) bool {
	return f[i].Rating > f[j].Rating
}
func (f FidScoreArray) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type UidAndFirScore struct {
	UidRating int
	FidRating int
}

func ReadLine(path string) (data []string) {
	info, err := ioutil.ReadFile(path)
	CheckError(err)
	reader := bufio.NewReader(bytes.NewReader(info))
	for {
		if v, err := reader.ReadString('\n'); err == nil {
			data = append(data, v)
		} else {
			if err == io.EOF {
				break
			}
			CheckError(err)
		}
	}
	return
}
func CheckInArray(data []int, k int) bool {
	for _, v := range data {
		if v == k {
			return true
		}
	}
	return false
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
