package model

type Node struct {
	Model
	Credential string `json:"credential"`
	PublicIP   string `json:"publicIp"`
	UId        string `json:"uid"`
	Role       string `json:"role"`
}

func NewNode(vm VM) *Node {
	return &Node{
		Model:      Model{Kind: KIND_NODE, Name: vm.Name},
		Credential: vm.Credential,
		PublicIP:   vm.PublicIP,
		UId:        vm.UId,
		Role:       vm.Role,
	}
}
