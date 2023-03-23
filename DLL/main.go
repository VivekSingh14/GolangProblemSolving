package main

import (
	"demo1/DLL/contracts"
	"demo1/DLL/dllnodes"
)

func main() {

	var dlllist contracts.StandardsMethods

	dlllist = dllnodes.NewDLinkedList()
	dlllist.Insert(2)
	dlllist.Insert(1)
	dlllist.Insert(5)
	dlllist.Insert(3)
	dlllist.Insert(9)

	dlllist.Traverse()

}
