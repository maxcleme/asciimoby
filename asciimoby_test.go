package asciimoby

import (
	"fmt"
)

func ExampleGenerate() {
	s := Generate()
	fmt.Println(s)
	// Output:
	//
	//	              ,;;:
	//	              ;ff1
	//	              iff1
	//	              ,;;:
	//	      ,11i:11i:11;
	//	      :fft;ff1iff1      :1
	//	      :tt1:tt1;tt1      tf1
	//	   ,,,.:::.::,,::,.,,. ,fff:
	//	  ,fft;ff1;ff1;ffiiff1 :fff1;;,
	//	  ,fft;fft;ff1iff1iff1 ,fffffff1
	//	  .iii,ii;,ii;:ii;:ii:  tffffff;
	//	:iiiiiiiiiiiiiiiiiiiiii1ffffff;
	//	fffffffffffffffffffffffffft;:.
	//	tfffffffffffffffffffffffffi
	//	tfffffffffffffffffffffffff.
	//	1ffffffffffffffffffffffffi
	//	;ffffffffffffffffffffffft.
	//	.fffffffffffffffffffffff:
	//	 ifffffffffffffffffffff;
	//	 ,ffffffffffffffffffff;
	//	  ;ffffffffffffffffft:
	//	   :tffffffffffffft;.
	//	    .;1fffffffft1:.
	//	       .,::::,,
}
