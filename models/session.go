package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/haskaalo/intribox/utils"
)

// ErrNotValidSessionToken Invalid token format
var ErrNotValidSessionToken = errors.New("Not a valid token")

const (
	// SessionExpireTime Session expire time for a token
	SessionExpireTime = time.Duration(336) * time.Hour // 2 weeks of inactivity

	// SessionHeaderName Session Token name
	SessionHeaderName = "X-Intribox-Token"

	// SessionPrefix Prefix used for a user session
	SessionPrefix = "session:"

	// SessionGroupPrefix Prefix used for set of user specific session
	SessionGroupPrefix = "session-group:"
)

// Session stored on Redis
type Session struct {
	UserID    int
	Selector  string
	Validator string
	CreatedAt int
}

// GetSessionBySelector Get Session model with a selector (string)
func GetSessionBySelector(selector string) (*Session, error) {
	sess := new(Session)
	hashSelector := utils.SHA1([]byte(selector))
	vals, err := r.HGetAll(SessionPrefix + hashSelector).Result()
	if len(vals) == 0 { // Only way to know if key doesn't exist
		return nil, redis.Nil
	}
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(vals["userid"])
	if err != nil {
		return nil, err
	}

	createdAt, err := strconv.Atoi(vals["createdat"])
	if err != nil {
		return nil, err
	}

	sess.UserID = userID
	sess.Selector = selector
	sess.Validator = vals["validator"]
	sess.CreatedAt = createdAt

	return sess, nil
}

// GetSessionByToken get Session with Token
func GetSessionByToken(token string) (*Session, error) {
	parsedToken, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	session, err := GetSessionBySelector(parsedToken.Selector)
	if err == redis.Nil {
		return nil, ErrNotValidSessionToken
	} else if err != nil {
		return nil, err
	}

	hashedValidator := utils.SHA1([]byte(parsedToken.Validator))
	if hashedValidator == session.Validator {
		return session, nil
	}

	return nil, ErrNotValidSessionToken
}

// DeleteSessionBySelector Delete a session by using a selector
func DeleteSessionBySelector(selector string) error {
	hashSelector := utils.SHA1([]byte(selector))
	_, err := r.Del(SessionPrefix + hashSelector).Result()

	return err
}

// ResetTimeSession Reset time of a session based on config.ini Expire time
func (s Session) ResetTimeSession() error {
	hashSelector := utils.SHA1([]byte(s.Selector))

	err := r.Expire(SessionPrefix+hashSelector, SessionExpireTime).Err()
	if err != nil {
		return err
	}

	err = r.Expire(SessionGroupPrefix+strconv.Itoa(s.UserID), SessionExpireTime).Err()
	if err != nil {
		return err
	}

	return nil
}

// InitiateSession Create a session for a user
// TODO: Allow a maximum of 10 active sessions. Otherwise, delete the oldest
func InitiateSession(uid int) (selector string, validator string, err error) {
	// Get active sessions
	activeSessions, err := r.SMembers(SessionGroupPrefix + strconv.Itoa(uid)).Result()
	if err != nil {
		return "", "", err
	}

	// Checks if the number of active session if higher than 10
	if len(activeSessions) >= 10 {
		err = DeleteOldestSession(uid)
		if err != nil {
			return "", "", err
		}
	}

	// s for selector
	// v for validator
	// s for selector
	s := utils.RandString(12)
	v := utils.RandString(50)
	hashS := utils.SHA1([]byte(s))

	// Create new session
	// Validator acts as some sort of key
	err = r.HMSet(SessionPrefix+hashS, map[string]interface{}{
		"userid":    uid,
		"validator": utils.SHA1([]byte(v)),
		"createdat": time.Now().Unix(),
	}).Err()
	if err != nil {
		return "", "", err
	}

	// Add the new session to a set of session ids for that user
	err = r.SAdd(SessionGroupPrefix+strconv.Itoa(uid), hashS).Err()
	if err != nil {
		return "", "", err
	}

	// Make sure the session expire being inactive for a while
	err = Session{
		Selector: s,
		UserID:   uid,
	}.ResetTimeSession()
	if err != nil {
		return "", "", err
	}

	return s, v, nil
}

// DeleteOldestSession This function deletes the session that will expire the soonest
// for a given userID
func DeleteOldestSession(uid int) error {
	// Get all active session for the user
	vals, err := r.SMembers(SessionGroupPrefix + strconv.Itoa(uid)).Result()
	if err != nil {
		return err
	}

	oldestSessHashSelector := ""
	oldestSessTTL := 1.7e+308

	// Get the session that will expire the soonest and set it
	for _, sessionSelector := range vals {
		duration := r.TTL(SessionPrefix + sessionSelector).Val().Seconds()
		if duration < oldestSessTTL {
			oldestSessHashSelector = sessionSelector
			oldestSessTTL = duration
		}
	}

	// Delete the oldest session if it exist
	if oldestSessHashSelector != "" {
		_, err = r.Del(SessionPrefix + oldestSessHashSelector).Result()
	}

	return err
}
