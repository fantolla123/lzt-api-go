package market

type DiscountModel struct {
	DiscountPercent int64 `json:"discount_percent"`
	DiscountUserId int64 `json:"discount_user_id"`
	MaxPrice int64 `json:"max_price"`
	MinPrice int64 `json:"min_price"`
	UserId int64 `json:"user_id"`
	CategoryId int64 `json:"category_id"`
	DiscountId int64 `json:"discount_id"`
}

type BalanceModel struct {
	Balance string `json:"balance"`
	BalanceId int64 `json:"balance_id"`
	CustomTitle any `json:"custom_title"`
	FullTitle string `json:"fullTitle"`
	MerchantId int64 `json:"merchant_id"`
	Title string `json:"title"`
	TypeValue string `json:"type"`
	UserId int64 `json:"user_id"`
}

type RespSystemInfo struct {
	Time int64 `json:"time"`
	LogId int64 `json:"log_id"`
	VisitorId int64 `json:"visitor_id"`
}

type UserModel struct {
	RegisterDate int64 `json:"register_date"`
	ViewUrl string `json:"view_url"`
	CanViewProfile bool `json:"can_view_profile"`
	CanViewProfilePosts bool `json:"can_view_profile_posts"`
	Like2Count int64 `json:"like2_count"`
	TelegramClient map[string]any `json:"telegram_client"`
	TrophyPoints int64 `json:"trophy_points"`
	ConvertedBalance int64 `json:"convertedBalance"`
	Dob map[string]any `json:"dob"`
	IsBanned bool `json:"is_banned"`
	UserAllowAskDiscount bool `json:"user_allow_ask_discount"`
	Balance string `json:"balance"`
	ContestCount int64 `json:"contest_count"`
	FeedbackData map[string]any `json:"feedback_data"`
	MaxDiscountPercent int64 `json:"max_discount_percent"`
	RestoreData map[string]any `json:"restore_data"`
	BumpItemPeriod int64 `json:"bump_item_period"`
	ConvWelcomeMessage string `json:"conv_welcome_message"`
	WarningPoints int64 `json:"warning_points"`
	CurrencyPhrase string `json:"currencyPhrase"`
	LastActivity int64 `json:"last_activity"`
	Username string `json:"username"`
	Hold string `json:"hold"`
	ConvertedDeposit int64 `json:"convertedDeposit"`
	ConvertedHold int64 `json:"convertedHold"`
	CustomFields map[string]any `json:"custom_fields"`
	IsStaff bool `json:"is_staff"`
	RestoreCount int64 `json:"restore_count"`
	UserTitle string `json:"user_title"`
	ActivityVisible bool `json:"activity_visible"`
	ImapData map[string]any `json:"imap_data"`
	IsAdmin bool `json:"is_admin"`
	IsFollowed bool `json:"is_followed"`
	IsSuperAdmin bool `json:"is_super_admin"`
	MessageCount int64 `json:"message_count"`
	PaidMailLeft int64 `json:"paid_mail_left"`
	CanEdit bool `json:"can_edit"`
	Homepage string `json:"homepage"`
	IsModerator bool `json:"is_moderator"`
	UserId int64 `json:"user_id"`
	Age int64 `json:"age"`
	CanWarn bool `json:"can_warn"`
	Rendered map[string]any `json:"rendered"`
	CanIgnore bool `json:"can_ignore"`
	Deposit int64 `json:"deposit"`
	IsIgnored bool `json:"is_ignored"`
	LikeCount int64 `json:"like_count"`
	CanPostProfile bool `json:"can_post_profile"`
	CustomAccountDownloadFormat string `json:"custom_account_download_format"`
	CustomTitle string `json:"custom_title"`
	MarketCustomTitle string `json:"market_custom_title"`
	PublicTags []map[string]any `json:"public_tags"`
	ShortLink string `json:"short_link"`
	Visible bool `json:"visible"`
	Balances []map[string]any `json:"balances"`
	JoinedDate int64 `json:"joined_date"`
	Location string `json:"location"`
	SoldItemsCount int64 `json:"sold_items_count"`
	Tags []map[string]any `json:"tags"`
	ActiveItemsCount int64 `json:"active_items_count"`
	CanFollow bool `json:"can_follow"`
	Currency string `json:"currency"`
}

