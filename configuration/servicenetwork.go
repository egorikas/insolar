//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package configuration

// ServiceNetwork is configuration for ServiceNetwork.
type ServiceNetwork struct {
	CacheDirectory   string
	ConsensusEnabled bool
}

// NewServiceNetwork creates a new ServiceNetwork configuration.
func NewServiceNetwork() ServiceNetwork {
	return ServiceNetwork{
		CacheDirectory:   "network_cache",
		ConsensusEnabled: true,
	}
}
