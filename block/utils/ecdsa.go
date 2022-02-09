package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int // 一般的な公開鍵のX座標
	S *big.Int // 署名生成に必要な一時的な秘密鍵や、送信者のトランザクションから導き出された値
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
