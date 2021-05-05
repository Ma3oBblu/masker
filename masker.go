package masker

import (
	"strings"
)

const (
	MPassword       = "password"
	MName           = "name"
	MEmail          = "email"
	MMobile         = "mobile"
	MLastFourDigits = "last_four_digits"
	MCreditCard     = "credit_card"
	MPassportSeries = "passport_series"
	MPassportNumber = "passport_number"
	MCode           = "code"
)

type Masker struct{}

func (m *Masker) overlay(str string, overlay string, start int, end int) string {
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

func (m *Masker) String(t string, i string) string {
	switch t {
	default:
		return i
	case MPassword:
		return m.Password(i)
	case MName:
		return m.Name(i)
	case MEmail:
		return m.Email(i)
	case MMobile:
		return m.Mobile(i)
	case MLastFourDigits:
		return m.LastFourDigits(i)
	case MCreditCard:
		return m.CreditCard(i)
	case MPassportSeries:
		return m.PassportSeries(i)
	case MPassportNumber:
		return m.PassportNumber(i)
	case MCode:
		return m.Code(i)
	}
}

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
		return m.overlay(i, "*", 1, 2)
	}

	if l > 3 {
		return m.overlay(i, "**", 1, 3)
	}

	return "*"
}

func (m *Masker) CreditCard(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}
	return m.overlay(i, "******", 6, 12)
}

func (m *Masker) Email(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}

	chunks := strings.Split(i, "@")

	return m.overlay(chunks[0], "****", 3, 7) + "@" + chunks[1]
}

func (m *Masker) Mobile(i string) string {
	if len(i) == 0 {
		return ""
	}
	return m.overlay(i, "***", 4, 7)
}

func (m *Masker) Password(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}
	return "************"
}

func (m *Masker) PassportSeries(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}
	return m.overlay(i, "**", 1, 3)
}

func (m *Masker) PassportNumber(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}
	return m.overlay(i, "****", 1, 5)
}

func (m *Masker) Code(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}
	if l == 1 {
		return "*"
	}
	if l == 2 || l == 3 {
		return m.overlay(i, strings.Repeat("*", l-1), 1, l)
	}
	return m.overlay(i, strings.Repeat("*", l-2), 1, l-1)
}

func (m *Masker) LastFourDigits(i string) string {
	l := len([]rune(i))
	if l == 0 {
		return ""
	}

	if l < 5 {
		return "****"
	}

	return m.overlay(i, strings.Repeat("*", l-4), 0, l-4)
}

// New create Masker
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
