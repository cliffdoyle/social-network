package main

import "fmt"

func main() {
	str := "hello there"
	res := ""
	word := ""

	//reverse the string
	for i := len(str) - 1; i >= 0; i-- {
		word += string(str[i])
		if str[i] == ' ' {
			res += word
			res+=" "
		}
		if str[i] != ' '{
			res += word
			
		}

	}

	fmt.Println(res)

}
