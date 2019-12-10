package consistent_hash_code

import (
	"testing"
)

func TestNewConsistent(t *testing.T) {
	hash := NewConsistent()
	hash.Add("1000")
	hash.Add("2000")
	hash.Add("3000")
	t.Log(hash.GetNode())
	t.Log(hash.Get("1"))
	t.Log(hash.Get("2"))
	t.Log(hash.Get("3"))
	t.Log(hash.Get("4"))
	t.Log(hash.Get("5"))
	t.Log(hash.Get("6"))
	t.Log(hash.Get("7"))
	t.Log(hash.Get("8"))
	t.Log(hash.Get("9"))
	t.Log(hash.Get("10"))
	t.Log("remove", hash.Remove("3000"))
	t.Log(hash.GetNode())
	t.Log(hash.Get("1"))
	t.Log(hash.Get("2"))
	t.Log(hash.Get("3"))
	t.Log(hash.Get("4"))
	t.Log(hash.Get("5"))
	t.Log(hash.Get("6"))
	t.Log(hash.Get("7"))
	t.Log(hash.Get("8"))
	t.Log(hash.Get("9"))
	t.Log(hash.Get("10"))
}
