package reprepro

import (
	"fmt"

	"pault.ag/go/debian/control"
	"pault.ag/go/debian/version"
)

type LogEntry struct {
	Action  string
	Suite   string
	Source  string
	Version version.Version
	Changes control.Changes
}

func (repo Repo) ParseLogEntry(params []string) (*LogEntry, error) {
	if len(params) != 6 {
		return nil, fmt.Errorf("Unknown input string format")
	}

	version, err := version.Parse(params[3])
	if err != nil {
		return nil, err
	}

	changes, err := control.ParseChangesFile(repo.Basedir + "/" + params[5])
	if err != nil {
		return nil, err
	}

	return &LogEntry{
		Action:  params[0],
		Suite:   params[1],
		Source:  params[2],
		Version: version,
		Changes: *changes,
	}, nil
}
