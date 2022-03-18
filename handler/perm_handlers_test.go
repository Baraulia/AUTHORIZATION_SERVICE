package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
	mock_service "stlab.itechart-group.com/go/food_delivery/authorization_service/service/mocks"
	"testing"
)

func TestHandler_getPermsByRoleId(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAuthorization, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAuthorization, token string)
	type mockBehavior func(s *mock_service.MockRolePerm, id int)

	testTable := []struct {
		name                   string
		input                  string
		id                     int
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehavior           mockBehavior
		mockBehaviorCheck      mockBehaviorCheck
		expectedStatusCode     int
		expectedRequestBody    string
	}{
		{
			name:       "OK",
			input:      "1",
			id:         1,
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthorization, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockRolePerm, id int) {
				s.EXPECT().GetPermsByRoleId(id).Return([]model.Permission{
					{
						ID:   1,
						Name: "test",
					},
					{
						ID:   2,
						Name: "test2",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Perms":[{"id":1,"name":"test"},{"id":2,"name":"test2"}]}`,
		},
		{
			name:         "invalid request",
			input:        "a",
			id:           1,
			mockBehavior: func(s *mock_service.MockRolePerm, id int) {},

			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
		},
		{
			name:  "server error",
			input: "1",
			id:    1,
			mockBehavior: func(s *mock_service.MockRolePerm, id int) {
				s.EXPECT().GetPermsByRoleId(id).Return(nil, fmt.Errorf("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			getPerm := mock_service.NewMockRolePerm(c)
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviorParseToken(auth, testCase.inputToken)
			testCase.mockBehaviorCheck(auth, testCase.inputPerms, testCase.inputRole)
			testCase.mockBehavior(getPerm, testCase.id)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getPerm}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/roles/%s/perms", testCase.input), nil)
			req.Header.Set("Authorization", "Bearer testToken")

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_getAllPerms(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRolePerm)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockRolePerm) {
				s.EXPECT().GetAllPerms().Return([]model.Permission{
					{
						ID:   1,
						Name: "test",
					},
					{
						ID:   2,
						Name: "test2",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Perms":[{"id":1,"name":"test"},{"id":2,"name":"test2"}]}`,
		},
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockRolePerm) {
				s.EXPECT().GetAllPerms().Return([]model.Permission{
					{
						ID:   1,
						Name: "test",
					},
					{
						ID:   2,
						Name: "test2",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Perms":[{"id":1,"name":"test"},{"id":2,"name":"test2"}]}`,
		},
		{
			name: "server error",
			mockBehavior: func(s *mock_service.MockRolePerm) {
				s.EXPECT().GetAllPerms().Return(nil, fmt.Errorf("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			getPerm := mock_service.NewMockRolePerm(c)
			testCase.mockBehavior(getPerm)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getPerm}
			handler := NewHandler(services, logger)

			//Init server
			r := gin.New()
			r.GET("/perms/", handler.getAllPerms)

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", "/perms/", nil)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_createPerm(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRolePerm, role *model.CreatePerm)

	testTable := []struct {
		name                string
		inputBody           string
		inputPerm           *model.CreatePerm
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"test"}`,
			inputPerm: &model.CreatePerm{Name: "test"},
			mockBehavior: func(s *mock_service.MockRolePerm, perm *model.CreatePerm) {
				s.EXPECT().CreatePermission(perm.Name).Return(1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Invalid request",
			inputBody:           `{"name":1}`,
			inputPerm:           &model.CreatePerm{},
			mockBehavior:        func(s *mock_service.MockRolePerm, perm *model.CreatePerm) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
		},
		{
			name:      "server error",
			inputBody: `{"name":"test"}`,
			inputPerm: &model.CreatePerm{Name: "test"},
			mockBehavior: func(s *mock_service.MockRolePerm, perm *model.CreatePerm) {
				s.EXPECT().CreatePermission(perm.Name).Return(0, fmt.Errorf("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			getPerm := mock_service.NewMockRolePerm(c)
			testCase.mockBehavior(getPerm, testCase.inputPerm)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getPerm}
			handler := NewHandler(services, logger)

			//Init server
			r := gin.New()
			r.POST("/perms/", handler.createPerm)

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/perms/", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
