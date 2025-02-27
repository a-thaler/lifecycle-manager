// Code generated by MockGen. DO NOT EDIT.
// Source: object.go

// Package mock_v2 is a generated GoMock package.
package mock_v2

import (
	"github.com/kyma-project/lifecycle-manager/internal/declarative/v2"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
)

// MockObject is a mock of Object interface.
type MockObject struct {
	ctrl     *gomock.Controller
	recorder *MockObjectMockRecorder
}

// MockObjectMockRecorder is the mock recorder for MockObject.
type MockObjectMockRecorder struct {
	mock *MockObject
}

// NewMockObject creates a new mock instance.
func NewMockObject(ctrl *gomock.Controller) *MockObject {
	mock := &MockObject{ctrl: ctrl}
	mock.recorder = &MockObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObject) EXPECT() *MockObjectMockRecorder {
	return m.recorder
}

// DeepCopyObject mocks base method.
func (m *MockObject) DeepCopyObject() runtime.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopyObject")
	ret0, _ := ret[0].(runtime.Object)
	return ret0
}

// DeepCopyObject indicates an expected call of DeepCopyObject.
func (mr *MockObjectMockRecorder) DeepCopyObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyObject", reflect.TypeOf((*MockObject)(nil).DeepCopyObject))
}

// GetAnnotations mocks base method.
func (m *MockObject) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockObjectMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockObject)(nil).GetAnnotations))
}

// GetCreationTimestamp mocks base method.
func (m *MockObject) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp.
func (mr *MockObjectMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockObject)(nil).GetCreationTimestamp))
}

// GetDeletionGracePeriodSeconds mocks base method.
func (m *MockObject) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds.
func (mr *MockObjectMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockObject)(nil).GetDeletionGracePeriodSeconds))
}

// GetDeletionTimestamp mocks base method.
func (m *MockObject) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp.
func (mr *MockObjectMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockObject)(nil).GetDeletionTimestamp))
}

// GetFinalizers mocks base method.
func (m *MockObject) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers.
func (mr *MockObjectMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockObject)(nil).GetFinalizers))
}

// GetGenerateName mocks base method.
func (m *MockObject) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName.
func (mr *MockObjectMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockObject)(nil).GetGenerateName))
}

// GetGeneration mocks base method.
func (m *MockObject) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration.
func (mr *MockObjectMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockObject)(nil).GetGeneration))
}

// GetLabels mocks base method.
func (m *MockObject) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockObjectMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockObject)(nil).GetLabels))
}

// GetManagedFields mocks base method.
func (m *MockObject) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields.
func (mr *MockObjectMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockObject)(nil).GetManagedFields))
}

// GetName mocks base method.
func (m *MockObject) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockObjectMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockObject)(nil).GetName))
}

// GetNamespace mocks base method.
func (m *MockObject) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockObjectMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockObject)(nil).GetNamespace))
}

// GetObjectKind mocks base method.
func (m *MockObject) GetObjectKind() schema.ObjectKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectKind")
	ret0, _ := ret[0].(schema.ObjectKind)
	return ret0
}

// GetObjectKind indicates an expected call of GetObjectKind.
func (mr *MockObjectMockRecorder) GetObjectKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectKind", reflect.TypeOf((*MockObject)(nil).GetObjectKind))
}

// GetOwnerReferences mocks base method.
func (m *MockObject) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences.
func (mr *MockObjectMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockObject)(nil).GetOwnerReferences))
}

// GetResourceVersion mocks base method.
func (m *MockObject) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion.
func (mr *MockObjectMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockObject)(nil).GetResourceVersion))
}

// GetSelfLink mocks base method.
func (m *MockObject) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink.
func (mr *MockObjectMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockObject)(nil).GetSelfLink))
}

// GetStatus mocks base method.
func (m *MockObject) GetStatus() v2.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(v2.Status)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockObjectMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockObject)(nil).GetStatus))
}

// GetUID mocks base method.
func (m *MockObject) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID.
func (mr *MockObjectMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockObject)(nil).GetUID))
}

