# 最长回文子串

字符串为s，长度为n，记LPS(i,j)为子串s[i,j]是否为回文子串。则有
LPS(i,j) = LPS(i+1,j-1)+2 if s[i]=s[j] & LPS(i+1,j-1)为回文子串 or 不是回文子串

## 参考

最长回文子串wiki https://en.wikipedia.org/wiki/Longest_palindromic_substring
