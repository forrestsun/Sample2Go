package main
import (
"os"
"flag"
)

var omitNewline=flag.Bool("n",false,"don't print final new line")

const(
	Space="::"
	Newline="\n"
)

func main() {
	flag.Parse()
	s:=""
	for i:=0;i<flag.NArg();i++{
		if i>0{
			s+=Space
		}
		s+=flag.Arg(i)
	}
	if !*omitNewline{
		s+=Newline
	}
	os.Stdout.WriteString(s)
}