package exporter

import (
	"github.com/infinityworks/go-common/logger"
	"github.com/prometheus/client_golang/prometheus"
)

// Describe - loops through the API metrics and passes them to prometheus.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, m := range e.APIMetrics {
		ch <- m
	}

}

// Collect function, called on by Prometheus Client library
// This function is called when a scrape is peformed on the /metrics page
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log := logger.Start(e.Config.Config)

	client := e.newClient()

	// Orgs
	e.gatherByOrg(client)

	// Users
	e.gatherByUser(client)

	// Explicit Repos
	e.gatherByRepo(client)

	// Set prometheus gauge metrics using the data gathered
	err := e.processMetrics(ch)

	if err != nil {
		log.Error("Error Processing Metrics", err)
		return
	}

	log.Info("All Metrics successfully collected")

}