type InvoiceModel struct {
	ResendAttempts int64 `json:"resend_attempts"`
	Amount int64 `json:"amount"`
	InvoiceId int64 `json:"invoice_id"`
	PayerUserId int64 `json:"payer_user_id"`
	UrlCallback string `json:"url_callback"`
	AdditionalData string `json:"additional_data"`
	IsTest bool `json:"is_test"`
	Url string `json:"url"`
	UserId int64 `json:"user_id"`
	Comment string `json:"comment"`
	ExpiresAt int64 `json:"expires_at"`
	InvoiceDate int64 `json:"invoice_date"`
	MerchantId int64 `json:"merchant_id"`
	PaidDate int64 `json:"paid_date"`
	PaymentId string `json:"payment_id"`
	Status string `json:"status"`
	UrlSuccess string `json:"url_success"`
}

type ConfirmationCodeModel struct {
	Item ItemModel `json:"item"`
	CodeData map[string]any `json:"codeData"`
}

type ItemListModel struct {
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []ItemFromListModel `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	StickyItems []ItemFromListModel `json:"stickyItems"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
}

type ItemFromListModel struct {
	IsSticky *int64 `json:"is_sticky,omitempty"`
	ItemOrigin *string `json:"item_origin,omitempty"`
	ResaleItemOrigin *string `json:"resale_item_origin,omitempty"`
	CanUnstickItem *bool `json:"canUnstickItem,omitempty"`
	CanResellItemAfterPurchase *bool `json:"canResellItemAfterPurchase,omitempty"`
	PublishedDate *int64 `json:"published_date,omitempty"`
	Price *int64 `json:"price,omitempty"`
	AllowAskDiscount *int64 `json:"allow_ask_discount,omitempty"`
	IsIgnored *int64 `json:"isIgnored,omitempty"`
	CanValidateAccount *bool `json:"canValidateAccount,omitempty"`
	CanViewAccountLink *bool `json:"canViewAccountLink,omitempty"`
	Nsb *int64 `json:"nsb,omitempty"`
	CanStickItem *bool `json:"canStickItem,omitempty"`
	NoteText *string `json:"note_text,omitempty"`
	DescriptionHtmlEn *string `json:"description_html_en,omitempty"`
	CategoryId *int64 `json:"category_id,omitempty"`
	CanUpdateItemStats *bool `json:"canUpdateItemStats,omitempty"`
	DescriptionHtml *string `json:"description_html,omitempty"`
	Guarantee *bool `json:"guarantee,omitempty"`
	CanEditItem *bool `json:"canEditItem,omitempty"`
	ItemId *int64 `json:"item_id,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	ExtendedGuarantee *int64 `json:"extended_guarantee,omitempty"`
	ItemDomain *string `json:"item_domain,omitempty"`
	RubPrice *int64 `json:"rub_price,omitempty"`
	Title *string `json:"title,omitempty"`
	RefreshedDate *int64 `json:"refreshed_date,omitempty"`
	TitleEn *string `json:"title_en,omitempty"`
	ShowGetEmailCodeButton *bool `json:"showGetEmailCodeButton,omitempty"`
	PriceCurrency *string `json:"price_currency,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UpdateStatDate *int64 `json:"update_stat_date,omitempty"`
	CanViewLoginData *bool `json:"canViewLoginData,omitempty"`
	CanBumpItem *bool `json:"canBumpItem,omitempty"`
	ItemOriginPhrase *string `json:"itemOriginPhrase,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionEn *string `json:"description_en,omitempty"`
	CanOpenItem *bool `json:"canOpenItem,omitempty"`
	CanCloseItem *bool `json:"canCloseItem,omitempty"`
	CanDeleteItem *bool `json:"canDeleteItem,omitempty"`
	BumpSettings map[string]any `json:"bumpSettings,omitempty"`
	CanBuyItem *bool `json:"canBuyItem,omitempty"`
	Seller map[string]any `json:"seller,omitempty"`
	CanViewEmailLoginData *bool `json:"canViewEmailLoginData,omitempty"`
	ViewCount *int64 `json:"view_count,omitempty"`
}

type ExtraModel struct {
	ConfirmationCode *string `json:"confirmationCode,omitempty"`
	Dota2Mmr *int64 `json:"dota2_mmr,omitempty"`
	EaGames *bool `json:"ea_games,omitempty"`
	UplayGames *bool `json:"uplay_games,omitempty"`
	Warframe *bool `json:"warframe,omitempty"`
	HonkaiCurrency *int64 `json:"honkai_currency,omitempty"`
	Region *string `json:"region,omitempty"`
	MfaFile *string `json:"mfa_file,omitempty"`
	ZenlessCurrency *int64 `json:"zenless_currency,omitempty"`
	TelegramClient *string `json:"telegramClient,omitempty"`
	TelegramJson *string `json:"telegramJson,omitempty"`
	CheckHypixelBan *bool `json:"checkHypixelBan,omitempty"`
	System *string `json:"system,omitempty"`
	Cookies *string `json:"cookies,omitempty"`
	ArkAscended *bool `json:"ark_ascended,omitempty"`
	GenshinCurrency *int64 `json:"genshin_currency,omitempty"`
	Password *string `json:"password,omitempty"`
	CheckChannels *bool `json:"checkChannels,omitempty"`
	CloseItem *bool `json:"close_item,omitempty"`
	Service *string `json:"service,omitempty"`
	LoginWithoutCookies *bool `json:"login_without_cookies,omitempty"`
	CookieLogin *bool `json:"cookie_login,omitempty"`
	TheQuarry *bool `json:"the_quarry,omitempty"`
	Ark *bool `json:"ark,omitempty"`
	CheckSpam *bool `json:"checkSpam,omitempty"`
	Proxy *string `json:"proxy,omitempty"`
}

type ItemModel struct {
	Login string `json:"login"`
	Information string `json:"information"`
	MarketCustomTitle string `json:"market_custom_title"`
	GetEmailCodeDisplayLogin any `json:"getEmailCodeDisplayLogin"`
	RubPrice int64 `json:"rub_price"`
	AccountLink string `json:"accountLink"`
	Tags map[string]any `json:"tags"`
	ViewCount int64 `json:"view_count"`
	Description string `json:"description"`
	EditDate int64 `json:"edit_date"`
	PriceCurrency string `json:"price_currency"`
	CanChangePassword bool `json:"canChangePassword"`
	IsTrusted bool `json:"isTrusted"`
	CanCheckAiPrice bool `json:"canCheckAiPrice"`
	CanCheckAutoBuyPrice bool `json:"canCheckAutoBuyPrice"`
	ItemOrigin string `json:"item_origin"`
	Nsb int64 `json:"nsb"`
	ItemDomain string `json:"item_domain"`
	InCart any `json:"in_cart"`
	PriceWithSellerFee float64 `json:"priceWithSellerFee"`
	CanUpdateItemStats bool `json:"canUpdateItemStats"`
	AiPriceCheckDate int64 `json:"aiPriceCheckDate"`
	NeedToRequireVideoToViewLoginData bool `json:"needToRequireVideoToViewLoginData"`
	ContentId any `json:"content_id"`
	DeleteUserId int64 `json:"delete_user_id"`
	Buyer map[string]any `json:"buyer"`
	BumpSettings map[string]any `json:"bumpSettings"`
	AiPrice int64 `json:"aiPrice"`
	Seller map[string]any `json:"seller"`
	RefreshedDate int64 `json:"refreshed_date"`
	ContentType any `json:"content_type"`
	UserAllowAskDiscount int64 `json:"user_allow_ask_discount"`
	BuyerAvatarDate int64 `json:"buyer_avatar_date"`
	CanAskDiscount bool `json:"canAskDiscount"`
	IsBirthdayToday bool `json:"isBirthdayToday"`
	Guarantee map[string]any `json:"guarantee"`
	CanViewItemViews bool `json:"canViewItemViews"`
	CategoryId int64 `json:"category_id"`
	IsSticky int64 `json:"is_sticky"`
	PriceWithSellerFeeLabel string `json:"priceWithSellerFeeLabel"`
	DescriptionEnHtml string `json:"descriptionEnHtml"`
	DescriptionEn string `json:"description_en"`
	InformationEn string `json:"information_en"`
	EmailType string `json:"email_type"`
	BuyerUserGroupId int64 `json:"buyer_user_group_id"`
	IsPersonalAccount bool `json:"isPersonalAccount"`
	Deposit int64 `json:"deposit"`
	ItemId int64 `json:"item_id"`
	MaxDiscountPercent int64 `json:"max_discount_percent"`
	CanResellItemAfterPurchase bool `json:"canResellItemAfterPurchase"`
	CanCheckGuarantee bool `json:"canCheckGuarantee"`
	Title string `json:"title"`
	DeleteReason string `json:"delete_reason"`
	CanViewLoginData bool `json:"canViewLoginData"`
	CanReportItem bool `json:"canReportItem"`
	IsSmallExf bool `json:"isSmallExf"`
	ExternalAuth []any `json:"externalAuth"`
	DescriptionHtml string `json:"descriptionHtml"`
	CartPrice any `json:"cart_price"`
	CanValidateAccount bool `json:"canValidateAccount"`
	CanViewAccountLink bool `json:"canViewAccountLink"`
	AccountLinks []map[string]any `json:"accountLinks"`
	UniqueKeyExists bool `json:"uniqueKeyExists"`
	DescriptionPlain string `json:"descriptionPlain"`
	ResaleItemOrigin string `json:"resale_item_origin"`
	CopyFormatData map[string]any `json:"copyFormatData"`
	CustomFields map[string]any `json:"customFields"`
	ItemState string `json:"item_state"`
	Price int64 `json:"price"`
	TempEmail string `json:"temp_email"`
	FeedbackData string `json:"feedback_data"`
	IsFave any `json:"is_fave"`
	CanResellItem bool `json:"canResellItem"`
	ImagePreviewLinks []string `json:"imagePreviewLinks"`
	ItemOriginPhrase string `json:"itemOriginPhrase"`
	TitleEn string `json:"title_en"`
	BuyerDisplayIconGroupId int64 `json:"buyer_display_icon_group_id"`
	ShowGetEmailCodeButton bool `json:"showGetEmailCodeButton"`
	VisitorIsAuthor bool `json:"visitorIsAuthor"`
	AutoBuyPrice int64 `json:"autoBuyPrice"`
	AutoBuyPriceCheckDate int64 `json:"autoBuyPriceCheckDate"`
	PendingDeletionDate int64 `json:"pending_deletion_date"`
	AllowAskDiscount int64 `json:"allow_ask_discount"`
	EmailProvider string `json:"email_provider"`
	DeleteDate int64 `json:"delete_date"`
	DeleteUsername string `json:"delete_username"`
	BuyerUniqBanner string `json:"buyer_uniq_banner"`
	ExtraPrices []map[string]any `json:"extraPrices"`
	CanShareItem bool `json:"canShareItem"`
	PublishedDate int64 `json:"published_date"`
	UpdateStatDate int64 `json:"update_stat_date"`
	LoginData map[string]any `json:"loginData"`
	AccountLastActivity int64 `json:"account_last_activity"`
	IsIgnored bool `json:"isIgnored"`
	DescriptionEnPlain string `json:"descriptionEnPlain"`
	ExtendedGuarantee int64 `json:"extended_guarantee"`
	NoteText string `json:"note_text"`
	CanViewEmailLoginData bool `json:"canViewEmailLoginData"`
	CanChangeEmailPassword bool `json:"canChangeEmailPassword"`
	CanViewAccountLoginAndTempEmail bool `json:"canViewAccountLoginAndTempEmail"`
}

type CategoryGamesResponse struct {
	Games []map[string]any `json:"games,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ProfileClaimsResponse struct {
	Claims []map[string]any `json:"claims"`
	Stats map[string]any `json:"stats"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingCreateClaimResponse struct {
	Thread map[string]any `json:"thread"`
	SystemInfo map[string]any `json:"system_info"`
}

type CategoryGiftsResponse struct {
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategorySocialClubResponse struct {
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
}

type CategoryInstagramResponse struct {
	StickyItems []any `json:"stickyItems"`
	TotalItems int64 `json:"totalItems"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	CacheTTL int64 `json:"cacheTTL"`
}

