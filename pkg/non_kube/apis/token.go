package apis

import (
	"bufio"
	"bytes"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

type Token struct {
	Link   *v1alpha1.Link
	Secret *v1.Secret
}

func (t *Token) Marshal() ([]byte, error) {
	s := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	buffer := new(bytes.Buffer)
	writer := bufio.NewWriter(buffer)
	_, _ = writer.Write([]byte("---\n"))
	err := s.Encode(t.Secret, writer)
	if err != nil {
		return nil, err
	}
	_, _ = writer.Write([]byte("---\n"))
	err = s.Encode(t.Link, writer)
	if err != nil {
		return nil, err
	}
	writer.Flush()
	return buffer.Bytes(), nil
}
