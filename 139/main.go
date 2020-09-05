// https://leetcode.com/problems/word-break/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	r := make([][]int, 0)
	ls := len(s)
	for i := 0; i < len(wordDict); i++ {
		j := 0
		lw := len(wordDict[i])
		for j < ls {
			k := strings.Index(s[j:], wordDict[i])
			if k < 0 {
				break
			}
			r = append(r, []int{j + k, j + k + lw - 1})
			j += k + 1
		}
	}
	if len(r) < 1 {
		return false
	}
	sort.Slice(r, func(i, j int) bool {
		if r[i][0] == r[j][0] {
			return r[i][1] <= r[j][1]
		}
		return r[i][0] < r[j][0]
	})
	for i := 0; i < len(r); i++ {
		k := ls
		for j := 0; j < i; j++ {
			if r[i][0]-r[j][1] == 1 && r[j][0] < k {
				k = r[j][0]
			}
		}
		if k < ls {
			r[i][0] = k
		}
		if r[i][0] == 0 && r[i][1] == ls-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "cat"}))
	fmt.Println(wordBreak("abcd", []string{"a", "abc", "d", "cd"}))
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
	fmt.Println(wordBreak("fohhemkkaecojceoaejkkoedkofhmohkcjmkggcmnami", []string{"kfomka", "hecagbngambii", "anobmnikj", "c", "nnkmfelneemfgcl", "ah", "bgomgohl", "lcbjbg", "ebjfoiddndih", "hjknoamjbfhckb", "eioldlijmmla", "nbekmcnakif", "fgahmihodolmhbi", "gnjfe", "hk", "b", "jbfgm", "ecojceoaejkkoed", "cemodhmbcmgl", "j", "gdcnjj", "kolaijoicbc", "liibjjcini", "lmbenj", "eklingemgdjncaa", "m", "hkh", "fblb", "fk", "nnfkfanaga", "eldjml", "iejn", "gbmjfdooeeko", "jafogijka", "ngnfggojmhclkjd", "bfagnfclg", "imkeobcdidiifbm", "ogeo", "gicjog", "cjnibenelm", "ogoloc", "edciifkaff", "kbeeg", "nebn", "jdd", "aeojhclmdn", "dilbhl", "dkk", "bgmck", "ohgkefkadonafg", "labem", "fheoglj", "gkcanacfjfhogjc", "eglkcddd", "lelelihakeh", "hhjijfiodfi", "enehbibnhfjd", "gkm", "ggj", "ag", "hhhjogk", "lllicdhihn", "goakjjnk", "lhbn", "fhheedadamlnedh", "bin", "cl", "ggjljjjf", "fdcdaobhlhgj", "nijlf", "i", "gaemagobjfc", "dg", "g", "jhlelodgeekj", "hcimohlni", "fdoiohikhacgb", "k", "doiaigclm", "bdfaoncbhfkdbjd", "f", "jaikbciac", "cjgadmfoodmba", "molokllh", "gfkngeebnggo", "lahd", "n", "ehfngoc", "lejfcee", "kofhmoh", "cgda", "de", "kljnicikjeh", "edomdbibhif", "jehdkgmmofihdi", "hifcjkloebel", "gcghgbemjege", "kobhhefbbb", "aaikgaolhllhlm", "akg", "kmmikgkhnn", "dnamfhaf", "mjhj", "ifadcgmgjaa", "acnjehgkflgkd", "bjj", "maihjn", "ojakklhl", "ign", "jhd", "kndkhbebgh", "amljjfeahcdlfdg", "fnboolobch", "gcclgcoaojc", "kfokbbkllmcd", "fec", "dljma", "noa", "cfjie", "fohhemkka", "bfaldajf", "nbk", "kmbnjoalnhki", "ccieabbnlhbjmj", "nmacelialookal", "hdlefnbmgklo", "bfbblofk", "doohocnadd", "klmed", "e", "hkkcmbljlojkghm", "jjiadlgf", "ogadjhambjikce", "bglghjndlk", "gackokkbhj", "oofohdogb", "leiolllnjj", "edekdnibja", "gjhglilocif", "ccfnfjalchc", "gl", "ihee", "cfgccdmecem", "mdmcdgjelhgk", "laboglchdhbk", "ajmiim", "cebhalkngloae", "hgohednmkahdi", "ddiecjnkmgbbei", "ajaengmcdlbk", "kgg", "ndchkjdn", "heklaamafiomea", "ehg", "imelcifnhkae", "hcgadilb", "elndjcodnhcc", "nkjd", "gjnfkogkjeobo", "eolega", "lm", "jddfkfbbbhia", "cddmfeckheeo", "bfnmaalmjdb", "fbcg", "ko", "mojfj", "kk", "bbljjnnikdhg", "l", "calbc", "mkekn", "ejlhdk", "hkebdiebecf", "emhelbbda", "mlba", "ckjmih", "odfacclfl", "lgfjjbgookmnoe", "begnkogf", "gakojeblk", "bfflcmdko", "cfdclljcg", "ho", "fo", "acmi", "oemknmffgcio", "mlkhk", "kfhkndmdojhidg", "ckfcibmnikn", "dgoecamdliaeeoa", "ocealkbbec", "kbmmihb", "ncikad", "hi", "nccjbnldneijc", "hgiccigeehmdl", "dlfmjhmioa", "kmff", "gfhkd", "okiamg", "ekdbamm", "fc", "neg", "cfmo", "ccgahikbbl", "khhoc", "elbg", "cbghbacjbfm", "jkagbmfgemjfg", "ijceidhhajmja", "imibemhdg", "ja", "idkfd", "ndogdkjjkf", "fhic", "ooajkki", "fdnjhh", "ba", "jdlnidngkfffbmi", "jddjfnnjoidcnm", "kghljjikbacd", "idllbbn", "d", "mgkajbnjedeiee", "fbllleanknmoomb", "lom", "kofjmmjm", "mcdlbglonin", "gcnboanh", "fggii", "fdkbmic", "bbiln", "cdjcjhonjgiagkb", "kooenbeoongcle", "cecnlfbaanckdkj", "fejlmog", "fanekdneoaammb", "maojbcegdamn", "bcmanmjdeabdo", "amloj", "adgoej", "jh", "fhf", "cogdljlgek", "o", "joeiajlioggj", "oncal", "lbgg", "elainnbffk", "hbdi", "femcanllndoh", "ke", "hmib", "nagfahhljh", "ibifdlfeechcbal", "knec", "oegfcghlgalcnno", "abiefmjldmln", "mlfglgni", "jkofhjeb", "ifjbneblfldjel", "nahhcimkjhjgb", "cdgkbn", "nnklfbeecgedie", "gmllmjbodhgllc", "hogollongjo", "fmoinacebll", "fkngbganmh", "jgdblmhlmfij", "fkkdjknahamcfb", "aieakdokibj", "hddlcdiailhd", "iajhmg", "jenocgo", "embdib", "dghbmljjogka", "bahcggjgmlf", "fb", "jldkcfom", "mfi", "kdkke", "odhbl", "jin", "kcjmkggcmnami", "kofig", "bid", "ohnohi", "fcbojdgoaoa", "dj", "ifkbmbod", "dhdedohlghk", "nmkeakohicfdjf", "ahbifnnoaldgbj", "egldeibiinoac", "iehfhjjjmil", "bmeimi", "ombngooicknel", "lfdkngobmik", "ifjcjkfnmgjcnmi", "fmf", "aoeaa", "an", "ffgddcjblehhggo", "hijfdcchdilcl", "hacbaamkhblnkk", "najefebghcbkjfl", "hcnnlogjfmmjcma", "njgcogemlnohl", "ihejh", "ej", "ofn", "ggcklj", "omah", "hg", "obk", "giig", "cklna", "lihaiollfnem", "ionlnlhjckf", "cfdlijnmgjoebl", "dloehimen", "acggkacahfhkdne", "iecd", "gn", "odgbnalk", "ahfhcd", "dghlag", "bchfe", "dldblmnbifnmlo", "cffhbijal", "dbddifnojfibha", "mhh", "cjjol", "fed", "bhcnf", "ciiibbedklnnk", "ikniooicmm", "ejf", "ammeennkcdgbjco", "jmhmd", "cek", "bjbhcmda", "kfjmhbf", "chjmmnea", "ifccifn", "naedmco", "iohchafbega", "kjejfhbco", "anlhhhhg"}))
	fmt.Println(wordBreak("bccdbacdbdacddabbaaaadababadad", []string{"cbc", "bcda", "adb", "ddca", "bad", "bbb", "dad", "dac", "ba", "aa", "bd", "abab", "bb", "dbda", "cb", "caccc", "d", "dd", "aadb", "cc", "b", "bcc", "bcd", "cd", "cbca", "bbd", "ddd", "dabb", "ab", "acd", "a", "bbcc", "cdcbd", "cada", "dbca", "ac", "abacd", "cba", "cdb", "dbac", "aada", "cdcda", "cdc", "dbc", "dbcb", "bdb", "ddbdd", "cadaa", "ddbc", "babb"}))
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aa"}))
}
