* 单元测试的种类：
* 1. 使用标准库 testing 的场景
* 2. 使用 convey 三方库来 单元测试
* 3. 使用 assert 来左单元测试。
* 4. 使用 gomonkey 做mock 测试(打桩管理).字节开源mockey也挺好用(https://gitee.com/ByteDance/mockey).  需要添加运行选型: -gcflags=all=-l

* 单元测试覆盖率计算:
* 方式1： go test -coverprofile=coverage.out -gcflags=all=-l ./...
* 方式2: 生成网页结果： go tool cover -html=coverage.out -o coverage.html 
