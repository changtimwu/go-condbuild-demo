package xcat

import "fmt"
import "../halif"

type halimpl struct {
}

func (xh *halimpl) PortEnable(pno int, enabled bool) {
	s := "Disable"
	if enabled {
		s = "Enabled"
	}
	fmt.Printf("XCAT: %s port %d\n", s, pno)
}

func (xh *halimpl) GetBoardInfo() halif.BoardInfo {
	return halif.BoardInfo{Description: "12 GE"}
}

func GetImpl() halif.SwixHalInterface {
	return new(halimpl)
}
