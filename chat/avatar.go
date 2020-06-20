package main

import (
	"errors"
)

// ErrNoAvatarURL is error that when Avatar instance can not return Avatar URL
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

// Avatar is type of user profile image
type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
