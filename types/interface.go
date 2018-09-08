package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Interface is a generated struct representing the Sophos Interface Endpoint
// GET /api/nodes/interface
type Interface struct {
	InterfaceBridge   InterfaceBridge   `json:"interface_bridge"`
	InterfaceEthernet InterfaceEthernet `json:"interface_ethernet"`
	InterfaceGroup    InterfaceGroup    `json:"interface_group"`
	InterfacePpp3G    InterfacePpp3G    `json:"interface_ppp3g"`
	InterfacePppmodem InterfacePppmodem `json:"interface_pppmodem"`
	InterfacePppoa    InterfacePppoa    `json:"interface_pppoa"`
	InterfacePppoe    InterfacePppoe    `json:"interface_pppoe"`
	InterfaceTunnel   InterfaceTunnel   `json:"interface_tunnel"`
	InterfaceVlan     InterfaceVlan     `json:"interface_vlan"`
}

var defsInterface = map[string]sophos.RestObject{
	"InterfaceBridge":   &InterfaceBridge{},
	"InterfaceEthernet": &InterfaceEthernet{},
	"InterfaceGroup":    &InterfaceGroup{},
	"InterfacePpp3G":    &InterfacePpp3G{},
	"InterfacePppmodem": &InterfacePppmodem{},
	"InterfacePppoa":    &InterfacePppoa{},
	"InterfacePppoe":    &InterfacePppoe{},
	"InterfaceTunnel":   &InterfaceTunnel{},
	"InterfaceVlan":     &InterfaceVlan{},
}

// RestObjects implements the sophos.Node interface and returns a map of Interface's Objects
func (Interface) RestObjects() map[string]sophos.RestObject { return defsInterface }

// GetPath implements sophos.RestGetter
func (*Interface) GetPath() string { return "/api/nodes/interface" }

// RefRequired implements sophos.RestGetter
func (*Interface) RefRequired() (string, bool) { return "", false }

var defInterface = &sophos.Definition{Description: "interface", Name: "interface", Link: "/api/definitions/interface"}

// Definition returns the /api/definitions struct of Interface
func (Interface) Definition() sophos.Definition { return *defInterface }

// ApiRoutes returns all known Interface Paths
func (Interface) ApiRoutes() []string {
	return []string{
		"/api/objects/interface/bridge/",
		"/api/objects/interface/bridge/{ref}",
		"/api/objects/interface/bridge/{ref}/usedby",
		"/api/objects/interface/ethernet/",
		"/api/objects/interface/ethernet/{ref}",
		"/api/objects/interface/ethernet/{ref}/usedby",
		"/api/objects/interface/group/",
		"/api/objects/interface/group/{ref}",
		"/api/objects/interface/group/{ref}/usedby",
		"/api/objects/interface/ppp3g/",
		"/api/objects/interface/ppp3g/{ref}",
		"/api/objects/interface/ppp3g/{ref}/usedby",
		"/api/objects/interface/pppmodem/",
		"/api/objects/interface/pppmodem/{ref}",
		"/api/objects/interface/pppmodem/{ref}/usedby",
		"/api/objects/interface/pppoa/",
		"/api/objects/interface/pppoa/{ref}",
		"/api/objects/interface/pppoa/{ref}/usedby",
		"/api/objects/interface/pppoe/",
		"/api/objects/interface/pppoe/{ref}",
		"/api/objects/interface/pppoe/{ref}/usedby",
		"/api/objects/interface/tunnel/",
		"/api/objects/interface/tunnel/{ref}",
		"/api/objects/interface/tunnel/{ref}/usedby",
		"/api/objects/interface/vlan/",
		"/api/objects/interface/vlan/{ref}",
		"/api/objects/interface/vlan/{ref}/usedby",
	}
}

