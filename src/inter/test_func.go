package inter

import "fmt"

func TestFunc1() {
	DefaultFuncTestServer = FuncTestHaveParam(func(i interface{}) {
		fmt.Println("--- test func have param ---", i)
	})
	DefaultFuncTestServer.TestFunc("have param")
}

func TestFunc2() {
	DefaultFuncTestServer = FuncTestNoParam(func() {
		fmt.Println("--- test func no param ---")
	})
	DefaultFuncTestServer.TestFunc("no param")
}
