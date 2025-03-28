package helper

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type queryParams struct {
	Key string
	Val interface{}
}

type TokenInfo struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
		arr  []queryParams
	)

	for k, v := range params {
		arr = append(arr, queryParams{Key: k, Val: v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i].Key) > len(arr[j].Key)
	})

	for _, v := range arr {
		if v.Key != "" && strings.Contains(namedQuery, ":"+v.Key) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+v.Key, "$"+strconv.Itoa(i))
			args = append(args, v.Val)
			i++
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func GenerateJWT(m map[string]interface{}, tokenExpireTime time.Duration, tokenSecretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range m {
		claims[key] = value
	}

	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(tokenExpireTime).Unix()

	return token.SignedString([]byte(tokenSecretKey))
}

func ParseClaims(token string, secretKey string) (TokenInfo, error) {
	claims, err := ExtractClaims(token, secretKey)
	if err != nil {
		return TokenInfo{}, err
	}

	userID, ok := claims["id"].(string)
	if !ok || userID == "" {
		return TokenInfo{}, errors.New("cannot parse 'id' field")
	}

	role, ok := claims["role"].(string)
	if !ok {
		roleFloat, ok := claims["role"].(float64)
		if !ok {
			return TokenInfo{}, errors.New("cannot parse 'role' field")
		}
		role = strconv.Itoa(int(roleFloat))
	}

	return TokenInfo{UserID: userID, Role: role}, nil
}

func ExtractClaims(tokenString string, tokenSecretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tokenSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractToken(bearer string) (string, error) {
	parts := strings.Split(bearer, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid token format, expected 'Bearer <token>'")
	}
	return parts[1], nil
}
