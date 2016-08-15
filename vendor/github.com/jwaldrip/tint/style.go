package tint

import (
	"fmt"
	"strconv"
	"strings"
)

// Style is the struct that holds information on how to style a string
type Style struct {
	Color      int
	Background int
	Bold       bool
	Dim        bool
	Blink      bool
	Underline  bool
	Inverted   bool
	Use256     bool
}

// Render the styled string
func (s *Style) Render(str string) string {
	var strs []string
	if s.Bold {
		strs = append(strs, strconv.Itoa(bold))
	}
	if s.Dim {
		strs = append(strs, strconv.Itoa(dim))
	}
	if s.Underline {
		strs = append(strs, strconv.Itoa(underline))
	}
	if s.Blink {
		strs = append(strs, strconv.Itoa(blink))
	}
	if s.Inverted {
		strs = append(strs, strconv.Itoa(invert))
	}
	if s.Use256 {
		strs = append(strs, "48;5")
		strs = append(strs, s.getColor())
		strs = append(strs, "38;5")
		strs = append(strs, s.getBackground())
	} else {
		if s.Color != 0 && s.Color != DefaultColor {
			strs = append(strs, s.getColor())
		}
		if s.Background != 0 && s.Background != DefaultColor {
			strs = append(strs, s.getBackground())
		}
	}
	if len(strs) == 0 {
		return str
	}
	return fmt.Sprintf(
		"%s[%sm%s%s[%sm",
		escape,
		strings.Join(strs, ";"),
		str,
		escape,
		strconv.Itoa(reset),
	)
}

func (s *Style) getColor() string {
	clr := s.Color
	return strconv.Itoa(clr)
}

func (s *Style) getBackground() string {
	bg := s.Background
	return strconv.Itoa(bg + 10)
}
