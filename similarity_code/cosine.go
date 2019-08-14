package similarity_code

import "math"

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
