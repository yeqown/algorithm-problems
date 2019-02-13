package branch_bound

/*
In an election, the i-th vote was cast for persons[i] at time times[i].
Now, we would like to implement the following query function: TopVotedCandidate.q(int t) will return the number of the person that was leading the election at time t.
Votes cast at time t will count towards our query.  In the case of a tie, the most recent vote (among tied candidates) wins.



Example 1:

Input: ["TopVotedCandidate","q","q","q","q","q","q"], [[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]
Output: [null,0,1,1,0,0,1]
Explanation:
	At time 3, the votes are [0], and 0 is leading.
	At time 12, the votes are [0,1,1], and 1 is leading.
	At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
	This continues for 3 more queries at time 15, 24, and 8.


Note:
	1 <= persons.length = times.length <= 5000
	0 <= persons[i] <= persons.length
	times is a strictly increasing array with all elements in [0, 10^9].
	TopVotedCandidate.q is called at most 10000 times per test case.
	TopVotedCandidate.q(int t) is always called with t >= times[0].
*/

type timeLeadingNode struct {
	StartTime int
	EndTime   int
	Leading   int
}

// TopVotedCandidate ...
type TopVotedCandidate struct {
	nodes []timeLeadingNode
	count int
}

// Constructor ...
func Constructor(persons []int, times []int) TopVotedCandidate {
	nodes := make([]timeLeadingNode, len(times))
	dict := make(map[int]int)

	lastLeading := -1 // 前一个leading
	maxVotes := 0     // 当前时间的最多票数
	leading := -1     // 当前leading

	for idx, person := range persons {
		dict[person]++
		if dict[person] >= maxVotes {
			leading = person
			maxVotes = dict[person]
		} else {
			leading = lastLeading
		}
		lastLeading = leading

		// finnale node
		if idx+1 == len(persons) {
			nodes[idx] = timeLeadingNode{
				StartTime: times[idx],
				EndTime:   times[idx] + 1,
				Leading:   leading,
			}
			break
		}

		nodes[idx] = timeLeadingNode{
			StartTime: times[idx],
			EndTime:   times[idx+1] - 1,
			Leading:   leading,
		}
	}

	return TopVotedCandidate{
		nodes: nodes,
		count: len(nodes),
	}
}

// Q ...
func (c *TopVotedCandidate) Q(t int) int {
	leading := -1
	for i := 0; i < c.count; i++ {
		leading = c.nodes[i].Leading
		if t <= c.nodes[i].EndTime {
			// fmt.Println(t, i, c.nodes[i])
			break
		}
	}
	return leading
}
