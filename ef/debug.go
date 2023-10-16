package ef

import "fmt"

var _debug bool
var _printer func(args ...any)

func EnableDebug() {
	_debug = true
}

func SetPrinter(printer func(args ...any)) {
	_printer = printer
}

func DebugPrint(args ...any) {
	if !_debug {
		return
	}

	if _printer != nil {
		_printer(args...)
	} else {
		fmt.Println(args...)
	}
}
