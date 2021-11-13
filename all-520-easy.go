package main

import "unicode"

func detectCapitalUse(word string) bool { 
    n := len(word)
    // 如果只有一个字符一定是对的，直接返回
    if n == 1 {
        return true
    }

    // 根据首字母是否大写划分
    if unicode.IsUpper(rune(word[0])) {
        // 第二个字符判断，是全大写，还是首字母大写其余小写
        if unicode.IsUpper(rune(word[1])) {
            for i := 2; i < n; i++ {
                if unicode.IsLower(rune(word[i])) {
                    return false
                }
            }
        } else {
            for i := 2; i < n; i++ {
                if unicode.IsUpper(rune(word[i])) {
                    return false
                }
            }
        }
    } else {
        // 全小写的情况
        for i := 1; i < n; i++ {
            if unicode.IsUpper(rune(word[i])) {
                return false
            }
        }
    }

    return true
}