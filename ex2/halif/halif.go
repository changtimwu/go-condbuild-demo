package halif

type BoardInfo struct {
	Description string
}

type SwixHalInterface interface {
	PortEnable(pno int, enabled bool)
	GetBoardInfo() BoardInfo
}
