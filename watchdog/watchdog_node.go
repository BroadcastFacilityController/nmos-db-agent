package watchdog

import (
	"context"
	"encoding/json"

	"github.com/BroadcastFacilityController/nmos-db-agent/nmosdb"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/google/uuid"
)

func (w *Watchdog) handleV1_3Node(ctx context.Context, newNode is04v1_3.Node) error {
	idParsed, err := uuid.Parse(newNode.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newNode.Tags)
	if err != nil {
		return err
	}
	apiVersions := make([]string, len(newNode.API.Versions))
	for i, vers := range newNode.API.Versions {
		apiVersions[i] = vers.String()
	}
	apiEndpointsJson, err := json.Marshal(newNode.API.Endpoints)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newNode.Caps)
	if err != nil {
		return err
	}
	servicesJson, err := json.Marshal(newNode.Services)
	if err != nil {
		return err
	}
	clocksJson, err := json.Marshal(newNode.Clocks)
	if err != nil {
		return err
	}
	interfacesJson, err := json.Marshal(newNode.Interfaces)
	if err != nil {
		return err
	}
	apiVersion := "v1.3"
	_, err = w.db.Queries.UpsertNode(ctx, nmosdb.UpsertNodeParams{
		ID:              idParsed,
		ResourceVersion: newNode.Version,
		Label:           newNode.Label,
		Description:     strPtrOrNil(newNode.Description),
		Tags:            tagsJson,
		ApiVersions:     apiVersions,
		ApiEndpoints:    apiEndpointsJson,
		Caps:            capsJson,
		Services:        servicesJson,
		Clocks:          clocksJson,
		Interfaces:      interfacesJson,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_2Node(ctx context.Context, newNode is04v1_2.Node) error {
	idParsed, err := uuid.Parse(newNode.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newNode.Tags)
	if err != nil {
		return err
	}
	apiVersions := make([]string, len(newNode.API.Versions))
	for i, vers := range newNode.API.Versions {
		apiVersions[i] = vers.String()
	}
	apiEndpointsJson, err := json.Marshal(newNode.API.Endpoints)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newNode.Caps)
	if err != nil {
		return err
	}
	servicesJson, err := json.Marshal(newNode.Services)
	if err != nil {
		return err
	}
	clocksJson, err := json.Marshal(newNode.Clocks)
	if err != nil {
		return err
	}
	interfacesJson, err := json.Marshal(newNode.Interfaces)
	if err != nil {
		return err
	}
	apiVersion := "v1.2"
	_, err = w.db.Queries.UpsertNode(ctx, nmosdb.UpsertNodeParams{
		ID:              idParsed,
		ResourceVersion: newNode.Version,
		Label:           newNode.Label,
		Description:     strPtrOrNil(newNode.Description),
		Tags:            tagsJson,
		ApiVersions:     apiVersions,
		ApiEndpoints:    apiEndpointsJson,
		Caps:            capsJson,
		Services:        servicesJson,
		Clocks:          clocksJson,
		Interfaces:      interfacesJson,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_1Node(ctx context.Context, newNode is04v1_1.Node) error {
	idParsed, err := uuid.Parse(newNode.ID)
	if err != nil {
		return err
	}
	tagsJson, err := json.Marshal(newNode.Tags)
	if err != nil {
		return err
	}
	apiVersions := make([]string, len(newNode.API.Versions))
	for i, vers := range newNode.API.Versions {
		apiVersions[i] = vers.String()
	}
	apiEndpointsJson, err := json.Marshal(newNode.API.Endpoints)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newNode.Caps)
	if err != nil {
		return err
	}
	servicesJson, err := json.Marshal(newNode.Services)
	if err != nil {
		return err
	}
	clocksJson, err := json.Marshal(newNode.Clocks)
	if err != nil {
		return err
	}
	apiVersion := "v1.1"
	_, err = w.db.Queries.UpsertNode(ctx, nmosdb.UpsertNodeParams{
		ID:              idParsed,
		ResourceVersion: newNode.Version,
		Label:           newNode.Label,
		Description:     strPtrOrNil(newNode.Description),
		Tags:            tagsJson,
		ApiVersions:     apiVersions,
		ApiEndpoints:    apiEndpointsJson,
		Caps:            capsJson,
		Services:        servicesJson,
		Clocks:          clocksJson,
		Interfaces:      nil,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) handleV1_0Node(ctx context.Context, newNode is04v1_0.Node) error {
	idParsed, err := uuid.Parse(newNode.ID)
	if err != nil {
		return err
	}
	capsJson, err := json.Marshal(newNode.Caps)
	if err != nil {
		return err
	}
	servicesJson, err := json.Marshal(newNode.Services)
	if err != nil {
		return err
	}
	description := ""
	apiVersion := "v1.0"
	_, err = w.db.Queries.UpsertNode(ctx, nmosdb.UpsertNodeParams{
		ID:              idParsed,
		ResourceVersion: newNode.Version,
		Label:           newNode.Label,
		Description:     strPtrOrNil(description),
		Tags:            nil,
		ApiVersions:     nil,
		ApiEndpoints:    nil,
		Caps:            capsJson,
		Services:        servicesJson,
		Clocks:          nil,
		Interfaces:      nil,
		MetaApiVersion:  apiVersion,
	})
	if err != nil {
		return err
	}
	return nil
}
