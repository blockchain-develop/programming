# 最长无重复字符子串

s为字符串，s[i,j]表示为s中从i到j索引的子串，LNRS(i,j)表示为子串s[i,j]是否为无重复字符子串，则有
if LNRS(i, j-1)为无重复字符子串且s[j]不在s[i,j-1]中 
    LNRS(i,j) = true 
else
    LNRS(i,j) = false 

