package inter

import "fmt"

type FuncTestHaveParam func(interface{})
type FuncTestNoParam func()

var DefaultFuncTestServer funcTestServer

type funcTestServer interface {
	TestFunc(i interface{})
}

func (ft FuncTestHaveParam) TestFunc(i interface{}) {
	fmt.Println("into func test have param...")
	ft(i)
}

func (ft FuncTestNoParam) TestFunc(i interface{}) {
	fmt.Println("into func test no param ...")
	ft()
}

// 直接调用接口DefaultFuncTestServer的方式来调用方法
