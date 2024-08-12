package models

import (
	"fmt"
)

type Sentence struct {

	// // timestamp that the sentence was received
	// Timestamp time.Time

	// // source identifier
	// Source string

	// ais v4.10+ TAG block e.g. \s:HE0002,c:1474107266,n:5*20\
	TagBlock string

	// nmea/ais sentence e.g. !AIVDM,1,1,,B,13PRrB0000OvbS@NhA9=oPbr0<0u,0*58
	Content string

	// Sequence id number
	Id uint64
}

func (s *Sentence) IsEmpty() bool {

	return len(s.Content) == 0

}

func (s *Sentence) AsString() string {

	if s.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%s%s\n\r", s.TagBlock, s.Content)

}
