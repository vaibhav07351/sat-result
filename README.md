# sat-result

SAT Results Application

A Go application to manage SAT results using GORM and PostgreSQL. This CLI-based application allows you to insert, view, update, and delete student records, as well as retrieve student rankings based on SAT scores.

Features
Insert Data: Add new student records with SAT scores and calculate pass/fail status.
View All Data: Display all student records in JSON format.
Get Rank: Retrieve the rank of a student based on their SAT score.
Update Score: Update the SAT score of a student.
Delete Record: Remove a student record by name.

Prerequisites
Go (1.16 or later)
PostgreSQL
GORM (Go ORM library)
PostgreSQL driver for GORM

Installation:
git clone https://github.com/vaibhav07351/sat-result.git
cd sat-result

Setup PostgreSQL
Ensure PostgreSQL is installed and running.
Create a new database:
createdb database sat_result;

go mod tidy

go run .
