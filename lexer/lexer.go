package san

// StateFn represents the current lewer's state
type StateFn func(lx *Lexer) StateFn

// Lexer is a SAN lexer
type Lexer struct {
	input             string
	currentTokenStart int64
	start             int64
	pos               int64
	tokens            chan Token
	line              int64
	col               int64
	state             StateFn
	stack             []StateFn
}

func Lex(input string) []Token {
	ret := []Token{}
	return ret
}

func NewLexer(input string) *Lexer {
	lx := &Lexer{
		input:  input,
		state:  lexTop,
		line:   1,
		tokens: make(chan Token),
		stack:  make([]StateFn, 0, 10),
	}
	return lx
}

// lexTop consumes elements at the top level of SAN data.
func lexTop(lx *Lexer) StateFn {
	/*
		r := lx.next()
		if isWhitespace(r) || isNL(r) {
			return lexSkip(lx, lexTop)
		}
		switch r {
		case commentStart:
			lx.push(lexTop)
			return lexCommentStart
		case tableStart:
			return lexTableStart
		case eof:
			if lx.pos > lx.start {
				return lx.errorf("unexpected EOF")
			}
			lx.emit(itemEOF)
			return nil
		}

		// At this point, the only valid item can be a key, so we back up
		// and let the key lexer do the rest.
		lx.backup()
		lx.push(lexTopEnd)
		return lexKeyStart
	*/
	return nil
}

// skip ignores all slurped input and moves on to the next state.
func (lx *Lexer) skip(nextState StateFn) StateFn {
	return func(lx *Lexer) StateFn {
		lx.ignore()
		return nextState
	}
}

// ignore skips over the pending input before this point.
func (lx *Lexer) ignore() {
	lx.start = lx.pos
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isHexadecimal(r rune) bool {
	return (r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'f') ||
		(r >= 'A' && r <= 'F')
}

func isBareKeyChar(r rune) bool {
	return (r >= 'A' && r <= 'Z') ||
		(r >= 'a' && r <= 'z') ||
		(r >= '0' && r <= '9') ||
		r == '_' ||
		r == '-'
}

// isWhitespace returns true if `r` is a whitespace character according to the spec.
func isWhitespace(r rune) bool {
	return r == ' '
}

// isWhitespace returns true if `r` is a newline character according to the spec.
func isNL(r rune) bool {
	return r == '\n' || r == '\r'
}
