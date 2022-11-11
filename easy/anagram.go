package main

func main() {
	isAnagram("anagram", "nagaram")
}

func isAnagram(s string, t string) bool {
	hash := make([]int, 26)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
		hash[t[i]-'a']--
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}

	return true
}
