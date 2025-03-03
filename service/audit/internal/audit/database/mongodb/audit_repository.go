package mongodb

import (
	"context"

	"github.com/charmingruby/impr/service/audit/internal/audit/core/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	AUDIT_COLLECTION = "audits"
)

type AuditRepository struct {
	db *mongo.Database
}

func NewAuditRepository(db *mongo.Database) *AuditRepository {
	return &AuditRepository{
		db: db,
	}
}

func (r *AuditRepository) Create(audit model.Audit) error {
	collection := r.db.Collection(AUDIT_COLLECTION)

	_, err := collection.InsertOne(context.Background(), audit)
	if err != nil {
		return err
	}

	return nil
}
