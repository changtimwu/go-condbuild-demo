package linkstreet

import "fmt"
import "../halif"

type halimpl struct {
}

func (lh *halimpl) PortEnable(pno int, enabled bool) {
	s := "Disable"
	if enabled {
		s = "Enabled"
	}
	fmt.Printf("LinkStreet: %s port %d\n", s, pno)
}

func (lh *halimpl) GetBoardInfo() halif.BoardInfo {
	return halif.BoardInfo{Description: "10/100 FE + GE"}
}

func GetImpl() halif.SwixHalInterface {
	return new(halimpl)
}
