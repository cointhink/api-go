package lxd

import "time"

type LxcStatus struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Operation  string `json:"operation"`
	ErrorCode  int    `json:"error_code"`
	Error      string `json:"error"`
	Metadata   struct {
		Architecture string `json:"architecture"`
		Config       struct {
			ImageArchitecture      string `json:"image.architecture"`
			ImageDescription       string `json:"image.description"`
			ImageOs                string `json:"image.os"`
			ImageRelease           string `json:"image.release"`
			ImageSerial            string `json:"image.serial"`
			VolatileApplyTemplate  string `json:"volatile.apply_template"`
			VolatileBaseImage      string `json:"volatile.base_image"`
			VolatileEth0Hwaddr     string `json:"volatile.eth0.hwaddr"`
			VolatileEth0Name       string `json:"volatile.eth0.name"`
			VolatileIdmapBase      string `json:"volatile.idmap.base"`
			VolatileIdmapNext      string `json:"volatile.idmap.next"`
			VolatileLastStateIdmap string `json:"volatile.last_state.idmap"`
		} `json:"config"`
		Devices struct {
		} `json:"devices"`
		Ephemeral      bool      `json:"ephemeral"`
		Profiles       []string  `json:"profiles"`
		Stateful       bool      `json:"stateful"`
		CreatedAt      time.Time `json:"created_at"`
		ExpandedConfig struct {
			ImageArchitecture      string `json:"image.architecture"`
			ImageDescription       string `json:"image.description"`
			ImageOs                string `json:"image.os"`
			ImageRelease           string `json:"image.release"`
			ImageSerial            string `json:"image.serial"`
			VolatileApplyTemplate  string `json:"volatile.apply_template"`
			VolatileBaseImage      string `json:"volatile.base_image"`
			VolatileEth0Hwaddr     string `json:"volatile.eth0.hwaddr"`
			VolatileEth0Name       string `json:"volatile.eth0.name"`
			VolatileIdmapBase      string `json:"volatile.idmap.base"`
			VolatileIdmapNext      string `json:"volatile.idmap.next"`
			VolatileLastStateIdmap string `json:"volatile.last_state.idmap"`
		} `json:"expanded_config"`
		ExpandedDevices struct {
			Eth0 struct {
				Nictype string `json:"nictype"`
				Parent  string `json:"parent"`
				Type    string `json:"type"`
			} `json:"eth0"`
			Root struct {
				Path string `json:"path"`
				Pool string `json:"pool"`
				Type string `json:"type"`
			} `json:"root"`
		} `json:"expanded_devices"`
		Name       string    `json:"name"`
		Status     string    `json:"status"`
		StatusCode int       `json:"status_code"`
		LastUsedAt time.Time `json:"last_used_at"`
	} `json:"metadata"`
}
