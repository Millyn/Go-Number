package main

import (
	"GuessGame/core"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var game core.GuessGame
	var gen core.Generator

	// 设置最大数
	gen.SetMaxNumber(100)

	// 设置本次谜底
	game.SetNumber(gen.Next())

	// 设置谜底的最大值和最小值
	game.SetBiggest(gen.GetMaxNumber())
	game.SetSmallest(1)
	game.SetGuessTime(10)

	fmt.Println("开始游戏了~")
	var a int
	for {
		fmt.Printf("请输入 %d 至 %d 的数字，你还剩余 %d 次", game.GetSmallest(), game.GetBiggest(), game.GetGuessTime()-game.GetCurrentGuessTime())
		stdin := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(stdin, &a)
		if err != nil {
			continue
		}
		game.SetGuess(a)
		if !game.IsValidNumberRange() {
			fmt.Println("你的输入有误，请重新输入")
			continue
		}

		game.OnGuess()

		if game.IsGameWon() {
			fmt.Println("你答对了")
			break
		}
		if game.IsGameLost() {
			fmt.Printf("次数耗尽，你失败了，正确答案是 %d", game.GetNumber())
			break
		}

		game.Check()
	}
}
