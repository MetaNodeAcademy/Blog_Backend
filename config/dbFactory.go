package config

import (
	"Blog_Backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitTables() {
	db := GetDBConn()
	err := db.AutoMigrate(model.User{}, model.Comment{}, model.Article{})
	if err != nil {
		log.Fatal(err)
	}
	////initData
	//user1 := repo.User{
	//	UserName: "David",
	//	Password: "123456",
	//	Email:    "test@gmail.com",
	//	Articles: []repo.Article{
	//		{Title: "Go 的基础语法学习", Content: "HelloWorld",
	//			CreatedAt: time.Now(), UpdatedAt: time.Now(),
	//			Comments: []repo.Comment{
	//				{Content: "写得一般般", CreatedAt: time.Now()},
	//				{Content: "哈哈哈哈哈", CreatedAt: time.Now()},
	//			}},
	//		{Title: "Go 的进阶", Content: "代理",
	//			CreatedAt: time.Now(), UpdatedAt: time.Now(),
	//			Comments: []repo.Comment{
	//				{Content: "牛批", CreatedAt: time.Now()},
	//				{Content: "6666", CreatedAt: time.Now()},
	//			}},
	//	},
	//}
	//user2 := repo.User{
	//	UserName: "Bruce",
	//	Password: "123456",
	//	Email:    "bruce@gmail.com",
	//	Articles: []repo.Article{
	//		{Title: "Gin 框架的介绍", Content: "2222",
	//			CreatedAt: time.Now(), UpdatedAt: time.Now(),
	//			Comments: []repo.Comment{
	//				{Content: "写得一般般", CreatedAt: time.Now()},
	//				{Content: "哈哈哈哈哈", CreatedAt: time.Now()},
	//			}},
	//	},
	//}
	//gorm.G[repo.User](db).Create(context.Background(), &user1)
	//gorm.G[repo.User](db).Create(context.Background(), &user2)
}

func GetDBConn() *gorm.DB {
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

	db, err := gorm.Open(sqlite.Open("./blog.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
