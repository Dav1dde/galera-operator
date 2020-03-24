// Copyright 2020 Orange SA
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
// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "galera-operator/pkg/apis/apigalera/v1beta2"
	"galera-operator/pkg/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type SqlV1beta2Interface interface {
	RESTClient() rest.Interface
	GalerasGetter
	GaleraBackupsGetter
	UpgradeConfigsGetter
	UpgradeRulesGetter
}

// SqlV1beta2Client is used to interact with features provided by the sql.databases group.
type SqlV1beta2Client struct {
	restClient rest.Interface
}

func (c *SqlV1beta2Client) Galeras(namespace string) GaleraInterface {
	return newGaleras(c, namespace)
}

func (c *SqlV1beta2Client) GaleraBackups(namespace string) GaleraBackupInterface {
	return newGaleraBackups(c, namespace)
}

func (c *SqlV1beta2Client) UpgradeConfigs(namespace string) UpgradeConfigInterface {
	return newUpgradeConfigs(c, namespace)
}

func (c *SqlV1beta2Client) UpgradeRules(namespace string) UpgradeRuleInterface {
	return newUpgradeRules(c, namespace)
}

// NewForConfig creates a new SqlV1beta2Client for the given config.
func NewForConfig(c *rest.Config) (*SqlV1beta2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SqlV1beta2Client{client}, nil
}

// NewForConfigOrDie creates a new SqlV1beta2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SqlV1beta2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SqlV1beta2Client for the given RESTClient.
func New(c rest.Interface) *SqlV1beta2Client {
	return &SqlV1beta2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SqlV1beta2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
