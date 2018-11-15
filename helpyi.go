package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	data ,_,_ := reader.ReadLine()
	str := string(data)
	n,_ := strconv.Atoi(str)
	data ,_,_ = reader.ReadLine()
	str = string(data)
	x := strings.Split(str," ")
	data ,_,_ = reader.ReadLine()
	str = string(data)
	y := strings.Split(str," ")
	xint,_:= strconv.Atoi(x[0])
	yint,_:= strconv.Atoi(y[0])
	min := xint + yint
	for i:=1 ; i < n ; i++ {
		xint,_= strconv.Atoi(x[i])
		yint,_= strconv.Atoi(y[i])
		newMin := xint + yint
		if newMin < min {
			min = newMin
		}
	}
	fmt.Println(min - 2)

}