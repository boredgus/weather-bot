package commands

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Message = tgbotapi.Chattable

type Replyable interface {
	Reply() (Message, error)
}

const MarkdownV2Mode = "MarkdownV2"

func ToJson(name, description string) string {
	return fmt.Sprintf(`{"command":"%v","description":"%v"}`, name, description)
}

func parseText(str string) string {
	str = strings.ReplaceAll(str, "-", "\\-")
	str = strings.ReplaceAll(str, ".", "\\.")
	str = strings.ReplaceAll(str, "(", "\\(")
	str = strings.ReplaceAll(str, ")", "\\)")
	return strings.ReplaceAll(str, "+", "\\+")
}
