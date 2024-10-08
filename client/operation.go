package client

import (
	"errors"

	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"

	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/message"
	"github.com/LagrangeDev/LagrangeGo/internal/proto"

	"github.com/LagrangeDev/LagrangeGo/client/entity"
	oidb2 "github.com/LagrangeDev/LagrangeGo/client/packets/oidb"
)

// FetchFriends 获取好友列表信息，使用token可以获取下一页的群成员信息
func (c *QQClient) FetchFriends(token uint32) ([]*entity.Friend, uint32, error) {
	pkt, err := oidb2.BuildFetchFriendsReq(token)
	if err != nil {
		return nil, 0, err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, 0, err
	}
	friends, token, err := oidb2.ParseFetchFriendsResp(resp)
	if err != nil {
		return nil, 0, err
	}
	return friends, token, nil
}

// FetchGroups 获取所有已加入的群的信息
func (c *QQClient) FetchGroups() ([]*entity.Group, error) {
	pkt, err := oidb2.BuildFetchGroupsReq()
	if err != nil {
		return nil, err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, err
	}
	groups, err := oidb2.ParseFetchGroupsResp(resp)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// FetchGroupMember 获取对应群的群成员信息
func (c *QQClient) FetchGroupMember(groupUin, memberUin uint32) (*entity.GroupMember, error) {
	pkt, err := oidb2.BuildFetchMemberReq(groupUin, c.GetUid(memberUin, groupUin))
	if err != nil {
		return nil, err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, err
	}
	members, err := oidb2.ParseFetchMemberResp(resp)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// FetchGroupMembers 获取对应群的所有群成员信息，使用token可以获取下一页的群成员信息
func (c *QQClient) FetchGroupMembers(groupUin uint32, token string) ([]*entity.GroupMember, string, error) {
	pkt, err := oidb2.BuildFetchMembersReq(groupUin, token)
	if err != nil {
		return nil, "", err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, "", err
	}
	members, newToken, err := oidb2.ParseFetchMembersResp(resp)
	if err != nil {
		return nil, "", err
	}
	return members, newToken, nil
}

// GroupRemark 设置群聊备注
func (c *QQClient) GroupRemark(groupUin uint32, remark string) error {
	pkt, err := oidb2.BuildGroupRemarkReq(groupUin, remark)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupRemarkResp(resp)
}

// GroupRename 设置群聊名称
func (c *QQClient) GroupRename(groupUin uint32, name string) error {
	pkt, err := oidb2.BuildGroupRenameReq(groupUin, name)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupRenameResp(resp)
}

// GroupMuteGlobal 群全员禁言
func (c *QQClient) GroupMuteGlobal(groupUin uint32, isMute bool) error {
	pkt, err := oidb2.BuildGroupMuteGlobalReq(groupUin, isMute)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupMuteGlobalResp(resp)
}

// GroupMuteMember 禁言群成员
func (c *QQClient) GroupMuteMember(groupUin, uin, duration uint32) error {
	uid := c.GetUid(uin, groupUin)
	if uid == "" {
		return errors.New("uid not found")
	}
	pkt, err := oidb2.BuildGroupMuteMemberReq(groupUin, duration, uid)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupMuteMemberResp(resp)
}

// GroupLeave 退出群聊
func (c *QQClient) GroupLeave(groupUin uint32) error {
	pkt, err := oidb2.BuildGroupLeaveReq(groupUin)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupLeaveResp(resp)
}

// GroupSetAdmin 设置群管理员
func (c *QQClient) GroupSetAdmin(groupUin, uin uint32, isAdmin bool) error {
	uid := c.GetUid(uin, groupUin)
	if uid == "" {
		return errors.New("uid not found")
	}
	pkt, err := oidb2.BuildGroupSetAdminReq(groupUin, uid, isAdmin)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	err = oidb2.ParseGroupSetAdminResp(resp)
	if err != nil {
		return err
	}
	if m := c.GetCachedMemberInfo(uin, groupUin); m != nil {
		m.Permission = entity.Admin
		c.cache.RefreshGroupMember(groupUin, m)
	}

	return nil
}

// GroupRenameMember 设置群成员昵称
func (c *QQClient) GroupRenameMember(groupUin, uin uint32, name string) error {
	uid := c.GetUid(uin, groupUin)
	if uid == "" {
		return errors.New("uid not found")
	}
	pkt, err := oidb2.BuildGroupRenameMemberReq(groupUin, uid, name)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	err = oidb2.ParseGroupRenameMemberResp(resp)
	if err != nil {
		return err
	}
	if m := c.GetCachedMemberInfo(uin, groupUin); m != nil {
		m.MemberCard = name
		c.cache.RefreshGroupMember(groupUin, m)
	}

	return nil
}

// GroupKickMember 踢出群成员，可选是否拒绝加群请求
func (c *QQClient) GroupKickMember(groupUin, uin uint32, rejectAddRequest bool) error {
	uid := c.GetUid(uin, groupUin)
	if uid == "" {
		return errors.New("uid not found")
	}
	pkt, err := oidb2.BuildGroupKickMemberReq(groupUin, uid, rejectAddRequest)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupKickMemberResp(resp)
}

// GroupSetSpecialTitle 设置群成员专属头衔
func (c *QQClient) GroupSetSpecialTitle(groupUin, uin uint32, title string) error {
	uid := c.GetUid(uin, groupUin)
	if uid == "" {
		return errors.New("uid not found")
	}
	pkt, err := oidb2.BuildGroupSetSpecialTitleReq(groupUin, uid, title)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseGroupSetSpecialTitleResp(resp)
}

// GroupPoke 戳一戳群友
func (c *QQClient) GroupPoke(groupUin, uin uint32) error {
	pkt, err := oidb2.BuildGroupPokeReq(groupUin, uin)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParsePokeResp(resp)
}

// FriendPoke 戳一戳好友
func (c *QQClient) FriendPoke(uin uint32) error {
	pkt, err := oidb2.BuildFriendPokeReq(uin)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParsePokeResp(resp)
}

// RecallGroupMessage 撤回群聊消息
func (c *QQClient) RecallGroupMessage(GrpUin, seq uint32) error {
	packet := message.GroupRecallMsg{
		Type:     1,
		GroupUin: GrpUin,
		Field3: &message.GroupRecallMsgField3{
			Sequence: seq,
			Field3:   0,
		},
		Field4: &message.GroupRecallMsgField4{Field1: 0},
	}
	pktData, err := proto.Marshal(&packet)
	if err != nil {
		return err
	}
	resp, err := c.sendUniPacketAndWait("trpc.msg.msg_svc.MsgService.SsoGroupRecallMsg", pktData)
	if err != nil {
		return err
	}
	if len(resp) == 0 {
		return errors.New("empty response data")
	}
	return nil
}

// GetPrivateImageUrl 获取私聊图片下载url
func (c *QQClient) GetPrivateImageUrl(node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildPrivateImageDownloadReq(c.GetUid(c.Uin), node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return "", err
	}
	return oidb2.ParsePrivateImageDownloadResp(resp)
}

// GetGroupImageUrl 获取群聊图片下载url
func (c *QQClient) GetGroupImageUrl(groupUin uint32, node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildGroupImageDownloadReq(groupUin, node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return "", err
	}
	return oidb2.ParseGroupImageDownloadResp(resp)
}

// GetPrivateRecordUrl 获取私聊语音下载url
func (c *QQClient) GetPrivateRecordUrl(node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildPrivateRecordDownloadReq(c.GetUid(c.Uin), node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return "", err
	}
	return oidb2.ParsePrivateRecordDownloadResp(resp)
}

// GetGroupRecordUrl 获取群聊语音下载url
func (c *QQClient) GetGroupRecordUrl(groupUin uint32, node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildGroupRecordDownloadReq(groupUin, node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return "", err
	}
	return oidb2.ParseGroupRecordDownloadResp(resp)
}

// FetchUserInfo 获取用户信息
func (c *QQClient) FetchUserInfo(uid string) (*entity.Friend, error) {
	pkt, err := oidb2.BuildFetchUserInfoReq(uid)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, err
	}
	return oidb2.ParseFetchUserInfoResp(resp)
}

// FetchUserInfoUin 通过uin获取用户信息
func (c *QQClient) FetchUserInfoUin(uin uint32) (*entity.Friend, error) {
	pkt, err := oidb2.BuildFetchUserInfoReq(uin)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, err
	}
	return oidb2.ParseFetchUserInfoResp(resp)
}

// GetGroupSystemMessages 获取加群请求信息
func (c *QQClient) GetGroupSystemMessages(groupUin ...uint32) ([]*entity.GroupJoinRequest, error) {
	pkt, err := oidb2.BuildFetchGroupSystemMessagesReq(20)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return nil, err
	}
	return oidb2.ParseFetchGroupSystemMessagesReq(resp, groupUin...)
}

// SetGroupRequest 处理加群请求
func (c *QQClient) SetGroupRequest(accept bool, sequence uint64, typ uint32, groupUin uint32, message string) error {
	pkt, err := oidb2.BuildSetGroupRequestReq(accept, sequence, typ, groupUin, message)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseSetGroupRequestResp(resp)
}

// SetFriendRequest 处理好友请求
func (c *QQClient) SetFriendRequest(accept bool, targetUid string) error {
	pkt, err := oidb2.BuildSetFriendRequest(accept, targetUid)
	if err != nil {
		return err
	}
	resp, err := c.sendOidbPacketAndWait(pkt)
	if err != nil {
		return err
	}
	return oidb2.ParseSetFriendRequestResp(resp)
}
