//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSize) DeepCopyInto(out *InstanceSize) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSize.
func (in *InstanceSize) DeepCopy() *InstanceSize {
	if in == nil {
		return nil
	}
	out := new(InstanceSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCCluster) DeepCopyInto(out *QCCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCCluster.
func (in *QCCluster) DeepCopy() *QCCluster {
	if in == nil {
		return nil
	}
	out := new(QCCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterList) DeepCopyInto(out *QCClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QCCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterList.
func (in *QCClusterList) DeepCopy() *QCClusterList {
	if in == nil {
		return nil
	}
	out := new(QCClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterSpec) DeepCopyInto(out *QCClusterSpec) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterSpec.
func (in *QCClusterSpec) DeepCopy() *QCClusterSpec {
	if in == nil {
		return nil
	}
	out := new(QCClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterStatus) DeepCopyInto(out *QCClusterStatus) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterStatus.
func (in *QCClusterStatus) DeepCopy() *QCClusterStatus {
	if in == nil {
		return nil
	}
	out := new(QCClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterTemplate) DeepCopyInto(out *QCClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterTemplate.
func (in *QCClusterTemplate) DeepCopy() *QCClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(QCClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterTemplateList) DeepCopyInto(out *QCClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QCClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterTemplateList.
func (in *QCClusterTemplateList) DeepCopy() *QCClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(QCClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterTemplateResource) DeepCopyInto(out *QCClusterTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterTemplateResource.
func (in *QCClusterTemplateResource) DeepCopy() *QCClusterTemplateResource {
	if in == nil {
		return nil
	}
	out := new(QCClusterTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterTemplateSpec) DeepCopyInto(out *QCClusterTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterTemplateSpec.
func (in *QCClusterTemplateSpec) DeepCopy() *QCClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(QCClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCClusterTemplateStatus) DeepCopyInto(out *QCClusterTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCClusterTemplateStatus.
func (in *QCClusterTemplateStatus) DeepCopy() *QCClusterTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(QCClusterTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCEIP) DeepCopyInto(out *QCEIP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCEIP.
func (in *QCEIP) DeepCopy() *QCEIP {
	if in == nil {
		return nil
	}
	out := new(QCEIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCLoadBalancer) DeepCopyInto(out *QCLoadBalancer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCLoadBalancer.
func (in *QCLoadBalancer) DeepCopy() *QCLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(QCLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachine) DeepCopyInto(out *QCMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachine.
func (in *QCMachine) DeepCopy() *QCMachine {
	if in == nil {
		return nil
	}
	out := new(QCMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineList) DeepCopyInto(out *QCMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QCMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineList.
func (in *QCMachineList) DeepCopy() *QCMachineList {
	if in == nil {
		return nil
	}
	out := new(QCMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineSpec) DeepCopyInto(out *QCMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.OSDiskSize != nil {
		in, out := &in.OSDiskSize, &out.OSDiskSize
		*out = new(int)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.SSHKeyID != nil {
		in, out := &in.SSHKeyID, &out.SSHKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineSpec.
func (in *QCMachineSpec) DeepCopy() *QCMachineSpec {
	if in == nil {
		return nil
	}
	out := new(QCMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineStatus) DeepCopyInto(out *QCMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.InstanceStatus != nil {
		in, out := &in.InstanceStatus, &out.InstanceStatus
		*out = new(QCResourceStatus)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineStatus.
func (in *QCMachineStatus) DeepCopy() *QCMachineStatus {
	if in == nil {
		return nil
	}
	out := new(QCMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineTemplate) DeepCopyInto(out *QCMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineTemplate.
func (in *QCMachineTemplate) DeepCopy() *QCMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(QCMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineTemplateList) DeepCopyInto(out *QCMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QCMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineTemplateList.
func (in *QCMachineTemplateList) DeepCopy() *QCMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(QCMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QCMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineTemplateResource) DeepCopyInto(out *QCMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineTemplateResource.
func (in *QCMachineTemplateResource) DeepCopy() *QCMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(QCMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineTemplateSpec) DeepCopyInto(out *QCMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineTemplateSpec.
func (in *QCMachineTemplateSpec) DeepCopy() *QCMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(QCMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCMachineTemplateStatus) DeepCopyInto(out *QCMachineTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCMachineTemplateStatus.
func (in *QCMachineTemplateStatus) DeepCopy() *QCMachineTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(QCMachineTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCNetwork) DeepCopyInto(out *QCNetwork) {
	*out = *in
	out.EIP = in.EIP
	if in.VxNets != nil {
		in, out := &in.VxNets, &out.VxNets
		*out = make([]QCVxNet, len(*in))
		copy(*out, *in)
	}
	out.VPC = in.VPC
	out.APIServerLoadbalancer = in.APIServerLoadbalancer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCNetwork.
func (in *QCNetwork) DeepCopy() *QCNetwork {
	if in == nil {
		return nil
	}
	out := new(QCNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCNetworkResources) DeepCopyInto(out *QCNetworkResources) {
	*out = *in
	out.APIServerLoadbalancersRef = in.APIServerLoadbalancersRef
	out.APIServerLoadbalancersListenerRef = in.APIServerLoadbalancersListenerRef
	out.EIPRef = in.EIPRef
	out.RouterRef = in.RouterRef
	if in.VxNetsRef != nil {
		in, out := &in.VxNetsRef, &out.VxNetsRef
		*out = make([]VxNetRef, len(*in))
		copy(*out, *in)
	}
	out.SecurityGroupRef = in.SecurityGroupRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCNetworkResources.
func (in *QCNetworkResources) DeepCopy() *QCNetworkResources {
	if in == nil {
		return nil
	}
	out := new(QCNetworkResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCResourceReference) DeepCopyInto(out *QCResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCResourceReference.
func (in *QCResourceReference) DeepCopy() *QCResourceReference {
	if in == nil {
		return nil
	}
	out := new(QCResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCVPC) DeepCopyInto(out *QCVPC) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCVPC.
func (in *QCVPC) DeepCopy() *QCVPC {
	if in == nil {
		return nil
	}
	out := new(QCVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QCVxNet) DeepCopyInto(out *QCVxNet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QCVxNet.
func (in *QCVxNet) DeepCopy() *QCVxNet {
	if in == nil {
		return nil
	}
	out := new(QCVxNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VxNetRef) DeepCopyInto(out *VxNetRef) {
	*out = *in
	out.ResourceRef = in.ResourceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VxNetRef.
func (in *VxNetRef) DeepCopy() *VxNetRef {
	if in == nil {
		return nil
	}
	out := new(VxNetRef)
	in.DeepCopyInto(out)
	return out
}
