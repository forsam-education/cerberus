package database

import (
	"fmt"
	"github.com/forsam-education/kerberos/models"
	"github.com/forsam-education/kerberos/utils"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

func generateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"!@#$%^&*()_+=")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func generateFirstUser() error {
	plainPassword := generateRandomPassword(18)

	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), viper.GetInt(utils.PasswordHashCost))
	if err != nil {
		return err
	}

	user := models.User{
		Username: "Admin",
		Password: string(password),
		Email:    "root@localhost",
	}

	err = user.InsertG(boil.Infer())

	utils.Logger.Info(fmt.Sprintf("First user credentials: Admin - %s", plainPassword), nil)

	return err
}

func isFirstStart() bool {
	count, err := models.Users().CountG()
	if err != nil {
		utils.LogAndForceExit(err)
	}

	return count < 1
}

// HandleFirstStart execute operations when you start the server with an empty database.
func HandleFirstStart() error {
	if !isFirstStart() {
		return nil
	}
	utils.Logger.Info("First start detected, generating default data...", nil)
	utils.Logger.Info("Generating first admin user...", nil)

	return generateFirstUser()
}
