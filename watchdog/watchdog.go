package watchdog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/database"
	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/BroadcastFacilityController/nmos-go-client"
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/sirupsen/logrus"
)

type Watchdog struct {
	UseTLS                   bool
	sockets                  []*Websocket
	registry                 *nmos.NMOSEndpoint
	subscriptionNodeV1_0     *SubscriptionNode
	subscriptionNodeV1_1     *SubscriptionNode
	subscriptionNodeV1_2     *SubscriptionNode
	subscriptionNodeV1_3     *SubscriptionNode
	subscriptionDeviceV1_0   *SubscriptionDevice
	subscriptionDeviceV1_1   *SubscriptionDevice
	subscriptionDeviceV1_2   *SubscriptionDevice
	subscriptionDeviceV1_3   *SubscriptionDevice
	subscriptionSourceV1_0   *SubscriptionSource
	subscriptionSourceV1_1   *SubscriptionSource
	subscriptionSourceV1_2   *SubscriptionSource
	subscriptionSourceV1_3   *SubscriptionSource
	subscriptionFlowV1_0     *SubscriptionFlow
	subscriptionFlowV1_1     *SubscriptionFlow
	subscriptionFlowV1_2     *SubscriptionFlow
	subscriptionFlowV1_3     *SubscriptionFlow
	subscriptionSenderV1_0   *SubscriptionSender
	subscriptionSenderV1_1   *SubscriptionSender
	subscriptionSenderV1_2   *SubscriptionSender
	subscriptionSenderV1_3   *SubscriptionSender
	subscriptionReceiverV1_0 *SubscriptionReceiver
	subscriptionReceiverV1_1 *SubscriptionReceiver
	subscriptionReceiverV1_2 *SubscriptionReceiver
	subscriptionReceiverV1_3 *SubscriptionReceiver
}

func NewWatchdog(address string, port int, useTLS bool) (*Watchdog, error) {
	url := ""
	if useTLS {
		url += "https://"
	} else {
		url += "http://"
	}
	url += fmt.Sprintf("%s:%d/x-nmos", address, port)
	endpoint, err := nmos.NewNMOSEndpoint(url)
	if err != nil {
		return nil, err
	}
	wd := Watchdog{
		UseTLS:   useTLS,
		registry: endpoint,
		sockets:  make([]*Websocket, 0),
	}
	return &wd, nil
}

func (w *Watchdog) FetchAllFromHttp() error {
	err := w.UpdateNodesFromHttp()
	if err != nil {
		return err
	}
	err = w.UpdateDevicesFromHttp()
	if err != nil {
		return err
	}
	err = w.UpdateSourcesFromHttp()
	if err != nil {
		return err
	}
	err = w.UpdateFlowsFromHttp()
	if err != nil {
		return err
	}
	err = w.UpdateSendersFromHttp()
	if err != nil {
		return err
	}
	err = w.UpdateReceiversFromHttp()
	if err != nil {
		return err
	}
	return nil
}

func (w *Watchdog) UpdateNodesFromHttp() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			nodes, err := w.registry.IS04().V1_0().QueryGetNodes()
			if err != nil {
				return err
			}
			for _, node := range nodes {
				line, err := schema.NewNodeFromV1_0(node)
				if err != nil {
					return err
				}
				logrus.Debugf("HTTP Watchdog Node V1.0: Updating node %s", line.ID)
				database.UpsertNode(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			nodes, err := w.registry.IS04().V1_1().QueryGetNodes()
			if err != nil {
				return err
			}
			for _, node := range nodes {
				line, err := schema.NewNodeFromV1_1(node)
				if err != nil {
					return err
				}
				logrus.Debugf("HTTP Watchdog Node V1.1: Updating node %s", line.ID)
				database.UpsertNode(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			nodes, err := w.registry.IS04().V1_2().QueryGetNodes()
			if err != nil {
				return err
			}
			for _, node := range nodes {
				line, err := schema.NewNodeFromV1_2(node)
				if err != nil {
					return err
				}
				logrus.Debugf("HTTP Watchdog Node V1.2: Updating node %s", line.ID)
				database.UpsertNode(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			nodes, err := w.registry.IS04().V1_3().QueryGetNodes()
			if err != nil {
				return err
			}
			for _, node := range nodes {
				line, err := schema.NewNodeFromV1_3(node)
				if err != nil {
					return err
				}
				logrus.Debugf("HTTP Watchdog Node V1.3: Updating node %s", line.ID)
				database.UpsertNode(context.Background(), line)
			}
		}
	}
	return nil
}

func (w *Watchdog) UpdateDevicesFromHttp() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			devices, err := w.registry.IS04().V1_0().QueryGetDevices()
			if err != nil {
				return err
			}
			for _, device := range devices {
				line, err := schema.NewDeviceFromV1_0(device)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), line.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.0" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Device V1.0: Updating device %s", line.ID)
				database.UpsertDevice(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			devices, err := w.registry.IS04().V1_1().QueryGetDevices()
			if err != nil {
				return err
			}
			for _, device := range devices {
				line, err := schema.NewDeviceFromV1_1(device)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), line.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.1" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Device V1.1: Updating device %s", line.ID)
				database.UpsertDevice(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			devices, err := w.registry.IS04().V1_2().QueryGetDevices()
			if err != nil {
				return err
			}
			for _, device := range devices {
				line, err := schema.NewDeviceFromV1_2(device)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), line.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.2" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Device V1.2: Updating device %s", line.ID)
				database.UpsertDevice(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			devices, err := w.registry.IS04().V1_3().QueryGetDevices()
			if err != nil {
				return err
			}
			for _, device := range devices {
				line, err := schema.NewDeviceFromV1_3(device)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), line.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.3" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Device V1.3: Updating device %s", line.ID)
				database.UpsertDevice(context.Background(), line)
			}
		}
	}
	return nil
}

