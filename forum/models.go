package forum

type RespUserModel struct {
	Balance string `json:"balance"`
	Currency string `json:"currency"`
	SecretAnswerRendered string `json:"secret_answer_rendered"`
	TrophyCount int64 `json:"trophy_count"`
	UserIsFollowed bool `json:"user_is_followed"`
	Fields []map[string]any `json:"fields"`
	UserExternalAuthentications []map[string]any `json:"user_external_authentications"`
	UsernameHtml string `json:"username_html"`
	Hold string `json:"hold"`
	UserUnreadNotificationCount int64 `json:"user_unread_notification_count"`
	ConvWelcomeMessage string `json:"conv_welcome_message"`
	UserIsIgnored bool `json:"user_is_ignored"`
	UserGroups []map[string]any `json:"user_groups"`
	SecretAnswerFirstLetter string `json:"secret_answer_first_letter"`
	UserLikeCount int64 `json:"user_like_count"`
	IsBanned int64 `json:"is_banned"`
	UserTitle string `json:"user_title"`
	UserIsValid bool `json:"user_is_valid"`
	UserTimezoneOffset int64 `json:"user_timezone_offset"`
	SelfPermissions map[string]any `json:"self_permissions"`
	UserMessageCount int64 `json:"user_message_count"`
	UserRegisterDate int64 `json:"user_register_date"`
	UserLike2Count int64 `json:"user_like2_count"`
	ContestCount int64 `json:"contest_count"`
	DisplayBannerId int64 `json:"display_banner_id"`
	UserIsVisitor bool `json:"user_is_visitor"`
	CuratorTitles []string `json:"curator_titles"`
	UserFollowing map[string]any `json:"user_following"`
	Username string `json:"username"`
	CustomTitle string `json:"custom_title"`
	UserLastSeenDate int64 `json:"user_last_seen_date"`
	EditPermissions map[string]any `json:"edit_permissions"`
	Banner string `json:"banner"`
	UserId int64 `json:"user_id"`
	ShortLink string `json:"short_link"`
	UserUnreadConversationCount int64 `json:"user_unread_conversation_count"`
	Permissions map[string]any `json:"permissions"`
	Birthday map[string]any `json:"birthday"`
	UserEmail string `json:"user_email"`
	Links map[string]any `json:"links"`
	UserGroupId int64 `json:"user_group_id"`
	UserFollowers map[string]any `json:"user_followers"`
	UserDeposit int64 `json:"user_deposit"`
	UserIsVerified bool `json:"user_is_verified"`
	DisplayIconGroupId int64 `json:"display_icon_group_id"`
}

