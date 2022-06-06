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

type ProductRepoTestSuite struct {
	suite.Suite
	ProductRepo *repositories.ProductRepo
	mock        sqlmock.Sqlmock
}

var p = &models.Product{
	Id:     1,
	Name:   "TestName",
	MenuId: 1,
	Price:  1.5,
	Image:  "TestImage",
	Type:   "TestType",
}

func (suite *ProductRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.ProductRepo = repositories.NewProductRepo(db)
}

func TestUProductRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepoTestSuite))
}

func (suite *ProductRepoTestSuite) TestInsert() {
	query := "INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)"

	suite.mock.ExpectExec(query).
		WithArgs(p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image).
		WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.ProductRepo.Insert(p)
	assert.NoError(suite.T(), err)
}

func (suite *ProductRepoTestSuite) TestUpdatePrice() {
	query := "UPDATE products SET price=? WHERE id=?"

	suite.mock.ExpectExec(query).WithArgs(p.Price, p.Id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	_, err := suite.ProductRepo.UpdatePrice(p)
	assert.NoError(suite.T(), err)
}

//
func (suite *ProductRepoTestSuite) TestGetAll() {
	query := "SELECT id, menu_id, name, price, image, type FROM products"
	rows := sqlmock.NewRows([]string{"id", "menu_id", "name", "price", "image", "type"}).
		AddRow(p.Id, p.MenuId, p.Name, p.Price, p.Image, p.Type)
	suite.mock.ExpectQuery(query).WillReturnRows(rows)

	products, err := suite.ProductRepo.GetAll()
	assert.NotNil(suite.T(), products)
	assert.NotEmpty(suite.T(), products)
	assert.NoError(suite.T(), err)
}

func (suite *ProductRepoTestSuite) TestGetById() {
	query := "SELECT id, menu_id, name, price, image, type FROM products WHERE id = (?)"
	rows := sqlmock.NewRows([]string{"id", "menu_id", "name", "price", "image", "type"}).
		AddRow(p.Id, p.MenuId, p.Name, p.Price, p.Image, p.Type)
	suite.mock.ExpectQuery(query).WithArgs(p.Id).WillReturnRows(rows)

	product, err := suite.ProductRepo.GetById(p.Id)
	assert.Equal(suite.T(), product, p)
	assert.NotNil(suite.T(), product)
	assert.NoError(suite.T(), err)
}

func (suite *ProductRepoTestSuite) TestGetByMenuId() {
	query := "SELECT id, menu_id, name, price, image, type FROM products WHERE menu_id = (?)"
	rows := sqlmock.NewRows([]string{"id", "menu_id", "name", "price", "image", "type"}).
		AddRow(p.Id, p.MenuId, p.Name, p.Price, p.Image, p.Type)
	suite.mock.ExpectQuery(query).WithArgs(p.MenuId).WillReturnRows(rows)

	products, err := suite.ProductRepo.GetByMenuId(p.MenuId)
	assert.Equal(suite.T(), p.Name, products[0].Name)
	assert.NotNil(suite.T(), products)
	assert.NoError(suite.T(), err)
}
