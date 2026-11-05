package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Mahasiswa struct {
	Name string `json:"name"`
	Umur int    `json:"umur"`
}

func TestUji(t *testing.T) {

	mhs := []Mahasiswa{
		{Name: "Reddra", Umur: 20},
		{Name: "Rendi", Umur: 20},
	}

	b, _ := json.Marshal(mhs)

	fmt.Println(string(b))

	fmt.Println("hai")

}
