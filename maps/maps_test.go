package maps

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCount1(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)

	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount, ) {
		t.Error("codeCount mismatch : ", codeCount)
	}

	fmt.Println(codeCount)
}

func TestCount2(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		fmt.Println("push keys... ", key)
		keys = append(keys, int(key))
	}
	sort.Sort(keys)

	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
}

func TestHasDupeRune(t *testing.T) {
	if !hasDupeRune2("가나다가") {
		t.Fail()
	}
}
