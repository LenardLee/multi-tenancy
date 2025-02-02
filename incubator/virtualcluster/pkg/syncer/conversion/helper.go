/*
Copyright 2019 The Kubernetes Authors.

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

package conversion

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation"
)

const (
	LabelCluster = "tenancy.x-k8s.io/cluster"
)

var DefaultDeletionPolicy = v1.DeletePropagationBackground

func ToSuperMasterNamespace(cluster, ns string) string {
	targetNamespace := strings.Join([]string{cluster, ns}, "-");
	if len(targetNamespace) > validation.DNS1123SubdomainMaxLength {
		digest := sha256.Sum256([]byte(targetNamespace))
		return targetNamespace[0:57] + "-" + hex.EncodeToString(digest[0:])[0:5]
	}
	return targetNamespace
}

func GetOwner(obj runtime.Object) (cluster, namespace string) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return "", ""
	}

	cluster = meta.GetAnnotations()[LabelCluster]
	namespace = strings.TrimPrefix(meta.GetNamespace(), cluster+"-")
	return cluster, namespace
}

func BuildMetadata(targetNamespace string, obj runtime.Object) (runtime.Object, error) {
	target := obj.DeepCopyObject()
	m, err := meta.Accessor(target)
	if err != nil {
		return nil, err
	}

	resetMetadata(m)
	if len(targetNamespace) > 0 {
		m.SetNamespace(targetNamespace)
	}

	return target, nil
}

func BuildSuperMasterNamespace(cluster string, obj runtime.Object) (runtime.Object, error) {
	target := obj.DeepCopyObject()
	m, err := meta.Accessor(target)
	if err != nil {
		return nil, err
	}

	resetMetadata(m)

	targetName := strings.Join([]string{cluster, m.GetName()}, "-")
	m.SetName(targetName)
	return target, nil
}

func resetMetadata(obj v1.Object) {
	obj.SetGenerateName("")
	obj.SetSelfLink("")
	obj.SetUID("")
	obj.SetResourceVersion("")
	obj.SetGeneration(0)
	obj.SetCreationTimestamp(v1.Time{})
	obj.SetDeletionTimestamp(nil)
	obj.SetDeletionGracePeriodSeconds(nil)
	obj.SetOwnerReferences(nil)
	obj.SetFinalizers(nil)
	obj.SetClusterName("")
	obj.SetInitializers(nil)
}

func MutatePod(namespace string, pod *corev1.Pod) {
	pod.Status = corev1.PodStatus{}
	pod.Spec.NodeName = ""
	pod.Spec.ServiceAccountName = ""

	for i := range pod.Spec.Containers {
		for j, env := range pod.Spec.Containers[i].Env {
			if env.ValueFrom != nil && env.ValueFrom.FieldRef != nil && env.ValueFrom.FieldRef.FieldPath == "metadata.name" {
				pod.Spec.Containers[i].Env[j].ValueFrom = nil
				pod.Spec.Containers[i].Env[j].Value = pod.Name
			}
			if env.ValueFrom != nil && env.ValueFrom.FieldRef != nil && env.ValueFrom.FieldRef.FieldPath == "metadata.namespace" {
				pod.Spec.Containers[i].Env[j].ValueFrom = nil
				pod.Spec.Containers[i].Env[j].Value = namespace
			}
		}
	}
}
