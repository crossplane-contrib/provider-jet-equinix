/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cloudrouter "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/fabric/cloudrouter"
	connection "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/fabric/connection"
	network "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/fabric/network"
	routingprotocol "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/fabric/routingprotocol"
	serviceprofile "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/fabric/serviceprofile"
	bgpsession "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/bgpsession"
	connectionmetal "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/connection"
	device "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/device"
	devicenetworktype "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/devicenetworktype"
	gateway "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/gateway"
	ipattachment "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/ipattachment"
	organization "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/organization"
	organizationmember "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/organizationmember"
	port "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/port"
	portvlanattachment "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/portvlanattachment"
	project "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/project"
	projectapikey "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/projectapikey"
	projectsshkey "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/projectsshkey"
	reservedipblock "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/reservedipblock"
	spotmarketrequest "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/spotmarketrequest"
	sshkey "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/sshkey"
	userapikey "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/userapikey"
	virtualcircuit "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/virtualcircuit"
	vlan "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/vlan"
	vrf "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/metal/vrf"
	acltemplate "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/acltemplate"
	bgp "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/bgp"
	devicenetwork "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/device"
	devicelink "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/devicelink"
	file "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/file"
	sshkeynetwork "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/sshkey"
	sshuser "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/network/sshuser"
	providerconfig "github.com/crossplane-contrib/provider-jet-equinix/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudrouter.Setup,
		connection.Setup,
		network.Setup,
		routingprotocol.Setup,
		serviceprofile.Setup,
		bgpsession.Setup,
		connectionmetal.Setup,
		device.Setup,
		devicenetworktype.Setup,
		gateway.Setup,
		ipattachment.Setup,
		organization.Setup,
		organizationmember.Setup,
		port.Setup,
		portvlanattachment.Setup,
		project.Setup,
		projectapikey.Setup,
		projectsshkey.Setup,
		reservedipblock.Setup,
		spotmarketrequest.Setup,
		sshkey.Setup,
		userapikey.Setup,
		virtualcircuit.Setup,
		vlan.Setup,
		vrf.Setup,
		acltemplate.Setup,
		bgp.Setup,
		devicenetwork.Setup,
		devicelink.Setup,
		file.Setup,
		sshkeynetwork.Setup,
		sshuser.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
