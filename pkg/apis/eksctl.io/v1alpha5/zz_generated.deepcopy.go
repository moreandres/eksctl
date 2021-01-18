// +build !ignore_autogenerated

/*
Copyright 2018 Weaveworks. All rights reserved.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha5

import (
	ipnet "github.com/weaveworks/eksctl/pkg/utils/ipnet"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AZSubnetMapping) DeepCopyInto(out *AZSubnetMapping) {
	{
		in := &in
		*out = make(AZSubnetMapping, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AZSubnetMapping.
func (in AZSubnetMapping) DeepCopy() AZSubnetMapping {
	if in == nil {
		return nil
	}
	out := new(AZSubnetMapping)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AZSubnetSpec) DeepCopyInto(out *AZSubnetSpec) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AZSubnetSpec.
func (in *AZSubnetSpec) DeepCopy() *AZSubnetSpec {
	if in == nil {
		return nil
	}
	out := new(AZSubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.AttachPolicyARNs != nil {
		in, out := &in.AttachPolicyARNs, &out.AttachPolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AttachPolicy.DeepCopyInto(&out.AttachPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCloudWatch) DeepCopyInto(out *ClusterCloudWatch) {
	*out = *in
	if in.ClusterLogging != nil {
		in, out := &in.ClusterLogging, &out.ClusterLogging
		*out = new(ClusterCloudWatchLogging)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCloudWatch.
func (in *ClusterCloudWatch) DeepCopy() *ClusterCloudWatch {
	if in == nil {
		return nil
	}
	out := new(ClusterCloudWatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCloudWatchLogging) DeepCopyInto(out *ClusterCloudWatchLogging) {
	*out = *in
	if in.EnableTypes != nil {
		in, out := &in.EnableTypes, &out.EnableTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCloudWatchLogging.
func (in *ClusterCloudWatchLogging) DeepCopy() *ClusterCloudWatchLogging {
	if in == nil {
		return nil
	}
	out := new(ClusterCloudWatchLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(ClusterMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesNetworkConfig != nil {
		in, out := &in.KubernetesNetworkConfig, &out.KubernetesNetworkConfig
		*out = new(KubernetesNetworkConfig)
		**out = **in
	}
	if in.IAM != nil {
		in, out := &in.IAM, &out.IAM
		*out = new(ClusterIAM)
		(*in).DeepCopyInto(*out)
	}
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(ClusterVPC)
		(*in).DeepCopyInto(*out)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]*Addon, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Addon)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.PrivateCluster != nil {
		in, out := &in.PrivateCluster, &out.PrivateCluster
		*out = new(PrivateCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]*NodeGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ManagedNodeGroups != nil {
		in, out := &in.ManagedNodeGroups, &out.ManagedNodeGroups
		*out = make([]*ManagedNodeGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ManagedNodeGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.FargateProfiles != nil {
		in, out := &in.FargateProfiles, &out.FargateProfiles
		*out = make([]*FargateProfile, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FargateProfile)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CloudWatch != nil {
		in, out := &in.CloudWatch, &out.CloudWatch
		*out = new(ClusterCloudWatch)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretsEncryption != nil {
		in, out := &in.SecretsEncryption, &out.SecretsEncryption
		*out = new(SecretsEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(Git)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigList) DeepCopyInto(out *ClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigList.
func (in *ClusterConfigList) DeepCopy() *ClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpoints) DeepCopyInto(out *ClusterEndpoints) {
	*out = *in
	if in.PrivateAccess != nil {
		in, out := &in.PrivateAccess, &out.PrivateAccess
		*out = new(bool)
		**out = **in
	}
	if in.PublicAccess != nil {
		in, out := &in.PublicAccess, &out.PublicAccess
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpoints.
func (in *ClusterEndpoints) DeepCopy() *ClusterEndpoints {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIAM) DeepCopyInto(out *ClusterIAM) {
	*out = *in
	if in.ServiceRoleARN != nil {
		in, out := &in.ServiceRoleARN, &out.ServiceRoleARN
		*out = new(string)
		**out = **in
	}
	if in.ServiceRolePermissionsBoundary != nil {
		in, out := &in.ServiceRolePermissionsBoundary, &out.ServiceRolePermissionsBoundary
		*out = new(string)
		**out = **in
	}
	if in.FargatePodExecutionRoleARN != nil {
		in, out := &in.FargatePodExecutionRoleARN, &out.FargatePodExecutionRoleARN
		*out = new(string)
		**out = **in
	}
	if in.FargatePodExecutionRolePermissionsBoundary != nil {
		in, out := &in.FargatePodExecutionRolePermissionsBoundary, &out.FargatePodExecutionRolePermissionsBoundary
		*out = new(string)
		**out = **in
	}
	if in.WithOIDC != nil {
		in, out := &in.WithOIDC, &out.WithOIDC
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = make([]*ClusterIAMServiceAccount, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClusterIAMServiceAccount)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VPCResourceControllerPolicy != nil {
		in, out := &in.VPCResourceControllerPolicy, &out.VPCResourceControllerPolicy
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIAM.
func (in *ClusterIAM) DeepCopy() *ClusterIAM {
	if in == nil {
		return nil
	}
	out := new(ClusterIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIAMMeta) DeepCopyInto(out *ClusterIAMMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIAMMeta.
func (in *ClusterIAMMeta) DeepCopy() *ClusterIAMMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterIAMMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIAMServiceAccount) DeepCopyInto(out *ClusterIAMServiceAccount) {
	*out = *in
	in.ClusterIAMMeta.DeepCopyInto(&out.ClusterIAMMeta)
	if in.AttachPolicyARNs != nil {
		in, out := &in.AttachPolicyARNs, &out.AttachPolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AttachPolicy.DeepCopyInto(&out.AttachPolicy)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ClusterIAMServiceAccountStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIAMServiceAccount.
func (in *ClusterIAMServiceAccount) DeepCopy() *ClusterIAMServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ClusterIAMServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIAMServiceAccountStatus) DeepCopyInto(out *ClusterIAMServiceAccountStatus) {
	*out = *in
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIAMServiceAccountStatus.
func (in *ClusterIAMServiceAccountStatus) DeepCopy() *ClusterIAMServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterIAMServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMeta) DeepCopyInto(out *ClusterMeta) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMeta.
func (in *ClusterMeta) DeepCopy() *ClusterMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNAT) DeepCopyInto(out *ClusterNAT) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNAT.
func (in *ClusterNAT) DeepCopy() *ClusterNAT {
	if in == nil {
		return nil
	}
	out := new(ClusterNAT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSubnets) DeepCopyInto(out *ClusterSubnets) {
	*out = *in
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = make(AZSubnetMapping, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = make(AZSubnetMapping, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSubnets.
func (in *ClusterSubnets) DeepCopy() *ClusterSubnets {
	if in == nil {
		return nil
	}
	out := new(ClusterSubnets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVPC) DeepCopyInto(out *ClusterVPC) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = new(ClusterSubnets)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraCIDRs != nil {
		in, out := &in.ExtraCIDRs, &out.ExtraCIDRs
		*out = make([]*ipnet.IPNet, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.AutoAllocateIPv6 != nil {
		in, out := &in.AutoAllocateIPv6, &out.AutoAllocateIPv6
		*out = new(bool)
		**out = **in
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(ClusterNAT)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterEndpoints != nil {
		in, out := &in.ClusterEndpoints, &out.ClusterEndpoints
		*out = new(ClusterEndpoints)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicAccessCIDRs != nil {
		in, out := &in.PublicAccessCIDRs, &out.PublicAccessCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVPC.
func (in *ClusterVPC) DeepCopy() *ClusterVPC {
	if in == nil {
		return nil
	}
	out := new(ClusterVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfile) DeepCopyInto(out *FargateProfile) {
	*out = *in
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]FargateProfileSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfile.
func (in *FargateProfile) DeepCopy() *FargateProfile {
	if in == nil {
		return nil
	}
	out := new(FargateProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfileSelector) DeepCopyInto(out *FargateProfileSelector) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfileSelector.
func (in *FargateProfileSelector) DeepCopy() *FargateProfileSelector {
	if in == nil {
		return nil
	}
	out := new(FargateProfileSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Git) DeepCopyInto(out *Git) {
	*out = *in
	if in.Repo != nil {
		in, out := &in.Repo, &out.Repo
		*out = new(Repo)
		(*in).DeepCopyInto(*out)
	}
	in.Operator.DeepCopyInto(&out.Operator)
	if in.BootstrapProfile != nil {
		in, out := &in.BootstrapProfile, &out.BootstrapProfile
		*out = new(Profile)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Git.
func (in *Git) DeepCopy() *Git {
	if in == nil {
		return nil
	}
	out := new(Git)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in InlineDocument) DeepCopyInto(out *InlineDocument) {
	{
		in := &in
		clone := in.DeepCopy()
		*out = *clone
		return
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNetworkConfig) DeepCopyInto(out *KubernetesNetworkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNetworkConfig.
func (in *KubernetesNetworkConfig) DeepCopy() *KubernetesNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(KubernetesNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplate) DeepCopyInto(out *LaunchTemplate) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplate.
func (in *LaunchTemplate) DeepCopy() *LaunchTemplate {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedNodeGroup) DeepCopyInto(out *ManagedNodeGroup) {
	*out = *in
	if in.NodeGroupBase != nil {
		in, out := &in.NodeGroupBase, &out.NodeGroupBase
		*out = new(NodeGroupBase)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LaunchTemplate != nil {
		in, out := &in.LaunchTemplate, &out.LaunchTemplate
		*out = new(LaunchTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedNodeGroup.
func (in *ManagedNodeGroup) DeepCopy() *ManagedNodeGroup {
	if in == nil {
		return nil
	}
	out := new(ManagedNodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsCollection) DeepCopyInto(out *MetricsCollection) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsCollection.
func (in *MetricsCollection) DeepCopy() *MetricsCollection {
	if in == nil {
		return nil
	}
	out := new(MetricsCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.NodeGroupBase != nil {
		in, out := &in.NodeGroupBase, &out.NodeGroupBase
		*out = new(NodeGroupBase)
		(*in).DeepCopyInto(*out)
	}
	if in.InstancesDistribution != nil {
		in, out := &in.InstancesDistribution, &out.InstancesDistribution
		*out = new(NodeGroupInstancesDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.ASGMetricsCollection != nil {
		in, out := &in.ASGMetricsCollection, &out.ASGMetricsCollection
		*out = make([]MetricsCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CPUCredits != nil {
		in, out := &in.CPUCredits, &out.CPUCredits
		*out = new(string)
		**out = **in
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClassicLoadBalancerNames != nil {
		in, out := &in.ClassicLoadBalancerNames, &out.ClassicLoadBalancerNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TargetGroupARNs != nil {
		in, out := &in.TargetGroupARNs, &out.TargetGroupARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Bottlerocket != nil {
		in, out := &in.Bottlerocket, &out.Bottlerocket
		*out = new(NodeGroupBottlerocket)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeletExtraConfig != nil {
		in, out := &in.KubeletExtraConfig, &out.KubeletExtraConfig
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupBase) DeepCopyInto(out *NodeGroupBase) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScalingConfig != nil {
		in, out := &in.ScalingConfig, &out.ScalingConfig
		*out = new(ScalingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int)
		**out = **in
	}
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		*out = new(NodeGroupSSH)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IAM != nil {
		in, out := &in.IAM, &out.IAM
		*out = new(NodeGroupIAM)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = new(NodeGroupSGs)
		(*in).DeepCopyInto(*out)
	}
	if in.ASGSuspendProcesses != nil {
		in, out := &in.ASGSuspendProcesses, &out.ASGSuspendProcesses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EBSOptimized != nil {
		in, out := &in.EBSOptimized, &out.EBSOptimized
		*out = new(bool)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
	if in.VolumeEncrypted != nil {
		in, out := &in.VolumeEncrypted, &out.VolumeEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.VolumeKmsKeyID != nil {
		in, out := &in.VolumeKmsKeyID, &out.VolumeKmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.VolumeIOPS != nil {
		in, out := &in.VolumeIOPS, &out.VolumeIOPS
		*out = new(int)
		**out = **in
	}
	if in.VolumeThroughput != nil {
		in, out := &in.VolumeThroughput, &out.VolumeThroughput
		*out = new(int)
		**out = **in
	}
	if in.PreBootstrapCommands != nil {
		in, out := &in.PreBootstrapCommands, &out.PreBootstrapCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OverrideBootstrapCommand != nil {
		in, out := &in.OverrideBootstrapCommand, &out.OverrideBootstrapCommand
		*out = new(string)
		**out = **in
	}
	if in.DisableIMDSv1 != nil {
		in, out := &in.DisableIMDSv1, &out.DisableIMDSv1
		*out = new(bool)
		**out = **in
	}
	if in.DisablePodIMDS != nil {
		in, out := &in.DisablePodIMDS, &out.DisablePodIMDS
		*out = new(bool)
		**out = **in
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(Placement)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupBase.
func (in *NodeGroupBase) DeepCopy() *NodeGroupBase {
	if in == nil {
		return nil
	}
	out := new(NodeGroupBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupBottlerocket) DeepCopyInto(out *NodeGroupBottlerocket) {
	*out = *in
	if in.EnableAdminContainer != nil {
		in, out := &in.EnableAdminContainer, &out.EnableAdminContainer
		*out = new(bool)
		**out = **in
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupBottlerocket.
func (in *NodeGroupBottlerocket) DeepCopy() *NodeGroupBottlerocket {
	if in == nil {
		return nil
	}
	out := new(NodeGroupBottlerocket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupIAM) DeepCopyInto(out *NodeGroupIAM) {
	*out = *in
	if in.AttachPolicyARNs != nil {
		in, out := &in.AttachPolicyARNs, &out.AttachPolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.WithAddonPolicies.DeepCopyInto(&out.WithAddonPolicies)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupIAM.
func (in *NodeGroupIAM) DeepCopy() *NodeGroupIAM {
	if in == nil {
		return nil
	}
	out := new(NodeGroupIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupIAMAddonPolicies) DeepCopyInto(out *NodeGroupIAMAddonPolicies) {
	*out = *in
	if in.ImageBuilder != nil {
		in, out := &in.ImageBuilder, &out.ImageBuilder
		*out = new(bool)
		**out = **in
	}
	if in.AutoScaler != nil {
		in, out := &in.AutoScaler, &out.AutoScaler
		*out = new(bool)
		**out = **in
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = new(bool)
		**out = **in
	}
	if in.CertManager != nil {
		in, out := &in.CertManager, &out.CertManager
		*out = new(bool)
		**out = **in
	}
	if in.AppMesh != nil {
		in, out := &in.AppMesh, &out.AppMesh
		*out = new(bool)
		**out = **in
	}
	if in.AppMeshPreview != nil {
		in, out := &in.AppMeshPreview, &out.AppMeshPreview
		*out = new(bool)
		**out = **in
	}
	if in.EBS != nil {
		in, out := &in.EBS, &out.EBS
		*out = new(bool)
		**out = **in
	}
	if in.FSX != nil {
		in, out := &in.FSX, &out.FSX
		*out = new(bool)
		**out = **in
	}
	if in.EFS != nil {
		in, out := &in.EFS, &out.EFS
		*out = new(bool)
		**out = **in
	}
	if in.AWSLoadBalancerController != nil {
		in, out := &in.AWSLoadBalancerController, &out.AWSLoadBalancerController
		*out = new(bool)
		**out = **in
	}
	if in.XRay != nil {
		in, out := &in.XRay, &out.XRay
		*out = new(bool)
		**out = **in
	}
	if in.CloudWatch != nil {
		in, out := &in.CloudWatch, &out.CloudWatch
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupIAMAddonPolicies.
func (in *NodeGroupIAMAddonPolicies) DeepCopy() *NodeGroupIAMAddonPolicies {
	if in == nil {
		return nil
	}
	out := new(NodeGroupIAMAddonPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupInstancesDistribution) DeepCopyInto(out *NodeGroupInstancesDistribution) {
	*out = *in
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		*out = new(float64)
		**out = **in
	}
	if in.OnDemandBaseCapacity != nil {
		in, out := &in.OnDemandBaseCapacity, &out.OnDemandBaseCapacity
		*out = new(int)
		**out = **in
	}
	if in.OnDemandPercentageAboveBaseCapacity != nil {
		in, out := &in.OnDemandPercentageAboveBaseCapacity, &out.OnDemandPercentageAboveBaseCapacity
		*out = new(int)
		**out = **in
	}
	if in.SpotInstancePools != nil {
		in, out := &in.SpotInstancePools, &out.SpotInstancePools
		*out = new(int)
		**out = **in
	}
	if in.SpotAllocationStrategy != nil {
		in, out := &in.SpotAllocationStrategy, &out.SpotAllocationStrategy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupInstancesDistribution.
func (in *NodeGroupInstancesDistribution) DeepCopy() *NodeGroupInstancesDistribution {
	if in == nil {
		return nil
	}
	out := new(NodeGroupInstancesDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSGs) DeepCopyInto(out *NodeGroupSGs) {
	*out = *in
	if in.AttachIDs != nil {
		in, out := &in.AttachIDs, &out.AttachIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WithShared != nil {
		in, out := &in.WithShared, &out.WithShared
		*out = new(bool)
		**out = **in
	}
	if in.WithLocal != nil {
		in, out := &in.WithLocal, &out.WithLocal
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSGs.
func (in *NodeGroupSGs) DeepCopy() *NodeGroupSGs {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSGs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSSH) DeepCopyInto(out *NodeGroupSSH) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = new(bool)
		**out = **in
	}
	if in.PublicKeyPath != nil {
		in, out := &in.PublicKeyPath, &out.PublicKeyPath
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyName != nil {
		in, out := &in.PublicKeyName, &out.PublicKeyName
		*out = new(string)
		**out = **in
	}
	if in.SourceSecurityGroupIDs != nil {
		in, out := &in.SourceSecurityGroupIDs, &out.SourceSecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnableSSM != nil {
		in, out := &in.EnableSSM, &out.EnableSSM
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSSH.
func (in *NodeGroupSSH) DeepCopy() *NodeGroupSSH {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operator) DeepCopyInto(out *Operator) {
	*out = *in
	if in.CommitOperatorManifests != nil {
		in, out := &in.CommitOperatorManifests, &out.CommitOperatorManifests
		*out = new(bool)
		**out = **in
	}
	if in.WithHelm != nil {
		in, out := &in.WithHelm, &out.WithHelm
		*out = new(bool)
		**out = **in
	}
	if in.AdditionalFluxArgs != nil {
		in, out := &in.AdditionalFluxArgs, &out.AdditionalFluxArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalHelmOperatorArgs != nil {
		in, out := &in.AdditionalHelmOperatorArgs, &out.AdditionalHelmOperatorArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operator.
func (in *Operator) DeepCopy() *Operator {
	if in == nil {
		return nil
	}
	out := new(Operator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCluster) DeepCopyInto(out *PrivateCluster) {
	*out = *in
	if in.AdditionalEndpointServices != nil {
		in, out := &in.AdditionalEndpointServices, &out.AdditionalEndpointServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCluster.
func (in *PrivateCluster) DeepCopy() *PrivateCluster {
	if in == nil {
		return nil
	}
	out := new(PrivateCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repo) DeepCopyInto(out *Repo) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repo.
func (in *Repo) DeepCopy() *Repo {
	if in == nil {
		return nil
	}
	out := new(Repo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingConfig) DeepCopyInto(out *ScalingConfig) {
	*out = *in
	if in.DesiredCapacity != nil {
		in, out := &in.DesiredCapacity, &out.DesiredCapacity
		*out = new(int)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingConfig.
func (in *ScalingConfig) DeepCopy() *ScalingConfig {
	if in == nil {
		return nil
	}
	out := new(ScalingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsEncryption) DeepCopyInto(out *SecretsEncryption) {
	*out = *in
	if in.KeyARN != nil {
		in, out := &in.KeyARN, &out.KeyARN
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsEncryption.
func (in *SecretsEncryption) DeepCopy() *SecretsEncryption {
	if in == nil {
		return nil
	}
	out := new(SecretsEncryption)
	in.DeepCopyInto(out)
	return out
}
