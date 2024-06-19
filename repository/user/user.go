package users

import (
	"database/sql"
	"fmt"

	"github.com/pradiptarana/book-online-store/model"
)

// UserRepository is responsible for storing and retrieving user information
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(
	db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// SignUp adds a new user to the repository
func (ur *UserRepository) SignUp(us *model.User) error {
	stmt, err := ur.db.Prepare("insert into `user` (`username`, `password`) values (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(&us.Username, &us.Password)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// // Login verifies the username and password for login
// func (ur *UserRepository) Login(username, password string) (bool, error) {
// 	// Check if the username exists
// 	user, exists := ur.users[username]
// 	if !exists {
// 		return false, fmt.Errorf("user %s not found", username)
// 	}

// 	// Check if the password matches
// 	if user.Password == password {
// 		fmt.Printf("User %s logged in successfully!\n", username)
// 		return true, nil
// 	}

// 	return false, fmt.Errorf("incorrect password for user %s", username)
// }

func (ur *UserRepository) GetUser(username string) (*model.User, error) {
	stmt, err := ur.db.Prepare("select * from user where username = ?")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var result = &model.User{}
	err = stmt.QueryRow(username).Scan(&result.Id, &result.Username, &result.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}
