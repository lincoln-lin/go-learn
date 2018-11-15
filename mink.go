package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"fmt"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	data ,_,_ := reader.ReadLine()
	str := string(data)
	arr := strings.Split(str," ")
	intArr := make([]int,len(arr))
	for i:=0; i< len(arr) ; i++ {
		intArr[i],_ =  strconv.Atoi(arr[i])
	}
	k := intArr[len(intArr)-1]
	intArr = intArr[:len(intArr)-1]
	Minarr := intArr[:k]
	intArr = intArr[k:]
	sort.Ints(Minarr)
	for i:=0; i< len(intArr) ; i++ {
		if intArr[i] < Minarr[k-1] {
			Minarr[k-1] = intArr[i]
			sort.Ints(Minarr)
		}
	}
	for k,v:=range Minarr{
		if k==0 {
			fmt.Printf("%d", v)
		}else {
			fmt.Printf(" %d", v)
		}
	}

}
