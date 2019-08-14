package similarity_code

import (
	"github.com/huichen/sego"
	"testing"
)

func TestCosine(t *testing.T) {
	t.Log(Cosine([]float64{1, 1, 1, 1, 1}, []float64{1, 0, 1, 1, 1, 1, 1, 1}))
}

func TestSego(t *testing.T) {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("./dictionary.txt")
	segments := segmenter.Segment([]byte("梦是美好的吗?我不能确定，但也不是悲伤的。天使的彩色羽翼张开，羽毛飘落，每一片羽毛都有一个梦，你的梦是什么颜色的羽毛?"))
	a := sego.SegmentsToSlice(segments, false)
	b := sego.SegmentsToString(segments, true)
	for _, v := range a {
		t.Log("a >>>>>>", v)
	}
	for _, v := range b {
		t.Log("b >>>>>>", v)
	}
}
