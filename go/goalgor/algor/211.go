type WordDictionary struct {
	Words map[int][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		Words: map[int][]string{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	n := len(word)
	if _, ok := this.Words[n]; !ok {
		this.Words[n] = []string{word}
	} else {
		this.Words[n] = append(this.Words[n], word)
	}
}

func (this *WordDictionary) Search(word string) bool {
	n := len(word)
	wList, ok := this.Words[n]
	if !ok {
		return false
	}
	matcher := regexp.MustCompile("^" + word + "$")
	for _, v := range wList {
		if matcher.MatchString(v) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
