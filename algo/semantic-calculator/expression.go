package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	lOpenParenthese  = 1
	lCloseParenthese = 2
	lOperator        = 3
	lOperand         = 4
)

const (
	opAddition        = "+"
	opSubtraction     = "-"
	opDivision        = ":"
	opMultiplication  = "*"
	opOpenParenthese  = "("
	opCloseParenthese = ")"
)

type Operator struct {
	s        string
	priority uint8
}

var opMap = map[string]*Operator{
	opAddition:        {s: opAddition, priority: 1},
	opDivision:        {s: opDivision, priority: 2},
	opMultiplication:  {s: opMultiplication, priority: 2},
	opSubtraction:     {s: opSubtraction, priority: 1},
	opOpenParenthese:  {s: opOpenParenthese, priority: 1},
	opCloseParenthese: {s: opCloseParenthese, priority: 1},
}

func isOperator(s string) bool {
	_, found := opMap[s]
	return found
}

func canEval(op, prevOp *Operator) bool {

	if prevOp == nil {
		return true
	}

	cannot := (op.priority < prevOp.priority) ||
				(op.s == opSubtraction && prevOp.s == opSubtraction) ||
					(op.s == opDivision && prevOp.s == opDivision)
	return !cannot

}

func (o *Operator) eval(leftOperand, rightOperand float64) float64 {

	res := 0.0

	switch o.s {
	case opAddition:
		res = leftOperand + rightOperand
	case opSubtraction:
		res = leftOperand - rightOperand
	case opMultiplication:
		res = leftOperand * rightOperand
	case opDivision:
		res = leftOperand / rightOperand
	}
	return res
}

func getOp(s string) *Operator {
	if o, ok := opMap[s]; ok {
		return o
	}
	return nil
}

type Token struct {
	lType int
	value []rune
	str string
}

func valueToOperand(value float64) *Token {
	return &Token{
		lType: lOperand,
		value: []rune(fmt.Sprintf("%f", value)),
		str: fmt.Sprintf("%f", value),
	}
}

func (t *Token) toValue() (*float64, error) {
	if t.lType == lOperand {
		v, err := strconv.ParseFloat(string(t.value), 64)
		if err != nil {
			return nil, err
		}
		return &v, nil
	}
	return nil, errors.New("not an operand")
}

func (l *Token) add(r rune) {
	l.value = append(l.value, r)
	l.str = string(l.value)
}

func parseTokens(runes []rune) ([]*Token, error) {
	var tokens []*Token
	var token *Token

	for _, r := range runes {

		if token == nil {
			token = &Token{}
			tokens = append(tokens, token)
		}

		if unicode.IsDigit(r) {
			token.add(r)
			token.lType = lOperand
		} else {

			if token.lType == lOperand {
				token = &Token{}
				tokens = append(tokens, token)
			}
			token.add(r)
			if r == '(' {
				token.lType = lOpenParenthese
			} else if r == ')' {
				token.lType = lCloseParenthese
			} else if isOperator(string(r)) {
				token.lType = lOperator
			} else {
				return nil, errors.New("illegal symbol")
			}
			token = nil

		}

	}

	return tokens, nil

}

func evalRecurse(opToken *Token, opStack, operandStack *Stack) error {

	op := getOp(string(opToken.value))

	var prevOp *Operator
	if p := opStack.Peek(); p != nil {
		prevOp = getOp(string(opStack.Peek().(*Token).value))
	}

	rightOperandToken := operandStack.Pop().(*Token)
	leftOperandToken := operandStack.Pop().(*Token)

	if canEval(op, prevOp) {

		leftValue, err := leftOperandToken.toValue()
		if err != nil {
			return err
		}
		rightValue, err := rightOperandToken.toValue()
		if err != nil {
			return err
		}
		operandStack.Push(valueToOperand(op.eval(*leftValue, *rightValue)))

	} else {

		operandStack.Push(leftOperandToken)

		prevOpToken := opStack.Pop().(*Token)

		err := evalRecurse(prevOpToken, opStack, operandStack)
		if err != nil {
			return err
		}

		opStack.Push(opToken)
		operandStack.Push(rightOperandToken)

	}
	return nil
}


func evaluate(opStack, operandStack *Stack, tokens []*Token) (*float64, error) {

	for _, t := range tokens {

		if t.lType == lOperand {
			operandStack.Push(t)
		}

		if t.lType == lOperator || t.lType == lOpenParenthese {
			opStack.Push(t)
		}

		if t.lType == lCloseParenthese {

			for {

				opToken := opStack.Pop().(*Token)

				if opToken.lType == lOpenParenthese {
					break
				}

				if err := evalRecurse(opToken, opStack, operandStack); err != nil {
					return nil, err
				}

			}

		}

	}

	res, _ := operandStack.Pop().(*Token).toValue()
	return res, nil

}

func parseAndEval(expr string) (*float64, error) {

	runes := []rune(fmt.Sprintf("(%s)", strings.ReplaceAll(expr, " ", "")))

	tokens, err := parseTokens(runes)
	if err != nil {
		return nil, err
	}

	opStack := NewStack()
	operandStack := NewStack()

	return evaluate(opStack, operandStack, tokens)

}