type CategoryBattleNetResponse struct {
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	SearchUrl string `json:"searchUrl"`
}

type CategoryChatGPTResponse struct {
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
}

type ListStatesResponse struct {
	UserItemStates map[string]any `json:"userItemStates"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingGetLetters2Response struct {
	Email string `json:"email"`
	Letters []map[string]any `json:"letters"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryWotBlitzResponse struct {
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
}

type CategoryUplayResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
}

type CategoryTikTokResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	SystemInfo RespSystemInfo `json:"system_info"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
}

type ManagingSteamValueResponse struct {
	Query *string `json:"query,omitempty"`
	Data map[string]any `json:"data,omitempty"`
	AppId *int64 `json:"appId,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ManagingChangePasswordResponse struct {
	NewPassword string `json:"new_password"`
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

type CategoryRiotResponse struct {
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	ServerTime int64 `json:"serverTime"`
	Items []map[string]any `json:"items"`
}

type CategoryVpnResponse struct {
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
}

type ManagingAIPriceResponse struct {
	Price int64 `json:"price"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingSteamGetMafileResponse struct {
	MaFile map[string]any `json:"maFile"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryMihoyoResponse struct {
	Page int64 `json:"page"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
}

type PaymentsCurrencyResponse struct {
	CurrencyList map[string]any `json:"currencyList"`
	LastUpdate int64 `json:"lastUpdate"`
	VisitorCurrency string `json:"visitorCurrency"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategorySupercellResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
}

type ManagingGetResponse struct {
	CanCloseItem bool `json:"canCloseItem"`
	CanOpenItem bool `json:"canOpenItem"`
	CanEditItem bool `json:"canEditItem"`
	CanDeleteItem bool `json:"canDeleteItem"`
	CanViewItemHistory bool `json:"canViewItemHistory"`
	CanChangeOwner bool `json:"canChangeOwner"`
	Item ItemModel `json:"item"`
	CanBuyItem bool `json:"canBuyItem"`
	CanReportItem bool `json:"canReportItem"`
	SameItemsIds []int64 `json:"sameItemsIds"`
	FaveCount bool `json:"faveCount"`
	ShowToFavouritesButton bool `json:"showToFavouritesButton"`
	SameItemsCount int64 `json:"sameItemsCount"`
	SystemInfo RespSystemInfo `json:"system_info"`
	CanStickItem bool `json:"canStickItem"`
	CanUnstickItem bool `json:"canUnstickItem"`
	CannotBuyItemError string `json:"cannotBuyItemError"`
	CanCancelConfirmedBuy bool `json:"canCancelConfirmedBuy"`
	IsVisibleItem bool `json:"isVisibleItem"`
	CanViewLoginData bool `json:"canViewLoginData"`
	ItemLink string `json:"itemLink"`
}

type ManagingSteamInventoryValueResponse struct {
	Query *string `json:"query,omitempty"`
	Data map[string]any `json:"data,omitempty"`
	AppId *int64 `json:"appId,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ManagingImageResponse struct {
	Base64 string `json:"base64"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PaymentsInvoiceListResponse struct {
	Invoices []InvoiceModel `json:"invoices"`
	Count int64 `json:"count"`
	Page int64 `json:"page"`
	PerPage int64 `json:"perPage"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategorySteamResponse struct {
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	Items []map[string]any `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
}

type CategoryEAResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
}

type CategoryEpicGamesResponse struct {
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
}

type CategoryParamsResponse struct {
	Params []map[string]any `json:"params,omitempty"`
	BaseParams map[string]any `json:"base_params,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
	Category map[string]any `json:"category,omitempty"`
}

type ManagingAutoBuyPriceResponse struct {
	Price int64 `json:"price"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CartAddResponse struct {
	Success bool `json:"success"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CartDeleteResponse struct {
	Success bool `json:"success"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingTempEmailPasswordResponse struct {
	Item map[string]any `json:"item"`
}

type CategoryDiscordResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryWarfaceResponse struct {
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	PerPage int64 `json:"perPage"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
}

type CategoryListResponse struct {
	Category map[string]any `json:"category"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingBulkGetResponse struct {
	Items []map[string]any `json:"items"`
	LeftItemId []int64 `json:"left_item_id"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PublishingAddResponse struct {
	Status string `json:"status"`
	Item ItemModel `json:"item"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ProfileGetResponse struct {
	User UserModel `json:"user"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryTelegramResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
}

type CustomDiscountsEditResponse struct {
	Total int64 `json:"total"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Discounts []DiscountModel `json:"discounts"`
}

type CustomDiscountsGetResponse struct {
	Discounts []DiscountModel `json:"discounts"`
	Total int64 `json:"total"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CustomDiscountsCreateResponse struct {
	Discount DiscountModel `json:"discount"`
	Total int64 `json:"total"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PublishingFastSellResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Item ItemModel `json:"item"`
	ItemLink string `json:"itemLink"`
}

type AutoPaymentsListResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Payments map[string]any `json:"payments"`
}

type CategoryEscapeFromTarkovResponse struct {
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
	SearchUrl string `json:"searchUrl"`
	Items []map[string]any `json:"items"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
}

type ManagingSteamUpdateValueResponse struct {
	Status string `json:"status"`
	Item ItemModel `json:"item"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PurchasingConfirmResponse struct {
	Status *string `json:"status,omitempty"`
	Item map[string]any `json:"item"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingTelegramCodeResponse struct {
	Item ItemModel `json:"item"`
	Codes map[string]any `json:"codes"`
}

type PaymentsFeeResponse struct {
	CommissionPercentage int64 `json:"commission_percentage"`
	SpentCurrentMonth int64 `json:"spentCurrentMonth"`
	Calculator map[string]any `json:"calculator"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryWotResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryMinecraftResponse struct {
	Items []map[string]any `json:"items"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	CacheTTL int64 `json:"cacheTTL"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type BatchResponse struct {
	Jobs map[string]any `json:"jobs"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type CategoryRobloxResponse struct {
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PaymentsHistoryResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	PerPage string `json:"perPage"`
	PageNavLink string `json:"pageNavLink"`
	PageNavParams map[string]any `json:"pageNavParams"`
	PeriodLabel string `json:"periodLabel"`
	PaymentStats any `json:"paymentStats"`
	HasNextPage bool `json:"hasNextPage"`
	LastOperationId int64 `json:"lastOperationId"`
	NextPageHref string `json:"nextPageHref"`
	Payments map[string]any `json:"payments"`
	Page int64 `json:"page"`
	PeriodLabelPhrase string `json:"periodLabelPhrase"`
	FilterDatesDefault bool `json:"filterDatesDefault"`
	Input map[string]any `json:"input"`
}

type AutoPaymentsCreateResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	AutoPaymentId int64 `json:"auto_payment_id"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryFortniteResponse struct {
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
}

type CategoryHytaleResponse struct {
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
}

type ManagingCheckGuaranteeResponse struct {
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PaymentsInvoiceGetResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Invoice InvoiceModel `json:"invoice"`
}

type PaymentsInvoiceCreateResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Invoice InvoiceModel `json:"invoice"`
}

type PaymentsPayoutServicesResponse struct {
	Systems []map[string]any `json:"systems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ProxyGetResponse struct {
	Proxies []map[string]any `json:"proxies"`
	SystemInfo RespSystemInfo `json:"system_info"`
}
