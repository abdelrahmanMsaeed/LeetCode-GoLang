func isAnagram(s string, t string) bool {
    str1 := strings.Split(s, "")
    str2 := strings.Split(t, "")
	sort.Strings(str1)
	sort.Strings(str2)

	return strings.Join(str1, "") == strings.Join(str2, "")    
}