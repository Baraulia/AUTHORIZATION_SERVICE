package handler

import (
	"bytes"
	"fmt"
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

func TestHandler_getRoleById(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAuthUser, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAuthUser, token string)
	type mockBehavior func(s *mock_service.MockAuthUser, id int)

	testTable := []struct {
		name                   string
		input                  string
		id                     int
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehaviorCheck      mockBehaviorCheck
		mockBehavior           mockBehavior
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
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser, id int) {
				s.EXPECT().GetRoleById(id).Return(&model.Role{
					ID:   1,
					Name: "test",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"name":"test"}`,
		},
		{
			name:       "invalid request",
			input:      "a",
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior:        func(s *mock_service.MockAuthUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
		},
		{
			name:       "non-existent id",
			input:      "1",
			id:         1,
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser, id int) {
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
			getRole := mock_service.NewMockAuthUser(c)
			testCase.mockBehaviorParseToken(getRole, testCase.inputToken)
			testCase.mockBehaviorCheck(getRole, testCase.inputPerms, testCase.inputRole)
			testCase.mockBehavior(getRole, testCase.id)
			logger := logging.GetLogger()
			services := &service.Service{AuthUser: getRole}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/roles/%s", testCase.input), nil)
			req.Header.Set("Authorization", "Bearer testToken")

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_getAllRoles(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAuthUser, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAuthUser, token string)
	type mockBehavior func(s *mock_service.MockAuthUser)

	testTable := []struct {
		name                   string
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehaviorCheck      mockBehaviorCheck
		mockBehavior           mockBehavior
		expectedStatusCode     int
		expectedRequestBody    string
	}{
		{
			name:       "OK",
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser) {
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
			name:       "server error",
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser) {
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
			getRoles := mock_service.NewMockAuthUser(c)
			testCase.mockBehaviorParseToken(getRoles, testCase.inputToken)
			testCase.mockBehaviorCheck(getRoles, testCase.inputPerms, testCase.inputRole)
			testCase.mockBehavior(getRoles)
			logger := logging.GetLogger()
			services := &service.Service{AuthUser: getRoles}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", "/roles/", nil)
			req.Header.Set("Authorization", "Bearer testToken")

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_createRole(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAuthUser, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAuthUser, token string)
	type mockBehavior func(s *mock_service.MockAuthUser, role *model.CreateRole)

	testTable := []struct {
		name                   string
		inputBody              string
		inputRole              *model.CreateRole
		inputPerms             string
		inputAuthRole          string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehaviorCheck      mockBehaviorCheck
		mockBehavior           mockBehavior
		expectedStatusCode     int
		expectedRequestBody    string
	}{
		{
			name:          "OK",
			inputBody:     `{"name":"test"}`,
			inputRole:     &model.CreateRole{Name: "test"},
			inputPerms:    "",
			inputAuthRole: "Superadmin",
			inputToken:    "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser, role *model.CreateRole) {
				s.EXPECT().CreateRole(role.Name).Return(1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:          "server error",
			inputBody:     `{"name":"test"}`,
			inputRole:     &model.CreateRole{Name: "test"},
			inputPerms:    "",
			inputAuthRole: "Superadmin",
			inputToken:    "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser, role *model.CreateRole) {
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
			getRole := mock_service.NewMockAuthUser(c)
			testCase.mockBehaviorParseToken(getRole, testCase.inputToken)
			testCase.mockBehaviorCheck(getRole, testCase.inputPerms, testCase.inputAuthRole)
			testCase.mockBehavior(getRole, testCase.inputRole)
			logger := logging.GetLogger()
			services := &service.Service{AuthUser: getRole}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/roles/", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_bindRoleWithPerms(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAuthUser, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAuthUser, token string)
	type mockBehavior func(s *mock_service.MockAuthUser, role *model.BindRoleWithPermission)

	testTable := []struct {
		name                   string
		inputRoleId            string
		inputBody              string
		input                  *model.BindRoleWithPermission
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehaviorCheck      mockBehaviorCheck
		mockBehavior           mockBehavior
		expectedStatusCode     int
	}{
		{
			name:      "OK",
			inputBody: `{"role_id":1,"permissions_id":[1,2,3]}`,
			input: &model.BindRoleWithPermission{
				RoleId:        1,
				PermissionsId: []int{1, 2, 3},
			},
			inputRoleId: "1",
			inputPerms:  "",
			inputRole:   "Superadmin",
			inputToken:  "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser, input *model.BindRoleWithPermission) {
				s.EXPECT().BindRoleWithPerms(input).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:       "Invalid request",
			inputBody:  `{"role_id":"a","permissions_id":[1,2,3]}`,
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior:       func(s *mock_service.MockAuthUser, input *model.BindRoleWithPermission) {},
			expectedStatusCode: 400,
		},
		{
			name:      "server error",
			inputBody: `{"role_id":1,"permissions_id":[1,2,3]}`,
			input: &model.BindRoleWithPermission{
				RoleId:        1,
				PermissionsId: []int{1, 2, 3},
			},
			inputPerms: "",
			inputRole:  "Superadmin",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAuthUser, token string) {
				s.EXPECT().ParseToken(token).Return(&authProto.UserRole{
					UserId:      1,
					Role:        "Superadmin",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAuthUser, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Superadmin", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAuthUser, input *model.BindRoleWithPermission) {
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
			getRole := mock_service.NewMockAuthUser(c)
			testCase.mockBehaviorParseToken(getRole, testCase.inputToken)
			testCase.mockBehaviorCheck(getRole, testCase.inputPerms, testCase.inputRole)
			testCase.mockBehavior(getRole, testCase.input)
			logger := logging.GetLogger()
			services := &service.Service{AuthUser: getRole}
			handler := NewHandler(services, logger)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", fmt.Sprintf("/roles/%s/perms", testCase.inputRoleId), bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}

}
