package linkedlist

import "testing"

func asStringPtr(str string) *string {
	s := new(string)
	*s = str
	return s
}

func TestEmpty(t *testing.T) {
	ll := NewLinkedList[string]()

	got := ll.Size()
	want := 0
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestEmptyIterator(t *testing.T) {
	ll := NewLinkedList[*string]()
	it := ll.NewIterator()

	got := it.Next()
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 := it.Get()
	if got2 != nil {
		t.Errorf("got %v, wanted %v", got2, nil)
	}

	got = it.Prev()
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	if got2 != nil {
		t.Errorf("got %v, wanted %v", got2, nil)
	}
}

func TestOne(t *testing.T) {
	ll := NewLinkedList[string]()
	ll.Append(asStringPtr("test string"))

	got := ll.Size()
	want := 1
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestOneIterator(t *testing.T) {
	data := asStringPtr("test string")
	ll := NewLinkedList[string]()
	ll.Append(data)
	it := ll.NewIterator()

	got := it.Next()
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 := it.Get()
	want2 := data
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Next()
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = nil
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Prev()
	want = true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = data
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Prev()
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = nil
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}
}

func TestTwo(t *testing.T) {
	ll := NewLinkedList[string]()
	ll.Append(asStringPtr("test string"))
	ll.Append(asStringPtr("second string"))

	got := ll.Size()
	want := 2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTwoIterator(t *testing.T) {
	data := asStringPtr("test string")
	data2 := asStringPtr("second string")
	ll := NewLinkedList[string]()
	ll.Append(data)
	ll.Append(data2)
	it := ll.NewIterator()

	got := it.Next()
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 := it.Get()
	want2 := data
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Next()
	want = true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = data2
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Next()
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = nil
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Prev()
	want = true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = data2
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Prev()
	want = true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = data
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got = it.Prev()
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 = it.Get()
	want2 = nil
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}
}
