package xcat

import "fmt"

func PortEnable(pno int, enabled bool) {
	s := "Disable"
	if enabled {
		s = "Enabled"
	}
	fmt.Printf("XCAT: %s port %d\n", s, pno)
}