func (w *Watchdog) UpdateSourcesFromHttp() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			sources, err := w.registry.IS04().V1_0().QueryGetSources()
			if err != nil {
				return err
			}
			for _, source := range sources {
				line, err := schema.NewSourceFromV1_0(&source)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.0" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Source V1.0: Updating source %s", line.ID)
				database.UpsertSource(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			sources, err := w.registry.IS04().V1_1().QueryGetSources()
			if err != nil {
				return err
			}
			for _, source := range sources {
				line, err := schema.NewSourceFromV1_1(&source)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.1" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Source V1.1: Updating source %s", line.ID)
				database.UpsertSource(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			sources, err := w.registry.IS04().V1_2().QueryGetSources()
			if err != nil {
				return err
			}
			for _, source := range sources {
				line, err := schema.NewSourceFromV1_2(&source)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.2" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Source V1.2: Updating source %s", line.ID)
				database.UpsertSource(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			sources, err := w.registry.IS04().V1_3().QueryGetSources()
			if err != nil {
				return err
			}
			for _, source := range sources {
				line, err := schema.NewSourceFromV1_3(&source)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.3" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Source V1.3: Updating source %s", line.ID)
				database.UpsertSource(context.Background(), line)
			}
		}
	}
	return nil
}

func (w *Watchdog) UpdateFlowsFromHttp() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			flows, err := w.registry.IS04().V1_0().QueryGetFlows()
			if err != nil {
				return err
			}
			for _, flow := range flows {
				line, err := schema.NewFlowFromV1_0(&flow)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.0" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Flow V1.0: Updating flow %s", line.ID)
				database.UpsertFlow(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			flows, err := w.registry.IS04().V1_1().QueryGetFlows()
			if err != nil {
				return err
			}
			for _, flow := range flows {
				line, err := schema.NewFlowFromV1_1(&flow)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.1" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Flow V1.1: Updating flow %s", line.ID)
				database.UpsertFlow(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			flows, err := w.registry.IS04().V1_2().QueryGetFlows()
			if err != nil {
				return err
			}
			for _, flow := range flows {
				line, err := schema.NewFlowFromV1_2(&flow)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.2" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Flow V1.2: Updating flow %s", line.ID)
				database.UpsertFlow(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			flows, err := w.registry.IS04().V1_3().QueryGetFlows()
			if err != nil {
				return err
			}
			for _, flow := range flows {
				line, err := schema.NewFlowFromV1_3(&flow)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.3" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Flow V1.3: Updating flow %s", line.ID)
				database.UpsertFlow(context.Background(), line)
			}
		}
	}
	return nil
}

