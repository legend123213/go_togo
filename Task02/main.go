package main

import (
	"fmt"
	"strings"
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
str := strings.ToUpper(s)
i := 0
j := len(s)-1
print(str[i])
for i<=j {
	if str[i]<65 || str[i]>90{
		i++
		continue
	}
	if str[j]<65 || str[j]>90{
		j--
		continue
	}
	fmt.Println(i,j)
	if str[i]!= str[j]{
		return false
	}
	i++
	j--
}
return true
}

func main(){
	fmt.Println(WordFrequency("abel wendmu"))
	fmt.Println(Palindrom("abe ba"))

}
