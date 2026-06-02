package main

import "fmt"

type Patient struct {
	ID   int
	Name string
}

type Visit struct {
	Date    int
	Service string
	Cost    float64
}

const NMAX int = 999

var Patients [NMAX]Patient

var Visits [NMAX]Visit

var count int = 0

//Display the menu and handle the user inputs then connected to the other functions to the patients menue, show patient and statistics
func main() {
	var n int
	fmt.Println("=====Mangement Menu=====")
	fmt.Println("1. Patients")
	fmt.Println("2. Statistics")
	fmt.Println("3. Exit")
	fmt.Println("Choose: ")
	fmt.Scan(&n)

	switch n {
	case 1:
		patientsMenu()
	case 2:
		statistics()
	case 3:
		fmt.Println("Exiting")
	}
}

//Display the menu for the patients and handle the user inputs then connected to the other functions to add, edit, delete and back to the main menu
func patientsMenu() {
	var i, choice int
	fmt.Println("===  List of Patients  ===")
	if countPatient == 0 {
		fmt.Println("No patients found")
	}
	for i := 0; i < countPatient; i++ {
		fmt.Printf("ID: %d, Name: %s, Visit: %d \n", Patients[i].ID, Patients[i].Name, Visits[i].Date)
	}
	fmt.Println("===   Patients Menu    ===")
	fmt.Println("1. Add Patient")
	fmt.Println("2. edit Patient")
	fmt.Println("3. Delete Patient")
	fmt.Println("4. Add Visit")
	fmt.Println("6. Back to Main Menu")
	fmt.Print("Choose: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		addPatient()
	case 2:
		editPatient()
	case 3:
		deletePatient()
	case 4:
		addVisit()
	case 5:
		main()
	}
}

func addPatient() {
	var n, i, limit int
	fmt.Print("How many patients do you want to add? ")
	fmt.Scan(&n)
	limit = count + n
	for i = count; i < limit; i++ {
		fmt.Print("ID: ")
		fmt.Scan(&Patients[i].ID)
		fmt.Print("Name: ")
		fmt.Scan(&Patients[i].Name)
		fmt.Print("Date: ")
		fmt.Scan(&Visits[i].Date)
		fmt.Print("Type of Service: ")
		fmt.Scan(&Visits[i].Service)
		fmt.Print("Cost: ")
		fmt.Scan(&Visits[i].Cost)
		count = count + 1
	}
}

func editPatient() {
	var Target string
	var IDs, found, i int
	found = -1
	fmt.Println(showPatient())
	fmt.Print("Enter patient ID or Name to edit: ")
	fmt.Scan(&Target)
	fmt.Scan(&IDs)
	for i = 0; i < count && found == -1; i++ {
		if Patients[i].Name == Target || Patients[i].ID == IDs {
			found = i
		}
	}
	if found != -1 {
		fmt.Print("ID: ")
		fmt.Scan(&Patients[found].ID)
		fmt.Print("Name: ")
		fmt.Scan(&Patients[found].Name)
		fmt.Print("Date: ")
		fmt.Scan(&Visits[found].Date)
		fmt.Print("Type of Service: ")
		fmt.Scan(&Visits[found].Service)
		fmt.Print("Cost: ")
		fmt.Scan(&Visits[found].Cost)
	} else {
		fmt.Println("Patient not found")
	}
}
func deletePatient() {

}
func showPatient() {

}
func statistics() {

}
