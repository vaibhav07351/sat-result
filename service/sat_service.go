package service

import (
	"encoding/json"
	"fmt"
	"sat-result/db"
	"sat-result/model"
)

func InsertData() {
	var student model.Student

	fmt.Println("Enter student name:")
	fmt.Scanln(&student.Name)
	fmt.Println("Enter address:")
	fmt.Scanln(&student.Address)
	fmt.Println("Enter city:")
	fmt.Scanln(&student.City)
	fmt.Println("Enter country:")
	fmt.Scanln(&student.Country)
	fmt.Println("Enter pincode:")
	fmt.Scanln(&student.Pincode)
	fmt.Println("Enter SAT score:")
	fmt.Scanln(&student.SATScore)

	// Calculate Pass/Fail
	if student.SATScore > 30 {
		student.Passed = "Pass"
	} else {
		student.Passed = "Fail"
	}

	// Insert into the database using GORM
	result := db.Db.Create(&student)
	if result.Error != nil {
		fmt.Println("Failed to insert student:", result.Error)
	} else {
		fmt.Println("Student inserted successfully!")
	}
}

func ViewData() {
	var students []model.Student

	result := db.Db.Find(&students)
	if result.Error != nil {
		fmt.Println("Error retrieving data:", result.Error)
		return
	}

	if len(students)==0{
		fmt.Println("No student records exists!")
		return
	}

	// Convert student slice to JSON
    studentJSON, err := json.MarshalIndent(students, "", "    ")
    if err != nil {
        fmt.Println("Error converting data to JSON:", err)
        return
    }

    // Print JSON data
    fmt.Println(string(studentJSON))

}

func GetRank() {
	var name string
	fmt.Println("Enter student name:")
	fmt.Scanln(&name)
	var students []model.Student

    // Fetch all students by SATScore in descending order
    result := db.Db.Order("sat_score desc").Find(&students)
    if result.Error != nil {
        fmt.Println("Error retrieving data:", result.Error)
        return
    }

    
    rank := -1
    for i, student := range students {
        if student.Name == name {
            rank = i + 1 // Rank starts from 1
            break
        }
    }

    if rank == -1 {
        fmt.Println("Student not found.")
    } else {
        fmt.Printf("Rank of %s: %d\n", name, rank)
    }
}

func UpdateData() {
	var name string
	var newScore int
	fmt.Println("Enter student name:")
	fmt.Scanln(&name)
	fmt.Println("Enter new SAT score:")
	fmt.Scanln(&newScore)
	var student model.Student

	result := db.Db.First(&student, "name = ?", name)
	if result.Error != nil {
		fmt.Println("Student not found:", result.Error)
		return
	}

	// Update SAT score and pass/fail status
	student.SATScore = newScore
	if newScore > 30 {
		student.Passed = "Pass"
	} else {
		student.Passed = "Fail"
	}

	db.Db.Save(&student)
	fmt.Println("Student's SAT score updated successfully!")
}

func DeleteRecord() {
	var name string
	fmt.Println("Enter student name:")
	fmt.Scanln(&name)
	var student model.Student

	result := db.Db.First(&student, "name = ?", name)
	if result.Error != nil {
		fmt.Println("Student not found:", result.Error)
		return
	}

	result = db.Db.Where("name = ?", name).Delete(&student)
	if result.Error != nil {
		fmt.Println("Failed to delete student:", result.Error)
	} else {
		fmt.Println("Student deleted successfully!")
	}
}
