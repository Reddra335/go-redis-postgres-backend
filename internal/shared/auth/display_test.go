package auth

import (
	"fmt"
	"testing"
)

func TestDisplay(t *testing.T) {

	tampil, _ := GenerateJwt(234, "reddra", "ranger@gmail.com")

	fmt.Println("isi: ", tampil)

}
