package bot

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func (b *Bot) handleCalcRegCommand(ctx telebot.Context) {
	var fst, snd float64
	ch := make(chan bool)

	ctx.Send("Введите баллы за 1-ую половину триместра (1st Attestation):")
	getPercentageResponse(&fst, ctx, ch)

	<-ch // Wait for user's response

	ctx.Send("Введите баллы за 2-ую половину триместра (2nd Attestation):")
	getPercentageResponse(&snd, ctx, ch)

	<-ch // Wait for user's response
	finalScore := (70 - (fst*0.3 + snd*0.3)) / 0.4
	if (fst+snd)/2 < 50 {
		ctx.Send("GG  Retake")
		return
	}
	ctx.Send(fmt.Sprintf("🔴 Для сохранения стипендии   (>70)  \n %.2f%% на файнале. ", finalScore))
}

func (b *Bot) handleStartCommand(ctx telebot.Context) {
	var (
		menu    = &telebot.ReplyMarkup{ResizeKeyboard: true}
		btnHelp = menu.Text("Рассчитать нужное количество баллов на файнале")
	)

	menu.Reply(
		menu.Row(btnHelp),
	)

	// send the welcome message with the reply markup
	ctx.Send(fmt.Sprintf("Hello, %s! Welcome to my bot.", ctx.Sender().FirstName), menu)
}
