package dao

import (
	"fmt"
	"goCommunity/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func StartDataBase() error {
	var err error
	DB, err = InitDB()
	if err != nil {
		log.Println("[DATABASE STARTUP ERROR]", err)
		return err
	}
	err = CreatedTable()
	if err != nil {
		log.Println("[DATABASE TABLE ERROR]", err)
	}
	return nil
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "root:Gfy20051208@tcp(127.0.0.1:3308)/community?charset=utf8mb4&parseTime=True&loc=Local"

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
		Logger: logger.Default.LogMode(logger.Info), // 打印所有 SQL
	}

	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 设置通用数据库选项
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取通用数据库对象失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接的最大生命周期

	return db, nil
}
func CreatedTable() error {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
