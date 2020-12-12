package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type student struct{
	name string
}

func parseLine(line string)(stud student , err error){
	nameRE:=regexp.MustCompile("([a-z]|[A-Z]|[0-9]|_)+")
	scoreRE:=regexp.MustCompile("[1-3][0-9]")

	stud.name=nameRE.FindString(line)
	nameEndingIndex :=nameRE.FindStringIndex(line)[1]
	score :=scoreRE.FindAllString(line[nameEndingIndex:] , -1)

	for _, x:= range score{
		score, conv := strconv.Atoi(x)

		if score < 18 || score > 32 || score == 32 {
			err=errors.New("the score "+x+" is invalid")
			return
		} else if conv !=nil{
			fmt.Println(conv)
		}else {
			fmt.Println("score "+x+" is valid")
		}
	}
	return
}

func main() {
	_, err:= parseLine("  paolo_mario30_boldi 18 18 29")

	if err !=nil{
		fmt.Fprintln(os.Stderr,err)
	}
}

