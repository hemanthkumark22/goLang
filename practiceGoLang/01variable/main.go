package main

import "fmt"

func variableData() {
	// var companyName string = "TechM"
	// fmt.Println("CompanyName:" + companyName)
	// var companyID int = 12345
	// fmt.Println("CompanyID:", companyID)
	// var companyRevenue float64 = 123.456
	// fmt.Println("CompanyRevenue:", companyRevenue)
	var (
		empName  = "hemanth"
		empName  = "hemanth"
		salaryPA = 150000.50
	)

	fmt.Printf("%s getting %.2f$ of salaryPA\n", empName, salaryPA)

}
