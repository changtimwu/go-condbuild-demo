// +build linkstreet

package main

import (
	"./halif"
	"./linkstreet"
)

func GetHal() halif.SwixHalInterface {
	return xcat.GetImpl()
}
