package forum

import (
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

	return raw, nil
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

	return raw, nil
}

type PostsUnlikeParams struct {
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

	return raw, nil
}

type PostsLikesParams struct {
	Page *int64
	Limit *int64
}

func (c *Client) PostsCommentsReport(postCommentId int64, message string, opts ...func(*PostsCommentsReportParams)) (map[string]any, error) {
	p := &PostsCommentsReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments/report"


	body := map[string]any{}
	body["post_comment_id"] = postCommentId
	body["message"] = message

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type PostsCommentsReportParams struct {
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

	return raw, nil
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

	return raw, nil
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

	return raw, nil
}

type UsersUnignoreParams struct {
}

func (c *Client) ProfilePostsCommentsReport(commentId int64, message string, opts ...func(*ProfilePostsCommentsReportParams)) (map[string]any, error) {
	p := &ProfilePostsCommentsReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments/{comment_id}/report"
	path = strings.Replace(path, "{comment_id}", fmt.Sprintf("%v", commentId), 1)


	body := map[string]any{}
	body["message"] = message

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ProfilePostsCommentsReportParams struct {
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

	return raw, nil
}

type ThreadsFollowedParams struct {
	Total *bool
	FieldsInclude *[]string
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

	return raw, nil
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

	return raw, nil
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

	return raw, nil
}

type UsersUnfollowParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersSecretAnswerTypesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersSecretAnswerTypesParams struct {
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

	return raw, nil
}

type ConversationsMessagesGetParams struct {
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

	return raw, nil
}

type TagsFindParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxOnlineResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxOnlineParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxReportReasonsResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxReportReasonsParams struct {
}

func (c *Client) ChatboxReport(messageId int64, reason string, opts ...func(*ChatboxReportParams)) (map[string]any, error) {
	p := &ChatboxReportParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages/report"


	body := map[string]any{}
	body["message_id"] = messageId
	body["reason"] = reason

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ChatboxReportParams struct {
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

	return raw, nil
}

type TagsGetParams struct {
	Page *int64
	Limit *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result FormsListResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type FormsListParams struct {
	Page *int64
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

	return raw, nil
}

type UsersFollowingsParams struct {
	Order *string
	Page *int64
	Limit *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ProfilePostsReportReasonsResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["message"] = message

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ProfilePostsReportParams struct {
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

	return raw, nil
}

type ProfilePostsCommentsGetParams struct {
}

func (c *Client) ConversationsKick(conversationId int64, userId int64, opts ...func(*ConversationsKickParams)) (map[string]any, error) {
	p := &ConversationsKickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/kick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	body := map[string]any{}
	body["user_id"] = userId

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ConversationsKickParams struct {
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

	return raw, nil
}

type ConversationsReadParams struct {
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

	return raw, nil
}

type TagsPopularParams struct {
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

	return raw, nil
}

type CategoriesGetParams struct {
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

	return raw, nil
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


	body := map[string]any{}
	if p.Alert != nil {
		body["alert"] = *p.Alert
	}
	if p.Email != nil {
		body["email"] = *p.Email
	}
	if p.PrefixIds != nil {
		body["prefix_ids"] = *p.PrefixIds
	}
	if p.MinimalContestAmount != nil {
		body["minimal_contest_amount"] = *p.MinimalContestAmount
	}
	if p.Post != nil {
		body["post"] = *p.Post
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ForumsFollowParams struct {
	Alert *bool
	Email *bool
	PrefixIds *[]int64
	MinimalContestAmount *int64
	Post *bool
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

	return raw, nil
}

type ForumsUnfollowParams struct {
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

	return raw, nil
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

func (c *Client) ThreadsCreate(postBody string, forumId int64, opts ...func(*ThreadsCreateParams)) (*ThreadsCreateResponse, error) {
	p := &ThreadsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads"


	body := map[string]any{}
	body["post_body"] = postBody
	body["forum_id"] = forumId
	if p.Title != nil {
		body["title"] = *p.Title
	}
	if p.ReplyGroup != nil {
		body["reply_group"] = *p.ReplyGroup
	}
	if p.CommentIgnoreGroup != nil {
		body["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.TitleEn != nil {
		body["title_en"] = *p.TitleEn
	}
	if p.DontAlertFollowers != nil {
		body["dont_alert_followers"] = *p.DontAlertFollowers
	}
	if p.WatchThread != nil {
		body["watch_thread"] = *p.WatchThread
	}
	if p.ScheduleDate != nil {
		body["schedule_date"] = *p.ScheduleDate
	}
	if p.ScheduleTime != nil {
		body["schedule_time"] = *p.ScheduleTime
	}
	if p.WatchThreadState != nil {
		body["watch_thread_state"] = *p.WatchThreadState
	}
	if p.PrefixId != nil {
		body["prefix_id"] = *p.PrefixId
	}
	if p.Tags != nil {
		body["tags"] = *p.Tags
	}
	if p.HideContacts != nil {
		body["hide_contacts"] = *p.HideContacts
	}
	if p.WatchThreadEmail != nil {
		body["watch_thread_email"] = *p.WatchThreadEmail
	}
	if p.AllowAskHiddenContent != nil {
		body["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ThreadsCreateResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsCreateParams struct {
	Title *string
	ReplyGroup *int64
	CommentIgnoreGroup *bool
	TitleEn *string
	DontAlertFollowers *bool
	WatchThread *bool
	ScheduleDate *string
	ScheduleTime *string
	WatchThreadState *bool
	PrefixId *[]int64
	Tags *[]string
	HideContacts *bool
	WatchThreadEmail *bool
	AllowAskHiddenContent *bool
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ThreadsGetResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	if p.Title != nil {
		body["title"] = *p.Title
	}
	if p.TitleEn != nil {
		body["title_en"] = *p.TitleEn
	}
	if p.Tags != nil {
		body["tags"] = *p.Tags
	}
	if p.DiscussionOpen != nil {
		body["discussion_open"] = *p.DiscussionOpen
	}
	if p.HideContacts != nil {
		body["hide_contacts"] = *p.HideContacts
	}
	if p.PrefixId != nil {
		body["prefix_id"] = *p.PrefixId
	}
	if p.AllowAskHiddenContent != nil {
		body["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.ReplyGroup != nil {
		body["reply_group"] = *p.ReplyGroup
	}
	if p.CommentIgnoreGroup != nil {
		body["comment_ignore_group"] = *p.CommentIgnoreGroup
	}

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ThreadsEditResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsEditParams struct {
	Title *string
	TitleEn *string
	Tags *[]string
	DiscussionOpen *bool
	HideContacts *bool
	PrefixId *[]int64
	AllowAskHiddenContent *bool
	ReplyGroup *int64
	CommentIgnoreGroup *bool
}

func (c *Client) ThreadsDelete(threadId int64, opts ...func(*ThreadsDeleteParams)) (map[string]any, error) {
	p := &ThreadsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	body := map[string]any{}
	if p.Reason != nil {
		body["reason"] = *p.Reason
	}

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ThreadsDeleteParams struct {
	Reason *string
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsStarResponse
	if err := json.Unmarshal(data, &result); err != nil {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsUnstarResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsUnstarParams struct {
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

	return raw, nil
}

type NotificationsListParams struct {
	TypeValue *string
	Page *int64
	Limit *int64
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

	return raw, nil
}

type NotificationsGetParams struct {
}

func (c *Client) SearchResults(searchId string, opts ...func(*SearchResultsParams)) (map[string]any, error) {
	p := &SearchResultsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/{search_id}/results"
	path = strings.Replace(path, "{search_id}", fmt.Sprintf("%v", searchId), 1)


	body := map[string]any{}
	if p.Page != nil {
		body["page"] = *p.Page
	}
	if p.Limit != nil {
		body["limit"] = *p.Limit
	}

	raw, err := c.DoRequest("GET", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
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

	return raw, nil
}

type PagesListParams struct {
	ParentPageId *int64
	Order *string
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

	return raw, nil
}

type PagesGetParams struct {
}

func (c *Client) ProfilePostsCreate(userId string, postBody string, opts ...func(*ProfilePostsCreateParams)) (map[string]any, error) {
	p := &ProfilePostsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts"


	body := map[string]any{}
	body["user_id"] = userId
	body["post_body"] = postBody

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ProfilePostsCreateParams struct {
}

func (c *Client) ConversationsSearch(opts ...func(*ConversationsSearchParams)) (*ConversationsSearchResponse, error) {
	p := &ConversationsSearchParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/search"


	body := map[string]any{}
	if p.Q != nil {
		body["q"] = *p.Q
	}
	if p.ConversationId != nil {
		body["conversation_id"] = *p.ConversationId
	}
	if p.SearchRecipients != nil {
		body["search_recipients"] = *p.SearchRecipients
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsSearchResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsSearchParams struct {
	Q *string
	ConversationId *int64
	SearchRecipients *bool
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxGetIgnoreResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["user_id"] = userId

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ChatboxPostIgnoreParams struct {
}

func (c *Client) ChatboxDeleteIgnore(userId string, opts ...func(*ChatboxDeleteIgnoreParams)) (map[string]any, error) {
	p := &ChatboxDeleteIgnoreParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/ignore"


	body := map[string]any{}
	body["user_id"] = userId

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ChatboxDeleteIgnoreParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ForumsListResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ForumsListParams struct {
	ParentCategoryId *int64
	ParentForumId *int64
	Order *string
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

	return raw, nil
}

type ForumsGetParams struct {
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

	return raw, nil
}

type ThreadsHideParams struct {
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

	return raw, nil
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

	return raw, nil
}

type ThreadsUnstarParams struct {
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

	return raw, nil
}

type ThreadsFinishParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result PostsReportReasonsResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["message"] = message

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type PostsReportParams struct {
}

func (c *Client) ProfilePostsEdit(profilePostId int64, opts ...func(*ProfilePostsEditParams)) (map[string]any, error) {
	p := &ProfilePostsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%v", profilePostId), 1)


	body := map[string]any{}
	if p.PostBody != nil {
		body["post_body"] = *p.PostBody
	}
	if p.DisableComments != nil {
		body["disable_comments"] = *p.DisableComments
	}

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ProfilePostsEditParams struct {
	PostBody *string
	DisableComments *bool
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

	return raw, nil
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

	return raw, nil
}

type ProfilePostsGetParams struct {
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

	return raw, nil
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


	body := map[string]any{}
	body["profile_post_id"] = profilePostId
	body["comment_body"] = commentBody

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ProfilePostsCommentsCreateParams struct {
}

func (c *Client) ProfilePostsCommentsEdit(commentId int64, commentBody string, opts ...func(*ProfilePostsCommentsEditParams)) (*ProfilePostsCommentsEditResponse, error) {
	p := &ProfilePostsCommentsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/profile-posts/comments"


	body := map[string]any{}
	body["comment_id"] = commentId
	body["comment_body"] = commentBody

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ProfilePostsCommentsEditResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["comment_id"] = commentId

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ProfilePostsCommentsDeleteParams struct {
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

	return raw, nil
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


	body := map[string]any{}
	body["post_body"] = postBody
	if p.ThreadId != nil {
		body["thread_id"] = *p.ThreadId
	}
	if p.QuotePostId != nil {
		body["quote_post_id"] = *p.QuotePostId
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result PostsCreateResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PostsCreateParams struct {
	ThreadId *int64
	QuotePostId *int64
}

func (c *Client) UsersEdit(userId string, opts ...func(*UsersEditParams)) (map[string]any, error) {
	p := &UsersEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	body := map[string]any{}
	if p.AllowSendPersonalConversation != nil {
		body["allow_send_personal_conversation"] = *p.AllowSendPersonalConversation
	}
	if p.AllowInviteGroup != nil {
		body["allow_invite_group"] = *p.AllowInviteGroup
	}
	if p.AllowReceiveNewsFeed != nil {
		body["allow_receive_news_feed"] = *p.AllowReceiveNewsFeed
	}
	if p.SecretAnswer != nil {
		body["secret_answer"] = *p.SecretAnswer
	}
	if p.LanguageId != nil {
		body["language_id"] = *p.LanguageId
	}
	if p.Fields != nil {
		body["fields"] = *p.Fields
	}
	if p.DisplayBannerId != nil {
		body["display_banner_id"] = *p.DisplayBannerId
	}
	if p.UserDobMonth != nil {
		body["user_dob_month"] = *p.UserDobMonth
	}
	if p.Gender != nil {
		body["gender"] = *p.Gender
	}
	if p.ActivityVisible != nil {
		body["activity_visible"] = *p.ActivityVisible
	}
	if p.ShowDobYear != nil {
		body["show_dob_year"] = *p.ShowDobYear
	}
	if p.AllowViewProfile != nil {
		body["allow_view_profile"] = *p.AllowViewProfile
	}
	if p.AllowPostProfile != nil {
		body["allow_post_profile"] = *p.AllowPostProfile
	}
	if p.Alert != nil {
		body["alert"] = *p.Alert
	}
	if p.UserTitle != nil {
		body["user_title"] = *p.UserTitle
	}
	if p.DisplayGroupId != nil {
		body["display_group_id"] = *p.DisplayGroupId
	}
	if p.DisplayIconGroupId != nil {
		body["display_icon_group_id"] = *p.DisplayIconGroupId
	}
	if p.ConvWelcomeMessage != nil {
		body["conv_welcome_message"] = *p.ConvWelcomeMessage
	}
	if p.UserDobYear != nil {
		body["user_dob_year"] = *p.UserDobYear
	}
	if p.ShortLink != nil {
		body["short_link"] = *p.ShortLink
	}
	if p.Timezone != nil {
		body["timezone"] = *p.Timezone
	}
	if p.HideUsernameChangeLogs != nil {
		body["hide_username_change_logs"] = *p.HideUsernameChangeLogs
	}
	if p.Username != nil {
		body["username"] = *p.Username
	}
	if p.UserDobDay != nil {
		body["user_dob_day"] = *p.UserDobDay
	}
	if p.SecretAnswerType != nil {
		body["secret_answer_type"] = *p.SecretAnswerType
	}
	if p.ReceiveAdminEmail != nil {
		body["receive_admin_email"] = *p.ReceiveAdminEmail
	}
	if p.ShowDobDate != nil {
		body["show_dob_date"] = *p.ShowDobDate
	}

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type UsersEditParams struct {
	AllowSendPersonalConversation *string
	AllowInviteGroup *string
	AllowReceiveNewsFeed *string
	SecretAnswer *string
	LanguageId *int64
	Fields *map[string]any
	DisplayBannerId *int64
	UserDobMonth *int64
	Gender *string
	ActivityVisible *bool
	ShowDobYear *bool
	AllowViewProfile *string
	AllowPostProfile *string
	Alert *map[string]any
	UserTitle *string
	DisplayGroupId *int64
	DisplayIconGroupId *int64
	ConvWelcomeMessage *string
	UserDobYear *int64
	ShortLink *string
	Timezone *string
	HideUsernameChangeLogs *bool
	Username *string
	UserDobDay *int64
	SecretAnswerType *int64
	ReceiveAdminEmail *bool
	ShowDobDate *bool
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersGetResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersGetParams struct {
	FieldsInclude *[]string
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersLikesResponse
	if err := json.Unmarshal(data, &result); err != nil {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersIgnoredResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersIgnoredParams struct {
	Total *bool
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersTrophiesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersTrophiesParams struct {
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

	return raw, nil
}

type ProfilePostsListParams struct {
	PostsUserId *int64
	Page *int64
	Limit *int64
	FieldsInclude *[]string
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsMessagesListResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	if p.ReplyMessageId != nil {
		body["reply_message_id"] = *p.ReplyMessageId
	}
	body["message_body"] = messageBody

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsMessagesCreateResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsMessagesCreateParams struct {
	ReplyMessageId *int64
}

func (c *Client) ConversationsInvite(conversationId int64, recipients []string, opts ...func(*ConversationsInviteParams)) (map[string]any, error) {
	p := &ConversationsInviteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/invite"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)


	body := map[string]any{}
	body["recipients"] = recipients

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ConversationsInviteParams struct {
}

func (c *Client) ThreadsPollVote(threadId int64, opts ...func(*ThreadsPollVoteParams)) (map[string]any, error) {
	p := &ThreadsPollVoteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/poll/votes"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	body := map[string]any{}
	if p.ResponseId != nil {
		body["response_id"] = *p.ResponseId
	}
	if p.ResponseIds != nil {
		body["response_ids"] = *p.ResponseIds
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ThreadsPollVoteParams struct {
	ResponseId *int64
	ResponseIds *[]int64
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

	return raw, nil
}

type CategoriesListParams struct {
	ParentCategoryId *int64
	ParentForumId *int64
	Order *string
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

	return raw, nil
}

type ForumsFollowedParams struct {
	Total *bool
}

func (c *Client) ThreadsMove(threadId int64, nodeId string, opts ...func(*ThreadsMoveParams)) (map[string]any, error) {
	p := &ThreadsMoveParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/move"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	body := map[string]any{}
	body["node_id"] = nodeId
	if p.Title != nil {
		body["title"] = *p.Title
	}
	if p.TitleEn != nil {
		body["title_en"] = *p.TitleEn
	}
	if p.PrefixId != nil {
		body["prefix_id"] = *p.PrefixId
	}
	if p.ApplyThreadPrefix != nil {
		body["apply_thread_prefix"] = *p.ApplyThreadPrefix
	}
	if p.SendAlert != nil {
		body["send_alert"] = *p.SendAlert
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ThreadsMoveParams struct {
	Title *string
	TitleEn *string
	PrefixId *[]int64
	ApplyThreadPrefix *bool
	SendAlert *bool
}

func (c *Client) UsersAvatarUpload(userId string, avatar io.Reader, opts ...func(*UsersAvatarUploadParams)) (map[string]any, error) {
	p := &UsersAvatarUploadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/avatar"
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
	files["avatar"] = avatar

	raw, err := c.DoMultipart("POST", path, nil, fields, files)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type UsersAvatarUploadParams struct {
	Y *int64
	Crop *int64
	X *int64
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

	return raw, nil
}

type UsersAvatarDeleteParams struct {
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

	return raw, nil
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

	return raw, nil
}

type ProfilePostsUnstickParams struct {
}

func (c *Client) ConversationsMessagesEdit(conversationId int64, messageId int64, messageBody string, opts ...func(*ConversationsMessagesEditParams)) (*ConversationsMessagesEditResponse, error) {
	p := &ConversationsMessagesEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/{conversation_id}/messages/{message_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%v", conversationId), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%v", messageId), 1)


	body := map[string]any{}
	body["message_body"] = messageBody

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsMessagesEditResponse
	if err := json.Unmarshal(data, &result); err != nil {
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

	return raw, nil
}

type ConversationsMessagesDeleteParams struct {
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

	return raw, nil
}

type TagsListParams struct {
	Page *int64
	Limit *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ForumsGroupedResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ForumsGroupedParams struct {
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

	return raw, nil
}

type ThreadsNavigationParams struct {
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

	return raw, nil
}

type UsersFindParams struct {
	Username *string
	CustomFields *map[string]any
	FieldsInclude *[]string
}

func (c *Client) ConversationsSave(link string, opts ...func(*ConversationsSaveParams)) (map[string]any, error) {
	p := &ConversationsSaveParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/save"


	body := map[string]any{}
	body["link"] = link

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ConversationsSaveParams struct {
}

func (c *Client) SearchAll(opts ...func(*SearchAllParams)) (*SearchAllResponse, error) {
	p := &SearchAllParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search"


	body := map[string]any{}
	if p.UserId != nil {
		body["user_id"] = *p.UserId
	}
	if p.Page != nil {
		body["page"] = *p.Page
	}
	if p.Limit != nil {
		body["limit"] = *p.Limit
	}
	if p.Q != nil {
		body["q"] = *p.Q
	}
	if p.Tag != nil {
		body["tag"] = *p.Tag
	}
	if p.ForumId != nil {
		body["forum_id"] = *p.ForumId
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result SearchAllResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type SearchAllParams struct {
	UserId *string
	Page *int64
	Limit *int64
	Q *string
	Tag *string
	ForumId *int64
}

func (c *Client) SearchProfilePosts(opts ...func(*SearchProfilePostsParams)) (map[string]any, error) {
	p := &SearchProfilePostsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/profile-posts"


	body := map[string]any{}
	if p.Limit != nil {
		body["limit"] = *p.Limit
	}
	if p.Q != nil {
		body["q"] = *p.Q
	}
	if p.UserId != nil {
		body["user_id"] = *p.UserId
	}
	if p.Page != nil {
		body["page"] = *p.Page
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type SearchProfilePostsParams struct {
	Limit *int64
	Q *string
	UserId *int64
	Page *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxIndexResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxIndexParams struct {
	RoomId *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxGetMessagesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxGetMessagesParams struct {
	BeforeMessageId *int64
}

func (c *Client) ChatboxPostMessage(roomId int64, message string, opts ...func(*ChatboxPostMessageParams)) (*ChatboxPostMessageResponse, error) {
	p := &ChatboxPostMessageParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatbox/messages"


	body := map[string]any{}
	body["room_id"] = roomId
	if p.ReplyMessageId != nil {
		body["reply_message_id"] = *p.ReplyMessageId
	}
	body["message"] = message

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxPostMessageResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["message_id"] = messageId
	body["message"] = message

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxEditMessageResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["message_id"] = messageId

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ChatboxDeleteMessageParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result LinksListResponse
	if err := json.Unmarshal(data, &result); err != nil {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result LinksGetResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type LinksGetParams struct {
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

	return raw, nil
}

type ThreadsUnreadParams struct {
	Limit *int64
	ForumId *int64
	DataLimit *int64
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

	return raw, nil
}

type UsersFieldsParams struct {
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

	return raw, nil
}

type ConversationsMessagesStickParams struct {
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

	return raw, nil
}

type ConversationsMessagesUnstickParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsAlertsEnableResponse
	if err := json.Unmarshal(data, &result); err != nil {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsAlertsDisableResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsAlertsDisableParams struct {
}

func (c *Client) SearchThreads(opts ...func(*SearchThreadsParams)) (map[string]any, error) {
	p := &SearchThreadsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/threads"


	body := map[string]any{}
	if p.Limit != nil {
		body["limit"] = *p.Limit
	}
	if p.DataLimit != nil {
		body["data_limit"] = *p.DataLimit
	}
	if p.Q != nil {
		body["q"] = *p.Q
	}
	if p.Tag != nil {
		body["tag"] = *p.Tag
	}
	if p.ForumId != nil {
		body["forum_id"] = *p.ForumId
	}
	if p.UserId != nil {
		body["user_id"] = *p.UserId
	}
	if p.Page != nil {
		body["page"] = *p.Page
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type SearchThreadsParams struct {
	Limit *int64
	DataLimit *int64
	Q *string
	Tag *string
	ForumId *int64
	UserId *string
	Page *int64
}

func (c *Client) SearchPosts(opts ...func(*SearchPostsParams)) (map[string]any, error) {
	p := &SearchPostsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/posts"


	body := map[string]any{}
	if p.Limit != nil {
		body["limit"] = *p.Limit
	}
	if p.DataLimit != nil {
		body["data_limit"] = *p.DataLimit
	}
	if p.Q != nil {
		body["q"] = *p.Q
	}
	if p.Tag != nil {
		body["tag"] = *p.Tag
	}
	if p.ForumId != nil {
		body["forum_id"] = *p.ForumId
	}
	if p.UserId != nil {
		body["user_id"] = *p.UserId
	}
	if p.Page != nil {
		body["page"] = *p.Page
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type SearchPostsParams struct {
	Limit *int64
	DataLimit *int64
	Q *string
	Tag *string
	ForumId *int64
	UserId *string
	Page *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result AssetsCssResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type AssetsCssParams struct {
	Css *[]string
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersClaimsResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersClaimsParams struct {
	TypeValue *string
	ClaimState *string
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

	return raw, nil
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

	return raw, nil
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

	return raw, nil
}

type ProfilePostsLikeParams struct {
}

func (c *Client) ConversationsStart(userId string, opts ...func(*ConversationsStartParams)) (*ConversationsStartResponse, error) {
	p := &ConversationsStartParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations/start"


	body := map[string]any{}
	body["user_id"] = userId

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsStartResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsStartParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result FormsCreateResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type FormsCreateParams struct {
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

	return raw, nil
}

type ThreadsBumpParams struct {
}

func (c *Client) ThreadsFollow(threadId int64, opts ...func(*ThreadsFollowParams)) (map[string]any, error) {
	p := &ThreadsFollowParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%v", threadId), 1)


	body := map[string]any{}
	if p.Email != nil {
		body["email"] = *p.Email
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ThreadsFollowParams struct {
	Email *bool
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

	return raw, nil
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

	return raw, nil
}

type ThreadsFollowersParams struct {
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

	return raw, nil
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


	body := map[string]any{}
	body["post_id"] = postId
	body["comment_body"] = commentBody

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type PostsCommentsCreateParams struct {
}

func (c *Client) PostsCommentsEdit(postCommentId int64, commentBody string, opts ...func(*PostsCommentsEditParams)) (*PostsCommentsEditResponse, error) {
	p := &PostsCommentsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/posts/comments"


	body := map[string]any{}
	body["post_comment_id"] = postCommentId
	body["comment_body"] = commentBody

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result PostsCommentsEditResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	body["post_comment_id"] = postCommentId
	if p.Reason != nil {
		body["reason"] = *p.Reason
	}

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type PostsCommentsDeleteParams struct {
	Reason *string
}

func (c *Client) UsersAvatarCrop(userId string, opts ...func(*UsersAvatarCropParams)) (map[string]any, error) {
	p := &UsersAvatarCropParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/avatar/crop"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	body := map[string]any{}
	if p.X != nil {
		body["x"] = *p.X
	}
	if p.Y != nil {
		body["y"] = *p.Y
	}
	if p.Crop != nil {
		body["crop"] = *p.Crop
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type UsersAvatarCropParams struct {
	X *int64
	Y *int64
	Crop *int64
}

func (c *Client) UsersBackgroundCrop(userId string, opts ...func(*UsersBackgroundCropParams)) (map[string]any, error) {
	p := &UsersBackgroundCropParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/background/crop"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userId), 1)


	body := map[string]any{}
	if p.X != nil {
		body["x"] = *p.X
	}
	if p.Y != nil {
		body["y"] = *p.Y
	}
	if p.Crop != nil {
		body["crop"] = *p.Crop
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type UsersBackgroundCropParams struct {
	X *int64
	Y *int64
	Crop *int64
}

func (c *Client) ConversationsDelete(deleteType string, conversationId int64, opts ...func(*ConversationsDeleteParams)) (map[string]any, error) {
	p := &ConversationsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations"


	body := map[string]any{}
	body["delete_type"] = deleteType
	body["conversation_id"] = conversationId

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ConversationsDeleteParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsListResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	if p.Recipients != nil {
		body["recipients"] = *p.Recipients
	}
	if p.Title != nil {
		body["title"] = *p.Title
	}
	if p.OpenInvite != nil {
		body["open_invite"] = *p.OpenInvite
	}
	if p.AllowStickyMessages != nil {
		body["allow_sticky_messages"] = *p.AllowStickyMessages
	}
	if p.RecipientId != nil {
		body["recipient_id"] = *p.RecipientId
	}
	if p.IsGroup != nil {
		body["is_group"] = *p.IsGroup
	}
	if p.AllowEditMessages != nil {
		body["allow_edit_messages"] = *p.AllowEditMessages
	}
	if p.AllowDeleteOwnMessages != nil {
		body["allow_delete_own_messages"] = *p.AllowDeleteOwnMessages
	}
	if p.MessageBody != nil {
		body["message_body"] = *p.MessageBody
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsCreateResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsCreateParams struct {
	Recipients *[]string
	Title *string
	OpenInvite *bool
	AllowStickyMessages *bool
	RecipientId *int64
	IsGroup *bool
	AllowEditMessages *bool
	AllowDeleteOwnMessages *bool
	MessageBody *string
}

func (c *Client) ConversationsUpdate(conversationId int64, opts ...func(*ConversationsUpdateParams)) (map[string]any, error) {
	p := &ConversationsUpdateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/conversations"


	body := map[string]any{}
	body["conversation_id"] = conversationId
	if p.Title != nil {
		body["title"] = *p.Title
	}
	if p.OpenInvite != nil {
		body["open_invite"] = *p.OpenInvite
	}
	if p.HistoryOpen != nil {
		body["history_open"] = *p.HistoryOpen
	}
	if p.AllowEditMessages != nil {
		body["allow_edit_messages"] = *p.AllowEditMessages
	}
	if p.AllowStickyMessages != nil {
		body["allow_sticky_messages"] = *p.AllowStickyMessages
	}
	if p.AllowDeleteOwnMessages != nil {
		body["allow_delete_own_messages"] = *p.AllowDeleteOwnMessages
	}

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ConversationsUpdateParams struct {
	Title *string
	OpenInvite *bool
	HistoryOpen *bool
	AllowEditMessages *bool
	AllowStickyMessages *bool
	AllowDeleteOwnMessages *bool
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ConversationsReadAllResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ConversationsReadAllParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ChatboxGetLeaderboardResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatboxGetLeaderboardParams struct {
	Duration *string
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

	return raw, nil
}

type NavigationListParams struct {
	Parent *int64
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

	return raw, nil
}

type ThreadsPollGetParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersListResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersListParams struct {
	Page *int64
	Limit *int64
	FieldsInclude *[]string
}

func (c *Client) NotificationsRead(opts ...func(*NotificationsReadParams)) (map[string]any, error) {
	p := &NotificationsReadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/notifications/read"


	body := map[string]any{}
	if p.NotificationId != nil {
		body["notification_id"] = *p.NotificationId
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type NotificationsReadParams struct {
	NotificationId *int64
}

func (c *Client) SearchTagged(opts ...func(*SearchTaggedParams)) (map[string]any, error) {
	p := &SearchTaggedParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/tagged"


	body := map[string]any{}
	if p.Limit != nil {
		body["limit"] = *p.Limit
	}
	if p.Tag != nil {
		body["tag"] = *p.Tag
	}
	if p.Tags != nil {
		body["tags"] = *p.Tags
	}
	if p.Page != nil {
		body["page"] = *p.Page
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type SearchTaggedParams struct {
	Limit *int64
	Tag *string
	Tags *[]string
	Page *int64
}

func (c *Client) UsersBackgroundUpload(userId string, background io.Reader, opts ...func(*UsersBackgroundUploadParams)) (map[string]any, error) {
	p := &UsersBackgroundUploadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/users/{user_id}/background"
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
	files["background"] = background

	raw, err := c.DoMultipart("POST", path, nil, fields, files)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type UsersBackgroundUploadParams struct {
	X *int64
	Y *int64
	Crop *int64
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

	return raw, nil
}

type UsersBackgroundDeleteParams struct {
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

	return raw, nil
}

type UsersContentsParams struct {
	Page *int64
	Limit *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result UsersSAResetResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type UsersSAResetParams struct {
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

	return raw, nil
}

type UsersSACancelResetParams struct {
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

	return raw, nil
}

type ThreadsRecentParams struct {
	Days *int64
	Limit *int64
	ForumId *int64
	DataLimit *int64
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result PostsGetResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	if p.PostBody != nil {
		body["post_body"] = *p.PostBody
	}

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result PostsEditResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	if p.Reason != nil {
		body["reason"] = *p.Reason
	}

	raw, err := c.DoRequest("DELETE", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type PostsDeleteParams struct {
	Reason *string
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

	return raw, nil
}

type ConversationsGetParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result BatchExecuteResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type BatchExecuteParams struct {
}

func (c *Client) ThreadsCreateContest(requireLikeCount int64, postBody string, contestType string, requireTotalLikeCount int64, prizeType string, opts ...func(*ThreadsCreateContestParams)) (*ThreadsCreateContestResponse, error) {
	p := &ThreadsCreateContestParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/contests"


	body := map[string]any{}
	if p.Tags != nil {
		body["tags"] = *p.Tags
	}
	if p.DontAlertFollowers != nil {
		body["dont_alert_followers"] = *p.DontAlertFollowers
	}
	if p.CommentIgnoreGroup != nil {
		body["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.WatchThreadState != nil {
		body["watch_thread_state"] = *p.WatchThreadState
	}
	if p.WatchThreadEmail != nil {
		body["watch_thread_email"] = *p.WatchThreadEmail
	}
	body["require_like_count"] = requireLikeCount
	body["post_body"] = postBody
	body["contest_type"] = contestType
	if p.IsMoneyPlaces != nil {
		body["is_money_places"] = *p.IsMoneyPlaces
	}
	if p.PrizeDataPlaces != nil {
		body["prize_data_places"] = *p.PrizeDataPlaces
	}
	if p.ReplyGroup != nil {
		body["reply_group"] = *p.ReplyGroup
	}
	if p.HideContacts != nil {
		body["hide_contacts"] = *p.HideContacts
	}
	if p.AllowAskHiddenContent != nil {
		body["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.LengthValue != nil {
		body["length_value"] = *p.LengthValue
	}
	if p.LengthOption != nil {
		body["length_option"] = *p.LengthOption
	}
	if p.CountWinners != nil {
		body["count_winners"] = *p.CountWinners
	}
	if p.PrizeDataMoney != nil {
		body["prize_data_money"] = *p.PrizeDataMoney
	}
	if p.ScheduleDate != nil {
		body["schedule_date"] = *p.ScheduleDate
	}
	if p.ScheduleTime != nil {
		body["schedule_time"] = *p.ScheduleTime
	}
	if p.WatchThread != nil {
		body["watch_thread"] = *p.WatchThread
	}
	body["require_total_like_count"] = requireTotalLikeCount
	if p.Title != nil {
		body["title"] = *p.Title
	}
	if p.TitleEn != nil {
		body["title_en"] = *p.TitleEn
	}
	body["prize_type"] = prizeType
	if p.PrizeDataUpgrade != nil {
		body["prize_data_upgrade"] = *p.PrizeDataUpgrade
	}
	if p.SecretAnswer != nil {
		body["secret_answer"] = *p.SecretAnswer
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ThreadsCreateContestResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsCreateContestParams struct {
	Tags *[]string
	DontAlertFollowers *bool
	CommentIgnoreGroup *bool
	WatchThreadState *bool
	WatchThreadEmail *bool
	IsMoneyPlaces *bool
	PrizeDataPlaces *[]float64
	ReplyGroup *int64
	HideContacts *bool
	AllowAskHiddenContent *bool
	LengthValue *int64
	LengthOption *string
	CountWinners *int64
	PrizeDataMoney *float64
	ScheduleDate *string
	ScheduleTime *string
	WatchThread *bool
	Title *string
	TitleEn *string
	PrizeDataUpgrade *int64
	SecretAnswer *string
}

func (c *Client) ThreadsClaim(asAmount float64, transferType string, postBody string, asIsMarketDeal bool, asResponder string, opts ...func(*ThreadsClaimParams)) (*ThreadsClaimResponse, error) {
	p := &ThreadsClaimParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/claims"


	body := map[string]any{}
	if p.AsMarketItemId != nil {
		body["as_market_item_id"] = *p.AsMarketItemId
	}
	body["as_amount"] = asAmount
	body["transfer_type"] = transferType
	if p.ReplyGroup != nil {
		body["reply_group"] = *p.ReplyGroup
	}
	if p.CommentIgnoreGroup != nil {
		body["comment_ignore_group"] = *p.CommentIgnoreGroup
	}
	if p.WatchThread != nil {
		body["watch_thread"] = *p.WatchThread
	}
	if p.WatchThreadEmail != nil {
		body["watch_thread_email"] = *p.WatchThreadEmail
	}
	body["post_body"] = postBody
	if p.AllowAskHiddenContent != nil {
		body["allow_ask_hidden_content"] = *p.AllowAskHiddenContent
	}
	if p.AsData != nil {
		body["as_data"] = *p.AsData
	}
	if p.Currency != nil {
		body["currency"] = *p.Currency
	}
	if p.AsFundsReceipt != nil {
		body["as_funds_receipt"] = *p.AsFundsReceipt
	}
	if p.AsTgLoginScreenshot != nil {
		body["as_tg_login_screenshot"] = *p.AsTgLoginScreenshot
	}
	if p.Tags != nil {
		body["tags"] = *p.Tags
	}
	if p.ScheduleDate != nil {
		body["schedule_date"] = *p.ScheduleDate
	}
	body["as_is_market_deal"] = asIsMarketDeal
	if p.DontAlertFollowers != nil {
		body["dont_alert_followers"] = *p.DontAlertFollowers
	}
	if p.WatchThreadState != nil {
		body["watch_thread_state"] = *p.WatchThreadState
	}
	body["as_responder"] = asResponder
	if p.PayClaim != nil {
		body["pay_claim"] = *p.PayClaim
	}
	if p.HideContacts != nil {
		body["hide_contacts"] = *p.HideContacts
	}
	if p.ScheduleTime != nil {
		body["schedule_time"] = *p.ScheduleTime
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ThreadsClaimResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ThreadsClaimParams struct {
	AsMarketItemId *int64
	ReplyGroup *int64
	CommentIgnoreGroup *bool
	WatchThread *bool
	WatchThreadEmail *bool
	AllowAskHiddenContent *bool
	AsData *string
	Currency *string
	AsFundsReceipt *string
	AsTgLoginScreenshot *string
	Tags *[]string
	ScheduleDate *string
	DontAlertFollowers *bool
	WatchThreadState *bool
	PayClaim *string
	HideContacts *bool
	ScheduleTime *string
}

func (c *Client) SearchUsers(opts ...func(*SearchUsersParams)) (*SearchUsersResponse, error) {
	p := &SearchUsersParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/search/users"


	body := map[string]any{}
	if p.Q != nil {
		body["q"] = *p.Q
	}

	raw, err := c.DoRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result SearchUsersResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type SearchUsersParams struct {
	Q *string
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result OAuthTokenResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type OAuthTokenParams struct {
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

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result ForumsGetFeedOptionsResponse
	if err := json.Unmarshal(data, &result); err != nil {
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


	body := map[string]any{}
	if p.NodeIds != nil {
		body["node_ids"] = *p.NodeIds
	}
	if p.Keywords != nil {
		body["keywords"] = *p.Keywords
	}

	raw, err := c.DoRequest("PUT", path, nil, body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type ForumsEditFeedOptionsParams struct {
	NodeIds *[]int64
	Keywords *[]string
}
