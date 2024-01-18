//
// Copyright (c) 2022 One Track Consulting
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

//go:build !include_nats_messaging

package messaging_test

import (
	"testing"

	"github.com/silenceli/go-mod-messaging/v3/messaging"
	"github.com/silenceli/go-mod-messaging/v3/pkg/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var natsConfig = types.MessageBusConfig{
	Broker: types.HostInfo{
		Host:     "localhost",
		Port:     6379,
		Protocol: "redis",
	},
}

func TestNewMessageClientNatsCore(t *testing.T) {
	messageBusConfig := natsConfig
	messageBusConfig.Type = messaging.NatsCore
	messageBusConfig.Broker = types.HostInfo{Host: uuid.NewString(), Port: 37, Protocol: "nats"}

	_, err := messaging.NewMessageClient(messageBusConfig)

	require.Error(t, err)
}

func TestNewMessageClientNatsJetstream(t *testing.T) {
	messageBusConfig := natsConfig
	messageBusConfig.Type = messaging.NatsJetStream
	messageBusConfig.Broker = types.HostInfo{Host: uuid.NewString(), Port: 37, Protocol: "nats"}

	_, err := messaging.NewMessageClient(messageBusConfig)

	require.Error(t, err)
}
