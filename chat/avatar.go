package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
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

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	// Gravatarのガイドラインに従った処理
	// メールアドレスに含まれる大文字を小文字に変換
	// 結果の文字列をMD5を用いてハッシュ値を算出
	// ハッシュ値をGravatarのURLに埋め込む
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			fmt.Println(email)
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}
