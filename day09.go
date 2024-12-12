package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := []byte(scanner.Text())
	data = append(data, '0')
	type file struct {
		id    int
		dat   int
		space int
		prev  *file
		next  *file
	}
	var head, tail *file
	fileId := 0
	for i := 0; i < len(data); i += 2 {
		a, b := int(data[i]-'0'), int(data[i+1]-'0')
		newFile := &file{fileId, a, b, nil, nil}
		fileId++
		if head == nil {
			head = newFile
		}
		if tail == nil {
			tail = newFile
		} else {
			newFile.prev = tail
			tail.next = newFile
			tail = newFile
		}
	}
	left, right := head, tail
	for left != nil && right != nil && left != right {
		for left != nil && left.space == 0 {
			left = left.next
		}
		for right != nil && right.dat == 0 {
			right = right.prev
		}
		if left == nil || right == nil || left == right {
			break
		}
		if right.dat >= left.space {
			newFile := &file{right.id, left.space, 0, nil, nil}
			next := left.next
			prev := left
			prev.next = newFile
			newFile.prev = prev
			if next != nil {
				next.prev = newFile
			}
			newFile.next = next
			right.dat -= left.space
			left.space = 0
		} else if right.dat < left.space {
			newFile := &file{right.id, right.dat, left.space - right.dat, nil, nil}
			next := left.next
			prev := left
			prev.next = newFile
			newFile.prev = prev
			if next != nil {
				next.prev = newFile
			}
			newFile.next = next
			right.dat = 0
			left.space = 0
		}
		/*pos = 0
		for tmp := head; tmp != nil; tmp = tmp.next {
			for i := pos; i < pos + tmp.dat; i++ {
				fmt.Print(tmp.id)
			}
			for i := 0; i < tmp.space; i++ {
				fmt.Print(".")
			}
			pos += tmp.dat + tmp.space
		}
		fmt.Println()*/
	}
	pos := 0
	p1 := int64(0)
	for tmp := head; tmp != nil; tmp = tmp.next {
		for i := pos; i < pos+tmp.dat; i++ {
			p1 += int64(tmp.id * i)
		}
		pos += tmp.dat + tmp.space
	}
	fmt.Println("Part 1:", p1)

	head, tail = nil, nil
	fileId = 0
	for i := 0; i < len(data); i += 2 {
		a, b := int(data[i]-'0'), int(data[i+1]-'0')
		newFile := &file{fileId, a, b, nil, nil}
		fileId++
		if head == nil {
			head = newFile
		}
		if tail == nil {
			tail = newFile
		} else {
			newFile.prev = tail
			tail.next = newFile
			tail = newFile
		}
	}
	fileId--
	left, right = head, tail
	for left != nil && right != nil {
		for right = head; right != nil; right = right.next {
			if right.id == fileId {
				break
			}
		}
		fileId--
		for left != nil && left.space == 0 {
			left = left.next
		}
		if left == nil || right == nil {
			break
		}
		found := false
		for tmp := left; tmp != nil; tmp = tmp.next {
			if tmp == right {
				found = true
				break
			}
		}
		if !found {
			break
		}
		if left != right && right.dat <= left.space {
			newFile := &file{right.id, right.dat, left.space - right.dat, nil, nil}
			next := left.next
			prev := left
			prev.next = newFile
			newFile.prev = prev
			if next != nil {
				next.prev = newFile
			}
			newFile.next = next
			right.dat = 0
			right.space += newFile.dat
			left.space = 0
		} else if left != right && right.dat > left.space {
			var left2 *file
			for left2 = left; left2 != nil && left2 != right; left2 = left2.next {
				if right.dat <= left2.space {
					break
				}
			}
			if left2 != nil && left2 != right {
				newFile := &file{right.id, right.dat, left2.space - right.dat, nil, nil}
				next := left2.next
				prev := left2
				prev.next = newFile
				newFile.prev = prev
				if next != nil {
					next.prev = newFile
				}
				newFile.next = next
				right.dat = 0
				right.space += newFile.dat
				left2.space = 0
			}
		}
		/*pos = 0
		for tmp := head; tmp != nil; tmp = tmp.next {
			for i := pos; i < pos + tmp.dat; i++ {
				fmt.Print(tmp.id)
			}
			for i := 0; i < tmp.space; i++ {
				fmt.Print(".")
			}
			pos += tmp.dat + tmp.space
		}
		fmt.Println()*/
	}
	pos = 0
	p2 := int64(0)
	for tmp := head; tmp != nil; tmp = tmp.next {
		for i := pos; i < pos+tmp.dat; i++ {
			p2 += int64(tmp.id * i)
		}
		pos += tmp.dat + tmp.space
	}
	fmt.Println("Part 2:", p2)
}
