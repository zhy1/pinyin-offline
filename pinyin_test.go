package pinyin

import (
	"testing"
)

func TestConvert(t *testing.T) {
	str, err := New("张中诚").Split("").Mode(InitialsInCapitals).Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}

	str, err = New("张中诚").Split(" ").Mode(WithoutTone).Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}

	str, err = New("张中诚").Split("-").Mode(Tone).Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}

	str, err = New("张中诚").Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}
}

