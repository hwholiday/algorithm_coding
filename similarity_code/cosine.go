package similarity_code

import (
	"github.com/huichen/sego"
	"math"
)

//比对两片文章的相似度

func ComPare(a string, b string) float64 {
	var (
		aArray    = Divide(a)
		bArray    = Divide(b)
		aArrayLen = len(aArray)
		bArrayLen = len(bArray)
		aMap      = make(map[string]int, aArrayLen)
		bMap      = make(map[string]int, bArrayLen)
		allMap    = make(map[string]string, bArrayLen+aArrayLen)
		AllArray  = append(aArray, bArray...)
		aA        []float64
		bA        []float64
	)
	//计算出全部词的并集
	for _, v := range AllArray {
		allMap[v] = ""
	}
	//开始统计每个次出现的次数
	for _, vA := range aArray {
		if _, ok := aMap[vA]; ok {
			aMap[vA] += 1
		} else {
			aMap[vA] = 1
		}
	}
	for _, vB := range bArray {
		if _, ok := bMap[vB]; ok {
			bMap[vB] += 1
		} else {
			bMap[vB] = 1
		}
	}
	//计算每个key出现的次数
	for key, _ := range allMap {
		if _, ok := aMap[key]; ok {
			aA = append(aA, float64(aMap[key]))
		} else {
			aA = append(aA, 0)
		}
		if _, ok := bMap[key]; ok {
			bA = append(bA, float64(bMap[key]))
		} else {
			bA = append(bA, 0)
		}
	}
	//计算相似度
	return Cosine(aA, bA)
}

func Divide(a string) []string {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("./dictionary.txt")
	segments := segmenter.Segment([]byte(a))
	return sego.SegmentsToSlice(segments, true)
}

//向量空间余弦相似度(Cosine Similarity)
func Cosine(a []float64, b []float64) float64 {
	var (
		aLen  = len(a)
		bLen  = len(b)
		sum   = 0.0
		sa    = 0.0
		sb    = 0.0
		count = 0
	)
	if aLen > bLen {
		count = aLen
	} else {
		count = bLen
	}
	for i := 0; i < count; i++ {
		if i >= aLen {
			sb += math.Pow(b[i], 2)
			continue
		}
		if i >= bLen {
			sa += math.Pow(a[i], 2)
			continue
		}
		sum += a[i] * b[i]
		sa += math.Pow(a[i], 2)
		sb += math.Pow(b[i], 2)
	}
	return sum / (math.Sqrt(sa) * math.Sqrt(sb))
}
