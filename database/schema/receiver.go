package schema

import (
	"encoding/json"
	"fmt"
	"time"

	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	uuid "github.com/google/uuid"
)

type Receiver struct {
	// Resource Core
	ID              uuid.UUID
	ResourceVersion string
	Label           string
	Description     *string
	Tags            []byte // JSONB

	// Receiver Core
	DeviceID           uuid.UUID
	Transport          string
	InterfaceBindings  []string
	SubscriptionSender *uuid.UUID
	SubscriptionActive *bool

	Format string
	Caps   []byte // JSONB

	// Database Helpers
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewReceiverFromV1_0(receiver *is04v1_0.Receiver) (*Receiver, error) {
	id, err := uuid.Parse(receiver.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tagsJSON, err := json.Marshal(receiver.Tags)
	if err != nil {
		return nil, err
	}

	deviceID, err := uuid.Parse(receiver.DeviceID)
	if err != nil {
		return nil, err
	}

	var senderID *uuid.UUID
	if receiver.Subscription.SenderID.Valid && receiver.Subscription.SenderID.ValueOrZero() != "" {
		senderIDTemp, err := uuid.Parse(receiver.Subscription.SenderID.String)
		if err != nil {
			return nil, err
		}
		senderID = &senderIDTemp
	} else {
		senderID = nil
	}

	capsJSON, err := json.Marshal(receiver.Caps)
	if err != nil {
		return nil, err
	}

	line := Receiver{
		ID:                 id,
		ResourceVersion:    receiver.Version,
		Label:              receiver.Label,
		Description:        strPtrOrNil(receiver.Description),
		Tags:               tagsJSON,
		DeviceID:           deviceID,
		Transport:          string(receiver.Transport),
		InterfaceBindings:  nil,
		SubscriptionSender: senderID,
		SubscriptionActive: boolZeroPtrOrNil(receiver.Subscription.Active),
		Format:             string(receiver.Format),
		Caps:               capsJSON,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	return &line, nil
}

func NewReceiverFromV1_1(receiver *is04v1_1.Receiver) (*Receiver, error) {
	switch receiver.Type {
	case is04v1_1.RECEIVER_TYPE_VIDEO:
		rx := receiver.ReceiverVideo
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: boolZeroPtrOrNil(rx.Subscription.Active),
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_1.RECEIVER_TYPE_AUDIO:
		rx := receiver.ReceiverAudio
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: boolZeroPtrOrNil(rx.Subscription.Active),
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_1.RECEIVER_TYPE_DATA:
		rx := receiver.ReceiverData
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: boolZeroPtrOrNil(rx.Subscription.Active),
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_1.RECEIVER_TYPE_MUX:
		rx := receiver.ReceiverMux
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: boolZeroPtrOrNil(rx.Subscription.Active),
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unrecoginized type %s", receiver.Type)
	}
}

func NewReceiverFromV1_2(receiver *is04v1_2.Receiver) (*Receiver, error) {
	switch receiver.Type {
	case is04v1_2.RECEIVER_TYPE_VIDEO:
		rx := receiver.ReceiverVideo
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_2.RECEIVER_TYPE_AUDIO:
		rx := receiver.ReceiverAudio
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_2.RECEIVER_TYPE_DATA:
		rx := receiver.ReceiverData
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_2.RECEIVER_TYPE_MUX:
		rx := receiver.ReceiverMux
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unrecoginized type %s", receiver.Type)
	}
}

func NewReceiverFromV1_3(receiver *is04v1_3.Receiver) (*Receiver, error) {
	switch receiver.Type {
	case is04v1_3.RECEIVER_TYPE_VIDEO:
		rx := receiver.ReceiverVideo
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_3.RECEIVER_TYPE_AUDIO:
		rx := receiver.ReceiverAudio
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_3.RECEIVER_TYPE_DATA:
		rx := receiver.ReceiverData
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	case is04v1_3.RECEIVER_TYPE_MUX:
		rx := receiver.ReceiverMux
		id, err := uuid.Parse(rx.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(rx.Tags)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(rx.DeviceID)
		if err != nil {
			return nil, err
		}

		var senderID *uuid.UUID
		if rx.Subscription.SenderID.Valid && rx.Subscription.SenderID.ValueOrZero() != "" {
			senderIDTemp, err := uuid.Parse(rx.Subscription.SenderID.String)
			if err != nil {
				return nil, err
			}
			senderID = &senderIDTemp
		} else {
			senderID = nil
		}

		capsJSON, err := json.Marshal(rx.Caps)
		if err != nil {
			return nil, err
		}

		line := Receiver{
			ID:                 id,
			ResourceVersion:    rx.Version,
			Label:              rx.Label,
			Description:        strPtrOrNil(rx.Description),
			Tags:               tagsJSON,
			DeviceID:           deviceID,
			Transport:          string(rx.Transport),
			InterfaceBindings:  nil,
			SubscriptionSender: senderID,
			SubscriptionActive: &rx.Subscription.Active,
			Format:             string(rx.Format),
			Caps:               capsJSON,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unrecoginized type %s", receiver.Type)
	}
}
