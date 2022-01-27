package main

import (
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/XiaoMengXinX/ProfileStatusSyncer/gh"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	c, err := gh.NewClient(os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	ghStatus, err := c.GetUserStatus(c.Login)
	if err != nil {
		log.Errorln(err)
	}
	data := utils.RequestData{Cookies: []*http.Cookie{{Name: "MUSIC_U", Value: os.Getenv("MUSIC_U")}}}
	result, err := api.SetUserStatus(data, fmt.Sprintf("%s%s", gh.Emojis.Shortname2Emoji(ghStatus.Data.User.Status.Emoji), ghStatus.Data.User.Status.Message))
	if err != nil {
		log.Errorln(err)
	}
	fmt.Println(result.Message)
	fmt.Println("Sync Success")
}
