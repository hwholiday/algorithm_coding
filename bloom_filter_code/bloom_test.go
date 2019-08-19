package bloom_filter_code

import (
	"github.com/spaolacci/murmur3"
	"testing"
)

func TestMurmurNew128(t *testing.T) {
	murmur := murmur3.New128()
	_, _ = murmur.Write([]byte("1"))
	v1, v2 := murmur.Sum128()
	t.Log("v1", v1, "v2", v2)
	t.Log("v1", v1%100, "v2", v2%100)
}

func TestNew(t *testing.T) {
	m := New(3355, 23)
	m.Add([]byte("a"))
	m.Add([]byte("b"))
	m.Add([]byte("c"))
	t.Log(m.Exist([]byte("a")))
	t.Log(m.Exist([]byte("c")))
	t.Log(m.Exist([]byte("d")))
}
