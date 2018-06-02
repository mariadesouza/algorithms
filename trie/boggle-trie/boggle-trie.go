/*
Given a dictionary, a method to do lookup in dictionary
and a M x N board where every cell has one character.
Find all possible words that can be formed by a sequence of adjacent characters.
Note that we can move to any of 8 adjacent characters,
but a word should not have multiple instances of same cell.

Example:

Input: dictionary[] = {"GEEKS", "FOR", "QUIZ", "GO"};
       boggle[][]   = {{'G','I','Z'},
                       {'U','E','K'},
                       {'Q','S','E'}};

Output:  Following words of dictionary are present
         GEEKS, QUIZ

Input:
The first line of input contains an integer T denoting the no of test cases . Then T test cases follow. Each test case contains an integer x denoting the no of words in the dictionary. Then in the next line are x space separated strings denoting the contents of the dictinory. In the next line are two integers N and M denoting the size of the boggle. The last line of each test case contains NxM space separated values of the boggle.

Output:
For each test case in a new line print the space separated sorted distinct words of the dictionary which could be formed from the boggle. If no word can be formed print -1.

Constraints:
1<=T<=10
1<=x<=10
1<=n,m<=7

Example:
Input:
1
4
GEEKS FOR QUIZ GO
3 3
G I Z U E K Q S E

Output:
GEEKS QUIZ

** For More Input/Output Examples Use 'Expected Output' option **
*/

package main

import (
	"fmt"

	"github.com/mariadesouza/data-structures-and-alogorithms/trie"
)

const (
	M = 3
	N = 3
)

func main() {
	fmt.Println("Start")

	dictionary := []string{"GEEKS", "FOR", "QUIZ", "GO"}
	root := trie.NewTrieNode()
	for _, word := range dictionary {
		root.Insert(word)
	}

	boggle := [M][N]rune{{'G', 'I', 'Z'},
		{'U', 'E', 'K'},
		{'Q', 'S', 'E'},
	}
	//search on the board

	var visited [M][N]bool

	crawler := root
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			index := boggle[i][j] - 'A'
			//fmt.Println(string(boggle[i][j]))
			if crawler.Links[index] != nil {
				fmt.Println(string(boggle[i][j]))
				searchWord(crawler.Links[index], boggle, i, j, string(boggle[i][j]), visited)
			}
		}
	}
}

func isSafe(i int, j int) bool {
	if i >= 0 && j >= 0 && i < M && j < N {
		return true
	}
	return false
}

func searchWord(root *trie.TrieNode, boggle [M][N]rune, i int, j int, dictstr string, visited [M][N]bool) {

	//fmt.Println(dictstr)
	if root.IsEndofWord == true && len(dictstr) > 0 {
		fmt.Println(dictstr)
	}

	if isSafe(i, j) && !visited[i][j] {

		visited[i][j] = true

		for k := 0; k < trie.SIZE; k++ {
			if root.Links[k] != nil {
				ch := k + 'A'
				if isSafe(i-1, j) && !visited[i-1][j] && boggle[i-1][j] == rune(ch) {
					searchWord(root.Links[k], boggle, i-1, j, dictstr+string(ch), visited)
				}
				if isSafe(i, j-1) && !visited[i][j-1] && boggle[i][j-1] == rune(ch) {
					searchWord(root.Links[k], boggle, i, j-1, dictstr+string(ch), visited)
				}
				if isSafe(i+1, j) && !visited[i+1][j] && boggle[i+1][j] == rune(ch) {
					searchWord(root.Links[k], boggle, i+1, j, dictstr+string(ch), visited)
				}
				if isSafe(i, j+1) && !visited[i][j+1] && boggle[i][j+1] == rune(ch) {
					searchWord(root.Links[k], boggle, i, j+1, dictstr+string(ch), visited)
				}
				if isSafe(i+1, j+1) && !visited[i+1][j+1] && boggle[i+1][j+1] == rune(ch) {
					searchWord(root.Links[k], boggle, i, j+1, dictstr+string(ch), visited)
				}
				if isSafe(i-1, j-1) && !visited[i-1][j-1] && boggle[i-1][j-1] == rune(ch) {
					searchWord(root.Links[k], boggle, i-1, j-1, dictstr+string(ch), visited)
				}
				if isSafe(i-1, j+1) && !visited[i-1][j+1] && boggle[i-1][j+1] == rune(ch) {
					searchWord(root.Links[k], boggle, i-1, j+1, dictstr+string(ch), visited)
				}
				if isSafe(i+1, j-1) && !visited[i+1][j-1] && boggle[i+1][j-1] == rune(ch) {
					searchWord(root.Links[k], boggle, i+1, j-1, dictstr+string(ch), visited)
				}
			}
		}
		visited[i][j] = false
	}

}
