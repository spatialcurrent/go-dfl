// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package shell

type History struct {
	Lines  []string
	Cursor int
}

func (h *History) Get(i int) string {
	return h.Lines[i]
}

func (h *History) Forward() {
	h.Cursor--
}

func (h *History) Back() {
	h.Cursor++
}

func (h *History) Last() string {
	if len(h.Lines) == 0 {
		return ""
	}
	return h.Lines[len(h.Lines)-1]
}

func (h *History) Line() string {
	return h.Lines[len(h.Lines)-h.Cursor]
}

func (h *History) Len() int {
	return len(h.Lines)
}

func (h *History) Empty() bool {
	return len(h.Lines) == 0
}

func (h *History) Push(line string) {
	h.Lines = append(h.Lines, line)
}
