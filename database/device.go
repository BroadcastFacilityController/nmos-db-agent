package database

import (
	"context"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpsertDevice(ctx context.Context, device *schema.Device) error {
	now := time.Now().UTC()

	if device.CreatedAt.IsZero() {
		device.CreatedAt = now
	}
	device.UpdatedAt = now

	query := `
INSERT INTO public.devices (
	id,
	resource_version,
	label,
	description,
	tags,
	type,
	receivers,
	senders,
	node_id,
	controls,
	created_at,
	updated_at
)
VALUES (
	@id,
	@resource_version,
	@label,
	@description,
	@tags,
	@type,
	@receivers,
	@senders,
	@node_id,
	@controls,
	@created_at,
	@updated_at
)
ON CONFLICT (id) DO UPDATE SET
	resource_version = EXCLUDED.resource_version,
	label = EXCLUDED.label,
	description = EXCLUDED.description,
	tags = EXCLUDED.tags,
	type = EXCLUDED.type,
	receivers = EXCLUDED.receivers,
	senders = EXCLUDED.senders,
	node_id = EXCLUDED.node_id,
	controls = EXCLUDED.controls,
	updated_at = NOW()
WHERE
	devices.resource_version <> EXCLUDED.resource_version
`

	args := pgx.NamedArgs{
		"id":               device.ID,
		"resource_version": device.ResourceVersion,
		"label":            device.Label,
		"description":      device.Description,
		"tags":             device.Tags,
		"type":             device.Type,
		"receivers":        device.Receivers,
		"senders":          device.Senders,
		"node_id":          device.NodeID,
		"controls":         device.Controls,
		"created_at":       device.CreatedAt,
		"updated_at":       device.UpdatedAt,
	}

	_, err := db.db.Exec(ctx, query, args)
	return err
}

func SelectDeviceByID(ctx context.Context, deviceID uuid.UUID) (*schema.Device, error) {
	query := `
SELECT * FROM devices WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": deviceID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	device, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[schema.Device])
	if err != nil {
		return nil, err
	}
	return &device, nil
}
