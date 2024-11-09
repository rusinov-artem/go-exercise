package e00002

import "fmt"

type st struct {
}

func (t st) private() {
	fmt.Println("This is private")
}
