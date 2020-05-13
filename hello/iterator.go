package main

//Repeat characters
func Repeat(character string, num int) string {
	finalResult := ""
	for i := 0; i < num; i++ {
		finalResult += character
	}
	return finalResult
}
