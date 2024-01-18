package arraysslices

import "testing"

// go test -cover
func TestSum(t *testing.T) {
	// t.Run("coleção de 5 numeros", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		resultado := Sum(numbers)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numbers)
		}
	})
}
