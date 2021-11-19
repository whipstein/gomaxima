package maxima

import "fmt"

func main() {
	m := NewMaxima()
	defer m.Close()

	fmt.Printf(m.Set("a", "5"))
	fmt.Printf("\n")
	fmt.Printf(m.Set("b", "3"))
	fmt.Printf("\n")
	fmt.Printf(m.Eval("a+b"))
	fmt.Printf("\n")
	fmt.Printf(m.Set("c", "d+e"))
	fmt.Printf("\n")
	fmt.Printf(m.Eval("c"))
	fmt.Printf("\n")
}
