package standard_library

import (
	"fmt"
	"testing"
)

func Test_Copy(t *testing.T) {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{7, 8, 9, 10}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	//a6: [1 2 3 4 5 6]
	//a4: [7 8 9 10]
	//a6: [7 8 9 10 5 6]
	//a4: [7 8 9 10]
}

func Test_Copy_v2(t *testing.T) {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{7, 8, 9, 10}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a4, a6)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	//a6: [1 2 3 4 5 6]
	//a4: [7 8 9 10]
	//a6: [1 2 3 4 5 6]
	//a4: [1 2 3 4]
}
