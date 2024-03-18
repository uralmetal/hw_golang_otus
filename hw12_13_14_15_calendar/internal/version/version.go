package version

import (
	"encoding/json"
	"fmt"
)

var (
	release   = "UNKNOWN"
	buildDate = "UNKNOWN"
	gitHash   = "UNKNOWN"
)

type versionFormat int

const (
	PlainVersion versionFormat = iota
	JSONVersion
)

func GenerateVersion(format versionFormat) string {
	if format == PlainVersion {
		return fmt.Sprintf("Release: %s\nBuild Date: %s\nGit Hash: %s", release, buildDate, gitHash)
	}
	if format == JSONVersion {
		output, err := json.Marshal(struct {
			Release   string
			BuildDate string
			GitHash   string
		}{
			Release:   release,
			BuildDate: buildDate,
			GitHash:   gitHash,
		})
		if err == nil {
			return string(output)
		}
	}
	return "Unknown version"
}

func PrintVersion(format versionFormat) {
	fmt.Println(GenerateVersion(format))
}
