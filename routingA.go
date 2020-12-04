package routingA

type RoutingA []interface{}

func (ra RoutingA) Unfold() (rs []Routing, ds []Define) {
	for i := range ra {
		switch r := ra[i].(type) {
		case Routing:
			rs = append(rs, r)
		case Define:
			ds = append(ds, r)
		}
	}
	return
}