func (w *Watchdog) UpdateSendersFromHttp() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			senders, err := w.registry.IS04().V1_0().QueryGetSenders()
			if err != nil {
				return err
			}
			for _, sender := range senders {
				req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef, nil)
				if err != nil {
					return err
				}
				req.Header.Add("Accept", "application/sdp,application/json,application/xml")
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				line, err := schema.NewSenderFromV1_0(&sender, body)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.0" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Sender V1.0: Updating sender %s", line.ID)
				database.UpsertSender(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			senders, err := w.registry.IS04().V1_1().QueryGetSenders()
			if err != nil {
				return err
			}
			for _, sender := range senders {
				req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef, nil)
				if err != nil {
					return err
				}
				req.Header.Add("Accept", "application/sdp,application/json,application/xml")
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				line, err := schema.NewSenderFromV1_1(&sender, body)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.1" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Sender V1.1: Updating sender %s", line.ID)
				database.UpsertSender(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			senders, err := w.registry.IS04().V1_2().QueryGetSenders()
			if err != nil {
				return err
			}
			for _, sender := range senders {
				req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef, nil)
				if err != nil {
					return err
				}
				req.Header.Add("Accept", "application/sdp,application/json,application/xml")
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				line, err := schema.NewSenderFromV1_2(&sender, body)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.2" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Sender V1.2: Updating sender %s", line.ID)
				database.UpsertSender(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			senders, err := w.registry.IS04().V1_3().QueryGetSenders()
			if err != nil {
				return err
			}
			for _, sender := range senders {
				var body []byte
				if sender.ManifestHRef.Valid {
					req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef.ValueOrZero(), nil)
					if err != nil {
						return err
					}
					req.Header.Add("Accept", "application/sdp,application/json,application/xml")
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						return err
					}
					body, err = io.ReadAll(resp.Body)
					if err != nil {
						return err
					}
				} else {
					body = nil
				}

				line, err := schema.NewSenderFromV1_3(&sender, body)
				if err != nil {
					return err
				}
				device, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
				if err != nil {
					return err
				}
				node, err := database.SelectNodeByID(context.Background(), device.NodeID)
				if err != nil {
					return err
				}
				versions := make([]string, len(node.APIVersions))
				copy(versions, node.SupportedVersions)
				slices.SortFunc(versions, func(a string, b string) int {
					return strings.Compare(a, b)
				})
				if versions[len(versions)-1] != "v1.3" {
					continue
				}
				logrus.Debugf("HTTP Watchdog Sender V1.3: Updating sender %s", line.ID)
				database.UpsertSender(context.Background(), line)
			}
		}
	}
	return nil
}

