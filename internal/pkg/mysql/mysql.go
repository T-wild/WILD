package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"wild/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wild/configs"
)

var DB *gorm.DB

func Connect(config *configs.MySQLConfig) *gorm.DB {

	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		panic(`😫: Connected failed, check your Mysql with ` + address)
	}

	// Migrate the schema
	migrateErr := db.AutoMigrate(&models.User{})
	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your Mysql with ` + address)
	}

	// export DB
	//DB = db

	zap.L().Info(`🍟: Successfully connected to Mysql at ` + address)

	return db

}

func InitMysql() error {

	DB = Connect(configs.Conf.MySQLConfig)
	return nil
}
