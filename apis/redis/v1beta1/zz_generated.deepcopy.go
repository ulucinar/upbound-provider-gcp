//go:build (redis || all) && !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceInitParameters) DeepCopyInto(out *InstanceInitParameters) {
	*out = *in
	if in.AlternativeLocationID != nil {
		in, out := &in.AlternativeLocationID, &out.AlternativeLocationID
		*out = new(string)
		**out = **in
	}
	if in.AuthEnabled != nil {
		in, out := &in.AuthEnabled, &out.AuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizedNetwork != nil {
		in, out := &in.AuthorizedNetwork, &out.AuthorizedNetwork
		*out = new(string)
		**out = **in
	}
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKeyRef != nil {
		in, out := &in.CustomerManagedKeyRef, &out.CustomerManagedKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomerManagedKeySelector != nil {
		in, out := &in.CustomerManagedKeySelector, &out.CustomerManagedKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LocationID != nil {
		in, out := &in.LocationID, &out.LocationID
		*out = new(string)
		**out = **in
	}
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = make([]MaintenancePolicyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemorySizeGb != nil {
		in, out := &in.MemorySizeGb, &out.MemorySizeGb
		*out = new(float64)
		**out = **in
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = make([]PersistenceConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ReadReplicasMode != nil {
		in, out := &in.ReadReplicasMode, &out.ReadReplicasMode
		*out = new(string)
		**out = **in
	}
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(float64)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
	if in.SecondaryIPRange != nil {
		in, out := &in.SecondaryIPRange, &out.SecondaryIPRange
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceInitParameters.
func (in *InstanceInitParameters) DeepCopy() *InstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
	if in.AlternativeLocationID != nil {
		in, out := &in.AlternativeLocationID, &out.AlternativeLocationID
		*out = new(string)
		**out = **in
	}
	if in.AuthEnabled != nil {
		in, out := &in.AuthEnabled, &out.AuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizedNetwork != nil {
		in, out := &in.AuthorizedNetwork, &out.AuthorizedNetwork
		*out = new(string)
		**out = **in
	}
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.CurrentLocationID != nil {
		in, out := &in.CurrentLocationID, &out.CurrentLocationID
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LocationID != nil {
		in, out := &in.LocationID, &out.LocationID
		*out = new(string)
		**out = **in
	}
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = make([]MaintenancePolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaintenanceSchedule != nil {
		in, out := &in.MaintenanceSchedule, &out.MaintenanceSchedule
		*out = make([]MaintenanceScheduleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemorySizeGb != nil {
		in, out := &in.MemorySizeGb, &out.MemorySizeGb
		*out = new(float64)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = make([]PersistenceConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistenceIAMIdentity != nil {
		in, out := &in.PersistenceIAMIdentity, &out.PersistenceIAMIdentity
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ReadEndpoint != nil {
		in, out := &in.ReadEndpoint, &out.ReadEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ReadEndpointPort != nil {
		in, out := &in.ReadEndpointPort, &out.ReadEndpointPort
		*out = new(float64)
		**out = **in
	}
	if in.ReadReplicasMode != nil {
		in, out := &in.ReadReplicasMode, &out.ReadReplicasMode
		*out = new(string)
		**out = **in
	}
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(float64)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
	if in.SecondaryIPRange != nil {
		in, out := &in.SecondaryIPRange, &out.SecondaryIPRange
		*out = new(string)
		**out = **in
	}
	if in.ServerCACerts != nil {
		in, out := &in.ServerCACerts, &out.ServerCACerts
		*out = make([]ServerCACertsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.AlternativeLocationID != nil {
		in, out := &in.AlternativeLocationID, &out.AlternativeLocationID
		*out = new(string)
		**out = **in
	}
	if in.AuthEnabled != nil {
		in, out := &in.AuthEnabled, &out.AuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizedNetwork != nil {
		in, out := &in.AuthorizedNetwork, &out.AuthorizedNetwork
		*out = new(string)
		**out = **in
	}
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKeyRef != nil {
		in, out := &in.CustomerManagedKeyRef, &out.CustomerManagedKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomerManagedKeySelector != nil {
		in, out := &in.CustomerManagedKeySelector, &out.CustomerManagedKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LocationID != nil {
		in, out := &in.LocationID, &out.LocationID
		*out = new(string)
		**out = **in
	}
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = make([]MaintenancePolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemorySizeGb != nil {
		in, out := &in.MemorySizeGb, &out.MemorySizeGb
		*out = new(float64)
		**out = **in
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = make([]PersistenceConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ReadReplicasMode != nil {
		in, out := &in.ReadReplicasMode, &out.ReadReplicasMode
		*out = new(string)
		**out = **in
	}
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(float64)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
	if in.SecondaryIPRange != nil {
		in, out := &in.SecondaryIPRange, &out.SecondaryIPRange
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenancePolicyInitParameters) DeepCopyInto(out *MaintenancePolicyInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]WeeklyMaintenanceWindowInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenancePolicyInitParameters.
func (in *MaintenancePolicyInitParameters) DeepCopy() *MaintenancePolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenancePolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenancePolicyObservation) DeepCopyInto(out *MaintenancePolicyObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]WeeklyMaintenanceWindowObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenancePolicyObservation.
func (in *MaintenancePolicyObservation) DeepCopy() *MaintenancePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenancePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenancePolicyParameters) DeepCopyInto(out *MaintenancePolicyParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]WeeklyMaintenanceWindowParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenancePolicyParameters.
func (in *MaintenancePolicyParameters) DeepCopy() *MaintenancePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenancePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceScheduleInitParameters) DeepCopyInto(out *MaintenanceScheduleInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceScheduleInitParameters.
func (in *MaintenanceScheduleInitParameters) DeepCopy() *MaintenanceScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceScheduleObservation) DeepCopyInto(out *MaintenanceScheduleObservation) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ScheduleDeadlineTime != nil {
		in, out := &in.ScheduleDeadlineTime, &out.ScheduleDeadlineTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceScheduleObservation.
func (in *MaintenanceScheduleObservation) DeepCopy() *MaintenanceScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceScheduleParameters) DeepCopyInto(out *MaintenanceScheduleParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceScheduleParameters.
func (in *MaintenanceScheduleParameters) DeepCopy() *MaintenanceScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesInitParameters) DeepCopyInto(out *NodesInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesInitParameters.
func (in *NodesInitParameters) DeepCopy() *NodesInitParameters {
	if in == nil {
		return nil
	}
	out := new(NodesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesObservation) DeepCopyInto(out *NodesObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesObservation.
func (in *NodesObservation) DeepCopy() *NodesObservation {
	if in == nil {
		return nil
	}
	out := new(NodesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesParameters) DeepCopyInto(out *NodesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesParameters.
func (in *NodesParameters) DeepCopy() *NodesParameters {
	if in == nil {
		return nil
	}
	out := new(NodesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfigInitParameters) DeepCopyInto(out *PersistenceConfigInitParameters) {
	*out = *in
	if in.PersistenceMode != nil {
		in, out := &in.PersistenceMode, &out.PersistenceMode
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotPeriod != nil {
		in, out := &in.RdbSnapshotPeriod, &out.RdbSnapshotPeriod
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotStartTime != nil {
		in, out := &in.RdbSnapshotStartTime, &out.RdbSnapshotStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfigInitParameters.
func (in *PersistenceConfigInitParameters) DeepCopy() *PersistenceConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfigObservation) DeepCopyInto(out *PersistenceConfigObservation) {
	*out = *in
	if in.PersistenceMode != nil {
		in, out := &in.PersistenceMode, &out.PersistenceMode
		*out = new(string)
		**out = **in
	}
	if in.RdbNextSnapshotTime != nil {
		in, out := &in.RdbNextSnapshotTime, &out.RdbNextSnapshotTime
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotPeriod != nil {
		in, out := &in.RdbSnapshotPeriod, &out.RdbSnapshotPeriod
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotStartTime != nil {
		in, out := &in.RdbSnapshotStartTime, &out.RdbSnapshotStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfigObservation.
func (in *PersistenceConfigObservation) DeepCopy() *PersistenceConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfigParameters) DeepCopyInto(out *PersistenceConfigParameters) {
	*out = *in
	if in.PersistenceMode != nil {
		in, out := &in.PersistenceMode, &out.PersistenceMode
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotPeriod != nil {
		in, out := &in.RdbSnapshotPeriod, &out.RdbSnapshotPeriod
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotStartTime != nil {
		in, out := &in.RdbSnapshotStartTime, &out.RdbSnapshotStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfigParameters.
func (in *PersistenceConfigParameters) DeepCopy() *PersistenceConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCACertsInitParameters) DeepCopyInto(out *ServerCACertsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCACertsInitParameters.
func (in *ServerCACertsInitParameters) DeepCopy() *ServerCACertsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCACertsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCACertsObservation) DeepCopyInto(out *ServerCACertsObservation) {
	*out = *in
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ExpireTime != nil {
		in, out := &in.ExpireTime, &out.ExpireTime
		*out = new(string)
		**out = **in
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.Sha1Fingerprint != nil {
		in, out := &in.Sha1Fingerprint, &out.Sha1Fingerprint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCACertsObservation.
func (in *ServerCACertsObservation) DeepCopy() *ServerCACertsObservation {
	if in == nil {
		return nil
	}
	out := new(ServerCACertsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCACertsParameters) DeepCopyInto(out *ServerCACertsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCACertsParameters.
func (in *ServerCACertsParameters) DeepCopy() *ServerCACertsParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCACertsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeInitParameters) DeepCopyInto(out *StartTimeInitParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(float64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeInitParameters.
func (in *StartTimeInitParameters) DeepCopy() *StartTimeInitParameters {
	if in == nil {
		return nil
	}
	out := new(StartTimeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeObservation) DeepCopyInto(out *StartTimeObservation) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(float64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeObservation.
func (in *StartTimeObservation) DeepCopy() *StartTimeObservation {
	if in == nil {
		return nil
	}
	out := new(StartTimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeParameters) DeepCopyInto(out *StartTimeParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(float64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeParameters.
func (in *StartTimeParameters) DeepCopy() *StartTimeParameters {
	if in == nil {
		return nil
	}
	out := new(StartTimeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyMaintenanceWindowInitParameters) DeepCopyInto(out *WeeklyMaintenanceWindowInitParameters) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = make([]StartTimeInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyMaintenanceWindowInitParameters.
func (in *WeeklyMaintenanceWindowInitParameters) DeepCopy() *WeeklyMaintenanceWindowInitParameters {
	if in == nil {
		return nil
	}
	out := new(WeeklyMaintenanceWindowInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyMaintenanceWindowObservation) DeepCopyInto(out *WeeklyMaintenanceWindowObservation) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = make([]StartTimeObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyMaintenanceWindowObservation.
func (in *WeeklyMaintenanceWindowObservation) DeepCopy() *WeeklyMaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(WeeklyMaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyMaintenanceWindowParameters) DeepCopyInto(out *WeeklyMaintenanceWindowParameters) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = make([]StartTimeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyMaintenanceWindowParameters.
func (in *WeeklyMaintenanceWindowParameters) DeepCopy() *WeeklyMaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(WeeklyMaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}