type RespThreadModel struct {
	LastPost map[string]any `json:"last_post"`
	ThreadId int64 `json:"thread_id"`
	ThreadIsSticky bool `json:"thread_is_sticky"`
	Permissions map[string]any `json:"permissions"`
	Restrictions map[string]any `json:"restrictions"`
	Contest map[string]any `json:"contest"`
	ThreadPostCount int64 `json:"thread_post_count"`
	ThreadIsPublished bool `json:"thread_is_published"`
	ThreadIsDeleted bool `json:"thread_is_deleted"`
	ThreadTitle string `json:"thread_title"`
	ThreadViewCount int64 `json:"thread_view_count"`
	CreatorUserId int64 `json:"creator_user_id"`
	CreatorUsername string `json:"creator_username"`
	ThreadCreateDate int64 `json:"thread_create_date"`
	ForumId int64 `json:"forum_id"`
	ThreadUpdateDate int64 `json:"thread_update_date"`
	UserIsIgnored bool `json:"user_is_ignored"`
	ThreadIsClosed bool `json:"thread_is_closed"`
	ThreadIsFollowed bool `json:"thread_is_followed"`
	ThreadIsStarred bool `json:"thread_is_starred"`
	FirstPost map[string]any `json:"first_post"`
	ThreadPrefixes []any `json:"thread_prefixes"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	ThreadTags map[string]any `json:"thread_tags"`
	Links map[string]any `json:"links"`
	NodeTitle string `json:"node_title"`
}

type RespPostCommentModel struct {
	PostId int64 `json:"post_id"`
	PostCommentIsPublished bool `json:"post_comment_is_published"`
	PostCommentIsDeleted bool `json:"post_comment_is_deleted"`
	PostCommentUpdateDate int64 `json:"post_comment_update_date"`
	PostCommentId int64 `json:"post_comment_id"`
	ThreadId int64 `json:"thread_id"`
	PostCommentBodyHtml string `json:"post_comment_body_html"`
	Links map[string]any `json:"links"`
	PosterUserId int64 `json:"poster_user_id"`
	PostCommentCreateDate int64 `json:"post_comment_create_date"`
	PostCommentBody string `json:"post_comment_body"`
	UserIsIgnored bool `json:"user_is_ignored"`
	Permissions map[string]any `json:"permissions"`
	PosterUsername string `json:"poster_username"`
	PosterUsernameHtml string `json:"poster_username_html"`
	PostCommentBodyPlainText string `json:"post_comment_body_plain_text"`
	PostCommentLikeCount int64 `json:"post_comment_like_count"`
}

type RespProfilePostModel struct {
	TimelineUserId int64 `json:"timeline_user_id"`
	PosterUserId int64 `json:"poster_user_id"`
	PostCommentCount int64 `json:"post_comment_count"`
	PostIsPublished bool `json:"post_is_published"`
	PostIsSticked bool `json:"post_is_sticked"`
	Links map[string]any `json:"links"`
	PostIsDeleted bool `json:"post_is_deleted"`
	TimelineUser RespUserModel `json:"timeline_user"`
	PosterUsername string `json:"poster_username"`
	PostBodyPlainText string `json:"post_body_plain_text"`
	PostLikeCount int64 `json:"post_like_count"`
	TimelineUsername string `json:"timeline_username"`
	PostIsLiked bool `json:"post_is_liked"`
	Permissions map[string]any `json:"permissions"`
	PosterUsernameHtml string `json:"poster_username_html"`
	PostCreateDate int64 `json:"post_create_date"`
	PostBody string `json:"post_body"`
	PostBodyHtml string `json:"post_body_html"`
	PostCommentsIsDisabled int64 `json:"post_comments_is_disabled"`
	UserIsIgnored bool `json:"user_is_ignored"`
	ProfilePostId int64 `json:"profile_post_id"`
}

type RespProfilePostCommentModel struct {
	CommentUsername string `json:"comment_username"`
	CommentCreateDate int64 `json:"comment_create_date"`
	CommentBody string `json:"comment_body"`
	CommentBodyHtml string `json:"comment_body_html"`
	CommentBodyPlainText string `json:"comment_body_plain_text"`
	UserIsIgnored bool `json:"user_is_ignored"`
	ProfilePostId int64 `json:"profile_post_id"`
	CommentUsernameHtml string `json:"comment_username_html"`
	TimelineUserId int64 `json:"timeline_user_id"`
	Links map[string]any `json:"links"`
	Permissions map[string]any `json:"permissions"`
	CommentId int64 `json:"comment_id"`
	CommentUserId int64 `json:"comment_user_id"`
}

type RespConversationModel struct {
	Recipients []map[string]any `json:"recipients"`
	Links map[string]any `json:"links"`
	ConversationTitle string `json:"conversation_title"`
	CreatorUserId int64 `json:"creator_user_id"`
	Alerts int64 `json:"alerts"`
	ConversationIsOpen bool `json:"conversation_is_open"`
	Permissions map[string]any `json:"permissions"`
	ConversationMessageCount int64 `json:"conversation_message_count"`
	ConversationId int64 `json:"conversation_id"`
	CreatorUsername string `json:"creator_username"`
	ConversationUpdateDate int64 `json:"conversation_update_date"`
	ConversationLastReadDate int64 `json:"conversation_last_read_date"`
	ConversationOnlineCount int64 `json:"conversation_online_count"`
	IsStarred int64 `json:"is_starred"`
	IsGroup int64 `json:"is_group"`
	ConversationIsNew bool `json:"conversation_is_new"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	ConversationCreateDate int64 `json:"conversation_create_date"`
	IsUnread int64 `json:"is_unread"`
	CreatorIsIgnored bool `json:"creator_is_ignored"`
	ConversationIsDeleted bool `json:"conversation_is_deleted"`
	Recipient map[string]any `json:"recipient"`
}

