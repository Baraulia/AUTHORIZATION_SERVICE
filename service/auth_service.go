package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"strings"
	"time"
)

var Secret string

const AccessTokenTTL = time.Minute * 15
const RefreshTokenTTL = time.Hour * 24 * 30

type MyClaims struct {
	UserId int32
	Role   string
	jwt.StandardClaims
}

func (a *AuthUserService) GenerateTokensByAuthUser(user *authProto.User) (*authProto.GeneratedTokens, error) {
	expired := time.Now().Add(AccessTokenTTL)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		UserId:         user.UserId,
		Role:           user.Role,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expired.Unix()},
	})
	accessTokenString, err := accessToken.SignedString([]byte(Secret))
	if err != nil {
		a.logger.Errorf("GenerateTokensByAuthUser:%s", err)
		return nil, fmt.Errorf("GenerateTokensByAuthUser:%w", err)
	}

	refreshExpired := time.Now().Add(RefreshTokenTTL)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		UserId:         user.UserId,
		StandardClaims: jwt.StandardClaims{ExpiresAt: refreshExpired.Unix()},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(Secret))
	if err != nil {
		a.logger.Errorf("GenerateTokensByAuthUser:%s", err)
		return nil, fmt.Errorf("GenerateTokensByAuthUser:%w", err)
	}

	return &authProto.GeneratedTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil

}

func (a *AuthUserService) ParseToken(token string) (*authProto.UserRole, error) {
	claims, err := ParseGWTToken(token)
	if err != nil {
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	role, err := a.repo.RolePerm.GetRoleByName(claims.Role)
	if err != nil {
		return nil, err
	}
	perms, err := a.repo.RolePerm.GetPermsByRoleId(role.ID)
	if err != nil {
		return nil, err
	}
	var slicePerms []string
	for _, perm := range perms {
		slicePerms = append(slicePerms, perm.Name)
	}
	stringPerms := strings.Join(slicePerms[:], ",")
	return &authProto.UserRole{
		UserId:      claims.UserId,
		Role:        role.Name,
		Permissions: stringPerms,
	}, nil
}

func (a *AuthUserService) RefreshTokens(refreshToken string) (*authProto.GeneratedTokens, error) {
	claims, err := ParseGWTToken(refreshToken)
	if err != nil {
		return nil, err
	}
	if claims.StandardClaims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	role, err := a.repo.RolePerm.GetRoleByUserId(int(claims.UserId))
	if err != nil {
		return nil, err
	}
	return a.GenerateTokensByAuthUser(&authProto.User{UserId: claims.UserId, Role: role.Name})
}

func (a *AuthUserService) CheckRole(neededRoles []string, givenRole string) error {
	neededRolesString := strings.Join(neededRoles, ",")
	if !strings.Contains(neededRolesString, givenRole) {
		return fmt.Errorf("not enough rights")
	}
	return nil
}

func (a *AuthUserService) CheckRights(neededPerms []string, givenPerms string) error {
	if neededPerms != nil {
		ok := true
		for _, perm := range neededPerms {
			if !strings.Contains(givenPerms, perm) {
				ok = false
				return fmt.Errorf("not enough rights")
			} else {
				continue
			}
		}
		if ok == true {
			return nil
		}
	}
	return nil
}

func ParseGWTToken(token string) (*MyClaims, error) {
	parseToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("ParseGWTToken:%w", err)
	}
	claims, ok := parseToken.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("error while parsing token")
	}
	return claims, nil
}
