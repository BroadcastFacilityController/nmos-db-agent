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
)

var ErrMalformedSourceType error = errors.New("malformed source type")

func (w *Watchdog) handleV1_3Source(ctx context.Context, newSource is04v1_3.Source) error {
	apiVersion := "v1.3"
	switch newSource.Type {
	case is04v1_3.SOURCE_TYPE_GENERIC:
		src := newSource.SourceGeneric
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.SOURCE_TYPE_AUDIO:
		src := newSource.SourceAudio
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		audioChannelsJson, err := json.Marshal(src.Channels)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			AudioChannels:   audioChannelsJson,
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.SOURCE_TYPE_DATA:
		src := newSource.SourceData
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrMalformedSourceType
	}
}

func (w *Watchdog) handleV1_2Source(ctx context.Context, newSource is04v1_2.Source) error {
	apiVersion := "v1.2"
	switch newSource.Type {
	case is04v1_2.SOURCE_TYPE_GENERIC:
		src := newSource.SourceGeneric
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.SOURCE_TYPE_AUDIO:
		src := newSource.SourceAudio
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		audioChannelsJson, err := json.Marshal(src.Channels)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			AudioChannels:   audioChannelsJson,
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrMalformedSourceType
	}
}

func (w *Watchdog) handleV1_1Source(ctx context.Context, newSource is04v1_1.Source) error {
	apiVersion := "v1.1"
	switch newSource.Type {
	case is04v1_1.SOURCE_TYPE_GENERIC:
		src := newSource.SourceGeneric
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.SOURCE_TYPE_AUDIO:
		src := newSource.SourceAudio
		if src == nil {
			return ErrMalformedSourceType
		}
		idParsed, err := uuid.Parse(src.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(src.Tags)
		if err != nil {
			return err
		}
		grainRateJson, err := json.Marshal(src.GrainRate)
		if err != nil {
			return err
		}
		capsJson, err := json.Marshal(src.Caps)
		if err != nil {
			return err
		}
		deviceIdParsed, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return err
		}
		audioChannelsJson, err := json.Marshal(src.Channels)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
			ID:              idParsed,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJson,
			GrainRate:       grainRateJson,
			Caps:            capsJson,
			DeviceID:        uuidPtrOrNil(deviceIdParsed),
			Parents:         src.Parents,
			ClockName:       src.ClockName.Ptr(),
			Format:          string(src.Format),
			AudioChannels:   audioChannelsJson,
			MetaApiVersion:  apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrMalformedSourceType
	}
}

func (w *Watchdog) handleV1_0Source(ctx context.Context, newSource is04v1_0.Source) error {
	apiVersion := "v1.1"
	idParsed, err := uuid.Parse(newSource.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newSource.Tags)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newSource.Caps)
	if err != nil {
		return err
	}
	deviceIdParsed, err := uuid.Parse(newSource.DeviceID)
	if err != nil {
		return err
	}
	_, err = w.db.Queries.UpsertSource(ctx, nmosdb.UpsertSourceParams{
		ID:              idParsed,
		ResourceVersion: newSource.Version,
		Label:           newSource.Label,
		Description:     strPtrOrNil(newSource.Description),
		Tags:            tagsJson,
		GrainRate:       nil,
		Caps:            capsJson,
		DeviceID:        uuidPtrOrNil(deviceIdParsed),
		Parents:         newSource.Parents,
		ClockName:       nil,
		Format:          string(newSource.Format),
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}
