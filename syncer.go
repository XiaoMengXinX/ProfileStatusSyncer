package main

import (
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/XiaoMengXinX/ProfileStatusSyncer/gh"
	"github.com/rivo/uniseg"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type logFormatter struct{}

// Format is a formatter for logrus
func (s *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	var msg string
	msg = fmt.Sprintf("%s [%s] %s (%s:%d)\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message, path.Base(entry.Caller.File), entry.Caller.Line)
	return []byte(msg), nil
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	log.SetFormatter(new(logFormatter))
	log.SetReportCaller(true)
}

var githubToken = os.Getenv("GITHUB_TOKEN")
var neteaseCookie = os.Getenv("MUSIC_U")
var mode = os.Getenv("MODE")

const (
	// Github2NeteaseMode is to sync GitHun profile status to NeteaseCloud Music
	Github2NeteaseMode = "GitHub2Netease"
	// Netease2GithubMode is to sync NeteaseCloud Music profile status to GitHun
	Netease2GithubMode = "Netease2GitHub"
	// KeepNeteaseStatusMode is to maintain the profile status of NeteaseCloud Music
	KeepNeteaseStatusMode = "KeepNeteaseStatus"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Errorln(err)
		}
	}()

	if githubToken == "" && neteaseCookie == "" {
		err = fmt.Errorf("GITHUB_TOKEN or MUSIC_U is empty ")
		return
	}

	c, err := gh.NewClient(githubToken)
	if err != nil {
		return
	}

	n := utils.RequestData{Cookies: []*http.Cookie{{Name: "MUSIC_U", Value: neteaseCookie}}}
	loginStatus, err := api.GetLoginStatus(n)
	if err != nil {
		return
	}
	if loginStatus.Profile.UserId == 0 {
		err = fmt.Errorf("Check NeteaseCloud Music login status failed ")
		return
	}
	userID := loginStatus.Profile.UserId

	switch mode {
	case Github2NeteaseMode:
		err = syncGitHub2Netease(c, n)
	case Netease2GithubMode:
		err = syncNetease2GitHub(c, n, userID)
	case KeepNeteaseStatusMode:
		err = keepNeteaseStatus(n, userID)
	default:
		err = syncGitHub2Netease(c, n)
	}

	if err == nil {
		log.Println("Sync or update profile status successfully")
	}
}

func syncGitHub2Netease(c *gh.Client, n utils.RequestData) (err error) {
	status, err := c.GetUserStatus(c.Login)
	if err != nil {
		return err
	}
	if status.Data.User.Status.Emoji == "" && status.Data.User.Status.Message == "" {
		return fmt.Errorf("GitHub status is empty ")

	}
	_, err = api.SetUserStatus(n, fmt.Sprintf("%s %s", gh.Emojis.Shortname2Emoji(status.Data.User.Status.Emoji), status.Data.User.Status.Message))
	if err != nil {
		return err
	}
	return err
}

func syncNetease2GitHub(c *gh.Client, n utils.RequestData, userID int) (err error) {
	status, err := api.GetUserStatus(n, userID)
	if err != nil {
		return err
	}
	if status.Data.Content.Content == "" {
		return fmt.Errorf("NeteaseCloud Music profile status is empty ")
	}
	gr := uniseg.NewGraphemes(status.Data.Content.Content)
	var emoji gh.Emoji
	for gr.Next() {
		if emoji = gh.Emojis.GetEmoji(gr.Str()); emoji.Shortname != "" {
			break
		}
	}
	err = c.SetUserStatus(emoji.Emoji, strings.Replace(status.Data.Content.Content, emoji.Emoji, "", 1))
	if err != nil {
		return err
	}
	return err
}

func keepNeteaseStatus(n utils.RequestData, userID int) (err error) {
	status, err := api.GetUserStatus(n, userID)
	if err != nil {
		return err
	}
	if status.Data.Content.Content == "" {
		return fmt.Errorf("NeteaseCloud Music profile status is empty ")
	}
	_, err = api.SetUserStatus(n, status.Data.Content.Content)
	return err
}
