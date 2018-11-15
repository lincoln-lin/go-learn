package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	str := string(data)
	arr := strings.Split(str, " ")
	M, _ := strconv.Atoi(arr[0])
	N, _ := strconv.Atoi(arr[1])
	F := ""
	if  M < 0 {
		M = M* -1
		F = "-"
	}
	if N < 17 {
		values := []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G",
		}
		var r string
		for {
			l := M % N
			M = int(M / N)
			r = values[l] + r
			if M == 0 {
				break
			}
		}
		fmt.Println(F+r)
	}

}

