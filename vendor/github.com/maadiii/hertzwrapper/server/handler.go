package server

import (
	"fmt"
	"strconv"
	"strings"
)

type Handler[IN any, OUT any] struct {
	action      func(*Context, IN) (OUT, error)
	path        string
	method      string
	status      string
	contentType string
}

func (h *Handler[IN, OUT]) fix() {
	name := funcPathAndName(h.action)
	desc := h.getFixedDescriberArray()

	for _, d := range desc {
		if strings.HasPrefix(d, "/") {
			d = strings.TrimRight(d, "/")
			h.path = d

			continue
		}

		if strings.HasPrefix(d, "[") && strings.HasSuffix(d, "]") {
			verb, ok := methods[strings.ToUpper(d)]
			if !ok {
				panic(fmt.Sprintf("%s has wrong @action describer, Invalid VERB", name))
			}

			h.method = verb

			continue
		}

		if _, err := strconv.Atoi(d); err == nil {
			h.status = d

			continue
		}

		h.contentType = d
	}
}

func (h *Handler[IN, OUT]) getFixedDescriberArray() []string {
	comment := funcDescription(h.action)
	comments := strings.Split(comment, "\n")

	for _, actionDescription := range comments {
		if !strings.HasPrefix(actionDescription, "@action") {
			continue
		}

		desc := strings.Split(actionDescription, " ")
		index := 0

		for i := 0; i < len(desc); i++ {
			if len(desc[i]) != 0 {
				desc[index] = desc[i]
				index++
			}
		}

		desc = desc[1:index]

		return desc
	}

	return []string{}
}
