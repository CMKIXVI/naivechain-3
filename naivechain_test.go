package naivechain_test

import (
	"fmt"
	"testing"

	"github.com/blainsmith/naivechain"
)

func TestNew(t *testing.T) {
	bc := naivechain.New([]byte("genesis block"), nil)

	// bc.Print()

	bc.Write([]byte("foo"))

	// bc.Print()

	bc.Write([]byte("bar"))

	bc.Print()

	fmt.Println(bc.Tail().Previous().Previous())

	t.Fail()
}
