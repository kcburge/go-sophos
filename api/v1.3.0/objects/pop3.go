package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Pop3 is a generated struct representing the Sophos Pop3 Endpoint
// GET /api/nodes/pop3
type Pop3 struct {
	AllowedClients            []interface{} `json:"allowed_clients"`
	AllowedServers            []string      `json:"allowed_servers"`
	CffAsMarker               string        `json:"cff_as_marker"`
	CffAv                     int64         `json:"cff_av"`
	CffAvAction               string        `json:"cff_av_action"`
	CffAvEngines              string        `json:"cff_av_engines"`
	CffFileExtensions         []string      `json:"cff_file_extensions"`
	DirectlyDeleteQuarantined int64         `json:"directly_delete_quarantined"`
	Exceptions                []interface{} `json:"exceptions"`
	KnownServers              []interface{} `json:"known_servers"`
	MaxMessageSize            int64         `json:"max_message_size"`
	Prefetching               struct {
		Interval           int64 `json:"interval"`
		OptimizeStorage    int64 `json:"optimize_storage"`
		Status             int64 `json:"status"`
		StorageMinHoldDays int64 `json:"storage_min_hold_days"`
	} `json:"prefetching"`
	QuarantineUnscannable int64         `json:"quarantine_unscannable"`
	SandboxMaxFilesizeMb  int64         `json:"sandbox_max_filesize_mb"`
	SandboxScanStatus     int64         `json:"sandbox_scan_status"`
	ScanTLS               int64         `json:"scan_tls"`
	SenderBlacklist       []interface{} `json:"sender_blacklist"`
	Spam                  string        `json:"spam"`
	SpamExpressions       []interface{} `json:"spam_expressions"`
	Spamplus              string        `json:"spamplus"`
	Spamstatus            int64         `json:"spamstatus"`
	Status                int64         `json:"status"`
	TLSCert               string        `json:"tls_cert"`
	TransparentSkip       []interface{} `json:"transparent_skip"`
	TransparentSkipAutoPf int64         `json:"transparent_skip_auto_pf"`
	UserCharset           string        `json:"user_charset"`
}

var _ sophos.Endpoint = &Pop3{}

var defsPop3 = map[string]sophos.RestObject{
	"Pop3Account":   &Pop3Account{},
	"Pop3Exception": &Pop3Exception{},
	"Pop3Group":     &Pop3Group{},
	"Pop3Server":    &Pop3Server{},
}

// RestObjects implements the sophos.Node interface and returns a map of Pop3's Objects
func (Pop3) RestObjects() map[string]sophos.RestObject { return defsPop3 }

// GetPath implements sophos.RestGetter
func (*Pop3) GetPath() string { return "/api/nodes/pop3" }

// RefRequired implements sophos.RestGetter
func (*Pop3) RefRequired() (string, bool) { return "", false }

var defPop3 = &sophos.Definition{Description: "pop3", Name: "pop3", Link: "/api/definitions/pop3"}

// Definition returns the /api/definitions struct of Pop3
func (Pop3) Definition() sophos.Definition { return *defPop3 }

// ApiRoutes returns all known Pop3 Paths
func (Pop3) ApiRoutes() []string {
	return []string{
		"/api/objects/pop3/account/",
		"/api/objects/pop3/account/{ref}",
		"/api/objects/pop3/account/{ref}/usedby",
		"/api/objects/pop3/exception/",
		"/api/objects/pop3/exception/{ref}",
		"/api/objects/pop3/exception/{ref}/usedby",
		"/api/objects/pop3/group/",
		"/api/objects/pop3/group/{ref}",
		"/api/objects/pop3/group/{ref}/usedby",
		"/api/objects/pop3/server/",
		"/api/objects/pop3/server/{ref}",
		"/api/objects/pop3/server/{ref}/usedby",
	}
}

// References returns the Pop3's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Pop3) References() []string {
	return []string{
		"REF_Pop3Account",
		"REF_Pop3Exception",
		"REF_Pop3Group",
		"REF_Pop3Server",
	}
}

// Pop3Accounts is an Sophos Endpoint subType and implements sophos.RestObject
type Pop3Accounts []Pop3Account

// Pop3Account represents a UTM POP3 account
type Pop3Account struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	// Server description: REF(pop3/server)
	Server   string `json:"server"`
	Username string `json:"username"`
}

var _ sophos.RestGetter = &Pop3Account{}

// GetPath implements sophos.RestObject and returns the Pop3Accounts GET path
// Returns all available pop3/account objects
func (*Pop3Accounts) GetPath() string { return "/api/objects/pop3/account/" }

// RefRequired implements sophos.RestObject
func (*Pop3Accounts) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the Pop3Accounts GET path
// Returns all available account types
func (p *Pop3Account) GetPath() string {
	return fmt.Sprintf("/api/objects/pop3/account/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *Pop3Account) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the Pop3Account DELETE path
// Creates or updates the complete object account
func (*Pop3Account) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/account/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the Pop3Account PATCH path
// Changes to parts of the object account types
func (*Pop3Account) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/account/%s", ref)
}

// PostPath implements sophos.RestObject and returns the Pop3Account POST path
// Create a new pop3/account object
func (*Pop3Account) PostPath() string {
	return "/api/objects/pop3/account/"
}

// PutPath implements sophos.RestObject and returns the Pop3Account PUT path
// Creates or updates the complete object account
func (*Pop3Account) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/account/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*Pop3Account) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/account/%s/usedby", ref)
}

