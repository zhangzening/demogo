package util

import (
	"bufio"
	"fmt"
	"os"
)

type readFromFile struct{}

type readFromStdin struct{}

type writeToFile struct{}

type writeToStdout struct{}

var (
	ReadFromFile  readFromFile
	ReadFromStdin readFromStdin
	WriteToStdout writeToStdout
	WriteToFile   writeToFile
	DefaultReader Reader
	DefaultWriter Writer
)

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

func (r *readFromFile) Read() {
	file, e := os.Open("D://test.txt")
	if e != nil {
		fmt.Println("open file fail : ", e)
		return
	}
	// 文件关闭
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		context, i := reader.ReadString('\n')
		if i != nil {
			fmt.Println("read file fail : ", i)
			return
		}
		if context != "" {
			fmt.Println(" read from file : ", context)
		} else {
			break
		}
	}
}

func (r *readFromStdin) Read() {
	reader := bufio.NewReader(os.Stdin)
	readString, e := reader.ReadString('\n')
	if e != nil {
		fmt.Println("read from stdin fail : ", e)
		return
	}

	fmt.Println("read context : ", readString)
}

func (w *writeToStdout) Write() {

	fmt.Fprintln(os.Stdout, "this is a test of write rizi to stdout")

}

func (w *writeToFile) Write() {
	file, err := os.OpenFile("D://test.txt", os.O_CREATE|os.O_WRONLY, 664)
	if err != nil {
		fmt.Println("write to file fail : ", err)
	}

	fmt.Fprintln(file, "this is a test of write rizi to file")
}
