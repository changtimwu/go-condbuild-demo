// +build xcat

package main

import (
	"./xcat"
)

func PortEnable(pno int, enabled bool) {
	xcat.PortEnable(pno, enabled)
}
