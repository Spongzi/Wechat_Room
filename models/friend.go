package models

// Friend 朋友表要完成的内容有：
/*
	1. 要存储自己的Id，和朋友的Id
	2. 是否为好友 1 为好友 0 为非好友
	3. 添加好友的时候的备注
	4. 好友的备注(user 对 friend)
	5. 好友的备注(friend 对 user)
*/
type Friend struct {
	UserId       uint   `json:"user_id" gorm:"type:int"`                // 用于存储自己的Id
	FriendId     uint   `json:"friend_id" gorm:"type:int"`              // 用于存储朋友的Id
	IsFriends    int    `json:"is_friends" gorm:"type:int"`             // 是否为好朋友，即有没有同意添加好友
	IsAgree      int    `json:"is_agree" gorm:"type:int"`               // 是否同意
	AddFriendMsg string `json:"add_friend_msg" gorm:"type:varchar(64)"` // 添加好友的备注信息
	RemarkUser   string `json:"remark_user" gorm:"type:varchar(64)"`    // 好友的备注 user --> friend
	RemarkFriend string `json:"remark_friend" gorm:"type:varchar(64)"`  // 好友的备注 friend --> user
}