type RespConversationMessageModel struct {
	CreatorUsername string `json:"creator_username"`
	MessageIsUnread int64 `json:"message_is_unread"`
	MessageNeedTranslate bool `json:"message_need_translate"`
	MessageIsSystem bool `json:"message_is_system"`
	MessageEditDate int64 `json:"message_edit_date"`
	MessageBodyHtml string `json:"message_body_html"`
	MessageId int64 `json:"message_id"`
	ConversationId int64 `json:"conversation_id"`
	CreatorUsernameHtml string `json:"creator_username_html"`
	MessageBodyPlainText string `json:"message_body_plain_text"`
	Permissions map[string]any `json:"permissions"`
	UserIsIgnored bool `json:"user_is_ignored"`
	Links map[string]any `json:"links"`
	CreatorUserId int64 `json:"creator_user_id"`
	MessageCreateDate int64 `json:"message_create_date"`
	MessageBody string `json:"message_body"`
}

type UserIDModel struct {
}

type RespNotificationModel struct {
	CreatorUsernameHtml string `json:"creator_username_html"`
	Links map[string]any `json:"links"`
	NotificationId int64 `json:"notification_id"`
	ContentType string `json:"content_type"`
	ContentId int64 `json:"content_id"`
	NotificationIsUnread bool `json:"notification_is_unread"`
	CreatorUsername string `json:"creator_username"`
	NotificationType string `json:"notification_type"`
	NotificationHtml string `json:"notification_html"`
	NotificationCreateDate int64 `json:"notification_create_date"`
	ContentAction string `json:"content_action"`
	CreatorUserId int64 `json:"creator_user_id"`
}

type RespLinkModel struct {
	LinkId int64 `json:"link_id"`
	LinkTitle string `json:"link_title"`
	LinkDescription string `json:"link_description"`
	Links map[string]any `json:"links"`
	Permissions map[string]any `json:"permissions"`
}

type RespChatboxMessageModel struct {
	Date int64 `json:"date"`
	MessageId int64 `json:"message_id"`
	MessageJson string `json:"messageJson"`
	MessageRaw string `json:"messageRaw"`
	Room map[string]any `json:"room"`
	User map[string]any `json:"user"`
	IsDeleted bool `json:"is_deleted"`
	Message string `json:"message"`
	CanReport bool `json:"can_report"`
}

type RespPostModel struct {
	PostBody string `json:"post_body"`
	PostBodyPlainText string `json:"post_body_plain_text"`
	SignatureHtml string `json:"signature_html"`
	SignaturePlainText string `json:"signature_plain_text"`
	PostUpdateDate int64 `json:"post_update_date"`
	Links map[string]any `json:"links"`
	Permissions map[string]any `json:"permissions"`
	PostId int64 `json:"post_id"`
	PosterUserId int64 `json:"poster_user_id"`
	PosterUsernameHtml string `json:"poster_username_html"`
	PostLikeCount int64 `json:"post_like_count"`
	PostIsDeleted bool `json:"post_is_deleted"`
	ThreadIsDeleted bool `json:"thread_is_deleted"`
	PostCreateDate int64 `json:"post_create_date"`
	Signature string `json:"signature"`
	UserIsIgnored bool `json:"user_is_ignored"`
	PostIsPublished bool `json:"post_is_published"`
	ThreadId int64 `json:"thread_id"`
	PosterUsername string `json:"poster_username"`
	PostBodyHtml string `json:"post_body_html"`
	PostIsFirstPost bool `json:"post_is_first_post"`
}

