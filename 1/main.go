package main

const a = "Hello World!"

var (
	b bool
	c int
	d string
	e float64
)

func main() {
	print(b)
	b = true
	print(b)
	print(a)
	print(c)

	print(d)
	print(e)

}
