package ch1

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	var s, sep string
	timeCounter := time.Now()
	//s = os.Args[0]

	for i := 1; i < len(os.Args); i++ {
		if s == "" {
			sep = ""
		} else {
			sep = " "
		}
		s += sep + os.Args[i]
		//fmt.Println("Argument " + strconv.Itoa(i) + " : " + "\"" + os.Args[i] + "\"")
	}
	fmt.Println(s)

	var timeBreak1 = time.Since(timeCounter).Seconds()
	fmt.Printf("%.9fs elapsed\n", timeBreak1)

	timeBreak1 = time.Since(timeCounter).Seconds() - timeBreak1

	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)

	var timeBreak2 = time.Since(timeCounter).Seconds() - timeBreak1
	fmt.Printf("%.9fs elapsed\n", timeBreak2)
/*
	var res [16]string

	for i := 0; i <= 15; i++ {
		res[i] = strconv.Itoa(i)
		s = strings.Join(res[0:], " ")
		//sep = " "
	}
	fmt.Println(s)
*/
}


