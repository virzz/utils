package password_test

import (
	"fmt"
	"testing"

	"github.com/virzz/utils/password"
)

func TestGenerate(t *testing.T) {
	fmt.Println(password.Generate(16, password.Alpha))
	fmt.Println(password.Generate(16, password.Word))
	fmt.Println(password.Generate(16, password.All))
}

func TestCheck(t *testing.T) {
	p := password.Generate(16, password.All)
	fmt.Println(p)
	s, c := password.Check(p)
	fmt.Println(s, " ", c)
	// fmt.Println(password.Generate(16, password.Word))
	// fmt.Println(password.Generate(16, password.All))
}

func TestEncrypt(t *testing.T) {
	fmt.Println(password.Encrypt(password.Generate(16, password.Word)))
}