func (w *Watchdog) UpdateReceiversFromHttp() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			receivers, err := w.registry.IS04().V1_0().QueryGetReceivers()
			if err != nil {
				return err
			}
			for _, receiver := range receivers {
				line, err := schema.NewReceiverFromV1_0(&receiver)
				if err != nil {
					return err
				}
				database.UpsertReceiver(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			receivers, err := w.registry.IS04().V1_1().QueryGetReceivers()
			if err != nil {
				return err
			}
			for _, receiver := range receivers {
				line, err := schema.NewReceiverFromV1_1(&receiver)
				if err != nil {
					return err
				}
				database.UpsertReceiver(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			receivers, err := w.registry.IS04().V1_2().QueryGetReceivers()
			if err != nil {
				return err
			}
			for _, receiver := range receivers {
				line, err := schema.NewReceiverFromV1_2(&receiver)
				if err != nil {
					return err
				}
				database.UpsertReceiver(context.Background(), line)
			}
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			receivers, err := w.registry.IS04().V1_3().QueryGetReceivers()
			if err != nil {
				return err
			}
			for _, receiver := range receivers {
				line, err := schema.NewReceiverFromV1_3(&receiver)
				if err != nil {
					return err
				}
				database.UpsertReceiver(context.Background(), line)
			}
		}
	}
	return nil
}

func (w *Watchdog) SubscribeNodes() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			ws, err := NewWebsocketV1_0(w.registry, is04v1_0.QUERY_NODES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionNodeV1_0, err = NewSubscriptionNode(ws, *common.NewAPIVersion(1, 0))
			if err != nil {
				return err
			}
			w.subscriptionNodeV1_0.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionNodeV1_0.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			ws, err := NewWebsocketV1_1(w.registry, is04v1_1.QUERY_NODES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionNodeV1_1, err = NewSubscriptionNode(ws, *common.NewAPIVersion(1, 1))
			if err != nil {
				return err
			}
			w.subscriptionNodeV1_1.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionNodeV1_1.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			ws, err := NewWebsocketV1_2(w.registry, is04v1_2.QUERY_NODES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionNodeV1_2, err = NewSubscriptionNode(ws, *common.NewAPIVersion(1, 2))
			if err != nil {
				return err
			}
			w.subscriptionNodeV1_2.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionNodeV1_2.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			ws, err := NewWebsocketV1_3(w.registry, is04v1_3.QueryAPISubscriptionResourcePath("/nodes"))
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionNodeV1_3, err = NewSubscriptionNode(ws, *common.NewAPIVersion(1, 3))
			if err != nil {
				return err
			}
			w.subscriptionNodeV1_3.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionNodeV1_3.ws.AwaitFirstMessage(waitCtx)
		}
	}
	return nil
}

func (w *Watchdog) SubscribeDevices() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			ws, err := NewWebsocketV1_0(w.registry, is04v1_0.QUERY_DEVICES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionDeviceV1_0, err = NewSubscriptionDevice(ws, *common.NewAPIVersion(1, 0))
			if err != nil {
				return err
			}
			w.subscriptionDeviceV1_0.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionDeviceV1_0.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			ws, err := NewWebsocketV1_1(w.registry, is04v1_1.QUERY_DEVICES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionDeviceV1_1, err = NewSubscriptionDevice(ws, *common.NewAPIVersion(1, 1))
			if err != nil {
				return err
			}
			w.subscriptionDeviceV1_1.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionDeviceV1_1.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			ws, err := NewWebsocketV1_2(w.registry, is04v1_2.QUERY_DEVICES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionDeviceV1_2, err = NewSubscriptionDevice(ws, *common.NewAPIVersion(1, 2))
			if err != nil {
				return err
			}
			w.subscriptionDeviceV1_2.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionDeviceV1_2.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			ws, err := NewWebsocketV1_3(w.registry, is04v1_3.QueryAPISubscriptionResourcePath("/devices"))
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionDeviceV1_3, err = NewSubscriptionDevice(ws, *common.NewAPIVersion(1, 3))
			if err != nil {
				return err
			}
			w.subscriptionDeviceV1_3.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionDeviceV1_3.ws.AwaitFirstMessage(waitCtx)
		}
	}
	return nil
}

func (w *Watchdog) SubscribeSources() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			ws, err := NewWebsocketV1_0(w.registry, is04v1_0.QUERY_SOURCES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSourceV1_0, err = NewSubscriptionSource(ws, *common.NewAPIVersion(1, 0))
			if err != nil {
				return err
			}
			w.subscriptionSourceV1_0.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSourceV1_0.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			ws, err := NewWebsocketV1_1(w.registry, is04v1_1.QUERY_SOURCES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSourceV1_1, err = NewSubscriptionSource(ws, *common.NewAPIVersion(1, 1))
			if err != nil {
				return err
			}
			w.subscriptionSourceV1_1.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSourceV1_1.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			ws, err := NewWebsocketV1_2(w.registry, is04v1_2.QUERY_SOURCES)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSourceV1_2, err = NewSubscriptionSource(ws, *common.NewAPIVersion(1, 2))
			if err != nil {
				return err
			}
			w.subscriptionSourceV1_2.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSourceV1_2.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			ws, err := NewWebsocketV1_3(w.registry, is04v1_3.QueryAPISubscriptionResourcePath("/sources"))
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSourceV1_3, err = NewSubscriptionSource(ws, *common.NewAPIVersion(1, 3))
			if err != nil {
				return err
			}
			w.subscriptionSourceV1_3.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSourceV1_3.ws.AwaitFirstMessage(waitCtx)
		}
	}
	return nil
}

func (w *Watchdog) SubscribeFlows() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			ws, err := NewWebsocketV1_0(w.registry, is04v1_0.QUERY_FLOWS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionFlowV1_0, err = NewSubscriptionFlow(ws, *common.NewAPIVersion(1, 0))
			if err != nil {
				return err
			}
			w.subscriptionFlowV1_0.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionFlowV1_0.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			ws, err := NewWebsocketV1_1(w.registry, is04v1_1.QUERY_FLOWS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionFlowV1_1, err = NewSubscriptionFlow(ws, *common.NewAPIVersion(1, 1))
			if err != nil {
				return err
			}
			w.subscriptionFlowV1_1.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionFlowV1_1.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			ws, err := NewWebsocketV1_2(w.registry, is04v1_2.QUERY_FLOWS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionFlowV1_2, err = NewSubscriptionFlow(ws, *common.NewAPIVersion(1, 2))
			if err != nil {
				return err
			}
			w.subscriptionFlowV1_2.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionFlowV1_2.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			ws, err := NewWebsocketV1_3(w.registry, is04v1_3.QueryAPISubscriptionResourcePath("/flows"))
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionFlowV1_3, err = NewSubscriptionFlow(ws, *common.NewAPIVersion(1, 3))
			if err != nil {
				return err
			}
			w.subscriptionFlowV1_3.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionFlowV1_3.ws.AwaitFirstMessage(waitCtx)
		}
	}
	return nil
}

