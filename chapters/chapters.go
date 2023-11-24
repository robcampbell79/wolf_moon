package chapters

import(
	_ "fmt"
)

func chapter1(st int) (string, int) {
	var chpt1 []string
	var st0 string = ""
	var st1 string = ""
	var mode int = 0
	// mode: 5 = story, 11 = interactive

	st0 = "---------------------------Full Moon Madness---------------------------\n\n" +
			"\tThe town of Talbot is a town in the Pacific Northwest that\n" +
			"\tis secluded by dense forest. This seclusion has created a\n" +
			"\tvery insular town where everyone seems to know evryone else,\n" +
			"\tthat is until a stranger showed up."

	st1 = "---------------------------Full Moon Madness---------------------------\n\n" +
			"\tThe killings started small, just pets here and there and then\n" +
			"\tsome homeless people would disappear. But eventually, people \n" +
			"\tfrom aound town would go missing. It's obvious we have a fucking\n" +
			"\twerewolf, and you need to find the furry fuck face"

	chpt1 = append(chpt1, st0)

	chpt1 = append(chpt1, st1)

	if st < 2 {
		mode = 5
	} else {
		mode = 11
	}

	return chpt1[st], mode
}

func GetChapter(chpt int, step int) (string, int) {
	var story string = ""
	var mode int = 0

	switch chpt {
	case 1:
		story, mode = chapter1(step)
	}

	return story, mode
}