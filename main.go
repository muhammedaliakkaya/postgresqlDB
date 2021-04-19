package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

const (
	HOST     = "localhost"
	DATABASE = "DB"
	USER     = "postgres"
	PASSWORD = "1"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

//db baglantısını gerceklestirecegimiz alan
func Connect() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	var err error
	db, err = gorm.Open(postgres.Open(vt), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10) //max bağlanılacak kişi sayısı
}

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	Connect()
	db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

}

/*
//2 methodu birden cagirma metodu
func ConnectAndMigrate() {
	Connect()
	Migrate()
}
*/

/*
//db ile models verilerini haberdar edecek method
func Migrate() {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Account{})
	db.AutoMigrate(models.Developer{})
	db.AutoMigrate(models.Score{})
}
*/