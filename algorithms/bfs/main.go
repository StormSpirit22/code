package main

import "fmt"

func main() {
	fmt.Println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
}

func openLock(deadends []string, target string) int {
	deadMap := make(map[string]bool)
	for _, d := range deadends {
		deadMap[d] = true
	}
	start := "0000"

	var lock []string
	lock = append(lock, start)
	steps := 1
	visited := make(map[string]bool)

	for len(lock) > 0 {
		n := len(lock)
		for i := 0; i < n; i++ {
			tmp := lock[i]
			if deadMap[tmp] {
				continue
			}
			for j := range tmp {
				back := []byte(tmp)
				back[j] = up(back[j])
				bs := string(back)
				if deadMap[bs] || visited[bs]{
					continue
				}
				if bs == target {
					return steps
				}
				visited[bs] = true
				lock = append(lock, bs)

				back2 := []byte(tmp)
				back2[j] = down(tmp[j])
				bs2 := string(back2)
				if bs2 == target {
					return steps
				}
				if deadMap[bs2] || visited[bs2] {
					continue
				}
				visited[bs2] = true
				lock = append(lock, bs2)
			}
		}
		lock = lock[n:]
		steps++
	}
	return -1
}

func up(s byte) byte {
	a := s - '0'
	if a != 9 {
		return byte(a+1) + '0'
	}
	return '0'
}

func down(s byte) byte {
	a := s - '0'
	if a != 0 {
		return byte(a-1) + '0'
	}
	return '9'
}
