package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// IpsecConnection is a generated struct representing the Sophos IpsecConnection Endpoint
// GET /api/nodes/ipsec_connection
type IpsecConnection struct {
	IpsecConnectionAmazonVpc        IpsecConnectionAmazonVpc        `json:"ipsec_connection_amazon_vpc"`
	IpsecConnectionGroup            IpsecConnectionGroup            `json:"ipsec_connection_group"`
	IpsecConnectionL2Tp             IpsecConnectionL2Tp             `json:"ipsec_connection_l2tp"`
	IpsecConnectionRoadwarriorCa    IpsecConnectionRoadwarriorCa    `json:"ipsec_connection_roadwarrior_ca"`
	IpsecConnectionRoadwarriorCisco IpsecConnectionRoadwarriorCisco `json:"ipsec_connection_roadwarrior_cisco"`
	IpsecConnectionRoadwarriorPsk   IpsecConnectionRoadwarriorPsk   `json:"ipsec_connection_roadwarrior_psk"`
	IpsecConnectionRoadwarriorX509  IpsecConnectionRoadwarriorX509  `json:"ipsec_connection_roadwarrior_x509"`
	IpsecConnectionSiteToSite       IpsecConnectionSiteToSite       `json:"ipsec_connection_site_to_site"`
}

var defsIpsecConnection = map[string]sophos.RestObject{
	"IpsecConnectionAmazonVpc":        &IpsecConnectionAmazonVpc{},
	"IpsecConnectionGroup":            &IpsecConnectionGroup{},
	"IpsecConnectionL2Tp":             &IpsecConnectionL2Tp{},
	"IpsecConnectionRoadwarriorCa":    &IpsecConnectionRoadwarriorCa{},
	"IpsecConnectionRoadwarriorCisco": &IpsecConnectionRoadwarriorCisco{},
	"IpsecConnectionRoadwarriorPsk":   &IpsecConnectionRoadwarriorPsk{},
	"IpsecConnectionRoadwarriorX509":  &IpsecConnectionRoadwarriorX509{},
	"IpsecConnectionSiteToSite":       &IpsecConnectionSiteToSite{},
}

// RestObjects implements the sophos.Node interface and returns a map of IpsecConnection's Objects
func (IpsecConnection) RestObjects() map[string]sophos.RestObject { return defsIpsecConnection }

// GetPath implements sophos.RestGetter
func (*IpsecConnection) GetPath() string { return "/api/nodes/ipsec_connection" }

// RefRequired implements sophos.RestGetter
func (*IpsecConnection) RefRequired() (string, bool) { return "", false }

var defIpsecConnection = &sophos.Definition{Description: "ipsec_connection", Name: "ipsec_connection", Link: "/api/definitions/ipsec_connection"}

// Definition returns the /api/definitions struct of IpsecConnection
func (IpsecConnection) Definition() sophos.Definition { return *defIpsecConnection }

// ApiRoutes returns all known IpsecConnection Paths
func (IpsecConnection) ApiRoutes() []string {
	return []string{
		"/api/objects/ipsec_connection/amazon_vpc/",
		"/api/objects/ipsec_connection/amazon_vpc/{ref}",
		"/api/objects/ipsec_connection/amazon_vpc/{ref}/usedby",
		"/api/objects/ipsec_connection/group/",
		"/api/objects/ipsec_connection/group/{ref}",
		"/api/objects/ipsec_connection/group/{ref}/usedby",
		"/api/objects/ipsec_connection/l2tp/",
		"/api/objects/ipsec_connection/l2tp/{ref}",
		"/api/objects/ipsec_connection/l2tp/{ref}/usedby",
		"/api/objects/ipsec_connection/roadwarrior_ca/",
		"/api/objects/ipsec_connection/roadwarrior_ca/{ref}",
		"/api/objects/ipsec_connection/roadwarrior_ca/{ref}/usedby",
		"/api/objects/ipsec_connection/roadwarrior_cisco/",
		"/api/objects/ipsec_connection/roadwarrior_cisco/{ref}",
		"/api/objects/ipsec_connection/roadwarrior_cisco/{ref}/usedby",
		"/api/objects/ipsec_connection/roadwarrior_psk/",
		"/api/objects/ipsec_connection/roadwarrior_psk/{ref}",
		"/api/objects/ipsec_connection/roadwarrior_psk/{ref}/usedby",
		"/api/objects/ipsec_connection/roadwarrior_x509/",
		"/api/objects/ipsec_connection/roadwarrior_x509/{ref}",
		"/api/objects/ipsec_connection/roadwarrior_x509/{ref}/usedby",
		"/api/objects/ipsec_connection/site_to_site/",
		"/api/objects/ipsec_connection/site_to_site/{ref}",
		"/api/objects/ipsec_connection/site_to_site/{ref}/usedby",
	}
}

