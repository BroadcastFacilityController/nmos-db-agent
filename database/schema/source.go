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

type Source struct {
	ID              uuid.UUID
	ResourceVersion string
	Label           string
	Description     *string
	Tags            []byte

	GrainRate []byte
	Caps      []byte
	DeviceID  uuid.UUID
	Parents   []string
	ClockName *string

	Format        string
	AudioChannels []byte

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSourceFromV1_0(src *is04v1_0.Source) (*Source, error) {
	id, err := uuid.Parse(src.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	deviceID, err := uuid.Parse(src.DeviceID)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(src.Tags)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(src.Caps)
	if err != nil {
		return nil, err
	}

	line := Source{
		ID:              id,
		ResourceVersion: src.Version,
		Label:           src.Label,
		Description:     strPtrOrNil(src.Description),
		Tags:            tagsJSON,
		GrainRate:       nil,
		Caps:            capsJSON,
		DeviceID:        deviceID,
		Parents:         src.Parents,
		ClockName:       nil,
		Format:          string(src.Format),
		AudioChannels:   nil,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	return &line, nil
}

func NewSourceFromV1_1(source *is04v1_1.Source) (*Source, error) {
	switch source.Type {
	case is04v1_1.SOURCE_TYPE_GENERIC:
		src := source.SourceGeneric
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   nil,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	case is04v1_1.SOURCE_TYPE_AUDIO:
		src := source.SourceAudio
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		audioChannelsJSON, err := json.Marshal(src.Channels)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   audioChannelsJSON,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unable to determine source type %s", source.Type)
	}
}

func NewSourceFromV1_2(source *is04v1_2.Source) (*Source, error) {
	switch source.Type {
	case is04v1_2.SOURCE_TYPE_GENERIC:
		src := source.SourceGeneric
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   nil,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	case is04v1_2.SOURCE_TYPE_AUDIO:
		src := source.SourceAudio
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		audioChannelsJSON, err := json.Marshal(src.Channels)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   audioChannelsJSON,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unable to determine source type %s", source.Type)
	}
}

func NewSourceFromV1_3(source *is04v1_3.Source) (*Source, error) {
	switch source.Type {
	case is04v1_3.SOURCE_TYPE_GENERIC:
		src := source.SourceGeneric
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   nil,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	case is04v1_3.SOURCE_TYPE_DATA:
		src := source.SourceData
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   nil,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	case is04v1_3.SOURCE_TYPE_AUDIO:
		src := source.SourceAudio
		id, err := uuid.Parse(src.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		deviceID, err := uuid.Parse(src.DeviceID)
		if err != nil {
			return nil, err
		}

		tagsJSON, err := json.Marshal(src.Tags)
		if err != nil {
			return nil, err
		}

		grainRateJSON, err := json.Marshal(src.GrainRate)
		if err != nil {
			return nil, err
		}

		capsJSON, err := json.Marshal(src.Caps)
		if err != nil {
			return nil, err
		}

		audioChannelsJSON, err := json.Marshal(src.Channels)
		if err != nil {
			return nil, err
		}

		line := Source{
			ID:              id,
			ResourceVersion: src.Version,
			Label:           src.Label,
			Description:     strPtrOrNil(src.Description),
			Tags:            tagsJSON,
			GrainRate:       grainRateJSON,
			Caps:            capsJSON,
			DeviceID:        deviceID,
			Parents:         src.Parents,
			ClockName:       strNilPtrOrNil(src.ClockName),
			Format:          string(src.Format),
			AudioChannels:   audioChannelsJSON,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unable to determine source type %s", source.Type)
	}
}
