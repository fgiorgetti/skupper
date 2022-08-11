package client

import "github.com/skupperproject/skupper/api/types"

type VanClientCommon struct {
	Namespace string
	Client    types.VanClientInterface
}
