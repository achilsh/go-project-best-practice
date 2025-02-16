package unit_test_demo

import (
	"context"
	"fmt"
	"testing"
	"time"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/dolthub/vitess/go/vt/proto/query"
	_ "github.com/dolthub/vitess/go/vt/proto/query"
)

type MysqlMockUser struct {
	Email    string
	PhoneNum string
	Name     string
}

var (
	dbName    = "mydb"
	tableName = "mytable"
	address   = "localhost"
	port      = 3306
)

func createTestDatabase() *memory.DbProvider {
	// 创建一个内存 mysql 数据库
	db := memory.NewDatabase(dbName)
	db.BaseDatabase.EnablePrimaryKeyIndexes()

	// 使用一个已创建的db 句柄来初始化一个 db provider.
	pro := memory.NewDBProvider(db)
	session := memory.NewSession(sql.NewBaseSession(), pro)
	ctx := sql.NewContext(context.Background(), sql.WithSession(session))

	// 创建一个内存表
	table := memory.NewTable(db, tableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name", Type: types.Text, Nullable: false, Source: tableName, PrimaryKey: true},
		{Name: "email", Type: types.Text, Nullable: false, Source: tableName, PrimaryKey: true},
		{Name: "phone_numbers", Type: types.JSON, Nullable: false, Source: tableName},
		{Name: "created_at", Type: types.MustCreateDatetimeType(query.Type_DATETIME, 6), Nullable: false, Source: tableName},
	}), db.GetForeignKeyCollection())

	// 将内存表添加到数据库中
	db.AddTable(tableName, table)

	creationTime := time.Unix(0, 1667304000000001000).UTC()

	// 插入数据记录
	_ = table.Insert(ctx, sql.NewRow("Jane Deo", "janedeo@gmail.com", types.MustJSON(`["556-565-566", "777-777-777"]`), creationTime))
	_ = table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", types.MustJSON(`[]`), creationTime))
	_ = table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", types.MustJSON(`["555-555-555"]`), creationTime))
	_ = table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", types.MustJSON(`[]`), creationTime))

	return pro
}

func TestRunMockMysql(t *testing.T) {
	pro := createTestDatabase()
	engine := sqle.NewDefault(pro)

	session := memory.NewSession(sql.NewBaseSession(), pro)
	ctx := sql.NewContext(context.Background(), sql.WithSession(session))
	ctx.SetCurrentDatabase(dbName)

	// This variable may be found in the "users_example.go" file. Please refer to that file for a walkthrough on how to
	// set up the "mysql" database to allow user creation and user checking when establishing connections. This is set
	// to false for this example, but feel free to play around with it and see how it works.

	if enableUsers {
		if err := enableUserAccounts(ctx, engine); err != nil {
			panic(err)
		}
	}

	// 配置 内存 Mysql server的服务的连接信息，指定后，client 可以按配置连接。
	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
	}

	// 创建 内存 mysql server 实例
	s, err := server.NewServer(config, engine, memory.NewSessionBuilder(pro), nil)
	if err != nil {
		panic(err)
	}
	
	// 启动内存mysql 服务
	if err = s.Start(); err != nil {
		panic(err)
	}
}
