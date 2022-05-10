package tests

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductIngredientRepoTestSuite struct {
	suite.Suite
	ProductIngredientRepo *repositories.ProductIngredientRepo
	mock                  sqlmock.Sqlmock
}

func (suite *ProductIngredientRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.ProductIngredientRepo = repositories.NewProductIngredientRepo(db)
}

func TestProductIngredientRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductIngredientRepoTestSuite))
}

func (suite *ProductIngredientRepoTestSuite) TestInsert() {
	var pId, iId int64 = 1, 1
	query := "INSERT INTO product_ingredients(product_id, ingredient_id) VALUES(?, ?)"
	suite.mock.ExpectExec(query).WithArgs(pId, iId).WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.ProductIngredientRepo.Insert(pId, iId)
	assert.NoError(suite.T(), err)
}