// SetAnnotations mocks base method.
func (m *MockObject) SetAnnotations(annotations map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", annotations)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockObjectMockRecorder) SetAnnotations(annotations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockObject)(nil).SetAnnotations), annotations)
}

// SetCreationTimestamp mocks base method.
func (m *MockObject) SetCreationTimestamp(timestamp v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", timestamp)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp.
func (mr *MockObjectMockRecorder) SetCreationTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockObject)(nil).SetCreationTimestamp), timestamp)
}

// SetDeletionGracePeriodSeconds mocks base method.
func (m *MockObject) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds.
func (mr *MockObjectMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockObject)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// SetDeletionTimestamp mocks base method.
func (m *MockObject) SetDeletionTimestamp(timestamp *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", timestamp)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp.
func (mr *MockObjectMockRecorder) SetDeletionTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockObject)(nil).SetDeletionTimestamp), timestamp)
}

// SetFinalizers mocks base method.
func (m *MockObject) SetFinalizers(finalizers []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", finalizers)
}

// SetFinalizers indicates an expected call of SetFinalizers.
func (mr *MockObjectMockRecorder) SetFinalizers(finalizers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockObject)(nil).SetFinalizers), finalizers)
}

// SetGenerateName mocks base method.
func (m *MockObject) SetGenerateName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", name)
}

// SetGenerateName indicates an expected call of SetGenerateName.
func (mr *MockObjectMockRecorder) SetGenerateName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockObject)(nil).SetGenerateName), name)
}

// SetGeneration mocks base method.
func (m *MockObject) SetGeneration(generation int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", generation)
}

// SetGeneration indicates an expected call of SetGeneration.
func (mr *MockObjectMockRecorder) SetGeneration(generation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockObject)(nil).SetGeneration), generation)
}

// SetLabels mocks base method.
func (m *MockObject) SetLabels(labels map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", labels)
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockObjectMockRecorder) SetLabels(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockObject)(nil).SetLabels), labels)
}

// SetManagedFields mocks base method.
func (m *MockObject) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", managedFields)
}

// SetManagedFields indicates an expected call of SetManagedFields.
func (mr *MockObjectMockRecorder) SetManagedFields(managedFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockObject)(nil).SetManagedFields), managedFields)
}

// SetName mocks base method.
func (m *MockObject) SetName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", name)
}

// SetName indicates an expected call of SetName.
func (mr *MockObjectMockRecorder) SetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockObject)(nil).SetName), name)
}

// SetNamespace mocks base method.
func (m *MockObject) SetNamespace(namespace string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", namespace)
}

// SetNamespace indicates an expected call of SetNamespace.
func (mr *MockObjectMockRecorder) SetNamespace(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockObject)(nil).SetNamespace), namespace)
}

// SetOwnerReferences mocks base method.
func (m *MockObject) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences.
func (mr *MockObjectMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockObject)(nil).SetOwnerReferences), arg0)
}

// SetResourceVersion mocks base method.
func (m *MockObject) SetResourceVersion(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", version)
}

// SetResourceVersion indicates an expected call of SetResourceVersion.
func (mr *MockObjectMockRecorder) SetResourceVersion(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockObject)(nil).SetResourceVersion), version)
}

// SetSelfLink mocks base method.
func (m *MockObject) SetSelfLink(selfLink string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", selfLink)
}

// SetSelfLink indicates an expected call of SetSelfLink.
func (mr *MockObjectMockRecorder) SetSelfLink(selfLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockObject)(nil).SetSelfLink), selfLink)
}

// SetStatus mocks base method.
func (m *MockObject) SetStatus(arg0 v2.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatus", arg0)
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockObjectMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockObject)(nil).SetStatus), arg0)
}

// SetUID mocks base method.
func (m *MockObject) SetUID(uid types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", uid)
}

// SetUID indicates an expected call of SetUID.
func (mr *MockObjectMockRecorder) SetUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockObject)(nil).SetUID), uid)
}
