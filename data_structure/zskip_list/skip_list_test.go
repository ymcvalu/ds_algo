package zskip_list

import (
	"k8s.io/apimachinery/pkg/util/rand"
	"testing"
)

var table = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 's', 'y', 'z'}

func randomKey() []byte {
	var c = 0
	for c < 3 {
		c = rand.Intn(10)
	}

	byts := make([]byte, c)
	for i := 0; i < c; i++ {
		byts[i] = table[rand.Intn(26)]
	}
	return byts
}

func randomScore() float64 {
	return float64(rand.Intn(100))
}

func randomNew(size int) *SkipList {
	sl := New()
	for i := 0; i < size; i++ {
		sl.Insert(randomScore(), randomKey())
	}
	return sl
}

func TestInsert(t *testing.T) {
	sl := randomNew(10)
	sl.print()
	sl.print_rev()
}

func TestRandomLevel(t *testing.T) {
	levelCnt := make([]int, 64)
	for i := 0; i < 20; i++ {
		levelCnt[randomLevel()-1]++
	}
	t.Log(levelCnt)
}

func TestGetByRank(t *testing.T) {
	sl := randomNew(10)

	sl.print()

	t.Log(sl.GetByRank(0))
	t.Log(sl.GetByRank(-1))

	t.Log(sl.GetByRank(sl.Length() / 2))
	t.Log(sl.GetByRank(-2))
}

func TestDelete(t *testing.T) {
	sl := randomNew(10)
	r := rand.Intn(sl.length)
	score, key, _ := sl.GetByRank(r)

	sl.print()
	t.Logf("remove node at rank[%d], score: %f, key: %s", r, score, string(key))
	if !sl.Delete(score, key) {
		t.Error("can't find the node to delete")
	}

	sl.print()
	sl.print_rev()
}
