package f

import "fmt"

func f1(a int) error {
	if a == 3 {
		return fmt.Errorf("err 3")
	}
	return nil
}

func f2(b int) (int, error) {
	if b == 3 {
		return 0, fmt.Errorf("b err 3")
	}
	return 0, nil
}

type D struct {
	name string
}

func f3(b int) (int, *D, error) {
	if b > 0 {
		return 0, &D{
			name: "foo",
		}, nil
	}
	return 0, nil, fmt.Errorf("empty")
}

type E struct {
	id   int
	name string
	*F
}

type F struct {
	id int
}

func f4(b int) (int, string, *D, *E, error) {
	if b > 0 {

	}
	return 0, "", &D{
			name: "foo",
		}, &E{
			id:   0,
			name: "foo-E",
			F:    &F{id: 100},
		}, nil
}
