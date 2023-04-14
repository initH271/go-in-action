package test_interface

type animal interface {
	PrintInfo()
}

type Cat int

func invoke(a animal) {
	a.PrintInfo()
}

func (c *Cat) PrintInfo() {
	println("I'm a cat.")
}
