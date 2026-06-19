package auth

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null"                           json:"email"`
	Name      string    `gorm:"default:null"                                   json:"name,omitempty"`
	AvatarURL string    `gorm:"default:null"                                   json:"avatar_url,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                 json:"created_at"`
}

type Session struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"not null;index"                                 json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null"                           json:"-"`
	ExpiresAt time.Time `gorm:"not null"                                       json:"expires_at"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                 json:"created_at"`
}

type OTPCode struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string    `gorm:"not null"                                       json:"email"`
	Code      string    `gorm:"not null"                                       json:"-"`
	ExpiresAt time.Time `gorm:"not null"                                       json:"expires_at"`
	Used      bool      `gorm:"default:false"                                  json:"used"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                 json:"created_at"`
}

type Passkey struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID       string    `gorm:"not null;index"                                 json:"user_id"`
	CredentialID string    `gorm:"uniqueIndex;not null"                           json:"credential_id"`
	PublicKey    []byte    `gorm:"not null"                                       json:"-"`
	SignCount    int64     `gorm:"default:0"                                      json:"sign_count"`
	CreatedAt    time.Time `gorm:"autoCreateTime"                                 json:"created_at"`
}

type OAuthAccount struct {
	ID             string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID         string `gorm:"not null;index"                                 json:"user_id"`
	Provider       string `gorm:"not null"                                       json:"provider"`
	ProviderUserID string `gorm:"not null"                                       json:"provider_user_id"`
}
