package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/legend123213/go_togo/Task01/utils"
)

func getInput(r *bufio.Reader, s string) (string, error) {
	if len(s) != 0 {
		fmt.Printf(s+" ")
	}
	input, err := r.ReadString('\n')
	return input, err
}
func getSubjects() map[string]float64 {
	input := bufio.NewReader(os.Stdin)
	numberSubjects,err := getInput(input,"How many Subjects did you take?")
	numberSubject,err_ := strconv.Atoi(numberSubjects[:len(numberSubjects)-1])
	if err!=nil || err_ != nil{
		panic("Please Enter Number try again")
	}
	subjects := make(map[string]float64)
	i:= 0
	for i<numberSubject{
		nameSubject,_:=getInput(input,"Subject Name:")
		sco,_ := getInput(input,"Score of the Subject:")
		score,err_score := strconv.ParseFloat(strings.TrimSpace(sco),64)
		if err_score!=nil{
			fmt.Println("⚠️Please Enter Number")
			continue
		}
		if score<0 || score>100{
			fmt.Println("⚠️Please Enter with Range 0-100")
			continue
		}
		subjects[strings.TrimSpace(nameSubject)]=score
		i+=1
	}
	return subjects
}
func main() {
	input := bufio.NewReader(os.Stdin)
	name, err := getInput(input, "Please Enter your name:")
	if err != nil {
		panic("⚠️wrong input")
	}
	scores := getSubjects()
	average := utils.Average(scores)

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------")

	subjectsDetails := "\n"
	for subject, score := range scores {
		s := strconv.FormatFloat(score, 'f', 2, 64)
		subjectsDetails += fmt.Sprintf("%-27v ...$%v \n", subject+":",s)
	}

	result := fmt.Sprintf("NAME: %s\nScores \n%s\n", name, subjectsDetails)
	fmt.Println(result)
	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("Average Score: %.2f\n", average)
	fmt.Println("----------------------------------------------------------------")
	
}
