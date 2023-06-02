package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getHeader(ctx *gin.Context, key string) ([]string, error) {
	value := ctx.Request.Header[key]
	if len(value) == 0 {
		return nil, fmt.Errorf("%v header not found", key)
	}
	return value, nil
}

func GetClientIPFromContext(ctx *gin.Context) (string, error) {
	xForwardedFor, err := getHeader(ctx, "X-Forwarded-For")
	if err != nil {
		return "", err
	}
	xForwardedForValue := xForwardedFor[0]
	xForwardedForValueArr := strings.Split(xForwardedForValue, ",")
	if len(xForwardedForValueArr) == 0 {
		return "", errors.New("X-Forwarded-For values not found")
	}
	clientIP := xForwardedForValueArr[0]
	return strings.TrimSpace(clientIP), nil
}

func GetBearerTokenFromContext(ctx *gin.Context) (string, error) {
	authorization, err := getHeader(ctx, "Authorization")
	if err != nil {
		return "", err
	}
	authorizationValue := authorization[0]

	authorizationValueArr := strings.Split(authorizationValue, " ")
	if len(authorizationValueArr) != 2 {
		return "", fmt.Errorf("invalid authorization value")
	}
	if authorizationValueArr[0] != "Bearer" {
		return "", fmt.Errorf("invalid bearer authorization value")
	}
	token := authorizationValueArr[1]
	return token, nil
}

func GetCustomHeaderValueFromContext(ctx *gin.Context, key string) (string, error) {
	valueArr, err := getHeader(ctx, key)
	if err != nil {
		return "", err
	}
	value := valueArr[0]

	return value, nil
}

func GetPredefinedValueFromContext(ctx *gin.Context, key string) (string, error) {
	value, ok := ctx.Get(key)
	if !ok {
		return "", fmt.Errorf("%v not found in predefined values", key)
	}
	valueStr, _ := value.(string)
	return valueStr, nil
}

func ValidateContextBody(ctx *gin.Context, obj interface{}) (bool, error) {
	if err := ctx.BindJSON(obj); err != nil {
		return false, err
	}
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		return false, err
	}
	return true, nil
}

func ValidateContextForm(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		return false, err
	}
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		return false, err
	}
	return true, nil
}

func ValidateContextQuery(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.ShouldBindUri(obj); err != nil {
		return false, err
	}
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		return false, err
	}
	return true, nil
}
