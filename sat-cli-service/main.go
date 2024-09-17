package satcliservice

import (
	"fmt"
	"sat-result/db"
	"sat-result/service"
)

func Run() {

	db.Connect()

	for {
		
		fmt.Println("--------------------------------------------------------")
		fmt.Println("Here is your menu to select option from->")
		fmt.Println()
		var option int
		fmt.Println("1. Insert SAT data")
		fmt.Println("2. View all SAT data")
		fmt.Println("3. Get SAT rank")
		fmt.Println("4. Update SAT score")
		fmt.Println("5. Delete SAT record")
		fmt.Println("6. Exit")
		fmt.Println("--------------------------------------------------------")

		fmt.Println("Please enter option number from above menu : ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			service.InsertData()
			fmt.Println()
		case 2:
			service.ViewData()
			fmt.Println()
		case 3:
			service.GetRank()
			fmt.Println()
		case 4:
			service.UpdateData()
			fmt.Println()
		case 5:
			service.DeleteRecord()
			fmt.Println()
		case 6:
			return
		default:
			fmt.Println("Invalid option number entered. Please Try again!")
		}

	}

}
