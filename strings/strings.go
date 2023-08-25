package main

import "fmt"
import "unicode/utf8"

func stringAsBytes() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}

func stringLiterals() {
	const placeOfInterest = `⌘`
	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)

	fmt.Printf("\nquoted string: ")
	fmt.Printf("%+q", placeOfInterest)

	fmt.Printf("\nhex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}

	fmt.Printf("\n")
}

func rangeLoops() {
	const nihongo = "\xbd日本語"
	fmt.Printf("plain string with len: ")
	fmt.Printf("%s, len: %d\n", nihongo, len(nihongo))
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d and length: %d\n", runeValue, index, len(string(runeValue)))

	}
}

func utf8Library() {
	const nihongo = "日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func byExample() {
	var examineRune = func(r rune) {
		if r == 't' {
			fmt.Println("Found tee")
		} else if r == 'ส' {
			fmt.Println("Found so sua")
		}
	}

	const s = "สวัสดี"

	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d with length: %d\n", runeValue, idx, len(string(runeValue)))
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d with length: %d\n", runeValue, i, width)
		w = width
		examineRune(runeValue)
	}
}

func main() {
	stringAsBytes()
	stringLiterals()
	rangeLoops()
	utf8Library()
	byExample()
}
