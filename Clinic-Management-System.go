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

type Patients [NMAX]Patient

type Visits [NMAX]Visit

var countPatient int = 0

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
	var choice int
	fmt.Println("===  List of Patients  ===")
	if countPatient == 0 {
		fmt.Println("No patients found")
	}
	for i := 0; i < countPatient; i++ {
		fmt.Printf("ID: %d, Name: %s, Date: %d\n", patients[i].ID, patients[i].Name, visits[i].Date)
	}
	fmt.Println("===   Patients Menu    ===")
	fmt.Println("1. Search for Patient (Name or ID)")
	fmt.Println("2. Add Patient")
	fmt.Println("3. Back to Main Menu")
	fmt.Print("Choose: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		searchPatient()
	case 2:
		addPatient()
	case 3:
		main()
	}
}

func searchPatient() {
	var method int
	fmt.Println("Search by:")
	fmt.Println("1. Name (Sequential Search)")
	fmt.Println("2. ID   (Binary Search)")
	fmt.Print("Choose: ")
	fmt.Scan(&method)

	switch method {

	//Sequential search by Name
	case 1:
		var targetName string
		fmt.Print("Enter patient name: ")
		fmt.Scan(&targetName)

		found := false
		for i := 0; i < countPatient; i++ {
			if patients[i].Name == targetName {
				fmt.Printf("Found — ID: %d | Name: %s | Date: %d | Service: %s | Cost: %.2f\n",
					patients[i].ID, patients[i].Name,
					visits[i].Date, visits[i].Service, visits[i].Cost)
				found = true
			}
		}
		if !found {
			fmt.Println("Patient not found.")
		}

	// Binary search by ID
	case 2:
    var targetID int
    fmt.Print("Enter patient ID: ")
    fmt.Scan(&targetID)

	sortedIdx := make([]int, countPatient)
    for i := 0; i < countPatient; i++ {
        sortedIdx[i] = i
    }
    sort.Slice(sortedIdx, func(a, b int) bool {
        return patients[sortedIdx[a]].ID < patients[sortedIdx[b]].ID
    })

    // Binary search, loop continues until range is finished or match is found
    lo, hi, foundIdx := 0, countPatient-1, -1
    for lo <= hi && foundIdx == -1 {
        mid := (lo + hi) / 2
        midID := patients[sortedIdx[mid]].ID
        if midID == targetID {
            foundIdx = sortedIdx[mid]
        } else if midID < targetID {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }

    if foundIdx != -1 {
        fmt.Printf("Found — ID: %d | Name: %s | Date: %d | Service: %s | Cost: %.2f\n",
            patients[foundIdx].ID, patients[foundIdx].Name,
            visits[foundIdx].Date, visits[foundIdx].Service, visits[foundIdx].Cost)
    } else {
        fmt.Println("Patient not found.")
    }

	default:
		fmt.Println("Invalid choice.")
	}
}
//Adding new patients to the list
func addPatient() {
	var n, i, limit int
	//Dealt with the user inputs, To the amouth of patients added and the data of the patients 
	fmt.Print("How many patients do you want to add? ")
	fmt.Scan(&n)
	limit = countPatient + n
	//Loops till the limit
	for i = countPatient; i < limit; i++ {
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
		countPatient = countPatient + 1
	}
}
//The user could edit patient data by using name or ID
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
//Deleting patient data by using the ID
func deletePatient() {
	var target int
	var found, i int
	
	fmt.Println("Patient ID:")
	fmt.Scan(&target)
	found = -1
	i = 0
	for i < countPatient && found = -1 {
		if Patients[i].ID == target {
			found = i
		}
		i = i + 1
	}
	
	if found != -1 {
		for i = found; i < countPatient; i++ {
			Patients[i] = Patients[i+1]
			Visits[i] = Visits[i+1]
		}
		countPatient = countPatient - 1
	} else {
		fmt.Parintln("Patient not found")
	}
}
func addVisit() {
	
}
//Showing the most used service and total visits in a day
func statistics() {
	
}
