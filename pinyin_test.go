package pinyin

import "testing"

func TestConvert(t *testing.T) {
	str := New("张中诚").Split("").Mode(InitialsInCapitals).Convert()
	t.Log(str)
	str = New("张中诚").Split(" ").Mode(WithoutTone).Convert()
	t.Log(str)
	str = New("张中诚").Split("-").Mode(Tone).Convert()
	t.Log(str)
}
