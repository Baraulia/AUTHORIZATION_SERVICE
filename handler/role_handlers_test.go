package handler

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
	mock_service "stlab.itechart-group.com/go/food_delivery/authorization_service/service/mocks"
	"testing"
)

func TestHandler_getRoleById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRolePerm, id int)

	testTable := []struct {
		name                string
		input               string
		id                  int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:  "OK",
			input: "1",
			id:    1,
			mockBehavior: func(s *mock_service.MockRolePerm, id int) {
				s.EXPECT().GetRoleById(id).Return(&model.Role{
					ID:   1,
					Name: "test",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"name":"test"}`,
		},
		{
			name:                "invalid request",
			input:               "a",
			mockBehavior:        func(s *mock_service.MockRolePerm, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
		},
		{
			name:  "non-existent id",
			input: "1",
			id:    1,
			mockBehavior: func(s *mock_service.MockRolePerm, id int) {
				s.EXPECT().GetRoleById(id).Return(nil, fmt.Errorf("server error"))
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
			getRole := mock_service.NewMockRolePerm(c)
			testCase.mockBehavior(getRole, testCase.id)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getRole}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/roles/%s", testCase.input), nil)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_getAllRoles(t *testing.T) {
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
				s.EXPECT().GetAllRoles().Return([]model.Role{
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
			expectedRequestBody: `{"Roles":[{"id":1,"name":"test"},{"id":2,"name":"test2"}]}`,
		},
		{
			name: "server error",
			mockBehavior: func(s *mock_service.MockRolePerm) {
				s.EXPECT().GetAllRoles().Return(nil, fmt.Errorf("server error"))
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
			getRoles := mock_service.NewMockRolePerm(c)
			testCase.mockBehavior(getRoles)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getRoles}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", "/roles/", nil)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_createRole(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRolePerm, role *model.CreateRole)

	testTable := []struct {
		name                string
		inputBody           string
		inputRole           *model.CreateRole
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"test"}`,
			inputRole: &model.CreateRole{Name: "test"},
			mockBehavior: func(s *mock_service.MockRolePerm, role *model.CreateRole) {
				s.EXPECT().CreateRole(role.Name).Return(1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "server error",
			inputBody: `{"name":"test"}`,
			inputRole: &model.CreateRole{Name: "test"},
			mockBehavior: func(s *mock_service.MockRolePerm, role *model.CreateRole) {
				s.EXPECT().CreateRole(role.Name).Return(0, fmt.Errorf("server error"))
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
			getRole := mock_service.NewMockRolePerm(c)
			testCase.mockBehavior(getRole, testCase.inputRole)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getRole}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/roles/", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_bindRoleWithPerms(t *testing.T) {
	type mockBehavior func(s *mock_service.MockRolePerm, role *model.BindRoleWithPermission)

	testTable := []struct {
		name               string
		inputRoleId        string
		inputBody          string
		input              *model.BindRoleWithPermission
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"role_id":1,"permissions_id":[1,2,3]}`,
			input: &model.BindRoleWithPermission{
				RoleId:        1,
				PermissionsId: []int{1, 2, 3},
			},
			inputRoleId: "1",
			mockBehavior: func(s *mock_service.MockRolePerm, input *model.BindRoleWithPermission) {
				s.EXPECT().BindRoleWithPerms(input).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:               "Invalid request",
			inputBody:          `{"role_id":"a","permissions_id":[1,2,3]}`,
			mockBehavior:       func(s *mock_service.MockRolePerm, input *model.BindRoleWithPermission) {},
			expectedStatusCode: 400,
		},
		{
			name:      "server error",
			inputBody: `{"role_id":1,"permissions_id":[1,2,3]}`,
			input: &model.BindRoleWithPermission{
				RoleId:        1,
				PermissionsId: []int{1, 2, 3},
			},
			mockBehavior: func(s *mock_service.MockRolePerm, input *model.BindRoleWithPermission) {
				s.EXPECT().BindRoleWithPerms(input).Return(fmt.Errorf("server error"))
			},
			expectedStatusCode: 500,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			getRole := mock_service.NewMockRolePerm(c)
			testCase.mockBehavior(getRole, testCase.input)
			logger := logging.GetLogger()
			services := &service.Service{RolePerm: getRole}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", fmt.Sprintf("/roles/%s/perms", testCase.inputRoleId), bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}

}
