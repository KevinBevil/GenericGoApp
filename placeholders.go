package placeholders

type placeholder [5]string
// Each digit to be put into our clock multi-dimensional array at the appropriate
// time
zero := placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

one := placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

two := placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

three := placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

four := placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

five := placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

six := placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

seven := placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

eight := placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

nine := placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

colon := placeholder{
	"   ",
	" █ ",
	"   ",
	" █ ",
	"   ",
}
colon2 := placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

	// Each of the 10 digits held here for convenience
	digits := [10]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}