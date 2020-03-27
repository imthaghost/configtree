package models

// Visitor ...
type Visitor struct {
	IP     string
	Remote string
}

// CreateVisitor ...
func CreateVisitor(ip string, remote string) Visitor {
	visitor := new(Visitor)
	visitor.IP = ip
	visitor.Remote = remote
	return *visitor
}
