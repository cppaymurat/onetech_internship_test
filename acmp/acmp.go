package acmp

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Difficulty(url string) float64 {
	client := http.Client{}
	request, ok := http.NewRequest("GET", url, nil)
	if ok != nil {
		return -1
	}
	request.AddCookie(
		&http.Cookie{
			Name:  "English",
			Value: "1"}) // устанавливаем куки, чтоб при обработке string незафакапиться при виде кириллицы
	response, ok := client.Do(request) // делаем запрос
	if ok != nil {
		return -1 // если неудачно, сворачиваемся
	}
	asBytes, _ := ioutil.ReadAll(response.Body)
	asString := string(asBytes) // переводим тело в текстовый вид
	toSearch := "Difficulty: "  // ищем оттуда ключевое слово сложность
	index := strings.Index(asString, toSearch)
	result := float64(0) // здесь будем результат хранить
	if index == -1 {
		return -1 // если не смогли найти ключевое слово, то возвращаем -1
	} else {
		for i := index + len(toSearch); asString[i] != '%'; i += 1 { // пока не увидим процент, будем все конвертировать в число
			result = result*10 + float64(asString[i]-'0')
		}
	}
	return result // возвращаем результат
}
