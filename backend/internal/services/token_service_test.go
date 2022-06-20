package services

import (
	config "github.com/RuslanNadyrshyn/GoLangCourse/backend/cfg"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/helper"
	helpers2 "github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/tests/helpers"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const userID int64 = 1

type TokenServiceTestSuite struct {
	suite.Suite
	cfg          *config.Config
	tokenService *TokenService
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	suite.cfg = &config.Config{
		AccessSecret:           "access",
		AccessLifetimeMinutes:  1,
		RefreshSecret:          "refresh",
		RefreshLifetimeMinutes: 1,
	}
	suite.tokenService = NewTokenService(suite.cfg)
}

func (suite *TokenServiceTestSuite) TearDownSuite() {
}

func (suite *TokenServiceTestSuite) SetupTest() {
}

func (suite *TokenServiceTestSuite) TearDownTest() {
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}

func (suite *TokenServiceTestSuite) TestGetTokenFromBearerString() {
	testCases := []helpers2.TestCaseGetBearerToken{
		{
			Name:         "Get token successfully",
			BearerString: "Bearer test_token",
			Want:         "test_token",
		},
		{
			Name:         "Return empty token if Bearer prefix is absent",
			BearerString: "Beare test_token",
			Want:         "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			got := suite.tokenService.GetTokenFromBearerString(testCase.BearerString)
			assert.Equal(suite.T(), testCase.Want, got)
		})
	}
}

func (suite *TokenServiceTestSuite) TestGenerateAccessToken() {
	tokenString, err := suite.tokenService.GenerateAccessToken(userID)

	assert.NoError(suite.T(), err)

	token, err := jwt.ParseWithClaims(tokenString, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(suite.cfg.AccessSecret), nil
	})
	assert.NoError(suite.T(), err)

	claims, ok := token.Claims.(*helper.JwtCustomClaims)
	assert.True(suite.T(), ok)
	assert.True(suite.T(), token.Valid)

	got := claims.ID
	assert.Equal(suite.T(), userID, got)

	expireTime := time.Unix(claims.ExpiresAt, 0)
	assert.WithinDuration(suite.T(), time.Now().Add(time.Minute*time.Duration(suite.cfg.AccessLifetimeMinutes)), expireTime, time.Second)
}

func (suite *TokenServiceTestSuite) TestValidateAccessToken() {
	tokenString, _ := suite.tokenService.GenerateAccessToken(userID)
	refreshTokenString, _ := suite.tokenService.GenerateRefreshToken(userID)
	invalidTokenString := tokenString + "a"

	suite.cfg.AccessLifetimeMinutes = 0
	expiredTokenString, _ := suite.tokenService.GenerateAccessToken(userID)

	testCases := []helpers2.TestCaseValidate{
		{
			Name:         "Valid token",
			AccessToken:  tokenString,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			Name:         "Invalid token",
			AccessToken:  invalidTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
		{
			Name:         "Token with non-expected signature",
			AccessToken:  refreshTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
		{
			Name:         "Expired token",
			AccessToken:  expiredTokenString,
			WantError:    true,
			WantErrorMsg: "token is expired",
			WantID:       0,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotClaims, err := suite.tokenService.ValidateAccessToken(testCase.AccessToken)

			helpers2.AssertTokenResponse(suite.T(), testCase, gotClaims, err)
		})
	}
}

func (suite *TokenServiceTestSuite) TestRefreshToken() {
	accessToken, _ := suite.tokenService.GenerateAccessToken(userID)
	refreshToken, _ := suite.tokenService.GenerateRefreshToken(userID)

	testCases := []helpers2.TestCaseRefresh{
		{
			Name:         "Generated new valid refresh tokens",
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			newAccToken, newRefToken, err := suite.tokenService.RefreshToken(
				testCase.AccessToken,
				testCase.RefreshToken,
				suite.cfg.AccessSecret,
				suite.cfg.RefreshSecret,
			)
			token, err := jwt.ParseWithClaims(newAccToken, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(suite.cfg.AccessSecret), nil
			})
			assert.NoError(suite.T(), err)
			claims, ok := token.Claims.(*helper.JwtCustomClaims)
			assert.True(suite.T(), ok)
			assert.True(suite.T(), token.Valid)
			got := claims.ID
			assert.Equal(suite.T(), userID, got)

			token, err = jwt.ParseWithClaims(newRefToken, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(suite.cfg.RefreshSecret), nil
			})
			assert.NoError(suite.T(), err)
			claims, ok = token.Claims.(*helper.JwtCustomClaims)
			assert.True(suite.T(), ok)
			assert.True(suite.T(), token.Valid)
			got = claims.ID
			assert.Equal(suite.T(), userID, got)
			_, err = suite.tokenService.ValidateAccessToken(newAccToken)
			assert.NoError(suite.T(), err)
			_, err = suite.tokenService.ValidateRefreshToken(newRefToken)
			assert.NoError(suite.T(), err)
		})
	}
}
