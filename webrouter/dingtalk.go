package webrouter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"

	"github.com/timonwong/prometheus-webhook-dingtalk/config"
	"github.com/timonwong/prometheus-webhook-dingtalk/models"
	"github.com/timonwong/prometheus-webhook-dingtalk/template"
)

type DingTalkResource struct {
	Logger     log.Logger
	Template   *template.Template
	Targets    map[string]config.Target
	HttpClient *http.Client
}

func (rs *DingTalkResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/{name}/send", rs.SendNotification)
	return r
}

func (rs *DingTalkResource) SendNotification(w http.ResponseWriter, r *http.Request) {
	logger := rs.Logger
	targetName := chi.URLParam(r, "name")
	target, ok := rs.Targets[targetName]
	if !ok {
		level.Warn(logger).Log("msg", fmt.Sprintf("target %s not found", targetName))
		http.NotFound(w, r)
		return
	}

	if target.URL == "" {
		level.Warn(logger).Log("msg", fmt.Sprintf("target %s url is empty", targetName))
		http.NotFound(w, r)
		return
	}

	var promMessage models.WebhookMessage
	if err := json.NewDecoder(r.Body).Decode(&promMessage); err != nil {
		level.Error(logger).Log("msg", "Cannot decode prometheus webhook JSON request", "err", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	notification, err := rs.buildDingTalkNotification(&target, &promMessage)
	if err != nil {
		level.Error(logger).Log("msg", "Failed to build notification", "err", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	robotResp, err := rs.sendDingTalkNotification(&target, notification)
	if err != nil {
		level.Error(logger).Log("msg", "Failed to send notification", "err", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if robotResp.ErrorCode != 0 {
		level.Error(logger).Log("msg", "Failed to send notification to DingTalk", "respCode", robotResp.ErrorCode, "respMsg", robotResp.ErrorMessage)
		http.Error(w, "Unable to talk to DingTalk", http.StatusBadRequest)
		return
	}

	io.WriteString(w, "OK") // nolint: errcheck
}

func (rs *DingTalkResource) buildDingTalkNotification(target *config.Target, m *models.WebhookMessage) (*models.DingTalkNotification, error) {
	title, err := rs.Template.ExecuteTextString(`{{ template "ding.link.title" . }}`, m)
	if err != nil {
		return nil, err
	}
	content, err := rs.Template.ExecuteTextString(`{{ template "ding.link.content" . }}`, m)
	if err != nil {
		return nil, err
	}

	notification := &models.DingTalkNotification{
		MessageType: "markdown",
		Markdown: &models.DingTalkNotificationMarkdown{
			Title: title,
			Text:  content,
		},
	}

	// Build mention
	if target.Mention != nil {
		notification.At = &models.DingTalkNotificationAt{
			IsAtAll:   target.Mention.All,
			AtMobiles: target.Mention.Mobiles,
		}
	}

	return notification, nil
}

func (rs *DingTalkResource) sendDingTalkNotification(target *config.Target, notification *models.DingTalkNotification) (*models.DingTalkNotificationResponse, error) {
	targetURL := target.URL
	// Calculate signature when secret is provided
	if target.Secret != "" {
		timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
		stringToSign := []byte(timestamp + "\n" + target.Secret)

		mac := hmac.New(sha256.New, []byte(target.Secret))
		mac.Write(stringToSign) // nolint: errcheck
		signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

		u, err := url.Parse(targetURL)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse target url")
		}

		qs := u.Query()
		qs.Set("timestamp", timestamp)
		qs.Set("sign", signature)
		u.RawQuery = qs.Encode()

		targetURL = u.String()
	}

	body, err := json.Marshal(&notification)
	if err != nil {
		return nil, errors.Wrap(err, "error encoding DingTalk request")
	}

	httpReq, err := http.NewRequest("POST", targetURL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building DingTalk request")
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := rs.HttpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "error sending notification to DingTalk")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("unacceptable response code %d", resp.StatusCode)
	}

	var robotResp models.DingTalkNotificationResponse
	enc := json.NewDecoder(resp.Body)
	if err := enc.Decode(&robotResp); err != nil {
		return nil, errors.Wrap(err, "error decoding response from DingTalk")
	}

	return &robotResp, nil
}
