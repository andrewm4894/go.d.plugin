// SPDX-License-Identifier: GPL-3.0-or-later

package wmi

import (
	"github.com/netdata/go.d.plugin/pkg/prometheus"
)

const (
	metricLogonType = "windows_logon_logon_type"
)

func (w *WMI) collectLogon(mx map[string]int64, pms prometheus.Metrics) {
	if !w.cache.collection[collectorLogon] {
		w.cache.collection[collectorLogon] = true
		w.addLogonCharts()
	}

	for _, pm := range pms.FindByName(metricLogonType) {
		if v := pm.Labels.Get("status"); v != "" {
			mx["logon_type_"+v+"_sessions"] = int64(pm.Value)
		}
	}
}
