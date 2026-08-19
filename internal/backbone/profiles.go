package backbone

// Profile describes an approved FTN transport profile. It intentionally stores
// policy metadata, not private keys or executable tunnel commands.
type Profile struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Transport Transport `json:"transport"`
    Purpose string `json:"purpose"`
    Enabled bool `json:"enabled"`
    Experimental bool `json:"experimental"`
}

var DefaultProfiles = []Profile{
    {ID: "wg-standard", Name: "WireGuard", Transport: WireGuard, Purpose: "standard secure tunnel", Enabled: true},
    {ID: "wgq-research", Name: "WGQ", Transport: WGQ, Purpose: "experimental quantum/PQC research abstraction", Enabled: false, Experimental: true},
    {ID: "amneziawg", Name: "AmneziaWG", Transport: AmneziaWG, Purpose: "FTN managed transport profile", Enabled: false, Experimental: false},
}
