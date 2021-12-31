package discord

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Int63n(16777215))
	}
}

func TestSendNotifyDiscordContent(t *testing.T) {
	webhook := "https://discord.com/api/webhooks/851721788760391700/JLeqEJkhH1AgGr9rkhJ0FEvcdwqPbZWpoopZcYwdxgDeM91pnhI2nruUncBw-Qo874_k"
	title := fmt.Sprintf("** %s ** bought for %s CKB($%s).", "xxxx.bit", "100,000,102", "189,834.2302")
	err := SendNotifyDiscordContent(webhook, title)
	if err != nil {
		t.Fatal()
	}
}

func TestSendNotifyDiscordEmbeds(t *testing.T) {
	webhook := "https://discord.com/api/webhooks/851721788760391700/JLeqEJkhH1AgGr9rkhJ0FEvcdwqPbZWpoopZcYwdxgDeM91pnhI2nruUncBw-Qo874_k"
	url := "https://551c3fbbfa.bestdas.com/stats"
	title := fmt.Sprintf("** %s ** bought for %s CKB($%s).", "xxxx.bit", "100,000,102", "189,834.2302")

	embed := Embed{
		Title:       title,
		URL:         url,
		Description: "what a awesome message",
		Color:       rand.Int63n(16777215),
	}
	err := SendNotifyDiscordEmbeds(webhook, embed)
	if err != nil {
		t.Fatal(err)
	}
}
