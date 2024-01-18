package arraysslices

import (
	"reflect"
	"testing"
)

// go test -cover
func TestSum(t *testing.T) {
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		resultado := Sum(numbers)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expect := []int{3, 9}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result %v expect %v", result, expect)
	}
}
