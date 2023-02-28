package bot

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func (b *Bot) handleCalcScoreCommand(ctx telebot.Context) error {
	var fst, snd float64
	ch := make(chan bool)

	ctx.Send("Введите баллы за 1-ую половину триместра (1st Attestation):")
	getPercentageResponse(&fst, ctx, ch)

	<-ch // Wait for user's response

	ctx.Send("Введите баллы за 2-ую половину триместра (2nd Attestation):")
	getPercentageResponse(&snd, ctx, ch)

	<-ch // Wait for user's response

	firstAttestationScore := fst * 0.3
	secondAttestationScore := snd * 0.3
	finalScoreNeededForScholarship := (70 - (firstAttestationScore + secondAttestationScore)) / 0.4
	finalScoreNeededForIncreasedScholarship := (90.5 - (firstAttestationScore + secondAttestationScore)) / 0.4
	scoreIf100 := (fst+snd)*0.3 + 40

	MapValue(&finalScoreNeededForScholarship, 50, 100)
	MapValue(&finalScoreNeededForIncreasedScholarship, 50, 100)

	if (fst+snd)/2 < 50 {
		ctx.Send("🔴 GG  Retake")
		return nil
	}

	ctx.Send(fmt.Sprintf("🔴 Для сохранения стипендии (>70)  \n %.2f%% на файнале.", finalScoreNeededForScholarship))

	if finalScoreNeededForIncreasedScholarship > 100 {
		ctx.Send("🔵 Невозможно набрать 90> баллов.")
	} else {
		ctx.Send(fmt.Sprintf("🔵 Для получения повышенной стипендии (>90)  \n %.2f%% на файнале.", finalScoreNeededForIncreasedScholarship))
	}

	ctx.Send(fmt.Sprintf("⚪️ Если получишь на файнале 100 : \n %.2f%%.", scoreIf100))

	return nil
}

func (b *Bot) handleStartCommand(ctx telebot.Context) error {
	ctx.Send(fmt.Sprintf("Hello, %s! Welcome to my bot.", ctx.Sender().FirstName), b.menus.HomeMenu)
	return nil
}
func (b *Bot) HandleHelpCommand(ctx telebot.Context) error {
	ctx.Send(fmt.Sprintf("🤖 I can calculate you scores.\n"), b.menus.HelpMenu)

	return nil
}
