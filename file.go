package e00003

func MyFN() {
	s := R()
	hello(1, R(), s)
}

func R() string {
	return "string"
}
