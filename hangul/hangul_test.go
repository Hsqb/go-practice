package hangul

import (
	"fmt"
	"strconv"
)

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("GO 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	// Output:
	// ea b0 80 eb 82 98 eb 8b a4
	//
}

func Example_modifyByte() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_str2Num() {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")
	fmt.Println(i)
	k, err = strconv.ParseInt("cc7ffd", 16, 32)
	fmt.Println(k)
	k, err = strconv.ParseInt("0xcc7ffd", 0, 32)
	fmt.Println(k)
	f, err = strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
	s = strconv.Itoa(340)
	fmt.Println(s)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)
	fmt.Println(err)
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)
	// Output:
	// 350
	// 13402109
	// 13402109
	// 3.14
	// 340
	// cc7fdd
	// <nil>
	// 3.14
	// cc7fdd
}

func Example_array() {
	var fruits = [4]string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		if !HasConsonantSuffix(fruit) {
			fmt.Printf("%s는 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s은 맛있다.\n", fruit)
		}

	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 수박은 맛있다.
}

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
}
