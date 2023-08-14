package dao

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// 这里的 User 直接对应 数据库表结构
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 设置为唯一索引
	Email    sql.NullString `gorm:"unique"`
	Password string

	//Phone *string
	Phone sql.NullString `gorm:"unique"`

	// 创建时间
	Ctime int64
	// 更新时间
	Utime int64
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	// 存 毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now

	return dao.db.WithContext(ctx).Create(&u).Error
}