type RespSystemInfo struct {
	VisitorId int64 `json:"visitor_id"`
	Time int64 `json:"time"`
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

type UsersLikesResponse struct {
	Page int64 `json:"page"`
	PerPage int64 `json:"perPage"`
	ContentType string `json:"contentType"`
	TotalLikes int64 `json:"totalLikes"`
	Likes map[string]any `json:"likes"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsListResponse struct {
	Conversations []RespConversationModel `json:"conversations"`
	CanStart bool `json:"can_start"`
	Folders []map[string]any `json:"folders"`
	Links map[string]any `json:"links"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsCreateResponse struct {
	Conversation RespConversationModel `json:"conversation"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxGetLeaderboardResponse struct {
	Leaderboard []map[string]any `json:"leaderboard"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsCommentsEditResponse struct {
	Comment map[string]any `json:"comment"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsClaimResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersIgnoredResponse struct {
	Users []map[string]any `json:"users"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsMessagesListResponse struct {
	MessagesTotal int64 `json:"messages_total"`
	Links map[string]any `json:"links"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Messages []RespConversationMessageModel `json:"messages"`
}

type ConversationsMessagesCreateResponse struct {
	Message RespConversationMessageModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersSAResetResponse struct {
	Success bool `json:"success"`
	WaitingTime string `json:"waiting_time"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxOnlineResponse struct {
	Users []map[string]any `json:"users"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type FormsListResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Forms []map[string]any `json:"forms"`
	FormsPerPage int64 `json:"formsPerPage"`
	Page int64 `json:"page"`
	TotalForms int64 `json:"totalForms"`
}

type ForumsGroupedResponse struct {
	Data map[string]any `json:"data"`
	Tabs []map[string]any `json:"tabs"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsReadAllResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ForumsListResponse struct {
	Forums []map[string]any `json:"forums"`
	ForumsTotal int64 `json:"forums_total"`
	Tabs []map[string]any `json:"tabs"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersListResponse struct {
	UsersTotal int64 `json:"users_total"`
	Links map[string]any `json:"links"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Users []RespUserModel `json:"users"`
}

type UsersSecretAnswerTypesResponse struct {
	Data []map[string]any `json:"data"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type AssetsCssResponse struct {
	Contents string `json:"contents"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersGetResponse struct {
	User RespUserModel `json:"user"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ForumsGetFeedOptionsResponse struct {
	Forums []map[string]any `json:"forums"`
	ExcludedForumsIds []int64 `json:"excluded_forums_ids"`
	DefaultExcludedForumsIds []int64 `json:"default_excluded_forums_ids"`
	Keywords string `json:"keywords"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsCreateContestResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersClaimsResponse struct {
	Stats map[string]any `json:"stats"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Claims []map[string]any `json:"claims"`
}

type ProfilePostsReportReasonsResponse struct {
	Reasons []string `json:"reasons"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsStarResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsUnstarResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsStartResponse struct {
	Conversation RespConversationModel `json:"conversation"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type SearchUsersResponse struct {
	Users []RespUserModel `json:"users"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type BatchExecuteResponse struct {
	Jobs map[string]any `json:"jobs"`
}

type ChatboxReportReasonsResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Reasons []string `json:"reasons"`
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

type ThreadsGetResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsEditResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsMessagesEditResponse struct {
	Message RespConversationModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type FormsCreateResponse struct {
	Message string `json:"message"`
	Content map[string]any `json:"content"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxIndexResponse struct {
	Commands []string `json:"commands"`
	RoomsOnline map[string]any `json:"roomsOnline"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Rooms []map[string]any `json:"rooms"`
	Ban any `json:"ban"`
	Ignore []map[string]any `json:"ignore"`
	Permissions map[string]any `json:"permissions"`
}

type ChatboxPostMessageResponse struct {
	Message RespChatboxMessageModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxEditMessageResponse struct {
	Message RespChatboxMessageModel `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxGetMessagesResponse struct {
	Messages []RespChatboxMessageModel `json:"messages"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ThreadsCreateResponse struct {
	Thread RespThreadModel `json:"thread"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ChatboxGetIgnoreResponse struct {
	Ignored []map[string]any `json:"ignored"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsSearchResponse struct {
	Conversations []RespConversationModel `json:"conversations"`
	Recipients bool `json:"recipients"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsCreateResponse struct {
	Post RespPostModel `json:"post"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ProfilePostsCommentsEditResponse struct {
	Comment map[string]any `json:"comment"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsAlertsEnableResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ConversationsAlertsDisableResponse struct {
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Status string `json:"status"`
}

type SearchAllResponse struct {
	Data []map[string]any `json:"data"`
	DataTotal int64 `json:"data_total"`
	Users []RespUserModel `json:"users"`
	Links map[string]any `json:"links"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type UsersTrophiesResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Trophies []map[string]any `json:"trophies"`
}

type PostsGetResponse struct {
	Post RespPostModel `json:"post"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsEditResponse struct {
	Post RespPostModel `json:"post"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PostsReportReasonsResponse struct {
	Reasons []string `json:"reasons"`
	SystemInfo RespSystemInfo `json:"system_info"`
}
