package core

import (
	"math/rand"
	"time"
)

type NumberGenerator interface {
	Next() int               // 获取一个随机数
	SetMaxNumber(number int) // 设置最大的数随机数
	GetMaxNumber() int       // 获取最大的数 最大的数
}

type Game interface {
	SetNumber(number int)     // 设置谜底数字
	GetNumber() int           // 获取谜底数字
	GetGuess() int            // 获取本地猜测的数字
	SetGuess(guess int)       // 设置本次猜测的数字
	GetSmallest() int         // 获取最小数
	SetSmallest(number int)   // 设置最小数
	GetBiggest() int          // 获取最大数
	SetBiggest(number int)    // 设置最大值
	SetGuessTime(number int)  // 设置猜测次数
	GetGuessTime() int		  // 获取本轮猜测总次数
	GetCurrentGuessTime() int // 当前猜测次数
	OnGuess()			  // 消耗一次答题
	Reset()                   // 重置
	Check()                   // 检查猜测
	IsValidNumberRange() bool // 输入是否在有效范围
	IsGameWon() bool          // 游戏是否成功
	IsGameLost() bool         // 游戏是否失败
}

type Generator struct {
	maxNumber int
}

func (gen *Generator) Next() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(gen.GetMaxNumber()) + 1
}

func (gen *Generator) GetMaxNumber() int {
	return gen.maxNumber
}

func (gen *Generator) SetMaxNumber(number int) {
	gen.maxNumber = number
}

type GuessGame struct {
	number             int
	count              int
	currentGuessNumber int
	currentCount       int
	answer             []int
	smallest           int
	biggest            int
}

func (g *GuessGame) SetSmallest(number int) {
	g.smallest = number
}

func (g *GuessGame) SetBiggest(number int) {
	g.biggest = number
}

func (g *GuessGame) SetNumber(number int) {
	g.number = number
}

func (g *GuessGame) GetNumber() int {
	return g.number
}

func (g *GuessGame) GetGuess() int {
	return g.currentGuessNumber
}

func (g *GuessGame) SetGuess(guess int) {
	g.currentGuessNumber = guess
}

func (g *GuessGame) GetSmallest() int {
	return g.smallest
}

func (g *GuessGame) GetBiggest() int {
	return g.biggest
}

func (g *GuessGame) SetGuessTime(number int) {
	g.count = number
}

func (g *GuessGame) GetCurrentGuessTime() int {
	return g.currentCount
}

func (g *GuessGame) GetGuessTime() int {
	return g.count
}

func (g *GuessGame) Reset() {

}

func (g *GuessGame) Check() {
	if g.currentGuessNumber >= g.number {
		g.biggest = g.currentGuessNumber
	} else {
		g.smallest = g.currentGuessNumber
	}
}

func (g *GuessGame) IsValidNumberRange() bool {
	if g.currentGuessNumber <= g.biggest && g.currentGuessNumber >= g.smallest {
		return true
	} else {
		return false
	}
}

func (g *GuessGame) IsGameWon() bool {
	if g.currentGuessNumber == g.number {
		return true
	}
	return false
}

func (g *GuessGame) IsGameLost() bool {
	if g.currentCount == g.count {
		return true
	}
	return false
}

func (g *GuessGame) OnGuess() {
	g.currentCount++
}
