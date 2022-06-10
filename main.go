package main

import (
	"fmt"
	"log"
	"os/exec"

	tb "github.com/nsf/termbox-go"
)

func reset() {
	tb.Sync() // cosmestic purpose
}
func speak(speakString string, printKey string) {
	fmt.Printf("%s  ", printKey)
	cmd := exec.Command(`espeak`, "-v", "f4", "-s", "140", "-p", "90", "-a", "100", speakString)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func newCharacter(input string) (output string) {

	switch input {
	case "`":
		return "backtick"
	case "^":
		return "carat"
	case "(":
		return "left parenthesis"
	case ")":
		return "right parenthesis"
	case "-":
		return "hyphen or minus"
	case "_":
		return "underscore "
	case "{":
		return "left braces"
	case "}":
		return "right braces"
	case "[":
		return "open bracket"
	case "]":
		return "close bracket"
	case "|":
		return "pipe"
	case ";":
		return "semicolon"
	case `"`:
		return "quotation mark"
	case "'":
		return "apostrophe"
	case "<":
		return "left angle brackets"
	case ">":
		return "right angle brackets"
	case ",":
		return "comma"
	case ".":
		return "period or decimal"
	case "/":
		return "forward slash or divide"
	case "*":
		return "asterisk or multiply"
	case "?":
		return "question mark"

	default:
		return input
	}
}

func main() {

	fmt.Println("lello")
	err := tb.Init()
	if err != nil {
		panic(err)
	}

	defer tb.Close()

	fmt.Println("Enter any key on QWERTY keyboard it will make helpful for those can  not see  and  want to type  but can not do ," +
		" and also it will be useful whom that they has no reminded the  keyboard keys   in their memory..")
	fmt.Printf("To exit press Ctrl + any key.\n\n")

	for {

		a := func() string {
			switch eventKey := tb.PollEvent(); eventKey.Type {
			case tb.EventKey:
				switch eventKey.Key {
				case tb.KeyEsc:

					reset()
					return "Escape"
				case tb.KeyF1:
					reset()

					return "f 1"
				case tb.KeyF2:
					reset()

					return "f 2"
				case tb.KeyF3:
					reset()

					return "f 3"
				case tb.KeyF4:
					reset()

					return "f 4"
				case tb.KeyF5:
					reset()

					return "f 5"
				case tb.KeyF6:
					reset()

					return "f 6"
				case tb.KeyF7:
					reset()

					return "f 7"
				case tb.KeyF8:
					reset()

					return "f 8"
				case tb.KeyF9:
					reset()

					return "f 9"
				case tb.KeyF10:
					reset()

					return "f 10"
				case tb.KeyF11:
					reset()

					return "f 11"
				case tb.KeyF12:
					reset()

					return "f 12"
				case tb.KeyInsert:
					reset()

					return "insert"
				case tb.KeyDelete:
					reset()

					return "delete"
				case tb.KeyHome:
					reset()

					return "home"
				case tb.KeyEnd:
					reset()

					return "end"
				case tb.KeyPgup:
					reset()

					return "page up"
				case tb.KeyPgdn:
					reset()

					return "page down"
				case tb.KeyArrowUp:
					reset()

					return "arrow up "
				case tb.KeyArrowDown:
					reset()

					return "arrow down"
				case tb.KeyArrowLeft:
					reset()

					return "arrow left"
				case tb.KeyArrowRight:
					reset()

					return "arrow right"
				case tb.KeySpace:
					reset()

					return "space"
				case tb.KeyBackspace:
					reset()

					return "back space"
				case tb.KeyEnter:
					reset()

					return "enter"
				case tb.KeyTab:
					reset()

					return "tab"
				case tb.KeyCtrlZ:
					reset()

					return "control z"

				default:
					//Read single key except above
					reset()
					return string(eventKey.Ch)
				}
			case tb.EventError:
				panic(eventKey.Err)

			}
			return ""
		}()

		newChar := newCharacter(a)
		speak(newChar, a)
	}

}
