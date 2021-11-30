package main

// import (
// 	"net/http"
// )

// Volume - volume properties
type Volume struct {
	ID                       string               `json:"id,omitempty"`
	Name                     string               `json:"name,omitempty"`
	Description              string               `json:"description,omitempty"`
	Type                     string               `json:"type,omitempty"`
	WWN                      string               `json:"wwn,omitempty"`
	NSID                     int                  `json:"nsid,omitempty"`
	NGUID                    string               `json:"nguid,omitempty"`
	ApplianceID              string               `json:"appliance_id,omitempty"`
	State                    string               `json:"state,omitempty"`
	Size                     int                  `json:"size,omitempty"`
	NodeAffinity             string               `json:"node_affinity,omitempty"`
	CreationTimeStamp        string               `json:"creation_timestamp,omitempty"`
	ProtectionPolicyID       string               `json:"protection_policy_id,omitempty"`
	PerformancePolicyID      string               `json:"performance_policy_id,omitempty"`
	IsReplicationDestination bool                 `json:"is_replication_destination,omitempty"`
	MigrationSessionID       string               `json:"migration_session_id,omitempty"`
	ProtectionData           ProtectionDataStruct `json:"protection_data,omitempty"`
	LocationHistory          string               `json:"location_history,omitempty"`
	TypeL10N                 string               `json:"type_l10n,omitempty"`
	StateL10N                string               `json:"state_l10n,omitempty"`
	NodeAffinityL10N         string               `json:"node_affinity_l10n,omitempty"`
}
