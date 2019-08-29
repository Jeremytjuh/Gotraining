package swapper

//Swapper swaps the values of string1 and string2 and returns the swapped values
func Swapper(s1 string, s2 string) (string, string) {
	s1, s2 = s2, s1
	return s1, s2
}
