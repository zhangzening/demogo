package inter

func TestStruct1() {
	DefaultStructTest.TestStruct("struct test1...")
}

func TestStruct2() {
	server := initServer()
	server.TestStruct("struct test2...")
}
