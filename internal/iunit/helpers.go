package iunit

import (
	"rooms/internal/ihelp"
)

func Radio(label string, bp []BodyPart) BodyPart {
	var str_opts []string

	for _, v := range bp {
		str_opts = append(str_opts, v.String())
	}

	res := ihelp.Radio(label, str_opts)

	for k, v := range bodyParts {
		if v == res {
			return k
		}
	}

	return Unknown
}
