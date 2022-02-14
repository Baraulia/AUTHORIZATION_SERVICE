package handler

import (
	"bytes"
	"errors"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/service"
	mock_service "github.com/Baraulia/AUTHORIZATION_SERVICE/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)


func TestHandler_authUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user model.User)
	testTable := []struct {
		name                string         //the name of the test
		inputBody           string         //the body of the request
		inputUser           model.User //the structure which we send to the service
		mockBehavior        mockBehavior
		expectedStatusCode  int    //expected code
		expectedRequestBody string //expected response
	}{
		{
			name:      "OK",
			inputBody: `{"email":"aaaaaaa@gmail.com", "password":"aaa111"}`,
			inputUser: model.User{
				Email:    "aaaaaaa@gmail.com",
				Password: "aaa111",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().GenerateToken(user.Email, user.Password).Return("token", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"token":"token"}`,
		},
		{
			name:                "Empty fields",
			inputBody:           `{"email":"aaaaaaa@gmail.com"}`,
			inputUser:           model.User{},
			mockBehavior:        func(s *mock_service.MockAuthorization, user model.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "No Input",
			inputBody:            `{}`,
			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
			expectedStatusCode:   400,
			expectedRequestBody: `{"error":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"email":"aaaaaaa@gmail.com", "password":"aaa111"}`,
			inputUser: model.User{
				Email:    "aaaaaaa@gmail.com",
				Password: "aaa111",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().GenerateToken(user.Email, user.Password).Return("", errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)
			logger := logging.GetLogger()
			services := &service.Service{Authorization: auth}
			handler := NewHandler(services, logger)

			//Init server
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
