/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package pulsar

import (
	"net"
	"testing"

	"github.com/insolar/insolar/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockListener struct {
	mock.Mock
}

func (mock *mockListener) Accept() (net.Conn, error) {
	panic("implement me")
}

func (mock *mockListener) Close() error {
	panic("implement me")
}

func (mock *mockListener) Addr() net.Addr {
	panic("implement me")
}

func TestNewPulsar_WithoutNeighbours(t *testing.T) {
	assertObj := assert.New(t)
	config := configuration.Pulsar{ConnectionType: "testType", ListenAddress: "listedAddress"}
	actualConnectionType := ""
	actualAddress := ""

	result, _ := NewPulsar(config, func(connectionType string, address string) (net.Listener, error) {
		actualConnectionType = connectionType
		actualAddress = address
		return &mockListener{}, nil
	})

	assertObj.Equal("testType", actualConnectionType)
	assertObj.Equal("listedAddress", actualAddress)
	assertObj.IsType(result.Sock, &mockListener{})
	assertObj.NotNil(result.PrivateKey)
}

func TestNewPulsar_WithtNeighbours(t *testing.T) {
	assertObj := assert.New(t)
	config := configuration.Pulsar{
		ConnectionType: "testType",
		ListenAddress:  "listedAddress",
		NodesAddresses: []*configuration.PulsarNodeAddress{
			{ConnectionType: "tcp", Address: "first"},
			{ConnectionType: "pct", Address: "second"},
		},
	}

	result, _ := NewPulsar(config, func(connectionType string, address string) (net.Listener, error) {
		return &mockListener{}, nil
	})

	assertObj.Equal(2, len(result.Neighbours))

	assertObj.Equal("tcp", result.Neighbours["first"].ConnectionType.String())
	assertObj.Equal("pct", result.Neighbours["second"].ConnectionType.String())
}
