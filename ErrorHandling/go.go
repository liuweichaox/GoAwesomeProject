package main

import (
	"errors"
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// go语言的错误，纵观全局就是error接口，只要实现了Error()string方法的任何类型，都属于错误。
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	WrapErrDemo()
	SelfErrDemo()
}

// 错误嵌套，普通类型举例
func WrapErrDemo() {
	errInit := errors.New("这是一个根错误")
	errWrap1 := fmt.Errorf("嵌套1层:%w", errInit)
	errWrap2 := fmt.Errorf("嵌套2层:%w", errWrap1)
	errWrap3 := fmt.Errorf("嵌套3层:%w", errWrap2)
	errWrap4 := fmt.Errorf("嵌套4层:%w", errWrap3)
	fmt.Println("最终错误", errWrap4)

	errUnwrap1 := errors.Unwrap(errWrap4)
	fmt.Println("解开1层", errUnwrap1)

	errUnwrap2 := errors.Unwrap(errUnwrap1)
	fmt.Println("解开2层", errUnwrap2)

	errUnwrap3 := errors.Unwrap(errUnwrap2)
	fmt.Println("解开3层", errUnwrap3)

	errUnwrap4 := errors.Unwrap(errUnwrap3)
	fmt.Println("解开4层", errUnwrap4)

	errUnwrap5 := errors.Unwrap(errUnwrap4)
	fmt.Println("解开5层", errUnwrap5)

	res := errors.Is(errWrap4, errInit)
	fmt.Println(res)
	res = errors.Is(errWrap4, errWrap2)
	fmt.Println(res)

}

// 自定义错误类型
type SelfErr struct {
	msg string
	err string
}

func (s *SelfErr) Error() string {
	return fmt.Sprintf("SelfError msg:%s,err:%s", s.msg, s.err)
}

// 自定义错误类型举例
func SelfErrDemo() {
	err := &SelfErr{
		msg: "自定义错误的msg",
		err: "自定义错误的err",
	}

	errWrap1 := fmt.Errorf("嵌套1层:%w", err)
	errWrap2 := fmt.Errorf("嵌套2层:%w", errWrap1)
	errWrap3 := fmt.Errorf("嵌套3层:%w", errWrap2)
	fmt.Println(errWrap3)

	//错误都是指针类型
	var resErr *SelfErr
	//要拿到原错误，要用错误指针类型的指针
	res := errors.As(errWrap3, &resErr)
	//打印结果
	fmt.Println(res)
	//最终对比
	if resErr == err {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
