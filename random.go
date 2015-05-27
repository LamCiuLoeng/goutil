package goutil

import "crypto/rand"

func RandFromList(size int, list string) string {
	var bytes = make([]byte, size)
	rand.Read(bytes)
	l := byte(len(list))
	for k, v := range bytes {
		bytes[k] = list[v%l]
	}
	return string(bytes)

}

func RandNumber(size int) string {
	return RandFromList(size, "0123456789")

}

func RandAlpha(size int) string {
	return RandFromList(size, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func RandAlphaNum(size int) string {
	return RandFromList(size, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

}
