package market

type ItemFromListModel struct {
	CanDeleteItem *bool `json:"canDeleteItem,omitempty"`
	CanBumpItem *bool `json:"canBumpItem,omitempty"`
	PriceCurrency *string `json:"price_currency,omitempty"`
	CanResellItemAfterPurchase *bool `json:"canResellItemAfterPurchase,omitempty"`
	CanViewAccountLink *bool `json:"canViewAccountLink,omitempty"`
	DescriptionHtml *string `json:"description_html,omitempty"`
	Description *string `json:"description,omitempty"`
	Nsb *int64 `json:"nsb,omitempty"`
	CanStickItem *bool `json:"canStickItem,omitempty"`
	CanBuyItem *bool `json:"canBuyItem,omitempty"`
	CanValidateAccount *bool `json:"canValidateAccount,omitempty"`
	ItemOriginPhrase *string `json:"itemOriginPhrase,omitempty"`
	NoteText *string `json:"note_text,omitempty"`
	ViewCount *int64 `json:"view_count,omitempty"`
	IsSticky *int64 `json:"is_sticky,omitempty"`
	CanCloseItem *bool `json:"canCloseItem,omitempty"`
	CanEditItem *bool `json:"canEditItem,omitempty"`
	BumpSettings map[string]any `json:"bumpSettings,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Seller map[string]any `json:"seller,omitempty"`
	RefreshedDate *int64 `json:"refreshed_date,omitempty"`
	ResaleItemOrigin *string `json:"resale_item_origin,omitempty"`
	Title *string `json:"title,omitempty"`
	ItemDomain *string `json:"item_domain,omitempty"`
	Guarantee *bool `json:"guarantee,omitempty"`
	CanViewEmailLoginData *bool `json:"canViewEmailLoginData,omitempty"`
	RubPrice *int64 `json:"rub_price,omitempty"`
	DescriptionHtmlEn *string `json:"description_html_en,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	PublishedDate *int64 `json:"published_date,omitempty"`
	UpdateStatDate *int64 `json:"update_stat_date,omitempty"`
	ItemOrigin *string `json:"item_origin,omitempty"`
	ExtendedGuarantee *int64 `json:"extended_guarantee,omitempty"`
	TitleEn *string `json:"title_en,omitempty"`
	ShowGetEmailCodeButton *bool `json:"showGetEmailCodeButton,omitempty"`
	Price *int64 `json:"price,omitempty"`
	DescriptionEn *string `json:"description_en,omitempty"`
	CanUnstickItem *bool `json:"canUnstickItem,omitempty"`
	ItemId *int64 `json:"item_id,omitempty"`
	CategoryId *int64 `json:"category_id,omitempty"`
	AllowAskDiscount *int64 `json:"allow_ask_discount,omitempty"`
	IsIgnored *int64 `json:"isIgnored,omitempty"`
	CanViewLoginData *bool `json:"canViewLoginData,omitempty"`
	CanUpdateItemStats *bool `json:"canUpdateItemStats,omitempty"`
	CanOpenItem *bool `json:"canOpenItem,omitempty"`
}

