package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
    ChargePercent int
}

var pmsetOutput = regexp.MustCompile("([0-9]+)%")

func ParsePmsetOutput(text string) ( Status, error){
    matches := pmsetOutput.FindStringSubmatch(text)
    if len(matches) < 2 {
        return Status{}, fmt.Errorf("failed to parse pmset output: %q", text)
    }
    charge, err := strconv.Atoi(matches[1])
    if err != nil {
        return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
    }
    return Status{ChargePercent: charge}, nil
}

func GetPmsetOutput() (text string, err error) {
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