// References returns the Interface's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Interface) References() []string {
	return []string{
		"REF_InterfaceBridge",
		"REF_InterfaceEthernet",
		"REF_InterfaceGroup",
		"REF_InterfacePpp3G",
		"REF_InterfacePppmodem",
		"REF_InterfacePppoa",
		"REF_InterfacePppoe",
		"REF_InterfaceTunnel",
		"REF_InterfaceVlan",
	}
}

// InterfaceBridge is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceBridge []interface{}

// GetPath implements sophos.RestObject and returns the InterfaceBridge GET path
// Returns all available interface/bridge objects
func (*InterfaceBridge) GetPath() string { return "/api/objects/interface/bridge/" }

// RefRequired implements sophos.RestObject
func (*InterfaceBridge) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the InterfaceBridge DELETE path
// Creates or updates the complete object bridge
func (*InterfaceBridge) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceBridge PATCH path
// Changes to parts of the object bridge types
func (*InterfaceBridge) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceBridge POST path
// Create a new interface/bridge object
func (*InterfaceBridge) PostPath() string {
	return "/api/objects/interface/bridge/"
}

// PutPath implements sophos.RestObject and returns the InterfaceBridge PUT path
// Creates or updates the complete object bridge
func (*InterfaceBridge) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceBridge) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/bridge/%s/usedby", ref)
}

// InterfaceEthernets is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceEthernets []InterfaceEthernet

// InterfaceEthernet is a generated Sophos object
type InterfaceEthernet struct {
	Locked              string        `json:"_locked"`
	Reference           string        `json:"_ref"`
	_type               string        `json:"_type"`
	AdditionalAddresses []interface{} `json:"additional_addresses"`
	Bandwidth           int64         `json:"bandwidth"`
	Comment             string        `json:"comment"`
	Inbandwidth         int64         `json:"inbandwidth"`
	Itfhw               string        `json:"itfhw"`
	Link                bool          `json:"link"`
	Mtu                 int64         `json:"mtu"`
	MtuAutoDiscovery    bool          `json:"mtu_auto_discovery"`
	Name                string        `json:"name"`
	Outbandwidth        int64         `json:"outbandwidth"`
	PrimaryAddress      string        `json:"primary_address"`
	Proxyarp            bool          `json:"proxyarp"`
	Proxyndp            bool          `json:"proxyndp"`
	Status              bool          `json:"status"`
}

// GetPath implements sophos.RestObject and returns the InterfaceEthernets GET path
// Returns all available interface/ethernet objects
func (*InterfaceEthernets) GetPath() string { return "/api/objects/interface/ethernet/" }

// RefRequired implements sophos.RestObject
func (*InterfaceEthernets) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceEthernets GET path
// Returns all available ethernet types
func (i *InterfaceEthernet) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceEthernet) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceEthernet DELETE path
// Creates or updates the complete object ethernet
func (*InterfaceEthernet) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceEthernet PATCH path
// Changes to parts of the object ethernet types
func (*InterfaceEthernet) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceEthernet POST path
// Create a new interface/ethernet object
func (*InterfaceEthernet) PostPath() string {
	return "/api/objects/interface/ethernet/"
}

// PutPath implements sophos.RestObject and returns the InterfaceEthernet PUT path
// Creates or updates the complete object ethernet
func (*InterfaceEthernet) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceEthernet) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ethernet/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *InterfaceEthernet) GetType() string { return i._type }

// InterfaceGroups is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceGroups []InterfaceGroup

// InterfaceGroup is a generated Sophos object
type InterfaceGroup struct {
	Locked           string        `json:"_locked"`
	Reference        string        `json:"_ref"`
	_type            string        `json:"_type"`
	Comment          string        `json:"comment"`
	Link             bool          `json:"link"`
	Members          []interface{} `json:"members"`
	Name             string        `json:"name"`
	PrimaryAddresses string        `json:"primary_addresses"`
}

// GetPath implements sophos.RestObject and returns the InterfaceGroups GET path
// Returns all available interface/group objects
func (*InterfaceGroups) GetPath() string { return "/api/objects/interface/group/" }

