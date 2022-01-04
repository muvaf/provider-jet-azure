/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LinkedServerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LinkedServerParameters struct {

	// +crossplane:generate:reference:type=Cache
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LinkedRedisCacheID *string `json:"linkedRedisCacheId,omitempty" tf:"linked_redis_cache_id,omitempty"`

	// +kubebuilder:validation:Optional
	LinkedRedisCacheIDRef *v1.Reference `json:"linkedRedisCacheIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LinkedRedisCacheIDSelector *v1.Selector `json:"linkedRedisCacheIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation" tf:"linked_redis_cache_location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ServerRole *string `json:"serverRole" tf:"server_role,omitempty"`

	// +crossplane:generate:reference:type=Cache
	// +kubebuilder:validation:Optional
	TargetRedisCacheName *string `json:"targetRedisCacheName,omitempty" tf:"target_redis_cache_name,omitempty"`

	// +kubebuilder:validation:Optional
	TargetRedisCacheNameRef *v1.Reference `json:"targetRedisCacheNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetRedisCacheNameSelector *v1.Selector `json:"targetRedisCacheNameSelector,omitempty" tf:"-"`
}

// LinkedServerSpec defines the desired state of LinkedServer
type LinkedServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServerParameters `json:"forProvider"`
}

// LinkedServerStatus defines the observed state of LinkedServer.
type LinkedServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServer is the Schema for the LinkedServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServerSpec   `json:"spec"`
	Status            LinkedServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServerList contains a list of LinkedServers
type LinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServer `json:"items"`
}

// Repository type metadata.
var (
	LinkedServer_Kind             = "LinkedServer"
	LinkedServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServer_Kind}.String()
	LinkedServer_KindAPIVersion   = LinkedServer_Kind + "." + CRDGroupVersion.String()
	LinkedServer_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServer_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServer{}, &LinkedServerList{})
}
