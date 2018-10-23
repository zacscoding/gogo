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
	strToNumeric2()
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
	fmt.Sscanf("57", "%d", &num)	// num == 57
	fmt.Println("num : ", num)
	var s string
	s = fmt.Sprint(3.14)	// s = "3.14"
	fmt.Println("s : ", s)
	s = fmt.Sprintf("%x", 13402077)	// s == "cc7fdd"
	fmt.Println("s : ", s)
}
