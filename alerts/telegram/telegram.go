package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/finnan444/utils/math/ints"
	"github.com/finnan444/utils/transport"
	"github.com/valyala/fasthttp"
)

// Pathes
const (
	pathSendMessage    = "/sendMessage"
	pathGetWebhookInfo = "/getWebhookInfo"
	pathSetWebhook     = "/setWebhook"
)

// Config describes telegram config
type Config struct {
	URL      string            `json:"url"`
	Chats    map[string]string `json:"chats"`
	BotID    string            `json:"botId"`
	BotToken string            `json:"botToken"`
}

// Alert is Telegram structure
type Alert struct {
	config  *Config
	timeout time.Duration
	botURL  string
}

type getWebhookInfoResponse struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int      `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       int      `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

// InitTelegram initialize Telegram
func InitTelegram(config *Config, connectionTimeoutSeconds time.Duration) *Alert {
	return &Alert{
		botURL:  fmt.Sprintf(config.URL, config.BotID, config.BotToken),
		timeout: connectionTimeoutSeconds * time.Second,
		config:  config,
	}
}

// PostMessage do a post request to Telegram with message param and using infoLevel
func (mngr *Alert) PostMessage(message string, infoLevel string) error {
	var err error

	chatID, ok := mngr.config.Chats[infoLevel]
	if !ok {
		if _, err = strconv.Atoi(infoLevel); err == nil {
			chatID = infoLevel
		} else {
			return err
		}
	}
	if chatID != "" {
		runes := []rune(message)
		client := transport.GetHTTPClient()
		defer transport.PutHTTPClient(client)

		for i, l := 0, len(runes); i < l; i += 4096 {
			if err := mngr.makeRequest(chatID, string(runes[i:ints.MinInt(l, i+4096)]), client); err != nil {
				return err
			}
		}
	} else {
		log.Printf("[Telegram] Unknown info level %s", infoLevel)
		return fmt.Errorf("[Telegram] Unknown info level %s", infoLevel)
	}

	return nil
}

func (mngr *Alert) makeRequest(chatID string, text string, client *fasthttp.Client) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(mngr.botURL + pathSendMessage)
	req.Header.SetMethod(transport.Post)
	req.Header.SetContentType(transport.ApplicationJSON)
	reqBody := map[string]string{
		"chat_id":    chatID,
		"parse_mode": "Markdown",
		"text":       text,
	}

	reqBytes, err := json.Marshal(&reqBody)
	if err != nil {
		log.Printf("[Telegram] Error marshaling request %v", err)
		return err
	}

	req.SetBody(reqBytes)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = client.DoTimeout(req, resp, mngr.timeout)
	if err != nil {
		log.Printf("[Telegram] error posting message %v", err)
		return err
	} else if resp.StatusCode() != 200 {
		log.Printf("[Telegram] error posting message %s", resp.Body())
		return fmt.Errorf("[Telegram] error posting message %s", resp.Body())
	}

	return nil
}

// RegisterWebhook do smth
func (mngr *Alert) RegisterWebhook(url string) {
	client := transport.GetHTTPClient()
	sc, body, err := client.GetTimeout(nil, mngr.botURL+pathGetWebhookInfo, time.Second*5)
	transport.PutHTTPClient(client)
	if err != nil {
		log.Printf("[Telegram] Error checking webhook info %v", err)
	} else {
		if sc == 200 {
			response := &getWebhookInfoResponse{}
			err = json.Unmarshal(body, response)
			if err != nil {
				log.Printf("[Telegram] Error unmarshalling webhook info response %v", err)
			} else {
				if response.URL == "" || response.URL != url {
					// TODO: delete webhook
					if url != "" {
						mngr.setWebhook(url)
					}
				}
			}
		} else {
			log.Printf("[Telegram] Error quering webhook info code %d\nbody: %s", sc, body)
		}
	}
}

// setWebhook do smth
func (mngr *Alert) setWebhook(url string) {
	client := transport.GetHTTPClient()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(mngr.botURL + pathSetWebhook)
	req.Header.SetMethod(transport.Post)
	req.Header.SetContentType(transport.ApplicationJSON)
	reqBody := map[string]interface{}{
		"url":             url,
		"allowed_updates": []string{"message"},
	}
	reqBytes, err := json.Marshal(&reqBody)
	if err != nil {
		log.Printf("[Telegram] Error marshaling request %v", err)
	} else {
		req.SetBody(reqBytes)
		resp := fasthttp.AcquireResponse()
		err = client.DoTimeout(req, resp, mngr.timeout)
		if err != nil {
			log.Printf("[Telegram] error setting webhook %v", err)
		} else if resp.StatusCode() != 200 {
			log.Printf("[Telegram] error setting webhook %s", resp.Body())
		}
		fasthttp.ReleaseResponse(resp)
	}
	fasthttp.ReleaseRequest(req)
	transport.PutHTTPClient(client)
}
