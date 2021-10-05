package main

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

// 核心想法：使用peek操作时标记一下是否进行过peek操作，如果进行过peek操作，那么下一次的next就不需要进行了，只要返回当前保存的值就可以了
type PeekingIterator struct {
	iter      *Iterator
	hasPeeked bool
	peekedVal int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasPeeked || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if !this.hasPeeked {
		return this.iter.next()
	}
	this.hasPeeked = false
	return this.peekedVal
}

func (this *PeekingIterator) peek() int {
	if this.hasPeeked {
		return this.peekedVal
	}
	this.peekedVal = this.iter.next()
	this.hasPeeked = true
	return this.peekedVal
}
