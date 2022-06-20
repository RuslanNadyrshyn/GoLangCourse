package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/database/Connection"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OrderProductRepoTestSuite struct {
	suite.Suite
	OrderProductRepo *repositories.OrderProductRepo
	mock             sqlmock.Sqlmock
}

var op = models.OrderProduct{
	Id:        1,
	OrderId:   1,
	ProductId: 1,
	Count:     1,
	Price:     1.5,
}

func (suite *OrderProductRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.OrderProductRepo = repositories.NewOrderProductRepo(db)
}

func TestOrderProductRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderProductRepoTestSuite))
}

func (suite *OrderProductRepoTestSuite) TestInsert() {

	query := "INSERT order_product(order_id, product_id, count, price) VALUES(?, ?, ?, ?)"

	suite.mock.ExpectExec(query).
		WithArgs(op.OrderId, op.ProductId, op.Count, op.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.OrderProductRepo.Insert(op.OrderId, op.ProductId, op.Count, op.Price)
	assert.NoError(suite.T(), err)
}

func (suite *OrderProductRepoTestSuite) TestGetByOrderId() {
	query := "SELECT id, order_id, product_id, count, price FROM order_product WHERE order_id = (?)"

	rows := sqlmock.NewRows([]string{"id", "order_id", "product_id", "count", "price"}).
		AddRow(op.Id, op.OrderId, op.ProductId, op.Count, op.Price)
	suite.mock.ExpectQuery(query).WithArgs(op.OrderId).WillReturnRows(rows)

	orderProd, err := suite.OrderProductRepo.GetByOrderId(op.OrderId)
	assert.NotNil(suite.T(), orderProd)
	assert.NoError(suite.T(), err)
}
