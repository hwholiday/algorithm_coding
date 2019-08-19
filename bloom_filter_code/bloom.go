package bloom_filter_code

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

//计算最佳的配置
//https://hur.st/bloomfilter/
//n代表元素的个数
//p代表假阳率
//m代表位图长度
//k代表hash函数的个数
type BloomFilter struct {
	m uint64 //数组集合大小
	k uint32 //hash函数个数
	b *bitset.BitSet
}

func New(m uint64, k uint32) *BloomFilter {
	return &BloomFilter{
		m, k, bitset.New(uint(m)),
	}
}

func (f *BloomFilter) Add(data []byte) {
	for i := uint32(0); i < f.k; i++ {
		f.b.Set(f.locate(data, i))
	}
}
func (f *BloomFilter) Exist(data []byte) bool {
	for i := uint32(0); i < f.k; i++ {
		if !f.b.Test(f.locate(data, i)) { //一个不存在就绝对不存在
			return false
		}
	}
	return true
}

func (f *BloomFilter) locate(data []byte, seed uint32) uint {
	return getHash(data, seed) % uint(f.m)
}

func getHash(data []byte, seed uint32) uint {
	m := murmur3.New64WithSeed(seed)
	_, _ = m.Write(data)
	return uint(m.Sum64())
}
