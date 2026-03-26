package schema

import (
	"encoding/json"
	"sort"
	"time"

	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	uuid "github.com/google/uuid"
)

type Node struct {
	ID              uuid.UUID
	ResourceVersion string
	Label           string
	Description     *string
	Tags            []byte

	HRef     *string
	Hostname *string

	APIVersions  []string
	APIEndpoints []byte

	Caps       []byte
	Services   []byte
	Clocks     []byte
	Interfaces []byte

	SupportedVersions []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewNodeFromV1_0(node is04v1_0.Node) (*Node, error) {
	id, err := uuid.Parse(node.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	servicesJSON, err := json.Marshal(node.Services)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(node.Caps)
	if err != nil {
		return nil, err
	}

	n := Node{
		ID:              id,
		ResourceVersion: node.Version,
		Label:           node.Label,

		HRef:     strPtrOrNil(node.HRef),
		Hostname: strPtrOrNil(node.Hostname),

		Caps:     capsJSON,
		Services: servicesJSON,

		SupportedVersions: []string{"v1.0"},

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}

func NewNodeFromV1_1(node is04v1_1.Node) (*Node, error) {
	id, err := uuid.Parse(node.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	servicesJSON, err := json.Marshal(node.Services)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(node.Caps)
	if err != nil {
		return nil, err
	}

	clocksJSON, err := json.Marshal(node.Clocks)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(node.Tags)
	if err != nil {
		return nil, err
	}

	endpointsJSON, err := json.Marshal(node.API.Endpoints)
	if err != nil {
		return nil, err
	}

	apiVersions := make([]string, len(node.API.Versions))
	for i, v := range node.API.Versions {
		apiVersions[i] = v.String()
	}
	sort.Slice(apiVersions, func(i int, j int) bool {
		return apiVersions[i] > apiVersions[j]
	})
	n := Node{
		ID:              id,
		ResourceVersion: node.Version,
		Label:           node.Label,
		Description:     strPtrOrNil(node.Description),
		Tags:            tagsJSON,

		HRef:     strPtrOrNil(node.HRef),
		Hostname: strPtrOrNil(node.Hostname),

		APIVersions:  apiVersions,
		APIEndpoints: endpointsJSON,

		Caps:     capsJSON,
		Services: servicesJSON,
		Clocks:   clocksJSON,

		SupportedVersions: apiVersions,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}

func NewNodeFromV1_2(node is04v1_2.Node) (*Node, error) {
	id, err := uuid.Parse(node.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	servicesJSON, err := json.Marshal(node.Services)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(node.Caps)
	if err != nil {
		return nil, err
	}

	clocksJSON, err := json.Marshal(node.Clocks)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(node.Tags)
	if err != nil {
		return nil, err
	}

	endpointsJSON, err := json.Marshal(node.API.Endpoints)
	if err != nil {
		return nil, err
	}

	interfacesJSON, err := json.Marshal(node.Interfaces)
	if err != nil {
		return nil, err
	}

	apiVersions := make([]string, len(node.API.Versions))
	for i, v := range node.API.Versions {
		apiVersions[i] = v.String()
	}
	n := Node{
		ID:              id,
		ResourceVersion: node.Version,
		Label:           node.Label,
		Description:     strPtrOrNil(node.Description),
		Tags:            tagsJSON,

		HRef:     strPtrOrNil(node.HRef),
		Hostname: strPtrOrNil(node.Hostname),

		APIVersions:  apiVersions,
		APIEndpoints: endpointsJSON,

		Caps:       capsJSON,
		Services:   servicesJSON,
		Clocks:     clocksJSON,
		Interfaces: interfacesJSON,

		SupportedVersions: apiVersions,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}

func NewNodeFromV1_3(node is04v1_3.Node) (*Node, error) {
	id, err := uuid.Parse(node.ID)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	servicesJSON, err := json.Marshal(node.Services)
	if err != nil {
		return nil, err
	}

	capsJSON, err := json.Marshal(node.Caps)
	if err != nil {
		return nil, err
	}

	clocksJSON, err := json.Marshal(node.Clocks)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(node.Tags)
	if err != nil {
		return nil, err
	}

	endpointsJSON, err := json.Marshal(node.API.Endpoints)
	if err != nil {
		return nil, err
	}

	interfacesJSON, err := json.Marshal(node.Interfaces)
	if err != nil {
		return nil, err
	}

	apiVersions := make([]string, len(node.API.Versions))
	for i, v := range node.API.Versions {
		apiVersions[i] = v.String()
	}
	n := Node{
		ID:              id,
		ResourceVersion: node.Version,
		Label:           node.Label,
		Description:     strPtrOrNil(node.Description),
		Tags:            tagsJSON,

		HRef:     strPtrOrNil(node.HRef),
		Hostname: strPtrOrNil(node.Hostname),

		APIVersions:  apiVersions,
		APIEndpoints: endpointsJSON,

		Caps:       capsJSON,
		Services:   servicesJSON,
		Clocks:     clocksJSON,
		Interfaces: interfacesJSON,

		SupportedVersions: apiVersions,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &n, nil
}
