package database

import (
	"context"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpsertNode(ctx context.Context, node *schema.Node) error {
	now := time.Now().UTC()

	if node.CreatedAt.IsZero() {
		node.CreatedAt = now
	}
	node.UpdatedAt = now

	query := `
INSERT INTO nodes (
	id,
	resource_version,
	label,
	description,
	tags,
	href,
	hostname,
	api_versions,
	api_endpoints,
	caps,
	services,
	clocks,
	interfaces,
	supported_versions,
	created_at,
	updated_at
)
VALUES (
	@id,
	@resource_version,
	@label,
	@description,
	@tags,
	@href,
	@hostname,
	@api_versions,
	@api_endpoints,
	@caps,
	@services,
	@clocks,
	@interfaces,
	@supported_versions,
	@created_at,
	@updated_at
)
ON CONFLICT (id) DO UPDATE SET
	resource_version = EXCLUDED.resource_version,
	label = EXCLUDED.label,
	description = EXCLUDED.description,
	tags = EXCLUDED.tags,
	href = EXCLUDED.href,
	hostname = EXCLUDED.hostname,
	api_versions = EXCLUDED.api_versions,
	api_endpoints = EXCLUDED.api_endpoints,
	caps = EXCLUDED.caps,
	services = EXCLUDED.services,
	clocks = EXCLUDED.clocks,
	interfaces = EXCLUDED.interfaces,
	supported_versions = EXCLUDED.supported_versions,
	updated_at = NOW()
WHERE
	nodes.resource_version <> EXCLUDED.resource_version
`

	args := pgx.NamedArgs{
		"id":                 node.ID,
		"resource_version":   node.ResourceVersion,
		"label":              node.Label,
		"description":        node.Description,
		"tags":               node.Tags,
		"href":               node.HRef,
		"hostname":           node.Hostname,
		"api_versions":       node.APIVersions,
		"api_endpoints":      node.APIEndpoints,
		"caps":               node.Caps,
		"services":           node.Services,
		"clocks":             node.Clocks,
		"interfaces":         node.Interfaces,
		"supported_versions": node.SupportedVersions,
		"created_at":         node.CreatedAt,
		"updated_at":         node.UpdatedAt,
	}

	_, err := db.db.Exec(ctx, query, args)
	return err
}

func SelectNodeByID(ctx context.Context, nodeID uuid.UUID) (*schema.Node, error) {
	query := `
SELECT * FROM nodes WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": nodeID,
	}
	rows, err := db.db.Query(ctx, query, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	node, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[schema.Node])
	if err != nil {
		return nil, err
	}
	return &node, nil
}
