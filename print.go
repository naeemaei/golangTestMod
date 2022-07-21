package Printer

func Print(s ...string) {
	for _, v := range s {
		println(v)
	}
}
