#yaml-conf

## ДЗ №3
## Модуль для чтения конфигурации из YAML-файла в структуру с помощью библиотеки [go-yaml/yaml](https://github.com/go-yaml/yaml)

Используется в https://github.com/mulla159/gb-go-step-2/tree/main/homework3

1. Создал новый [проект](https://github.com/mulla159/yaml-conf) с использованием инструментария go mod.
```Bash
go mod init github.com/mulla159/yaml-conf
go mod tidy
```
2. Опубликовал [проект](https://github.com/mulla159/yaml-conf) в репозитории, установив номер версии, указывающий на активный этап разработки библиотеки.
```Bash
git init
git log
git tag -a v0.0.1 515b194e073a06846e09f711ea5fb78d21a57394 -m init
git push origin --tags
```
3. добавил для теста пару модулей;
```Bash
go get github.com/valyala/fasthttp
go get github.com/gorilla/websocket
go mod tidy
git tag -a v0.0.2 515b194e073a06846e09f711ea5fb78d21a57394 -m "fix"
git push origin --tags
```
4. Сделал изменения в проекте, запушил их с мажорным обновлением версии пакета.
* удалил импорт библиотек fasthttp и websocket;
* удалил один из возвращаемых параметров структуры с конфигом;
* переименовал один из возвращаемых параметров структуры с конфигом;
```Bash
git tag -a 1.0.0 -m "изменил конфиг"
git push origin --tags
```
5. Очистил неиспользуемые библиотеки
```Bash
go mod tidy
git tag -a v1.0.1 -m "убрал неиспользуемые библиотеки"
git push origin --tags
```
