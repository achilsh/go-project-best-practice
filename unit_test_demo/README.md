* 单元测试的种类：
* 1. 使用标准库 testing 的场景
* 2. 使用 convey 三方库来 单元测试
* 3. 使用 assert 来左单元测试。
* 4. 使用 gomonkey 做 mock 测试(打桩管理).字节开源mockey也挺好用(https://gitee.com/ByteDance/mockey).  需要添加运行选型: -gcflags=all=-l

* 单元测试覆盖率计算:
* 方式1： go test -coverprofile=coverage.out -gcflags=all=-l ./...
* 方式2: 生成网页结果： go tool cover -html=coverage.out -o coverage.html 

* TestMain 使用介绍，包括使用场景：
  a) 在做测试的时候，可能会在测试之前做些准备工作，例如创建数据库连接等；在测试之后做些清理工作，例如关闭数据库连接、清理测试文件等。
* 使用 TestMain 函数的过程：
  b) 在 _test.go文件中添加 TestMain 函数，其入参为 *testing.M。
  c) TestMain是一个特殊的函数（相当于 main 函数），测试用例在执行时，会先执行 TestMain 函数，在 TestMain 中调用 m.Run() 函数来执行普通的测试函数。在m.Run()函数前面可以编写准备逻辑，在m.Run()后面可以编写清理逻辑。

*
* sqlmock 使和介绍：模拟数据库连接，数据库是项目中比较常见的依赖，在遇到数据库依赖时都可以用它。使用 开源库：  github.com/DATA-DOG/go-sqlmock； demo 在本目录下的 sql_mock_test.go 文件有描述。

* mock mysql 接口操作方式：1） 使用内存数据库 sqlite, 2) 使用上面面的sqlmock  3) 使用内存 mysql 服务： https://github.com/dolthub/go-mysql-server 下面将使用第三种方式来mock mysql的操作。


#

* 通过压测 testing.B 得到: for _, v := range items {} 数组或者slice分片中每个元素时；存在copy数据 分片中每个元素，有性能问题。
* 测试用例编写可采用 table-driven 的设计模式，它将多组测试数据以表格形式组织，与具体的测试逻辑分离开来。
