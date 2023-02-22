package bot

import (
	telebot "gopkg.in/telebot.v3"
	"time"
)

type Bot struct {
	bot *telebot.Bot
}

func NewBot(token string) (*Bot, error) {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	return &Bot{bot: b}, nil
}

func (b *Bot) HandleCommand(command string, handler func(ctx telebot.Context)) {
	b.bot.Handle(command, func(ctx telebot.Context) error {
		handler(ctx)
		return nil
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
	b.HandleCommand("Рассчитать нужное количество баллов на файнале", b.handleCalcRegCommand)
	b.HandleMessage(func(ctx telebot.Context) {
		b.SendMessage(ctx.Chat().ID, "I'm sorry, I don't understand.")
	})
}

func (b *Bot) Start() {
	b.bot.Start()
}
