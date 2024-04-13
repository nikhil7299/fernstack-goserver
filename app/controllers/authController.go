package controllers

import (
	"log"
	"server/app/models"
	mongoDB "server/platform/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserDTO struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Type            string `json:"type"`
	IsPhoneVerified bool   `json:"isPhoneVerified"`
}

func SignUpEmail(c *fiber.Ctx) error {
	newUser := new(CreateUserDTO)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"msg": err.Error()})
	}

	// userColl := mongoDB.GetCollection("users")
	userColl := mongoDB.MongoDatabase.Collection("users")
	emailFilter := bson.M{"email": newUser.Email}

	var existingUser models.User
	err := userColl.FindOne(c.Context(), emailFilter).Decode(&existingUser)

	log.Println(existingUser)

	if err != nil && err == mongo.ErrNoDocuments {
		log.Print("User Not found")
		// return nil
	} else {
		if existingUser.Password == "" {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{"msg": "Account created with Google, LogIn with Google"})
		}
		return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{"msg": "User already registered"})

	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)
	newUser.Type = "user"
	newUser.IsPhoneVerified = false
	newUser.Password = string(hashedPassword)
	userId, insertErr := userColl.InsertOne(c.Context(), newUser)
	if insertErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"msg": insertErr.Error()})
	}
	log.Println(userId.InsertedID)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{"msg": "User Registered Successfully"})
}

func LogInEmail(c *fiber.Ctx) error {
	return nil
}
func LogInGoogle(c *fiber.Ctx) error {
	return nil
}

func GetUser(c *fiber.Ctx) error {
	return nil
}
