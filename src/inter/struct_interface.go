package inter

import "fmt"

var (
	DefaultStructTest       structTest
	DefaultStructTestServer structTestServer
)

type structTestServer interface {
	TestStruct(i interface{})
}

type structTest struct{}

func (st *structTest) TestStruct(i interface{}) {
	fmt.Println("-- test struct interface --  ", i)
}

// 测试1：调用DefaultStructTest，直接调用TestStruct(i interface{})

// 测试2：调用DefaultStructTestServer，来调用TestStruct(i interface{})
func initServer() structTestServer {
	//DefaultStructTestServer = &DefaultStructTest
	DefaultStructTestServer = &structTest{}
	return DefaultStructTestServer
}
