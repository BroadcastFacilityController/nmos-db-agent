package database

import (
	"context"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpsertSender(ctx context.Context, sender *schema.Sender) error {
	now := time.Now().UTC()

	if sender.CreatedAt.IsZero() {
		sender.CreatedAt = now
	}
	sender.UpdatedAt = now

	query := `
INSERT INTO senders (
	id,
	resource_version,
	label,
	description,
	tags,
	caps,
	flow_id,
	transport,
	device_id,
	manifest_href,
	interface_bindings,
	subscription_receiver,
	subscription_active,
	transport_file,
	created_at,
	updated_at
)
VALUES (
	@id,
	@resource_version,
	@label,
	@description,
	@tags,
	@caps,
	@flow_id,
	@transport,
	@device_id,
	@manifest_href,
	@interface_bindings,
	@subscription_receiver,
	@subscription_active,
	@transport_file,
	@created_at,
	@updated_at
)
ON CONFLICT (id) DO UPDATE SET
	resource_version = EXCLUDED.resource_version,
	label = EXCLUDED.label,
	description = EXCLUDED.description,
	tags = EXCLUDED.tags,
	caps = EXCLUDED.caps,
	flow_id = EXCLUDED.flow_id,
	transport = EXCLUDED.transport,
	device_id = EXCLUDED.device_id,
	manifest_href = EXCLUDED.manifest_href,
	interface_bindings = EXCLUDED.interface_bindings,
	subscription_receiver = EXCLUDED.subscription_receiver,
	subscription_active = EXCLUDED.subscription_active,
	transport_file = EXCLUDED.transport_file,
	updated_at = NOW()
WHERE
	senders.resource_version <> EXCLUDED.resource_version
`

	args := pgx.NamedArgs{
		"id":                    sender.ID,
		"resource_version":      sender.ResourceVersion,
		"label":                 sender.Label,
		"description":           sender.Description,
		"tags":                  sender.Tags,
		"caps":                  sender.Caps,
		"flow_id":               sender.FlowID,
		"transport":             sender.Transport,
		"device_id":             sender.DeviceID,
		"manifest_href":         sender.ManifestHRef,
		"interface_bindings":    sender.InterfaceBindings,
		"subscription_receiver": sender.SubscriptionReceiver,
		"subscription_active":   sender.SubscriptionActive,
		"transport_file":        sender.TransportFile,
		"created_at":            sender.CreatedAt,
		"updated_at":            sender.UpdatedAt,
	}

	_, err := db.db.Exec(ctx, query, args)
	return err
}

func SelectSenderByID(ctx context.Context, senderID uuid.UUID) (*schema.Sender, error) {
	query := `
SELECT * FROM senders WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": senderID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	sender, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[schema.Sender])
	if err != nil {
		return nil, err
	}
	return &sender, nil
}
