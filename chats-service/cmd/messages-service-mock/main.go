package main

import (
	"encoding/json"
	"errors"
	"fmt"
	api "go13/pkg/ogen/messages-service"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func init() {
	mIdStr := os.Getenv("MESSAGE_ID")
	mId, err := strconv.Atoi(mIdStr)
	if err != nil {
		mId = 42
	}
	messageId = mId

	chIdStr := os.Getenv("CHAT_ID")
	chId, err := strconv.Atoi(chIdStr)
	if err != nil {
		chId = 42
	}
	chatId = chId

	sId := os.Getenv("SENDER_ID")
	if sId == "" {
		sId = "da4733f1-d3d1-4f96-ab31-c1ab5259b314"
	}
	senderId = sId
}

var (
	messageId int
	chatId    int
	senderId  string
	message   = api.Message{
		Message:       "test message",
		Edited:        false,
		SendTimestamp: 1516239022,
	}
)

type jwtClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
}

func main() {
	fmt.Println(messageId, chatId, senderId)
	message.ID = api.MessageId(messageId)
	message.SenderID = api.UserId(senderId)
	mux := http.NewServeMux()
	mux.HandleFunc("/messages/{id}", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		authS := strings.Split(auth, " ")
		if len(authS) < 2 || authS[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := authS[1]
		var claims jwtClaims
		_, err := jwt.ParseWithClaims(token, &claims, nil)
		if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId := claims.UserId
		if _, err := uuid.Parse(userId); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		chatIdString := r.URL.Query().Get("chatId")
		gotChatId, err := strconv.Atoi(chatIdString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if id != messageId || gotChatId != chatId {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Add("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", mux) //nolint:errcheck
}
