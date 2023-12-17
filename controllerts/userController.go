package controllers

import (
	"fmt"
	"context"
	"log"
	"net/http"
	"time"
	"strconv"
	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
	"githubmcom/go-playground/validator/v10"
	helper "github.com/jamyMarkos/gin-auth/helpers"
	"github.com/jamyMarkos/gin-auth/models"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()  {
	
}

func VerifyPassword()  {
	
}

func Signup()  {
	
}

func Login()  {
	
}


func GetUsers()  {
	
}

func GetUser()  {
	
}