type ItemModel struct {
	DescriptionPlain string `json:"descriptionPlain"`
	CanViewEmailLoginData bool `json:"canViewEmailLoginData"`
	ViewCount int64 `json:"view_count"`
	Information string `json:"information"`
	ContentId any `json:"content_id"`
	DeleteDate int64 `json:"delete_date"`
	MaxDiscountPercent int64 `json:"max_discount_percent"`
	InCart any `json:"in_cart"`
	Guarantee map[string]any `json:"guarantee"`
	PublishedDate int64 `json:"published_date"`
	PendingDeletionDate int64 `json:"pending_deletion_date"`
	IsSticky int64 `json:"is_sticky"`
	ExtendedGuarantee int64 `json:"extended_guarantee"`
	FeedbackData string `json:"feedback_data"`
	CanViewLoginData bool `json:"canViewLoginData"`
	CanViewItemViews bool `json:"canViewItemViews"`
	IsPersonalAccount bool `json:"isPersonalAccount"`
	CopyFormatData map[string]any `json:"copyFormatData"`
	UpdateStatDate int64 `json:"update_stat_date"`
	EmailType string `json:"email_type"`
	BuyerDisplayIconGroupId int64 `json:"buyer_display_icon_group_id"`
	AccountLink string `json:"accountLink"`
	CanReportItem bool `json:"canReportItem"`
	RubPrice int64 `json:"rub_price"`
	IsSmallExf bool `json:"isSmallExf"`
	AccountLastActivity int64 `json:"account_last_activity"`
	CanViewAccountLink bool `json:"canViewAccountLink"`
	IsTrusted bool `json:"isTrusted"`
	BumpSettings map[string]any `json:"bumpSettings"`
	CanCheckAutoBuyPrice bool `json:"canCheckAutoBuyPrice"`
	EditDate int64 `json:"edit_date"`
	UserAllowAskDiscount int64 `json:"user_allow_ask_discount"`
	CustomFields map[string]any `json:"customFields"`
	AutoBuyPrice int64 `json:"autoBuyPrice"`
	DescriptionEnHtml string `json:"descriptionEnHtml"`
	DeleteUserId int64 `json:"delete_user_id"`
	AccountLinks []map[string]any `json:"accountLinks"`
	VisitorIsAuthor bool `json:"visitorIsAuthor"`
	ExternalAuth []any `json:"externalAuth"`
	RefreshedDate int64 `json:"refreshed_date"`
	DescriptionEn string `json:"description_en"`
	NoteText string `json:"note_text"`
	CanResellItem bool `json:"canResellItem"`
	LoginData map[string]any `json:"loginData"`
	Buyer map[string]any `json:"buyer"`
	CanViewAccountLoginAndTempEmail bool `json:"canViewAccountLoginAndTempEmail"`
	AutoBuyPriceCheckDate int64 `json:"autoBuyPriceCheckDate"`
	DeleteReason string `json:"delete_reason"`
	CanChangeEmailPassword bool `json:"canChangeEmailPassword"`
	UniqueKeyExists bool `json:"uniqueKeyExists"`
	IsIgnored bool `json:"isIgnored"`
	CanCheckGuarantee bool `json:"canCheckGuarantee"`
	AiPrice int64 `json:"aiPrice"`
	DescriptionEnPlain string `json:"descriptionEnPlain"`
	Title string `json:"title"`
	EmailProvider string `json:"email_provider"`
	ContentType any `json:"content_type"`
	DeleteUsername string `json:"delete_username"`
	GetEmailCodeDisplayLogin any `json:"getEmailCodeDisplayLogin"`
	ImagePreviewLinks []string `json:"imagePreviewLinks"`
	Tags map[string]any `json:"tags"`
	TitleEn string `json:"title_en"`
	IsFave any `json:"is_fave"`
	PriceWithSellerFeeLabel string `json:"priceWithSellerFeeLabel"`
	ExtraPrices []map[string]any `json:"extraPrices"`
	AiPriceCheckDate int64 `json:"aiPriceCheckDate"`
	DescriptionHtml string `json:"descriptionHtml"`
	ShowGetEmailCodeButton bool `json:"showGetEmailCodeButton"`
	TempEmail string `json:"temp_email"`
	MarketCustomTitle string `json:"market_custom_title"`
	BuyerUniqBanner string `json:"buyer_uniq_banner"`
	CartPrice any `json:"cart_price"`
	CanAskDiscount bool `json:"canAskDiscount"`
	IsBirthdayToday bool `json:"isBirthdayToday"`
	CanShareItem bool `json:"canShareItem"`
	ItemState string `json:"item_state"`
	Login string `json:"login"`
	Nsb int64 `json:"nsb"`
	InformationEn string `json:"information_en"`
	BuyerAvatarDate int64 `json:"buyer_avatar_date"`
	Deposit int64 `json:"deposit"`
	Seller map[string]any `json:"seller"`
	PriceWithSellerFee float64 `json:"priceWithSellerFee"`
	PriceCurrency string `json:"price_currency"`
	ItemId int64 `json:"item_id"`
	CategoryId int64 `json:"category_id"`
	Price int64 `json:"price"`
	ResaleItemOrigin string `json:"resale_item_origin"`
	CanValidateAccount bool `json:"canValidateAccount"`
	CanCheckAiPrice bool `json:"canCheckAiPrice"`
	Description string `json:"description"`
	ItemOrigin string `json:"item_origin"`
	CanResellItemAfterPurchase bool `json:"canResellItemAfterPurchase"`
	CanChangePassword bool `json:"canChangePassword"`
	NeedToRequireVideoToViewLoginData bool `json:"needToRequireVideoToViewLoginData"`
	AllowAskDiscount int64 `json:"allow_ask_discount"`
	ItemDomain string `json:"item_domain"`
	BuyerUserGroupId int64 `json:"buyer_user_group_id"`
	CanUpdateItemStats bool `json:"canUpdateItemStats"`
	ItemOriginPhrase string `json:"itemOriginPhrase"`
}

