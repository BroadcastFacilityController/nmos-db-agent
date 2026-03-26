package database

import (
	"context"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpsertFlow(ctx context.Context, flow *schema.Flow) error {
	now := time.Now().UTC()

	if flow.CreatedAt.IsZero() {
		flow.CreatedAt = now
	}
	flow.UpdatedAt = now

	query := `
INSERT INTO flows (
	id,
	resource_version,
	label,
	description,
	tags,
	source_id,
	device_id,
	parents,
	grain_rate,
	format,
	media_type,
	frame_width,
	frame_height,
	interlace_mode,
	colorspace,
	transfer_characteristic,
	components,
	sample_rate,
	bit_depth,
	event_type,
	did_sdid,
	created_at,
	updated_at
)
VALUES (
	@id,
	@resource_version,
	@label,
	@description,
	@tags,
	@source_id,
	@device_id,
	@parents,
	@grain_rate,
	@format,
	@media_type,
	@frame_width,
	@frame_height,
	@interlace_mode,
	@colorspace,
	@transfer_characteristic,
	@components,
	@sample_rate,
	@bit_depth,
	@event_type,
	@did_sdid,
	@created_at,
	@updated_at
)
ON CONFLICT (id) DO UPDATE SET
	resource_version = EXCLUDED.resource_version,
	label = EXCLUDED.label,
	description = EXCLUDED.description,
	tags = EXCLUDED.tags,
	source_id = EXCLUDED.source_id,
	device_id = EXCLUDED.device_id,
	parents = EXCLUDED.parents,
	grain_rate = EXCLUDED.grain_rate,
	format = EXCLUDED.format,
	media_type = EXCLUDED.media_type,
	frame_width = EXCLUDED.frame_width,
	frame_height = EXCLUDED.frame_height,
	interlace_mode = EXCLUDED.interlace_mode,
	colorspace = EXCLUDED.colorspace,
	transfer_characteristic = EXCLUDED.transfer_characteristic,
	components = EXCLUDED.components,
	sample_rate = EXCLUDED.sample_rate,
	bit_depth = EXCLUDED.bit_depth,
	event_type = EXCLUDED.event_type,
	did_sdid = EXCLUDED.did_sdid,
	updated_at = EXCLUDED.updated_at
WHERE
	flows.resource_version <> EXCLUDED.resource_version
`

	args := pgx.NamedArgs{
		"id":                      flow.ID,
		"resource_version":        flow.ResourceVersion,
		"label":                   flow.Label,
		"description":             flow.Description,
		"tags":                    flow.Tags,
		"source_id":               flow.SourceID,
		"device_id":               flow.DeviceID,
		"parents":                 flow.Parents,
		"grain_rate":              flow.GrainRate,
		"format":                  flow.Format,
		"media_type":              flow.MediaType,
		"frame_width":             flow.FrameWidth,
		"frame_height":            flow.FrameHeight,
		"interlace_mode":          flow.InterlaceMode,
		"colorspace":              flow.Colorspace,
		"transfer_characteristic": flow.TransferCharacteristic,
		"components":              flow.Components,
		"sample_rate":             flow.SampleRate,
		"bit_depth":               flow.BitDepth,
		"event_type":              flow.EventType,
		"did_sdid":                flow.DID_SDID,
		"updated_at":              flow.UpdatedAt,
		"created_at":              flow.CreatedAt,
	}

	_, err := db.db.Exec(ctx, query, args)
	return err
}

func SelectFlowByID(ctx context.Context, flowID uuid.UUID) (*schema.Flow, error) {
	query := `
SELECT * FROM flows WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": flowID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	flow, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[schema.Flow])
	if err != nil {
		return nil, err
	}
	return &flow, nil
}
