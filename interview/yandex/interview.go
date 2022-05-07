package yandex

import "strconv"

// Необходимо исправить код:
// и объяснить, что выведет исправленный код.
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			print(i)
		}(i)
	}
	// wait
}

// Написать функцию получающую на вход произвольную строку, состоящую их печатных ASCII символов,
// и возвращающую эту же строку в которой все буквы в нижнем регистре заменены на соответствующие буквы в верхнем регистре.

func upperCase(in string) string {
	diff := 'a' - 'A'
	result := make([]rune, len(in))
	for _, sym := range in {
		if sym >= 'a' && sym <= 'z' {
			sym -= diff
		}
		result = append(result, sym)
	}

	return string(result)
}

/*
Дана скобочная грамматика: в круглых скобках записано выражение,
следом за ними идут квадратные скобки с неотрицательным целым числом.
Разворачивается в виде конкатенации строки в круглых скобках на
саму себя N раз, где N - число в квадратных скобках.

Чуть более формально, грамматика:

term ::= a-z*
term ::= (term)[\d+]
term ::= term1term2 - разворачивается в конкатенацию двух выражений.

Если видно что кандидат не справляется с задачей, можно убрать это
правило вывода - без него задача решается значительно легче в виде
простой линейной рекурсии.

На вход подается строка со скобочным выражением. Необходимо развернуть
его в строку и вернуть. Предлагается считать, что скобочная
последовательность на входе всегда корректна (не хочется здесь решать
задачу валидации).

Примеры:
С двумя правилами:

"" -> ""
"ab" -> "ab"
"(ab)[3]" -> "ababab"
"((ab)[2])[2]" -> "abababab"
"(()[1])[2]" -> ""
С тремя правилами:
"(a)[2](b)[2]" -> "aabb"
    ^
s{}{a}
arg{a}
"((a)[2]b)[3]" -> "aabaabaab"
       ^
"()[1]bc" -> "bc"
"(a)[0]bc" -> "bc"
"abc(d)[2]" -> "abcdd"
*/
/**
asdf((ab)[12])[3]a
                ^
stack[0] = ab

//stack[1] = abababababab
*/

func marshal(expr string) string {
	stack := [][]rune{{}} // len=1
	multBuf := make([]rune, 0)
	arg := []rune{}
	mult := 0
	state := 0 // 0 - вне, 1 - [, 2 - ], 3 - {, 4 - }
	for _, sym := range expr {
		switch sym {
		case '(':
			stack = append(stack, []rune{}) // stack {}{}
			state = 3
		case ')':
			arg = stack[len(stack)-1]    // arg = {}
			stack = stack[:len(stack)-1] // stack {}
		case '[':
			state = 1
		case ']':
			mult, _ = strconv.Atoi(string(multBuf)) // 1
			multBuf = []rune{}
			if len(arg) != 0 { //RESET multBuf
				val := []rune{}
				for i := 0; i < mult; i++ {
					val = append(val, arg...)
				}
				stack[len(stack)-1] = append(stack[len(stack)-1], val...)
			}
			state = 0
		default:
			switch state {
			case 1:
				multBuf = append(multBuf, sym) // multBuf 1
			case 3: // (
				stack[len(stack)-1] = append(stack[len(stack)-1], sym)
			case 4: // )
				state = 0
			case 0: // default
				stack[len(stack)-1] = append(stack[len(stack)-1], sym) // stack {bc}
			}
		}
	}

	return string(stack[len(stack)-1])
}
