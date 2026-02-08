package route

import "main/handlers"

func StartRoutesInit(h handlers.StartHandlers) {
	h.Bot.Handle("/start", h.StartMessage)
}
