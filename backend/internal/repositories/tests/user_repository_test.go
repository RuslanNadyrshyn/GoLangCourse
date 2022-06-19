package tests

import (
	"Delivery/backend/internal/repositories"
	"Delivery/backend/internal/repositories/database/Connection"
	"Delivery/backend/internal/repositories/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserRepoTestSuite struct {
	suite.Suite
	UserRepo *repositories.UserRepo
	mock     sqlmock.Sqlmock
}

var u = &models.User{
	Id:           1,
	Name:         "TestName",
	Email:        "TestEmail",
	PasswordHash: "TestHash",
}

func (suite *UserRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.UserRepo = repositories.NewUserRepo(db)
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) TestInsert() {
	query := "INSERT users(name, email, password_hash) VALUES(?, ?, ?)"

	suite.mock.ExpectExec(query).WithArgs(u.Name, u.Email, u.PasswordHash).WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.UserRepo.Insert(u)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestInsertName() {
	query := "INSERT users(name) VALUES(?)"

	suite.mock.ExpectExec(query).WithArgs(u.Name).WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.UserRepo.InsertName(u)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestGetByEmail() {
	query := "SELECT id, email, name FROM users WHERE email = ?"
	rows := sqlmock.NewRows([]string{"id", "email", "name"}).
		AddRow(u.Id, u.Email, u.Name)
	suite.mock.ExpectQuery(query).WithArgs(u.Email).WillReturnRows(rows)

	user, err := suite.UserRepo.GetByEmail(u.Email)
	assert.NotNil(suite.T(), user)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestGetById() {
	query := "SELECT id, ifnull(email, ''), name FROM users WHERE id = ?"
	rows := sqlmock.NewRows([]string{"id", "email", "name"}).
		AddRow(u.Id, u.Email, u.Name)
	suite.mock.ExpectQuery(query).WithArgs(u.Id).WillReturnRows(rows)

	user, err := suite.UserRepo.GetById(u.Id)
	assert.NotNil(suite.T(), user)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestUpdate() {
	query := "UPDATE users SET name=?, email=?, password_hash=? WHERE id=?"
	updateUser := models.User{
		Id:           1,
		Name:         "UpdatedTestName",
		Email:        "UpdatedTestEmail",
		PasswordHash: "UpdatedPasswordHash",
	}

	suite.mock.ExpectExec(query).WithArgs(updateUser.Name, updateUser.Email, updateUser.PasswordHash, updateUser.Id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	id, err := suite.UserRepo.Update(&updateUser)
	assert.Equal(suite.T(), u.Id, id)
	assert.NoError(suite.T(), err)
}
