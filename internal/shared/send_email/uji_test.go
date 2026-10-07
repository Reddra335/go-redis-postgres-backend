package sendemail

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestDisplay(t *testing.T) {

	u, _ := uuid.NewV7()
	res := u.ID()
	fmt.Println(res)
}
