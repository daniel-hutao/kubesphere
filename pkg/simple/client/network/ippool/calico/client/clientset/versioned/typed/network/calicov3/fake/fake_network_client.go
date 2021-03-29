/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	calicov3 "kubesphere.io/kubesphere/pkg/simple/client/network/ippool/calico/client/clientset/versioned/typed/network/calicov3"
)

type FakeCrdCalicov3 struct {
	*testing.Fake
}

func (c *FakeCrdCalicov3) BlockAffinities() calicov3.BlockAffinityInterface {
	return &FakeBlockAffinities{c}
}

func (c *FakeCrdCalicov3) IPAMBlocks() calicov3.IPAMBlockInterface {
	return &FakeIPAMBlocks{c}
}

func (c *FakeCrdCalicov3) IPPools() calicov3.IPPoolInterface {
	return &FakeIPPools{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrdCalicov3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
