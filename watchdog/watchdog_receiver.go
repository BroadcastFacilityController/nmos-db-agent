package watchdog

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/BroadcastFacilityController/nmos-db-agent/nmosdb"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (w *Watchdog) handleV1_3Receiver(ctx context.Context, newReceiver is04v1_3.Receiver) error {
	apiVersion := "v1.3"
	switch newReceiver.Type {
	case is04v1_3.RECEIVER_TYPE_VIDEO:
		receiver := newReceiver.ReceiverVideo
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.RECEIVER_TYPE_AUDIO:
		receiver := newReceiver.ReceiverAudio
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.RECEIVER_TYPE_DATA:
		receiver := newReceiver.ReceiverData
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.RECEIVER_TYPE_MUX:
		receiver := newReceiver.ReceiverMux
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown receiver type")
	}
}

func (w *Watchdog) handleV1_2Receiver(ctx context.Context, newReceiver is04v1_2.Receiver) error {
	apiVersion := "v1.2"
	switch newReceiver.Type {
	case is04v1_2.RECEIVER_TYPE_VIDEO:
		receiver := newReceiver.ReceiverVideo
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.RECEIVER_TYPE_AUDIO:
		receiver := newReceiver.ReceiverAudio
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.RECEIVER_TYPE_DATA:
		receiver := newReceiver.ReceiverData
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.RECEIVER_TYPE_MUX:
		receiver := newReceiver.ReceiverMux
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  receiver.InterfaceBindings,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Bool: receiver.Subscription.Active, Valid: true},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown receiver type")
	}
}

func (w *Watchdog) handleV1_1Receiver(ctx context.Context, newReceiver is04v1_1.Receiver) error {
	apiVersion := "v1.1"
	switch newReceiver.Type {
	case is04v1_1.RECEIVER_TYPE_VIDEO:
		receiver := newReceiver.ReceiverVideo
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  nil,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Valid: false},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.RECEIVER_TYPE_AUDIO:
		receiver := newReceiver.ReceiverAudio
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  nil,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Valid: false},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.RECEIVER_TYPE_DATA:
		receiver := newReceiver.ReceiverData
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  nil,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Valid: false},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.RECEIVER_TYPE_MUX:
		receiver := newReceiver.ReceiverMux
		idParsed, err := uuid.Parse(receiver.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(receiver.Tags)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
		if err != nil {
			return err
		}
		var subscriptionSender uuid.UUID
		if receiver.Subscription.SenderID.Valid {
			subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
			if err != nil {
				return err
			}
		}
		capsJson, err := json.Marshal(receiver.Caps)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
			ID:                 idParsed,
			ResourceVersion:    receiver.Version,
			Label:              receiver.Label,
			Description:        strPtrOrNil(receiver.Description),
			Tags:               tagsJson,
			DeviceID:           uuidPtrOrNil(deviceIDParsed),
			Transport:          strPtrOrNil(string(receiver.Transport)),
			InterfaceBindings:  nil,
			SubscriptionSender: uuidPtrOrNil(subscriptionSender),
			SubscriptionActive: pgtype.Bool{Valid: false},
			Format:             strPtrOrNil(string(receiver.Format)),
			Caps:               capsJson,
			MetaApiVersion:     apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown receiver type")
	}
}

func (w *Watchdog) handleV1_0Receiver(ctx context.Context, receiver is04v1_0.Receiver) error {
	apiVersion := "v1.0"
	idParsed, err := uuid.Parse(receiver.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(receiver.Tags)
	if err != nil {
		return err
	}
	deviceIDParsed, err := uuid.Parse(receiver.DeviceID)
	if err != nil {
		return err
	}
	var subscriptionSender uuid.UUID
	if receiver.Subscription.SenderID.Valid {
		subscriptionSender, err = uuid.Parse(receiver.Subscription.SenderID.String)
		if err != nil {
			return err
		}
	}
	capsJson, err := json.Marshal(receiver.Caps)
	if err != nil {
		return err
	}
	_, err = w.db.Queries.UpsertReceiver(ctx, nmosdb.UpsertReceiverParams{
		ID:                 idParsed,
		ResourceVersion:    receiver.Version,
		Label:              receiver.Label,
		Description:        strPtrOrNil(receiver.Description),
		Tags:               tagsJson,
		DeviceID:           uuidPtrOrNil(deviceIDParsed),
		Transport:          strPtrOrNil(string(receiver.Transport)),
		InterfaceBindings:  nil,
		SubscriptionSender: uuidPtrOrNil(subscriptionSender),
		SubscriptionActive: pgtype.Bool{Valid: false},
		Format:             strPtrOrNil(string(receiver.Format)),
		Caps:               capsJson,
		MetaApiVersion:     apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}
