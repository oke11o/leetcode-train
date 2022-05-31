package _461

/**
1461. Check If a String Contains All Binary Codes of Size K
https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
Medium
*/
func hasAllCodes(s string, k int) bool {
	cnt := 1 << k // math.Pow(2, float64(k))
	m := make(map[string]struct{}, cnt)
	for i := k - 1; i < len(s); i++ {
		key := s[i-k+1 : i+1]
		m[key] = struct{}{}
	}

	return len(m) == cnt
}

/**
For example, say s="11010110", and k=3, and we just finish calculating the hash of the first substring:
"110" (hash is 4+2=6, or 110). Now we want to know the next hash, which is the hash of "101".

We can start from the binary form of our hash, which is 110. First, we shift left, resulting 1100.
We do not need the first digit, so it is a good idea to do 1100 & 111 = 100.
The all-one 111 helps us to align the digits. Now we need to apply the lowest digit of "101", which is 1,
to our hash, and by using |, we get 100 | last_digit = 100 | 1 = 101.
*/
func hasAllCodes_hashFun(s string, k int) bool {
	need := 1 << k // math.Pow(2, float64(k))
	allOne := need - 1

	hashVal := 0
	got := make(map[int]struct{}, need)
	for i := k - 1; i < len(s); i++ {
		// calculate hash for s.substr(i-k+1,i+1)
		hashVal = ((hashVal << 1) & allOne) | int(s[i]-'0')
		_, ok := got[hashVal]
		// hash only available when i-k+1 > 0
		if i >= k-1 && !ok {
			got[hashVal] = struct{}{}
			need--
			if need == 0 {
				return true
			}
		}
	}

	return false
}
