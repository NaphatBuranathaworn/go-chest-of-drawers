package bus

import (
	"testing"
)

var f File

func Init() {

}

func TestGetFileSplitterUnderscore(t *testing.T) {
	f.name = "Screenshot_1608534453.png"
	expected := "_"

	actual, err := f.GetFileSplitter()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actual != expected {
		t.Fatalf("TestGetFileSplitterUnderscore : expected %s actual %s", expected, actual)
	}
}

func TestGetFileSplitterSpace(t *testing.T) {
	f.name = "Screen Shot 2564-04-01 at 11.08.26"
	expected := " "

	actual, err := f.GetFileSplitter()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actual != expected {
		t.Fatalf("TestGetFileSplitterSpace : expected %s actual %s", expected, actual)
	}
}

func TestSplitFileNameSpace(t *testing.T) {
	f.name = "Screen Shot 2564-04-01 at 11.08.26"
	sign := " "
	expected := 5

	arr, err := f.SplitFileName(sign)
    actual := len(arr)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actual != expected {
		t.Fatalf("TestSplitFileNameSpace : expected %d actual %d", expected, actual)
	}
}

func TestSplitFileNameUnderscore(t *testing.T) {
	f.name = "Screenshot_1608534453.png"
	sign := "_"
	expected := 2

	arr, err := f.SplitFileName(sign)
    actual := len(arr)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actual != expected {
		t.Fatalf("TestSplitFileNameSpace : expected %d actual %d", expected, actual)
	}
}

func TestFindIndexDateInFileName(t *testing.T) {
	f.name = "Screen Shot 2564-04-01 at 11.08.26"
	sign := " "
	arr, err := f.SplitFileName(sign)

	if err != nil {
		t.Fatalf(err.Error())
	}

	i, err := f.FindIndexDateInFileName(arr)
	actual := arr[i]
	expected := arr[2]

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actual != expected {
		t.Fatalf("TestFindIndexDateInFileName : expected %s actual %s", expected, actual)
	}
}

func TestFindIndexDateInFileNameNotFoundDate(t *testing.T) {
	f.name = "Screenshot_1608534453.png"
	sign := "_"
	arr, err := f.SplitFileName(sign)

	if err != nil {
		t.Fatalf(err.Error())
	}

	i, err := f.FindIndexDateInFileName(arr)
	actual := arr[i]
	expected := arr[0]

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actual != expected {
		t.Fatalf("TestFindIndexDateInFileNameNotFoundDate : expected %s actual %s", expected, actual)
	}
}

