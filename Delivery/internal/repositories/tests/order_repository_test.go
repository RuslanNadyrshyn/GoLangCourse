package tests

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	"Delivery/Delivery/internal/repositories/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type OrderRepoTestSuite struct {
	suite.Suite
	OrderRepo *repositories.OrderRepo
	mock      sqlmock.Sqlmock
}

var o = &models.Order{
	Id:      1,
	Price:   1.5,
	UserId:  1,
	Address: "TestAddress",
}

func (suite *OrderRepoTestSuite) SetupSuite() {
	db, mock := Connection.NewMock()
	suite.mock = mock
	suite.OrderRepo = repositories.NewOrderRepo(db)
}

func TestOrderRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepoTestSuite))
}

func (suite *OrderRepoTestSuite) TestInsert() {
	query := "INSERT orders(price, user_id, address) VALUES(?, ?, ?)"

	suite.mock.ExpectExec(query).WithArgs(o.Price, o.UserId, o.Address).WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := suite.OrderRepo.Insert(o)
	assert.NoError(suite.T(), err)
}

func (suite *OrderRepoTestSuite) TestGetById() {
	query := "SELECT id, price, user_id, address, created_at FROM orders WHERE id = (?)"
	rows := sqlmock.NewRows([]string{"id", "price", "user_id", "address", "created_at"}).
		AddRow(o.Id, o.Price, o.UserId, o.Address, time.Now().String())
	suite.mock.ExpectQuery(query).WithArgs(o.Id).WillReturnRows(rows)

	order, err := suite.OrderRepo.GetById(o.Id)
	assert.NotNil(suite.T(), order)
	assert.NoError(suite.T(), err)

	order, err = suite.OrderRepo.GetById(2)
	assert.Nil(suite.T(), order)
	assert.Error(suite.T(), err)
}

func (suite *OrderRepoTestSuite) TestGetByUserId() {
	query := "SELECT id, price, user_id, address, created_at " +
		"FROM orders WHERE user_id = (?)"

	rows := sqlmock.NewRows([]string{"id", "price", "user_id", "address", "created_at"}).
		AddRow(o.Id, o.Price, o.UserId, o.Address, time.Now().String())
	suite.mock.ExpectQuery(query).WithArgs(o.UserId).WillReturnRows(rows)

	order, err := suite.OrderRepo.GetByUserId(o.UserId)
	assert.NotNil(suite.T(), order)
	assert.NoError(suite.T(), err)

	order, err = suite.OrderRepo.GetByUserId(2)
	assert.Nil(suite.T(), order)
	assert.Error(suite.T(), err)
}