// References returns the IpsecConnection's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (IpsecConnection) References() []string {
	return []string{
		"REF_IpsecConnectionAmazonVpc",
		"REF_IpsecConnectionGroup",
		"REF_IpsecConnectionL2Tp",
		"REF_IpsecConnectionRoadwarriorCa",
		"REF_IpsecConnectionRoadwarriorCisco",
		"REF_IpsecConnectionRoadwarriorPsk",
		"REF_IpsecConnectionRoadwarriorX509",
		"REF_IpsecConnectionSiteToSite",
	}
}

// IpsecConnectionAmazonVpcs is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionAmazonVpcs []IpsecConnectionAmazonVpc

// IpsecConnectionAmazonVpc is a generated Sophos object
type IpsecConnectionAmazonVpc struct {
	Locked         string `json:"_locked"`
	Reference      string `json:"_ref"`
	_type          string `json:"_type"`
	Authentication string `json:"authentication"`
	Comment        string `json:"comment"`
	Interface      string `json:"interface"`
	Name           string `json:"name"`
	Policy         string `json:"policy"`
	Remote         string `json:"remote"`
}

// GetPath implements sophos.RestObject and returns the IpsecConnectionAmazonVpcs GET path
// Returns all available ipsec_connection/amazon_vpc objects
func (*IpsecConnectionAmazonVpcs) GetPath() string { return "/api/objects/ipsec_connection/amazon_vpc/" }

// RefRequired implements sophos.RestObject
func (*IpsecConnectionAmazonVpcs) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecConnectionAmazonVpcs GET path
// Returns all available amazon_vpc types
func (i *IpsecConnectionAmazonVpc) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_connection/amazon_vpc/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecConnectionAmazonVpc) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionAmazonVpc DELETE path
// Creates or updates the complete object amazon_vpc
func (*IpsecConnectionAmazonVpc) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/amazon_vpc/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionAmazonVpc PATCH path
// Changes to parts of the object amazon_vpc types
func (*IpsecConnectionAmazonVpc) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/amazon_vpc/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionAmazonVpc POST path
// Create a new ipsec_connection/amazon_vpc object
func (*IpsecConnectionAmazonVpc) PostPath() string {
	return "/api/objects/ipsec_connection/amazon_vpc/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionAmazonVpc PUT path
// Creates or updates the complete object amazon_vpc
func (*IpsecConnectionAmazonVpc) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/amazon_vpc/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionAmazonVpc) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/amazon_vpc/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecConnectionAmazonVpc) GetType() string { return i._type }

// IpsecConnectionGroup is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionGroup []interface{}

// GetPath implements sophos.RestObject and returns the IpsecConnectionGroup GET path
// Returns all available ipsec_connection/group objects
func (*IpsecConnectionGroup) GetPath() string { return "/api/objects/ipsec_connection/group/" }

