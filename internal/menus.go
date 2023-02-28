package internal

import "gopkg.in/telebot.v3"

type Menus struct {
	HomeMenu *telebot.ReplyMarkup
	HelpMenu *telebot.ReplyMarkup
}

func NewMenus() *Menus {
	homeMenu := newHomeMenu()
	helpMenu := newHelpMenu()

	return &Menus{HomeMenu: homeMenu, HelpMenu: helpMenu}
}

func newHomeMenu() *telebot.ReplyMarkup {
	homeMenu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	calcScoreBtn := homeMenu.Text("Рассчитать нужное количество баллов на файнале")
	helpBtn := homeMenu.Text("/help")
	homeMenu.Reply(
		homeMenu.Row(calcScoreBtn),
		homeMenu.Row(helpBtn),
	)
	return homeMenu
}

func newHelpMenu() *telebot.ReplyMarkup {
	helpMenu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	calcScoreBtn := helpMenu.Text("Рассчитать нужное количество баллов на файнале")
	helpMenu.Inline(
		helpMenu.Row(calcScoreBtn),
	)
	return helpMenu
}
