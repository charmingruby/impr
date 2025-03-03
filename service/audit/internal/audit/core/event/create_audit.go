package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/charmingruby/impr/service/audit/internal/audit/core/service"
	"github.com/charmingruby/impr/service/audit/pkg/logger"
)

func HandleCreateAudit(ctx context.Context, queue messaging.Subscriber, svc *service.Service) {
	err := queue.Subscribe(ctx, func(msg messaging.Message) error {
		logger.Log.Info(fmt.Sprintf("received %s event: id=%s", "audits.create", msg.Key))

		var params service.CreateAuditParams
		if err := json.Unmarshal(msg.Value, &params); err != nil {
			return err
		}

		if err := svc.CreateAudit(params); err != nil {
			return err
		}

		logger.Log.Info(fmt.Sprintf("processed %s event: id=%s", "audits.create", msg.Key))

		return nil
	})

	if err != nil {
		logger.Log.Error(fmt.Sprintf("error consuming messages: %v", err))
	}
}
