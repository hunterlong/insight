package insight

type Insight struct {
	Endpoint string
	Timeout  int
}

func New(endpoint string) *Insight {
	return &Insight{Endpoint: endpoint}
}

func (a *Insight) NewAddress(address string) *Address {
	return &Address{insight: a, Address: address}
}
