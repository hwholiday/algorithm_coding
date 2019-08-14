package similarity_code

import (
	"fmt"
	"testing"
)

func TestCosine(t *testing.T) {
	fmt.Println(Cosine([]float64{1, 1, 1, 1, 1}, []float64{1, 0, 1, 1, 1, 1, 1, 1}))
}
