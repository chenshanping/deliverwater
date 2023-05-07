package initialize

import (
	"deliverwater/global"
	"deliverwater/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
		m.CharSet,
		m.ParseTime,
		m.TimeZone,
	)
	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 彩色打印
			},
		),
		//是否跳过默认事务
		SkipDefaultTransaction: m.Gorm.SkipDefaultTx,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.Gorm.TablePrefix,
			SingularTable: m.Gorm.SingularTable,
		},
		// 执行任何SQL时都会创建一个prepared statement并将其缓存，以提高后续的效率
		PrepareStmt: m.Gorm.PrepareStmt,
		//在AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: m.Gorm.CloseForeignKey,
	}

	mysqlConfig := mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
	})
	db, err := gorm.Open(mysqlConfig, gormConfig)
	if err != nil {
		panic("failed to connect database")
	}
	global.DB = db
	MysqlAuto()
	log.Println("初始化数据库成功")

}

func MysqlAuto() {
	if err := global.DB.AutoMigrate(model.UserInfo{}); err != nil {
		log.Fatalln(fmt.Sprintf("数据库创建失败:%s", err))
	}
}
