// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/client-go/pkg/api"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	autoscaling "k8s.io/client-go/pkg/apis/autoscaling"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference,
		Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference,
		Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler,
		Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler,
		Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList,
		Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList,
		Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec,
		Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec,
		Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus,
		Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus,
		Convert_v1_MetricSpec_To_autoscaling_MetricSpec,
		Convert_autoscaling_MetricSpec_To_v1_MetricSpec,
		Convert_v1_MetricStatus_To_autoscaling_MetricStatus,
		Convert_autoscaling_MetricStatus_To_v1_MetricStatus,
		Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource,
		Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource,
		Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus,
		Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus,
		Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource,
		Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource,
		Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus,
		Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus,
		Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource,
		Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource,
		Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus,
		Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus,
		Convert_v1_Scale_To_autoscaling_Scale,
		Convert_autoscaling_Scale_To_v1_Scale,
		Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec,
		Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec,
		Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus,
		Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus,
	)
}

func autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

func Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in, out, s)
}

func autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

func Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]autoscaling.HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			if err := Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			if err := Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	// WARNING: in.TargetCPUUtilizationPercentage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	// WARNING: in.Metrics requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	// WARNING: in.CurrentCPUUtilizationPercentage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	// WARNING: in.CurrentMetrics requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_MetricSpec_To_autoscaling_MetricSpec(in *MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	out.Object = (*autoscaling.ObjectMetricSource)(unsafe.Pointer(in.Object))
	out.Pods = (*autoscaling.PodsMetricSource)(unsafe.Pointer(in.Pods))
	out.Resource = (*autoscaling.ResourceMetricSource)(unsafe.Pointer(in.Resource))
	return nil
}

func Convert_v1_MetricSpec_To_autoscaling_MetricSpec(in *MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	return autoConvert_v1_MetricSpec_To_autoscaling_MetricSpec(in, out, s)
}

func autoConvert_autoscaling_MetricSpec_To_v1_MetricSpec(in *autoscaling.MetricSpec, out *MetricSpec, s conversion.Scope) error {
	out.Type = MetricSourceType(in.Type)
	out.Object = (*ObjectMetricSource)(unsafe.Pointer(in.Object))
	out.Pods = (*PodsMetricSource)(unsafe.Pointer(in.Pods))
	out.Resource = (*ResourceMetricSource)(unsafe.Pointer(in.Resource))
	return nil
}

func Convert_autoscaling_MetricSpec_To_v1_MetricSpec(in *autoscaling.MetricSpec, out *MetricSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricSpec_To_v1_MetricSpec(in, out, s)
}

func autoConvert_v1_MetricStatus_To_autoscaling_MetricStatus(in *MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	out.Object = (*autoscaling.ObjectMetricStatus)(unsafe.Pointer(in.Object))
	out.Pods = (*autoscaling.PodsMetricStatus)(unsafe.Pointer(in.Pods))
	out.Resource = (*autoscaling.ResourceMetricStatus)(unsafe.Pointer(in.Resource))
	return nil
}

func Convert_v1_MetricStatus_To_autoscaling_MetricStatus(in *MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	return autoConvert_v1_MetricStatus_To_autoscaling_MetricStatus(in, out, s)
}

func autoConvert_autoscaling_MetricStatus_To_v1_MetricStatus(in *autoscaling.MetricStatus, out *MetricStatus, s conversion.Scope) error {
	out.Type = MetricSourceType(in.Type)
	out.Object = (*ObjectMetricStatus)(unsafe.Pointer(in.Object))
	out.Pods = (*PodsMetricStatus)(unsafe.Pointer(in.Pods))
	out.Resource = (*ResourceMetricStatus)(unsafe.Pointer(in.Resource))
	return nil
}

func Convert_autoscaling_MetricStatus_To_v1_MetricStatus(in *autoscaling.MetricStatus, out *MetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricStatus_To_v1_MetricStatus(in, out, s)
}

func autoConvert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.TargetValue = in.TargetValue
	return nil
}

func Convert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	return autoConvert_v1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in, out, s)
}

func autoConvert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.TargetValue = in.TargetValue
	return nil
}

func Convert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *ObjectMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_ObjectMetricSource_To_v1_ObjectMetricSource(in, out, s)
}

func autoConvert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.CurrentValue = in.CurrentValue
	return nil
}

func Convert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	return autoConvert_v1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in, out, s)
}

func autoConvert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *ObjectMetricStatus, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.CurrentValue = in.CurrentValue
	return nil
}

func Convert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *ObjectMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ObjectMetricStatus_To_v1_ObjectMetricStatus(in, out, s)
}

func autoConvert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.TargetAverageValue = in.TargetAverageValue
	return nil
}

func Convert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	return autoConvert_v1_PodsMetricSource_To_autoscaling_PodsMetricSource(in, out, s)
}

func autoConvert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *PodsMetricSource, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.TargetAverageValue = in.TargetAverageValue
	return nil
}

func Convert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *PodsMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_PodsMetricSource_To_v1_PodsMetricSource(in, out, s)
}

func autoConvert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

func Convert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	return autoConvert_v1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in, out, s)
}

func autoConvert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *PodsMetricStatus, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

func Convert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *PodsMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_PodsMetricStatus_To_v1_PodsMetricStatus(in, out, s)
}

func autoConvert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	out.Name = api.ResourceName(in.Name)
	out.TargetAverageUtilization = (*int32)(unsafe.Pointer(in.TargetAverageUtilization))
	out.TargetAverageValue = (*resource.Quantity)(unsafe.Pointer(in.TargetAverageValue))
	return nil
}

func Convert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	return autoConvert_v1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in, out, s)
}

func autoConvert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *ResourceMetricSource, s conversion.Scope) error {
	out.Name = api_v1.ResourceName(in.Name)
	out.TargetAverageUtilization = (*int32)(unsafe.Pointer(in.TargetAverageUtilization))
	out.TargetAverageValue = (*resource.Quantity)(unsafe.Pointer(in.TargetAverageValue))
	return nil
}

func Convert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *ResourceMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_ResourceMetricSource_To_v1_ResourceMetricSource(in, out, s)
}

func autoConvert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = api.ResourceName(in.Name)
	out.CurrentAverageUtilization = (*int32)(unsafe.Pointer(in.CurrentAverageUtilization))
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

func Convert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	return autoConvert_v1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in, out, s)
}

func autoConvert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *ResourceMetricStatus, s conversion.Scope) error {
	out.Name = api_v1.ResourceName(in.Name)
	out.CurrentAverageUtilization = (*int32)(unsafe.Pointer(in.CurrentAverageUtilization))
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

func Convert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *ResourceMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ResourceMetricStatus_To_v1_ResourceMetricStatus(in, out, s)
}

func autoConvert_v1_Scale_To_autoscaling_Scale(in *Scale, out *autoscaling.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Scale_To_autoscaling_Scale(in *Scale, out *autoscaling.Scale, s conversion.Scope) error {
	return autoConvert_v1_Scale_To_autoscaling_Scale(in, out, s)
}

func autoConvert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *Scale, s conversion.Scope) error {
	return autoConvert_autoscaling_Scale_To_v1_Scale(in, out, s)
}

func autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

func Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	return autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}

func autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

func Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *ScaleSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in, out, s)
}

func autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}

func Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	return autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in, out, s)
}

func autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}

func Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *ScaleStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in, out, s)
}
