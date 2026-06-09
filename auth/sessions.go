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
	Username  string
	ExpiresAt time.Time
}

// In-memory Session Store
var (
	sessions = make(map[string]Session)
	mu       sync.Mutex
)

func isAuhenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no session"})
			return
		}

		mu.Lock()
		session, exists := sessions[sessionID]
		mu.Unlock()

		if !exists || time.Now().After(session.ExpiresAt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			return
		}

		// attach user info to context
		c.Set("username", session.Username)
		c.Next()
	}
}

// Generate random session ID
func generateSessionID() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func createSession(username string) string {
	sessionID := generateSessionID()

	mu.Lock()
	sessions[sessionID] = Session{
		Username:  username,
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
