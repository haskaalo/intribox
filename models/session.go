package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"

	"github.com/haskaalo/intribox/modules/hash"
	"github.com/haskaalo/intribox/modules/random"
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
	hashSelector := hash.SHA1([]byte(selector))
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
		return &Session{}, err
	}

	session, err := GetSessionBySelector(parsedToken.Selector)
	if err == redis.Nil {
		return &Session{}, ErrNotValidSessionToken
	} else if err != nil {
		return &Session{}, err
	}

	hashedValidator := hash.SHA1([]byte(parsedToken.Validator))
	if hashedValidator == session.Validator {
		return session, nil
	}

	return &Session{}, ErrNotValidSessionToken
}

// DeleteSessionBySelector Delete a session by using a selector
func DeleteSessionBySelector(selector string) error {
	hashSelector := hash.SHA1([]byte(selector))
	_, err := r.Del(SessionPrefix + hashSelector).Result()

	return err
}

// ResetTimeSession Reset time of a session based on config.ini Expire time
func (s Session) ResetTimeSession() error {
	hashSelector := hash.SHA1([]byte(s.Selector))

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
func InitiateSession(uid int) (selector string, validator string, err error) {
	s := random.RandString(12)
	v := random.RandString(50)
	hashS := hash.SHA1([]byte(s))

	err = r.HMSet(SessionPrefix+hashS, map[string]interface{}{
		"userid":    uid,
		"validator": hash.SHA1([]byte(v)),
		"createdat": time.Now().Unix(),
	}).Err()
	if err != nil {
		return "", "", err
	}

	err = r.SAdd(SessionGroupPrefix+strconv.Itoa(uid), hashS).Err()
	if err != nil {
		return "", "", err
	}

	err = Session{
		Selector: s,
		UserID:   uid,
	}.ResetTimeSession()
	if err != nil {
		return "", "", err
	}

	return s, v, nil
}
