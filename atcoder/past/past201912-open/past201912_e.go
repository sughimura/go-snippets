// 17:43 - 18:06(例1->X)
// 18:39 - 19:30(例1->O,WA)
// 20:49 -

package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	// initialize f
	f := make([][]string, n+1)
	for i := 1; i <= n; i++ {
		f[i] = make([]string, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = "N"
		}
	}

	for i := 0; i < q; i++ {
		var command, a, b int
		fmt.Scan(&command)
		if command == 1 {
			fmt.Scan(&a, &b)
		} else {
			fmt.Scan(&a)
		}
		if command == 1 {
			f[a][b] = "Y"
		} else if command == 2 {
			for i := 1; i <= n; i++ {
				if f[i][a] == "Y" {
					f[a][i] = "Y"
				}
			}
		} else {
			wantToFollows := make([]bool, n+1)
			for x := 1; x <= n; x++ {
				if f[a][x] == "Y" {
					for y := 1; y <= n; y++ {
						if f[x][y] == "Y" {
							wantToFollows[y] = true
						}
					}
				}
			}
			for i := 1; i <= n; i++ {
				if wantToFollows[i] {
					f[a][i] = "Y"
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(f[i][j])
			if j == n {
				fmt.Print("\n")
			}
		}
	}
}
