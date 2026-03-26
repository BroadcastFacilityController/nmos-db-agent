package schema

import (
	"encoding/json"
	"time"

	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	uuid "github.com/google/uuid"
)

type Device struct {
	ID              uuid.UUID
	ResourceVersion string
	Label           string
	Description     *string // v1.1+
	Tags            []byte  // v1.1+

	Type      string
	Receivers []string
	Senders   []string
	NodeID    uuid.UUID
	Controls  []byte // V1.1+

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDeviceFromV1_0(device is04v1_0.Device) (*Device, error) {
	id, err := uuid.Parse(device.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	nodeID, err := uuid.Parse(device.NodeID)
	if err != nil {
		return nil, err
	}

	n := Device{
		ID:              id,
		ResourceVersion: device.Version,
		Label:           device.Label,
		Description:     nil,
		Tags:            nil,

		Type:      device.Type,
		Receivers: device.Receivers,
		Senders:   device.Senders,
		NodeID:    nodeID,
		Controls:  nil,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}

func NewDeviceFromV1_1(device is04v1_1.Device) (*Device, error) {
	id, err := uuid.Parse(device.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	nodeID, err := uuid.Parse(device.NodeID)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(device.Tags)
	if err != nil {
		return nil, err
	}

	controlsJSON, err := json.Marshal(device.Controls)
	if err != nil {
		return nil, err
	}

	n := Device{
		ID:              id,
		ResourceVersion: device.Version,
		Label:           device.Label,
		Description:     strPtrOrNil(device.Description),
		Tags:            tagsJSON,

		Type:      string(device.Type),
		Receivers: device.Receivers,
		Senders:   device.Senders,
		NodeID:    nodeID,
		Controls:  controlsJSON,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}

func NewDeviceFromV1_2(device is04v1_2.Device) (*Device, error) {
	id, err := uuid.Parse(device.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	nodeID, err := uuid.Parse(device.NodeID)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(device.Tags)
	if err != nil {
		return nil, err
	}

	controlsJSON, err := json.Marshal(device.Controls)
	if err != nil {
		return nil, err
	}

	n := Device{
		ID:              id,
		ResourceVersion: device.Version,
		Label:           device.Label,
		Description:     strPtrOrNil(device.Description),
		Tags:            tagsJSON,

		Type:      string(device.Type),
		Receivers: device.Receivers,
		Senders:   device.Senders,
		NodeID:    nodeID,
		Controls:  controlsJSON,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}

func NewDeviceFromV1_3(device is04v1_3.Device) (*Device, error) {
	id, err := uuid.Parse(device.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	nodeID, err := uuid.Parse(device.NodeID)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(device.Tags)
	if err != nil {
		return nil, err
	}

	controlsJSON, err := json.Marshal(device.Controls)
	if err != nil {
		return nil, err
	}

	n := Device{
		ID:              id,
		ResourceVersion: device.Version,
		Label:           device.Label,
		Description:     strPtrOrNil(device.Description),
		Tags:            tagsJSON,

		Type:      string(device.Type),
		Receivers: device.Receivers,
		Senders:   device.Senders,
		NodeID:    nodeID,
		Controls:  controlsJSON,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}
