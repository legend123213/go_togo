package main

import (
	"fmt"
)
func WordFrequency(s string) map[string]int{
 str:=s
 count := make(map[string]int)
 for _, ch := range str {
	count[string(ch)]+=1
 }
 return count
}

func Palindrom(s string) bool{
str := s
i := 0
j := len(s)-1
for i<j {
	if str[i]!= str[j]{
		return false
	}
}
return true
}

func main(){
	fmt.Println(WordFrequency("abel wendmu"))
	fmt.Println(Palindrom("abel wendmu"))

}
