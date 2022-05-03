package handlers

import (
	config "Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/handlers"
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	"Delivery/Delivery/internal/services"
	"Delivery/Delivery/tests"
	"Delivery/Delivery/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

const userID = 1

type UserHandlerTestSuite struct {
	suite.Suite
	accessToken string
	userHandler *handlers.UserHandler
	testSrv     *httptest.Server
	//user        *models.User
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	cfg := &config.Config{
		AccessSecret:          "access",
		AccessLifetimeMinutes: 1,
	}
	tokenService := services.NewTokenService(cfg)
	conn, _ := Connection.Connect()

	//suite.user = &models.User{
	//	Name:         "TestName",
	//	Email:        "TestEmail",
	//	PasswordHash: "TestPassword",
	//}

	suite.userHandler = handlers.NewUserHandler(tokenService, repositories.NewUserRepository(conn))
	suite.accessToken, _ = tokenService.GenerateAccessToken(userID)

	suite.testSrv = tests.Start(cfg, conn)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestWalkUserHandlerSignIn() {
	//t := suite.T()
	//handlerFunc := suite.userHandler.SignIn
	//
	//cases := []helpers.TestCaseHandler{
	//	{
	//		TestName: "Successfully sign in user profile",
	//		Request: helpers.Request{
	//			Method:    http.MethodPost,
	//			Url:       "/sign_in",
	//			AuthToken: "",
	//		},
	//		HandlerFunc: handlerFunc,
	//		Want: helpers.ExpectedSingInResponse{
	//			StatusCode: 200,
	//			User:       suite.user,
	//		},
	//	},
	//}
}

func (suite *UserHandlerTestSuite) TestWalkUserHandlerGetProfile() {
	t := suite.T()
	handlerFunc := suite.userHandler.GetProfile
	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "test-1@example.com",
			},
		},
		{
			TestName: "Unauthorized getting user profile",
			Request: helpers.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: "",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "invalid",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request, recorder := helpers.PrepareHandlerTestCase(test)
			test.HandlerFunc(recorder, request)

			assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
			if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
				if recorder.Code == http.StatusOK {
					helpers.AssertUserProfileResponse(t, recorder)
				}
			}
		})
	}
}

func (suite *UserHandlerTestSuite) TestWalkApiGetProfile() {
	t := suite.T()
	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "test-1@example.com",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			resp, err := suite.testSrv.Client().Get(suite.testSrv.URL + test.Request.Url)
			if assert.NoError(t, err) {
				assert.Equal(t, test.Want.StatusCode, resp.StatusCode)
			}
		})
	}
}
