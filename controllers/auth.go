package controllers

import (
	"golang-crud/db"
	"golang-crud/helper"
	"golang-crud/models"
	"golang-crud/redislib"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"user_name"`
	jwt.StandardClaims
}

var (
	DB   *gorm.DB = db.InitDb()
	data map[string]string
)

func Login(c *gin.Context) {
	user := models.User{}
	credentials := Credentials{}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		helper.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	if err := DB.Where("user_name = ?", credentials.Username).First(&user).Error; err != nil {
		helper.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	passwordFormated := []byte(credentials.Password)
	passwordHashedFormated := []byte(user.Password)

	if err := bcrypt.CompareHashAndPassword(passwordHashedFormated, passwordFormated); err != nil {
		helper.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(jwtKey))

	if err != nil {
		helper.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	var redisAuthKey string = "user_authen_" + strconv.Itoa(int(user.ID))
	redislib.SetKey(redisAuthKey, token)

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
