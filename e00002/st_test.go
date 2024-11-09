package e00002

import "testing"

func Test_Private(t *testing.T) {
	f := (st{}).private
	f()
}
