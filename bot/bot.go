package bot

import (
	telebot "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"os"
	"strings"
	"telegrambot/internal"
)

type Bot struct {
	bot   *telebot.Bot
	menus *internal.Menus
}

func NewBot(token string) (*Bot, error) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received webhook request: %v", r)
	})
	go http.ListenAndServe(":8080", nil)

	webhook := &telebot.Webhook{
		Listen:   "127.0.0.1:8443",
		Endpoint: &telebot.WebhookEndpoint{PublicURL: os.Getenv("PUBLIC_URL")},
	}

	spamProtected := telebot.NewMiddlewarePoller(webhook, func(upd *telebot.Update) bool {
		if upd.Message == nil {
			return true
		}

		if strings.Contains(upd.Message.Text, "spam") {
			return false
		}

		return true
	})

	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: spamProtected,
	})
	if err != nil {
		return nil, err
	}
	menus := internal.NewMenus()
	return &Bot{bot: b, menus: menus}, nil
}

func (b *Bot) HandleCommand(command string, handler func(ctx telebot.Context) error) {
	b.bot.Handle(command, func(ctx telebot.Context) error {
		return handler(ctx)
	})
}

// HandleMessage sets up a message handler for the bot.
func (b *Bot) HandleMessage(handler func(ctx telebot.Context)) {
	b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		handler(ctx)
		return nil
	})
}

// SendMessage sends a message to the specified chat.
func (b *Bot) SendMessage(chatID int64, text string) error {
	_, err := b.bot.Send(&telebot.User{ID: chatID}, text)
	return err
}

func (b *Bot) AddHandlers() {

	b.HandleCommand("/start", b.handleStartCommand)
	b.HandleCommand("/help", b.HandleHelpCommand)
	b.HandleCommand("Рассчитать нужное количество баллов на файнале", b.handleCalcScoreCommand)

	b.HandleMessage(func(ctx telebot.Context) {
		b.SendMessage(ctx.Chat().ID, "I'm sorry, I don't understand.")
	})
}

func (b *Bot) Start() {
	b.bot.Start()
}
