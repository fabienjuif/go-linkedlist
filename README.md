# go-linkedlist

![build & test](https://github.com/fabienjuif/go-linkedlist/actions/workflows/simple.yml/badge.svg)

> Simple LinkedList in Go

## Install

```sh
go get github.com/fabienjuif/go-linkedlist
```

## Thread safety

The Linked List is thread safe but the Iterator is not.

If you update the Linked List while Iterate over it you may notice strange behaviour.

## Example

```golang
package main

import (
	"fmt"

	"github.com/fabienjuif/go-linkedlist"
)

type Song struct {
	title  string
	artist string
}

func main() {
	// best 2000 rock songs according to https://us.napster.com/blog/post/the-50-best-rock-songs-of-2000
	ll := linkedlist.NewLinkedList[Song]()
	ll.Append(&Song{
		title:  "Californication",
		artist: "Red Hot Chili Peppers",
	})
	ll.Append(&Song{
		title:  "It's My Life",
		artist: "Bon Jovi",
	})
	ll.Append(&Song{
		title:  "Beautiful Day",
		artist: "U2",
	})

	fmt.Println("iterate in order")
	iterator := ll.NewIterator()
	for iterator.Next() {
		fmt.Printf("%v\n", iterator.Get())
	}

	fmt.Println("\niterate in reverse order")
	iterator = ll.NewIterator()
	for iterator.Prev() {
		fmt.Printf("%v\n", iterator.Get())
	}
}
```

```txt
iterate in order
&{Californication Red Hot Chili Peppers}
&{It's My Life Bon Jovi}
&{Beautiful Day U2}

iterate in reverse order
&{Beautiful Day U2}
&{It's My Life Bon Jovi}
&{Californication Red Hot Chili Peppers}
```
