package compliance

type Info struct {
	Data   []InfoData  `json:"data"`
	Paging interface{} `json:"paging"`
}
