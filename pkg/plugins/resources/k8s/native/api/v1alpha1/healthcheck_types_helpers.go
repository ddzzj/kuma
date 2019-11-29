package v1alpha1

import (
	proto "github.com/Kong/kuma/api/mesh/v1alpha1"
	"github.com/Kong/kuma/pkg/plugins/resources/k8s/native/pkg/model"
	"github.com/Kong/kuma/pkg/plugins/resources/k8s/native/pkg/registry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (pt *HealthCheck) GetObjectMeta() *metav1.ObjectMeta {
	return &pt.ObjectMeta
}

func (pt *HealthCheck) SetObjectMeta(m *metav1.ObjectMeta) {
	pt.ObjectMeta = *m
}

func (pt *HealthCheck) GetMesh() string {
	return pt.Mesh
}

func (pt *HealthCheck) SetMesh(mesh string) {
	pt.Mesh = mesh
}

func (pt *HealthCheck) GetSpec() map[string]interface{} {
	return pt.Spec
}

func (pt *HealthCheck) SetSpec(spec map[string]interface{}) {
	pt.Spec = spec
}

func (pt *HealthCheck) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *HealthCheckList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&proto.HealthCheck{}, &HealthCheck{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "HealthCheck",
		},
	})
	registry.RegisterListType(&proto.HealthCheck{}, &HealthCheckList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "HealthCheckList",
		},
	})
}
