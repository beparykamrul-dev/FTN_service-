package aethercore

// Policy controls how Aether-Core may select and use transport profiles.
type Policy struct {
	AllowProduction bool     `json:"allow_production"`
	AllowResearch   bool     `json:"allow_research"`
	ApprovalRequired bool    `json:"approval_required"`
	Preferred       []string `json:"preferred"`
}

func DefaultPolicy() Policy {
	return Policy{
		AllowProduction:  true,
		AllowResearch:    false,
		ApprovalRequired: true,
		Preferred:        []string{"wireguard", "amneziawg"},
	}
}
