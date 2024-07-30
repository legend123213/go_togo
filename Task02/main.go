package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)
func WordFrequency(s string) map[string]int{
 str:=strings.ToLower(strings.TrimSpace(s))
 arrWord:=strings.Split(str," ")
 count := make(map[string]int)
 for _, ch := range arrWord {
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
func getInput(r *bufio.Reader) (string,error){
	input,err := r.ReadString('\n')
	return input,err
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	str,err := getInput(reader)
	if err != nil{
		panic("wronge input")
	}
	fmt.Println(WordFrequency(str))
	fmt.Println(Palindrom(str))
}
