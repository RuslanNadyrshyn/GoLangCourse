package tests

import (
	"Delivery/backend/internal/repositories"
	"Delivery/backend/internal/repositories/database/Connection"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MenuRepoTestSuite struct {
	suite.Suite
	MenuRepo *repositories.MenuRepo
	mock     sqlmock.Sqlmock
}

var menuId, supplierId int64 = 324, 33

func (suite *MenuRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.MenuRepo = repositories.NewMenuRepo(db)
}

func TestMenuRepoTestSuite(t *testing.T) {
	suite.Run(t, new(MenuRepoTestSuite))
}

func (suite *MenuRepoTestSuite) TestInsert() {
	query := "INSERT INTO menus(supplier_id) VALUES(?)"

	suite.mock.ExpectExec(query).WithArgs(supplierId).WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.MenuRepo.Insert(supplierId)
	assert.NoError(suite.T(), err)
}

func (suite *MenuRepoTestSuite) TestGetIdBySupId() {
	query := "SELECT id FROM menus WHERE supplier_id = ?"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(menuId)
	suite.mock.ExpectQuery(query).WithArgs(supplierId).WillReturnRows(rows)

	_, err := suite.MenuRepo.GetIdBySupId(supplierId)
	assert.NoError(suite.T(), err)
}

func (suite *MenuRepoTestSuite) TestGetSupIdById() {
	query := "SELECT supplier_id FROM menus WHERE id = ?"
	rows := sqlmock.NewRows([]string{"supplier_id"}).AddRow(supplierId)
	suite.mock.ExpectQuery(query).WithArgs(menuId).WillReturnRows(rows)

	_, err := suite.MenuRepo.GetSupIdById(menuId)
	assert.NoError(suite.T(), err)
}

//
//func (suite *MenuRepoTestSuite) TestUpdate() {
//	query := "UPDATE users SET name=?, email=?, password_hash=? WHERE id=?"
//	updateUser := models.User{
//		Id:           1,
//		Name:         "UpdatedTestName",
//		Email:        "UpdatedTestEmail",
//		PasswordHash: "UpdatedPasswordHash",
//	}
//
//	suite.mock.ExpectExec(query).WithArgs(updateUser.Name, updateUser.Email, updateUser.PasswordHash, updateUser.Id).
//		WillReturnResult(sqlmock.NewResult(0, 0))
//
//	id, err := suite.MenuRepo.Update(&updateUser)
//	assert.Equal(suite.T(), u.Id, id)
//	assert.NoError(suite.T(), err)
//}
