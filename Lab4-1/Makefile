SHELL:=/bin/bash

programs:=echo cat head tail sort ls

.PHONY: test
test: $(patsubst %,test_%,$(programs))

test_echo: echo
	diff <(echo) <(./echo)
	diff <(echo foo) <(./echo foo)
	diff <(echo foo bar) <(./echo foo bar)

test_cat: cat
	diff <(echo foo) <(echo foo | ./cat)
	diff cat.go <(./cat cat.go)
	diff <(cat cat.go cat.go) <(./cat cat.go cat.go)

test_head: head
	diff <(head /dev/null) <(./head /dev/null)
	diff <(cat head.go | head) <(cat head.go | ./head)
	diff <(head head.go) <(./head head.go)
	diff <(head head.go head.go) <(./head head.go head.go head.go)

test_tail: tail
	diff <(tail /dev/null) <(./tail /dev/null)
	diff <(cat tail.go | tail) <(cat tail.go | ./tail)
	diff <(tail tail.go) <(./tail tail.go)
	diff <(tail tail.go tail.go) <(./tail tail.go tail.go tail.go)

test_sort: sort
	diff <(sort /dev/null) <(./sort /dev/null)
	diff <(cat sort.go | sort) <(cat sort.go | ./sort)
	diff <(sort sort.go) <(./sort sort.go)
	diff <(sort sort.go sort.go) <(./sort sort.go sort.go sort.go)

test_ls: ls
	diff <(ls) <(./ls)
	diff <(ls tail.go) <(./ls tail.go)

%: %.go
	go build -o $@ $^

.PHONY: clean
clean::
	$(RM) $(programs)
