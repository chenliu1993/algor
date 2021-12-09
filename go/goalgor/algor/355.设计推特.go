/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 */

// @lc code=start
type Message struct {
	ID     int
	UserID int
	Time   int
}

var counter int

type User struct {
	ID int // Should be the only one
	// followers => ID => followees
	Followees map[int]bool // These are followed by current User
}

type Twitter struct {
	Users    map[int]User // UserID:User
	Messages []Message
}

func Constructor() Twitter {
	return Twitter{
		Users:    map[int]User{},
		Messages: []Message{},
	}
}

func (this *Twitter) MakeUserAvailable(userId int) {
	if _, ok := this.Users[userId]; !ok {
		user := User{
			ID: userId,
			Followees: map[int]bool{
				userId: true,
			},
		}
		this.Users[userId] = user
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	counter++
	this.Messages = append(this.Messages, Message{
		ID:     tweetId,
		UserID: userId,
		Time:   counter,
	})
	sort.Slice(this.Messages, func(i, j int) bool {
		return this.Messages[i].Time > this.Messages[j].Time
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.MakeUserAvailable(userId)
	user := this.Users[userId]
	ans := []int{}
	for _, v := range this.Messages {
		if _, ok := user.Followees[v.UserID]; ok {
			ans = append(ans, v.ID)
		}
		if len(ans) == 10 {
			break
		}
	}
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.MakeUserAvailable(followerId)
	userer := this.Users[followerId]
	userer.Followees[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.MakeUserAvailable(followerId)
	userer := this.Users[followerId]
	delete(userer.Followees, followeeId)
	this.Users[followerId] = userer
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

