package model

import (
	"bytes"
	"encoding/json"
	"io"
)

// Testchain describes testchain data
type Testchain struct {
	ID       string          `json:"id"`
	Title    string          `json:"title"`
	NodeType string          `json:"node_type"`
	Status   string          `json:"status"`
	Deps     []string        `json:"deps"`
	Config   TestchainConfig `json:"config"`
}

// TestchainConfig describes config of Testchain
type TestchainConfig struct {
	Type          string `json:"type"`
	Accounts      uint64 `json:"accounts"`
	BlockMineTime uint64 `json:"block_mine_time"`
	CleanOnStop   bool   `json:"clean_on_stop"`
	SnapshotID    string `json:"snapshot_id"`
	DeployRef     string `json:"deploy_ref"`
	DeployStepID  uint64 `json:"deploy_step_id"`
}

// ToReader returns Testchain data as io.Reader
func (t *Testchain) ToReader() (io.Reader, error) {
	data := map[string]interface{}{
		"testchain": t,
	}
	return mapToBytesBuffer(data)
}

// DataForTakingSnapshotRequest returns map containing env ID and description as io.Reader
func DataForTakingSnapshotRequest(envID string, description string) (io.Reader, error) {
	data := map[string]interface{}{
		"id":          envID,
		"description": description,
	}
	return mapToBytesBuffer(data)
}

// DataForRevertSnapshotRequest returns map containing env ID and snapshot ID as io.Reader
func DataForRevertSnapshotRequest(envID string, snapshotID string) (io.Reader, error) {
	data := map[string]interface{}{
		"id":          envID,
		"snapshot_id": snapshotID,
	}
	return mapToBytesBuffer(data)
}

func mapToBytesBuffer(data map[string]interface{}) (*bytes.Buffer, error) {
	jsonB, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonB), nil
}
