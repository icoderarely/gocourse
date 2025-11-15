package intermediate

import "fmt"

// formatting verb: similar to placeholders in other language
// use printf() -> to print these general formatting verbs

func main() {
	// --- general verbs
	// %v   Prints the value in default format
	// %#v  Prints the value in Go-syntax format
	// %T   Prints the type
	// %%   Prints the % sign

	fmt.Println("------ General Verbs")
	i := 423_415.5
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	string := "Hello World!"
	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string) // "Hello World!" - how they appear in go source code
	fmt.Printf("%T\n", string)

	// --- integer verbs
	// %b   Base 2
	// %d   Base 10
	// %+d  Base 10 + show sign
	// %o   Base 8
	// %O   Base 8 with leading 0o
	// %x   Base 16, lowercase
	// %X   Base 16, uppercase
	// %#x  Base 16, with leading 0x
	// %4d  pad with spaces (widht 4, right justified)
	// %-4d pad with spaces (widht 4, left justified)
	// %04d pad with spaces (widht 4, with 0s)

	fmt.Println("------ Integer Verbs")
	int := 255
	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	// --- String verbs
	// %s   Prints the value as plain string
	// %q   double quoted string
	// %8s  plain string(width 8, right justified)
	// %-8s plain string(width 8, left justified)
	// %x   hex dump of byte values
	// % x  hex dump with spaces
	fmt.Println("------ String Verbs")
	txt := "icode"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// Boolean verbs
	// %t   Value of the boolean operator (same as %v)
	fmt.Println("------ Boolean Verbs")
	t := true
	f := false
	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", f)
	fmt.Printf("%T\n", f)

	// Floating verbs
	// %e   Scientific notation with 'e' as exponent
	// %f   Decimal point, no exponenet
	// %.2f Default width, precision 2
	// %6.2f width 6, precision 2
	// %g   Exponent as needed, only necessary digits
	fmt.Println("------ Float Verbs")

	flt := 089000.00900
	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)
}
