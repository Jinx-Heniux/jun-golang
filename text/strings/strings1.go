package strings

import (
	"fmt"
	"strings"
	"unicode"
)

func Strings1Example1() {
	a := "gopher"
	b := "hello world"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("壹", "一"))

	fmt.Println("###############################")

	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Println("###############################")

	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(len("谷歌中国"))
	fmt.Println(strings.Count("谷歌中国", ""))
	fmt.Println(strings.Count("fivevev", "vev"))

	fmt.Println("###############################")

	fmt.Printf("Fields are: %v\n", strings.Fields("  foo bar  baz   "))
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   ")) // 双引号括起来的字符串
	fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))

	fmt.Println("###############################")

	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))

	fmt.Println()
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))

	fmt.Println()
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println("###############################")

	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))

	fmt.Println("###############################")

	han := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 汉字
	}
	fmt.Println(strings.IndexFunc("Hello, world", han))
	fmt.Println(strings.IndexFunc("Hello, 世界", han))

	fmt.Println("###############################")

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println("###############################")
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写字母转小写
			return r + 32
		case r >= 'a' && r <= 'z': // 小写字母不处理
			return r
		case unicode.Is(unicode.Han, r): // 汉字换行
			return '\n'
		}
		return -1 // 过滤所有非字母、汉字的字符
	}
	fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))

	fmt.Println("###############################")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))

	fmt.Println("###############################")
	fmt.Println(strings.ToLower("HELLO WORLD"))
	fmt.Println(strings.ToLower("Ā Á Ǎ À"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "壹"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD"))
	fmt.Println(strings.ToLower("Önnek İş"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.ToUpper("ā á ǎ à"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "一"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world"))
	fmt.Println(strings.ToUpper("örnek iş"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))

	fmt.Println("###############################")
	fmt.Println(strings.Title("hElLo wOrLd"))
	fmt.Println(strings.ToTitle("hElLo wOrLd"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))
	fmt.Println(strings.Title("āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	fmt.Println("###############################")
	x := "!!!@@@你好,!@#$ Gophers###$$$"
	fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println(strings.TrimPrefix(x, "!"))
	fmt.Println(strings.TrimSuffix(x, "$"))

	f := func(r rune) bool {
		return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	}
	fmt.Println(strings.TrimFunc(x, f))
	fmt.Println(strings.TrimLeftFunc(x, f))
	fmt.Println(strings.TrimRightFunc(x, f))

	fmt.Println("###############################")
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	fmt.Println("###############################")
	bb := strings.Builder{}
	_ = bb.WriteByte('7')
	n, _ := bb.WriteRune('夕')
	fmt.Println(n)
	n, _ = bb.Write([]byte("Hello, World"))
	fmt.Println(n)
	n, _ = bb.WriteString("你好，世界")
	fmt.Println(n)
	fmt.Println(bb.Len())
	fmt.Println(bb.Cap())
	bb.Grow(100)
	fmt.Println(bb.Len())
	fmt.Println(bb.Cap())
	fmt.Println(bb.String())
	bb.Reset()
	fmt.Println(bb.String())

}
