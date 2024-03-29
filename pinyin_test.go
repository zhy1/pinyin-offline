package pinyin

import (
	"testing"
)

func TestConvert(t *testing.T) {
	str, err := New("我是中国人").Split("").Mode(InitialsInCapitals).Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}

	str, err = New("我是中国人").Split(" ").Mode(WithoutTone).Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}

	str, err = New("我是中国人").Split("-").Mode(Tone).Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}

	str, err = New("我是中国人").Convert()
	if err != nil {
		t.Error(err)
	}else{
		t.Log(str)
	}
}

