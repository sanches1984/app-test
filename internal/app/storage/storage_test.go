//go:build payment
// +build payment

package storage

import (
	"context"
	"github.com/sanches1984/referral-bot/internal/app/model"
	"github.com/stretchr/testify/assert"
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

func TestStorageGet(t *testing.T) {
	ctx := context.Background()
	now := time.Now().UTC()
	db, closer, err := New(ctx, "db_credentials.json")
	require.NoError(t, err)
	defer closer()

	row1 := &model.Referral{
		UserID:  1,
		ChatID:  2,
		Company: "aa",
		URL:     "bb",
		Comment: "cc",
		Created: now,
	}
	row2 := &model.Referral{
		UserID:  3,
		ChatID:  4,
		Company: "dd",
		URL:     "ee",
		Comment: "ff",
		Created: now.Add(time.Second),
	}

	err = db.AddReferral(ctx, row1)
	require.NoError(t, err)

	err = db.AddReferral(ctx, row2)
	require.NoError(t, err)

	list, err := db.GetReferrals(ctx)
	require.NoError(t, err)
	assert.Equal(t, []*model.Referral{row2, row1}, list)
}
