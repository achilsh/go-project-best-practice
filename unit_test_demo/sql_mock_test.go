package unit_test_demo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//主要介绍 sql mock testing 的使用。

// 先编写一个使用数据库的操作，比如使用 gorm 调用 mysql的场景：

type User struct {
	Id   int    `gorm:"id"`
	Name string `gorm:"name"`
}

func GetUser(db *gorm.DB, id int) (User, error) {
	var user User
	result := db.Where("id =?", id).First(&user)
	return user, result.Error
}

func TestGetUser(t *testing.T) {
	// 创建一个模拟数据库连接和mock对象
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("open sqlmock handle is fail, err: %v", err)
		return
	}
	defer db.Close()

	// 使用gorm连接到模拟数据库
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("create gorm db fail by mock db, err: %v", err)
		return
	}

	// 准备模拟的查询结果
	mUser := User{
		Id:   100,
		Name: "this is mock name",
	}
	resultMockRows := sqlmock.NewRows([]string{"id", "name"}).AddRow(mUser.Id, mUser.Name)

	// 期望执行的SQL查询并设置返回结果
	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE id =? ORDER BY `users`.`id` LIMIT \\?").WithArgs(100).WillReturnRows(resultMockRows)

	// 执行要测试的函数
	dstData, err := GetUser(gormdb, mUser.Id)
	if err != nil {
		t.Logf("get user data from mock db fail, err: %v", err)
		return
	}
	assert.Equal(t, dstData.Name, "this is mock name")

	// 验证是否所有期望的调用都已执行
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
