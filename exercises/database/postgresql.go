package database

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ErrorLogger struct {
	message string
	error   error
}
type ContactInfo struct {
	gorm.Model
	Address string
	Email   string
}

type User struct {
	gorm.Model
	Name          string
	ContactInfoId uint
	ContactInfo   ContactInfo // one 2 one
	Companies     []Company   // one 2 many
	Skills        []Skill     `gorm:"many2many:user_skills"`
}

type Company struct {
	gorm.Model
	Name   string
	UserID uint
}

type Skill struct {
	gorm.Model
	Name string
}

func connectToDB() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/sql_db"

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic(err)
	}
	return db
}

func genUuid() uuid.UUID {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(ErrorLogger{
			message: "error in generating uuid",
			error:   err,
		})
	}

	return uid
}

func genRandomStr() string {
	return strconv.Itoa(rand.Int())
}

func Postgresql() {

	db := connectToDB()
	db.AutoMigrate(User{}, ContactInfo{}, Company{})

	//defer db.DB().Close()

	defer func() {
		if val := recover(); val != nil {
			fmt.Println("===== RECOVERED")
			error, ok := val.(ErrorLogger)
			if ok {
				fmt.Println("[recovered error] ", error.message, error.error)
			} else {
				fmt.Println("recovered error: ", error)
			}

		}
	}()

	contactInfo1 := &ContactInfo{Address: "Address_" + genRandomStr(), Email: genRandomStr() + "_@gmail.com"}
	contactInfoCreationErr := db.Create(contactInfo1).Error
	if contactInfoCreationErr != nil {
		panic(
			ErrorLogger{
				message: "error in creating contact info",
				error:   contactInfoCreationErr,
			})
	}

	companies := []Company{
		{Name: "Ford_" + genRandomStr()},
		{Name: "Benz_" + genRandomStr()},
		{Name: "BMW_" + genRandomStr()},
	}

	user1 := &User{Name: "User_" + genRandomStr(), ContactInfo: *contactInfo1, Companies: companies}
	userCreationErr := db.Model(&User{}).Create(user1).Error
	if userCreationErr != nil {
		panic(
			ErrorLogger{
				message: "error in creating user: ",
				error:   userCreationErr,
			})
	}

	skills := []Skill{
		{Name: "s1"},
		{Name: "s2"},
		{Name: "s3"},
	}

	db.Model(&user1).Association("Skills").Append(skills)

	user2 := &User{}
	userQueryError := db.Model(&User{}).Where("ID = ?", 2).First(user2).Error
	if userQueryError != nil {
		panic(
			ErrorLogger{
				message: "error in finding user: ",
				error:   userQueryError,
			})
	}
	db.Model(&user2).Association("Skills").Append(skills)

	fmt.Println(user1.ID)

	tx := db.Begin()
	tx.Delete(&User{}, 7)
	tx.Rollback() // Commit()

}
