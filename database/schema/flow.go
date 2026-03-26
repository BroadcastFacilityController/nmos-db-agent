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

type Flow struct {
	// Resource Core
	ID              uuid.UUID
	ResourceVersion string
	Label           string
	Description     *string
	Tags            []byte

	// Flow Core
	SourceID  *uuid.UUID
	DeviceID  *uuid.UUID
	Parents   []uuid.UUID
	GrainRate *[]byte // JSONB

	Format    string
	MediaType *string

	// Flow Video
	FrameWidth             *int
	FrameHeight            *int
	InterlaceMode          *string
	Colorspace             *string
	TransferCharacteristic *string

	// Flow Video Raw
	Components []byte // JSONB

	// Flow Audio
	SampleRate []byte // JSONB

	// Flow Audio Raw
	BitDepth *int

	// Flow JSON Data
	EventType *string

	// Flow SDI_ANC Data
	DID_SDID []byte // JSONB

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewFlowFromV1_0(flow *is04v1_0.Flow) (*Flow, error) {
	id, err := uuid.Parse(flow.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tagsJSON, err := json.Marshal(flow.Tags)
	if err != nil {
		return nil, err
	}

	sourceID, err := uuid.Parse(flow.SourceID)
	if err != nil {
		return nil, err
	}

	parents := make([]uuid.UUID, len(flow.Parents))
	for i, parent := range flow.Parents {
		parents[i], err = uuid.Parse(parent)
		if err != nil {
			return nil, err
		}
	}

	line := Flow{
		ID:                     id,
		ResourceVersion:        flow.Version,
		Label:                  flow.Label,
		Description:            strPtrOrNil(flow.Description),
		Tags:                   tagsJSON,
		SourceID:               &sourceID,
		DeviceID:               nil,
		Parents:                parents,
		GrainRate:              nil,
		Format:                 string(flow.Format),
		MediaType:              nil,
		FrameWidth:             nil,
		FrameHeight:            nil,
		InterlaceMode:          nil,
		Colorspace:             nil,
		TransferCharacteristic: nil,
		Components:             nil,
		SampleRate:             nil,
		BitDepth:               nil,
		EventType:              nil,
		DID_SDID:               nil,

		CreatedAt: now,
		UpdatedAt: now,
	}
	return &line, nil
}

func NewFlowFromV1_1(flow *is04v1_1.Flow) (*Flow, error) {
	switch flow.Type {
	case is04v1_1.FLOW_TYPE_VIDEO_RAW:
		f := flow.FlowVideoRaw
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		componentsJSON, err := json.Marshal(f.Components)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             &f.FrameWidth,
			FrameHeight:            &f.FrameHeight,
			InterlaceMode:          strPtrOrNil(string(f.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(f.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(f.TransferCharacteristic)),
			Components:             componentsJSON,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_1.FLOW_TYPE_VIDEO_CODED:
		f := flow.FlowVideoCoded
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             &f.FrameWidth,
			FrameHeight:            &f.FrameHeight,
			InterlaceMode:          strPtrOrNil(string(f.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(f.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(f.TransferCharacteristic)),
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_1.FLOW_TYPE_AUDIO_RAW:
		f := flow.FlowAudioRaw
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		sampleRateJSON, err := json.Marshal(f.SampleRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJSON,
			BitDepth:               &f.BitDepth,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_1.FLOW_TYPE_AUDIO_CODED:
		f := flow.FlowAudioCoded
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		sampleRateJSON, err := json.Marshal(f.SampleRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJSON,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_1.FLOW_TYPE_DATA:
		f := flow.FlowData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_1.FLOW_TYPE_SDIANC_DATA:
		f := flow.FlowSDIANCData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		didSdid, err := json.Marshal(f.DID_SDID)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               didSdid,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_1.FLOW_TYPE_MUX:
		f := flow.FlowMux
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unknown flow type %s", flow.Type)
	}

}

func NewFlowFromV1_2(flow *is04v1_2.Flow) (*Flow, error) {
	switch flow.Type {
	case is04v1_2.FLOW_TYPE_VIDEO_RAW:
		f := flow.FlowVideoRaw
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		componentsJSON, err := json.Marshal(f.Components)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             &f.FrameWidth,
			FrameHeight:            &f.FrameHeight,
			InterlaceMode:          strPtrOrNil(string(f.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(f.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(f.TransferCharacteristic)),
			Components:             componentsJSON,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_2.FLOW_TYPE_VIDEO_CODED:
		f := flow.FlowVideoCoded
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             &f.FrameWidth,
			FrameHeight:            &f.FrameHeight,
			InterlaceMode:          strPtrOrNil(string(f.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(f.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(f.TransferCharacteristic)),
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_2.FLOW_TYPE_AUDIO_RAW:
		f := flow.FlowAudioRaw
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		sampleRateJSON, err := json.Marshal(f.SampleRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJSON,
			BitDepth:               &f.BitDepth,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_2.FLOW_TYPE_AUDIO_CODED:
		f := flow.FlowAudioCoded
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		sampleRateJSON, err := json.Marshal(f.SampleRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJSON,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_2.FLOW_TYPE_DATA:
		f := flow.FlowData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_2.FLOW_TYPE_SDIANC_DATA:
		f := flow.FlowSDIANCData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		didSdid, err := json.Marshal(f.DID_SDID)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               didSdid,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_2.FLOW_TYPE_MUX:
		f := flow.FlowMux
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unknown flow type %s", flow.Type)
	}

}

func NewFlowFromV1_3(flow *is04v1_3.Flow) (*Flow, error) {
	switch flow.Type {
	case is04v1_3.FLOW_TYPE_VIDEO_RAW:
		f := flow.FlowVideoRaw
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		componentsJSON, err := json.Marshal(f.Components)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             &f.FrameWidth,
			FrameHeight:            &f.FrameHeight,
			InterlaceMode:          strPtrOrNil(string(f.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(f.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(f.TransferCharacteristic)),
			Components:             componentsJSON,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_VIDEO_CODED:
		f := flow.FlowVideoCoded
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             &f.FrameWidth,
			FrameHeight:            &f.FrameHeight,
			InterlaceMode:          strPtrOrNil(string(f.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(f.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(f.TransferCharacteristic)),
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_AUDIO_RAW:
		f := flow.FlowAudioRaw
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		sampleRateJSON, err := json.Marshal(f.SampleRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJSON,
			BitDepth:               &f.BitDepth,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_AUDIO_CODED:
		f := flow.FlowAudioCoded
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		sampleRateJSON, err := json.Marshal(f.SampleRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJSON,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_DATA:
		f := flow.FlowData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_JSON_DATA:
		f := flow.FlowJSONData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              strPtrOrNil(f.EventType),
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_SDIANC_DATA:
		f := flow.FlowSDIANCData
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		didSdid, err := json.Marshal(f.DID_SDID)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               didSdid,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	case is04v1_3.FLOW_TYPE_MUX:
		f := flow.FlowMux
		id, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, err
		}
		now := time.Now()

		tagsJSON, err := json.Marshal(f.Tags)
		if err != nil {
			return nil, err
		}

		sourceID, err := uuid.Parse(f.SourceID)
		if err != nil {
			return nil, err
		}

		deviceID, err := uuid.Parse(f.DeviceID)
		if err != nil {
			return nil, err
		}

		parents := make([]uuid.UUID, len(f.Parents))
		for i, parent := range f.Parents {
			parents[i], err = uuid.Parse(parent)
			if err != nil {
				return nil, err
			}
		}

		grainRateJSON, err := json.Marshal(f.GrainRate)
		if err != nil {
			return nil, err
		}

		line := Flow{
			ID:                     id,
			ResourceVersion:        f.Version,
			Label:                  f.Label,
			Description:            strPtrOrNil(f.Description),
			Tags:                   tagsJSON,
			SourceID:               &sourceID,
			DeviceID:               &deviceID,
			Parents:                parents,
			GrainRate:              &grainRateJSON,
			Format:                 string(f.Format),
			MediaType:              strPtrOrNil(f.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DID_SDID:               nil,

			CreatedAt: now,
			UpdatedAt: now,
		}
		return &line, nil
	default:
		return nil, fmt.Errorf("unknown flow type %s", flow.Type)
	}
}
