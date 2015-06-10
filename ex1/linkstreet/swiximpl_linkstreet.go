package linkstreet

import "fmt"

func PortEnable(pno int, enabled bool) {
	s := "Disable"
	if enabled {
		s = "Enabled"
	}
	fmt.Printf("LinkStreet: %s port %d\n", s, pno)
}
