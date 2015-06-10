// +build linkstreet

package main

import (
	"./linkstreet"
)

func PortEnable(pno int, enabled bool) {
	linkstreet.PortEnable(pno, enabled)
}
