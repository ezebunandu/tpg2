package battery_test

import (
	"os"
	"testing"
    "github.com/ezebunandu/battery"
    "github.com/google/go-cmp/cmp"
)

func TestParsePmsetOutput_GetsChargePercent(t *testing.T){
    t.Parallel()
    data, err := os.ReadFile("testdata/pmset.txt")
    if err != nil {
        t.Fatal(err)
    }
    want := battery.Status{
        ChargePercent: 100,
    }
    got, err := battery.ParsePmsetOutput(string(data))
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}