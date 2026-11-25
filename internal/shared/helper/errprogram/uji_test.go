package errprogram

import (
	"fmt"
	"testing"
)

func TestDisplay(t *testing.T) {

	n := 10
	if n == 10 {
		fmt.Println("mantap")
		return
	} else if n == 20 {
		fmt.Println("ihir")
	}

}
