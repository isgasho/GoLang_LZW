package main

import (
	"fmt"

	"testing"
)

func TestByte_in_dbl_slice(t *testing.T) {
	dict := make([][]byte, 2) // создадим dict [[1 2] [3]]
	dict[0] = make([]byte, 2)
	dict[1] = make([]byte, 1)
	dict[0][0] = 1
	dict[0][1] = 2
	dict[1][0] = 3
	fmt.Println("dict  ", dict) //это потом уберем
	char := make([]byte, 2)     //char [3]
	char[0] = 1
	char[1] = 2
	fmt.Println("char  ", char) //и это уберем
	//проверяем функцию
	bl, id := byte_in_dbl_slice(dict, char) //в bl и id должны записаться true и 0 соответственно
	fmt.Println("bl ", bl, " id ", id)
	if (bl != true) || (id != 0) { //проверяем так ли это
		t.Error("ОШИБКА, БЛЯТЬ!")
	}
}

// тестим подходят ли словари, берем за пример input.txt
// чтобы что то сломать меняй значения что я добавляю в словарь и будет ошибка
func TestFill_in_dbl_dic(t *testing.T) {
	var dict [][]byte
	var res [][]byte
	var n int
	n = 7
	var path = "input.txt"

	res = append(res, []byte{97})
	res = append(res, []byte{98})
	res = append(res, []byte{99})
	res = append(res, []byte{100}) // ломать здесь !
	res = append(res, []byte{101})
	res = append(res, []byte{13})
	res = append(res, []byte{10})
	dict = fill_in_dbl_dic(dict, path)
	// fmt.Println(" res= ", res)
	// fmt.Println(" dict= ", dict)

	for i := 0; i < n; i++ {
		for j := 0; j < 1; j++ {
			// fmt.Println(" res el= ", res[i][j])
			// fmt.Println(" dict els= ", dict[i][j])
			if dict[i][j] != res[i][j] {
				t.Errorf("%d != %d , ошибка словаря йопта", res, dict)
			}
		}
	}
}

func TestCompress(t *testing.T) {
	var n int
	n = 13
	var path = "input.txt"

	var dict [][]byte
	dict = fill_in_dbl_dic(dict, path)
	res := []byte{0, 1, 0, 2, 7, 0, 3, 11, 10, 8, 4, 5, 6} // ломать тут
	message := compress(dict, path)

	for i := 0; i < n; i++ {
		if message[i] != res[i] {
			t.Errorf("%d != %d", message, res)
		}
	}
}
