package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yude/youbine/utils"
)

type payload struct {
	Text string `json:"content"`
}

var webhook_url string

func Initialize() {
	webhook_url = utils.GetEnv("DISCORD_WEBHOOK_URL", "")
}

func Post(content string) error {
	if (webhook_url=="") {
		return nil
	}

	p, err := json.Marshal(payload{Text: content})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", webhook_url, bytes.NewBuffer(p))
	if err != nil {
		fmt.Println("Webhook post error!:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Webhook request error!:", err)
		return err
	}
	if resp.StatusCode != 204 {
		fmt.Printf("Status code:%#v\n", resp)
	}
	return nil
}
