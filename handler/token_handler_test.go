package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
	mock_service "stlab.itechart-group.com/go/food_delivery/authorization_service/service/mocks"
	"testing"
)

func TestHandler_refreshToken(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                string
		headerName          string
		headerValue         string
		token               string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			headerName:  "refresh_token",
			headerValue: "token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().RefreshTokens(token).Return(&authProto.GeneratedTokens{
					AccessToken:  "accessToken",
					RefreshToken: "refreshToken",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"accessToken":"accessToken","refreshToken":"refreshToken"}`,
		},
		{
			name:                "Empty refresh header",
			headerName:          "refresh_token",
			headerValue:         "",
			token:               "",
			mockBehavior:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"empty refresh header"}`,
		},
		{
			name:        "Invalid refresh token",
			headerName:  "refresh_token",
			headerValue: "token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().RefreshTokens(token).Return(nil, errors.New("invalid token"))
			},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"invalid token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			get := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(get, testCase.token)
			logger := logging.GetLogger()
			services := &service.Service{Authorization: get}
			handler := NewHandler(services, logger)

			//Init server
			r := gin.New()
			r.GET("/refresh", handler.refreshToken)

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", "/refresh", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
