package main

import (
	"awesomeProject/Goroutines"
	"awesomeProject/internal/repositories/database"
	"awesomeProject/internal/repositories/database/Connection"
	"awesomeProject/internal/repositories/filesystem"
	"awesomeProject/internal/repositories/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	repo := filesystem.UserFileRepository{}

	user := repo.Insert(&models.User{
		Id:           0,
		Email:        "blabla@gmail.com",
		PasswordHash: "123",
		Name:         "Name",
	})

	fmt.Println(user)

	conn, err := Connection.Connect()
	if err != nil {
		log.Fatal(err)
	}

	//User
	userRepo := database.NewUserRepository(conn)
	userRepo.Insert(models.User{
		Email:        "Email2",
		PasswordHash: "Password",
		Name:         "Name",
	})

	newUser, err := userRepo.GetByEmail("Email")
	fmt.Println(newUser)

	//Supplier
	supplierRepo := database.NewSupplierRepository(conn)

	///////Goroutines
	test := Goroutines.GoRoutinesRepository{}
	sup := test.Insert("./Goroutines/Data/1.json")
	supplierRepo.Insert(sup)

	//////

	//newSupplier, err := supplierRepo.GetById(1)
	//fmt.Println(newSupplier)
	//
	//userRepo.Update(models.User{
	//	Email:        "NewEmail1",
	//	PasswordHash: "NewPass",
	//	Name:         "NewName",
	//}, "Email1")

}