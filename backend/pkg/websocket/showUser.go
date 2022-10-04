package websocket

import (
	"fmt"

	"github.com/doxanocap/golang-react/backend/pkg/controllers"
	"github.com/doxanocap/golang-react/backend/pkg/database"
	"github.com/doxanocap/golang-react/backend/pkg/models"
	"github.com/golang-jwt/jwt/v4"
)

func ShowCurrentUser(c *models.Client) models.User {
	cookie, err := c.Ctx.Cookie("jwt")
	if err != nil {
		return models.User{}
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(controllers.SecretKey), nil
	})
	if err != nil {
		panic(err)
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	user := models.User{0, "", "", "", []byte{}}
	res, err := database.DB.Query(fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", claims.Issuer))
	if err != nil {
		panic(err)
	}
	if res != nil {
		for res.Next() {
			err = res.Scan(&user.Id, &user.Token, &user.Username, &user.Email, &user.Password)
			if err != nil {
				panic(err)
			}
			break
		}
	}
	res.Close()
	return user
}
