// SPDX-License-Identifier: GPL-3.0-or-later

package wmi

import (
	"github.com/netdata/go.d.plugin/pkg/prometheus"
)

const (
	metricCollectorDuration = "windows_exporter_collector_duration_seconds"
	metricCollectorSuccess  = "windows_exporter_collector_success"
)

func (w *WMI) collectCollector(mx map[string]int64, pms prometheus.Metrics) {
	px := "collector_"
	for _, pm := range pms.FindByName(metricCollectorDuration) {
		if name := pm.Labels.Get("collector"); name != "" {
			mx[px+name+"_duration"] = int64(pm.Value * precision)
		}
	}
	for _, pm := range pms.FindByName(metricCollectorSuccess) {
		if name := pm.Labels.Get("collector"); name != "" {
			if pm.Value == 1 {
				mx[px+name+"_status_success"], mx[px+name+"_status_fail"] = 1, 0
			} else {
				mx[px+name+"_status_success"], mx[px+name+"_status_fail"] = 0, 1
			}
		}
	}
}
