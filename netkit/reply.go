package netkit

type reply struct {
	rAddr  string
	lAddr  string
	cnt    int
	banner string
	err    error
}
