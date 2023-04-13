package models

// 分配请求
type GetRoomIdRequest struct {
	Uid  int64  `json:"uid"`  // 用户ID
	Name string `json:"name"` // 名字
}

// 分配响应
type GetRoomIdRespones struct {
	Uid    int64 `json:"uid"`     // 用户ID
	RoomId int64 `json:"room_id"` // 房间ID
}
