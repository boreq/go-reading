package reading_test

import (
	"testing"
	"time"

	reading "github.com/boreq/go-reading"
	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	tests := []struct {
		Text         string
		ExpectedTime time.Duration
	}{
		{
			Text:         "",
			ExpectedTime: 0 * time.Minute,
		},
		{
			Text:         "This is a short sentence.",
			ExpectedTime: 1500 * time.Millisecond,
		},
		{
			Text: `
To all of which I do solemnly and sincerely promise and swear, without any
hesitation, mental reservation, or secret evasion of mind in me whatsoever;
binding myself under no less a penalty than that of having my body severed in
twain, my bowels taken thence, and with my body burned to ashes, and the ashes
thereof scattered to the four winds of Heaven, that there might remain neither
track, trace nor remembrance among man or womyn of so vile and perjured a
wretch as I should be, should I ever knowingly or willfully violate this, my
solemn Obligation of a Golang user. So help me God and make me steadfast to
keep and perform the same.
			`,
			ExpectedTime: 40 * time.Second,
		},
	}

	for _, test := range tests {
		readingTime := reading.Time(test.Text)
		t.Log(readingTime)
		if test.ExpectedTime == 0 {
			require.Equal(t, test.ExpectedTime, readingTime)
		} else {
			require.InEpsilon(t, test.ExpectedTime, readingTime, 0.1)
		}
	}
}
