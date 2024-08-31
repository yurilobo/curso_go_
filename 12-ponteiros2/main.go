package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhaVa1 := 10
	minhaVa2 := 20
	soma(&minhaVa1, &minhaVa2)
	println(minhaVa1)
}
