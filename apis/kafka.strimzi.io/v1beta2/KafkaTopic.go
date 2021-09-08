// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package v1beta2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// KafkaTopic
type KafkaTopic struct {
	// The specification of the topic.
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec *KafkaTopicSpec `json:"spec,omitempty"`

	// The status of the topic.
	Status *KafkaTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// KafkaTopicList contains a list of instances.
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// A list of Kafka objects.
	Items []KafkaTopic `json:"items,omitempty"`
}

func init() {
	SchemeBuilder.Register(&KafkaTopic{}, &KafkaTopicList{})
}

// The specification of the topic.
type KafkaTopicSpec struct {
	// The topic configuration.
	Config *apiextensions.JSON `json:"config,omitempty"`

	// The number of partitions the topic should have. This cannot be decreased after
	// topic creation. It can be increased after topic creation, but it is important
	// to understand the consequences that has, especially for topics with semantic
	// partitioning. When absent this will default to the broker configuration for
	// `num.partitions`.
	Partitions *int32 `json:"partitions,omitempty"`

	// The number of replicas the topic should have. When absent this will default to
	// the broker configuration for `default.replication.factor`.
	Replicas *int32 `json:"replicas,omitempty"`

	// The name of the topic. When absent this will default to the metadata.name of
	// the topic. It is recommended to not set this unless the topic name is not a
	// valid Kubernetes resource name.
	TopicName *string `json:"topicName,omitempty"`
}

// The topic configuration.
//type KafkaTopicSpecConfig map[string]interface{}

// The status of the topic.
type KafkaTopicStatus struct {
	// List of status conditions.
	Conditions []KafkaTopicStatusConditionsElem `json:"conditions,omitempty"`

	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int32 `json:"observedGeneration,omitempty"`

	// Topic name.
	TopicName *string `json:"topicName,omitempty"`
}

type KafkaTopicStatusConditionsElem struct {
	// Last time the condition of a type changed from one status to another. The
	// required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about the condition's last
	// transition.
	Message *string `json:"message,omitempty"`

	// The reason for the condition's last transition (a single word in CamelCase).
	Reason *string `json:"reason,omitempty"`

	// The status of the condition, either True, False or Unknown.
	Status *string `json:"status,omitempty"`

	// The unique identifier of a condition, used to distinguish between other
	// conditions in the resource.
	Type *string `json:"type,omitempty"`
}