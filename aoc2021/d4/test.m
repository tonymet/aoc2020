use Test::More tests => 1;
$result = `cat data/input.txt| PART=1 go run main.go|grep solution |tail -n 1`;
$result =~ s/.*solution/solution/g;
is($result , "solution: 39984\n", "solution pattern matches");
