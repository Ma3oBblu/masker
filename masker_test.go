package masker

import (
	"reflect"
	"testing"
)

func TestMasker_overlay(t *testing.T) {
	type args struct {
		str     string
		overlay string
		start   int
		end     int
	}
	tests := []struct {
		name string
		m    *Masker
		args args
		want string
	}{
		{
			name: "Empty Input",
			m:    New(),
			args: args{
				str:     "",
				overlay: "*",
				start:   0,
				end:     0,
			},
			want: "",
		},
		{
			name: "Correct",
			m:    New(),
			args: args{
				str:     "abcdefg",
				overlay: "***",
				start:   1,
				end:     5,
			},
			want: "a***fg",
		},
		{
			name: "Start Less Than 0",
			m:    New(),
			args: args{
				str:     "abcdefg",
				overlay: "***",
				start:   -1,
				end:     5,
			},
			want: "***fg",
		},
		{
			name: "Start Greater Than Length",
			m:    New(),
			args: args{
				str:     "abcdefg",
				overlay: "***",
				start:   30,
				end:     31,
			},
			want: "abcdefg***",
		},
		{
			name: "End Less Than 0",
			m:    New(),
			args: args{
				str:     "abcdefg",
				overlay: "***",
				start:   1,
				end:     -5,
			},
			want: "***bcdefg",
		},
		{
			name: "Start Less Than End",
			m:    New(),
			args: args{
				str:     "abcdefg",
				overlay: "***",
				start:   5,
				end:     1,
			},
			want: "a***fg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masker{}
			if got := m.overlay(tt.args.str, tt.args.overlay, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Masker.overlay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Masker
	}{
		{
			name: "New Instance",
			want: &Masker{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Russian Length 1",
			args: args{
				i: "А",
			},
			want: "*",
		},
		{
			name: "Russian Length 2",
			args: args{
				i: "Яр",
			},
			want: "Я*",
		},
		{
			name: "Russian Length 3",
			args: args{
				i: "Юра",
			},
			want: "Ю*а",
		},
		{
			name: "Russian Length 4",
			args: args{
				i: "Вика",
			},
			want: "В**а",
		},
		{
			name: "Russian Length 5",
			args: args{
				i: "Антон",
			},
			want: "А**он",
		},
		{
			name: "Russian Length 6",
			args: args{
				i: "Виктор",
			},
			want: "В**тор",
		},
		{
			name: "Russian Full Name",
			args: args{
				i: "Виктор Иванов",
			},
			want: "В**тор И**нов",
		},
		{
			name: "Russian Full Name With Spaces",
			args: args{
				i: "  Виктор   Иванов   ",
			},
			want: "В**тор И**нов",
		},
		{
			name: "English Length 4",
			args: args{
				i: "King",
			},
			want: "K**g",
		},
		{
			name: "English Full Name",
			args: args{
				i: "King Kong",
			},
			want: "K**g K**g",
		},
		{
			name: "English Full Name",
			args: args{
				i: "Charles Dickens",
			},
			want: "C**rles D**kens",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Name(tt.args.i); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreditCard(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "VISA JCB MasterCard",
			args: args{
				i: "1234567890123456",
			},
			want: "123456******3456",
		},
		{
			name: "American Express",
			args: args{
				i: "123456789012345",
			},
			want: "123456******345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreditCard(tt.args.i); got != tt.want {
				t.Errorf("CreditCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Correct",
			args: args{
				i: "test.mail@gmail.com",
			},
			want: "tes****il@gmail.com",
		},
		{
			name: "Address Less Than 3",
			args: args{
				i: "tt@gmail.com",
			},
			want: "tt****@gmail.com",
		},
		{
			name: "Address Less Equal 1",
			args: args{
				i: "t@gmail.com",
			},
			want: "t****@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Email(tt.args.i); got != tt.want {
				t.Errorf("Email() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMobile(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Correct",
			args: args{
				i: "79191232323",
			},
			want: "7919***2323",
		},
		{
			name: "Correct",
			args: args{
				i: "78432232323",
			},
			want: "7843***2323",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mobile(tt.args.i); got != tt.want {
				t.Errorf("Mobile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Correct",
			args: args{
				i: "1234567",
			},
			want: "************",
		},
		{
			name: "Correct",
			args: args{
				i: "abcd!@#$%321",
			},
			want: "************",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Password(tt.args.i); got != tt.want {
				t.Errorf("Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassportSeries(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Correct",
			args: args{
				i: "1234",
			},
			want: "1**4",
		},
		{
			name: "Correct",
			args: args{
				i: "9267",
			},
			want: "9**7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PassportSeries(tt.args.i); got != tt.want {
				t.Errorf("PassportSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassportNumber(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Correct",
			args: args{
				i: "123456",
			},
			want: "1****6",
		},
		{
			name: "Correct",
			args: args{
				i: "926734",
			},
			want: "9****4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PassportNumber(tt.args.i); got != tt.want {
				t.Errorf("PassportNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCode(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Length Equal 1",
			args: args{
				i: "1",
			},
			want: "*",
		},
		{
			name: "Length Equal 2",
			args: args{
				i: "12",
			},
			want: "1*",
		},
		{
			name: "Length Equal 3",
			args: args{
				i: "123",
			},
			want: "1**",
		},
		{
			name: "Length Equal 4",
			args: args{
				i: "1234",
			},
			want: "1**4",
		},
		{
			name: "Length Equal 5",
			args: args{
				i: "12345",
			},
			want: "1***5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code(tt.args.i); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastFourDigits(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty Input",
			args: args{
				i: "",
			},
			want: "",
		},
		{
			name: "Length Equal 1",
			args: args{
				i: "1",
			},
			want: "****",
		},
		{
			name: "Length Equal 2",
			args: args{
				i: "12",
			},
			want: "****",
		},
		{
			name: "Length Equal 3",
			args: args{
				i: "123",
			},
			want: "****",
		},
		{
			name: "Length Equal 4",
			args: args{
				i: "1234",
			},
			want: "****",
		},
		{
			name: "Length Equal 9",
			args: args{
				i: "79191232323",
			},
			want: "*******2323",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastFourDigits(tt.args.i); got != tt.want {
				t.Errorf("LastFourDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
