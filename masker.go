package masker

import (
	"strings"
)

// Masker структура для маскировки строк
type Masker struct{}

// mask заменяет символы в строке str на mask с позиции start до позиции end
func (m *Masker) mask(str string, overlay string, start int, end int) string {
	r := []rune(str)
	l := len(r)

	if l == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}
	if start > end {
		start, end = end, start
	}

	return string(r[:start]) + overlay + string(r[end:])
}

// maskExceptFirstLast оставляет только первый и последний символ
// использовать только со строками из ASCII символов (https://ru.wikipedia.org/wiki/ASCII)
func (m *Masker) maskExceptFirstLast(str string, overlay string, needLast bool) string {
	bytes := make([]byte, len(str))
	for i, k := range []byte(str) {
		if i == 0 || (i == len(str)-1 && needLast) {
			bytes[i] = k
			continue
		}
		bytes[i] = []byte(overlay)[0]
	}
	return string(bytes)
}

// maskExceptLastDigits оставляет count последних символов
// использовать только со строками из ASCII символов (https://ru.wikipedia.org/wiki/ASCII)
func (m *Masker) maskExceptLastDigits(str string, overlay string, count int) string {
	bytes := make([]byte, len(str))
	for i, k := range []byte(str) {
		if i >= len(str)-count {
			bytes[i] = k
			continue
		}
		bytes[i] = []byte(overlay)[0]
	}
	return string(bytes)
}

// Name маскирует второй и третий символ в строке. Может работать со строками из нескольких слов
func (m *Masker) Name(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}

	// если есть пробел
	i = strings.Trim(i, " ")
	if chunks := strings.Split(i, " "); len(chunks) > 1 {
		tmp := make([]string, 0, l)
		for _, v := range chunks {
			if v != "" {
				tmp = append(tmp, m.Name(v))
			}
		}
		return strings.Join(tmp, " ")
	}

	if l == 2 || l == 3 {
		return m.mask(i, "*", 1, 2)
	}

	if l > 3 {
		return m.mask(i, "**", 1, 3)
	}

	return "*"
}

// CreditCard маскирует 6 символов номера кредитной карты начиная с 7
func (m *Masker) CreditCard(i string) string {
	if len(i) == 0 {
		return ""
	}
	return m.mask(i, "******", 6, 12)
}

// Email маскирует логин в email, оставляя домен
func (m *Masker) Email(i string) string {
	if len(i) == 0 {
		return ""
	}

	chunks := strings.Split(i, "@")

	return m.mask(chunks[0], "****", 3, 7) + "@" + chunks[1]
}

// Mobile маскирует 11 значный номер телефона с ведущей цифрой 7
// оставляет первые 4 и последние 2 цифры
func (m *Masker) Mobile(i string) string {
	if len(i) == 0 {
		return ""
	}
	return m.mask(i, "*****", 4, 9)
}

// Password маскирует пароль заданной маской
func (m *Masker) Password(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}
	return "************"
}

// PassportSeries маскирует серию паспорта, оставляя первую и последнюю цифры
func (m *Masker) PassportSeries(i string) string {
	if len(i) == 0 {
		return ""
	}
	if len(i) < 4 {
		return "****"
	}
	return m.maskExceptFirstLast(i, "*", true)
}

// PassportNumber маскирует номер паспорта, оставляя первую и последнюю цифру
func (m *Masker) PassportNumber(i string) string {
	if len(i) == 0 {
		return ""
	}
	if len(i) < 6 {
		return "******"
	}
	return m.maskExceptFirstLast(i, "*", true)
}

// Code маскирует код из цифр
// для кодов состоящих из меньше чем 4 символов, оставляет первый символ
// для кодов большей длины оставляет первый и последний символ
func (m *Masker) Code(i string) string {
	if len(i) == 0 {
		return ""
	}

	if len(i) == 1 {
		return "*"
	}

	if len(i) < 4 {
		return m.maskExceptFirstLast(i, "*", false)
	}

	return m.maskExceptFirstLast(i, "*", true)
}

//LastFourDigits маскирует любую последовательность больше 5 символов, оставляя 4 последних символа
func (m *Masker) LastFourDigits(i string) string {
	if len(i) == 0 {
		return ""
	}

	if len(i) < 5 {
		return "****"
	}

	return m.maskExceptLastDigits(i, "*", 4)
}

// New создает Masker
func New() *Masker {
	return &Masker{}
}

var instance *Masker

func init() {
	instance = New()
}

func Name(i string) string {
	return instance.Name(i)
}

func CreditCard(i string) string {
	return instance.CreditCard(i)
}

func Email(i string) string {
	return instance.Email(i)
}

func Mobile(i string) string {
	return instance.Mobile(i)
}

func Password(i string) string {
	return instance.Password(i)
}

func PassportSeries(i string) string {
	return instance.PassportSeries(i)
}

func PassportNumber(i string) string {
	return instance.PassportNumber(i)
}

func Code(i string) string {
	return instance.Code(i)
}

func LastFourDigits(i string) string {
	return instance.LastFourDigits(i)
}
