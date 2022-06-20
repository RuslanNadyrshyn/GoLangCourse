package filesystem

import (
	"encoding/json"
	"fmt"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
	"io"
	"os"
)

type UserFileRepository struct {
}

func (ufr UserFileRepository) Insert(user *models.User) error {
	file, err := os.Create("./internal/datastore/files/user2.json")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	str, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(str))
	file.WriteString(string(str))

	return nil
}

func (ufr UserFileRepository) GetByEmail(_ string) (user models.User) {
	userData := []byte{}

	file, err := os.Open("./internal/datastore/files/user1.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		chunk := make([]byte, 64)
		n, err := file.Read(chunk)
		if err == io.EOF { //
			break //
		}
		if err != nil {
			panic(err)
		}

		userData = append(userData, chunk[:n]...)
		fmt.Println(userData)
	}

	err = json.Unmarshal(userData, &user)
	if err != nil {
		panic(err)
	}

	return user
}
