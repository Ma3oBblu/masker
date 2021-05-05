# Masker

[![Build Status](https://travis-ci.org/ma3obblu/masker.svg?branch=master)](https://travis-ci.org/ma3obblu/masker)
[![codecov](https://codecov.io/gh/ma3obblu/masker/branch/master/graph/badge.svg)](https://codecov.io/gh/ma3obblu/masker)
[![Go Report Card](https://goreportcard.com/badge/github.com/ma3obblu/masker)](https://goreportcard.com/report/github.com/ma3obblu/masker)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/ma3obblu/masker/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/ma3obblu/masker?status.svg)](https://godoc.org/github.com/ma3obblu/masker)
[![Release](https://img.shields.io/github/release/ma3obblu/masker.svg?style=flat-square)](https://github.com/ma3obblu/masker/releases/latest)

Masker простая утилита для маскировки чувствительных данных.

# Установка

```
$ go get -u github.com/ma3obblu/masker
```

# Примеры

Два способа работы с masker:

#### 1. Получить инстанс напрямую из пакета masker

``` golang
package main

import (
	masker "github.com/ma3obblu/masker"
)

func main() {
	masker.Name("Руслан")
	masker.Mobile("79191232323")
}
```

#### 2. Получить инстанс используя конструктор `masker.New()`

``` golang
package main

import (
	masker "github.com/ma3obblu/masker"
)

func main() {
	m := masker.New()
	m.Name("Руслан")
	m.Mobile("79191232323")
}
```

## Типы полей

|Type        |Description                                                                                            |
|:----------:|:------------------------------------------------------------------------------------------------------|
|Name|маскирует второй и третий символ в строке. Может работать со строками из нескольких слов|                                                          |
|CreditCard|маскирует 6 символов номера кредитной карты начиная с 7|
|Email|маскирует логин в email, оставляя домен|
|Mobile|маскирует 11 значный номер телефона с ведущей цифрой 7, оставляет первые 4 и последние 2 цифры|
|Password|всегда вернет "************"|
|PassportSeries|маскирует серию паспорта, оставляя первую и последнюю цифры|
|PassportNumber|маскирует номер паспорта, оставляя первую и последнюю цифру|
|Code|маскирует код из цифр. Для кодов состоящих из меньше чем 4 символов, оставляет первый символ. Для кодов большей длины оставляет первый и последний символ|
|LastFourDigits|маскирует любую последовательность больше 5 символов, оставляя 4 последних символа|