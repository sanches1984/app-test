package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/sanches1984/referral-bot/internal/app/model"
	"google.golang.org/api/iterator"

	"google.golang.org/api/option"
)

const (
	projectID     = "referral-bot-prod"
	databaseID    = "referral-db"
	collReferrals = "referrals"
)

type Storage struct {
	db *firestore.CollectionRef
}

func New(ctx context.Context, filename string) (*Storage, func() error, error) {
	store, err := firestore.NewClientWithDatabase(ctx, projectID, databaseID, option.WithCredentialsFile(filename))
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	return &Storage{
		db: store.Collection(collReferrals),
	}, store.Close, nil
}

func (s *Storage) AddReferral(ctx context.Context, rec *model.Referral) error {
	docs, err := s.db.Where("userId", "==", rec.UserID).Documents(ctx).GetAll()
	if err != nil {
		return err
	}
	if len(docs) == 0 {
		_, _, err := s.db.Add(ctx, rec)
		return err
	}

	_, err = s.db.Doc(docs[0].Ref.ID).Set(ctx, rec)
	return err
}

func (s *Storage) DeleteReferral(ctx context.Context, userID int64) error {
	docs := s.db.Where("userId", "==", userID).Documents(ctx)
	for {
		doc, err := docs.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return err
		}

		_, err = s.db.Doc(doc.Ref.ID).Delete(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) GetReferrals(ctx context.Context) ([]*model.Referral, error) {
	list := make([]*model.Referral, 0, 100)
	iter := s.db.OrderBy("created", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		list = append(list, model.NewReferral(doc.Data()))
	}

	return list, nil
}
