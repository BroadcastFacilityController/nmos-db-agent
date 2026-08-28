package watchdog

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/BroadcastFacilityController/nmos-db-agent/nmosdb"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (w *Watchdog) handleV1_3Sender(ctx context.Context, newSender is04v1_3.Sender) error {
	apiVersion := "v1.3"
	idParsed, err := uuid.Parse(newSender.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newSender.Tags)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newSender.Caps)
	if err != nil {
		return err
	}
	var flowID uuid.UUID
	if newSender.FlowID.Valid {
		flowID, err = uuid.Parse(newSender.FlowID.String)
		if err != nil {
			return err
		}
	}
	deviceIDParsed, err := uuid.Parse(newSender.DeviceID)
	if err != nil {
		return err
	}
	var manifestHRef *string
	var transportFile []byte
	if newSender.ManifestHRef.Valid {
		manifestHRef = &newSender.ManifestHRef.String
		resp, err := http.DefaultClient.Get(newSender.ManifestHRef.String)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = respBody
	}
	var subscriptionReceiver uuid.UUID
	if newSender.Subscription.ReceiverID.Valid {
		subscriptionReceiver, err = uuid.Parse(newSender.Subscription.ReceiverID.String)
		if err != nil {
			return err
		}
	}
	_, err = w.db.Queries.UpsertSender(ctx, nmosdb.UpsertSenderParams{
		ID:                   idParsed,
		ResourceVersion:      newSender.Version,
		Label:                newSender.Label,
		Description:          strPtrOrNil(newSender.Description),
		Tags:                 tagsJson,
		Caps:                 capsJson,
		FlowID:               uuidPtrOrNil(flowID),
		Transport:            string(newSender.Transport),
		DeviceID:             uuidPtrOrNil(deviceIDParsed),
		ManifestHref:         manifestHRef,
		InterfaceBindings:    newSender.InterfaceBindings,
		SubscriptionReceiver: uuidPtrOrNil(subscriptionReceiver),
		SubscriptionActive:   pgtype.Bool{Bool: newSender.Subscription.Active, Valid: true},
		TransportFile:        transportFile,
		MetaApiVersion:       apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_2Sender(ctx context.Context, newSender is04v1_2.Sender) error {
	apiVersion := "v1.2"
	idParsed, err := uuid.Parse(newSender.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newSender.Tags)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newSender.Caps)
	if err != nil {
		return err
	}
	var flowID uuid.UUID
	if newSender.FlowID != "" {
		flowID, err = uuid.Parse(newSender.FlowID)
		if err != nil {
			return err
		}
	}
	deviceIDParsed, err := uuid.Parse(newSender.DeviceID)
	if err != nil {
		return err
	}
	var transportFile []byte
	if newSender.ManifestHRef != "" {
		resp, err := http.DefaultClient.Get(newSender.ManifestHRef)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = respBody
	}
	var subscriptionReceiver uuid.UUID
	if newSender.Subscription.ReceiverID.Valid {
		subscriptionReceiver, err = uuid.Parse(newSender.Subscription.ReceiverID.String)
		if err != nil {
			return err
		}
	}
	_, err = w.db.Queries.UpsertSender(ctx, nmosdb.UpsertSenderParams{
		ID:                   idParsed,
		ResourceVersion:      newSender.Version,
		Label:                newSender.Label,
		Description:          strPtrOrNil(newSender.Description),
		Tags:                 tagsJson,
		Caps:                 capsJson,
		FlowID:               uuidPtrOrNil(flowID),
		Transport:            string(newSender.Transport),
		DeviceID:             uuidPtrOrNil(deviceIDParsed),
		ManifestHref:         strPtrOrNil(newSender.ManifestHRef),
		InterfaceBindings:    newSender.InterfaceBindings,
		SubscriptionReceiver: uuidPtrOrNil(subscriptionReceiver),
		SubscriptionActive:   pgtype.Bool{Bool: newSender.Subscription.Active, Valid: true},
		TransportFile:        transportFile,
		MetaApiVersion:       apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_1Sender(ctx context.Context, newSender is04v1_1.Sender) error {
	apiVersion := "v1.1"
	idParsed, err := uuid.Parse(newSender.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newSender.Tags)
	if err != nil {
		return err
	}
	var flowID uuid.UUID
	if newSender.FlowID != "" {
		flowID, err = uuid.Parse(newSender.FlowID)
		if err != nil {
			return err
		}
	}
	deviceIDParsed, err := uuid.Parse(newSender.DeviceID)
	if err != nil {
		return err
	}
	var transportFile []byte
	if newSender.ManifestHRef != "" {
		resp, err := http.DefaultClient.Get(newSender.ManifestHRef)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = respBody
	}
	_, err = w.db.Queries.UpsertSender(ctx, nmosdb.UpsertSenderParams{
		ID:                   idParsed,
		ResourceVersion:      newSender.Version,
		Label:                newSender.Label,
		Description:          strPtrOrNil(newSender.Description),
		Tags:                 tagsJson,
		Caps:                 nil,
		FlowID:               uuidPtrOrNil(flowID),
		Transport:            string(newSender.Transport),
		DeviceID:             uuidPtrOrNil(deviceIDParsed),
		ManifestHref:         strPtrOrNil(newSender.ManifestHRef),
		InterfaceBindings:    nil,
		SubscriptionReceiver: nil,
		SubscriptionActive:   pgtype.Bool{Valid: false},
		TransportFile:        transportFile,
		MetaApiVersion:       apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_0Sender(ctx context.Context, newSender is04v1_0.Sender) error {
	apiVersion := "v1.0"
	idParsed, err := uuid.Parse(newSender.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newSender.Tags)
	if err != nil {
		return err
	}
	var flowID uuid.UUID
	if newSender.FlowID != "" {
		flowID, err = uuid.Parse(newSender.FlowID)
		if err != nil {
			return err
		}
	}
	deviceIDParsed, err := uuid.Parse(newSender.DeviceID)
	if err != nil {
		return err
	}
	var transportFile []byte
	if newSender.ManifestHRef != "" {
		resp, err := http.DefaultClient.Get(newSender.ManifestHRef)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = respBody
	}
	_, err = w.db.Queries.UpsertSender(ctx, nmosdb.UpsertSenderParams{
		ID:                   idParsed,
		ResourceVersion:      newSender.Version,
		Label:                newSender.Label,
		Description:          strPtrOrNil(newSender.Description),
		Tags:                 tagsJson,
		Caps:                 nil,
		FlowID:               uuidPtrOrNil(flowID),
		Transport:            string(newSender.Transport),
		DeviceID:             uuidPtrOrNil(deviceIDParsed),
		ManifestHref:         strPtrOrNil(newSender.ManifestHRef),
		InterfaceBindings:    nil,
		SubscriptionReceiver: nil,
		SubscriptionActive:   pgtype.Bool{Valid: false},
		TransportFile:        transportFile,
		MetaApiVersion:       apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}
