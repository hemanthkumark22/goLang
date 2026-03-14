package main

import "fmt"

func main() {
	mvariableData()
}

func mvariableData() {
	// var companyName string = "TechM"
	// fmt.Println("CompanyName:" + companyName)
	// var companyID int = 12345
	// fmt.Println("CompanyID:", companyID)
	// var companyRevenue float64 = 123.456
	// fmt.Println("CompanyRevenue:", companyRevenue)
	var (
		empName      = "hemanth"
		salaryPA     = 150000.80
		empName1     = "hari"
		emp1salaryPA = 160000
	)

	fmt.Printf("%s getting %.2f of salaryPA\n", empName, salaryPA)
	fmt.Printf("%s getting %d$ of salaryPA\n", empName1, emp1salaryPA)
}
