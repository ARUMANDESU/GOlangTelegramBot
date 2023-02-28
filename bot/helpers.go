package bot

import (
	"gopkg.in/telebot.v3"
	"telegrambot/utils"
)

func getPercentageResponse(p *float64, ctx telebot.Context, ch chan bool) {
	ctx.Bot().Handle(telebot.OnText, func(context telebot.Context) error {
		percentage, err := utils.IsPercentage(context.Text())
		if err != nil {
			context.Reply("Некорректный ввод. Пожалуйста, введите проценты (например, 80.5 или 80.5%)")
			return err
		}
		*p = percentage
		ch <- true
		return nil
	})

}

func MapValue(value *float64, minNew, maxNew float64) {
	switch {
	case *value < minNew:
		*value = minNew
	case *value > maxNew:
		*value = maxNew

	}
}
