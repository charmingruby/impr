package event

import (
	"context"
	"fmt"

	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/charmingruby/impr/lib/pkg/parser"
	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"github.com/charmingruby/impr/service/audit/internal/audit/core/service"
	"github.com/charmingruby/impr/service/audit/pkg/logger"
)

func HandleCreateAudit(ctx context.Context, queue messaging.Subscriber, svc *service.Service) {
	err := queue.Subscribe(ctx, func(msg messaging.Message) error {
		logger.Log.Info(fmt.Sprintf("new event received, key: %s", msg.Key))

		var params pb.AuditLog
		if err := parser.BytesToProto(msg.Value, &params); err != nil {
			return err
		}

		logger.Log.Info(fmt.Sprintf("processing %s event: id=%s", params.Subject, msg.Key))

		if err := svc.CreateAudit(service.CreateAuditParams{
			Context:      params.Context,
			Subject:      params.Subject,
			Content:      params.Content,
			DispatchedAt: params.DispatchedAt.AsTime(),
		}); err != nil {
			return err
		}

		logger.Log.Info(fmt.Sprintf("processed %s event: id=%s", params.Subject, msg.Key))

		return nil
	})

	if err != nil {
		logger.Log.Error(fmt.Sprintf("failed to subscribe to audits.create: %v", err))
	}
}
