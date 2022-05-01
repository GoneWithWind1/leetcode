package weekgame0501

import "strings"
/*
6047. 移除指定数字得到的最大结果 显示英文描述
给你一个表示某个正整数的字符串 number 和一个字符 digit 。

从 number 中 恰好 移除 一个 等于 digit 的字符后，找出并返回按 十进制 表示 最大 的结果字符串。
生成的测试用例满足 digit 在 number 中出现至少一次。
输入：number = "123", digit = "3"
输出："12"
解释："123" 中只有一个 '3' ，在移除 '3' 之后，结果为 "12" 。

*/

func removeDigit(number string, digit byte) string {
    removeIndex := -1
    for index , ch := range number{
        if byte(ch) != digit {
            continue
        }
        if index < len(number) -1 && byte(number[index+1]) < digit {
            removeIndex = index
            break
        }
        removeIndex = index
    }
    return string(number[:removeIndex]) + string(number[removeIndex+1:])
}

