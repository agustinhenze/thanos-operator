// +build !ignore_autogenerated

// Copyright 2020 Banzai Cloud
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/banzaicloud/operator-tools/pkg/types"
	"github.com/banzaicloud/operator-tools/pkg/volume"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketWeb) DeepCopyInto(out *BucketWeb) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentOverrides != nil {
		in, out := &in.DeploymentOverrides, &out.DeploymentOverrides
		*out = new(types.DeploymentSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		**out = **in
	}
	out.HTTPGracePeriod = in.HTTPGracePeriod
	out.Refresh = in.Refresh
	out.Timeout = in.Timeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketWeb.
func (in *BucketWeb) DeepCopy() *BucketWeb {
	if in == nil {
		return nil
	}
	out := new(BucketWeb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Compactor) DeepCopyInto(out *Compactor) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentOverrides != nil {
		in, out := &in.DeploymentOverrides, &out.DeploymentOverrides
		*out = new(types.DeploymentSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		**out = **in
	}
	out.HTTPGracePeriod = in.HTTPGracePeriod
	if in.DataVolume != nil {
		in, out := &in.DataVolume, &out.DataVolume
		*out = new(volume.KubernetesVolume)
		(*in).DeepCopyInto(*out)
	}
	out.ConsistencyDelay = in.ConsistencyDelay
	out.RetentionResolutionRaw = in.RetentionResolutionRaw
	out.RetentionResolution5m = in.RetentionResolution5m
	out.RetentionResolution1h = in.RetentionResolution1h
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Compactor.
func (in *Compactor) DeepCopy() *Compactor {
	if in == nil {
		return nil
	}
	out := new(Compactor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSelector) DeepCopyInto(out *KubernetesSelector) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSelector.
func (in *KubernetesSelector) DeepCopy() *KubernetesSelector {
	if in == nil {
		return nil
	}
	out := new(KubernetesSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metrics) DeepCopyInto(out *Metrics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metrics.
func (in *Metrics) DeepCopy() *Metrics {
	if in == nil {
		return nil
	}
	out := new(Metrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStore) DeepCopyInto(out *ObjectStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStore.
func (in *ObjectStore) DeepCopy() *ObjectStore {
	if in == nil {
		return nil
	}
	out := new(ObjectStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreList) DeepCopyInto(out *ObjectStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreList.
func (in *ObjectStoreList) DeepCopy() *ObjectStoreList {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreSpec) DeepCopyInto(out *ObjectStoreSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.Compactor != nil {
		in, out := &in.Compactor, &out.Compactor
		*out = new(Compactor)
		(*in).DeepCopyInto(*out)
	}
	if in.BucketWeb != nil {
		in, out := &in.BucketWeb, &out.BucketWeb
		*out = new(BucketWeb)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreSpec.
func (in *ObjectStoreSpec) DeepCopy() *ObjectStoreSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreStatus) DeepCopyInto(out *ObjectStoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreStatus.
func (in *ObjectStoreStatus) DeepCopy() *ObjectStoreStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Query) DeepCopyInto(out *Query) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentOverrides != nil {
		in, out := &in.DeploymentOverrides, &out.DeploymentOverrides
		*out = new(types.DeploymentSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		**out = **in
	}
	if in.HTTPIngress != nil {
		in, out := &in.HTTPIngress, &out.HTTPIngress
		*out = new(Ingress)
		**out = **in
	}
	if in.GRPCIngress != nil {
		in, out := &in.GRPCIngress, &out.GRPCIngress
		*out = new(Ingress)
		**out = **in
	}
	out.QueryTimeout = in.QueryTimeout
	if in.QueryReplicaLabels != nil {
		in, out := &in.QueryReplicaLabels, &out.QueryReplicaLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SelectorLabels != nil {
		in, out := &in.SelectorLabels, &out.SelectorLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Stores != nil {
		in, out := &in.Stores, &out.Stores
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.StoreSDDNSInterval = in.StoreSDDNSInterval
	out.StoreUnhealthyTimeout = in.StoreUnhealthyTimeout
	out.QueryDefaultEvaluationInterval = in.QueryDefaultEvaluationInterval
	out.StoreResponseTimeout = in.StoreResponseTimeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Query.
func (in *Query) DeepCopy() *Query {
	if in == nil {
		return nil
	}
	out := new(Query)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryFrontend) DeepCopyInto(out *QueryFrontend) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentOverrides != nil {
		in, out := &in.DeploymentOverrides, &out.DeploymentOverrides
		*out = new(types.DeploymentSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		**out = **in
	}
	if in.HTTPIngress != nil {
		in, out := &in.HTTPIngress, &out.HTTPIngress
		*out = new(Ingress)
		**out = **in
	}
	if in.QueryRangePartialResponse != nil {
		in, out := &in.QueryRangePartialResponse, &out.QueryRangePartialResponse
		*out = new(bool)
		**out = **in
	}
	if in.QueryFrontendCompressResponses != nil {
		in, out := &in.QueryFrontendCompressResponses, &out.QueryFrontendCompressResponses
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryFrontend.
func (in *QueryFrontend) DeepCopy() *QueryFrontend {
	if in == nil {
		return nil
	}
	out := new(QueryFrontend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
	if in.StatefulsetOverrides != nil {
		in, out := &in.StatefulsetOverrides, &out.StatefulsetOverrides
		*out = new(types.StatefulsetSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		**out = **in
	}
	if in.HTTPIngress != nil {
		in, out := &in.HTTPIngress, &out.HTTPIngress
		*out = new(Ingress)
		**out = **in
	}
	if in.GRPCIngress != nil {
		in, out := &in.GRPCIngress, &out.GRPCIngress
		*out = new(Ingress)
		**out = **in
	}
	if in.DataVolume != nil {
		in, out := &in.DataVolume, &out.DataVolume
		*out = new(volume.KubernetesVolume)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AlertmanagersURLs != nil {
		in, out := &in.AlertmanagersURLs, &out.AlertmanagersURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AlertLabelDrop != nil {
		in, out := &in.AlertLabelDrop, &out.AlertLabelDrop
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Queries != nil {
		in, out := &in.Queries, &out.Queries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreEndpoint) DeepCopyInto(out *StoreEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreEndpoint.
func (in *StoreEndpoint) DeepCopy() *StoreEndpoint {
	if in == nil {
		return nil
	}
	out := new(StoreEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreEndpointList) DeepCopyInto(out *StoreEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoreEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreEndpointList.
func (in *StoreEndpointList) DeepCopy() *StoreEndpointList {
	if in == nil {
		return nil
	}
	out := new(StoreEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreEndpointSpec) DeepCopyInto(out *StoreEndpointSpec) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(KubernetesSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(Ingress)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreEndpointSpec.
func (in *StoreEndpointSpec) DeepCopy() *StoreEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(StoreEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreEndpointStatus) DeepCopyInto(out *StoreEndpointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreEndpointStatus.
func (in *StoreEndpointStatus) DeepCopy() *StoreEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(StoreEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreGateway) DeepCopyInto(out *StoreGateway) {
	*out = *in
	if in.MetaOverrides != nil {
		in, out := &in.MetaOverrides, &out.MetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadMetaOverrides != nil {
		in, out := &in.WorkloadMetaOverrides, &out.WorkloadMetaOverrides
		*out = new(types.MetaBase)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadOverrides != nil {
		in, out := &in.WorkloadOverrides, &out.WorkloadOverrides
		*out = new(types.PodSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerOverrides != nil {
		in, out := &in.ContainerOverrides, &out.ContainerOverrides
		*out = new(types.ContainerBase)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentOverrides != nil {
		in, out := &in.DeploymentOverrides, &out.DeploymentOverrides
		*out = new(types.DeploymentSpecBase)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		**out = **in
	}
	if in.TimeRanges != nil {
		in, out := &in.TimeRanges, &out.TimeRanges
		*out = make([]TimeRange, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreGateway.
func (in *StoreGateway) DeepCopy() *StoreGateway {
	if in == nil {
		return nil
	}
	out := new(StoreGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Thanos) DeepCopyInto(out *Thanos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Thanos.
func (in *Thanos) DeepCopy() *Thanos {
	if in == nil {
		return nil
	}
	out := new(Thanos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Thanos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosDiscovery) DeepCopyInto(out *ThanosDiscovery) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosDiscovery.
func (in *ThanosDiscovery) DeepCopy() *ThanosDiscovery {
	if in == nil {
		return nil
	}
	out := new(ThanosDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosList) DeepCopyInto(out *ThanosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Thanos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosList.
func (in *ThanosList) DeepCopy() *ThanosList {
	if in == nil {
		return nil
	}
	out := new(ThanosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosSpec) DeepCopyInto(out *ThanosSpec) {
	*out = *in
	if in.StoreGateway != nil {
		in, out := &in.StoreGateway, &out.StoreGateway
		*out = new(StoreGateway)
		(*in).DeepCopyInto(*out)
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(Rule)
		(*in).DeepCopyInto(*out)
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(Query)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryFrontend != nil {
		in, out := &in.QueryFrontend, &out.QueryFrontend
		*out = new(QueryFrontend)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosSpec.
func (in *ThanosSpec) DeepCopy() *ThanosSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosStatus) DeepCopyInto(out *ThanosStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosStatus.
func (in *ThanosStatus) DeepCopy() *ThanosStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRange) DeepCopyInto(out *TimeRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRange.
func (in *TimeRange) DeepCopy() *TimeRange {
	if in == nil {
		return nil
	}
	out := new(TimeRange)
	in.DeepCopyInto(out)
	return out
}
