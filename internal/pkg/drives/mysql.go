package drives

import (
	"fmt"
	"github.com/pro911/authApis/internal/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const dsnTpl = "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"

func MustInitMySQL() {
	var err error
	dsn := fmt.Sprintf(
		dsnTpl,
		conf.Conf.MySQL.Username,
		conf.Conf.MySQL.Password,
		conf.Conf.MySQL.Host,
		conf.Conf.MySQL.Port,
		conf.Conf.MySQL.DBName,
		conf.Conf.MySQL.Charset,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("fatal error db init: %w", err))
	}

	if conf.Conf.Base.IsDebug {
		db = db.Debug()
	}

	fmt.Println("mysql initialized")
}

func DB() *gorm.DB {
	return db
}
