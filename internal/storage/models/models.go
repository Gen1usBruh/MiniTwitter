package models

type SignInStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FollowStruct struct {
	FollowingID int32 `json:"following_id"`
}

type CreateTweetStruct struct {
	Content string `json:"content"`
	Media   []byte `json:"media"`
}

type CreateRetweetStruct struct {
	ParentTweetID   int32  `json:"parent_tweet_id"`
	ParentRetweetID int32  `json:"parent_retweet_id"`
	IsQuote         bool   `json:"is_quote"`
	Quote           string `json:"quote"`
}

type SelectTweetStruct struct {
	TweetID int32 `json:"id"`
}

type BaseResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error,omitempty"`
	ErrorCode int    `json:"errorCode"`
}
