//
// intcal.go
// Copyright (C) 2020 damon <damon@asheville>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const MONTHLY_RATE float64 = 1.5

	display_instructions(MONTHLY_RATE)

	fmt.Println("\nEnter the number of days late: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	days_late, _ := strconv.ParseFloat(scanner.Text(), 64)

	if scanner.Err() != nil {
		fmt.Println("There was an error")
	}

	fmt.Println("\nEnter the invoice amount: ")

	amount_scanner := bufio.NewScanner(os.Stdin)
	amount_scanner.Scan()
	invoice_amount, _ := strconv.ParseFloat(amount_scanner.Text(), 64)

	yearly_rate := MONTHLY_RATE * 12.0
	yearly_multiplier := yearly_rate / 100.0
	year_portion := days_late / 365.0
	partial_rate := year_portion * yearly_multiplier
	interest_amount := invoice_amount * partial_rate

	fmt.Println("Interest Owed:", interest_amount)
}

func display_instructions(monthly_rate float64) {
	fmt.Println("Calculate interest on a late invoice.")
	fmt.Println("This assumes", monthly_rate, "monthly interest.")
}
