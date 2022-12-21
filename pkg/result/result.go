package result

import (
	"encoding/json"
)

type ResultInfo struct {
	Path         string `json:"path"`
	Hash         string `json:"hash"`
	ExpectedHash string `json:"expected_hash"`
	Message      string `json:"message"`
}

type ResultInfos struct {
	ScanTime string       `json:"scan_time"`
	Results  []ResultInfo `json:"results"`
}

func ToJson(infos ResultInfos) (string, error) {
	result, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}
