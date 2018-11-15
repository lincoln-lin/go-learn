package main

import "fmt"

func main()  {
	fmt.Println(AddHex("48","18Z"))
}

/**
36进制加法
 */
func AddHex(a,b string) (r string) {
	hex := map[string]int{
		"0":0,
		"1":1,
		"2":2,
		"3":3,
		"4":4,
		"5":5,
		"6":6,
		"7":7,
		"8":8,
		"9":9,
		"A":10,
		"B":11,
		"C":12,
		"D":13,
		"E":14,
		"F":15,
		"G":16,
		"H":17,
		"I":18,
		"J":19,
		"K":20,
		"L":21,
		"M":22,
		"N":23,
		"O":24,
		"P":25,
		"Q":26,
		"R":27,
		"S":28,
		"T":29,
		"U":30,
		"V":31,
		"W":32,
		"X":33,
		"Y":34,
		"Z":35,
	}
	hex2 := map[int]string{}
	for k,v :=  range hex {
		hex2[v] = k
	}
	lena := len(a)
	lenb := len(b)
	var maxlen int
	if lena > lenb {
		maxlen = lena
	}else{
		maxlen = lenb
	}
	an,bn,p := 0 ,0 ,0

	for i :=0;i < maxlen; i++ {
		if lena-i >= 1  {
			an = hex[a[lena-1-i:lena-i]]
		}else{
			an = 0
		}
		if lenb-i >= 1  {
			bn = hex[b[lenb-1-i:lenb-i]]
		}else{
			bn = 0
		}

		n := an + bn + p
		if n >= 36 {
			p = 1
			n = n - 36
		} else {
			p = 0
		}
		r = hex2[n] + r
	}
	return r

}