package rest

import (
	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/db"
	"github.com/threefoldtech/tfgrid-sdk-go/gapi/pkg/types"
)

func ConvertNodes(dbNode []db.Node) (apiNode []types.Node) {
	for _, node := range dbNode {
		apiNode = append(apiNode, types.Node{
			NodeId: node.NodeId,
		})
	}

	return
}