// RefRequired implements sophos.RestObject
func (*InterfaceGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceGroups GET path
// Returns all available group types
func (i *InterfaceGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/group/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceGroup) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceGroup DELETE path
// Creates or updates the complete object group
func (*InterfaceGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceGroup PATCH path
// Changes to parts of the object group types
func (*InterfaceGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceGroup POST path
// Create a new interface/group object
func (*InterfaceGroup) PostPath() string {
	return "/api/objects/interface/group/"
}

// PutPath implements sophos.RestObject and returns the InterfaceGroup PUT path
// Creates or updates the complete object group
func (*InterfaceGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *InterfaceGroup) GetType() string { return i._type }

// InterfacePpp3G is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePpp3G []interface{}

// GetPath implements sophos.RestObject and returns the InterfacePpp3G GET path
// Returns all available interface/ppp3g objects
func (*InterfacePpp3G) GetPath() string { return "/api/objects/interface/ppp3g/" }

// RefRequired implements sophos.RestObject
func (*InterfacePpp3G) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the InterfacePpp3G DELETE path
// Creates or updates the complete object ppp3g
func (*InterfacePpp3G) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePpp3G PATCH path
// Changes to parts of the object ppp3g types
func (*InterfacePpp3G) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePpp3G POST path
// Create a new interface/ppp3g object
func (*InterfacePpp3G) PostPath() string {
	return "/api/objects/interface/ppp3g/"
}

// PutPath implements sophos.RestObject and returns the InterfacePpp3G PUT path
// Creates or updates the complete object ppp3g
func (*InterfacePpp3G) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePpp3G) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/ppp3g/%s/usedby", ref)
}

// InterfacePppmodem is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePppmodem []interface{}

// GetPath implements sophos.RestObject and returns the InterfacePppmodem GET path
// Returns all available interface/pppmodem objects
func (*InterfacePppmodem) GetPath() string { return "/api/objects/interface/pppmodem/" }

// RefRequired implements sophos.RestObject
func (*InterfacePppmodem) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the InterfacePppmodem DELETE path
// Creates or updates the complete object pppmodem
func (*InterfacePppmodem) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePppmodem PATCH path
// Changes to parts of the object pppmodem types
func (*InterfacePppmodem) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePppmodem POST path
// Create a new interface/pppmodem object
func (*InterfacePppmodem) PostPath() string {
	return "/api/objects/interface/pppmodem/"
}

// PutPath implements sophos.RestObject and returns the InterfacePppmodem PUT path
// Creates or updates the complete object pppmodem
func (*InterfacePppmodem) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePppmodem) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppmodem/%s/usedby", ref)
}

// InterfacePppoa is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePppoa []interface{}

// GetPath implements sophos.RestObject and returns the InterfacePppoa GET path
// Returns all available interface/pppoa objects
func (*InterfacePppoa) GetPath() string { return "/api/objects/interface/pppoa/" }

// RefRequired implements sophos.RestObject
func (*InterfacePppoa) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the InterfacePppoa DELETE path
// Creates or updates the complete object pppoa
func (*InterfacePppoa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePppoa PATCH path
// Changes to parts of the object pppoa types
func (*InterfacePppoa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePppoa POST path
// Create a new interface/pppoa object
func (*InterfacePppoa) PostPath() string {
	return "/api/objects/interface/pppoa/"
}

// PutPath implements sophos.RestObject and returns the InterfacePppoa PUT path
// Creates or updates the complete object pppoa
func (*InterfacePppoa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePppoa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoa/%s/usedby", ref)
}

// InterfacePppoe is an Sophos Endpoint subType and implements sophos.RestObject
type InterfacePppoe []interface{}

// GetPath implements sophos.RestObject and returns the InterfacePppoe GET path
// Returns all available interface/pppoe objects
func (*InterfacePppoe) GetPath() string { return "/api/objects/interface/pppoe/" }

// RefRequired implements sophos.RestObject
func (*InterfacePppoe) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the InterfacePppoe DELETE path
// Creates or updates the complete object pppoe
func (*InterfacePppoe) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfacePppoe PATCH path
// Changes to parts of the object pppoe types
func (*InterfacePppoe) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfacePppoe POST path
// Create a new interface/pppoe object
func (*InterfacePppoe) PostPath() string {
	return "/api/objects/interface/pppoe/"
}

// PutPath implements sophos.RestObject and returns the InterfacePppoe PUT path
// Creates or updates the complete object pppoe
func (*InterfacePppoe) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfacePppoe) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/pppoe/%s/usedby", ref)
}

