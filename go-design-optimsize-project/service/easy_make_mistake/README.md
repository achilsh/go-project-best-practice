* golang 容易犯的错误：
* 1. 在使用时，对于接口类型变量的判空，比如 x := call() 其中 x是一个接口类型变量，
* 如果对x 进行为空判断，比如： if x != nil {} 我们需要格外留意，否则稍有不慎，就有可能出现 “nil!= nil”
* 需要注意call()返回值如果返回为空，不仅要求返回的值的值是Nil,返回值的类型也应该是 nil,不能是其他定义的类型。
* 否者call()返回值不是nil.

* 2. json unmarshal int, 实际是float64类型

* 3. waitgroup使用add(1) 不当，在协程内部使用，而不是在go func(){}外面或前面调用 add(1).

* 4. channel使用不当，导致使用channel的协程阻塞而泄露.