// RefRequired implements sophos.RestObject
func (*IpsecConnectionGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionGroup DELETE path
// Creates or updates the complete object group
func (*IpsecConnectionGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionGroup PATCH path
// Changes to parts of the object group types
func (*IpsecConnectionGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionGroup POST path
// Create a new ipsec_connection/group object
func (*IpsecConnectionGroup) PostPath() string {
	return "/api/objects/ipsec_connection/group/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionGroup PUT path
// Creates or updates the complete object group
func (*IpsecConnectionGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/group/%s/usedby", ref)
}

// IpsecConnectionL2Tps is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionL2Tps []IpsecConnectionL2Tp

// IpsecConnectionL2Tp is a generated Sophos object
type IpsecConnectionL2Tp struct {
	Locked                    string   `json:"_locked"`
	Reference                 string   `json:"_ref"`
	_type                     string   `json:"_type"`
	AuthenticationType        string   `json:"authentication_type"`
	Certificate               string   `json:"certificate"`
	Comment                   string   `json:"comment"`
	Debug                     bool     `json:"debug"`
	Interface                 string   `json:"interface"`
	IPAssignmentDhcp          string   `json:"ip_assignment_dhcp"`
	IPAssignmentDhcpInterface string   `json:"ip_assignment_dhcp_interface"`
	IPAssignmentMode          string   `json:"ip_assignment_mode"`
	IPAssignmentPool          string   `json:"ip_assignment_pool"`
	IphoneConnectionName      string   `json:"iphone_connection_name"`
	IphoneHostname            string   `json:"iphone_hostname"`
	IphoneStatus              bool     `json:"iphone_status"`
	Name                      string   `json:"name"`
	Policy                    string   `json:"policy"`
	Psk                       string   `json:"psk"`
	Status                    bool     `json:"status"`
	Users                     []string `json:"users"`
}

// GetPath implements sophos.RestObject and returns the IpsecConnectionL2Tps GET path
// Returns all available ipsec_connection/l2tp objects
func (*IpsecConnectionL2Tps) GetPath() string { return "/api/objects/ipsec_connection/l2tp/" }

// RefRequired implements sophos.RestObject
func (*IpsecConnectionL2Tps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecConnectionL2Tps GET path
// Returns all available l2tp types
func (i *IpsecConnectionL2Tp) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_connection/l2tp/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecConnectionL2Tp) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionL2Tp DELETE path
// Creates or updates the complete object l2tp
func (*IpsecConnectionL2Tp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/l2tp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionL2Tp PATCH path
// Changes to parts of the object l2tp types
func (*IpsecConnectionL2Tp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/l2tp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionL2Tp POST path
// Create a new ipsec_connection/l2tp object
func (*IpsecConnectionL2Tp) PostPath() string {
	return "/api/objects/ipsec_connection/l2tp/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionL2Tp PUT path
// Creates or updates the complete object l2tp
func (*IpsecConnectionL2Tp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/l2tp/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionL2Tp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/l2tp/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecConnectionL2Tp) GetType() string { return i._type }

// IpsecConnectionRoadwarriorCa is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionRoadwarriorCa []interface{}

// GetPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCa GET path
// Returns all available ipsec_connection/roadwarrior_ca objects
func (*IpsecConnectionRoadwarriorCa) GetPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_ca/"
}

// RefRequired implements sophos.RestObject
func (*IpsecConnectionRoadwarriorCa) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCa DELETE path
// Creates or updates the complete object roadwarrior_ca
func (*IpsecConnectionRoadwarriorCa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_ca/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCa PATCH path
// Changes to parts of the object roadwarrior_ca types
func (*IpsecConnectionRoadwarriorCa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_ca/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCa POST path
// Create a new ipsec_connection/roadwarrior_ca object
func (*IpsecConnectionRoadwarriorCa) PostPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_ca/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCa PUT path
// Creates or updates the complete object roadwarrior_ca
func (*IpsecConnectionRoadwarriorCa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_ca/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionRoadwarriorCa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_ca/%s/usedby", ref)
}

// IpsecConnectionRoadwarriorCisco is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionRoadwarriorCisco []interface{}

// GetPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCisco GET path
// Returns all available ipsec_connection/roadwarrior_cisco objects
func (*IpsecConnectionRoadwarriorCisco) GetPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_cisco/"
}

// RefRequired implements sophos.RestObject
func (*IpsecConnectionRoadwarriorCisco) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCisco DELETE path
// Creates or updates the complete object roadwarrior_cisco
func (*IpsecConnectionRoadwarriorCisco) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_cisco/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCisco PATCH path
// Changes to parts of the object roadwarrior_cisco types
func (*IpsecConnectionRoadwarriorCisco) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_cisco/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCisco POST path
// Create a new ipsec_connection/roadwarrior_cisco object
func (*IpsecConnectionRoadwarriorCisco) PostPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_cisco/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorCisco PUT path
// Creates or updates the complete object roadwarrior_cisco
func (*IpsecConnectionRoadwarriorCisco) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_cisco/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionRoadwarriorCisco) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_cisco/%s/usedby", ref)
}

// IpsecConnectionRoadwarriorPsk is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionRoadwarriorPsk []interface{}

// GetPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorPsk GET path
// Returns all available ipsec_connection/roadwarrior_psk objects
func (*IpsecConnectionRoadwarriorPsk) GetPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_psk/"
}

// RefRequired implements sophos.RestObject
func (*IpsecConnectionRoadwarriorPsk) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorPsk DELETE path
// Creates or updates the complete object roadwarrior_psk
func (*IpsecConnectionRoadwarriorPsk) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_psk/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorPsk PATCH path
// Changes to parts of the object roadwarrior_psk types
func (*IpsecConnectionRoadwarriorPsk) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_psk/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorPsk POST path
// Create a new ipsec_connection/roadwarrior_psk object
func (*IpsecConnectionRoadwarriorPsk) PostPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_psk/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorPsk PUT path
// Creates or updates the complete object roadwarrior_psk
func (*IpsecConnectionRoadwarriorPsk) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_psk/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionRoadwarriorPsk) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_psk/%s/usedby", ref)
}

// IpsecConnectionRoadwarriorX509 is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionRoadwarriorX509 []interface{}

// GetPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorX509 GET path
// Returns all available ipsec_connection/roadwarrior_x509 objects
func (*IpsecConnectionRoadwarriorX509) GetPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_x509/"
}

// RefRequired implements sophos.RestObject
func (*IpsecConnectionRoadwarriorX509) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorX509 DELETE path
// Creates or updates the complete object roadwarrior_x509
func (*IpsecConnectionRoadwarriorX509) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_x509/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorX509 PATCH path
// Changes to parts of the object roadwarrior_x509 types
func (*IpsecConnectionRoadwarriorX509) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_x509/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorX509 POST path
// Create a new ipsec_connection/roadwarrior_x509 object
func (*IpsecConnectionRoadwarriorX509) PostPath() string {
	return "/api/objects/ipsec_connection/roadwarrior_x509/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionRoadwarriorX509 PUT path
// Creates or updates the complete object roadwarrior_x509
func (*IpsecConnectionRoadwarriorX509) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_x509/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionRoadwarriorX509) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/roadwarrior_x509/%s/usedby", ref)
}

