package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// printBytes()
	// printBytes2()
	// modifyBytes()
	// strCat()
	// strToNumeric()
	// strToNumeric2()
	// exampleArray()
	// exampleSlice()
	// sliceAppend()
	// sliceCap()
	// sliceCopy()
	addAndDelete()
}

func printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output :
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)

	// Output :
	//eab080eb8298eb8ba4
	//ea b0 80 eb 82 98 eb 8b a4
}

func modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))

	// Output :
	// 각나다
}

func strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

	s = fmt.Sprint(s, "ghi")
	fmt.Println(s)
	fmt.Println(*ps)

	s = fmt.Sprintf("%sjkl", s)
	fmt.Println(s)
	fmt.Println(*ps)

	s = strings.Join([]string{s, "mno"}, "")
	fmt.Println(s)
	fmt.Println(*ps)

	// Output :
	//abcdef
	//abcdef
	//abcdefghi
	//abcdefghi
	//abcdefghijkl
	//abcdefghijkl
	//abcdefghijklmno
	//abcdefghijklmno
}

func strToNumeric() {
	var i int
	var k int64
	var f float64
	var s string
	var err error
	i, err = strconv.Atoi("350") // i == 350
	fmt.Println("i : ", i)
	k, err = strconv.ParseInt("cc7fdd", 16, 32) // k == 13402077
	fmt.Println("k : ", k)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32) // k == 13402077
	fmt.Println("k : ", k)
	f, err = strconv.ParseFloat("3.14", 64) // f == 3.14
	fmt.Println("f : ", f)
	s = strconv.Itoa(340) // s == 340
	fmt.Println("s : ", s)
	s = strconv.FormatInt(13402077, 16) // s == cc7fdd
	fmt.Println("s : ", s)
	fmt.Println(err)
}

func strToNumeric2() {
	var num int
	fmt.Sscanf("57", "%d", &num) // num == 57
	fmt.Println("num : ", num)
	var s string
	s = fmt.Sprint(3.14) // s = "3.14"
	fmt.Println("s : ", s)
	s = fmt.Sprintf("%x", 13402077) // s == "cc7fdd"
	fmt.Println("s : ", s)
}

func exampleArray() {
	// fruits := [3]string{"apple", "banana", "tomato"}
	fruits := [...]string{"apple", "banana", "tomato"}

	// _ : index | fruit : value
	for _, fruit := range fruits {
		fmt.Printf("%s is delicious :)\n", fruit)
	}

	// Output
	//apple is delicious :)
	//banana is delicious :)
	//tomato is delicious :)
}

func exampleSlice() {
	// fruits == nill
	// var fruits []string

	// fruits == 빈 string 3개
	fruits := make([]string, 3)
	fmt.Printf("fruits[1] : %s", fruits[1])

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	fmt.Println(nums[:len(nums)-1])

	// Output
	// fruits[1] : [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
	// [1 2 3 4]
}

func sliceAppend() {
	fruits := []string{"apple", "banana", "tomato"}
	fmt.Println(fruits)

	fruits = append(fruits, "grape", "strawberry")
	fmt.Println(fruits)

	f1 := []string{"apple", "banana", "tomato"}
	f2 := []string{"grape", "strawberry"}

	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)

	fmt.Println(f3)
	fmt.Println(f4)

	// Output
	// [apple banana tomato]
	// [apple banana tomato grape strawberry]
	// [apple banana tomato grape strawberry]
	// [apple banana grape strawberry]
}

func sliceCap() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("len : ", len(nums))
	fmt.Println("cap : ", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len : ", len(sliced1))
	fmt.Println("cap : ", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len : ", len(sliced2))
	fmt.Println("cap : ", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len : ", len(sliced3))
	fmt.Println("cap : ", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)

	// Output:
	//[1 2 3 4 5]
	//len :  5
	//cap :  5
	//
	//[1 2 3]
	//len :  3
	//cap :  5
	//
	//[3 4 5]
	//len :  3
	//cap :  3
	//
	//[1 2 3 4]
	//len :  4
	//cap :  5
	//
	//[1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}

func sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))

	for i := range src {
		dest[i] = src[i]
	}

	fmt.Println(dest)
	// [30 20 50 10 40]

	dest2 := make([]int, len(src)-1)

	if n := copy(dest2, src); n != len(src) {
		fmt.Println("복사가 덜 됐습니다.")
	}
}

func addAndDelete() {
	a := []int{1, 2, 3, 4, 5, 6}

	a = addArray2(a, 3, 9)
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}
	b = deleteArray(b, 2, 2)

	fmt.Println(b)

	// output
	// [1 2 3 9 4 5 6]
	// [1 2 5]
}

func addArray(array []int, idx int, value int) []int {
	if idx < len(array) {
		array = append(array[:idx+1], array[idx:]...)
		array[idx] = value
	} else {
		array = append(array, value)
	}

	return array
}

func addArray2(array []int, idx int, value int) []int {
	array = append(array, value)

	copy(array[idx+1:], array[idx:])
	array[idx] = value

	return array
}

func deleteArray(array []int, idx int, count int) []int {
	return append(array[:idx], array[idx+count])
}
