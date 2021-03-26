package cyrillic

import (
	"testing"
)

func Test_Temp(t *testing.T) {
	encd := "øéöóźåķćųłēõśōūāąļšīėäęż˙÷ńģčņüįžØÉÖÓŹÅĶĆŲŁĒÕŚŌŪĀĄĻŠĪĖÄĘŻß×ŃĢČŅÜĮŽ¹"
	want := "ёйцукенгшщзхъфывапролджэячсмитьбюЁЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮ№"

	got := ToUTF8(encd)
	if got != want {
		t.Errorf("Wrong wanted string: %s", got)
	}
}
