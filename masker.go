package masker

import "strings"

const (
	MPassword   = "password"
	MName       = "name"
	MEmail      = "email"
	MMobile     = "mobile"
	MCreditCard = "creditcard"
)

type Masker struct{}

func (m *Masker) overlay(str string, overlay string, start int, end int) (overlayed string) {
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
		tmp := start
		start = end
		end = tmp
	}

	overlayed = ""
	overlayed += string(r[:start])
	overlayed += overlay
	overlayed += string(r[end:])
	return overlayed
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
	case MCreditCard:
		return m.CreditCard(i)
	}
}

func (m *Masker) Name(i string) string {
	l := len([]rune(i))

	if l == 0 {
		return ""
	}

	// if has space
	if strs := strings.Split(i, " "); len(strs) > 1 {
		tmp := make([]string, len(strs))
		for idx, str := range strs {
			tmp[idx] = m.Name(str)
		}
		return strings.Join(tmp, " ")
	}

	if l == 2 || l == 3 {
		return m.overlay(i, "**", 1, 2)
	}

	if l > 3 {
		return m.overlay(i, "**", 1, 3)
	}

	return "**"
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

	tmp := strings.Split(i, "@")
	addr := tmp[0]
	domain := tmp[1]

	addr = m.overlay(addr, "****", 3, 7)

	return addr + "@" + domain
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
