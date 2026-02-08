package markups

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/telebot.v3"
)

type Markup struct {
	Message string    `json:"message,omitempty"`
	Buttons []Element `json:"buttons,omitempty"`
}

type Button struct {
	Button string `json:"button,omitempty"`
	Name   string `json:"name,omitempty"`
}
type Element struct {
	Cmp   string      `json:"cmp,omitempty"`
	Group []GroupItem `json:"group,omitempty"`
	Button
}

type GroupItem struct {
	Button
}

func deserialize(fileName string) (string, *[]Element) {
	data := getFile(fileName)
	markup := Markup{}
	err := json.Unmarshal(data, &markup)
	if err != nil {
		fmt.Println(err.Error())
		return "", nil
	}
	return markup.Message, &markup.Buttons
}

func getFile(fileName string) []byte {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/markups/%s/markup.json", currentDir, fileName))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return file
}

func SetMarkupData(menu *telebot.ReplyMarkup, fileName string) (string, map[string]*telebot.Btn) {
	systemButtons := make(map[string]*telebot.Btn)
	message, deserializedMarkup := deserialize(fileName)

	if deserializedMarkup == nil {
		return "", nil
	}
	menuRows := make([]telebot.Row, 0, 0)
	for _, e := range *deserializedMarkup {
		if e.Cmp == "group" {
			menuBtn := make([]telebot.Btn, 0, 0)
			for _, ge := range e.Group {
				button := menu.Data(ge.Name, ge.Button.Button)
				systemButtons[ge.Button.Button] = &button
				menuBtn = append(menuBtn, button)
			}

			menuRows = append(menuRows, menu.Row(menuBtn...))
		}
		button := menu.Data(e.Name, e.Button.Button)

		systemButtons[e.Button.Button] = &button
		menuRows = append(menuRows, menu.Row(button))
	}

	menu.Inline(menuRows...)
	return message, systemButtons
}
