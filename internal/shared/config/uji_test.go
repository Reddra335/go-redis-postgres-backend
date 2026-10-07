package config

import (
	"fmt"
	"testing"
)

func TestUji(t *testing.T) {

	config := Load()
	fmt.Println(config.DbName)
}

func TestDB(t *testing.T) {
	InitDB()
}