type UserModel struct {
	ContestCount int64 `json:"contest_count"`
	Dob map[string]any `json:"dob"`
	Homepage string `json:"homepage"`
	IsSuperAdmin bool `json:"is_super_admin"`
	Like2Count int64 `json:"like2_count"`
	RegisterDate int64 `json:"register_date"`
	RestoreCount int64 `json:"restore_count"`
	ShortLink string `json:"short_link"`
	ActivityVisible bool `json:"activity_visible"`
	Balance string `json:"balance"`
	CanEdit bool `json:"can_edit"`
	Deposit int64 `json:"deposit"`
	Location string `json:"location"`
	ConvWelcomeMessage string `json:"conv_welcome_message"`
	IsIgnored bool `json:"is_ignored"`
	PaidMailLeft int64 `json:"paid_mail_left"`
	Username string `json:"username"`
	IsFollowed bool `json:"is_followed"`
	CanViewProfile bool `json:"can_view_profile"`
	IsModerator bool `json:"is_moderator"`
	MarketCustomTitle string `json:"market_custom_title"`
	PublicTags []map[string]any `json:"public_tags"`
	Tags []map[string]any `json:"tags"`
	Age int64 `json:"age"`
	ConvertedDeposit int64 `json:"convertedDeposit"`
	Hold string `json:"hold"`
	UserAllowAskDiscount bool `json:"user_allow_ask_discount"`
	WarningPoints int64 `json:"warning_points"`
	ConvertedBalance int64 `json:"convertedBalance"`
	CustomFields map[string]any `json:"custom_fields"`
	ImapData map[string]any `json:"imap_data"`
	JoinedDate int64 `json:"joined_date"`
	MaxDiscountPercent int64 `json:"max_discount_percent"`
	TrophyPoints int64 `json:"trophy_points"`
	ViewUrl string `json:"view_url"`
	Currency string `json:"currency"`
	IsBanned bool `json:"is_banned"`
	RestoreData map[string]any `json:"restore_data"`
	Visible bool `json:"visible"`
	CurrencyPhrase string `json:"currencyPhrase"`
	CustomTitle string `json:"custom_title"`
	FeedbackData map[string]any `json:"feedback_data"`
	UserTitle string `json:"user_title"`
	BumpItemPeriod int64 `json:"bump_item_period"`
	IsStaff bool `json:"is_staff"`
	LikeCount int64 `json:"like_count"`
	CanPostProfile bool `json:"can_post_profile"`
	CustomAccountDownloadFormat string `json:"custom_account_download_format"`
	UserId int64 `json:"user_id"`
	ActiveItemsCount int64 `json:"active_items_count"`
	Balances []map[string]any `json:"balances"`
	ConvertedHold int64 `json:"convertedHold"`
	CanViewProfilePosts bool `json:"can_view_profile_posts"`
	IsAdmin bool `json:"is_admin"`
	MessageCount int64 `json:"message_count"`
	Rendered map[string]any `json:"rendered"`
	TelegramClient map[string]any `json:"telegram_client"`
	LastActivity int64 `json:"last_activity"`
	CanIgnore bool `json:"can_ignore"`
	CanWarn bool `json:"can_warn"`
	CanFollow bool `json:"can_follow"`
	SoldItemsCount int64 `json:"sold_items_count"`
}

