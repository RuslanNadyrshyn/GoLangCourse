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

type IngredientRepoTestSuite struct {
	suite.Suite
	IngredientRepo *repositories.IngredientRepo
	mock           sqlmock.Sqlmock
}

var i = models.Ingredient{
	Id:   1,
	Name: "TestName",
}

func (suite *IngredientRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.IngredientRepo = repositories.NewIngredientRepo(db)
}

func TestIngredientRepoTestSuite(t *testing.T) {
	suite.Run(t, new(IngredientRepoTestSuite))
}

func (suite *IngredientRepoTestSuite) TestInsert() {
	query := "INSERT IGNORE INTO ingredients(name) VALUES (?)"

	suite.mock.ExpectExec(query).WithArgs(i.Name).WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.IngredientRepo.Insert(i.Name)
	assert.NoError(suite.T(), err)
}

func (suite *IngredientRepoTestSuite) TestGetById() {
	query := "SELECT name FROM ingredients WHERE id = (?)"

	rows := sqlmock.NewRows([]string{"name"}).AddRow(i.Name)
	suite.mock.ExpectQuery(query).WithArgs(i.Id).WillReturnRows(rows)

	_, err := suite.IngredientRepo.GetById(i.Id)

	assert.NoError(suite.T(), err)
}
