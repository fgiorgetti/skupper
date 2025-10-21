package client

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func MergeOwnerReferences(obj *v1.ObjectMeta, added []v1.OwnerReference) bool {
	changed := false
	byUid := map[types.UID]v1.OwnerReference{}
	original := obj.OwnerReferences
	for _, ref := range original {
		byUid[ref.UID] = ref
	}
	for _, ref := range added {
		if actual, ok := byUid[ref.UID]; !ok || actual != ref {
			original = append(original, ref)
			changed = true
		}
	}
	if changed {
		obj.OwnerReferences = original
	}
	return changed
}
