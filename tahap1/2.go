package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func No2() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Total belanja seorang customer: Rp ")
	scanner.Scan()

	totalGroceries, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pembeli membayar: Rp ")
	scanner.Scan()
	payment, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		panic(err)
	}

	if payment < totalGroceries {
		fmt.Println("false, kurang bayar")
		return
	}

	change := payment - totalGroceries

	roundedChange := math.Floor(change/100) * 100

	moneyDenomination := []int{100000, 50000,
		20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	fmt.Printf("Kembalian yang harus diberikan kasir: %.0f, dibulatkan menjadi %.0f\n", change, roundedChange)
	fmt.Println("Pecahan uang:")

	remainingChange := int(roundedChange)

	for _, v := range moneyDenomination {
		count := remainingChange / v

		if count > 0 {
			unit := "lembar"

			if v <= 500 {
				unit = "koin"
			}

			fmt.Printf("%d %s %d\n", count, unit, v)
			remainingChange -= v * count
		}
	}
}
