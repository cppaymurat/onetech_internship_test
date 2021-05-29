package acmp_concurrent

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
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

func Difficulties(urls []string) map[string]float64 {
	result := make(map[string]float64)
	waitGroups := sync.WaitGroup{}
	mutexOfMap := sync.RWMutex{}
	for _, url := range urls { // проходимся по урлам
		waitGroups.Add(1) // добавляем в вэйтгруппу +1
		go func(url string) {
			value := Difficulty(url) // находим сложность с race condition
			mutexOfMap.Lock()        // только во время записи в мапу лочим мьютекс
			result[url] = value      // записываем
			waitGroups.Add(-1)       // теперь можно не ждать эту горутину
			mutexOfMap.Unlock()      // анлочим мьютекс
		}(url) // с прошлой стажки узнал, что нужно данные передавать как параметры
		// иначе если к локальным переменным обращаться, можно случайно захватить не то, и обработать первый урл для всех горутин
	}
	waitGroups.Wait() // ждем пока горутины завершатся
	return result     // и возвращаем результат, с первого раза кстати прошло
	// да и быстрее в 20 раз
}
