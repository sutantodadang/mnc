package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func No1() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	var resultIndex []int

	mapSeen := make(map[string][]int)

	for i := 0; i < N; i++ {
		scanner.Scan()

		lowerWord := strings.ToLower(scanner.Text())

		mapSeen[lowerWord] = append(mapSeen[lowerWord], i+1)

	}

	for _, v := range mapSeen {
		if len(v) > 1 {
			resultIndex = append(resultIndex, v...)
			break
		}
	}

	if len(resultIndex) == 0 {
		println("false")
	}

	for _, v := range resultIndex {
		fmt.Printf("%d ", v)
	}

}
