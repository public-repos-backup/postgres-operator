// +build !ignore_autogenerated

/*
Copyright 2020 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSGCPConfiguration) DeepCopyInto(out *AWSGCPConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSGCPConfiguration.
func (in *AWSGCPConfiguration) DeepCopy() *AWSGCPConfiguration {
	if in == nil {
		return nil
	}
	out := new(AWSGCPConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneDescription) DeepCopyInto(out *CloneDescription) {
	*out = *in
	if in.S3ForcePathStyle != nil {
		in, out := &in.S3ForcePathStyle, &out.S3ForcePathStyle
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneDescription.
func (in *CloneDescription) DeepCopy() *CloneDescription {
	if in == nil {
		return nil
	}
	out := new(CloneDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesMetaConfiguration) DeepCopyInto(out *KubernetesMetaConfiguration) {
	*out = *in
	if in.SpiloFSGroup != nil {
		in, out := &in.SpiloFSGroup, &out.SpiloFSGroup
		*out = new(int64)
		**out = **in
	}
	if in.EnablePodDisruptionBudget != nil {
		in, out := &in.EnablePodDisruptionBudget, &out.EnablePodDisruptionBudget
		*out = new(bool)
		**out = **in
	}
	if in.EnableInitContainers != nil {
		in, out := &in.EnableInitContainers, &out.EnableInitContainers
		*out = new(bool)
		**out = **in
	}
	if in.EnableSidecars != nil {
		in, out := &in.EnableSidecars, &out.EnableSidecars
		*out = new(bool)
		**out = **in
	}
	out.OAuthTokenSecretName = in.OAuthTokenSecretName
	out.InfrastructureRolesSecretName = in.InfrastructureRolesSecretName
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InheritedLabels != nil {
		in, out := &in.InheritedLabels, &out.InheritedLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeReadinessLabel != nil {
		in, out := &in.NodeReadinessLabel, &out.NodeReadinessLabel
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CustomPodAnnotations != nil {
		in, out := &in.CustomPodAnnotations, &out.CustomPodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodToleration != nil {
		in, out := &in.PodToleration, &out.PodToleration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesMetaConfiguration.
func (in *KubernetesMetaConfiguration) DeepCopy() *KubernetesMetaConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubernetesMetaConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfiguration) DeepCopyInto(out *LoadBalancerConfiguration) {
	*out = *in
	if in.CustomServiceAnnotations != nil {
		in, out := &in.CustomServiceAnnotations, &out.CustomServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfiguration.
func (in *LoadBalancerConfiguration) DeepCopy() *LoadBalancerConfiguration {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingRESTAPIConfiguration) DeepCopyInto(out *LoggingRESTAPIConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingRESTAPIConfiguration.
func (in *LoggingRESTAPIConfiguration) DeepCopy() *LoggingRESTAPIConfiguration {
	if in == nil {
		return nil
	}
	out := new(LoggingRESTAPIConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow) DeepCopyInto(out *MaintenanceWindow) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow.
func (in *MaintenanceWindow) DeepCopy() *MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfiguration) DeepCopyInto(out *OperatorConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Configuration.DeepCopyInto(&out.Configuration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfiguration.
func (in *OperatorConfiguration) DeepCopy() *OperatorConfiguration {
	if in == nil {
		return nil
	}
	out := new(OperatorConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfigurationData) DeepCopyInto(out *OperatorConfigurationData) {
	*out = *in
	if in.EnableCRDValidation != nil {
		in, out := &in.EnableCRDValidation, &out.EnableCRDValidation
		*out = new(bool)
		**out = **in
	}
	if in.ShmVolume != nil {
		in, out := &in.ShmVolume, &out.ShmVolume
		*out = new(bool)
		**out = **in
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.PostgresUsersConfiguration = in.PostgresUsersConfiguration
	in.Kubernetes.DeepCopyInto(&out.Kubernetes)
	out.PostgresPodResources = in.PostgresPodResources
	out.Timeouts = in.Timeouts
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	out.AWSGCP = in.AWSGCP
	out.OperatorDebug = in.OperatorDebug
	in.TeamsAPI.DeepCopyInto(&out.TeamsAPI)
	out.LoggingRESTAPI = in.LoggingRESTAPI
	out.Scalyr = in.Scalyr
	out.LogicalBackup = in.LogicalBackup
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfigurationData.
func (in *OperatorConfigurationData) DeepCopy() *OperatorConfigurationData {
	if in == nil {
		return nil
	}
	out := new(OperatorConfigurationData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfigurationList) DeepCopyInto(out *OperatorConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfigurationList.
func (in *OperatorConfigurationList) DeepCopy() *OperatorConfigurationList {
	if in == nil {
		return nil
	}
	out := new(OperatorConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorDebugConfiguration) DeepCopyInto(out *OperatorDebugConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorDebugConfiguration.
func (in *OperatorDebugConfiguration) DeepCopy() *OperatorDebugConfiguration {
	if in == nil {
		return nil
	}
	out := new(OperatorDebugConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorLogicalBackupConfiguration) DeepCopyInto(out *OperatorLogicalBackupConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorLogicalBackupConfiguration.
func (in *OperatorLogicalBackupConfiguration) DeepCopy() *OperatorLogicalBackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(OperatorLogicalBackupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorTimeouts) DeepCopyInto(out *OperatorTimeouts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorTimeouts.
func (in *OperatorTimeouts) DeepCopy() *OperatorTimeouts {
	if in == nil {
		return nil
	}
	out := new(OperatorTimeouts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Patroni) DeepCopyInto(out *Patroni) {
	*out = *in
	if in.InitDB != nil {
		in, out := &in.InitDB, &out.InitDB
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PgHba != nil {
		in, out := &in.PgHba, &out.PgHba
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Slots != nil {
		in, out := &in.Slots, &out.Slots
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Patroni.
func (in *Patroni) DeepCopy() *Patroni {
	if in == nil {
		return nil
	}
	out := new(Patroni)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresPodResourcesDefaults) DeepCopyInto(out *PostgresPodResourcesDefaults) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresPodResourcesDefaults.
func (in *PostgresPodResourcesDefaults) DeepCopy() *PostgresPodResourcesDefaults {
	if in == nil {
		return nil
	}
	out := new(PostgresPodResourcesDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSpec) DeepCopyInto(out *PostgresSpec) {
	*out = *in
	in.PostgresqlParam.DeepCopyInto(&out.PostgresqlParam)
	out.Volume = in.Volume
	in.Patroni.DeepCopyInto(&out.Patroni)
	out.Resources = in.Resources
	if in.SpiloFSGroup != nil {
		in, out := &in.SpiloFSGroup, &out.SpiloFSGroup
		*out = new(int64)
		**out = **in
	}
	if in.EnableMasterLoadBalancer != nil {
		in, out := &in.EnableMasterLoadBalancer, &out.EnableMasterLoadBalancer
		*out = new(bool)
		**out = **in
	}
	if in.EnableReplicaLoadBalancer != nil {
		in, out := &in.EnableReplicaLoadBalancer, &out.EnableReplicaLoadBalancer
		*out = new(bool)
		**out = **in
	}
	if in.UseLoadBalancer != nil {
		in, out := &in.UseLoadBalancer, &out.UseLoadBalancer
		*out = new(bool)
		**out = **in
	}
	if in.ReplicaLoadBalancer != nil {
		in, out := &in.ReplicaLoadBalancer, &out.ReplicaLoadBalancer
		*out = new(bool)
		**out = **in
	}
	if in.AllowedSourceRanges != nil {
		in, out := &in.AllowedSourceRanges, &out.AllowedSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make(map[string]UserFlags, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(UserFlags, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.MaintenanceWindows != nil {
		in, out := &in.MaintenanceWindows, &out.MaintenanceWindows
		*out = make([]MaintenanceWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Clone.DeepCopyInto(&out.Clone)
	if in.Databases != nil {
		in, out := &in.Databases, &out.Databases
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PreparedDatabases != nil {
		in, out := &in.PreparedDatabases, &out.PreparedDatabases
		*out = make(map[string]PreparedDatabase, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]Sidecar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShmVolume != nil {
		in, out := &in.ShmVolume, &out.ShmVolume
		*out = new(bool)
		**out = **in
	}
	if in.StandbyCluster != nil {
		in, out := &in.StandbyCluster, &out.StandbyCluster
		*out = new(StandbyDescription)
		**out = **in
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSDescription)
		**out = **in
	}
	if in.InitContainersOld != nil {
		in, out := &in.InitContainersOld, &out.InitContainersOld
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSpec.
func (in *PostgresSpec) DeepCopy() *PostgresSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresStatus) DeepCopyInto(out *PostgresStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresStatus.
func (in *PostgresStatus) DeepCopy() *PostgresStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresUsersConfiguration) DeepCopyInto(out *PostgresUsersConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresUsersConfiguration.
func (in *PostgresUsersConfiguration) DeepCopy() *PostgresUsersConfiguration {
	if in == nil {
		return nil
	}
	out := new(PostgresUsersConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Postgresql) DeepCopyInto(out *Postgresql) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Postgresql.
func (in *Postgresql) DeepCopy() *Postgresql {
	if in == nil {
		return nil
	}
	out := new(Postgresql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Postgresql) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlList) DeepCopyInto(out *PostgresqlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Postgresql, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlList.
func (in *PostgresqlList) DeepCopy() *PostgresqlList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlParam) DeepCopyInto(out *PostgresqlParam) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlParam.
func (in *PostgresqlParam) DeepCopy() *PostgresqlParam {
	if in == nil {
		return nil
	}
	out := new(PostgresqlParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreparedDatabase) DeepCopyInto(out *PreparedDatabase) {
	*out = *in
	if in.PreparedSchemas != nil {
		in, out := &in.PreparedSchemas, &out.PreparedSchemas
		*out = make(map[string]PreparedSchema, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreparedDatabase.
func (in *PreparedDatabase) DeepCopy() *PreparedDatabase {
	if in == nil {
		return nil
	}
	out := new(PreparedDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreparedSchema) DeepCopyInto(out *PreparedSchema) {
	*out = *in
	if in.DefaultRoles != nil {
		in, out := &in.DefaultRoles, &out.DefaultRoles
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreparedSchema.
func (in *PreparedSchema) DeepCopy() *PreparedSchema {
	if in == nil {
		return nil
	}
	out := new(PreparedSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDescription) DeepCopyInto(out *ResourceDescription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDescription.
func (in *ResourceDescription) DeepCopy() *ResourceDescription {
	if in == nil {
		return nil
	}
	out := new(ResourceDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.ResourceRequests = in.ResourceRequests
	out.ResourceLimits = in.ResourceLimits
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalyrConfiguration) DeepCopyInto(out *ScalyrConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalyrConfiguration.
func (in *ScalyrConfiguration) DeepCopy() *ScalyrConfiguration {
	if in == nil {
		return nil
	}
	out := new(ScalyrConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidecar) DeepCopyInto(out *Sidecar) {
	*out = *in
	out.Resources = in.Resources
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]corev1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidecar.
func (in *Sidecar) DeepCopy() *Sidecar {
	if in == nil {
		return nil
	}
	out := new(Sidecar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyDescription) DeepCopyInto(out *StandbyDescription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyDescription.
func (in *StandbyDescription) DeepCopy() *StandbyDescription {
	if in == nil {
		return nil
	}
	out := new(StandbyDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSDescription) DeepCopyInto(out *TLSDescription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSDescription.
func (in *TLSDescription) DeepCopy() *TLSDescription {
	if in == nil {
		return nil
	}
	out := new(TLSDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamsAPIConfiguration) DeepCopyInto(out *TeamsAPIConfiguration) {
	*out = *in
	if in.TeamAPIRoleConfiguration != nil {
		in, out := &in.TeamAPIRoleConfiguration, &out.TeamAPIRoleConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProtectedRoles != nil {
		in, out := &in.ProtectedRoles, &out.ProtectedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostgresSuperuserTeams != nil {
		in, out := &in.PostgresSuperuserTeams, &out.PostgresSuperuserTeams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamsAPIConfiguration.
func (in *TeamsAPIConfiguration) DeepCopy() *TeamsAPIConfiguration {
	if in == nil {
		return nil
	}
	out := new(TeamsAPIConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in UserFlags) DeepCopyInto(out *UserFlags) {
	{
		in := &in
		*out = make(UserFlags, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFlags.
func (in UserFlags) DeepCopy() UserFlags {
	if in == nil {
		return nil
	}
	out := new(UserFlags)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
