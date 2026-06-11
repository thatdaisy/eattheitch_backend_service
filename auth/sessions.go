package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Session model
type Session struct {
	Email     string
	ExpiresAt time.Time
}

// In-memory Session Store
var (
	sessions = make(map[string]Session)
	mu       sync.Mutex
)

func IsAuhenticated() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionID, err := context.Cookie("session_id")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no session"})
			return
		}

		mu.Lock()
		session, exists := sessions[sessionID]
		mu.Unlock()

		if !exists || time.Now().After(session.ExpiresAt) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			return
		}

		context.Set("email", session.Email)
		context.Next()
	}
}

// Generate random session ID
func generateSessionID() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func createSession(email string) string {
	sessionID := generateSessionID()

	mu.Lock()
	sessions[sessionID] = Session{
		Email:     email,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}
	mu.Unlock()
	return sessionID
}

func deleteSession(sessionID string) {
	mu.Lock()
	delete(sessions, sessionID)
	mu.Unlock()
}
