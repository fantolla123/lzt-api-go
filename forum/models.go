package forum

type RespThreadModel struct {
	ThreadIsDeleted bool `json:"thread_is_deleted"`
	ThreadIsStarred bool `json:"thread_is_starred"`
	FirstPost map[string]any `json:"first_post"`
	Links map[string]any `json:"links"`
	Restrictions map[string]any `json:"restrictions"`
	ThreadTitle string `json:"thread_title"`
	CreatorUserId int64 `json:"creator_user_id"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	ThreadIsClosed bool `json:"thread_is_closed"`
	ThreadIsFollowed bool `json:"thread_is_followed"`
	ThreadPrefixes []any `json:"thread_prefixes"`
	ThreadTags map[string]any `json:"thread_tags"`
	NodeTitle string `json:"node_title"`
	ForumId int64 `json:"forum_id"`
	CreatorUsername string `json:"creator_username"`
	ThreadCreateDate int64 `json:"thread_create_date"`
	ThreadUpdateDate int64 `json:"thread_update_date"`
	ThreadPostCount int64 `json:"thread_post_count"`
	ThreadIsSticky bool `json:"thread_is_sticky"`
	LastPost map[string]any `json:"last_post"`
	Contest map[string]any `json:"contest"`
	ThreadIsPublished bool `json:"thread_is_published"`
	Permissions map[string]any `json:"permissions"`
	ThreadId int64 `json:"thread_id"`
	ThreadViewCount int64 `json:"thread_view_count"`
	UserIsIgnored bool `json:"user_is_ignored"`
}

type RespPostModel struct {
	PostBodyHtml string `json:"post_body_html"`
	Signature string `json:"signature"`
	PostLikeCount int64 `json:"post_like_count"`
	PostIsFirstPost bool `json:"post_is_first_post"`
	ThreadIsDeleted bool `json:"thread_is_deleted"`
	PostId int64 `json:"post_id"`
	PosterUserId int64 `json:"poster_user_id"`
	PostBodyPlainText string `json:"post_body_plain_text"`
	SignatureHtml string `json:"signature_html"`
	PostIsPublished bool `json:"post_is_published"`
	PostIsDeleted bool `json:"post_is_deleted"`
	PosterUsernameHtml string `json:"poster_username_html"`
	PostCreateDate int64 `json:"post_create_date"`
	PostBody string `json:"post_body"`
	SignaturePlainText string `json:"signature_plain_text"`
	Permissions map[string]any `json:"permissions"`
	PosterUsername string `json:"poster_username"`
	UserIsIgnored bool `json:"user_is_ignored"`
	PostUpdateDate int64 `json:"post_update_date"`
	Links map[string]any `json:"links"`
	ThreadId int64 `json:"thread_id"`
}

type RespProfilePostModel struct {
	PostIsLiked bool `json:"post_is_liked"`
	TimelineUserId int64 `json:"timeline_user_id"`
	PostBodyPlainText string `json:"post_body_plain_text"`
	PostIsPublished bool `json:"post_is_published"`
	Links map[string]any `json:"links"`
	Permissions map[string]any `json:"permissions"`
	PosterUsername string `json:"poster_username"`
	PostCommentCount int64 `json:"post_comment_count"`
	TimelineUsername string `json:"timeline_username"`
	PostIsDeleted bool `json:"post_is_deleted"`
	PostIsSticked bool `json:"post_is_sticked"`
	TimelineUser RespUserModel `json:"timeline_user"`
	ProfilePostId int64 `json:"profile_post_id"`
	PostCreateDate int64 `json:"post_create_date"`
	PostBody string `json:"post_body"`
	PostCommentsIsDisabled int64 `json:"post_comments_is_disabled"`
	UserIsIgnored bool `json:"user_is_ignored"`
	PosterUserId int64 `json:"poster_user_id"`
	PosterUsernameHtml string `json:"poster_username_html"`
	PostBodyHtml string `json:"post_body_html"`
	PostLikeCount int64 `json:"post_like_count"`
}

type RespSystemInfo struct {
	Time int64 `json:"time"`
	VisitorId int64 `json:"visitor_id"`
}

type UserIDModel struct {
}

type RespNotificationModel struct {
	NotificationId int64 `json:"notification_id"`
	ContentAction string `json:"content_action"`
	NotificationIsUnread bool `json:"notification_is_unread"`
	CreatorUserId int64 `json:"creator_user_id"`
	CreatorUsername string `json:"creator_username"`
	Links map[string]any `json:"links"`
	NotificationCreateDate int64 `json:"notification_create_date"`
	ContentType string `json:"content_type"`
	ContentId int64 `json:"content_id"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	NotificationType string `json:"notification_type"`
	NotificationHtml string `json:"notification_html"`
}