// InterfaceTunnel is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceTunnel []interface{}

// GetPath implements sophos.RestObject and returns the InterfaceTunnel GET path
// Returns all available interface/tunnel objects
func (*InterfaceTunnel) GetPath() string { return "/api/objects/interface/tunnel/" }

// RefRequired implements sophos.RestObject
func (*InterfaceTunnel) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the InterfaceTunnel DELETE path
// Creates or updates the complete object tunnel
func (*InterfaceTunnel) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceTunnel PATCH path
// Changes to parts of the object tunnel types
func (*InterfaceTunnel) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceTunnel POST path
// Create a new interface/tunnel object
func (*InterfaceTunnel) PostPath() string {
	return "/api/objects/interface/tunnel/"
}

// PutPath implements sophos.RestObject and returns the InterfaceTunnel PUT path
// Creates or updates the complete object tunnel
func (*InterfaceTunnel) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceTunnel) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/tunnel/%s/usedby", ref)
}

// InterfaceVlans is an Sophos Endpoint subType and implements sophos.RestObject
type InterfaceVlans []InterfaceVlan

// InterfaceVlan is a generated Sophos object
type InterfaceVlan struct {
	Locked              string        `json:"_locked"`
	Reference           string        `json:"_ref"`
	_type               string        `json:"_type"`
	AdditionalAddresses []interface{} `json:"additional_addresses"`
	Bandwidth           int64         `json:"bandwidth"`
	Comment             string        `json:"comment"`
	Inbandwidth         int64         `json:"inbandwidth"`
	Itfhw               string        `json:"itfhw"`
	Link                bool          `json:"link"`
	Macvlan             bool          `json:"macvlan"`
	Mtu                 int64         `json:"mtu"`
	MtuAutoDiscovery    bool          `json:"mtu_auto_discovery"`
	Name                string        `json:"name"`
	Outbandwidth        int64         `json:"outbandwidth"`
	PrimaryAddress      string        `json:"primary_address"`
	Proxyarp            bool          `json:"proxyarp"`
	Proxyndp            bool          `json:"proxyndp"`
	Status              bool          `json:"status"`
	Vlantag             int64         `json:"vlantag"`
}

// GetPath implements sophos.RestObject and returns the InterfaceVlans GET path
// Returns all available interface/vlan objects
func (*InterfaceVlans) GetPath() string { return "/api/objects/interface/vlan/" }

// RefRequired implements sophos.RestObject
func (*InterfaceVlans) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the InterfaceVlans GET path
// Returns all available vlan types
func (i *InterfaceVlan) GetPath() string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *InterfaceVlan) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the InterfaceVlan DELETE path
// Creates or updates the complete object vlan
func (*InterfaceVlan) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the InterfaceVlan PATCH path
// Changes to parts of the object vlan types
func (*InterfaceVlan) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", ref)
}

// PostPath implements sophos.RestObject and returns the InterfaceVlan POST path
// Create a new interface/vlan object
func (*InterfaceVlan) PostPath() string {
	return "/api/objects/interface/vlan/"
}

// PutPath implements sophos.RestObject and returns the InterfaceVlan PUT path
// Creates or updates the complete object vlan
func (*InterfaceVlan) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*InterfaceVlan) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/interface/vlan/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *InterfaceVlan) GetType() string { return i._type }
