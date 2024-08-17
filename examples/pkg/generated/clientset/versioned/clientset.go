/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen-v0.31. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	examplev1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1"
	examplev1alpha1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1alpha1"
	examplev1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
	examplev2 "acme.corp/pkg/generated/clientset/versioned/typed/example/v2"
	example3v1 "acme.corp/pkg/generated/clientset/versioned/typed/example3/v1"
	existinginterfacesv1 "acme.corp/pkg/generated/clientset/versioned/typed/existinginterfaces/v1"
	secondexamplev1 "acme.corp/pkg/generated/clientset/versioned/typed/secondexample/v1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ExampleV1() examplev1.ExampleV1Interface
	ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface
	ExampleV1beta1() examplev1beta1.ExampleV1beta1Interface
	ExampleV2() examplev2.ExampleV2Interface
	Example3V1() example3v1.Example3V1Interface
	ExistinginterfacesV1() existinginterfacesv1.ExistinginterfacesV1Interface
	SecondexampleV1() secondexamplev1.SecondexampleV1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	exampleV1            *examplev1.ExampleV1Client
	exampleV1alpha1      *examplev1alpha1.ExampleV1alpha1Client
	exampleV1beta1       *examplev1beta1.ExampleV1beta1Client
	exampleV2            *examplev2.ExampleV2Client
	example3V1           *example3v1.Example3V1Client
	existinginterfacesV1 *existinginterfacesv1.ExistinginterfacesV1Client
	secondexampleV1      *secondexamplev1.SecondexampleV1Client
}

// ExampleV1 retrieves the ExampleV1Client
func (c *Clientset) ExampleV1() examplev1.ExampleV1Interface {
	return c.exampleV1
}

// ExampleV1alpha1 retrieves the ExampleV1alpha1Client
func (c *Clientset) ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface {
	return c.exampleV1alpha1
}

// ExampleV1beta1 retrieves the ExampleV1beta1Client
func (c *Clientset) ExampleV1beta1() examplev1beta1.ExampleV1beta1Interface {
	return c.exampleV1beta1
}

// ExampleV2 retrieves the ExampleV2Client
func (c *Clientset) ExampleV2() examplev2.ExampleV2Interface {
	return c.exampleV2
}

// Example3V1 retrieves the Example3V1Client
func (c *Clientset) Example3V1() example3v1.Example3V1Interface {
	return c.example3V1
}

// ExistinginterfacesV1 retrieves the ExistinginterfacesV1Client
func (c *Clientset) ExistinginterfacesV1() existinginterfacesv1.ExistinginterfacesV1Interface {
	return c.existinginterfacesV1
}

// SecondexampleV1 retrieves the SecondexampleV1Client
func (c *Clientset) SecondexampleV1() secondexamplev1.SecondexampleV1Interface {
	return c.secondexampleV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.exampleV1, err = examplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV1alpha1, err = examplev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV1beta1, err = examplev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV2, err = examplev2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.example3V1, err = example3v1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.existinginterfacesV1, err = existinginterfacesv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.secondexampleV1, err = secondexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.exampleV1 = examplev1.New(c)
	cs.exampleV1alpha1 = examplev1alpha1.New(c)
	cs.exampleV1beta1 = examplev1beta1.New(c)
	cs.exampleV2 = examplev2.New(c)
	cs.example3V1 = example3v1.New(c)
	cs.existinginterfacesV1 = existinginterfacesv1.New(c)
	cs.secondexampleV1 = secondexamplev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
