package main

import "bytes"

/**
接口嵌套接口
*/
type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

func main() {

}