type InvoiceModel struct {
	UrlCallback string `json:"url_callback"`
	ExpiresAt int64 `json:"expires_at"`
	PayerUserId int64 `json:"payer_user_id"`
	AdditionalData string `json:"additional_data"`
	Amount int64 `json:"amount"`
	Comment string `json:"comment"`
	InvoiceDate int64 `json:"invoice_date"`
	InvoiceId int64 `json:"invoice_id"`
	IsTest bool `json:"is_test"`
	MerchantId int64 `json:"merchant_id"`
	UrlSuccess string `json:"url_success"`
	PaidDate int64 `json:"paid_date"`
	PaymentId string `json:"payment_id"`
	ResendAttempts int64 `json:"resend_attempts"`
	Url string `json:"url"`
	UserId int64 `json:"user_id"`
	Status string `json:"status"`
}

type DiscountModel struct {
	DiscountId int64 `json:"discount_id"`
	DiscountPercent int64 `json:"discount_percent"`
	DiscountUserId int64 `json:"discount_user_id"`
	MaxPrice int64 `json:"max_price"`
	MinPrice int64 `json:"min_price"`
	UserId int64 `json:"user_id"`
	CategoryId int64 `json:"category_id"`
}

type ConfirmationCodeModel struct {
	Item ItemModel `json:"item"`
	CodeData map[string]any `json:"codeData"`
}

type ExtraModel struct {
	LoginWithoutCookies *bool `json:"login_without_cookies,omitempty"`
	ZenlessCurrency *int64 `json:"zenless_currency,omitempty"`
	CheckChannels *bool `json:"checkChannels,omitempty"`
	CheckSpam *bool `json:"checkSpam,omitempty"`
	CheckHypixelBan *bool `json:"checkHypixelBan,omitempty"`
	Proxy *string `json:"proxy,omitempty"`
	CloseItem *bool `json:"close_item,omitempty"`
	Service *string `json:"service,omitempty"`
	System *string `json:"system,omitempty"`
	TheQuarry *bool `json:"the_quarry,omitempty"`
	Ark *bool `json:"ark,omitempty"`
	HonkaiCurrency *int64 `json:"honkai_currency,omitempty"`
	TelegramClient *string `json:"telegramClient,omitempty"`
	ConfirmationCode *string `json:"confirmationCode,omitempty"`
	CookieLogin *bool `json:"cookie_login,omitempty"`
	EaGames *bool `json:"ea_games,omitempty"`
	UplayGames *bool `json:"uplay_games,omitempty"`
	Warframe *bool `json:"warframe,omitempty"`
	ArkAscended *bool `json:"ark_ascended,omitempty"`
	GenshinCurrency *int64 `json:"genshin_currency,omitempty"`
	TelegramJson *string `json:"telegramJson,omitempty"`
	Region *string `json:"region,omitempty"`
	MfaFile *string `json:"mfa_file,omitempty"`
	Dota2Mmr *int64 `json:"dota2_mmr,omitempty"`
	Password *string `json:"password,omitempty"`
	Cookies *string `json:"cookies,omitempty"`
}

