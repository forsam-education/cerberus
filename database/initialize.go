package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/forsam-education/cerberus/models"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
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
		utils.Logger.StdErrorCritical(err, nil)
		os.Exit(1)
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

// Connect returns a database after initializing the connection and pinging the server.
func Connect() (*sql.DB, error) {
	dsn := buildDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	utils.Logger.Info("Connecting to database...", map[string]interface{}{"DSN": dsn})
	var dbErr error
	for i := 1; i <= 3; i++ {
		dbErr = db.Ping()
		if dbErr != nil {
			utils.Logger.Info(fmt.Sprintf("Attempt #%d failed, will retry in 10 seconds", i), map[string]interface{}{"Error": dbErr})
			if i < 3 {
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}

	if dbErr != nil {
		return nil, errors.New("can't connect to database after 3 attempts")
	}

	return db, nil
}
