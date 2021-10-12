// Copyright 2021 Synopsys, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hubclient

import (
	"encoding/json"
	"github.com/blackducksoftware/hub-client-go/hubapi"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (c *Client) CheckHubReadiness(hubUrl string) (error, *hubapi.HealthCheckStatus) {
	readinessUrl := hubapi.BuildUrl(hubUrl, hubapi.ReadinessApi)
	var status hubapi.HealthCheckStatus
	resp, err := c.httpClient.Get(readinessUrl)
	if err != nil {
		log.Error("Error fetching hub health status", err)
		return err, nil
	}

	if resp != nil {
		if err := json.NewDecoder(resp.Body).Decode(&status); err == nil {
			return nil, &status
		}
	}
	return errors.New("error fetching hub health status"), nil
}
