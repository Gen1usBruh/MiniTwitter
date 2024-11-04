package rest

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/rand"

	"github.com/Gen1usBruh/MiniTwitter/internal/logger/sl"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/models"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
	"github.com/Gen1usBruh/MiniTwitter/util/hash"
	"github.com/Gen1usBruh/MiniTwitter/util/validator"
)

// signUp is the handler for user sign up.
// @Summary User sign up
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param query_params query postgresdb.CreateUserParams true "Query parameters"
// @Success 200 {string} postgresdb.User "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /users/signup [post]
func (h *HandlerConfig) signUp(c *gin.Context) {
	const op = "internal.rest.auth.signUp"
	var rq postgresdb.CreateUserParams
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error binding: ", op), sl.Err(err))
		c.JSON(400, models.BaseResponse{Error: "Wrong data", ErrorCode: 400})
		return
	}
	if !validator.ValidateUserSignUp(rq) {
		c.JSON(400, models.BaseResponse{Error: "Wrong data", ErrorCode: 400})
		return
	}
	rq.Password, err = hash.HashPassword(rq.Password)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error Hash: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to Hash Password", ErrorCode: 500})
		return
	}
	new_user, err := h.Dep.Db.CreateUser(c.Request.Context(), rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error Sign up Database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create User", ErrorCode: 500})
		return
	}
	c.JSON(200, gin.H{"user": new_user})
}

// signIn is the handler for user login.
// @Summary User sign in
// @Description login an existing user
// @Tags Auth
// @Accept json
// @Produce json
// @Param query_params query models.SignInStruct true "Query parameters"
// @Success 200 {string} string "OK"
// @Failure 400 {object} models.BaseResponse "Bad request"
// @Failure 401 {object} models.BaseResponse "Unauthorized access"
// @Failure 500 {object} models.BaseResponse "Internal server error"
// @Router /users/login [post]
func (h *HandlerConfig) signIn(c *gin.Context) {
	var rq models.SignInStruct
	const op = "internal.rest.auth.signIn"
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error binding: ", op), sl.Err(err))
		c.JSON(400, models.BaseResponse{Error: "Bad Request", ErrorCode: 400})
		return
	}
	if !validator.ValidateUserSignIn(rq) {
		c.JSON(400, models.BaseResponse{Error: "Wrong data", ErrorCode: 400})
		return
	}
	user, err := h.Dep.Db.SelectUserSignIn(c.Request.Context(), rq.Email)
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error Sign up Database: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: "Failed to create User", ErrorCode: 500})
		return
	}
	if !hash.CheckPasswordHash(rq.Password, user.Password) {
		c.JSON(401, models.BaseResponse{Error: "Incorrect password", ErrorCode: 401})
		return
	}
	accessToken, refreshToken, err := h.GenerateTokens(strconv.Itoa(int(user.ID)))
	if err != nil {
		h.Dep.Sl.Error(fmt.Sprintf("%s | Error token in db: ", op), sl.Err(err))
		c.JSON(500, models.BaseResponse{Error: err.Error()})
		return
	}
	c.Writer.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.Writer.Header().Add("Content-Type", "application/json")
	c.JSON(200, accessToken)
}

func (h *HandlerConfig) GenerateTokens(userId string) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId,
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 15)},
	})
	accessToken, err := t.SignedString([]byte(h.Dep.Secret))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := NewRefreshToken()
	if err != nil {
		return "", "", err
	}
	arg := postgresdb.CreateTokenParams{}
	id_num, err := strconv.Atoi(userId)
    if err != nil {
        return "", "", err
    }
	arg.UserID = int32(id_num)
	arg.Token = refreshToken
	arg.ExpiresAt.Time = time.Now().Add(time.Hour * 24 * 30)
	arg.ExpiresAt.Valid = true
	if _, err := h.Dep.Db.CreateToken(context.TODO(), arg); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(uint64(time.Now().Unix()))
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (h *HandlerConfig) ParseToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(h.Dep.Secret), nil
	})
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errors.New("invalid tjwt oken")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid jwt claims")
	}
	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid jwt subject")
	}

	return subject, nil
}
