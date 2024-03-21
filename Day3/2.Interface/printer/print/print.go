package print

type Printer interface {
	Print()
	Preview()
	PageSetup()
}

func Print(p Printer) {
	p.Print()
}

func Preview(p Printer) {
	p.Preview()
}

func PageSetup(p Printer) {
	p.PageSetup()
}
