package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	// запускаем горутину для вызова
	go func(Input <-chan int, Output chan<- int) {
		for {
			x, ok := <-Input // смотрим можно ли из канала что то посчитать
			if ok {
				Output <- x * x // если можно, то добавляем ответ в выходной канал
			} else {
				break // иначе выходим, числа закончились
			}
		}
		close(Output) // и закрываем выходной поток
	}(c.Input, c.Output)
}