// Pop3Exceptions is an Sophos Endpoint subType and implements sophos.RestObject
type Pop3Exceptions []Pop3Exception

// Pop3Exception represents a UTM POP3 filter exception
type Pop3Exception struct {
	Locked    string        `json:"_locked"`
	Reference string        `json:"_ref"`
	_type     string        `json:"_type"`
	Client    []interface{} `json:"client"`
	Comment   string        `json:"comment"`
	Name      string        `json:"name"`
	Sender    []interface{} `json:"sender"`
	Skiplist  []interface{} `json:"skiplist"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &Pop3Exception{}

// GetPath implements sophos.RestObject and returns the Pop3Exceptions GET path
// Returns all available pop3/exception objects
func (*Pop3Exceptions) GetPath() string { return "/api/objects/pop3/exception/" }

// RefRequired implements sophos.RestObject
func (*Pop3Exceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the Pop3Exceptions GET path
// Returns all available exception types
func (p *Pop3Exception) GetPath() string {
	return fmt.Sprintf("/api/objects/pop3/exception/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *Pop3Exception) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the Pop3Exception DELETE path
// Creates or updates the complete object exception
func (*Pop3Exception) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the Pop3Exception PATCH path
// Changes to parts of the object exception types
func (*Pop3Exception) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the Pop3Exception POST path
// Create a new pop3/exception object
func (*Pop3Exception) PostPath() string {
	return "/api/objects/pop3/exception/"
}

// PutPath implements sophos.RestObject and returns the Pop3Exception PUT path
// Creates or updates the complete object exception
func (*Pop3Exception) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/exception/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*Pop3Exception) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/exception/%s/usedby", ref)
}

// Pop3Groups is an Sophos Endpoint subType and implements sophos.RestObject
type Pop3Groups []Pop3Group

// Pop3Group represents a UTM group
type Pop3Group struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &Pop3Group{}

// GetPath implements sophos.RestObject and returns the Pop3Groups GET path
// Returns all available pop3/group objects
func (*Pop3Groups) GetPath() string { return "/api/objects/pop3/group/" }

// RefRequired implements sophos.RestObject
func (*Pop3Groups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the Pop3Groups GET path
// Returns all available group types
func (p *Pop3Group) GetPath() string { return fmt.Sprintf("/api/objects/pop3/group/%s", p.Reference) }

// RefRequired implements sophos.RestObject
func (p *Pop3Group) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the Pop3Group DELETE path
// Creates or updates the complete object group
func (*Pop3Group) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the Pop3Group PATCH path
// Changes to parts of the object group types
func (*Pop3Group) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the Pop3Group POST path
// Create a new pop3/group object
func (*Pop3Group) PostPath() string {
	return "/api/objects/pop3/group/"
}

// PutPath implements sophos.RestObject and returns the Pop3Group PUT path
// Creates or updates the complete object group
func (*Pop3Group) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*Pop3Group) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/group/%s/usedby", ref)
}

// Pop3Servers is an Sophos Endpoint subType and implements sophos.RestObject
type Pop3Servers []Pop3Server

// Pop3Server represents a UTM POP3 server
type Pop3Server struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Name      string `json:"name"`
	// TlsCert description: REF(ca/host_key_cert)
	// TlsCert default value is ""
	TlsCert string        `json:"tls_cert"`
	Comment string        `json:"comment"`
	Hosts   []interface{} `json:"hosts"`
}

var _ sophos.RestGetter = &Pop3Server{}

// GetPath implements sophos.RestObject and returns the Pop3Servers GET path
// Returns all available pop3/server objects
func (*Pop3Servers) GetPath() string { return "/api/objects/pop3/server/" }

// RefRequired implements sophos.RestObject
func (*Pop3Servers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the Pop3Servers GET path
// Returns all available server types
func (p *Pop3Server) GetPath() string { return fmt.Sprintf("/api/objects/pop3/server/%s", p.Reference) }

// RefRequired implements sophos.RestObject
func (p *Pop3Server) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the Pop3Server DELETE path
// Creates or updates the complete object server
func (*Pop3Server) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/server/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the Pop3Server PATCH path
// Changes to parts of the object server types
func (*Pop3Server) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/server/%s", ref)
}

// PostPath implements sophos.RestObject and returns the Pop3Server POST path
// Create a new pop3/server object
func (*Pop3Server) PostPath() string {
	return "/api/objects/pop3/server/"
}

// PutPath implements sophos.RestObject and returns the Pop3Server PUT path
// Creates or updates the complete object server
func (*Pop3Server) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/server/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*Pop3Server) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/pop3/server/%s/usedby", ref)
}
