// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistry) DeepCopyInto(out *ApicurioRegistry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistry.
func (in *ApicurioRegistry) DeepCopy() *ApicurioRegistry {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApicurioRegistry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistryList) DeepCopyInto(out *ApicurioRegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApicurioRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistryList.
func (in *ApicurioRegistryList) DeepCopy() *ApicurioRegistryList {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApicurioRegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpec) DeepCopyInto(out *ApicurioRegistrySpec) {
	*out = *in
	out.Configuration = in.Configuration
	in.Deployment.DeepCopyInto(&out.Deployment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpec.
func (in *ApicurioRegistrySpec) DeepCopy() *ApicurioRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfiguration) DeepCopyInto(out *ApicurioRegistrySpecConfiguration) {
	*out = *in
	out.Sql = in.Sql
	out.Kafkasql = in.Kafkasql
	out.UI = in.UI
	out.Security = in.Security
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfiguration.
func (in *ApicurioRegistrySpecConfiguration) DeepCopy() *ApicurioRegistrySpecConfiguration {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationDataSource) DeepCopyInto(out *ApicurioRegistrySpecConfigurationDataSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationDataSource.
func (in *ApicurioRegistrySpecConfigurationDataSource) DeepCopy() *ApicurioRegistrySpecConfigurationDataSource {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationKafkaSecurity) DeepCopyInto(out *ApicurioRegistrySpecConfigurationKafkaSecurity) {
	*out = *in
	out.Tls = in.Tls
	out.Scram = in.Scram
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationKafkaSecurity.
func (in *ApicurioRegistrySpecConfigurationKafkaSecurity) DeepCopy() *ApicurioRegistrySpecConfigurationKafkaSecurity {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationKafkaSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationKafkaSecurityScram) DeepCopyInto(out *ApicurioRegistrySpecConfigurationKafkaSecurityScram) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationKafkaSecurityScram.
func (in *ApicurioRegistrySpecConfigurationKafkaSecurityScram) DeepCopy() *ApicurioRegistrySpecConfigurationKafkaSecurityScram {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationKafkaSecurityScram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationKafkaSecurityTls) DeepCopyInto(out *ApicurioRegistrySpecConfigurationKafkaSecurityTls) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationKafkaSecurityTls.
func (in *ApicurioRegistrySpecConfigurationKafkaSecurityTls) DeepCopy() *ApicurioRegistrySpecConfigurationKafkaSecurityTls {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationKafkaSecurityTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationKafkasql) DeepCopyInto(out *ApicurioRegistrySpecConfigurationKafkasql) {
	*out = *in
	out.Security = in.Security
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationKafkasql.
func (in *ApicurioRegistrySpecConfigurationKafkasql) DeepCopy() *ApicurioRegistrySpecConfigurationKafkasql {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationKafkasql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationSecurity) DeepCopyInto(out *ApicurioRegistrySpecConfigurationSecurity) {
	*out = *in
	out.Keycloak = in.Keycloak
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationSecurity.
func (in *ApicurioRegistrySpecConfigurationSecurity) DeepCopy() *ApicurioRegistrySpecConfigurationSecurity {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationSecurityKeycloak) DeepCopyInto(out *ApicurioRegistrySpecConfigurationSecurityKeycloak) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationSecurityKeycloak.
func (in *ApicurioRegistrySpecConfigurationSecurityKeycloak) DeepCopy() *ApicurioRegistrySpecConfigurationSecurityKeycloak {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationSecurityKeycloak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationSql) DeepCopyInto(out *ApicurioRegistrySpecConfigurationSql) {
	*out = *in
	out.DataSource = in.DataSource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationSql.
func (in *ApicurioRegistrySpecConfigurationSql) DeepCopy() *ApicurioRegistrySpecConfigurationSql {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationSql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationUI) DeepCopyInto(out *ApicurioRegistrySpecConfigurationUI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationUI.
func (in *ApicurioRegistrySpecConfigurationUI) DeepCopy() *ApicurioRegistrySpecConfigurationUI {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecDeployment) DeepCopyInto(out *ApicurioRegistrySpecDeployment) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecDeployment.
func (in *ApicurioRegistrySpecDeployment) DeepCopy() *ApicurioRegistrySpecDeployment {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecDeploymentMetadata) DeepCopyInto(out *ApicurioRegistrySpecDeploymentMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecDeploymentMetadata.
func (in *ApicurioRegistrySpecDeploymentMetadata) DeepCopy() *ApicurioRegistrySpecDeploymentMetadata {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecDeploymentMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistryStatus) DeepCopyInto(out *ApicurioRegistryStatus) {
	*out = *in
	out.Info = in.Info
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make([]ApicurioRegistryStatusManagedResource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistryStatus.
func (in *ApicurioRegistryStatus) DeepCopy() *ApicurioRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistryStatusInfo) DeepCopyInto(out *ApicurioRegistryStatusInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistryStatusInfo.
func (in *ApicurioRegistryStatusInfo) DeepCopy() *ApicurioRegistryStatusInfo {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistryStatusInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistryStatusManagedResource) DeepCopyInto(out *ApicurioRegistryStatusManagedResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistryStatusManagedResource.
func (in *ApicurioRegistryStatusManagedResource) DeepCopy() *ApicurioRegistryStatusManagedResource {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistryStatusManagedResource)
	in.DeepCopyInto(out)
	return out
}
