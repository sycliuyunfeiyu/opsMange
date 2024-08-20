package consulData

type ConsulRpcModel struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
	//Tags    []string            `json:"tags"`
	//Meta    map[string]string   `json:"Meta"`
	//Checks  map[string][]string `json:"checks"`

	Tags   []string            `json:"tags"`
	Meta   map[string]string   `json:"Meta"`
	Checks []map[string]string `json:"checks"`
	//Checks string `json:"checks"`

}
