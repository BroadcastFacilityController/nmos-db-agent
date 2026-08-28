package watchdog

import (
	"context"
	"encoding/json"

	"github.com/BroadcastFacilityController/nmos-db-agent/nmosdb"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/google/uuid"
)

func (w *Watchdog) handleV1_3Device(ctx context.Context, newDevice is04v1_3.Device) error {
	idParsed, err := uuid.Parse(newDevice.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newDevice.Tags)
	if err != nil {
		return err
	}
	nodeId, err := uuid.Parse(newDevice.NodeID)
	if err != nil {
		return err
	}
	controlsJson, err := json.Marshal(newDevice.Controls)
	if err != nil {
		return err
	}
	apiVersion := "v1.3"
	_, err = w.db.Queries.UpsertDevice(ctx, nmosdb.UpsertDeviceParams{
		ID:              idParsed,
		ResourceVersion: newDevice.Version,
		Label:           newDevice.Label,
		Description:     strPtrOrNil(newDevice.Description),
		Tags:            tagsJson,
		Type:            string(newDevice.Type),
		Receivers:       newDevice.Receivers,
		Senders:         newDevice.Senders,
		NodeID:          uuidPtrOrNil(nodeId),
		Controls:        controlsJson,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_2Device(ctx context.Context, newDevice is04v1_2.Device) error {
	idParsed, err := uuid.Parse(newDevice.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newDevice.Tags)
	if err != nil {
		return err
	}
	nodeId, err := uuid.Parse(newDevice.NodeID)
	if err != nil {
		return err
	}
	controlsJson, err := json.Marshal(newDevice.Controls)
	if err != nil {
		return err
	}
	apiVersion := "v1.2"
	_, err = w.db.Queries.UpsertDevice(ctx, nmosdb.UpsertDeviceParams{
		ID:              idParsed,
		ResourceVersion: newDevice.Version,
		Label:           newDevice.Label,
		Description:     strPtrOrNil(newDevice.Description),
		Tags:            tagsJson,
		Type:            string(newDevice.Type),
		Receivers:       newDevice.Receivers,
		Senders:         newDevice.Senders,
		NodeID:          uuidPtrOrNil(nodeId),
		Controls:        controlsJson,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_1Device(ctx context.Context, newDevice is04v1_1.Device) error {
	idParsed, err := uuid.Parse(newDevice.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newDevice.Tags)
	if err != nil {
		return err
	}
	nodeId, err := uuid.Parse(newDevice.NodeID)
	if err != nil {
		return err
	}
	controlsJson, err := json.Marshal(newDevice.Controls)
	if err != nil {
		return err
	}
	apiVersion := "v1.1"
	_, err = w.db.Queries.UpsertDevice(ctx, nmosdb.UpsertDeviceParams{
		ID:              idParsed,
		ResourceVersion: newDevice.Version,
		Label:           newDevice.Label,
		Description:     strPtrOrNil(newDevice.Description),
		Tags:            tagsJson,
		Type:            string(newDevice.Type),
		Receivers:       newDevice.Receivers,
		Senders:         newDevice.Senders,
		NodeID:          uuidPtrOrNil(nodeId),
		Controls:        controlsJson,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_0Device(ctx context.Context, newDevice is04v1_0.Device) error {
	idParsed, err := uuid.Parse(newDevice.ID)
	if err != nil {
		return err
	}
	nodeId, err := uuid.Parse(newDevice.NodeID)
	if err != nil {
		return err
	}
	description := ""
	apiVersion := "v1.0"
	_, err = w.db.Queries.UpsertDevice(ctx, nmosdb.UpsertDeviceParams{
		ID:              idParsed,
		ResourceVersion: newDevice.Version,
		Label:           newDevice.Label,
		Description:     &description,
		Tags:            make([]byte, 0),
		Type:            string(newDevice.Type),
		Receivers:       newDevice.Receivers,
		Senders:         newDevice.Senders,
		NodeID:          uuidPtrOrNil(nodeId),
		Controls:        make([]byte, 0),
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}
