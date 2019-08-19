package average_code

import (
	"fmt"
	"math/rand"
	"time"
)

//二倍均值法
//剩余红包金额为M，剩余人生为N，那么有如下的公式：
//每次抢到的金额 = 随机区间 (0, M/N *2)
//amount 分
func AverageAmount(amount, num int64) {
	overNum := num
	var i int64
	var val int64
	for i = 0; i < num; i++ {
		if overNum == 1 {
			val = amount
		} else {
			rand.Seed(time.Now().UnixNano())
			val = rand.Int63n(amount/overNum*2-1) + 1

		}
		fmt.Println("第:", i, "个人，抢的钱 :", val, "分")
		amount -= val
		overNum -= 1
	}
}
