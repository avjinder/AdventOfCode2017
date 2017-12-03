package DayOne

import (
	"strings"
	"strconv"
	"runtime"
	"path"
	"io/ioutil"
	"fmt"
)

//INVERSE CAPTCHA
func RunDayOne(){
	_, fn, _, ok := runtime.Caller(0)
	if !ok{
		return
	}

	fp := path.Join(path.Dir(fn), "dayOneInput.txt")
	fa, _ := ioutil.ReadFile(fp)

	result := findCheckSum(string(fa), 1)
	fmt.Printf("Inverse Checksum for 1st puzzle is %d\n", result)
}

func RunDayTwo(){
	_, fn, _, ok := runtime.Caller(0)
	if !ok{
		return
	}

	fp := path.Join(path.Dir(fn), "dayOneInput.txt")
	ba, _ := ioutil.ReadFile(fp)

	inputString := string(ba)
	//inputString := "12131415"
	result := findCheckSum(inputString, len(inputString) / 2)
	fmt.Printf("Inverse Checksum for 2nd puzzle is %d\n", result)
}

func findCheckSum(input string, step int) int{
	splitString := strings.Split(input, "")
	var sum int
	tempIndex := 0
	inputLen := len(input)

	for i, v := range splitString{
		tempIndex = i
		tempIndex = (tempIndex + step) % inputLen
		fmt.Println(tempIndex)
		if v == splitString[tempIndex]{
			if x, e := strconv.ParseInt(v, 10, 32); e == nil{
				sum += int(x)
			}
		}
	}

	return sum
}

