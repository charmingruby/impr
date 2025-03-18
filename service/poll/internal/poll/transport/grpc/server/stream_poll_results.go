package server

import (
	"fmt"
	"log"
	"time"

	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const pollInterval = 3 * time.Second

func (s *Service) StreamPollResults(payload *pb.PollPayload, stream pb.PollService_StreamPollResultsServer) error {
	pollID := payload.PollId
	var lastSent *model.PollSummary

	logger.Log.Debug(fmt.Sprintf("Starting stream for Poll ID: %s", pollID))

	for {
		select {
		case <-stream.Context().Done():
			log.Println("Client disconnected, ending stream.")
			return nil
		default:
			summary, err := s.domainService.GetPollSummary(pollID)
			if err != nil {
				logger.Log.Debug(fmt.Sprintf("Error fetching poll summary: %v", err))
				time.Sleep(pollInterval)
				continue
			}

			if lastSent != nil && compareSummaries(lastSent, summary) {
				time.Sleep(pollInterval)
				continue
			}

			response := &pb.PollResponse{
				PollId:    summary.PollID,
				ExpiresAt: timestamppb.New(*summary.ExpiresAt),
			}

			for _, option := range summary.Options {
				response.Options = append(response.Options, &pb.PollOption{
					Id:        option.PollOptionID,
					Content:   option.Content,
					VoteCount: int32(option.Votes),
				})
			}

			if err := stream.Send(response); err != nil {
				logger.Log.Debug(fmt.Sprintf("Error sending stream: %v", err))
				return err
			}

			lastSent = summary

			time.Sleep(pollInterval)
		}
	}
}

func compareSummaries(a, b *model.PollSummary) bool {
	if a.PollID != b.PollID || len(a.Options) != len(b.Options) {
		return false
	}

	for i := range a.Options {
		if a.Options[i].Votes != b.Options[i].Votes {
			return false
		}
	}

	return true
}
