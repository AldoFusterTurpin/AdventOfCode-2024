https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatch
https://pkg.go.dev/regexp#Regexp.FindStringSubmatch
https://regexr.com/

### Part 1

The regex we need for part 1:
mul -> literal string
\((\d{1,3})
	\( -> scape the opening parenthesis
	( -> open capture bloc
	\d{1,3} -> any digit with 1,2 or 3 digits
	) -> close capture group

, -> literal char
(\d{1,3})\) -> same capture group as above

----------------------------------------------------------------------------------
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
"s" is the above string ^

FindAllStringSubmatch(s) will return:
[
    ["mul(2,4)", "2", "4"],
    ["mul(5,5)", "5", "5"],
    ["mul(11,8)", "11", "8"],
    ["mul(8,5)", "8", "5"]
]
because it extracts all the submatches with the groups -> what we want!
https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatch


FindStringSubmatch(s) will return:
[
    "mul(2,4)",
    "2",
    "4"
]
because it extracts just the first submatch with the first groups -> we don't want this because we want all the submatches.
https://pkg.go.dev/regexp#Regexp.FindStringSubmatch


### Part 2
For the part 2, what we need to do conceptually is to first extract just the substrings that are inside do()...don't(), and forget about the rest. To do that easily, I simply add a literal do() at the begining of the input (as the multiplications are activated by default) and a don't() at the end of the input. If I don't add "don't()" at the end, I would need to think how to indicate in the regex that I want to match the next don't() OR the end of the line, and I thought that doing this trick was easier and simpler (just my feeling, not saying is the prefered way to do so).

So in simple worlds, the Part 2 is the same as part 1 just that we need to consider all the substrings that are between do() ... don't().

On important thing to consider is that we don't want the regex to be greedy, because we want to match the samllest matches to not erase by mistake parts that we need to consider. We can do that by adding "?" just after the capture group.

Example:
do()mult(2,2)don't()mul(3,3)do()mul(4,4)don't()

in the above string we want to extract -> "mult(2,2)" and "mul(4,4)",
we do NOT want to extract "mult(2,2)", "mul(4,4)" and "mul(3,3)", as "mul(3,3)" has a "don't()" just before it.

In other/simpler words: if we do a greedy search, we are basically taking everything between the first do() and the last don'(), but we want all the multiplications in the middle discarding the ones that have a don't before them.

So this is the regex that we will use:
do\(\)(.*?)don't\(\)

do\(\) -> literal "do()"
(.*?) -> anything in the middle, ? just means not greedy
don't\(\) -> literal "don't()"

[Go playground](https://go.dev/play/p/0ZMwHNqWeAA)
```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := `do\(\)(.*?)don't\(\)`
	re := regexp.MustCompile(regex)
	s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+do()don't()mul(32,64)do()mul(32,64]don't()(mul(11,8)undo()?mul(8,5))"
	res := re.FindAllStringSubmatch(s, -1)
	fmt.Println(len(res))
	for _, v := range res {
		fmt.Printf("%T%v\n", v, v)
	}
}
```

The key to solve this problem is to play with the regex and many inputs to see how the pattern matches the input. I used https://regexr.com/ and https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatch for that.

Lastly, I needed to merge the multiple lines of the input before applying the regex to treat it as a single string.