package aquestalk

import "fmt"

type errno int

func (nr errno) Error() string {
	switch nr {
	case 100:
		return "misc error (error code: 100)"
	case 101:
		return "out of memory"
	case 102:
		return "undefined reading symbols (error code: 102)"
	case 103:
		return "negative prosody data"
	case 104:
		return "undefined delimiters"
	case 105:
		return "undefined reading symbols (error code: 105)"
	case 106:
		return "illegal tags"
	case 107:
		return "too long tags"
	case 108:
		return "invalid tag values"
	case 109:
		return "failed to play wave (error code: 109)"
	case 110:
		return "failed to play wave (error code: 110)"
	case 111:
		return "no sound data to play"
	case 200:
		return "too long phonetic string (error code: 200)"
	case 201:
		return "too many reading symbols in a phrase"
	case 202:
		return "too long phonetic string (error code: 202)"
	case 203:
		return "heap memory exhaust"
	case 204:
		return "too long phonetic string (error code: 204)"
	default:
		return fmt.Sprintf("undefined error code: %d", nr)
	}
}

