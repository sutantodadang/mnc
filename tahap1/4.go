package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func No4() {

	const (
		totalLeave   = 14
		waitForLeave = 180
		maxLeave     = 3
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Jumlah cuti bersama: ")
	scanner.Scan()
	leaveDays, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tanggal join karyawan: ")
	scanner.Scan()
	joinDate, err := time.Parse("2006-01-02", scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tanggal rencana cuti: ")
	scanner.Scan()
	leaveDate, err := time.Parse("2006-01-02", scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Durasi cuti (hari): ")
	scanner.Scan()
	leaveDuration, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	allowedLeave := joinDate.AddDate(0, 0, waitForLeave)

	if leaveDate.Before(allowedLeave) {
		fmt.Println("false")
		fmt.Println("Alasan: Karena belum 180 hari sejak tanggal join karyawan")
		return
	}

	privateLeave := totalLeave - leaveDays

	if leaveDate.Year() == joinDate.Year() {
		totalWorkingDays := 365 - int(allowedLeave.Sub(time.Date(joinDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)).Hours()/24)
		privateLeave = totalWorkingDays * privateLeave / 365

	}

	if leaveDuration > privateLeave {
		fmt.Println("false")
		fmt.Printf("Alasan: Karena hanya boleh mengambil %d hari cuti\n", privateLeave)
		return
	}

	fmt.Println("true")

}