func (w *Watchdog) SubscribeSenders() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			ws, err := NewWebsocketV1_0(w.registry, is04v1_0.QUERY_SENDERS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSenderV1_0, err = NewSubscriptionSender(ws, *common.NewAPIVersion(1, 0))
			if err != nil {
				return err
			}
			w.subscriptionSenderV1_0.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSenderV1_0.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			ws, err := NewWebsocketV1_1(w.registry, is04v1_1.QUERY_SENDERS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSenderV1_1, err = NewSubscriptionSender(ws, *common.NewAPIVersion(1, 1))
			if err != nil {
				return err
			}
			w.subscriptionSenderV1_1.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSenderV1_1.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			ws, err := NewWebsocketV1_2(w.registry, is04v1_2.QUERY_SENDERS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSenderV1_2, err = NewSubscriptionSender(ws, *common.NewAPIVersion(1, 2))
			if err != nil {
				return err
			}
			w.subscriptionSenderV1_2.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSenderV1_2.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			ws, err := NewWebsocketV1_3(w.registry, is04v1_3.QueryAPISubscriptionResourcePath("/senders"))
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionSenderV1_3, err = NewSubscriptionSender(ws, *common.NewAPIVersion(1, 3))
			if err != nil {
				return err
			}
			w.subscriptionSenderV1_3.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionSenderV1_3.ws.AwaitFirstMessage(waitCtx)
		}
	}
	return nil
}

func (w *Watchdog) SubscribeReceivers() error {
	// Check if queryAPI is supported
	apis, err := w.registry.GetSupportedAPIs()
	if err != nil {
		return err
	}
	hasQuery := slices.Contains(apis, common.QUERY)
	if !hasQuery {
		return fmt.Errorf("nmos query endpoint is not supported")
	}
	queryAPIVersions, err := w.registry.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return err
	}
	for _, vers := range queryAPIVersions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			ws, err := NewWebsocketV1_0(w.registry, is04v1_0.QUERY_RECEIVERS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionReceiverV1_0, err = NewSubscriptionReceiver(ws, *common.NewAPIVersion(1, 0))
			if err != nil {
				return err
			}
			w.subscriptionReceiverV1_0.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionReceiverV1_0.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			ws, err := NewWebsocketV1_1(w.registry, is04v1_1.QUERY_RECEIVERS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionReceiverV1_1, err = NewSubscriptionReceiver(ws, *common.NewAPIVersion(1, 1))
			if err != nil {
				return err
			}
			w.subscriptionReceiverV1_1.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionReceiverV1_1.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			ws, err := NewWebsocketV1_2(w.registry, is04v1_2.QUERY_RECEIVERS)
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionReceiverV1_2, err = NewSubscriptionReceiver(ws, *common.NewAPIVersion(1, 2))
			if err != nil {
				return err
			}
			w.subscriptionReceiverV1_2.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionReceiverV1_2.ws.AwaitFirstMessage(waitCtx)
		}
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			ws, err := NewWebsocketV1_3(w.registry, is04v1_3.QueryAPISubscriptionResourcePath("/receivers"))
			if err != nil {
				return err
			}
			w.sockets = append(w.sockets, ws)
			w.subscriptionReceiverV1_3, err = NewSubscriptionReceiver(ws, *common.NewAPIVersion(1, 3))
			if err != nil {
				return err
			}
			w.subscriptionReceiverV1_3.Start()
			waitCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			w.subscriptionReceiverV1_3.ws.AwaitFirstMessage(waitCtx)
		}
	}
	return nil
}

func (w *Watchdog) TimeSinceLastMessage() time.Duration {
	lowest := time.Since(time.Unix(0, 0))
	for _, sock := range w.sockets {
		dur := time.Since(sock.lastMessage)
		if dur < lowest {
			lowest = dur
		}
	}
	return lowest
}
