package simple

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var(
	db *gorm.DB
	sqlDB * sql.DB
)

func OpenDB(dsn string,config * gorm.Config,maxIdleConns,maxOpenConns int,models ...interface{})(err error){
	if config==nil{
		config =&gorm.Config{}
	}

	if config.NamingStrategy == nil{
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix: "t_",
			SingularTable: true,
		}
	}

	if db,err = gorm.Open(mysql.Open(dsn),config);err!=nil{
		logrus.Errorf("opens database failed:%s ",err.Error())
		return err
	}

	if sqlDB,err=db.DB();err==nil{
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	}else{
		logrus.Error(err)
	}
	if err = db.AutoMigrate(models...); nil != err {
		logrus.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

func DB() *gorm.DB{
	return db
}

func CloseDB(){
	if sqlDB==nil{
		return
	}

	if err:=sqlDB.Close();err!=nil{
		logrus.Errorf("Disconnect from database failed:%s",err.Error())
	}
}