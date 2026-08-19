package controlplane

type Action struct {
    ID string `json:"id"`
    TargetType string `json:"target_type"`
    TargetID string `json:"target_id"`
    Operation string `json:"operation"`
    RequestedBy string `json:"requested_by"`
    RequiresApproval bool `json:"requires_approval"`
}

type ActionResult struct { ActionID string `json:"action_id"`; Status string `json:"status"`; Message string `json:"message,omitempty"` }

// Control actions are represented as intents. Execution belongs to an approved adapter.
func ValidateAction(a Action) bool { return a.ID!="" && a.TargetID!="" && a.Operation!="" && a.RequestedBy!="" }
