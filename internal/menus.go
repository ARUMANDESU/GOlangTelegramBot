package internal

import "gopkg.in/telebot.v3"

type Menus struct {
	HomeMenu *telebot.ReplyMarkup
}

func NewMenus() *Menus {
	homeMenu := newHomeMenu()

	return &Menus{HomeMenu: homeMenu}
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
