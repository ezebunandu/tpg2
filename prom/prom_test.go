package prom_test

import (
	"github.com/ezebunandu/prom"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestConfigFromYAML__CorrectlyParsesYAMLData(t *testing.T) {
	t.Parallel()

	want := prom.Config{
		Global: prom.GlobalConfig{
			ScrapeInterval:     15 * time.Second,
			EvaluationInterval: 30 * time.Second,
			ScrapeTimeout:      10 * time.Second,
			ExternalLabels: map[string]string{
				"monitor": "codelab",
				"foo":     "bar",
			},
		},
	}
	got, err := prom.ConfigFromYAML("testdata/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
