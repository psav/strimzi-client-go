// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package v1beta1

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// KafkaUser
type KafkaUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// The specification of the user.
	Spec *KafkaUserSpec `json:"spec,omitempty"`

	// The status of the Kafka User.
	Status *KafkaUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// KafkaUserList contains a list of instances.
type KafkaUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// A list of Kafka objects.
	Items []KafkaUser `json:"items,omitempty"`
}

func init() {
	SchemeBuilder.Register(&KafkaUser{}, &KafkaUserList{})
}

// The specification of the user.
type KafkaUserSpec struct {
	// Authentication mechanism enabled for this Kafka user.
	Authentication *KafkaUserSpecAuthentication `json:"authentication,omitempty"`

	// Authorization rules for this Kafka user.
	Authorization *KafkaUserSpecAuthorization `json:"authorization,omitempty"`

	// Quotas on requests to control the broker resources used by clients. Network
	// bandwidth and request rate quotas can be enforced.Kafka documentation for Kafka
	// User quotas can be found at
	// http://kafka.apache.org/documentation/#design_quotas.
	Quotas *KafkaUserSpecQuotas `json:"quotas,omitempty"`

	// Template to specify how Kafka User `Secrets` are generated.
	Template *KafkaUserSpecTemplate `json:"template,omitempty"`
}

// Authentication mechanism enabled for this Kafka user.
type KafkaUserSpecAuthentication struct {
	// Authentication type.
	Type KafkaUserSpecAuthenticationType `json:"type"`
}

type KafkaUserSpecAuthenticationType string

const KafkaUserSpecAuthenticationTypeScramSha512 KafkaUserSpecAuthenticationType = "scram-sha-512"
const KafkaUserSpecAuthenticationTypeTls KafkaUserSpecAuthenticationType = "tls"

// Authorization rules for this Kafka user.
type KafkaUserSpecAuthorization struct {
	// List of ACL rules which should be applied to this user.
	Acls []KafkaUserSpecAuthorizationAclsElem `json:"acls"`

	// Authorization type. Currently the only supported type is `simple`. `simple`
	// authorization type uses Kafka's `kafka.security.authorizer.AclAuthorizer` class
	// for authorization.
	Type KafkaUserSpecAuthorizationType `json:"type"`
}

type KafkaUserSpecAuthorizationAclsElem struct {
	// The host from which the action described in the ACL rule is allowed or denied.
	Host *string `json:"host,omitempty"`

	// Operation which will be allowed or denied. Supported operations are: Read,
	// Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs,
	// DescribeConfigs, IdempotentWrite and All.
	Operation KafkaUserSpecAuthorizationAclsElemOperation `json:"operation"`

	// Indicates the resource for which given ACL rule applies.
	Resource KafkaUserSpecAuthorizationAclsElemResource `json:"resource"`

	// The type of the rule. Currently the only supported type is `allow`. ACL rules
	// with type `allow` are used to allow user to execute the specified operations.
	// Default value is `allow`.
	Type *KafkaUserSpecAuthorizationAclsElemType `json:"type,omitempty"`
}

type KafkaUserSpecAuthorizationAclsElemOperation string

const KafkaUserSpecAuthorizationAclsElemOperationAll KafkaUserSpecAuthorizationAclsElemOperation = "All"
const KafkaUserSpecAuthorizationAclsElemOperationAlter KafkaUserSpecAuthorizationAclsElemOperation = "Alter"
const KafkaUserSpecAuthorizationAclsElemOperationAlterConfigs KafkaUserSpecAuthorizationAclsElemOperation = "AlterConfigs"
const KafkaUserSpecAuthorizationAclsElemOperationClusterAction KafkaUserSpecAuthorizationAclsElemOperation = "ClusterAction"
const KafkaUserSpecAuthorizationAclsElemOperationCreate KafkaUserSpecAuthorizationAclsElemOperation = "Create"
const KafkaUserSpecAuthorizationAclsElemOperationDelete KafkaUserSpecAuthorizationAclsElemOperation = "Delete"
const KafkaUserSpecAuthorizationAclsElemOperationDescribe KafkaUserSpecAuthorizationAclsElemOperation = "Describe"
const KafkaUserSpecAuthorizationAclsElemOperationDescribeConfigs KafkaUserSpecAuthorizationAclsElemOperation = "DescribeConfigs"
const KafkaUserSpecAuthorizationAclsElemOperationIdempotentWrite KafkaUserSpecAuthorizationAclsElemOperation = "IdempotentWrite"
const KafkaUserSpecAuthorizationAclsElemOperationRead KafkaUserSpecAuthorizationAclsElemOperation = "Read"
const KafkaUserSpecAuthorizationAclsElemOperationWrite KafkaUserSpecAuthorizationAclsElemOperation = "Write"

