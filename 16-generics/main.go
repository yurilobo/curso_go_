package main

type MyNumber int

//solucao usando constranct

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false

}

//solucao 1
// func Soma[T int | float64](m map[string]T) T {
// 	var soma T
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

func main() {
	m1 := map[string]int{"Yuri": 1000, "Ruan": 3500, "Chaves": 5443}
	m2 := map[string]float64{"Yuri": 10.36, "Ruan": 35.59, "Chaves": 54.43}

	m3 := map[string]MyNumber{"Yuri": 1000, "Ruan": 3500, "Chaves": 5443}

	println(Soma(m1))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(30, 30.00))
}
