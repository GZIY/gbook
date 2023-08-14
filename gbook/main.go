package main

import (
	"github.com/GZIY/gbook/gbook/internal/repository"
	"github.com/GZIY/gbook/gbook/internal/repository/dao"
	"github.com/GZIY/gbook/gbook/internal/service"
	"github.com/GZIY/gbook/gbook/internal/web"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 初始化
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/gbook"))
	if err != nil {
		// 只在初始化过程 panic
		// panic 相当于整个 goroutine 结束
		panic(err)
	}

	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)

	// 起 gin 服务
	server := gin.Default()
	u.RegisterRoutes(server)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