// Indicates the resource for which given ACL rule applies.
type KafkaUserSpecAuthorizationAclsElemResource struct {
	// Name of resource for which given ACL rule applies. Can be combined with
	// `patternType` field to use prefix pattern.
	Name *string `json:"name,omitempty"`

	// Describes the pattern used in the resource field. The supported types are
	// `literal` and `prefix`. With `literal` pattern type, the resource field will be
	// used as a definition of a full name. With `prefix` pattern type, the resource
	// name will be used only as a prefix. Default value is `literal`.
	PatternType *KafkaUserSpecAuthorizationAclsElemResourcePatternType `json:"patternType,omitempty"`

	// Resource type. The available resource types are `topic`, `group`, `cluster`,
	// and `transactionalId`.
	Type KafkaUserSpecAuthorizationAclsElemResourceType `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationAclsElemType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaUserSpecAuthorizationAclsElemType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaUserSpecAuthorizationAclsElemType, v)
	}
	*j = KafkaUserSpecAuthorizationAclsElemType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaUserSpecAuthorizationType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaUserSpecAuthorizationType, v)
	}
	*j = KafkaUserSpecAuthorizationType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationAclsElemResourcePatternType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaUserSpecAuthorizationAclsElemResourcePatternType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaUserSpecAuthorizationAclsElemResourcePatternType, v)
	}
	*j = KafkaUserSpecAuthorizationAclsElemResourcePatternType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationAclsElemOperation) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaUserSpecAuthorizationAclsElemOperation {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaUserSpecAuthorizationAclsElemOperation, v)
	}
	*j = KafkaUserSpecAuthorizationAclsElemOperation(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorization) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["acls"]; !ok || v == nil {
		return fmt.Errorf("field acls: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain KafkaUserSpecAuthorization
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaUserSpecAuthorization(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationAclsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["operation"]; !ok || v == nil {
		return fmt.Errorf("field operation: required")
	}
	if v, ok := raw["resource"]; !ok || v == nil {
		return fmt.Errorf("field resource: required")
	}
	type Plain KafkaUserSpecAuthorizationAclsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaUserSpecAuthorizationAclsElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthentication) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain KafkaUserSpecAuthentication
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaUserSpecAuthentication(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationAclsElemResourceType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaUserSpecAuthorizationAclsElemResourceType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaUserSpecAuthorizationAclsElemResourceType, v)
	}
	*j = KafkaUserSpecAuthorizationAclsElemResourceType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthenticationType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KafkaUserSpecAuthenticationType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KafkaUserSpecAuthenticationType, v)
	}
	*j = KafkaUserSpecAuthenticationType(v)
	return nil
}

type KafkaUserSpecAuthorizationAclsElemResourcePatternType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaUserSpecAuthorizationAclsElemResource) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain KafkaUserSpecAuthorizationAclsElemResource
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaUserSpecAuthorizationAclsElemResource(plain)
	return nil
}

const KafkaUserSpecAuthorizationAclsElemResourcePatternTypeLiteral KafkaUserSpecAuthorizationAclsElemResourcePatternType = "literal"
const KafkaUserSpecAuthorizationAclsElemResourcePatternTypePrefix KafkaUserSpecAuthorizationAclsElemResourcePatternType = "prefix"

type KafkaUserSpecAuthorizationAclsElemResourceType string

const KafkaUserSpecAuthorizationAclsElemResourceTypeCluster KafkaUserSpecAuthorizationAclsElemResourceType = "cluster"
const KafkaUserSpecAuthorizationAclsElemResourceTypeGroup KafkaUserSpecAuthorizationAclsElemResourceType = "group"
const KafkaUserSpecAuthorizationAclsElemResourceTypeTopic KafkaUserSpecAuthorizationAclsElemResourceType = "topic"
const KafkaUserSpecAuthorizationAclsElemResourceTypeTransactionalId KafkaUserSpecAuthorizationAclsElemResourceType = "transactionalId"

type KafkaUserSpecAuthorizationAclsElemType string

const KafkaUserSpecAuthorizationAclsElemTypeAllow KafkaUserSpecAuthorizationAclsElemType = "allow"
const KafkaUserSpecAuthorizationAclsElemTypeDeny KafkaUserSpecAuthorizationAclsElemType = "deny"

type KafkaUserSpecAuthorizationType string

const KafkaUserSpecAuthorizationTypeSimple KafkaUserSpecAuthorizationType = "simple"

// Quotas on requests to control the broker resources used by clients. Network
// bandwidth and request rate quotas can be enforced.Kafka documentation for Kafka
// User quotas can be found at
// http://kafka.apache.org/documentation/#design_quotas.
type KafkaUserSpecQuotas struct {
	// A quota on the maximum bytes per-second that each client group can fetch from a
	// broker before the clients in the group are throttled. Defined on a per-broker
	// basis.
	ConsumerByteRate *int32 `json:"consumerByteRate,omitempty"`

	// A quota on the maximum bytes per-second that each client group can publish to a
	// broker before the clients in the group are throttled. Defined on a per-broker
	// basis.
	ProducerByteRate *int32 `json:"producerByteRate,omitempty"`

	// A quota on the maximum CPU utilization of each client group as a percentage of
	// network and I/O threads.
	RequestPercentage *int32 `json:"requestPercentage,omitempty"`
}

// Template to specify how Kafka User `Secrets` are generated.
type KafkaUserSpecTemplate struct {
	// Template for KafkaUser resources. The template allows users to specify how the
	// `Secret` with password or TLS certificates is generated.
	Secret *KafkaUserSpecTemplateSecret `json:"secret,omitempty"`
}

// Template for KafkaUser resources. The template allows users to specify how the
// `Secret` with password or TLS certificates is generated.
type KafkaUserSpecTemplateSecret struct {
	// Metadata applied to the resource.
	Metadata *KafkaUserSpecTemplateSecretMetadata `json:"metadata,omitempty"`
}

// Metadata applied to the resource.
type KafkaUserSpecTemplateSecretMetadata struct {
	// Annotations added to the resource template. Can be applied to different
	// resources such as `StatefulSets`, `Deployments`, `Pods`, and `Services`.
	Annotations KafkaUserSpecTemplateSecretMetadataAnnotations `json:"annotations,omitempty"`

	// Labels added to the resource template. Can be applied to different resources
	// such as `StatefulSets`, `Deployments`, `Pods`, and `Services`.
	Labels KafkaUserSpecTemplateSecretMetadataLabels `json:"labels,omitempty"`
}

// Annotations added to the resource template. Can be applied to different
// resources such as `StatefulSets`, `Deployments`, `Pods`, and `Services`.
type KafkaUserSpecTemplateSecretMetadataAnnotations map[string]string

// Labels added to the resource template. Can be applied to different resources
// such as `StatefulSets`, `Deployments`, `Pods`, and `Services`.
type KafkaUserSpecTemplateSecretMetadataLabels map[string]string

// The status of the Kafka User.
type KafkaUserStatus struct {
	// List of status conditions.
	Conditions []KafkaUserStatusConditionsElem `json:"conditions,omitempty"`

	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int32 `json:"observedGeneration,omitempty"`

	// The name of `Secret` where the credentials are stored.
	Secret *string `json:"secret,omitempty"`

	// Username.
	Username *string `json:"username,omitempty"`
}

type KafkaUserStatusConditionsElem struct {
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

var enumValues_KafkaUserSpecAuthenticationType = []interface{}{
	"tls",
	"scram-sha-512",
}
var enumValues_KafkaUserSpecAuthorizationAclsElemOperation = []interface{}{
	"Read",
	"Write",
	"Create",
	"Delete",
	"Alter",
	"Describe",
	"ClusterAction",
	"AlterConfigs",
	"DescribeConfigs",
	"IdempotentWrite",
	"All",
}
var enumValues_KafkaUserSpecAuthorizationAclsElemResourcePatternType = []interface{}{
	"literal",
	"prefix",
}
var enumValues_KafkaUserSpecAuthorizationAclsElemResourceType = []interface{}{
	"topic",
	"group",
	"cluster",
	"transactionalId",
}
var enumValues_KafkaUserSpecAuthorizationAclsElemType = []interface{}{
	"allow",
	"deny",
}
var enumValues_KafkaUserSpecAuthorizationType = []interface{}{
	"simple",
}
