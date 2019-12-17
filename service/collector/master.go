package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelInstallation = "installation"
	labelClusterID    = "cluster_id"
)

var (
	ScheduleDesc *prometheus.Desc = prometheus.NewDesc(
		prometheus.BuildFQName("master_operator", "master", "info"),
		"Master description of the master operator master metric",
		[]string{
			labelInstallation,
			labelClusterID,
		},
		nil,
	)
)

type MasterConfig struct {
}

type Master struct {
}

func NewMaster(config MasterConfig) (*Master, error) {
	r := &Master{}

	return r, nil
}

func (r *Master) Collect(ch chan<- prometheus.Metric) error {
	return nil
}

func (r *Master) Describe(ch chan<- *prometheus.Desc) error {
	ch <- ScheduleDesc

	return nil
}
