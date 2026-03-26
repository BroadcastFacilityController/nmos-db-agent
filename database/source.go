package database

import (
	"context"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpsertSource(ctx context.Context, source *schema.Source) error {
	now := time.Now().UTC()

	if source.CreatedAt.IsZero() {
		source.CreatedAt = now
	}
	source.UpdatedAt = now

	query := `
INSERT INTO sources (
	id,
	resource_version,
	label,
	description,
	tags,
	grain_rate,
	caps,
	deviceid,
	parents,
	clock_name,
	format,
	audio_channels,
	created_at,
	updated_at
)
VALUES (
	@id,
	@resource_version,
	@label,
	@description,
	@tags,
	@grain_rate,
	@caps,
	@deviceid,
	@parents,
	@clock_name,
	@format,
	@audio_channels,
	@created_at,
	@updated_at
)
ON CONFLICT (id) DO UPDATE SET
	resource_version = EXCLUDED.resource_version,
	label = EXCLUDED.label,
	description = EXCLUDED.description,
	tags = EXCLUDED.tags,
	grain_rate = EXCLUDED.grain_rate,
	caps = EXCLUDED.caps,
	deviceid = EXCLUDED.deviceid,
	parents = EXCLUDED.parents,
	clock_name = EXCLUDED.clock_name,
	format = EXCLUDED.format,
	audio_channels = EXCLUDED.audio_channels,
	updated_at = NOW()
WHERE
	sources.resource_version <> EXCLUDED.resource_version
`

	args := pgx.NamedArgs{
		"id":               source.ID,
		"resource_version": source.ResourceVersion,
		"label":            source.Label,
		"description":      source.Description,
		"tags":             source.Tags,
		"grain_rate":       source.GrainRate,
		"caps":             source.Caps,
		"deviceid":         source.DeviceID,
		"parents":          source.Parents,
		"clock_name":       source.ClockName,
		"format":           source.Format,
		"created_at":       source.CreatedAt,
		"updated_at":       source.UpdatedAt,
	}

	_, err := db.db.Exec(ctx, query, args)
	return err
}

func SelectSourceByID(ctx context.Context, sourceID uuid.UUID) (*schema.Source, error) {
	query := `
SELECT * FROM sources WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": sourceID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	source, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[schema.Source])
	if err != nil {
		return nil, err
	}
	return &source, nil
}
