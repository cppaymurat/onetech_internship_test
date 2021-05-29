package quicksort

func Partition(a []int) int {
	left := 0
	right := len(a) - 1
	pivot := a[right] // выбираем любой элемент как опорный, совсем неважно какой, ведь
	// в среднем случае массив будет неотсортированный изначально
	for {
		for a[left] < pivot {
			left += 1 // пропускаем хорошие элементы слева
		}
		for a[right] > pivot {
			right -= 1 // пропускаем хорошие справа
		}
		if left >= right {
			return left // если и слева и справа все хорошие, и между левым и правым индексами больше не осталось элементов
			// то можем выбрать в качестве опорного любой элемент среди этих
		}
		a[left], a[right] = a[right], a[left] // а если между ними есть элементы, то факт:
		// a[left] >= pivot, a[right] <= pivot, оба не подходят, меняем местами эти элементы, теперь все ОК
		left += 1 // идем дальше
		right -= 1
	}
}

func QuickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	pivot := Partition(a) // запрашиваем кто опорный
	QuickSort(a[:pivot])  // сортируем срез с индексами 0 ... pivot - 1
	QuickSort(a[pivot:])  // сортируем срез с индексами pivot ... len(a) - 1
}
