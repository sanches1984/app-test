//go:build payment
// +build payment

package storage

import (
	"context"
	"github.com/sanches1984/referral-bot/internal/app/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStorageCreate(t *testing.T) {
	ctx := context.Background()
	db, closer, err := New(ctx, "db_credentials.json")
	require.NoError(t, err)
	defer closer()

	err = db.AddReferral(ctx, &model.Referral{
		UserID:  1,
		ChatID:  2,
		Company: "aa",
		URL:     "bb",
		Comment: "cc",
		Created: time.Now(),
	})
	require.NoError(t, err)
}

func TestStorageDelete(t *testing.T) {
	ctx := context.Background()
	db, closer, err := New(ctx, "db_credentials.json")
	require.NoError(t, err)
	defer closer()

	err = db.DeleteReferral(ctx, 3)
	require.NoError(t, err)
}
