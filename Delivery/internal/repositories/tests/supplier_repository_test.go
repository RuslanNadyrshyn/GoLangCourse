package tests

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	"Delivery/Delivery/internal/repositories/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SupplierRepoTestSuite struct {
	suite.Suite
	SupplierRepo *repositories.SupplierRepo
	mock         sqlmock.Sqlmock
}

var s = &models.Supplier{
	Id:    1,
	Name:  "TestName",
	Type:  "TestType",
	Image: "TestImage",
	WorkingHours: struct {
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	}{
		Opening: "TestOpening",
		Closing: "TestClosing",
	},
}

func (suite *SupplierRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.SupplierRepo = repositories.NewSupplierRepo(db)
}

func TestSupplierRepoTestSuite(t *testing.T) {
	suite.Run(t, new(SupplierRepoTestSuite))
}

func (suite *SupplierRepoTestSuite) TestInsert() {
	query := "INSERT INTO suppliers(id, name, type, image, opening, closing) VALUES(?, ?, ?, ?, ?, ?)"

	suite.mock.ExpectExec(query).
		WithArgs(s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing).
		WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.SupplierRepo.Insert(s)
	assert.NoError(suite.T(), err)
}

func (suite *SupplierRepoTestSuite) TestGetAll() {
	query := "SELECT id, name, type, image, opening, closing FROM suppliers"
	rows := sqlmock.NewRows([]string{"id", "name", "type", "image", "opening", "closing"}).
		AddRow(s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)
	suite.mock.ExpectQuery(query).WillReturnRows(rows)

	supplier, err := suite.SupplierRepo.GetAll()
	assert.NotNil(suite.T(), supplier)
	assert.NotEmpty(suite.T(), supplier)
	assert.NoError(suite.T(), err)
}

func (suite *SupplierRepoTestSuite) TestGetById() {
	query := "SELECT id, name, type, image, opening, closing FROM suppliers WHERE id = ?"
	rows := sqlmock.NewRows([]string{"id", "name", "type", "image", "opening", "closing"}).
		AddRow(s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)

	suite.mock.ExpectQuery(query).WithArgs(s.Id).WillReturnRows(rows)

	supplier, err := suite.SupplierRepo.GetById(s.Id)
	assert.NotNil(suite.T(), supplier)
	assert.NoError(suite.T(), err)

	supplier, err = suite.SupplierRepo.GetById(2)
	assert.Error(suite.T(), err)
}
