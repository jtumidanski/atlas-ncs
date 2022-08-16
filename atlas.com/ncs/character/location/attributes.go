package location

type inputDataContainer struct {
	Data inputDataBody `json:"data"`
}

type inputDataBody struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes inputAttributes `json:"attributes"`
}

type inputAttributes struct {
	Type string `json:"type"`
}

type attributes struct {
	Type     string `json:"type"`
	MapId    uint32 `json:"mapId"`
	PortalId uint32 `json:"portalId"`
}