type RespUserModel struct {
	UserDeposit int64 `json:"user_deposit"`
	UserIsFollowed bool `json:"user_is_followed"`
	TrophyCount int64 `json:"trophy_count"`
	UserExternalAuthentications []map[string]any `json:"user_external_authentications"`
	SelfPermissions map[string]any `json:"self_permissions"`
	Birthday map[string]any `json:"birthday"`
	SecretAnswerFirstLetter string `json:"secret_answer_first_letter"`
	UserFollowers map[string]any `json:"user_followers"`
	UsernameHtml string `json:"username_html"`
	UserLike2Count int64 `json:"user_like2_count"`
	CustomTitle string `json:"custom_title"`
	UserIsVerified bool `json:"user_is_verified"`
	UserIsVisitor bool `json:"user_is_visitor"`
	UserGroupId int64 `json:"user_group_id"`
	EditPermissions map[string]any `json:"edit_permissions"`
	Banner string `json:"banner"`
	IsBanned int64 `json:"is_banned"`
	DisplayBannerId int64 `json:"display_banner_id"`
	ConvWelcomeMessage string `json:"conv_welcome_message"`
	UserTitle string `json:"user_title"`
	Permissions map[string]any `json:"permissions"`
	ContestCount int64 `json:"contest_count"`
	DisplayIconGroupId int64 `json:"display_icon_group_id"`
	Balance string `json:"balance"`
	Hold string `json:"hold"`
	UserEmail string `json:"user_email"`
	UserTimezoneOffset int64 `json:"user_timezone_offset"`
	Username string `json:"username"`
	UserUnreadConversationCount int64 `json:"user_unread_conversation_count"`
	Fields []map[string]any `json:"fields"`
	Links map[string]any `json:"links"`
	UserId int64 `json:"user_id"`
	ShortLink string `json:"short_link"`
	UserUnreadNotificationCount int64 `json:"user_unread_notification_count"`
	UserIsValid bool `json:"user_is_valid"`
	UserLastSeenDate int64 `json:"user_last_seen_date"`
	UserIsIgnored bool `json:"user_is_ignored"`
	SecretAnswerRendered string `json:"secret_answer_rendered"`
	UserMessageCount int64 `json:"user_message_count"`
	UserRegisterDate int64 `json:"user_register_date"`
	UserLikeCount int64 `json:"user_like_count"`
	CuratorTitles []string `json:"curator_titles"`
	UserGroups []map[string]any `json:"user_groups"`
	UserFollowing map[string]any `json:"user_following"`
	Currency string `json:"currency"`
}

type RespPostCommentModel struct {
	PostCommentBodyHtml string `json:"post_comment_body_html"`
	PostCommentLikeCount int64 `json:"post_comment_like_count"`
	PostCommentIsPublished bool `json:"post_comment_is_published"`
	PostCommentCreateDate int64 `json:"post_comment_create_date"`
	UserIsIgnored bool `json:"user_is_ignored"`
	PostCommentUpdateDate int64 `json:"post_comment_update_date"`
	ThreadId int64 `json:"thread_id"`
	PosterUsernameHtml string `json:"poster_username_html"`
	PostCommentId int64 `json:"post_comment_id"`
	PostId int64 `json:"post_id"`
	PosterUserId int64 `json:"poster_user_id"`
	PostCommentBody string `json:"post_comment_body"`
	PostCommentBodyPlainText string `json:"post_comment_body_plain_text"`
	PostCommentIsDeleted bool `json:"post_comment_is_deleted"`
	Links map[string]any `json:"links"`
	Permissions map[string]any `json:"permissions"`
	PosterUsername string `json:"poster_username"`
}

