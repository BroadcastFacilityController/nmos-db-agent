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

type Sender struct {
	// Resource Core
	ID              uuid.UUID
	ResourceVersion string
	Label           string
	Description     *string
	Tags            []byte // JSONB

	// Sender
	Caps                 []byte // JSONB
	FlowID               *uuid.UUID
	Transport            string
	DeviceID             uuid.UUID
	ManifestHRef         *string
	InterfaceBindings    []string
	SubscriptionReceiver *uuid.UUID
	SubscriptionActive   *bool

	TransportFile []byte // Raw transport file from ManifestHRef (usually SDP which can be converted via string)

	// Database Helpers
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSenderFromV1_0(sender *is04v1_0.Sender, transportFile []byte) (*Sender, error) {
	id, err := uuid.Parse(sender.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tagsJSON, err := json.Marshal(sender.Tags)
	if err != nil {
		return nil, err
	}

	flowID, err := uuid.Parse(sender.FlowID)
	if err != nil {
		return nil, err
	}

	deviceID, err := uuid.Parse(sender.DeviceID)
	if err != nil {
		return nil, err
	}

	line := Sender{
		ID:                   id,
		ResourceVersion:      sender.Version,
		Label:                sender.Label,
		Description:          strPtrOrNil(sender.Description),
		Tags:                 tagsJSON,
		Caps:                 nil,
		FlowID:               &flowID,
		Transport:            string(sender.Transport),
		DeviceID:             deviceID,
		ManifestHRef:         strPtrOrNil(string(sender.ManifestHRef)),
		InterfaceBindings:    nil,
		TransportFile:        transportFile,
		SubscriptionReceiver: nil,
		SubscriptionActive:   nil,
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	return &line, nil
}

func NewSenderFromV1_1(sender *is04v1_1.Sender, transportFile []byte) (*Sender, error) {
	id, err := uuid.Parse(sender.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tagsJSON, err := json.Marshal(sender.Tags)
	if err != nil {
		return nil, err
	}

	flowID, err := uuid.Parse(sender.FlowID)
	if err != nil {
		return nil, err
	}

	deviceID, err := uuid.Parse(sender.DeviceID)
	if err != nil {
		return nil, err
	}

	line := Sender{
		ID:                   id,
		ResourceVersion:      sender.Version,
		Label:                sender.Label,
		Description:          strPtrOrNil(sender.Description),
		Tags:                 tagsJSON,
		Caps:                 nil,
		FlowID:               &flowID,
		Transport:            string(sender.Transport),
		DeviceID:             deviceID,
		ManifestHRef:         strPtrOrNil(string(sender.ManifestHRef)),
		InterfaceBindings:    nil,
		SubscriptionReceiver: nil,
		SubscriptionActive:   nil,
		TransportFile:        transportFile,
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	return &line, nil
}

func NewSenderFromV1_2(sender *is04v1_2.Sender, transportFile []byte) (*Sender, error) {
	id, err := uuid.Parse(sender.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tagsJSON, err := json.Marshal(sender.Tags)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(sender.Caps)
	if err != nil {
		return nil, err
	}

	flowID, err := uuid.Parse(sender.FlowID)
	if err != nil {
		return nil, err
	}

	deviceID, err := uuid.Parse(sender.DeviceID)
	if err != nil {
		return nil, err
	}

	subscriptionReceiverRaw := strNilPtrOrNil(sender.Subscription.ReceiverID)
	var subscriptionReceiver *uuid.UUID
	if subscriptionReceiverRaw != nil {
		parsed, err := uuid.Parse(*subscriptionReceiverRaw)
		if err != nil {
			return nil, err
		}
		subscriptionReceiver = &parsed
	} else {
		subscriptionReceiver = nil
	}

	line := Sender{
		ID:                   id,
		ResourceVersion:      sender.Version,
		Label:                sender.Label,
		Description:          strPtrOrNil(sender.Description),
		Tags:                 tagsJSON,
		Caps:                 capsJSON,
		FlowID:               &flowID,
		Transport:            string(sender.Transport),
		DeviceID:             deviceID,
		ManifestHRef:         strPtrOrNil(string(sender.ManifestHRef)),
		InterfaceBindings:    sender.InterfaceBindings,
		SubscriptionReceiver: subscriptionReceiver,
		SubscriptionActive:   &sender.Subscription.Active,
		TransportFile:        transportFile,
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	return &line, nil
}

func NewSenderFromV1_3(sender *is04v1_3.Sender, transportFile []byte) (*Sender, error) {
	id, err := uuid.Parse(sender.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tagsJSON, err := json.Marshal(sender.Tags)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(sender.Caps)
	if err != nil {
		return nil, err
	}

	var flowID uuid.UUID
	if sender.FlowID.Valid {
		flowID, err = uuid.Parse(sender.FlowID.String)
		if err != nil {
			return nil, err
		}
	}

	deviceID, err := uuid.Parse(sender.DeviceID)
	if err != nil {
		return nil, err
	}

	subscriptionReceiverRaw := strNilPtrOrNil(sender.Subscription.ReceiverID)
	var subscriptionReceiver *uuid.UUID
	if subscriptionReceiverRaw != nil {
		parsed, err := uuid.Parse(*subscriptionReceiverRaw)
		if err != nil {
			return nil, err
		}
		subscriptionReceiver = &parsed
	} else {
		subscriptionReceiver = nil
	}

	line := Sender{
		ID:                   id,
		ResourceVersion:      sender.Version,
		Label:                sender.Label,
		Description:          strPtrOrNil(sender.Description),
		Tags:                 tagsJSON,
		Caps:                 capsJSON,
		FlowID:               &flowID,
		Transport:            string(sender.Transport),
		DeviceID:             deviceID,
		ManifestHRef:         strNilPtrOrNil(sender.ManifestHRef),
		InterfaceBindings:    sender.InterfaceBindings,
		SubscriptionReceiver: subscriptionReceiver,
		SubscriptionActive:   &sender.Subscription.Active,
		TransportFile:        transportFile,
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	return &line, nil
}
