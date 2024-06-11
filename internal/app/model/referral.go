package model

import "time"

type Referral struct {
	UserID  int64     `firestore:"userId"`
	ChatID  int64     `firestore:"chatId"`
	Company string    `firestore:"company"`
	URL     string    `firestore:"url"`
	Comment string    `firestore:"comment"`
	Created time.Time `firestore:"created"`
}

func NewReferral(doc map[string]interface{}) *Referral {
	return &Referral{
		UserID:  doc["userId"].(int64),
		ChatID:  doc["chatId"].(int64),
		Company: doc["company"].(string),
		URL:     doc["url"].(string),
		Comment: doc["comment"].(string),
		Created: doc["created"].(time.Time),
	}
}
