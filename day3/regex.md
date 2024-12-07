https://regexr.com/ -> helped a bit

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
