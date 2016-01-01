package stringutil

//This is an inner function and isn't visible outside of the package
func reverseTwo(s string) string {
	//reverses letters in a string
	length := len(s)
	outString := []rune(s)
	for i := 0; i < (length / 2); i++ {
		outString[i], outString[length-1-i] = outString[length-1-i], outString[i]
	}
	return string(outString)
}
