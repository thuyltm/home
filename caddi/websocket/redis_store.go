package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
)

type RedisStore struct {
	Codecs      []securecookie.Codec
	Options     *sessions.Options
	RedisClient redis.Client
}

func NewRedisStore(keyPairs ...[]byte) *RedisStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rs := &RedisStore{
		Codecs: securecookie.CodecsFromPairs(keyPairs...),
		Options: &sessions.Options{
			Path:   "/",
			MaxAge: 60 * 60, //1hour
		},
		RedisClient: *redisClient,
	}
	return rs
}

var USERID = "userID"

func (s *RedisStore) Get(r *http.Request) (*goth.User, error) {
	var userId string
	if c, errCookie := r.Cookie(USERID); errCookie == nil {
		securecookie.DecodeMulti(USERID, c.Value, &userId, s.Codecs...)
	} else {
		return nil, fmt.Errorf("no user info")
	}
	var user goth.User
	values, err := s.RedisClient.HGetAll(context.TODO(), userId).Result()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user.Name = values["name"]
	user.UserID = values["id"]
	return &user, nil
}

func (s *RedisStore) Save(r *http.Request, w http.ResponseWriter, gothUser goth.User) error {
	var err error
	key := "user:" + gothUser.UserID
	s.RedisClient.Expire(context.TODO(), key, time.Duration(s.Options.MaxAge)*time.Minute)
	_, err = s.RedisClient.HSet(context.TODO(), key,
		"UserID", gothUser.UserID,
		"Name", gothUser.Name).Result()
	if err != nil {
		return err
	}
	encoded, err := securecookie.EncodeMulti(USERID, key, s.Codecs...)
	if err != nil {
		return err
	}
	http.SetCookie(w, sessions.NewCookie(USERID, encoded, s.Options))
	return nil
}
