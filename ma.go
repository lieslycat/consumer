package main

import (
		"github.com/lieslycat/no_interface"
		"github.com/lieslycat/no_imp"
)

func test(imp no_imp.Inter) {
		imp.Interface_func()
}

func main() {
		s = no_interface.No_Interface {
				Name: "xxx",
		}
		test(s)

}
