package event

import (
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/core/id"
	"github.com/charmingruby/impr/lib/pkg/parser"
	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateAuditMessageParams struct {
	Context      string    `json:"context"`
	Subject      string    `json:"subject"`
	Content      string    `json:"content"`
	DispatchedAt time.Time `json:"dispatched_at"`
}

func CreateAuditMessage(params CreateAuditMessageParams) ([]byte, error) {
	protoMsg := &pb.AuditLog{
		Id:           id.New(),
		Context:      params.Context,
		Subject:      params.Subject,
		Content:      params.Content,
		DispatchedAt: timestamppb.New(params.DispatchedAt),
	}

	msg, err := parser.ProtoToBytes(protoMsg)
	if err != nil {
		return nil, core_err.NewSerializationErr("create audit message", "bytes")
	}

	return msg, nil
}
