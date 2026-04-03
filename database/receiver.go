package database

import (
	"context"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpsertReceiver(ctx context.Context, receiver *schema.Receiver) error {
	now := time.Now().UTC()

	if receiver.CreatedAt.IsZero() {
		receiver.CreatedAt = now
	}
	receiver.UpdatedAt = now

	query := `
INSERT INTO receivers (
	id,
	resource_version,
	label,
	description,
	tags,
	device_id,
	transport,
	interface_bindings,
	subscription_sender,
	subscription_active,
	format,
	caps,
	created_at,
	updated_at
)
VALUES (
	@id,
	@resource_version,
	@label,
	@description,
	@tags,
	@device_id,
	@transport,
	@interface_bindings,
	@subscription_sender,
	@subscription_active,
	@format,
	@caps,
	@created_at,
	@updated_at
)
ON CONFLICT (id) DO UPDATE SET
	resource_version = EXCLUDED.resource_version,
	label = EXCLUDED.label,
	description = EXCLUDED.description,
	tags = EXCLUDED.tags,
	device_id = EXCLUDED.device_id,
	transport = EXCLUDED.transport,
	interface_bindings = EXCLUDED.interface_bindings,
	subscription_sender = EXCLUDED.subscription_sender,
	subscription_active = EXCLUDED.subscription_active,
	format = EXCLUDED.format,
	caps = EXCLUDED.caps,
	updated_at = NOW()
WHERE
	receivers.resource_version <> EXCLUDED.resource_version
`

	args := pgx.NamedArgs{
		"id":                  receiver.ID,
		"resource_version":    receiver.ResourceVersion,
		"label":               receiver.Label,
		"description":         receiver.Description,
		"tags":                receiver.Tags,
		"device_id":           receiver.DeviceID,
		"transport":           receiver.Transport,
		"interface_bindings":  receiver.InterfaceBindings,
		"subscription_sender": receiver.SubscriptionSender,
		"subscription_active": receiver.SubscriptionActive,
		"format":              receiver.Format,
		"caps":                receiver.Caps,
		"created_at":          receiver.CreatedAt,
		"updated_at":          receiver.UpdatedAt,
	}

	_, err := db.db.Exec(ctx, query, args)
	return err
}

func SelectReceiverByID(ctx context.Context, receiverID uuid.UUID) (*schema.Receiver, error) {
	query := `
SELECT * FROM receivers WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": receiverID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	receiver, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[schema.Receiver])
	if err != nil {
		return nil, err
	}
	return &receiver, nil
}

func SelectReceiversByFormat(ctx context.Context, format string) ([]schema.Receiver, error) {
	query := `
SELECT * FROM receivers WHERE format = @format
	`
	args := pgx.NamedArgs{
		"format": format,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	receivers, err := pgx.CollectRows(rows, pgx.RowToStructByName[schema.Receiver])
	if err != nil {
		return nil, err
	}
	return receivers, nil
}

func SelectReceiverDistinctFormats(ctx context.Context) ([]string, error) {
	query := `
SELECT DISTINCT format FROM receivers
	`
	rows, err := db.db.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	formats, err := pgx.CollectRows(rows, pgx.RowToStructByName[string])
	if err != nil {
		return nil, err
	}
	return formats, nil
}

func SelectReceiversByDeviceID(ctx context.Context, deviceID uuid.UUID) ([]schema.Receiver, error) {
	query := `
SELECT * FROM receivers WHERE device_id = @device_id
	`
	args := pgx.NamedArgs{
		"device_id": deviceID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	receivers, err := pgx.CollectRows(rows, pgx.RowToStructByName[schema.Receiver])
	if err != nil {
		return nil, err
	}
	return receivers, nil
}
