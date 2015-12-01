package key

import "github.com/veandco/go-sdl2/sdl"

type Key int

func (k Key) Name() string {
	return sdl.GetKeyName(sdl.Keycode(k))
}

func FromName(name string) Key {
	code := sdl.GetKeyFromName(name)
	return Key(code)
}

const (
	Unknown Key = sdl.K_UNKNOWN

	Return       Key = sdl.K_RETURN
	Escape       Key = sdl.K_ESCAPE
	Backspace    Key = sdl.K_BACKSPACE
	Tab          Key = sdl.K_TAB
	Space        Key = sdl.K_SPACE
	Exclaim      Key = sdl.K_EXCLAIM
	Hash         Key = sdl.K_HASH
	Percent      Key = sdl.K_PERCENT
	Dollar       Key = sdl.K_DOLLAR
	Ampersand    Key = sdl.K_AMPERSAND
	Quote        Key = sdl.K_QUOTE
	DoubleQuote  Key = sdl.K_QUOTEDBL
	LeftParen    Key = sdl.K_LEFTPAREN
	RightParen   Key = sdl.K_RIGHTPAREN
	Asterisk     Key = sdl.K_ASTERISK
	Plus         Key = sdl.K_PLUS
	Comma        Key = sdl.K_COMMA
	Minus        Key = sdl.K_MINUS
	Period       Key = sdl.K_PERIOD
	Slash        Key = sdl.K_SLASH
	Num0         Key = sdl.K_0
	Num1         Key = sdl.K_1
	Num2         Key = sdl.K_2
	Num3         Key = sdl.K_3
	Num4         Key = sdl.K_4
	Num5         Key = sdl.K_5
	Num6         Key = sdl.K_6
	Num7         Key = sdl.K_7
	Num8         Key = sdl.K_8
	Num9         Key = sdl.K_9
	Colon        Key = sdl.K_COLON
	Semicolon    Key = sdl.K_SEMICOLON
	Less         Key = sdl.K_LESS
	Equals       Key = sdl.K_EQUALS
	Greater      Key = sdl.K_GREATER
	Question     Key = sdl.K_QUESTION
	At           Key = sdl.K_AT
	LeftBracket  Key = sdl.K_LEFTBRACKET
	Backslash    Key = sdl.K_BACKSLASH
	RightBracket Key = sdl.K_RIGHTBRACKET
	Caret        Key = sdl.K_CARET
	Underscore   Key = sdl.K_UNDERSCORE
	Backquote    Key = sdl.K_BACKQUOTE

	A Key = sdl.K_a
	B Key = sdl.K_b
	C Key = sdl.K_c
	D Key = sdl.K_d
	E Key = sdl.K_e
	F Key = sdl.K_f
	G Key = sdl.K_g
	H Key = sdl.K_h
	I Key = sdl.K_i
	J Key = sdl.K_j
	K Key = sdl.K_k
	L Key = sdl.K_l
	M Key = sdl.K_m
	N Key = sdl.K_n
	O Key = sdl.K_o
	P Key = sdl.K_p
	Q Key = sdl.K_q
	R Key = sdl.K_r
	S Key = sdl.K_s
	T Key = sdl.K_t
	U Key = sdl.K_u
	V Key = sdl.K_v
	W Key = sdl.K_w
	X Key = sdl.K_x
	Y Key = sdl.K_y
	Z Key = sdl.K_z

	Capslock Key = sdl.K_CAPSLOCK

	F1  Key = sdl.K_F1
	F2  Key = sdl.K_F2
	F3  Key = sdl.K_F3
	F4  Key = sdl.K_F4
	F5  Key = sdl.K_F5
	F6  Key = sdl.K_F6
	F7  Key = sdl.K_F7
	F8  Key = sdl.K_F8
	F9  Key = sdl.K_F9
	F10 Key = sdl.K_F10
	F11 Key = sdl.K_F11
	F12 Key = sdl.K_F12

	Printscreen Key = sdl.K_PRINTSCREEN
	ScrollLock  Key = sdl.K_SCROLLLOCK
	Pause       Key = sdl.K_PAUSE
	Insert      Key = sdl.K_INSERT
	Home        Key = sdl.K_HOME
	PageUp      Key = sdl.K_PAGEUP
	Delete      Key = sdl.K_DELETE
	End         Key = sdl.K_END
	PageDown    Key = sdl.K_PAGEDOWN
	Right       Key = sdl.K_RIGHT
	Left        Key = sdl.K_LEFT
	Down        Key = sdl.K_DOWN
	Up          Key = sdl.K_UP

	NumlockClear   Key = sdl.K_NUMLOCKCLEAR
	KeypadDivide   Key = sdl.K_KP_DIVIDE
	KeypadMultiply Key = sdl.K_KP_MULTIPLY
	KeypadMinus    Key = sdl.K_KP_MINUS
	KeypadPlus     Key = sdl.K_KP_PLUS
	KeypadEnter    Key = sdl.K_KP_ENTER
	Keypad1        Key = sdl.K_KP_1
	Keypad2        Key = sdl.K_KP_2
	Keypad3        Key = sdl.K_KP_3
	Keypad4        Key = sdl.K_KP_4
	Keypad5        Key = sdl.K_KP_5
	Keypad6        Key = sdl.K_KP_6
	Keypad7        Key = sdl.K_KP_7
	Keypad8        Key = sdl.K_KP_8
	Keypad9        Key = sdl.K_KP_9
	Keypad0        Key = sdl.K_KP_0
	KeypadPeriod   Key = sdl.K_KP_PERIOD
	KeypadEquals   Key = sdl.K_KP_EQUALS
	KeypadComma    Key = sdl.K_KP_COMMA

	LeftCtrl   Key = sdl.K_LCTRL
	LeftShift  Key = sdl.K_LSHIFT
	LeftAlt    Key = sdl.K_LALT
	LeftGUI    Key = sdl.K_LGUI
	RightCtrl  Key = sdl.K_RCTRL
	RightShift Key = sdl.K_RSHIFT
	RightAlt   Key = sdl.K_RALT
	RightGUI   Key = sdl.K_RGUI
)