// IpsecConnectionSiteToSites is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecConnectionSiteToSites []IpsecConnectionSiteToSite

// IpsecConnectionSiteToSite is a generated Sophos object
type IpsecConnectionSiteToSite struct {
	Locked        string   `json:"_locked"`
	Reference     string   `json:"_ref"`
	_type         string   `json:"_type"`
	AutoPfIn      string   `json:"auto_pf_in"`
	AutoPfOut     string   `json:"auto_pf_out"`
	AutoPfrule    bool     `json:"auto_pfrule"`
	Bind          bool     `json:"bind"`
	Comment       string   `json:"comment"`
	Interface     string   `json:"interface"`
	Name          string   `json:"name"`
	Networks      []string `json:"networks"`
	Policy        string   `json:"policy"`
	RemoteGateway string   `json:"remote_gateway"`
	Status        bool     `json:"status"`
	StrictRouting bool     `json:"strict_routing"`
}

// GetPath implements sophos.RestObject and returns the IpsecConnectionSiteToSites GET path
// Returns all available ipsec_connection/site_to_site objects
func (*IpsecConnectionSiteToSites) GetPath() string {
	return "/api/objects/ipsec_connection/site_to_site/"
}

// RefRequired implements sophos.RestObject
func (*IpsecConnectionSiteToSites) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecConnectionSiteToSites GET path
// Returns all available site_to_site types
func (i *IpsecConnectionSiteToSite) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_connection/site_to_site/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecConnectionSiteToSite) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecConnectionSiteToSite DELETE path
// Creates or updates the complete object site_to_site
func (*IpsecConnectionSiteToSite) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/site_to_site/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecConnectionSiteToSite PATCH path
// Changes to parts of the object site_to_site types
func (*IpsecConnectionSiteToSite) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/site_to_site/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecConnectionSiteToSite POST path
// Create a new ipsec_connection/site_to_site object
func (*IpsecConnectionSiteToSite) PostPath() string {
	return "/api/objects/ipsec_connection/site_to_site/"
}

// PutPath implements sophos.RestObject and returns the IpsecConnectionSiteToSite PUT path
// Creates or updates the complete object site_to_site
func (*IpsecConnectionSiteToSite) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/site_to_site/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecConnectionSiteToSite) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_connection/site_to_site/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecConnectionSiteToSite) GetType() string { return i._type }
