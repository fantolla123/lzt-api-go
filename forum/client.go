package forum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"

	lztapi "github.com/fantolla123/lzt-api-go"
)

type Client struct {
	*lztapi.Client
}

func NewClient(token string, opts ...lztapi.Option) *Client {
	return &Client{
		Client: lztapi.NewClient(token, "https://prod-api.lolz.live", opts...),
	}
}

func (c *Client) ConversationsGet(conversationId int64, opts ...func(*ConversationsGetParams)) (map[string]any, error) {
	p := &ConversationsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsGetParams struct {
}

func (c *Client) SearchTagged(opts ...func(*SearchTaggedParams)) (map[string]any, error) {
	p := &SearchTaggedParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/tagged"


	bodyMap := map[string]any{}
	if p.Tag != nil {
		bodyMap["tag"] = *p.Tag
	}
	if p.Tags != nil {
		bodyMap["tags"] = *p.Tags
	}
	if p.Page != nil {
		bodyMap["page"] = *p.Page
	}
	if p.Limit != nil {
		bodyMap["limit"] = *p.Limit
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type SearchTaggedParams struct {
	Tag *string
	Tags *[]string
	Page *int64
	Limit *int64
}

func (c *Client) LinksList(opts ...func(*LinksListParams)) (*LinksListResponse, error) {
	p := &LinksListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/link-forums"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result LinksListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type LinksListParams struct {
}

func (c *Client) LinksGet(linkId int64, opts ...func(*LinksGetParams)) (*LinksGetResponse, error) {
	p := &LinksGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/link-forums/{link_id}"
	path = strings.Replace(path, "{link_id}", fmt.Sprintf("%v", linkId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result LinksGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type LinksGetParams struct {
}

func (c *Client) PostsCommentsReport(postCommentId int64, message string, opts ...func(*PostsCommentsReportParams)) (map[string]any, error) {
	p := &PostsCommentsReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments/report"


	bodyMap := map[string]any{}
	bodyMap["post_comment_id"] = postCommentId
	bodyMap["message"] = message
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsCommentsReportParams struct {
}

func (c *Client) UsersFollowers(userId string, opts ...func(*UsersFollowersParams)) (map[string]any, error) {
	p := &UsersFollowersParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/followers"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersFollowersParams struct {
	Order *string
	Page *int64
	Limit *int64
}

func (c *Client) UsersFollow(userId string, opts ...func(*UsersFollowParams)) (map[string]any, error) {
	p := &UsersFollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/followers"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersFollowParams struct {
}

func (c *Client) UsersUnfollow(userId string, opts ...func(*UsersUnfollowParams)) (map[string]any, error) {
	p := &UsersUnfollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/followers"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersUnfollowParams struct {
}

func (c *Client) UsersLikes(userId string, opts ...func(*UsersLikesParams)) (*UsersLikesResponse, error) {
	p := &UsersLikesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/likes"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.NodeId != nil {
		params.Set("node_id", fmt.Sprintf("%v", *p.NodeId))
	}
	if p.LikeType != nil {
		params.Set("like_type", fmt.Sprintf("%v", *p.LikeType))
	}
	if p.TypeValue != nil {
		params.Set("type", fmt.Sprintf("%v", *p.TypeValue))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.ContentType != nil {
		params.Set("content_type", fmt.Sprintf("%v", *p.ContentType))
	}
	if p.SearchUserId != nil {
		params.Set("search_user_id", fmt.Sprintf("%v", *p.SearchUserId))
	}
	if p.Stats != nil {
		params.Set("stats", fmt.Sprintf("%v", *p.Stats))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result UsersLikesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersLikesParams struct {
	NodeId *int64
	LikeType *string
	TypeValue *string
	Page *int64
	ContentType *string
	SearchUserId *int64
	Stats *bool
}

func (c *Client) ConversationsList(opts ...func(*ConversationsListParams)) (*ConversationsListResponse, error) {
	p := &ConversationsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations"

	params := url.Values{}
	if p.Folder != nil {
		params.Set("folder", fmt.Sprintf("%v", *p.Folder))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsListParams struct {
	Folder *string
	Page *int64
	Limit *int64
}

func (c *Client) ConversationsCreate(opts ...func(*ConversationsCreateParams)) (*ConversationsCreateResponse, error) {
	p := &ConversationsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations"


	bodyMap := map[string]any{}
	if p.RecipientId != nil {
		bodyMap["recipient_id"] = *p.RecipientId
	}
	if p.Recipients != nil {
		bodyMap["recipients"] = *p.Recipients
	}
	if p.IsGroup != nil {
		bodyMap["is_group"] = *p.IsGroup
	}
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.AllowEditMessages != nil {
		bodyMap["allow_edit_messages"] = *p.AllowEditMessages
	}
	if p.MessageBody != nil {
		bodyMap["message_body"] = *p.MessageBody
	}
	if p.OpenInvite != nil {
		bodyMap["open_invite"] = *p.OpenInvite
	}
	if p.AllowStickyMessages != nil {
		bodyMap["allow_sticky_messages"] = *p.AllowStickyMessages
	}
	if p.AllowDeleteOwnMessages != nil {
		bodyMap["allow_delete_own_messages"] = *p.AllowDeleteOwnMessages
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ConversationsCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsCreateParams struct {
	RecipientId *int64
	Recipients *[]string
	IsGroup *bool
	Title *string
	AllowEditMessages *bool
	MessageBody *string
	OpenInvite *bool
	AllowStickyMessages *bool
	AllowDeleteOwnMessages *bool
}

func (c *Client) ConversationsUpdate(conversationId int64, opts ...func(*ConversationsUpdateParams)) (map[string]any, error) {
	p := &ConversationsUpdateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations"


	bodyMap := map[string]any{}
	if p.AllowDeleteOwnMessages != nil {
		bodyMap["allow_delete_own_messages"] = *p.AllowDeleteOwnMessages
	}
	bodyMap["conversation_id"] = conversationId
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.OpenInvite != nil {
		bodyMap["open_invite"] = *p.OpenInvite
	}
	if p.HistoryOpen != nil {
		bodyMap["history_open"] = *p.HistoryOpen
	}
	if p.AllowEditMessages != nil {
		bodyMap["allow_edit_messages"] = *p.AllowEditMessages
	}
	if p.AllowStickyMessages != nil {
		bodyMap["allow_sticky_messages"] = *p.AllowStickyMessages
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsUpdateParams struct {
	AllowDeleteOwnMessages *bool
	Title *string
	OpenInvite *bool
	HistoryOpen *bool
	AllowEditMessages *bool
	AllowStickyMessages *bool
}

func (c *Client) ConversationsDelete(conversationId int64, deleteType string, opts ...func(*ConversationsDeleteParams)) (map[string]any, error) {
	p := &ConversationsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations"


	bodyMap := map[string]any{}
	bodyMap["conversation_id"] = conversationId
	bodyMap["delete_type"] = deleteType
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsDeleteParams struct {
}

func (c *Client) NotificationsGet(notificationId int64, opts ...func(*NotificationsGetParams)) (map[string]any, error) {
	p := &NotificationsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/notifications/{notification_id}/content"
	path = strings.Replace(path, "{notification_id}", fmt.Sprintf("%v", notificationId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type NotificationsGetParams struct {
}

func (c *Client) ChatboxGetLeaderboard(opts ...func(*ChatboxGetLeaderboardParams)) (*ChatboxGetLeaderboardResponse, error) {
	p := &ChatboxGetLeaderboardParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages/leaderboard"

	params := url.Values{}
	if p.Duration != nil {
		params.Set("duration", fmt.Sprintf("%v", *p.Duration))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ChatboxGetLeaderboardResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxGetLeaderboardParams struct {
	Duration *string
}

func (c *Client) ThreadsNavigation(threadId int64, opts ...func(*ThreadsNavigationParams)) (map[string]any, error) {
	p := &ThreadsNavigationParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/navigation"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsNavigationParams struct {
}

func (c *Client) PostsCommentsGet(postId int64, opts ...func(*PostsCommentsGetParams)) (map[string]any, error) {
	p := &PostsCommentsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments"

	params := url.Values{}
	params.Set("post_id", fmt.Sprintf("%v", postId))
	if p.Before != nil {
		params.Set("before", fmt.Sprintf("%v", *p.Before))
	}
	if p.BeforeComment != nil {
		params.Set("before_comment", fmt.Sprintf("%v", *p.BeforeComment))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsCommentsGetParams struct {
	Before *int64
	BeforeComment *int64
}

func (c *Client) PostsCommentsCreate(postId int64, commentBody string, opts ...func(*PostsCommentsCreateParams)) (map[string]any, error) {
	p := &PostsCommentsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments"


	bodyMap := map[string]any{}
	bodyMap["post_id"] = postId
	bodyMap["comment_body"] = commentBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsCommentsCreateParams struct {
}

func (c *Client) PostsCommentsEdit(postCommentId int64, commentBody string, opts ...func(*PostsCommentsEditParams)) (*PostsCommentsEditResponse, error) {
	p := &PostsCommentsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments"


	bodyMap := map[string]any{}
	bodyMap["post_comment_id"] = postCommentId
	bodyMap["comment_body"] = commentBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PostsCommentsEditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PostsCommentsEditParams struct {
}

func (c *Client) PostsCommentsDelete(postCommentId int64, opts ...func(*PostsCommentsDeleteParams)) (map[string]any, error) {
	p := &PostsCommentsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments"


	bodyMap := map[string]any{}
	bodyMap["post_comment_id"] = postCommentId
	if p.Reason != nil {
		bodyMap["reason"] = *p.Reason
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsCommentsDeleteParams struct {
	Reason *string
}

func (c *Client) ThreadsClaim(transferType string, postBody string, asIsMarketDeal bool, asAmount float64, asResponder string, opts ...func(*ThreadsClaimParams)) (*ThreadsClaimResponse, error) {
	p := &ThreadsClaimParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/claims"


	bodyMap := map[string]any{}
	if p.PayClaim != nil {
		bodyMap["pay_claim"] = *p.PayClaim
	}
	if p.AsFundsReceipt != nil {
		bodyMap["as_funds_receipt"] = *p.AsFundsReceipt
	}
	if p.Tags != nil {
		bodyMap["tags"] = *p.Tags
	}
	if p.ScheduleDate != nil {
		bodyMap["schedule_date"] = *p.ScheduleDate
	}
	if p.AsData != nil {
		bodyMap["as_data"] = *p.AsData
	}
	bodyMap["transfer_type"] = transferType
	if p.AsTgLoginScreenshot != nil {
		bodyMap["as_tg_login_screenshot"] = *p.AsTgLoginScreenshot
	}
	if p.HideContacts != nil {
		bodyMap["hide_contacts"] = *p.HideContacts
	}
	if p.AllowAskHiddenContent != nil {
		bodyMap["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.ReplyGroup != nil {
		bodyMap["reply_group"] = *p.ReplyGroup
	}
	if p.WatchThread != nil {
		bodyMap["watch_thread"] = *p.WatchThread
	}
	bodyMap["post_body"] = postBody
	bodyMap["as_is_market_deal"] = asIsMarketDeal
	if p.AsMarketItemId != nil {
		bodyMap["as_market_item_id"] = *p.AsMarketItemId
	}
	bodyMap["as_amount"] = asAmount
	if p.Currency != nil {
		bodyMap["currency"] = *p.Currency
	}
	if p.WatchThreadState != nil {
		bodyMap["watch_thread_state"] = *p.WatchThreadState
	}
	if p.WatchThreadEmail != nil {
		bodyMap["watch_thread_email"] = *p.WatchThreadEmail
	}
	bodyMap["as_responder"] = asResponder
	if p.CommentIgnoreGroup != nil {
		bodyMap["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.DontAlertFollowers != nil {
		bodyMap["dont_alert_followers"] = *p.DontAlertFollowers
	}
	if p.ScheduleTime != nil {
		bodyMap["schedule_time"] = *p.ScheduleTime
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ThreadsClaimResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsClaimParams struct {
	PayClaim *string
	AsFundsReceipt *string
	Tags *[]string
	ScheduleDate *string
	AsData *string
	AsTgLoginScreenshot *string
	HideContacts *bool
	AllowAskHiddenContent *bool
	ReplyGroup *int64
	WatchThread *bool
	AsMarketItemId *int64
	Currency *string
	WatchThreadState *bool
	WatchThreadEmail *bool
	CommentIgnoreGroup *bool
	DontAlertFollowers *bool
	ScheduleTime *string
}

func (c *Client) ThreadsPollGet(threadId int64, opts ...func(*ThreadsPollGetParams)) (map[string]any, error) {
	p := &ThreadsPollGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/poll"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsPollGetParams struct {
}

func (c *Client) UsersBackgroundUpload(userId string, background io.Reader, opts ...func(*UsersBackgroundUploadParams)) (map[string]any, error) {
	p := &UsersBackgroundUploadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/background"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	fields := map[string]string{}
	if p.Y != nil {
		fields["y"] = fmt.Sprintf("%v", *p.Y)
	}
	if p.Crop != nil {
		fields["crop"] = fmt.Sprintf("%v", *p.Crop)
	}
	if p.X != nil {
		fields["x"] = fmt.Sprintf("%v", *p.X)
	}
	files := map[string]io.Reader{}
	files["background"] = background

	raw, err := c.DoMultipart("POST", path, nil, fields, files)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersBackgroundUploadParams struct {
	Y *int64
	Crop *int64
	X *int64
}

func (c *Client) UsersBackgroundDelete(userId string, opts ...func(*UsersBackgroundDeleteParams)) (map[string]any, error) {
	p := &UsersBackgroundDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/background"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersBackgroundDeleteParams struct {
}

func (c *Client) UsersIgnored(opts ...func(*UsersIgnoredParams)) (*UsersIgnoredResponse, error) {
	p := &UsersIgnoredParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/ignored"

	params := url.Values{}
	if p.Total != nil {
		params.Set("total", fmt.Sprintf("%v", *p.Total))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result UsersIgnoredResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersIgnoredParams struct {
	Total *bool
}

func (c *Client) ConversationsMessagesList(conversationId int64, opts ...func(*ConversationsMessagesListParams)) (*ConversationsMessagesListResponse, error) {
	p := &ConversationsMessagesListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}
	if p.Before != nil {
		params.Set("before", fmt.Sprintf("%v", *p.Before))
	}
	if p.After != nil {
		params.Set("after", fmt.Sprintf("%v", *p.After))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsMessagesListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsMessagesListParams struct {
	Page *int64
	Limit *int64
	Order *string
	Before *int64
	After *int64
}

func (c *Client) ConversationsMessagesCreate(conversationId int64, messageBody string, opts ...func(*ConversationsMessagesCreateParams)) (*ConversationsMessagesCreateResponse, error) {
	p := &ConversationsMessagesCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	bodyMap := map[string]any{}
	if p.ReplyMessageId != nil {
		bodyMap["reply_message_id"] = *p.ReplyMessageId
	}
	bodyMap["message_body"] = messageBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ConversationsMessagesCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsMessagesCreateParams struct {
	ReplyMessageId *int64
}

func (c *Client) ConversationsRead(conversationId int64, opts ...func(*ConversationsReadParams)) (map[string]any, error) {
	p := &ConversationsReadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/read"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsReadParams struct {
}

func (c *Client) SearchPosts(opts ...func(*SearchPostsParams)) (map[string]any, error) {
	p := &SearchPostsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/posts"


	bodyMap := map[string]any{}
	if p.UserId != nil {
		bodyMap["user_id"] = *p.UserId
	}
	if p.Page != nil {
		bodyMap["page"] = *p.Page
	}
	if p.Limit != nil {
		bodyMap["limit"] = *p.Limit
	}
	if p.DataLimit != nil {
		bodyMap["data_limit"] = *p.DataLimit
	}
	if p.Q != nil {
		bodyMap["q"] = *p.Q
	}
	if p.Tag != nil {
		bodyMap["tag"] = *p.Tag
	}
	if p.ForumId != nil {
		bodyMap["forum_id"] = *p.ForumId
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type SearchPostsParams struct {
	UserId *string
	Page *int64
	Limit *int64
	DataLimit *int64
	Q *string
	Tag *string
	ForumId *int64
}

func (c *Client) SearchResults(searchId string, opts ...func(*SearchResultsParams)) (map[string]any, error) {
	p := &SearchResultsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/{search_id}/results"
	path = strings.Replace(path, "{search_id}", fmt.Sprintf("%v", searchId), 1)


	bodyMap := map[string]any{}
	if p.Page != nil {
		bodyMap["page"] = *p.Page
	}
	if p.Limit != nil {
		bodyMap["limit"] = *p.Limit
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("GET", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type SearchResultsParams struct {
	Page *int64
	Limit *int64
}

func (c *Client) PagesList(opts ...func(*PagesListParams)) (map[string]any, error) {
	p := &PagesListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/pages"

	params := url.Values{}
	if p.ParentPageId != nil {
		params.Set("parent_page_id", fmt.Sprintf("%v", *p.ParentPageId))
	}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PagesListParams struct {
	ParentPageId *int64
	Order *string
}

func (c *Client) ThreadsMove(threadId int64, nodeId string, opts ...func(*ThreadsMoveParams)) (map[string]any, error) {
	p := &ThreadsMoveParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/move"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	bodyMap := map[string]any{}
	bodyMap["node_id"] = nodeId
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.PrefixId != nil {
		bodyMap["prefix_id"] = *p.PrefixId
	}
	if p.ApplyThreadPrefix != nil {
		bodyMap["apply_thread_prefix"] = *p.ApplyThreadPrefix
	}
	if p.SendAlert != nil {
		bodyMap["send_alert"] = *p.SendAlert
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsMoveParams struct {
	Title *string
	TitleEn *string
	PrefixId *[]int64
	ApplyThreadPrefix *bool
	SendAlert *bool
}

func (c *Client) UsersSACancelReset(opts ...func(*UsersSACancelResetParams)) (map[string]any, error) {
	p := &UsersSACancelResetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/account/secret-answer/reset"


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersSACancelResetParams struct {
}

func (c *Client) UsersSAReset(opts ...func(*UsersSAResetParams)) (*UsersSAResetResponse, error) {
	p := &UsersSAResetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/account/secret-answer/reset"


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result UsersSAResetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersSAResetParams struct {
}

func (c *Client) ProfilePostsCreate(userId string, postBody string, opts ...func(*ProfilePostsCreateParams)) (map[string]any, error) {
	p := &ProfilePostsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts"


	bodyMap := map[string]any{}
	bodyMap["user_id"] = userId
	bodyMap["post_body"] = postBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsCreateParams struct {
}

func (c *Client) TagsList(opts ...func(*TagsListParams)) (map[string]any, error) {
	p := &TagsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/tags/list"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type TagsListParams struct {
	Page *int64
	Limit *int64
}

func (c *Client) ChatboxOnline(roomId int64, opts ...func(*ChatboxOnlineParams)) (*ChatboxOnlineResponse, error) {
	p := &ChatboxOnlineParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages/online"

	params := url.Values{}
	params.Set("room_id", fmt.Sprintf("%v", roomId))

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ChatboxOnlineResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxOnlineParams struct {
}

func (c *Client) FormsList(opts ...func(*FormsListParams)) (*FormsListResponse, error) {
	p := &FormsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forms"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result FormsListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type FormsListParams struct {
	Page *int64
}

func (c *Client) CategoriesList(opts ...func(*CategoriesListParams)) (map[string]any, error) {
	p := &CategoriesListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/categories"

	params := url.Values{}
	if p.ParentCategoryId != nil {
		params.Set("parent_category_id", fmt.Sprintf("%v", *p.ParentCategoryId))
	}
	if p.ParentForumId != nil {
		params.Set("parent_forum_id", fmt.Sprintf("%v", *p.ParentForumId))
	}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type CategoriesListParams struct {
	ParentCategoryId *int64
	ParentForumId *int64
	Order *string
}

func (c *Client) ForumsGrouped(opts ...func(*ForumsGroupedParams)) (*ForumsGroupedResponse, error) {
	p := &ForumsGroupedParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/grouped"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ForumsGroupedResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ForumsGroupedParams struct {
}

func (c *Client) UsersFields(opts ...func(*UsersFieldsParams)) (map[string]any, error) {
	p := &UsersFieldsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/fields"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersFieldsParams struct {
}

func (c *Client) ConversationsMessagesGet(messageId int64, opts ...func(*ConversationsMessagesGetParams)) (map[string]any, error) {
	p := &ConversationsMessagesGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/messages/{message_id}"
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%v", messageId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsMessagesGetParams struct {
}

func (c *Client) ConversationsReadAll(opts ...func(*ConversationsReadAllParams)) (*ConversationsReadAllResponse, error) {
	p := &ConversationsReadAllParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/read-all"


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsReadAllResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsReadAllParams struct {
}

func (c *Client) ForumsList(opts ...func(*ForumsListParams)) (*ForumsListResponse, error) {
	p := &ForumsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums"

	params := url.Values{}
	if p.ParentCategoryId != nil {
		params.Set("parent_category_id", fmt.Sprintf("%v", *p.ParentCategoryId))
	}
	if p.ParentForumId != nil {
		params.Set("parent_forum_id", fmt.Sprintf("%v", *p.ParentForumId))
	}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ForumsListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ForumsListParams struct {
	ParentCategoryId *int64
	ParentForumId *int64
	Order *string
}

func (c *Client) ThreadsBump(threadId int64, opts ...func(*ThreadsBumpParams)) (map[string]any, error) {
	p := &ThreadsBumpParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/bump"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsBumpParams struct {
}

func (c *Client) ThreadsFinish(threadId int64, opts ...func(*ThreadsFinishParams)) (map[string]any, error) {
	p := &ThreadsFinishParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/contests/{thread_id}/finish"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsFinishParams struct {
}

func (c *Client) UsersList(opts ...func(*UsersListParams)) (*UsersListResponse, error) {
	p := &UsersListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result UsersListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersListParams struct {
	Page *int64
	Limit *int64
	FieldsInclude *[]string
}

func (c *Client) UsersContents(userId string, opts ...func(*UsersContentsParams)) (map[string]any, error) {
	p := &UsersContentsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/timeline"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersContentsParams struct {
	Page *int64
	Limit *int64
}

func (c *Client) ProfilePostsDelete(profilePostId int64, opts ...func(*ProfilePostsDeleteParams)) (map[string]any, error) {
	p := &ProfilePostsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)

	params := url.Values{}
	if p.Reason != nil {
		params.Set("reason", fmt.Sprintf("%v", *p.Reason))
	}

	raw, err := c.DoRequest("DELETE", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsDeleteParams struct {
	Reason *string
}

func (c *Client) ProfilePostsGet(profilePostId int64, opts ...func(*ProfilePostsGetParams)) (map[string]any, error) {
	p := &ProfilePostsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsGetParams struct {
}

func (c *Client) ProfilePostsEdit(profilePostId int64, opts ...func(*ProfilePostsEditParams)) (map[string]any, error) {
	p := &ProfilePostsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	bodyMap := map[string]any{}
	if p.PostBody != nil {
		bodyMap["post_body"] = *p.PostBody
	}
	if p.DisableComments != nil {
		bodyMap["disable_comments"] = *p.DisableComments
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsEditParams struct {
	PostBody *string
	DisableComments *bool
}

func (c *Client) ProfilePostsCommentsGet(profilePostId int64, commentId int64, opts ...func(*ProfilePostsCommentsGetParams)) (map[string]any, error) {
	p := &ProfilePostsCommentsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/comments/{comment_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)
	path = strings.Replace(path, "{comment_id}", fmt.Sprintf("%v", commentId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsCommentsGetParams struct {
}

func (c *Client) UsersSecretAnswerTypes(opts ...func(*UsersSecretAnswerTypesParams)) (*UsersSecretAnswerTypesResponse, error) {
	p := &UsersSecretAnswerTypesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/secret-answer/types"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result UsersSecretAnswerTypesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersSecretAnswerTypesParams struct {
}

func (c *Client) AssetsCss(opts ...func(*AssetsCssParams)) (*AssetsCssResponse, error) {
	p := &AssetsCssParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/css"

	params := url.Values{}
	if p.Css != nil {
		params.Set("css", fmt.Sprintf("%v", *p.Css))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result AssetsCssResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type AssetsCssParams struct {
	Css *[]string
}

func (c *Client) ForumsFollowed(opts ...func(*ForumsFollowedParams)) (map[string]any, error) {
	p := &ForumsFollowedParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/followed"

	params := url.Values{}
	if p.Total != nil {
		params.Set("total", fmt.Sprintf("%v", *p.Total))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ForumsFollowedParams struct {
	Total *bool
}

func (c *Client) ThreadsStar(threadId int64, opts ...func(*ThreadsStarParams)) (map[string]any, error) {
	p := &ThreadsStarParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/star"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsStarParams struct {
}

func (c *Client) ThreadsUnstar(threadId int64, opts ...func(*ThreadsUnstarParams)) (map[string]any, error) {
	p := &ThreadsUnstarParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/star"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsUnstarParams struct {
}

func (c *Client) ThreadsFollowed(opts ...func(*ThreadsFollowedParams)) (map[string]any, error) {
	p := &ThreadsFollowedParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/followed"

	params := url.Values{}
	if p.Total != nil {
		params.Set("total", fmt.Sprintf("%v", *p.Total))
	}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsFollowedParams struct {
	Total *bool
	FieldsInclude *[]string
}

func (c *Client) UsersGet(userId string, opts ...func(*UsersGetParams)) (*UsersGetResponse, error) {
	p := &UsersGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result UsersGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersGetParams struct {
	FieldsInclude *[]string
}

func (c *Client) UsersEdit(userId string, opts ...func(*UsersEditParams)) (map[string]any, error) {
	p := &UsersEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	bodyMap := map[string]any{}
	if p.DisplayIconGroupId != nil {
		bodyMap["display_icon_group_id"] = *p.DisplayIconGroupId
	}
	if p.UserDobDay != nil {
		bodyMap["user_dob_day"] = *p.UserDobDay
	}
	if p.UserDobYear != nil {
		bodyMap["user_dob_year"] = *p.UserDobYear
	}
	if p.LanguageId != nil {
		bodyMap["language_id"] = *p.LanguageId
	}
	if p.ReceiveAdminEmail != nil {
		bodyMap["receive_admin_email"] = *p.ReceiveAdminEmail
	}
	if p.ShowDobYear != nil {
		bodyMap["show_dob_year"] = *p.ShowDobYear
	}
	if p.AllowViewProfile != nil {
		bodyMap["allow_view_profile"] = *p.AllowViewProfile
	}
	if p.AllowSendPersonalConversation != nil {
		bodyMap["allow_send_personal_conversation"] = *p.AllowSendPersonalConversation
	}
	if p.Gender != nil {
		bodyMap["gender"] = *p.Gender
	}
	if p.UserTitle != nil {
		bodyMap["user_title"] = *p.UserTitle
	}
	if p.ShowDobDate != nil {
		bodyMap["show_dob_date"] = *p.ShowDobDate
	}
	if p.AllowPostProfile != nil {
		bodyMap["allow_post_profile"] = *p.AllowPostProfile
	}
	if p.AllowInviteGroup != nil {
		bodyMap["allow_invite_group"] = *p.AllowInviteGroup
	}
	if p.AllowReceiveNewsFeed != nil {
		bodyMap["allow_receive_news_feed"] = *p.AllowReceiveNewsFeed
	}
	if p.Username != nil {
		bodyMap["username"] = *p.Username
	}
	if p.ActivityVisible != nil {
		bodyMap["activity_visible"] = *p.ActivityVisible
	}
	if p.HideUsernameChangeLogs != nil {
		bodyMap["hide_username_change_logs"] = *p.HideUsernameChangeLogs
	}
	if p.Alert != nil {
		bodyMap["alert"] = *p.Alert
	}
	if p.Fields != nil {
		bodyMap["fields"] = *p.Fields
	}
	if p.DisplayGroupId != nil {
		bodyMap["display_group_id"] = *p.DisplayGroupId
	}
	if p.DisplayBannerId != nil {
		bodyMap["display_banner_id"] = *p.DisplayBannerId
	}
	if p.ConvWelcomeMessage != nil {
		bodyMap["conv_welcome_message"] = *p.ConvWelcomeMessage
	}
	if p.UserDobMonth != nil {
		bodyMap["user_dob_month"] = *p.UserDobMonth
	}
	if p.SecretAnswer != nil {
		bodyMap["secret_answer"] = *p.SecretAnswer
	}
	if p.SecretAnswerType != nil {
		bodyMap["secret_answer_type"] = *p.SecretAnswerType
	}
	if p.ShortLink != nil {
		bodyMap["short_link"] = *p.ShortLink
	}
	if p.Timezone != nil {
		bodyMap["timezone"] = *p.Timezone
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersEditParams struct {
	DisplayIconGroupId *int64
	UserDobDay *int64
	UserDobYear *int64
	LanguageId *int64
	ReceiveAdminEmail *bool
	ShowDobYear *bool
	AllowViewProfile *string
	AllowSendPersonalConversation *string
	Gender *string
	UserTitle *string
	ShowDobDate *bool
	AllowPostProfile *string
	AllowInviteGroup *string
	AllowReceiveNewsFeed *string
	Username *string
	ActivityVisible *bool
	HideUsernameChangeLogs *bool
	Alert *map[string]any
	Fields *map[string]any
	DisplayGroupId *int64
	DisplayBannerId *int64
	ConvWelcomeMessage *string
	UserDobMonth *int64
	SecretAnswer *string
	SecretAnswerType *int64
	ShortLink *string
	Timezone *string
}

func (c *Client) UsersFollowings(userId string, opts ...func(*UsersFollowingsParams)) (map[string]any, error) {
	p := &UsersFollowingsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/followings"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersFollowingsParams struct {
	Order *string
	Page *int64
	Limit *int64
}

func (c *Client) ForumsGetFeedOptions(opts ...func(*ForumsGetFeedOptionsParams)) (*ForumsGetFeedOptionsResponse, error) {
	p := &ForumsGetFeedOptionsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/feed/options"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ForumsGetFeedOptionsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ForumsGetFeedOptionsParams struct {
}

func (c *Client) ForumsEditFeedOptions(opts ...func(*ForumsEditFeedOptionsParams)) (map[string]any, error) {
	p := &ForumsEditFeedOptionsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/feed/options"


	bodyMap := map[string]any{}
	if p.NodeIds != nil {
		bodyMap["node_ids"] = *p.NodeIds
	}
	if p.Keywords != nil {
		bodyMap["keywords"] = *p.Keywords
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ForumsEditFeedOptionsParams struct {
	NodeIds *[]int64
	Keywords *[]string
}

func (c *Client) ThreadsCreateContest(postBody string, contestType string, prizeType string, requireTotalLikeCount int64, requireLikeCount int64, opts ...func(*ThreadsCreateContestParams)) (*ThreadsCreateContestResponse, error) {
	p := &ThreadsCreateContestParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/contests"


	bodyMap := map[string]any{}
	if p.ScheduleTime != nil {
		bodyMap["schedule_time"] = *p.ScheduleTime
	}
	bodyMap["post_body"] = postBody
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	bodyMap["contest_type"] = contestType
	bodyMap["prize_type"] = prizeType
	if p.CountWinners != nil {
		bodyMap["count_winners"] = *p.CountWinners
	}
	if p.PrizeDataMoney != nil {
		bodyMap["prize_data_money"] = *p.PrizeDataMoney
	}
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.PrizeDataUpgrade != nil {
		bodyMap["prize_data_upgrade"] = *p.PrizeDataUpgrade
	}
	if p.HideContacts != nil {
		bodyMap["hide_contacts"] = *p.HideContacts
	}
	if p.WatchThread != nil {
		bodyMap["watch_thread"] = *p.WatchThread
	}
	if p.WatchThreadEmail != nil {
		bodyMap["watch_thread_email"] = *p.WatchThreadEmail
	}
	bodyMap["require_total_like_count"] = requireTotalLikeCount
	if p.LengthValue != nil {
		bodyMap["length_value"] = *p.LengthValue
	}
	if p.IsMoneyPlaces != nil {
		bodyMap["is_money_places"] = *p.IsMoneyPlaces
	}
	if p.Tags != nil {
		bodyMap["tags"] = *p.Tags
	}
	if p.CommentIgnoreGroup != nil {
		bodyMap["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.DontAlertFollowers != nil {
		bodyMap["dont_alert_followers"] = *p.DontAlertFollowers
	}
	if p.WatchThreadState != nil {
		bodyMap["watch_thread_state"] = *p.WatchThreadState
	}
	if p.LengthOption != nil {
		bodyMap["length_option"] = *p.LengthOption
	}
	if p.PrizeDataPlaces != nil {
		bodyMap["prize_data_places"] = *p.PrizeDataPlaces
	}
	bodyMap["require_like_count"] = requireLikeCount
	if p.SecretAnswer != nil {
		bodyMap["secret_answer"] = *p.SecretAnswer
	}
	if p.ReplyGroup != nil {
		bodyMap["reply_group"] = *p.ReplyGroup
	}
	if p.AllowAskHiddenContent != nil {
		bodyMap["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.ScheduleDate != nil {
		bodyMap["schedule_date"] = *p.ScheduleDate
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ThreadsCreateContestResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsCreateContestParams struct {
	ScheduleTime *string
	Title *string
	CountWinners *int64
	PrizeDataMoney *float64
	TitleEn *string
	PrizeDataUpgrade *int64
	HideContacts *bool
	WatchThread *bool
	WatchThreadEmail *bool
	LengthValue *int64
	IsMoneyPlaces *bool
	Tags *[]string
	CommentIgnoreGroup *bool
	DontAlertFollowers *bool
	WatchThreadState *bool
	LengthOption *string
	PrizeDataPlaces *[]float64
	SecretAnswer *string
	ReplyGroup *int64
	AllowAskHiddenContent *bool
	ScheduleDate *string
}

func (c *Client) ThreadsHide(threadId int64, opts ...func(*ThreadsHideParams)) (map[string]any, error) {
	p := &ThreadsHideParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/hide"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsHideParams struct {
}

func (c *Client) ThreadsPollVote(threadId int64, opts ...func(*ThreadsPollVoteParams)) (map[string]any, error) {
	p := &ThreadsPollVoteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/poll/votes"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	bodyMap := map[string]any{}
	if p.ResponseId != nil {
		bodyMap["response_id"] = *p.ResponseId
	}
	if p.ResponseIds != nil {
		bodyMap["response_ids"] = *p.ResponseIds
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsPollVoteParams struct {
	ResponseId *int64
	ResponseIds *[]int64
}

func (c *Client) TagsGet(tagId int64, opts ...func(*TagsGetParams)) (map[string]any, error) {
	p := &TagsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/tags/{tag_id}"
	path = strings.Replace(path, "{tag_id}", fmt.Sprintf("%v", tagId), 1)

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type TagsGetParams struct {
	Page *int64
	Limit *int64
}

func (c *Client) ForumsFollowers(forumId int64, opts ...func(*ForumsFollowersParams)) (map[string]any, error) {
	p := &ForumsFollowersParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/{forum_id}/followers"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%v", forumId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ForumsFollowersParams struct {
}

func (c *Client) ForumsFollow(forumId int64, opts ...func(*ForumsFollowParams)) (map[string]any, error) {
	p := &ForumsFollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/{forum_id}/followers"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%v", forumId), 1)


	bodyMap := map[string]any{}
	if p.Post != nil {
		bodyMap["post"] = *p.Post
	}
	if p.Alert != nil {
		bodyMap["alert"] = *p.Alert
	}
	if p.Email != nil {
		bodyMap["email"] = *p.Email
	}
	if p.PrefixIds != nil {
		bodyMap["prefix_ids"] = *p.PrefixIds
	}
	if p.MinimalContestAmount != nil {
		bodyMap["minimal_contest_amount"] = *p.MinimalContestAmount
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ForumsFollowParams struct {
	Post *bool
	Alert *bool
	Email *bool
	PrefixIds *[]int64
	MinimalContestAmount *int64
}

func (c *Client) ForumsUnfollow(forumId int64, opts ...func(*ForumsUnfollowParams)) (map[string]any, error) {
	p := &ForumsUnfollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/{forum_id}/followers"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%v", forumId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ForumsUnfollowParams struct {
}

func (c *Client) ThreadsUnfollow(threadId int64, opts ...func(*ThreadsUnfollowParams)) (map[string]any, error) {
	p := &ThreadsUnfollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsUnfollowParams struct {
}

func (c *Client) ThreadsFollowers(threadId int64, opts ...func(*ThreadsFollowersParams)) (map[string]any, error) {
	p := &ThreadsFollowersParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsFollowersParams struct {
}

func (c *Client) ThreadsFollow(threadId int64, opts ...func(*ThreadsFollowParams)) (map[string]any, error) {
	p := &ThreadsFollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	bodyMap := map[string]any{}
	if p.Email != nil {
		bodyMap["email"] = *p.Email
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsFollowParams struct {
	Email *bool
}

func (c *Client) ThreadsUnread(opts ...func(*ThreadsUnreadParams)) (map[string]any, error) {
	p := &ThreadsUnreadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/new"

	params := url.Values{}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.ForumId != nil {
		params.Set("forum_id", fmt.Sprintf("%v", *p.ForumId))
	}
	if p.DataLimit != nil {
		params.Set("data_limit", fmt.Sprintf("%v", *p.DataLimit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsUnreadParams struct {
	Limit *int64
	ForumId *int64
	DataLimit *int64
}

func (c *Client) UsersClaims(userId string, opts ...func(*UsersClaimsParams)) (*UsersClaimsResponse, error) {
	p := &UsersClaimsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/claims"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.TypeValue != nil {
		params.Set("type", fmt.Sprintf("%v", *p.TypeValue))
	}
	if p.ClaimState != nil {
		params.Set("claim_state", fmt.Sprintf("%v", *p.ClaimState))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result UsersClaimsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersClaimsParams struct {
	TypeValue *string
	ClaimState *string
}

func (c *Client) ProfilePostsReportReasons(profilePostId int64, opts ...func(*ProfilePostsReportReasonsParams)) (*ProfilePostsReportReasonsResponse, error) {
	p := &ProfilePostsReportReasonsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/report"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ProfilePostsReportReasonsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ProfilePostsReportReasonsParams struct {
}

func (c *Client) ProfilePostsReport(profilePostId int64, message string, opts ...func(*ProfilePostsReportParams)) (map[string]any, error) {
	p := &ProfilePostsReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/report"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	bodyMap := map[string]any{}
	bodyMap["message"] = message
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsReportParams struct {
}

func (c *Client) ConversationsKick(conversationId int64, userId int64, opts ...func(*ConversationsKickParams)) (map[string]any, error) {
	p := &ConversationsKickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/kick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	bodyMap := map[string]any{}
	bodyMap["user_id"] = userId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsKickParams struct {
}

func (c *Client) ConversationsStar(conversationId int64, opts ...func(*ConversationsStarParams)) (*ConversationsStarResponse, error) {
	p := &ConversationsStarParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/star"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsStarResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsStarParams struct {
}

func (c *Client) ConversationsUnstar(conversationId int64, opts ...func(*ConversationsUnstarParams)) (*ConversationsUnstarResponse, error) {
	p := &ConversationsUnstarParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/star"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsUnstarResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsUnstarParams struct {
}

func (c *Client) TagsFind(tag string, opts ...func(*TagsFindParams)) (map[string]any, error) {
	p := &TagsFindParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/tags/find"

	params := url.Values{}
	params.Set("tag", fmt.Sprintf("%v", tag))

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type TagsFindParams struct {
}

func (c *Client) CategoriesGet(categoryId int64, opts ...func(*CategoriesGetParams)) (map[string]any, error) {
	p := &CategoriesGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/categories/{category_id}"
	path = strings.Replace(path, "{category_id}", fmt.Sprintf("%v", categoryId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type CategoriesGetParams struct {
}

func (c *Client) ConversationsStart(userId string, opts ...func(*ConversationsStartParams)) (*ConversationsStartResponse, error) {
	p := &ConversationsStartParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/start"


	bodyMap := map[string]any{}
	bodyMap["user_id"] = userId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ConversationsStartResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsStartParams struct {
}

func (c *Client) SearchUsers(opts ...func(*SearchUsersParams)) (*SearchUsersResponse, error) {
	p := &SearchUsersParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/users"


	bodyMap := map[string]any{}
	if p.Q != nil {
		bodyMap["q"] = *p.Q
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result SearchUsersResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type SearchUsersParams struct {
	Q *string
}

func (c *Client) BatchExecute(opts ...func(*BatchExecuteParams)) (*BatchExecuteResponse, error) {
	p := &BatchExecuteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/batch"


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result BatchExecuteResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type BatchExecuteParams struct {
}

func (c *Client) ChatboxReport(messageId int64, reason string, opts ...func(*ChatboxReportParams)) (map[string]any, error) {
	p := &ChatboxReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages/report"


	bodyMap := map[string]any{}
	bodyMap["message_id"] = messageId
	bodyMap["reason"] = reason
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ChatboxReportParams struct {
}

func (c *Client) ChatboxReportReasons(messageId int64, opts ...func(*ChatboxReportReasonsParams)) (*ChatboxReportReasonsResponse, error) {
	p := &ChatboxReportReasonsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages/report"

	params := url.Values{}
	params.Set("message_id", fmt.Sprintf("%v", messageId))

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ChatboxReportReasonsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxReportReasonsParams struct {
}

func (c *Client) OAuthToken(opts ...func(*OAuthTokenParams)) (*OAuthTokenResponse, error) {
	p := &OAuthTokenParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/oauth/token"


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result OAuthTokenResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type OAuthTokenParams struct {
}

func (c *Client) ThreadsGet(threadId int64, opts ...func(*ThreadsGetParams)) (*ThreadsGetResponse, error) {
	p := &ThreadsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)

	params := url.Values{}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ThreadsGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsGetParams struct {
	FieldsInclude *[]string
}

func (c *Client) ThreadsEdit(threadId int64, opts ...func(*ThreadsEditParams)) (*ThreadsEditResponse, error) {
	p := &ThreadsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	bodyMap := map[string]any{}
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.PrefixId != nil {
		bodyMap["prefix_id"] = *p.PrefixId
	}
	if p.Tags != nil {
		bodyMap["tags"] = *p.Tags
	}
	if p.AllowAskHiddenContent != nil {
		bodyMap["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.ReplyGroup != nil {
		bodyMap["reply_group"] = *p.ReplyGroup
	}
	if p.CommentIgnoreGroup != nil {
		bodyMap["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.DiscussionOpen != nil {
		bodyMap["discussion_open"] = *p.DiscussionOpen
	}
	if p.HideContacts != nil {
		bodyMap["hide_contacts"] = *p.HideContacts
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ThreadsEditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsEditParams struct {
	Title *string
	PrefixId *[]int64
	Tags *[]string
	AllowAskHiddenContent *bool
	ReplyGroup *int64
	CommentIgnoreGroup *bool
	TitleEn *string
	DiscussionOpen *bool
	HideContacts *bool
}

func (c *Client) ThreadsDelete(threadId int64, opts ...func(*ThreadsDeleteParams)) (map[string]any, error) {
	p := &ThreadsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	bodyMap := map[string]any{}
	if p.Reason != nil {
		bodyMap["reason"] = *p.Reason
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsDeleteParams struct {
	Reason *string
}

func (c *Client) ProfilePostsList(userId string, opts ...func(*ProfilePostsListParams)) (map[string]any, error) {
	p := &ProfilePostsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/profile-posts"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.PostsUserId != nil {
		params.Set("posts_user_id", fmt.Sprintf("%v", *p.PostsUserId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsListParams struct {
	PostsUserId *int64
	Page *int64
	Limit *int64
	FieldsInclude *[]string
}

func (c *Client) ConversationsMessagesEdit(conversationId int64, messageId int64, messageBody string, opts ...func(*ConversationsMessagesEditParams)) (*ConversationsMessagesEditResponse, error) {
	p := &ConversationsMessagesEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages/{message_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%v", messageId), 1)


	bodyMap := map[string]any{}
	bodyMap["message_body"] = messageBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ConversationsMessagesEditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsMessagesEditParams struct {
}

func (c *Client) ConversationsMessagesDelete(conversationId int64, messageId int64, opts ...func(*ConversationsMessagesDeleteParams)) (map[string]any, error) {
	p := &ConversationsMessagesDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages/{message_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%v", messageId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsMessagesDeleteParams struct {
}

func (c *Client) ConversationsMessagesUnstick(conversationId int64, messageId int64, opts ...func(*ConversationsMessagesUnstickParams)) (map[string]any, error) {
	p := &ConversationsMessagesUnstickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages/{message_id}/stick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%v", messageId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsMessagesUnstickParams struct {
}

func (c *Client) ConversationsMessagesStick(conversationId int64, messageId int64, opts ...func(*ConversationsMessagesStickParams)) (map[string]any, error) {
	p := &ConversationsMessagesStickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages/{message_id}/stick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%v", messageId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsMessagesStickParams struct {
}

func (c *Client) TagsPopular(opts ...func(*TagsPopularParams)) (map[string]any, error) {
	p := &TagsPopularParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/tags"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type TagsPopularParams struct {
}

func (c *Client) FormsCreate(opts ...func(*FormsCreateParams)) (*FormsCreateResponse, error) {
	p := &FormsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forms/save"


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result FormsCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type FormsCreateParams struct {
}

func (c *Client) UsersAvatarCrop(userId string, opts ...func(*UsersAvatarCropParams)) (map[string]any, error) {
	p := &UsersAvatarCropParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/avatar/crop"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	bodyMap := map[string]any{}
	if p.X != nil {
		bodyMap["x"] = *p.X
	}
	if p.Y != nil {
		bodyMap["y"] = *p.Y
	}
	if p.Crop != nil {
		bodyMap["crop"] = *p.Crop
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersAvatarCropParams struct {
	X *int64
	Y *int64
	Crop *int64
}

func (c *Client) ProfilePostsStick(profilePostId int64, opts ...func(*ProfilePostsStickParams)) (map[string]any, error) {
	p := &ProfilePostsStickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/stick"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsStickParams struct {
}

func (c *Client) ProfilePostsUnstick(profilePostId int64, opts ...func(*ProfilePostsUnstickParams)) (map[string]any, error) {
	p := &ProfilePostsUnstickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/stick"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsUnstickParams struct {
}

func (c *Client) ProfilePostsUnlike(profilePostId int64, opts ...func(*ProfilePostsUnlikeParams)) (map[string]any, error) {
	p := &ProfilePostsUnlikeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/likes"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsUnlikeParams struct {
}

func (c *Client) ProfilePostsLikes(profilePostId int64, opts ...func(*ProfilePostsLikesParams)) (map[string]any, error) {
	p := &ProfilePostsLikesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/likes"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsLikesParams struct {
}

func (c *Client) ProfilePostsLike(profilePostId int64, opts ...func(*ProfilePostsLikeParams)) (map[string]any, error) {
	p := &ProfilePostsLikeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}/likes"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsLikeParams struct {
}

func (c *Client) ConversationsInvite(conversationId int64, recipients []string, opts ...func(*ConversationsInviteParams)) (map[string]any, error) {
	p := &ConversationsInviteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/invite"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	bodyMap := map[string]any{}
	bodyMap["recipients"] = recipients
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsInviteParams struct {
}

func (c *Client) NotificationsRead(opts ...func(*NotificationsReadParams)) (map[string]any, error) {
	p := &NotificationsReadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/notifications/read"


	bodyMap := map[string]any{}
	if p.NotificationId != nil {
		bodyMap["notification_id"] = *p.NotificationId
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type NotificationsReadParams struct {
	NotificationId *int64
}

func (c *Client) SearchProfilePosts(opts ...func(*SearchProfilePostsParams)) (map[string]any, error) {
	p := &SearchProfilePostsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/profile-posts"


	bodyMap := map[string]any{}
	if p.UserId != nil {
		bodyMap["user_id"] = *p.UserId
	}
	if p.Page != nil {
		bodyMap["page"] = *p.Page
	}
	if p.Limit != nil {
		bodyMap["limit"] = *p.Limit
	}
	if p.Q != nil {
		bodyMap["q"] = *p.Q
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type SearchProfilePostsParams struct {
	UserId *int64
	Page *int64
	Limit *int64
	Q *string
}

func (c *Client) ChatboxIndex(opts ...func(*ChatboxIndexParams)) (*ChatboxIndexResponse, error) {
	p := &ChatboxIndexParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox"

	params := url.Values{}
	if p.RoomId != nil {
		params.Set("room_id", fmt.Sprintf("%v", *p.RoomId))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ChatboxIndexResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxIndexParams struct {
	RoomId *int64
}

func (c *Client) ChatboxPostMessage(roomId int64, message string, opts ...func(*ChatboxPostMessageParams)) (*ChatboxPostMessageResponse, error) {
	p := &ChatboxPostMessageParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages"


	bodyMap := map[string]any{}
	bodyMap["room_id"] = roomId
	if p.ReplyMessageId != nil {
		bodyMap["reply_message_id"] = *p.ReplyMessageId
	}
	bodyMap["message"] = message
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ChatboxPostMessageResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxPostMessageParams struct {
	ReplyMessageId *int64
}

func (c *Client) ChatboxEditMessage(messageId int64, message string, opts ...func(*ChatboxEditMessageParams)) (*ChatboxEditMessageResponse, error) {
	p := &ChatboxEditMessageParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages"


	bodyMap := map[string]any{}
	bodyMap["message_id"] = messageId
	bodyMap["message"] = message
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ChatboxEditMessageResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxEditMessageParams struct {
}

func (c *Client) ChatboxDeleteMessage(messageId int64, opts ...func(*ChatboxDeleteMessageParams)) (map[string]any, error) {
	p := &ChatboxDeleteMessageParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages"


	bodyMap := map[string]any{}
	bodyMap["message_id"] = messageId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ChatboxDeleteMessageParams struct {
}

func (c *Client) ChatboxGetMessages(roomId int64, opts ...func(*ChatboxGetMessagesParams)) (*ChatboxGetMessagesResponse, error) {
	p := &ChatboxGetMessagesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages"

	params := url.Values{}
	params.Set("room_id", fmt.Sprintf("%v", roomId))
	if p.BeforeMessageId != nil {
		params.Set("before_message_id", fmt.Sprintf("%v", *p.BeforeMessageId))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ChatboxGetMessagesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxGetMessagesParams struct {
	BeforeMessageId *int64
}

func (c *Client) NavigationList(opts ...func(*NavigationListParams)) (map[string]any, error) {
	p := &NavigationListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/navigation"

	params := url.Values{}
	if p.Parent != nil {
		params.Set("parent", fmt.Sprintf("%v", *p.Parent))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type NavigationListParams struct {
	Parent *int64
}

func (c *Client) ThreadsList(opts ...func(*ThreadsListParams)) (map[string]any, error) {
	p := &ThreadsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads"

	params := url.Values{}
	if p.ForumId != nil {
		params.Set("forum_id", fmt.Sprintf("%v", *p.ForumId))
	}
	if p.Tab != nil {
		params.Set("tab", fmt.Sprintf("%v", *p.Tab))
	}
	if p.State != nil {
		params.Set("state", fmt.Sprintf("%v", *p.State))
	}
	if p.Period != nil {
		params.Set("period", fmt.Sprintf("%v", *p.Period))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.TitleOnly != nil {
		params.Set("title_only", fmt.Sprintf("%v", *p.TitleOnly))
	}
	if p.CreatorUserId != nil {
		params.Set("creator_user_id", fmt.Sprintf("%v", *p.CreatorUserId))
	}
	if p.Sticky != nil {
		params.Set("sticky", fmt.Sprintf("%v", *p.Sticky))
	}
	if p.PrefixIds != nil {
		params.Set("prefix_ids[]", fmt.Sprintf("%v", *p.PrefixIds))
	}
	if p.PrefixIdsNot != nil {
		params.Set("prefix_ids_not[]", fmt.Sprintf("%v", *p.PrefixIdsNot))
	}
	if p.ThreadTagId != nil {
		params.Set("thread_tag_id", fmt.Sprintf("%v", *p.ThreadTagId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}
	if p.Direction != nil {
		params.Set("direction", fmt.Sprintf("%v", *p.Direction))
	}
	if p.ThreadCreateDate != nil {
		params.Set("thread_create_date", fmt.Sprintf("%v", *p.ThreadCreateDate))
	}
	if p.ThreadUpdateDate != nil {
		params.Set("thread_update_date", fmt.Sprintf("%v", *p.ThreadUpdateDate))
	}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsListParams struct {
	ForumId *int64
	Tab *string
	State *string
	Period *string
	Title *string
	TitleOnly *bool
	CreatorUserId *int64
	Sticky *bool
	PrefixIds *[]int64
	PrefixIdsNot *[]int64
	ThreadTagId *int64
	Page *int64
	Limit *int64
	Order *string
	Direction *string
	ThreadCreateDate *int64
	ThreadUpdateDate *int64
	FieldsInclude *[]string
}

func (c *Client) ThreadsCreate(forumId int64, postBody string, opts ...func(*ThreadsCreateParams)) (*ThreadsCreateResponse, error) {
	p := &ThreadsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads"


	bodyMap := map[string]any{}
	if p.CommentIgnoreGroup != nil {
		bodyMap["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.WatchThreadState != nil {
		bodyMap["watch_thread_state"] = *p.WatchThreadState
	}
	bodyMap["forum_id"] = forumId
	if p.Tags != nil {
		bodyMap["tags"] = *p.Tags
	}
	if p.HideContacts != nil {
		bodyMap["hide_contacts"] = *p.HideContacts
	}
	if p.ReplyGroup != nil {
		bodyMap["reply_group"] = *p.ReplyGroup
	}
	if p.ScheduleDate != nil {
		bodyMap["schedule_date"] = *p.ScheduleDate
	}
	if p.WatchThread != nil {
		bodyMap["watch_thread"] = *p.WatchThread
	}
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.AllowAskHiddenContent != nil {
		bodyMap["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.DontAlertFollowers != nil {
		bodyMap["dont_alert_followers"] = *p.DontAlertFollowers
	}
	if p.ScheduleTime != nil {
		bodyMap["schedule_time"] = *p.ScheduleTime
	}
	bodyMap["post_body"] = postBody
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.PrefixId != nil {
		bodyMap["prefix_id"] = *p.PrefixId
	}
	if p.WatchThreadEmail != nil {
		bodyMap["watch_thread_email"] = *p.WatchThreadEmail
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ThreadsCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsCreateParams struct {
	CommentIgnoreGroup *bool
	WatchThreadState *bool
	Tags *[]string
	HideContacts *bool
	ReplyGroup *int64
	ScheduleDate *string
	WatchThread *bool
	Title *string
	AllowAskHiddenContent *bool
	DontAlertFollowers *bool
	ScheduleTime *string
	TitleEn *string
	PrefixId *[]int64
	WatchThreadEmail *bool
}

func (c *Client) ThreadsRecent(opts ...func(*ThreadsRecentParams)) (map[string]any, error) {
	p := &ThreadsRecentParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/recent"

	params := url.Values{}
	if p.Days != nil {
		params.Set("days", fmt.Sprintf("%v", *p.Days))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.ForumId != nil {
		params.Set("forum_id", fmt.Sprintf("%v", *p.ForumId))
	}
	if p.DataLimit != nil {
		params.Set("data_limit", fmt.Sprintf("%v", *p.DataLimit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ThreadsRecentParams struct {
	Days *int64
	Limit *int64
	ForumId *int64
	DataLimit *int64
}

func (c *Client) PostsLikes(postId int64, opts ...func(*PostsLikesParams)) (map[string]any, error) {
	p := &PostsLikesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}/likes"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsLikesParams struct {
	Page *int64
	Limit *int64
}

func (c *Client) PostsLike(postId int64, opts ...func(*PostsLikeParams)) (map[string]any, error) {
	p := &PostsLikeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}/likes"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsLikeParams struct {
}

func (c *Client) PostsUnlike(postId int64, opts ...func(*PostsUnlikeParams)) (map[string]any, error) {
	p := &PostsUnlikeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}/likes"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsUnlikeParams struct {
}

func (c *Client) UsersBackgroundCrop(userId string, opts ...func(*UsersBackgroundCropParams)) (map[string]any, error) {
	p := &UsersBackgroundCropParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/background/crop"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	bodyMap := map[string]any{}
	if p.Y != nil {
		bodyMap["y"] = *p.Y
	}
	if p.Crop != nil {
		bodyMap["crop"] = *p.Crop
	}
	if p.X != nil {
		bodyMap["x"] = *p.X
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersBackgroundCropParams struct {
	Y *int64
	Crop *int64
	X *int64
}

func (c *Client) ChatboxGetIgnore(opts ...func(*ChatboxGetIgnoreParams)) (*ChatboxGetIgnoreResponse, error) {
	p := &ChatboxGetIgnoreParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/ignore"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ChatboxGetIgnoreResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxGetIgnoreParams struct {
}

func (c *Client) ChatboxPostIgnore(userId string, opts ...func(*ChatboxPostIgnoreParams)) (map[string]any, error) {
	p := &ChatboxPostIgnoreParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/ignore"


	bodyMap := map[string]any{}
	bodyMap["user_id"] = userId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ChatboxPostIgnoreParams struct {
}

func (c *Client) ChatboxDeleteIgnore(userId string, opts ...func(*ChatboxDeleteIgnoreParams)) (map[string]any, error) {
	p := &ChatboxDeleteIgnoreParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/ignore"


	bodyMap := map[string]any{}
	bodyMap["user_id"] = userId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ChatboxDeleteIgnoreParams struct {
}

func (c *Client) UsersFind(opts ...func(*UsersFindParams)) (map[string]any, error) {
	p := &UsersFindParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/find"

	params := url.Values{}
	if p.Username != nil {
		params.Set("username", fmt.Sprintf("%v", *p.Username))
	}
	if p.CustomFields != nil {
		params.Set("custom_fields", fmt.Sprintf("%v", *p.CustomFields))
	}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersFindParams struct {
	Username *string
	CustomFields *map[string]any
	FieldsInclude *[]string
}

func (c *Client) ProfilePostsCommentsReport(commentId int64, message string, opts ...func(*ProfilePostsCommentsReportParams)) (map[string]any, error) {
	p := &ProfilePostsCommentsReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments/{comment_id}/report"
	path = strings.Replace(path, "{comment_id}", fmt.Sprintf("%v", commentId), 1)


	bodyMap := map[string]any{}
	bodyMap["message"] = message
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsCommentsReportParams struct {
}

func (c *Client) ConversationsSave(link string, opts ...func(*ConversationsSaveParams)) (map[string]any, error) {
	p := &ConversationsSaveParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/save"


	bodyMap := map[string]any{}
	bodyMap["link"] = link
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ConversationsSaveParams struct {
}

func (c *Client) ConversationsSearch(opts ...func(*ConversationsSearchParams)) (*ConversationsSearchResponse, error) {
	p := &ConversationsSearchParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/search"


	bodyMap := map[string]any{}
	if p.Q != nil {
		bodyMap["q"] = *p.Q
	}
	if p.ConversationId != nil {
		bodyMap["conversation_id"] = *p.ConversationId
	}
	if p.SearchRecipients != nil {
		bodyMap["search_recipients"] = *p.SearchRecipients
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ConversationsSearchResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsSearchParams struct {
	Q *string
	ConversationId *int64
	SearchRecipients *bool
}

func (c *Client) ForumsGet(forumId int64, opts ...func(*ForumsGetParams)) (map[string]any, error) {
	p := &ForumsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/forums/{forum_id}"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%v", forumId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ForumsGetParams struct {
}

func (c *Client) PagesGet(pageId int64, opts ...func(*PagesGetParams)) (map[string]any, error) {
	p := &PagesGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/pages/{page_id}"
	path = strings.Replace(path, "{page_id}", fmt.Sprintf("%v", pageId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PagesGetParams struct {
}

func (c *Client) PostsList(opts ...func(*PostsListParams)) (map[string]any, error) {
	p := &PostsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts"

	params := url.Values{}
	if p.ThreadId != nil {
		params.Set("thread_id", fmt.Sprintf("%v", *p.ThreadId))
	}
	if p.PageOfPostId != nil {
		params.Set("page_of_post_id", fmt.Sprintf("%v", *p.PageOfPostId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.Order != nil {
		params.Set("order", fmt.Sprintf("%v", *p.Order))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsListParams struct {
	ThreadId *int64
	PageOfPostId *int64
	Page *int64
	Limit *int64
	Order *string
}

func (c *Client) PostsCreate(postBody string, opts ...func(*PostsCreateParams)) (*PostsCreateResponse, error) {
	p := &PostsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts"


	bodyMap := map[string]any{}
	bodyMap["post_body"] = postBody
	if p.ThreadId != nil {
		bodyMap["thread_id"] = *p.ThreadId
	}
	if p.QuotePostId != nil {
		bodyMap["quote_post_id"] = *p.QuotePostId
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PostsCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PostsCreateParams struct {
	ThreadId *int64
	QuotePostId *int64
}

func (c *Client) ProfilePostsCommentsList(profilePostId int64, opts ...func(*ProfilePostsCommentsListParams)) (map[string]any, error) {
	p := &ProfilePostsCommentsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments"

	params := url.Values{}
	params.Set("profile_post_id", fmt.Sprintf("%v", profilePostId))
	if p.Before != nil {
		params.Set("before", fmt.Sprintf("%v", *p.Before))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsCommentsListParams struct {
	Before *int64
	Limit *int64
}

func (c *Client) ProfilePostsCommentsCreate(profilePostId int64, commentBody string, opts ...func(*ProfilePostsCommentsCreateParams)) (map[string]any, error) {
	p := &ProfilePostsCommentsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments"


	bodyMap := map[string]any{}
	bodyMap["profile_post_id"] = profilePostId
	bodyMap["comment_body"] = commentBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsCommentsCreateParams struct {
}

func (c *Client) ProfilePostsCommentsEdit(commentId int64, commentBody string, opts ...func(*ProfilePostsCommentsEditParams)) (*ProfilePostsCommentsEditResponse, error) {
	p := &ProfilePostsCommentsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments"


	bodyMap := map[string]any{}
	bodyMap["comment_id"] = commentId
	bodyMap["comment_body"] = commentBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ProfilePostsCommentsEditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ProfilePostsCommentsEditParams struct {
}

func (c *Client) ProfilePostsCommentsDelete(commentId int64, opts ...func(*ProfilePostsCommentsDeleteParams)) (map[string]any, error) {
	p := &ProfilePostsCommentsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments"


	bodyMap := map[string]any{}
	bodyMap["comment_id"] = commentId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ProfilePostsCommentsDeleteParams struct {
}

func (c *Client) ConversationsAlertsEnable(conversationId int64, opts ...func(*ConversationsAlertsEnableParams)) (*ConversationsAlertsEnableResponse, error) {
	p := &ConversationsAlertsEnableParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/alerts"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsAlertsEnableResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsAlertsEnableParams struct {
}

func (c *Client) ConversationsAlertsDisable(conversationId int64, opts ...func(*ConversationsAlertsDisableParams)) (*ConversationsAlertsDisableResponse, error) {
	p := &ConversationsAlertsDisableParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/alerts"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConversationsAlertsDisableResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsAlertsDisableParams struct {
}

func (c *Client) NotificationsList(opts ...func(*NotificationsListParams)) (map[string]any, error) {
	p := &NotificationsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/notifications"

	params := url.Values{}
	if p.TypeValue != nil {
		params.Set("type", fmt.Sprintf("%v", *p.TypeValue))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type NotificationsListParams struct {
	TypeValue *string
	Page *int64
	Limit *int64
}

func (c *Client) SearchAll(opts ...func(*SearchAllParams)) (*SearchAllResponse, error) {
	p := &SearchAllParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search"


	bodyMap := map[string]any{}
	if p.Page != nil {
		bodyMap["page"] = *p.Page
	}
	if p.Limit != nil {
		bodyMap["limit"] = *p.Limit
	}
	if p.Q != nil {
		bodyMap["q"] = *p.Q
	}
	if p.Tag != nil {
		bodyMap["tag"] = *p.Tag
	}
	if p.ForumId != nil {
		bodyMap["forum_id"] = *p.ForumId
	}
	if p.UserId != nil {
		bodyMap["user_id"] = *p.UserId
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result SearchAllResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type SearchAllParams struct {
	Page *int64
	Limit *int64
	Q *string
	Tag *string
	ForumId *int64
	UserId *string
}

func (c *Client) SearchThreads(opts ...func(*SearchThreadsParams)) (map[string]any, error) {
	p := &SearchThreadsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/threads"


	bodyMap := map[string]any{}
	if p.ForumId != nil {
		bodyMap["forum_id"] = *p.ForumId
	}
	if p.UserId != nil {
		bodyMap["user_id"] = *p.UserId
	}
	if p.Page != nil {
		bodyMap["page"] = *p.Page
	}
	if p.Limit != nil {
		bodyMap["limit"] = *p.Limit
	}
	if p.DataLimit != nil {
		bodyMap["data_limit"] = *p.DataLimit
	}
	if p.Q != nil {
		bodyMap["q"] = *p.Q
	}
	if p.Tag != nil {
		bodyMap["tag"] = *p.Tag
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type SearchThreadsParams struct {
	ForumId *int64
	UserId *string
	Page *int64
	Limit *int64
	DataLimit *int64
	Q *string
	Tag *string
}

func (c *Client) UsersIgnore(userId string, opts ...func(*UsersIgnoreParams)) (map[string]any, error) {
	p := &UsersIgnoreParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/ignore"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersIgnoreParams struct {
}

func (c *Client) UsersIgnoreEdit(userId string, opts ...func(*UsersIgnoreEditParams)) (map[string]any, error) {
	p := &UsersIgnoreEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/ignore"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)

	params := url.Values{}
	if p.IgnoreConversations != nil {
		params.Set("ignore_conversations", fmt.Sprintf("%v", *p.IgnoreConversations))
	}
	if p.IgnoreContent != nil {
		params.Set("ignore_content", fmt.Sprintf("%v", *p.IgnoreContent))
	}
	if p.RestrictViewProfile != nil {
		params.Set("restrict_view_profile", fmt.Sprintf("%v", *p.RestrictViewProfile))
	}

	raw, err := c.DoRequest("PUT", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersIgnoreEditParams struct {
	IgnoreConversations *bool
	IgnoreContent *bool
	RestrictViewProfile *bool
}

func (c *Client) UsersUnignore(userId string, opts ...func(*UsersUnignoreParams)) (map[string]any, error) {
	p := &UsersUnignoreParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/ignore"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersUnignoreParams struct {
}

func (c *Client) UsersTrophies(userId string, opts ...func(*UsersTrophiesParams)) (*UsersTrophiesResponse, error) {
	p := &UsersTrophiesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/trophies"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result UsersTrophiesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersTrophiesParams struct {
}

func (c *Client) PostsGet(postId int64, opts ...func(*PostsGetParams)) (*PostsGetResponse, error) {
	p := &PostsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result PostsGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PostsGetParams struct {
}

func (c *Client) PostsEdit(postId int64, opts ...func(*PostsEditParams)) (*PostsEditResponse, error) {
	p := &PostsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	bodyMap := map[string]any{}
	if p.PostBody != nil {
		bodyMap["post_body"] = *p.PostBody
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PostsEditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PostsEditParams struct {
	PostBody *string
}

func (c *Client) PostsDelete(postId int64, opts ...func(*PostsDeleteParams)) (map[string]any, error) {
	p := &PostsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	bodyMap := map[string]any{}
	if p.Reason != nil {
		bodyMap["reason"] = *p.Reason
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsDeleteParams struct {
	Reason *string
}

func (c *Client) PostsReportReasons(postId int64, opts ...func(*PostsReportReasonsParams)) (*PostsReportReasonsResponse, error) {
	p := &PostsReportReasonsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}/report"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result PostsReportReasonsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PostsReportReasonsParams struct {
}

func (c *Client) PostsReport(postId int64, message string, opts ...func(*PostsReportParams)) (map[string]any, error) {
	p := &PostsReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/{post_id}/report"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%v", postId), 1)


	bodyMap := map[string]any{}
	bodyMap["message"] = message
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type PostsReportParams struct {
}

func (c *Client) UsersAvatarUpload(userId string, avatar io.Reader, opts ...func(*UsersAvatarUploadParams)) (map[string]any, error) {
	p := &UsersAvatarUploadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/avatar"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	fields := map[string]string{}
	if p.X != nil {
		fields["x"] = fmt.Sprintf("%v", *p.X)
	}
	if p.Y != nil {
		fields["y"] = fmt.Sprintf("%v", *p.Y)
	}
	if p.Crop != nil {
		fields["crop"] = fmt.Sprintf("%v", *p.Crop)
	}
	files := map[string]io.Reader{}
	files["avatar"] = avatar

	raw, err := c.DoMultipart("POST", path, nil, fields, files)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersAvatarUploadParams struct {
	X *int64
	Y *int64
	Crop *int64
}

func (c *Client) UsersAvatarDelete(userId string, opts ...func(*UsersAvatarDeleteParams)) (map[string]any, error) {
	p := &UsersAvatarDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/avatar"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	raw, err := c.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type UsersAvatarDeleteParams struct {
}