type RespProfilePostCommentModel struct {
	Permissions map[string]any `json:"permissions"`
	ProfilePostId int64 `json:"profile_post_id"`
	CommentUsername string `json:"comment_username"`
	CommentUsernameHtml string `json:"comment_username_html"`
	CommentCreateDate int64 `json:"comment_create_date"`
	CommentBodyHtml string `json:"comment_body_html"`
	CommentBodyPlainText string `json:"comment_body_plain_text"`
	CommentId int64 `json:"comment_id"`
	CommentUserId int64 `json:"comment_user_id"`
	CommentBody string `json:"comment_body"`
	UserIsIgnored bool `json:"user_is_ignored"`
	TimelineUserId int64 `json:"timeline_user_id"`
	Links map[string]any `json:"links"`
}

type RespConversationModel struct {
	ConversationCreateDate int64 `json:"conversation_create_date"`
	ConversationUpdateDate int64 `json:"conversation_update_date"`
	Alerts int64 `json:"alerts"`
	Permissions map[string]any `json:"permissions"`
	ConversationMessageCount int64 `json:"conversation_message_count"`
	ConversationIsOpen bool `json:"conversation_is_open"`
	Recipients []map[string]any `json:"recipients"`
	ConversationId int64 `json:"conversation_id"`
	ConversationTitle string `json:"conversation_title"`
	ConversationOnlineCount int64 `json:"conversation_online_count"`
	IsStarred int64 `json:"is_starred"`
	Recipient map[string]any `json:"recipient"`
	Links map[string]any `json:"links"`
	CreatorUserId int64 `json:"creator_user_id"`
	ConversationIsDeleted bool `json:"conversation_is_deleted"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	ConversationLastReadDate int64 `json:"conversation_last_read_date"`
	IsGroup int64 `json:"is_group"`
	IsUnread int64 `json:"is_unread"`
	ConversationIsNew bool `json:"conversation_is_new"`
	CreatorIsIgnored bool `json:"creator_is_ignored"`
	CreatorUsername string `json:"creator_username"`
}

type RespConversationMessageModel struct {
	Permissions map[string]any `json:"permissions"`
	MessageNeedTranslate bool `json:"message_need_translate"`
	MessageEditDate int64 `json:"message_edit_date"`
	MessageBodyHtml string `json:"message_body_html"`
	Links map[string]any `json:"links"`
	ConversationId int64 `json:"conversation_id"`
	CreatorUserId int64 `json:"creator_user_id"`
	CreatorUsername string `json:"creator_username"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	MessageCreateDate int64 `json:"message_create_date"`
	MessageBodyPlainText string `json:"message_body_plain_text"`
	UserIsIgnored bool `json:"user_is_ignored"`
	MessageId int64 `json:"message_id"`
	MessageIsUnread int64 `json:"message_is_unread"`
	MessageIsSystem bool `json:"message_is_system"`
	MessageBody string `json:"message_body"`
}

type RespLinkModel struct {
	LinkId int64 `json:"link_id"`
	LinkTitle string `json:"link_title"`
	LinkDescription string `json:"link_description"`
	Links map[string]any `json:"links"`
	Permissions map[string]any `json:"permissions"`
}

type RespChatboxMessageModel struct {
	CanReport bool `json:"can_report"`
	Date int64 `json:"date"`
	MessageJson string `json:"messageJson"`
	MessageRaw string `json:"messageRaw"`
	IsDeleted bool `json:"is_deleted"`
	Message string `json:"message"`
	MessageId int64 `json:"message_id"`
	Room map[string]any `json:"room"`
	User map[string]any `json:"user"`
}

