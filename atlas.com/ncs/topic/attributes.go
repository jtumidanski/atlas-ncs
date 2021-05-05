package topic

type DataContainer struct {
	Data Data `json:"data"`
}

type Data struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Name string `json:"name"`
}
