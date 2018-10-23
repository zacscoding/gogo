package hangul

import (
	"fmt"
	"testing"
)

func TestHasConsonantSuffix(t *testing.T) {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))

	// Output
	//false
	//true
	//false

}