package linkedlist

import (
	"testing"
)

func TestGenerateSingle(t *testing.T) {

	single := GenSingleList(4)

	single.Append(Node{Data: 10})

	single.String()

	n, _ := single.Remove(1)
	t.Logf("%+v\n",n)

	single.String()
}
