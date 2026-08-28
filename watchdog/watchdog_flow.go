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

var ErrMalformedFlowType error = errors.New("malformed flow type")

func (w *Watchdog) handleV1_3Flow(ctx context.Context, newFlow is04v1_3.Flow) error {
	apiVersion := "v1.3"
	switch newFlow.Type {
	case is04v1_3.FLOW_TYPE_VIDEO_RAW:
		flow := newFlow.FlowVideoRaw
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		// Video Raw
		componentsJson, err := json.Marshal(flow.Components)
		if err != nil {
			return err
		}
		frameWidth := int32(flow.FrameWidth)
		frameHeight := int32(flow.FrameHeight)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              &flow.MediaType,
			FrameWidth:             &frameWidth,
			FrameHeight:            &frameHeight,
			InterlaceMode:          strPtrOrNil(string(flow.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(flow.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(flow.TransferCharacteristic)),
			Components:             componentsJson,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_VIDEO_CODED:
		flow := newFlow.FlowVideoCoded
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		frameWidth := int32(flow.FrameWidth)
		frameHeight := int32(flow.FrameHeight)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              &flow.MediaType,
			FrameWidth:             &frameWidth,
			FrameHeight:            &frameHeight,
			InterlaceMode:          strPtrOrNil(string(flow.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(flow.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(flow.TransferCharacteristic)),
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_AUDIO_RAW:
		flow := newFlow.FlowAudioRaw
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		sampleRateJson, err := json.Marshal(flow.SampleRate)
		if err != nil {
			return err
		}
		bitDepth := int32(flow.BitDepth)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJson,
			BitDepth:               &bitDepth,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_AUDIO_CODED:
		flow := newFlow.FlowAudioCoded
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		sampleRateJson, err := json.Marshal(flow.SampleRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJson,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_DATA:
		flow := newFlow.FlowData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_JSON_DATA:
		flow := newFlow.FlowJSONData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              strPtrOrNil(flow.EventType),
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_SDIANC_DATA:
		flow := newFlow.FlowSDIANCData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		did_sdid_json, err := json.Marshal(flow.DID_SDID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                did_sdid_json,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_3.FLOW_TYPE_MUX:
		flow := newFlow.FlowMux
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrMalformedFlowType
	}
}

func (w *Watchdog) handleV1_2Flow(ctx context.Context, newFlow is04v1_2.Flow) error {
	apiVersion := "v1.2"
	switch newFlow.Type {
	case is04v1_2.FLOW_TYPE_VIDEO_RAW:
		flow := newFlow.FlowVideoRaw
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		// Video Raw
		componentsJson, err := json.Marshal(flow.Components)
		if err != nil {
			return err
		}
		frameWidth := int32(flow.FrameWidth)
		frameHeight := int32(flow.FrameHeight)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              &flow.MediaType,
			FrameWidth:             &frameWidth,
			FrameHeight:            &frameHeight,
			InterlaceMode:          strPtrOrNil(string(flow.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(flow.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(flow.TransferCharacteristic)),
			Components:             componentsJson,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.FLOW_TYPE_VIDEO_CODED:
		flow := newFlow.FlowVideoCoded
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		frameWidth := int32(flow.FrameWidth)
		frameHeight := int32(flow.FrameHeight)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              &flow.MediaType,
			FrameWidth:             &frameWidth,
			FrameHeight:            &frameHeight,
			InterlaceMode:          strPtrOrNil(string(flow.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(flow.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(flow.TransferCharacteristic)),
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.FLOW_TYPE_AUDIO_RAW:
		flow := newFlow.FlowAudioRaw
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		sampleRateJson, err := json.Marshal(flow.SampleRate)
		if err != nil {
			return err
		}
		bitDepth := int32(flow.BitDepth)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJson,
			BitDepth:               &bitDepth,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.FLOW_TYPE_AUDIO_CODED:
		flow := newFlow.FlowAudioCoded
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		sampleRateJson, err := json.Marshal(flow.SampleRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJson,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.FLOW_TYPE_DATA:
		flow := newFlow.FlowData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.FLOW_TYPE_SDIANC_DATA:
		flow := newFlow.FlowSDIANCData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		did_sdid_json, err := json.Marshal(flow.DID_SDID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                did_sdid_json,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_2.FLOW_TYPE_MUX:
		flow := newFlow.FlowMux
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrMalformedFlowType
	}
}

func (w *Watchdog) handleV1_1Flow(ctx context.Context, newFlow is04v1_1.Flow) error {
	apiVersion := "v1.1"
	switch newFlow.Type {
	case is04v1_1.FLOW_TYPE_VIDEO_RAW:
		flow := newFlow.FlowVideoRaw
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		// Video Raw
		componentsJson, err := json.Marshal(flow.Components)
		if err != nil {
			return err
		}
		frameWidth := int32(flow.FrameWidth)
		frameHeight := int32(flow.FrameHeight)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              &flow.MediaType,
			FrameWidth:             &frameWidth,
			FrameHeight:            &frameHeight,
			InterlaceMode:          strPtrOrNil(string(flow.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(flow.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(flow.TransferCharacteristic)),
			Components:             componentsJson,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.FLOW_TYPE_VIDEO_CODED:
		flow := newFlow.FlowVideoCoded
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		frameWidth := int32(flow.FrameWidth)
		frameHeight := int32(flow.FrameHeight)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              &flow.MediaType,
			FrameWidth:             &frameWidth,
			FrameHeight:            &frameHeight,
			InterlaceMode:          strPtrOrNil(string(flow.InterlaceMode)),
			Colorspace:             strPtrOrNil(string(flow.Colorspace)),
			TransferCharacteristic: strPtrOrNil(string(flow.TransferCharacteristic)),
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.FLOW_TYPE_AUDIO_RAW:
		flow := newFlow.FlowAudioRaw
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		sampleRateJson, err := json.Marshal(flow.SampleRate)
		if err != nil {
			return err
		}
		bitDepth := int32(flow.BitDepth)
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJson,
			BitDepth:               &bitDepth,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.FLOW_TYPE_AUDIO_CODED:
		flow := newFlow.FlowAudioCoded
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		sampleRateJson, err := json.Marshal(flow.SampleRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             sampleRateJson,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.FLOW_TYPE_DATA:
		flow := newFlow.FlowData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.FLOW_TYPE_SDIANC_DATA:
		flow := newFlow.FlowSDIANCData
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		did_sdid_json, err := json.Marshal(flow.DID_SDID)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                did_sdid_json,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	case is04v1_1.FLOW_TYPE_MUX:
		flow := newFlow.FlowMux
		if flow == nil {
			return ErrMalformedFlowType
		}
		// Resource Core
		idParsed, err := uuid.Parse(flow.ID)
		if err != nil {
			return err
		}
		tagsJson, err := json.Marshal(flow.Tags)
		if err != nil {
			return err
		}
		// Flow core
		sourceIDParsed, err := uuid.Parse(flow.SourceID)
		if err != nil {
			return err
		}
		deviceIDParsed, err := uuid.Parse(flow.DeviceID)
		if err != nil {
			return err
		}
		parentsParsed := make([]uuid.UUID, len(flow.Parents))
		for i, parent := range flow.Parents {
			parentsParsed[i], err = uuid.Parse(parent)
			if err != nil {
				return err
			}
		}
		grainRateJson, err := json.Marshal(flow.GrainRate)
		if err != nil {
			return err
		}
		_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
			ID:                     idParsed,
			ResourceVersion:        flow.Version,
			Label:                  flow.Label,
			Description:            strPtrOrNil(flow.Description),
			Tags:                   tagsJson,
			SourceID:               uuidPtrOrNil(sourceIDParsed),
			DeviceID:               uuidPtrOrNil(deviceIDParsed),
			Parents:                parentsParsed,
			GrainRate:              grainRateJson,
			Format:                 string(flow.Format),
			MediaType:              strPtrOrNil(flow.MediaType),
			FrameWidth:             nil,
			FrameHeight:            nil,
			InterlaceMode:          nil,
			Colorspace:             nil,
			TransferCharacteristic: nil,
			Components:             nil,
			SampleRate:             nil,
			BitDepth:               nil,
			EventType:              nil,
			DidSdid:                nil,
			MetaApiVersion:         apiVersion,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return ErrMalformedFlowType
	}
}

func (w *Watchdog) handleV1_0Flow(ctx context.Context, newFlow is04v1_0.Flow) error {
	apiVersion := "v1.0"
	idParsed, err := uuid.Parse(newFlow.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newFlow.Tags)
	if err != nil {
		return err
	}
	sourceIDParsed, err := uuid.Parse(newFlow.SourceID)
	if err != nil {
		return err
	}
	parentsParsed := make([]uuid.UUID, len(newFlow.Parents))
	for i, parent := range newFlow.Parents {
		parentsParsed[i], err = uuid.Parse(parent)
		if err != nil {
			return err
		}
	}
	_, err = w.db.Queries.UpsertFlow(ctx, nmosdb.UpsertFlowParams{
		ID:                     idParsed,
		ResourceVersion:        newFlow.Version,
		Label:                  newFlow.Label,
		Description:            strPtrOrNil(newFlow.Description),
		Tags:                   tagsJson,
		SourceID:               uuidPtrOrNil(sourceIDParsed),
		DeviceID:               nil,
		Parents:                parentsParsed,
		GrainRate:              nil,
		Format:                 string(newFlow.Format),
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
		DidSdid:                nil,
		MetaApiVersion:         apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}
