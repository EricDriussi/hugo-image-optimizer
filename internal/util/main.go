package util

import "strings"

func StringIsInArray(arr []string, str string) bool {
	for _, a := range arr {
		if strings.Contains(a, str) {
			return true
		}
	}
	return false
}

func ByteArrayContainsByteArray(smol *[]byte, chungus *[]byte) bool {
	str_len := len(*smol)
	arr_len := len(*chungus)

	for i := 0; i <= arr_len-str_len; i++ {
		arr_chunk := (*chungus)[i : i+str_len]
		if ByteArraysAreEqual(*smol, arr_chunk) {
			return true
		}
	}
	return false
}

func ByteArrayContainsString(str string, arr *[]byte) bool {
	str_as_byte_array := []byte(str)
	str_len := len(str_as_byte_array)
	arr_len := len(*arr)

	for i := 0; i <= arr_len-str_len; i++ {
		arr_chunk := (*arr)[i : i+str_len]
		if ByteArraysAreEqual(str_as_byte_array, arr_chunk) {
			return true
		}
	}
	return false
}

func ByteArraysAreEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
