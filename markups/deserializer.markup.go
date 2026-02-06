package markups

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/telebot.v3"
)

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

func deserialize(fileName string) *[]Element {
	data := getFile("start")
	elements := []Element{}
	err := json.Unmarshal(data, &elements)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &elements
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

func SetMarkupData(menu *telebot.ReplyMarkup, fileName string) {
	deserializedMarkup := deserialize(fileName)

	if deserializedMarkup == nil {
		return
	}
	menuRows := make([]telebot.Row, 0, 0)
	for _, e := range *deserializedMarkup {
		if e.Cmp == "group" {
			menuBtn := make([]telebot.Btn, 0, 0)
			for _, ge := range e.Group {
				menuBtn = append(menuBtn, menu.Data(ge.Name, ge.Button.Button))
			}

			menuRows = append(menuRows, menu.Row(menuBtn...))
		}
		menuRows = append(menuRows, menu.Row(menu.Data(e.Name, e.Button.Button)))
	}

	menu.Inline(menuRows...)
}
