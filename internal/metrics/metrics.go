package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createQuestionCounter prometheus.Counter
	updateQuestionCounter prometheus.Counter
	removeQuestionCounter prometheus.Counter
)

func RegisterMetrics() {
	createQuestionCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_question_api_create_count_total",
		Help: "The total create question",
	})

	updateQuestionCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_question_api_update_count_total",
		Help: "The total update question",
	})

	removeQuestionCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_question_api_remove_count_total",
		Help: "The total remove question",
	})
}

func CreateQuestionCounterInc() {
	createQuestionCounter.Inc()
}

func UpdateQuestionCounterInc() {
	updateQuestionCounter.Inc()
}

func RemoveQuestionCounterInc() {
	removeQuestionCounter.Inc()
}
