package array

/*
Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.
Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

Input:
letters = ["c", "f", "j"]
target = "a"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "c"
Output: "f"

Input:
letters = ["c", "f", "j"]
target = "d"
Output: "f"

Input:
letters = ["c", "f", "j"]
target = "g"
Output: "j"
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	dict := make(map[byte]bool)
	var (
		want byte
		dist int
	)

	for _, letter := range letters {
		dict[letter] = true
	}

	dist = int('z' - target)
	for i := 1; i <= dist; i++ {
		want = target + byte(i)
		if dict[want] {
			goto end
		}
	}

	dist = int(target - 'a')
	for i := 0; i < dist; i++ {
		want = 'a' + byte(i)
		if dict[want] {
			break
		}
	}
end:
	return want
}
