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

	days_late := get_numeric_input_at_prompt("Enter the number of days late:")
	invoice_amount := get_numeric_input_at_prompt("Enter the number of days late:")
	interest_amount := calculate_interest(days_late, invoice_amount, MONTHLY_RATE)

	fmt.Println("Interest Owed:", fmt.Sprintf("%.2f", interest_amount))
}

func display_instructions(monthly_rate float64) {
	fmt.Println("Calculate interest on a late invoice.")
	fmt.Println("This assumes", monthly_rate, "monthly interest.")
}

func get_numeric_input_at_prompt(prompt string) float64 {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	captured_value, _ := strconv.ParseFloat(scanner.Text(), 64)

	if scanner.Err() != nil {
		fmt.Println("There was an error")
	}

	return captured_value
}

func calculate_interest(days_late float64, invoice_amount float64, monthly_rate float64) float64 {
	const MONTHS_IN_A_YEAR float64 = 12.0
	const DAYS_IN_A_YEAR float64 = 365.0

	yearly_rate := monthly_rate * MONTHS_IN_A_YEAR
	yearly_multiplier := yearly_rate / 100.0
	year_portion := days_late / DAYS_IN_A_YEAR
	partial_rate := year_portion * yearly_multiplier

	return invoice_amount * partial_rate
}
