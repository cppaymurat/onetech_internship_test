package full_outer_join

import (
	"io/ioutil"
	"sort"
	"strings"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	metInFirstFile := make(map[string]bool)
	metInSecondFile := make(map[string]bool)
	firstFile, err := ioutil.ReadFile(f1Path)
	if err != nil {
		panic("Не удалось открыть файл: " + f1Path)
	}
	text := string(firstFile)
	allLines := make([]string, 0)
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i += 1 {
		allLines = append(allLines, lines[i])
		metInFirstFile[lines[i]] = true
	}
	secondFile, err := ioutil.ReadFile(f2Path)
	if err != nil {
		panic("Не удалось открыть файл: " + f2Path)
	}
	text = string(secondFile)
	lines = strings.Split(text, "\n")
	for i := 0; i < len(lines); i += 1 {
		allLines = append(allLines, lines[i])
		metInSecondFile[lines[i]] = true
	}
	// все нужное посчитали, случаи когда файл не открываются учли
	// теперь словари metInFirstFile, metInSecondFile хранят для строки-ключа true, если видели такую строку
	// осталось пройтись по всем строкам, и найти те, что встречаются лишь в одном файле
	resultLines := make([]string, 0)
	for i := 0; i < len(allLines); i += 1 {
		if _, ok := metInFirstFile[allLines[i]]; ok == true { // если в первом есть, а во втором нет
			if _, ok := metInSecondFile[allLines[i]]; ok == false {
				resultLines = append(resultLines, allLines[i])
			}
		} else { // если в первом нет, то ОК
			resultLines = append(resultLines, allLines[i])
		}
	}
	sort.Strings(resultLines)                                        // отсортировали
	resultAsString := strings.Join(resultLines, "\n")                // слили воедино
	er := ioutil.WriteFile(resultPath, []byte(resultAsString), 0644) // 0644 нагуглил, сорри
	if er != nil {                                                   // если не удалось записать, то паникуем
		panic("At the disco! Nicotine!")
	}
}
