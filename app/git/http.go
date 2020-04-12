package git

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/lissteron/deployer/app/deploy"
	"github.com/lissteron/deployer/pkg/github"
	"github.com/spf13/viper"
)

const (
	maxBodySize = 1024 * 1024
)

func WebhookRequest(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	if viper.GetBool("DEBUG") {
		logger.Println("[debug]", "get request")
		logger.Println("[debug]", "X-GitHub-Event", r.Header.Get("X-GitHub-Event"))
		logger.Println("[debug]", "X-GitHub-Delivery", r.Header.Get("X-GitHub-Delivery"))
		logger.Println("[debug]", "X-Hub-Signature", r.Header.Get("X-Hub-Signature"))
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Println("[error]", err)
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	if viper.GetBool("DEBUG") {
		logger.Println("[debug]", "body")
		logger.Println(string(body))
	}

	if !checkSign(logger, body, r.Header.Get("X-Hub-Signature")) {
		logger.Println("[error]", "bad signature", r.Header.Get("X-Hub-Signature"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	go func() {
		switch r.Header.Get("X-GitHub-Event") {
		case "push":
			event := new(github.PushEvent)

			if err := event.UnmarshalJSON(body); err != nil {
				logger.Println("[error]", err)
				w.WriteHeader(http.StatusBadRequest)

				return
			}

			deploy.ProcessPush(logger, event)
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func checkSign(logger *log.Logger, body []byte, sign string) bool {
	h := hmac.New(sha1.New, []byte(viper.GetString("HMAC_SECRET")))

	if _, err := h.Write(body); err != nil {
		logger.Println("[error]", err)
		return false
	}

	sig := "sha1=" + hex.EncodeToString(h.Sum(nil))

	logger.Println("[debug]", sign, sig)

	return strings.EqualFold(sign, sig)
}
