package main

import (
	"C"
	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"fmt"
	"unsafe"
)

var dbs map[string]DB

type DB struct {
	p *sqlitex.Pool
}


//export TestGoFunction
func TestGoFunction(a string) *C.char {
	goString := "hello " + a
	return C.CString(goString)
}

func init()  {
	dbs = make(map[string]DB)
}

//export Open
func Open(uri *C.char, flags int) (bool, unsafe.Pointer) {
	u := C.GoString(uri)
	if d, ok := dbs[u]; ok {
		fmt.Println("Found open connection")
		return true, unsafe.Pointer(d.p)
	}
	dbPool, err := sqlitex.Open(u, sqlite.OpenFlags(flags),2)
	if err != nil {
		return false, nil
	}
	ptr := unsafe.Pointer(dbPool)
	dbs[u] = DB{p: dbPool}
	return true, ptr
}

//export Exec
func Exec() {}

//export Query
func Query() {}

//export PrepareStatement
func PrepareStatement() {}

//export Close
func Close() {}

// func main() {}