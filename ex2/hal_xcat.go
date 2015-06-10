// +build xcat

package main

import (
	"./halif"
	"./xcat"
)

func GetHal() halif.SwixHalInterface {
	return xcat.GetImpl()
}
