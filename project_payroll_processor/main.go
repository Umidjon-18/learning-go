package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculateSalary() float64
}

type SalaryEmployee struct {
	Name   string
	Annual float64
}

func (se SalaryEmployee) String() string {
	return fmt.Sprintf("Salary employee: %s (Annual: $%.2f)", se.Name, se.Annual)
}

func (se SalaryEmployee) CalculateSalary() float64 {
	return se.Annual / 12
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly employee: %s (Hourly rate: %.2f) (Hours worked: %.2f)", he.Name, he.HourlyRate, he.HoursWorked)
}

func (he HourlyEmployee) CalculateSalary() float64 {
	return he.HourlyRate * he.HoursWorked
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	SalesAmount    int
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission employee: %s (Base salary: %.2f) (Base salary: %.2f) (Base salary: %d)", ce.Name, ce.BaseSalary, ce.CommissionRate, ce.SalesAmount)
}

func (ce CommissionEmployee) CalculateSalary() float64 {
	return ce.BaseSalary + ce.CommissionRate*float64(ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("	-Processing: %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("---Processing Payroll---")
	totalPay := 0.0
	for _, employee := range employees {
		PrintEmployeeSummary(employee)
		pay := employee.CalculateSalary()
		fmt.Printf("Monthly pay: $%.2f\n", pay)
		totalPay += pay
	}
	fmt.Printf("Total monthly pay: $%.2f\n", totalPay)
	fmt.Println("------------------------")
}

func main() {
	se := SalaryEmployee{
		Name:   "Bob",
		Annual: 12000,
	}
	he := HourlyEmployee{
		Name:        "John",
		HourlyRate:  120,
		HoursWorked: 80,
	}
	ce := CommissionEmployee{
		Name:           "Dave",
		BaseSalary:     1000,
		CommissionRate: 100,
		SalesAmount:    100,
	}

	payrollList := []Payable{
		se,
		he,
		ce,
	}
	ProcessPayroll(payrollList)
}
