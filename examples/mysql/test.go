package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	Id       uint32     `gorm:"primarykey"`
	Nickname *string    `gorm:"column:nickname"`
	Profile  *string    `gorm:"column:profile"`
	Email    *string    `gorm:"column:email"`
	Birthday *time.Time `gorm:"column:birthday"`
}

func (u *User) String() string {
	return fmt.Sprintf("User[Id=%d,Nickname=%s,Profile=%s,Email=%s,Birthday=%v]",
		u.Id, *u.Nickname, *u.Profile, *u.Email, u.Birthday)
}

func (u *User) TableName() string {
	return "t_user"
}
func ConnectMySQL(host, port, user, pass, dbname string, config *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return gorm.Open(mysql.Open(dsn), config)
}
func SetConnect(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大打开连接数
	sqlDB.SetMaxIdleConns(100)                 // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置空闲连接最大存活时间
	return nil
}
func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := ConnectMySQL("127.0.0.1", "3306",
		"root", "rootroot", "scootor_mt",
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		fmt.Println("gorm connect mysql err")
	}
	err = SetConnect(db)
	if err != nil {
		fmt.Println("gorm SetConnect  err")
	}
	var users []User
	db.Find(&users)
	for _, i2 := range users {
		fmt.Println(i2.String())
	}
}