type UsersSecretAnswerTypesResponse struct {
	Data []map[string]any `json:"data"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxOnlineResponse struct {
	Users []map[string]any `json:"users"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxReportReasonsResponse struct {
	Reasons []string `json:"reasons"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type FormsListResponse struct {
	FormsPerPage int64 `json:"formsPerPage"`
	Page int64 `json:"page"`
	TotalForms int64 `json:"totalForms"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Forms []map[string]any `json:"forms"`
}

type ProfilePostsReportReasonsResponse struct {
	Reasons []string `json:"reasons"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsCreateResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsGetResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsEditResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Thread RespThreadModel `json:"thread"`
}

type ConversationsStarResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type ConversationsUnstarResponse struct {
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Status string `json:"status"`
}

type ConversationsSearchResponse struct {
	Conversations []RespConversationModel `json:"conversations"`
	Recipients bool `json:"recipients"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxGetIgnoreResponse struct {
	Ignored []map[string]any `json:"ignored"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ForumsListResponse struct {
	ForumsTotal int64 `json:"forums_total"`
	Tabs []map[string]any `json:"tabs"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Forums []map[string]any `json:"forums"`
}

type PostsReportReasonsResponse struct {
	Reasons []string `json:"reasons"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ProfilePostsCommentsEditResponse struct {
	Comment map[string]any `json:"comment"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsCreateResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Post RespPostModel `json:"post"`
}

type UsersGetResponse struct {
	User RespUserModel `json:"user"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersLikesResponse struct {
	Page int64 `json:"page"`
	PerPage int64 `json:"perPage"`
	ContentType string `json:"contentType"`
	TotalLikes int64 `json:"totalLikes"`
	Likes map[string]any `json:"likes"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersIgnoredResponse struct {
	Users []map[string]any `json:"users"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersTrophiesResponse struct {
	Trophies []map[string]any `json:"trophies"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsMessagesListResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Messages []RespConversationMessageModel `json:"messages"`
	MessagesTotal int64 `json:"messages_total"`
	Links map[string]any `json:"links"`
}

type ConversationsMessagesCreateResponse struct {
	Message RespConversationMessageModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsMessagesEditResponse struct {
	Message RespConversationModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ForumsGroupedResponse struct {
	Data map[string]any `json:"data"`
	Tabs []map[string]any `json:"tabs"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type SearchAllResponse struct {
	Data []map[string]any `json:"data"`
	DataTotal int64 `json:"data_total"`
	Users []RespUserModel `json:"users"`
	Links map[string]any `json:"links"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxIndexResponse struct {
	Rooms []map[string]any `json:"rooms"`
	Ban any `json:"ban"`
	Ignore []map[string]any `json:"ignore"`
	Permissions map[string]any `json:"permissions"`
	Commands []string `json:"commands"`
	RoomsOnline map[string]any `json:"roomsOnline"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxGetMessagesResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Messages []RespChatboxMessageModel `json:"messages"`
}

type ChatboxPostMessageResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Message RespChatboxMessageModel `json:"message"`
}

type ChatboxEditMessageResponse struct {
	Message RespChatboxMessageModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type LinksListResponse struct {
	LinkForums []RespLinkModel `json:"link-forums"`
	LinkForumsTotal int64 `json:"link-forums_total"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type LinksGetResponse struct {
	LinkForum RespLinkModel `json:"link-forum"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsAlertsEnableResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsAlertsDisableResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type AssetsCssResponse struct {
	Contents string `json:"contents"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersClaimsResponse struct {
	Claims []map[string]any `json:"claims"`
	Stats map[string]any `json:"stats"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsStartResponse struct {
	Conversation RespConversationModel `json:"conversation"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type FormsCreateResponse struct {
	Message string `json:"message"`
	Content map[string]any `json:"content"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsCommentsEditResponse struct {
	Comment map[string]any `json:"comment"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsListResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Conversations []RespConversationModel `json:"conversations"`
	CanStart bool `json:"can_start"`
	Folders []map[string]any `json:"folders"`
	Links map[string]any `json:"links"`
}

type ConversationsCreateResponse struct {
	Conversation RespConversationModel `json:"conversation"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsReadAllResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type ChatboxGetLeaderboardResponse struct {
	Leaderboard []map[string]any `json:"leaderboard"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersListResponse struct {
	Users []RespUserModel `json:"users"`
	UsersTotal int64 `json:"users_total"`
	Links map[string]any `json:"links"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersSAResetResponse struct {
	Success bool `json:"success"`
	WaitingTime string `json:"waiting_time"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsGetResponse struct {
	Post RespPostModel `json:"post"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsEditResponse struct {
	Post RespPostModel `json:"post"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type BatchExecuteResponse struct {
	Jobs map[string]any `json:"jobs"`
}

type ThreadsCreateContestResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsClaimResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type SearchUsersResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Users []RespUserModel `json:"users"`
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

type ForumsGetFeedOptionsResponse struct {
	DefaultExcludedForumsIds []int64 `json:"default_excluded_forums_ids"`
	Keywords string `json:"keywords"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Forums []map[string]any `json:"forums"`
	ExcludedForumsIds []int64 `json:"excluded_forums_ids"`
}