type ItemListModel struct {
	Items []ItemFromListModel `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []ItemFromListModel `json:"stickyItems"`
	Page int64 `json:"page"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type RespSystemInfo struct {
	VisitorId int64 `json:"visitor_id"`
	Time int64 `json:"time"`
	LogId int64 `json:"log_id"`
}

type BalanceModel struct {
	UserId int64 `json:"user_id"`
	Balance string `json:"balance"`
	BalanceId int64 `json:"balance_id"`
	CustomTitle any `json:"custom_title"`
	FullTitle string `json:"fullTitle"`
	MerchantId int64 `json:"merchant_id"`
	Title string `json:"title"`
	TypeValue string `json:"type"`
}

type CategoryUplayResponse struct {
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryParamsResponse struct {
	Category map[string]any `json:"category,omitempty"`
	Params []map[string]any `json:"params,omitempty"`
	BaseParams map[string]any `json:"base_params,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ProfileGetResponse struct {
	User UserModel `json:"user"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryEscapeFromTarkovResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingGetResponse struct {
	IsVisibleItem bool `json:"isVisibleItem"`
	CanViewLoginData bool `json:"canViewLoginData"`
	CanChangeOwner bool `json:"canChangeOwner"`
	Item ItemModel `json:"item"`
	CanBuyItem bool `json:"canBuyItem"`
	CanDeleteItem bool `json:"canDeleteItem"`
	FaveCount bool `json:"faveCount"`
	ItemLink string `json:"itemLink"`
	CanStickItem bool `json:"canStickItem"`
	CanUnstickItem bool `json:"canUnstickItem"`
	CannotBuyItemError string `json:"cannotBuyItemError"`
	CanOpenItem bool `json:"canOpenItem"`
	CanCancelConfirmedBuy bool `json:"canCancelConfirmedBuy"`
	CanViewItemHistory bool `json:"canViewItemHistory"`
	ShowToFavouritesButton bool `json:"showToFavouritesButton"`
	SameItemsIds []int64 `json:"sameItemsIds"`
	SameItemsCount int64 `json:"sameItemsCount"`
	SystemInfo RespSystemInfo `json:"system_info"`
	CanCloseItem bool `json:"canCloseItem"`
	CanReportItem bool `json:"canReportItem"`
	CanEditItem bool `json:"canEditItem"`
}

type CartAddResponse struct {
	Success bool `json:"success"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CartDeleteResponse struct {
	Success bool `json:"success"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type AutoPaymentsListResponse struct {
	Payments map[string]any `json:"payments"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategorySupercellResponse struct {
	Items []map[string]any `json:"items"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
}

type CategoryWotBlitzResponse struct {
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	SystemInfo RespSystemInfo `json:"system_info"`
	HasNextPage bool `json:"hasNextPage"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
}

type CategorySocialClubResponse struct {
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
}

type CategoryWarfaceResponse struct {
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	StickyItems []any `json:"stickyItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
}

type ProfileClaimsResponse struct {
	Claims []map[string]any `json:"claims"`
	Stats map[string]any `json:"stats"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingCreateClaimResponse struct {
	SystemInfo map[string]any `json:"system_info"`
	Thread map[string]any `json:"thread"`
}

type PurchasingConfirmResponse struct {
	Status *string `json:"status,omitempty"`
	Item map[string]any `json:"item"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryEAResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PublishingFastSellResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Item ItemModel `json:"item"`
	ItemLink string `json:"itemLink"`
}

type PaymentsPayoutServicesResponse struct {
	Systems []map[string]any `json:"systems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryWotResponse struct {
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
}

type CategoryGiftsResponse struct {
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
}

type CategoryDiscordResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
}

type CategoryInstagramResponse struct {
	TotalItemsPrice any `json:"totalItemsPrice"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
}

type ManagingAutoBuyPriceResponse struct {
	Price int64 `json:"price"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingTelegramCodeResponse struct {
	Codes map[string]any `json:"codes"`
	Item ItemModel `json:"item"`
}

type ManagingTempEmailPasswordResponse struct {
	Item map[string]any `json:"item"`
}

type CategorySteamResponse struct {
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	Items []map[string]any `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
}

type CategoryRiotResponse struct {
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	SearchUrl string `json:"searchUrl"`
	Items []map[string]any `json:"items"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryMinecraftResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	Page int64 `json:"page"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	StickyItems []any `json:"stickyItems"`
}

type ManagingSteamValueResponse struct {
	Query *string `json:"query,omitempty"`
	Data map[string]any `json:"data,omitempty"`
	AppId *int64 `json:"appId,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type PaymentsCurrencyResponse struct {
	LastUpdate int64 `json:"lastUpdate"`
	VisitorCurrency string `json:"visitorCurrency"`
	SystemInfo RespSystemInfo `json:"system_info"`
	CurrencyList map[string]any `json:"currencyList"`
}

type CategoryMihoyoResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	HasNextPage bool `json:"hasNextPage"`
	WasCached bool `json:"wasCached"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryTelegramResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
}

type CategoryListResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Category map[string]any `json:"category"`
}

type PaymentsHistoryResponse struct {
	PeriodLabelPhrase string `json:"periodLabelPhrase"`
	FilterDatesDefault bool `json:"filterDatesDefault"`
	Input map[string]any `json:"input"`
	Page int64 `json:"page"`
	PaymentStats any `json:"paymentStats"`
	HasNextPage bool `json:"hasNextPage"`
	LastOperationId int64 `json:"lastOperationId"`
	NextPageHref string `json:"nextPageHref"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Payments map[string]any `json:"payments"`
	PerPage string `json:"perPage"`
	PageNavLink string `json:"pageNavLink"`
	PageNavParams map[string]any `json:"pageNavParams"`
	PeriodLabel string `json:"periodLabel"`
}

type BatchResponse struct {
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
	Jobs map[string]any `json:"jobs"`
}

type ManagingBulkGetResponse struct {
	LeftItemId []int64 `json:"left_item_id"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
}

type PaymentsFeeResponse struct {
	CommissionPercentage int64 `json:"commission_percentage"`
	SpentCurrentMonth int64 `json:"spentCurrentMonth"`
	Calculator map[string]any `json:"calculator"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type AutoPaymentsCreateResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	AutoPaymentId int64 `json:"auto_payment_id"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingSteamInventoryValueResponse struct {
	Query *string `json:"query,omitempty"`
	Data map[string]any `json:"data,omitempty"`
	AppId *int64 `json:"appId,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type CustomDiscountsCreateResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Discount DiscountModel `json:"discount"`
	Total int64 `json:"total"`
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

type PaymentsInvoiceGetResponse struct {
	Invoice InvoiceModel `json:"invoice"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PaymentsInvoiceCreateResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Invoice InvoiceModel `json:"invoice"`
}

type CategoryFortniteResponse struct {
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
}

type CategoryChatGPTResponse struct {
	TotalItems int64 `json:"totalItems"`
	Page int64 `json:"page"`
	LastModified int64 `json:"lastModified"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
}

type PublishingAddResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	Status string `json:"status"`
	Item ItemModel `json:"item"`
}

type ManagingCheckGuaranteeResponse struct {
	Message string `json:"message"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryEpicGamesResponse struct {
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	SearchUrl string `json:"searchUrl"`
	HasNextPage bool `json:"hasNextPage"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
}

type CategoryTikTokResponse struct {
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryBattleNetResponse struct {
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	PerPage int64 `json:"perPage"`
}

type CategoryRobloxResponse struct {
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	CacheTTL int64 `json:"cacheTTL"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	TotalItems int64 `json:"totalItems"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
	Items []map[string]any `json:"items"`
}

type CategoryGamesResponse struct {
	Games []map[string]any `json:"games,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ListStatesResponse struct {
	SystemInfo RespSystemInfo `json:"system_info"`
	UserItemStates map[string]any `json:"userItemStates"`
}

type ManagingImageResponse struct {
	Base64 string `json:"base64"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingGetLetters2Response struct {
	Email string `json:"email"`
	Letters []map[string]any `json:"letters"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryVpnResponse struct {
	Items []map[string]any `json:"items"`
	TotalItems int64 `json:"totalItems"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	LastModified int64 `json:"lastModified"`
	ServerTime int64 `json:"serverTime"`
	SearchUrl string `json:"searchUrl"`
	Page int64 `json:"page"`
	WasCached bool `json:"wasCached"`
	CacheTTL int64 `json:"cacheTTL"`
	StickyItems []any `json:"stickyItems"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type CategoryHytaleResponse struct {
	LastModified int64 `json:"lastModified"`
	SearchUrl string `json:"searchUrl"`
	StickyItems []any `json:"stickyItems"`
	Items []map[string]any `json:"items"`
	TotalItemsPrice any `json:"totalItemsPrice"`
	WasCached bool `json:"wasCached"`
	ServerTime int64 `json:"serverTime"`
	SystemInfo RespSystemInfo `json:"system_info"`
	TotalItems int64 `json:"totalItems"`
	HasNextPage bool `json:"hasNextPage"`
	PerPage int64 `json:"perPage"`
	Page int64 `json:"page"`
	CacheTTL int64 `json:"cacheTTL"`
}

type ManagingAIPriceResponse struct {
	Price int64 `json:"price"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type PaymentsInvoiceListResponse struct {
	Invoices []InvoiceModel `json:"invoices"`
	Count int64 `json:"count"`
	Page int64 `json:"page"`
	PerPage int64 `json:"perPage"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ProxyGetResponse struct {
	Proxies []map[string]any `json:"proxies"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingSteamUpdateValueResponse struct {
	Status string `json:"status"`
	Item ItemModel `json:"item"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingSteamGetMafileResponse struct {
	MaFile map[string]any `json:"maFile"`
	SystemInfo RespSystemInfo `json:"system_info"`
}

type ManagingChangePasswordResponse struct {
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	NewPassword string `json:"new_password"`
}
