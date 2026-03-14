package market

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	lztapi "github.com/fantolla123/lzt-api-go"
)

type Client struct {
	*lztapi.Client
}

func NewClient(token string, opts ...lztapi.Option) *Client {
	return &Client{
		Client: lztapi.NewClient(token, "https://prod-api.lzt.market", opts...),
	}
}

func (c *Client) CategoryGames(categoryName string, opts ...func(*CategoryGamesParams)) (*CategoryGamesResponse, error) {
	p := &CategoryGamesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{categoryName}/games"
	path = strings.Replace(path, "{categoryName}", fmt.Sprintf("%v", categoryName), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryGamesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryGamesParams struct {
}

func (c *Client) ListUser(opts ...func(*ListUserParams)) (*ItemListModel, error) {
	p := &ListUserParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/user/items"

	params := url.Values{}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.CategoryId != nil {
		params.Set("category_id", fmt.Sprintf("%v", *p.CategoryId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Show != nil {
		params.Set("show", fmt.Sprintf("%v", *p.Show))
	}
	if p.DeleteReason != nil {
		params.Set("delete_reason", fmt.Sprintf("%v", *p.DeleteReason))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Login != nil {
		params.Set("login", fmt.Sprintf("%v", *p.Login))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.Username != nil {
		params.Set("username", fmt.Sprintf("%v", *p.Username))
	}
	if p.PublishedStartDate != nil {
		params.Set("published_startDate", fmt.Sprintf("%v", *p.PublishedStartDate))
	}
	if p.PublishedEndDate != nil {
		params.Set("published_endDate", fmt.Sprintf("%v", *p.PublishedEndDate))
	}
	if p.FilterByPublishedDate != nil {
		params.Set("filter_by_published_date", fmt.Sprintf("%v", *p.FilterByPublishedDate))
	}
	if p.PaidStartDate != nil {
		params.Set("paid_startDate", fmt.Sprintf("%v", *p.PaidStartDate))
	}
	if p.PaidEndDate != nil {
		params.Set("paid_endDate", fmt.Sprintf("%v", *p.PaidEndDate))
	}
	if p.FilterByBuyerOperationDate != nil {
		params.Set("filter_by_buyer_operation_date", fmt.Sprintf("%v", *p.FilterByBuyerOperationDate))
	}
	if p.DeleteStartDate != nil {
		params.Set("delete_startDate", fmt.Sprintf("%v", *p.DeleteStartDate))
	}
	if p.DeleteEndDate != nil {
		params.Set("delete_endDate", fmt.Sprintf("%v", *p.DeleteEndDate))
	}
	if p.FilterByDeleteDate != nil {
		params.Set("filter_by_delete_date", fmt.Sprintf("%v", *p.FilterByDeleteDate))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ListUserParams struct {
	UserId *int64
	CategoryId *int64
	Page *int64
	Show *string
	DeleteReason *string
	Title *string
	Pmin *int64
	Pmax *int64
	Login *string
	Origin *[]string
	NotOrigin *[]string
	OrderBy *string
	Sb *bool
	SbByMe *bool
	Nsb *bool
	NsbByMe *bool
	Username *string
	PublishedStartDate *string
	PublishedEndDate *string
	FilterByPublishedDate *bool
	PaidStartDate *string
	PaidEndDate *string
	FilterByBuyerOperationDate *bool
	DeleteStartDate *string
	DeleteEndDate *string
	FilterByDeleteDate *bool
}

func (c *Client) ListFavorites(opts ...func(*ListFavoritesParams)) (*ItemListModel, error) {
	p := &ListFavoritesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/fave"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Show != nil {
		params.Set("show", fmt.Sprintf("%v", *p.Show))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ListFavoritesParams struct {
	Page *int64
	Show *string
	Title *string
	Pmin *int64
	Pmax *int64
	Origin *[]string
	NotOrigin *[]string
	OrderBy *string
	Sb *bool
	SbByMe *bool
	Nsb *bool
	NsbByMe *bool
}

func (c *Client) ProfileClaims(opts ...func(*ProfileClaimsParams)) (*ProfileClaimsResponse, error) {
	p := &ProfileClaimsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/claims"

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

	var result ProfileClaimsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ProfileClaimsParams struct {
	TypeValue *string
	ClaimState *string
}

func (c *Client) ManagingCreateClaim(itemId int64, postBody string, opts ...func(*ManagingCreateClaimParams)) (*ManagingCreateClaimResponse, error) {
	p := &ManagingCreateClaimParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/claims"


	bodyMap := map[string]any{}
	bodyMap["item_id"] = itemId
	bodyMap["post_body"] = postBody
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ManagingCreateClaimResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingCreateClaimParams struct {
}

func (c *Client) ManagingSteamSDA(itemId int64, opts ...func(*ManagingSteamSDAParams)) (map[string]any, error) {
	p := &ManagingSteamSDAParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/confirm-sda"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Id != nil {
		bodyMap["id"] = *p.Id
	}
	if p.Nonce != nil {
		bodyMap["nonce"] = *p.Nonce
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

type ManagingSteamSDAParams struct {
	Id *int64
	Nonce *int64
}

func (c *Client) CategoryGifts(opts ...func(*CategoryGiftsParams)) (*CategoryGiftsResponse, error) {
	p := &CategoryGiftsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/gifts"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Subscription != nil {
		params.Set("subscription", fmt.Sprintf("%v", *p.Subscription))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryGiftsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryGiftsParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Subscription *string
	SubscriptionLength *int64
	SubscriptionPeriod *string
}

func (c *Client) CategorySocialClub(opts ...func(*CategorySocialClubParams)) (*CategorySocialClubResponse, error) {
	p := &CategorySocialClubParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/socialclub"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.LevelMin != nil {
		params.Set("level_min", fmt.Sprintf("%v", *p.LevelMin))
	}
	if p.LevelMax != nil {
		params.Set("level_max", fmt.Sprintf("%v", *p.LevelMax))
	}
	if p.CashMin != nil {
		params.Set("cash_min", fmt.Sprintf("%v", *p.CashMin))
	}
	if p.CashMax != nil {
		params.Set("cash_max", fmt.Sprintf("%v", *p.CashMax))
	}
	if p.BankCashMin != nil {
		params.Set("bank_cash_min", fmt.Sprintf("%v", *p.BankCashMin))
	}
	if p.BankCashMax != nil {
		params.Set("bank_cash_max", fmt.Sprintf("%v", *p.BankCashMax))
	}
	if p.Game != nil {
		params.Set("game[]", fmt.Sprintf("%v", *p.Game))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategorySocialClubResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategorySocialClubParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Daybreak *int64
	LevelMin *int64
	LevelMax *int64
	CashMin *int64
	CashMax *int64
	BankCashMin *int64
	BankCashMax *int64
	Game *[]string
}

func (c *Client) CategoryInstagram(opts ...func(*CategoryInstagramParams)) (*CategoryInstagramResponse, error) {
	p := &CategoryInstagramParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/instagram"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Cookies != nil {
		params.Set("cookies", fmt.Sprintf("%v", *p.Cookies))
	}
	if p.LoginWithoutCookies != nil {
		params.Set("login_without_cookies", fmt.Sprintf("%v", *p.LoginWithoutCookies))
	}
	if p.FollowersMin != nil {
		params.Set("followers_min", fmt.Sprintf("%v", *p.FollowersMin))
	}
	if p.FollowersMax != nil {
		params.Set("followers_max", fmt.Sprintf("%v", *p.FollowersMax))
	}
	if p.PostMin != nil {
		params.Set("post_min", fmt.Sprintf("%v", *p.PostMin))
	}
	if p.PostMax != nil {
		params.Set("post_max", fmt.Sprintf("%v", *p.PostMax))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryInstagramResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryInstagramParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Tel *string
	Country *[]string
	NotCountry *[]string
	Cookies *string
	LoginWithoutCookies *string
	FollowersMin *int64
	FollowersMax *int64
	PostMin *int64
	PostMax *int64
	Reg *int64
	RegPeriod *string
}

func (c *Client) CategoryBattleNet(opts ...func(*CategoryBattleNetParams)) (*CategoryBattleNetResponse, error) {
	p := &CategoryBattleNetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/battlenet"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Eg != nil {
		params.Set("eg", fmt.Sprintf("%v", *p.Eg))
	}
	if p.Game != nil {
		params.Set("game[]", fmt.Sprintf("%v", *p.Game))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.EditBtag != nil {
		params.Set("edit_btag", fmt.Sprintf("%v", *p.EditBtag))
	}
	if p.ChangeableFn != nil {
		params.Set("changeable_fn", fmt.Sprintf("%v", *p.ChangeableFn))
	}
	if p.RealId != nil {
		params.Set("real_id", fmt.Sprintf("%v", *p.RealId))
	}
	if p.ParentControl != nil {
		params.Set("parent_control", fmt.Sprintf("%v", *p.ParentControl))
	}
	if p.NoBans != nil {
		params.Set("no_bans", fmt.Sprintf("%v", *p.NoBans))
	}
	if p.BalanceMin != nil {
		params.Set("balance_min", fmt.Sprintf("%v", *p.BalanceMin))
	}
	if p.BalanceMax != nil {
		params.Set("balance_max", fmt.Sprintf("%v", *p.BalanceMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryBattleNetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryBattleNetParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Eg *int64
	Game *[]int64
	Daybreak *int64
	Country *[]string
	NotCountry *[]string
	Tel *string
	EditBtag *string
	ChangeableFn *string
	RealId *string
	ParentControl *string
	NoBans *string
	BalanceMin *int64
	BalanceMax *int64
}

func (c *Client) CategoryChatGPT(opts ...func(*CategoryChatGPTParams)) (*CategoryChatGPTResponse, error) {
	p := &CategoryChatGPTParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/chatgpt"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Subscription != nil {
		params.Set("subscription[]", fmt.Sprintf("%v", *p.Subscription))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}
	if p.Autorenewal != nil {
		params.Set("autorenewal", fmt.Sprintf("%v", *p.Autorenewal))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Transactions != nil {
		params.Set("transactions", fmt.Sprintf("%v", *p.Transactions))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.OpenaiTier != nil {
		params.Set("openai_tier[]", fmt.Sprintf("%v", *p.OpenaiTier))
	}
	if p.OpenaiBalanceMin != nil {
		params.Set("openai_balance_min", fmt.Sprintf("%v", *p.OpenaiBalanceMin))
	}
	if p.OpenaiBalanceMax != nil {
		params.Set("openai_balance_max", fmt.Sprintf("%v", *p.OpenaiBalanceMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryChatGPTResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryChatGPTParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Subscription *[]string
	SubscriptionLength *int64
	SubscriptionPeriod *string
	Autorenewal *string
	Tel *string
	Transactions *string
	Reg *int64
	RegPeriod *string
	OpenaiTier *[]string
	OpenaiBalanceMin *int64
	OpenaiBalanceMax *int64
}

func (c *Client) ListStates(opts ...func(*ListStatesParams)) (*ListStatesResponse, error) {
	p := &ListStatesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/user/item-states"

	params := url.Values{}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ListStatesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ListStatesParams struct {
	UserId *int64
}

func (c *Client) ManagingEmailCode(itemId int64, opts ...func(*ManagingEmailCodeParams)) (*ConfirmationCodeModel, error) {
	p := &ManagingEmailCodeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/email-code"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConfirmationCodeModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingEmailCodeParams struct {
}

func (c *Client) ManagingGetLetters2(opts ...func(*ManagingGetLetters2Params)) (*ManagingGetLetters2Response, error) {
	p := &ManagingGetLetters2Params{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/letters2"

	params := url.Values{}
	if p.EmailPassword != nil {
		params.Set("email_password", fmt.Sprintf("%v", *p.EmailPassword))
	}
	if p.Email != nil {
		params.Set("email", fmt.Sprintf("%v", *p.Email))
	}
	if p.Password != nil {
		params.Set("password", fmt.Sprintf("%v", *p.Password))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingGetLetters2Response
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingGetLetters2Params struct {
	EmailPassword *string
	Email *string
	Password *string
	Limit *int64
}

func (c *Client) CategoryWotBlitz(opts ...func(*CategoryWotBlitzParams)) (*CategoryWotBlitzResponse, error) {
	p := &CategoryWotBlitzParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/wot-blitz"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.BattlesMin != nil {
		params.Set("battles_min", fmt.Sprintf("%v", *p.BattlesMin))
	}
	if p.BattlesMax != nil {
		params.Set("battles_max", fmt.Sprintf("%v", *p.BattlesMax))
	}
	if p.GoldMin != nil {
		params.Set("gold_min", fmt.Sprintf("%v", *p.GoldMin))
	}
	if p.GoldMax != nil {
		params.Set("gold_max", fmt.Sprintf("%v", *p.GoldMax))
	}
	if p.SilverMin != nil {
		params.Set("silver_min", fmt.Sprintf("%v", *p.SilverMin))
	}
	if p.SilverMax != nil {
		params.Set("silver_max", fmt.Sprintf("%v", *p.SilverMax))
	}
	if p.TopMin != nil {
		params.Set("top_min", fmt.Sprintf("%v", *p.TopMin))
	}
	if p.TopMax != nil {
		params.Set("top_max", fmt.Sprintf("%v", *p.TopMax))
	}
	if p.PremMin != nil {
		params.Set("prem_min", fmt.Sprintf("%v", *p.PremMin))
	}
	if p.PremMax != nil {
		params.Set("prem_max", fmt.Sprintf("%v", *p.PremMax))
	}
	if p.TopPremMin != nil {
		params.Set("top_prem_min", fmt.Sprintf("%v", *p.TopPremMin))
	}
	if p.TopPremMax != nil {
		params.Set("top_prem_max", fmt.Sprintf("%v", *p.TopPremMax))
	}
	if p.WinPmin != nil {
		params.Set("win_pmin", fmt.Sprintf("%v", *p.WinPmin))
	}
	if p.WinPmax != nil {
		params.Set("win_pmax", fmt.Sprintf("%v", *p.WinPmax))
	}
	if p.Tank != nil {
		params.Set("tank[]", fmt.Sprintf("%v", *p.Tank))
	}
	if p.Region != nil {
		params.Set("region[]", fmt.Sprintf("%v", *p.Region))
	}
	if p.NotRegion != nil {
		params.Set("not_region[]", fmt.Sprintf("%v", *p.NotRegion))
	}
	if p.Premium != nil {
		params.Set("premium", fmt.Sprintf("%v", *p.Premium))
	}
	if p.PremiumExpiration != nil {
		params.Set("premium_expiration", fmt.Sprintf("%v", *p.PremiumExpiration))
	}
	if p.PremiumExpirationPeriod != nil {
		params.Set("premium_expiration_period", fmt.Sprintf("%v", *p.PremiumExpirationPeriod))
	}
	if p.Clan != nil {
		params.Set("clan", fmt.Sprintf("%v", *p.Clan))
	}
	if p.ClanRole != nil {
		params.Set("clan_role[]", fmt.Sprintf("%v", *p.ClanRole))
	}
	if p.NotClanRole != nil {
		params.Set("not_clan_role[]", fmt.Sprintf("%v", *p.NotClanRole))
	}
	if p.ClanGoldMin != nil {
		params.Set("clan_gold_min", fmt.Sprintf("%v", *p.ClanGoldMin))
	}
	if p.ClanGoldMax != nil {
		params.Set("clan_gold_max", fmt.Sprintf("%v", *p.ClanGoldMax))
	}
	if p.ClanCreditsMin != nil {
		params.Set("clan_credits_min", fmt.Sprintf("%v", *p.ClanCreditsMin))
	}
	if p.ClanCreditsMax != nil {
		params.Set("clan_credits_max", fmt.Sprintf("%v", *p.ClanCreditsMax))
	}
	if p.ClanCrystalMin != nil {
		params.Set("clan_crystal_min", fmt.Sprintf("%v", *p.ClanCrystalMin))
	}
	if p.ClanCrystalMax != nil {
		params.Set("clan_crystal_max", fmt.Sprintf("%v", *p.ClanCrystalMax))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryWotBlitzResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryWotBlitzParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Tel *string
	Daybreak *int64
	BattlesMin *int64
	BattlesMax *int64
	GoldMin *int64
	GoldMax *int64
	SilverMin *int64
	SilverMax *int64
	TopMin *int64
	TopMax *int64
	PremMin *int64
	PremMax *int64
	TopPremMin *int64
	TopPremMax *int64
	WinPmin *int64
	WinPmax *int64
	Tank *[]int64
	Region *[]string
	NotRegion *[]string
	Premium *string
	PremiumExpiration *int64
	PremiumExpirationPeriod *string
	Clan *string
	ClanRole *[]string
	NotClanRole *[]string
	ClanGoldMin *int64
	ClanGoldMax *int64
	ClanCreditsMin *int64
	ClanCreditsMax *int64
	ClanCrystalMin *int64
	ClanCrystalMax *int64
	Country *[]string
	NotCountry *[]string
}

func (c *Client) CategoryUplay(opts ...func(*CategoryUplayParams)) (*CategoryUplayResponse, error) {
	p := &CategoryUplayParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/uplay"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Game != nil {
		params.Set("game[]", fmt.Sprintf("%v", *p.Game))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.Gmin != nil {
		params.Set("gmin", fmt.Sprintf("%v", *p.Gmin))
	}
	if p.Gmax != nil {
		params.Set("gmax", fmt.Sprintf("%v", *p.Gmax))
	}
	if p.Subscription != nil {
		params.Set("subscription", fmt.Sprintf("%v", *p.Subscription))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}
	if p.R6LevelMin != nil {
		params.Set("r6_level_min", fmt.Sprintf("%v", *p.R6LevelMin))
	}
	if p.R6LevelMax != nil {
		params.Set("r6_level_max", fmt.Sprintf("%v", *p.R6LevelMax))
	}
	if p.R6RankMin != nil {
		params.Set("r6_rank_min", fmt.Sprintf("%v", *p.R6RankMin))
	}
	if p.R6RankMax != nil {
		params.Set("r6_rank_max", fmt.Sprintf("%v", *p.R6RankMax))
	}
	if p.R6OperatorsMin != nil {
		params.Set("r6_operators_min", fmt.Sprintf("%v", *p.R6OperatorsMin))
	}
	if p.R6OperatorsMax != nil {
		params.Set("r6_operators_max", fmt.Sprintf("%v", *p.R6OperatorsMax))
	}
	if p.R6Ban != nil {
		params.Set("r6_ban", fmt.Sprintf("%v", *p.R6Ban))
	}
	if p.R6Smin != nil {
		params.Set("r6_smin", fmt.Sprintf("%v", *p.R6Smin))
	}
	if p.R6Smax != nil {
		params.Set("r6_smax", fmt.Sprintf("%v", *p.R6Smax))
	}
	if p.R6Skin != nil {
		params.Set("r6_skin[]", fmt.Sprintf("%v", *p.R6Skin))
	}
	if p.R6Operator != nil {
		params.Set("r6_operator[]", fmt.Sprintf("%v", *p.R6Operator))
	}
	if p.XboxConnected != nil {
		params.Set("xbox_connected", fmt.Sprintf("%v", *p.XboxConnected))
	}
	if p.PsnConnected != nil {
		params.Set("psn_connected", fmt.Sprintf("%v", *p.PsnConnected))
	}
	if p.SteamConnected != nil {
		params.Set("steam_connected", fmt.Sprintf("%v", *p.SteamConnected))
	}
	if p.BalanceMin != nil {
		params.Set("balance_min", fmt.Sprintf("%v", *p.BalanceMin))
	}
	if p.BalanceMax != nil {
		params.Set("balance_max", fmt.Sprintf("%v", *p.BalanceMax))
	}
	if p.Transactions != nil {
		params.Set("transactions", fmt.Sprintf("%v", *p.Transactions))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryUplayResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryUplayParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Game *[]string
	Country *[]string
	NotCountry *[]string
	Daybreak *int64
	Gmin *int64
	Gmax *int64
	Subscription *string
	SubscriptionLength *int64
	SubscriptionPeriod *string
	R6LevelMin *int64
	R6LevelMax *int64
	R6RankMin *int64
	R6RankMax *int64
	R6OperatorsMin *int64
	R6OperatorsMax *int64
	R6Ban *string
	R6Smin *int64
	R6Smax *int64
	R6Skin *[]string
	R6Operator *[]string
	XboxConnected *string
	PsnConnected *string
	SteamConnected *string
	BalanceMin *float64
	BalanceMax *float64
	Transactions *string
	Reg *int64
	RegPeriod *string
}

func (c *Client) CategoryTikTok(opts ...func(*CategoryTikTokParams)) (*CategoryTikTokResponse, error) {
	p := &CategoryTikTokParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/tiktok"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.FollowersMin != nil {
		params.Set("followers_min", fmt.Sprintf("%v", *p.FollowersMin))
	}
	if p.FollowersMax != nil {
		params.Set("followers_max", fmt.Sprintf("%v", *p.FollowersMax))
	}
	if p.PostMin != nil {
		params.Set("post_min", fmt.Sprintf("%v", *p.PostMin))
	}
	if p.PostMax != nil {
		params.Set("post_max", fmt.Sprintf("%v", *p.PostMax))
	}
	if p.LikeMin != nil {
		params.Set("like_min", fmt.Sprintf("%v", *p.LikeMin))
	}
	if p.LikeMax != nil {
		params.Set("like_max", fmt.Sprintf("%v", *p.LikeMax))
	}
	if p.CoinsMin != nil {
		params.Set("coins_min", fmt.Sprintf("%v", *p.CoinsMin))
	}
	if p.CoinsMax != nil {
		params.Set("coins_max", fmt.Sprintf("%v", *p.CoinsMax))
	}
	if p.CookieLogin != nil {
		params.Set("cookie_login", fmt.Sprintf("%v", *p.CookieLogin))
	}
	if p.Verified != nil {
		params.Set("verified", fmt.Sprintf("%v", *p.Verified))
	}
	if p.Email != nil {
		params.Set("email", fmt.Sprintf("%v", *p.Email))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryTikTokResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryTikTokParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Tel *string
	Reg *int64
	RegPeriod *string
	FollowersMin *int64
	FollowersMax *int64
	PostMin *int64
	PostMax *int64
	LikeMin *int64
	LikeMax *int64
	CoinsMin *int64
	CoinsMax *int64
	CookieLogin *string
	Verified *string
	Email *string
}

func (c *Client) ManagingSteamValue(link string, opts ...func(*ManagingSteamValueParams)) (*ManagingSteamValueResponse, error) {
	p := &ManagingSteamValueParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/steam-value"

	params := url.Values{}
	params.Set("link", fmt.Sprintf("%v", link))
	if p.AppId != nil {
		params.Set("app_id", fmt.Sprintf("%v", *p.AppId))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.IgnoreCache != nil {
		params.Set("ignore_cache", fmt.Sprintf("%v", *p.IgnoreCache))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingSteamValueResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingSteamValueParams struct {
	AppId *int64
	Currency *string
	IgnoreCache *bool
}

func (c *Client) PublishingCheck(itemId int64, opts ...func(*PublishingCheckParams)) (map[string]any, error) {
	p := &PublishingCheckParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/goods/check"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.RandomProxy != nil {
		bodyMap["random_proxy"] = *p.RandomProxy
	}
	if p.Login != nil {
		bodyMap["login"] = *p.Login
	}
	if p.Password != nil {
		bodyMap["password"] = *p.Password
	}
	if p.LoginPassword != nil {
		bodyMap["login_password"] = *p.LoginPassword
	}
	if p.EmailLoginData != nil {
		bodyMap["email_login_data"] = *p.EmailLoginData
	}
	if p.ResellItemId != nil {
		bodyMap["resell_item_id"] = *p.ResellItemId
	}
	if p.HasEmailLoginData != nil {
		bodyMap["has_email_login_data"] = *p.HasEmailLoginData
	}
	if p.EmailType != nil {
		bodyMap["email_type"] = *p.EmailType
	}
	if p.Extra != nil {
		bodyMap["extra"] = *p.Extra
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

type PublishingCheckParams struct {
	RandomProxy *bool
	Login *string
	Password *string
	LoginPassword *string
	EmailLoginData *string
	ResellItemId *int64
	HasEmailLoginData *bool
	EmailType *string
	Extra *ExtraModel
}

func (c *Client) ManagingDeclineVideoRecording(itemId int64, iVoluntarilyAndWithFullAwarenessOfMyActionsWaiveAnyClaimsRegardingThisItem bool, opts ...func(*ManagingDeclineVideoRecordingParams)) (map[string]any, error) {
	p := &ManagingDeclineVideoRecordingParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/decline-video-recording"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["i_voluntarily_and_with_full_awareness_of_my_actions_waive_any_claims_regarding_this_item"] = iVoluntarilyAndWithFullAwarenessOfMyActionsWaiveAnyClaimsRegardingThisItem
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

type ManagingDeclineVideoRecordingParams struct {
}

func (c *Client) ManagingChangePassword(itemId int64, opts ...func(*ManagingChangePasswordParams)) (*ManagingChangePasswordResponse, error) {
	p := &ManagingChangePasswordParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/change-password"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Cancel != nil {
		bodyMap["_cancel"] = *p.Cancel
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ManagingChangePasswordResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingChangePasswordParams struct {
	Cancel *int64
}

func (c *Client) PaymentsBalanceList(opts ...func(*PaymentsBalanceListParams)) (map[string]any, error) {
	p := &PaymentsBalanceListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/exchange"


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

type PaymentsBalanceListParams struct {
}

func (c *Client) PaymentsBalanceExchange(fromBalance string, toBalance string, amount int64, opts ...func(*PaymentsBalanceExchangeParams)) (map[string]any, error) {
	p := &PaymentsBalanceExchangeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/exchange"


	bodyMap := map[string]any{}
	bodyMap["from_balance"] = fromBalance
	bodyMap["to_balance"] = toBalance
	bodyMap["amount"] = amount
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

type PaymentsBalanceExchangeParams struct {
}

func (c *Client) CategoryRiot(opts ...func(*CategoryRiotParams)) (*CategoryRiotResponse, error) {
	p := &CategoryRiotParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/riot"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Rmin != nil {
		params.Set("rmin", fmt.Sprintf("%v", *p.Rmin))
	}
	if p.Rmax != nil {
		params.Set("rmax", fmt.Sprintf("%v", *p.Rmax))
	}
	if p.LastRmin != nil {
		params.Set("last_rmin", fmt.Sprintf("%v", *p.LastRmin))
	}
	if p.LastRmax != nil {
		params.Set("last_rmax", fmt.Sprintf("%v", *p.LastRmax))
	}
	if p.PreviousRmin != nil {
		params.Set("previous_rmin", fmt.Sprintf("%v", *p.PreviousRmin))
	}
	if p.PreviousRmax != nil {
		params.Set("previous_rmax", fmt.Sprintf("%v", *p.PreviousRmax))
	}
	if p.WeaponSkin != nil {
		params.Set("weaponSkin[]", fmt.Sprintf("%v", *p.WeaponSkin))
	}
	if p.Buddy != nil {
		params.Set("buddy[]", fmt.Sprintf("%v", *p.Buddy))
	}
	if p.Agent != nil {
		params.Set("agent[]", fmt.Sprintf("%v", *p.Agent))
	}
	if p.Champion != nil {
		params.Set("champion[]", fmt.Sprintf("%v", *p.Champion))
	}
	if p.Skin != nil {
		params.Set("skin[]", fmt.Sprintf("%v", *p.Skin))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.ValorantLevelMin != nil {
		params.Set("valorant_level_min", fmt.Sprintf("%v", *p.ValorantLevelMin))
	}
	if p.ValorantLevelMax != nil {
		params.Set("valorant_level_max", fmt.Sprintf("%v", *p.ValorantLevelMax))
	}
	if p.LolLevelMin != nil {
		params.Set("lol_level_min", fmt.Sprintf("%v", *p.LolLevelMin))
	}
	if p.LolLevelMax != nil {
		params.Set("lol_level_max", fmt.Sprintf("%v", *p.LolLevelMax))
	}
	if p.InvMin != nil {
		params.Set("inv_min", fmt.Sprintf("%v", *p.InvMin))
	}
	if p.InvMax != nil {
		params.Set("inv_max", fmt.Sprintf("%v", *p.InvMax))
	}
	if p.VpMin != nil {
		params.Set("vp_min", fmt.Sprintf("%v", *p.VpMin))
	}
	if p.VpMax != nil {
		params.Set("vp_max", fmt.Sprintf("%v", *p.VpMax))
	}
	if p.ValorantSmin != nil {
		params.Set("valorant_smin", fmt.Sprintf("%v", *p.ValorantSmin))
	}
	if p.ValorantSmax != nil {
		params.Set("valorant_smax", fmt.Sprintf("%v", *p.ValorantSmax))
	}
	if p.ValorantRankType != nil {
		params.Set("valorant_rank_type[]", fmt.Sprintf("%v", *p.ValorantRankType))
	}
	if p.Amin != nil {
		params.Set("amin", fmt.Sprintf("%v", *p.Amin))
	}
	if p.Amax != nil {
		params.Set("amax", fmt.Sprintf("%v", *p.Amax))
	}
	if p.ValorantRegion != nil {
		params.Set("valorant_region[]", fmt.Sprintf("%v", *p.ValorantRegion))
	}
	if p.ValorantNotRegion != nil {
		params.Set("valorant_not_region[]", fmt.Sprintf("%v", *p.ValorantNotRegion))
	}
	if p.LolRegion != nil {
		params.Set("lol_region[]", fmt.Sprintf("%v", *p.LolRegion))
	}
	if p.LolNotRegion != nil {
		params.Set("lol_not_region[]", fmt.Sprintf("%v", *p.LolNotRegion))
	}
	if p.Knife != nil {
		params.Set("knife", fmt.Sprintf("%v", *p.Knife))
	}
	if p.LolSmin != nil {
		params.Set("lol_smin", fmt.Sprintf("%v", *p.LolSmin))
	}
	if p.LolSmax != nil {
		params.Set("lol_smax", fmt.Sprintf("%v", *p.LolSmax))
	}
	if p.ChampionMin != nil {
		params.Set("champion_min", fmt.Sprintf("%v", *p.ChampionMin))
	}
	if p.ChampionMax != nil {
		params.Set("champion_max", fmt.Sprintf("%v", *p.ChampionMax))
	}
	if p.WinRateMin != nil {
		params.Set("win_rate_min", fmt.Sprintf("%v", *p.WinRateMin))
	}
	if p.WinRateMax != nil {
		params.Set("win_rate_max", fmt.Sprintf("%v", *p.WinRateMax))
	}
	if p.BlueMin != nil {
		params.Set("blue_min", fmt.Sprintf("%v", *p.BlueMin))
	}
	if p.BlueMax != nil {
		params.Set("blue_max", fmt.Sprintf("%v", *p.BlueMax))
	}
	if p.OrangeMin != nil {
		params.Set("orange_min", fmt.Sprintf("%v", *p.OrangeMin))
	}
	if p.OrangeMax != nil {
		params.Set("orange_max", fmt.Sprintf("%v", *p.OrangeMax))
	}
	if p.MythicMin != nil {
		params.Set("mythic_min", fmt.Sprintf("%v", *p.MythicMin))
	}
	if p.MythicMax != nil {
		params.Set("mythic_max", fmt.Sprintf("%v", *p.MythicMax))
	}
	if p.RiotMin != nil {
		params.Set("riot_min", fmt.Sprintf("%v", *p.RiotMin))
	}
	if p.RiotMax != nil {
		params.Set("riot_max", fmt.Sprintf("%v", *p.RiotMax))
	}
	if p.Email != nil {
		params.Set("email", fmt.Sprintf("%v", *p.Email))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.ValorantKnifeMin != nil {
		params.Set("valorant_knife_min", fmt.Sprintf("%v", *p.ValorantKnifeMin))
	}
	if p.ValorantKnifeMax != nil {
		params.Set("valorant_knife_max", fmt.Sprintf("%v", *p.ValorantKnifeMax))
	}
	if p.RpMin != nil {
		params.Set("rp_min", fmt.Sprintf("%v", *p.RpMin))
	}
	if p.RpMax != nil {
		params.Set("rp_max", fmt.Sprintf("%v", *p.RpMax))
	}
	if p.FaMin != nil {
		params.Set("fa_min", fmt.Sprintf("%v", *p.FaMin))
	}
	if p.FaMax != nil {
		params.Set("fa_max", fmt.Sprintf("%v", *p.FaMax))
	}
	if p.LolRank != nil {
		params.Set("lol_rank[]", fmt.Sprintf("%v", *p.LolRank))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryRiotResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryRiotParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Rmin *int64
	Rmax *int64
	LastRmin *int64
	LastRmax *int64
	PreviousRmin *int64
	PreviousRmax *int64
	WeaponSkin *[]string
	Buddy *[]string
	Agent *[]string
	Champion *[]string
	Skin *[]string
	Country *[]string
	NotCountry *[]string
	Daybreak *int64
	ValorantLevelMin *int64
	ValorantLevelMax *int64
	LolLevelMin *int64
	LolLevelMax *int64
	InvMin *int64
	InvMax *int64
	VpMin *int64
	VpMax *int64
	ValorantSmin *int64
	ValorantSmax *int64
	ValorantRankType *[]string
	Amin *int64
	Amax *int64
	ValorantRegion *[]string
	ValorantNotRegion *[]string
	LolRegion *[]string
	LolNotRegion *[]string
	Knife *bool
	LolSmin *int64
	LolSmax *int64
	ChampionMin *int64
	ChampionMax *int64
	WinRateMin *int64
	WinRateMax *int64
	BlueMin *int64
	BlueMax *int64
	OrangeMin *int64
	OrangeMax *int64
	MythicMin *int64
	MythicMax *int64
	RiotMin *int64
	RiotMax *int64
	Email *string
	Tel *string
	ValorantKnifeMin *int64
	ValorantKnifeMax *int64
	RpMin *int64
	RpMax *int64
	FaMin *int64
	FaMax *int64
	LolRank *[]string
}

func (c *Client) CategoryVpn(opts ...func(*CategoryVpnParams)) (*CategoryVpnResponse, error) {
	p := &CategoryVpnParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/vpn"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Service != nil {
		params.Set("service[]", fmt.Sprintf("%v", *p.Service))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}
	if p.Autorenewal != nil {
		params.Set("autorenewal", fmt.Sprintf("%v", *p.Autorenewal))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryVpnResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryVpnParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Service *[]string
	SubscriptionLength *int64
	SubscriptionPeriod *string
	Autorenewal *string
}

func (c *Client) ManagingSteamPreview(itemId int64, opts ...func(*ManagingSteamPreviewParams)) (map[string]any, error) {
	p := &ManagingSteamPreviewParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/steam-preview"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)

	params := url.Values{}
	if p.TypeValue != nil {
		params.Set("type", fmt.Sprintf("%v", *p.TypeValue))
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

type ManagingSteamPreviewParams struct {
	TypeValue *string
}

func (c *Client) ManagingAIPrice(itemId int64, opts ...func(*ManagingAIPriceParams)) (*ManagingAIPriceResponse, error) {
	p := &ManagingAIPriceParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/ai-price"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingAIPriceResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingAIPriceParams struct {
}

func (c *Client) PurchasingCheck(itemId int64, opts ...func(*PurchasingCheckParams)) (map[string]any, error) {
	p := &PurchasingCheckParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/check-account"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type PurchasingCheckParams struct {
}

func (c *Client) ManagingSteamGetMafile(itemId int64, opts ...func(*ManagingSteamGetMafileParams)) (*ManagingSteamGetMafileResponse, error) {
	p := &ManagingSteamGetMafileParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/mafile"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingSteamGetMafileResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingSteamGetMafileParams struct {
}

func (c *Client) ManagingSteamAddMafile(itemId int64, opts ...func(*ManagingSteamAddMafileParams)) (map[string]any, error) {
	p := &ManagingSteamAddMafileParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/mafile"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingSteamAddMafileParams struct {
}

func (c *Client) ManagingSteamRemoveMafile(itemId int64, opts ...func(*ManagingSteamRemoveMafileParams)) (map[string]any, error) {
	p := &ManagingSteamRemoveMafileParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/mafile"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingSteamRemoveMafileParams struct {
}

func (c *Client) ManagingTransfer(itemId int64, username string, secretAnswer string, opts ...func(*ManagingTransferParams)) (map[string]any, error) {
	p := &ManagingTransferParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/change-owner"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["username"] = username
	bodyMap["secret_answer"] = secretAnswer
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

type ManagingTransferParams struct {
}

func (c *Client) ImapCreate(port int64, secure bool, domain string, imapServer string, opts ...func(*ImapCreateParams)) (map[string]any, error) {
	p := &ImapCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/imap"


	bodyMap := map[string]any{}
	bodyMap["port"] = port
	bodyMap["secure"] = secure
	bodyMap["domain"] = domain
	bodyMap["imap_server"] = imapServer
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

type ImapCreateParams struct {
}

func (c *Client) ImapDelete(domain string, opts ...func(*ImapDeleteParams)) (map[string]any, error) {
	p := &ImapDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/imap"


	bodyMap := map[string]any{}
	bodyMap["domain"] = domain
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

type ImapDeleteParams struct {
}

func (c *Client) CategoryMihoyo(opts ...func(*CategoryMihoyoParams)) (*CategoryMihoyoResponse, error) {
	p := &CategoryMihoyoParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/mihoyo"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Email != nil {
		params.Set("email", fmt.Sprintf("%v", *p.Email))
	}
	if p.Ea != nil {
		params.Set("ea", fmt.Sprintf("%v", *p.Ea))
	}
	if p.Region != nil {
		params.Set("region", fmt.Sprintf("%v", *p.Region))
	}
	if p.NotRegion != nil {
		params.Set("not_region", fmt.Sprintf("%v", *p.NotRegion))
	}
	if p.GenshinCharacter != nil {
		params.Set("genshin_character[]", fmt.Sprintf("%v", *p.GenshinCharacter))
	}
	if p.GenshinCharacterConstellations != nil {
		params.Set("genshin_character_constellations", fmt.Sprintf("%v", *p.GenshinCharacterConstellations))
	}
	if p.GenshinCharacterConstellationsMax != nil {
		params.Set("genshin_character_constellations_max", fmt.Sprintf("%v", *p.GenshinCharacterConstellationsMax))
	}
	if p.GenshinWeapon != nil {
		params.Set("genshin_weapon[]", fmt.Sprintf("%v", *p.GenshinWeapon))
	}
	if p.GenshinCharMin != nil {
		params.Set("genshin_char_min", fmt.Sprintf("%v", *p.GenshinCharMin))
	}
	if p.GenshinCharMax != nil {
		params.Set("genshin_char_max", fmt.Sprintf("%v", *p.GenshinCharMax))
	}
	if p.GenshinLegendaryMin != nil {
		params.Set("genshin_legendary_min", fmt.Sprintf("%v", *p.GenshinLegendaryMin))
	}
	if p.GenshinLegendaryMax != nil {
		params.Set("genshin_legendary_max", fmt.Sprintf("%v", *p.GenshinLegendaryMax))
	}
	if p.GenshinLevelMin != nil {
		params.Set("genshin_level_min", fmt.Sprintf("%v", *p.GenshinLevelMin))
	}
	if p.GenshinLevelMax != nil {
		params.Set("genshin_level_max", fmt.Sprintf("%v", *p.GenshinLevelMax))
	}
	if p.GenshinLegendaryWeaponMin != nil {
		params.Set("genshin_legendary_weapon_min", fmt.Sprintf("%v", *p.GenshinLegendaryWeaponMin))
	}
	if p.GenshinLegendaryWeaponMax != nil {
		params.Set("genshin_legendary_weapon_max", fmt.Sprintf("%v", *p.GenshinLegendaryWeaponMax))
	}
	if p.ConstellationsMin != nil {
		params.Set("constellations_min", fmt.Sprintf("%v", *p.ConstellationsMin))
	}
	if p.ConstellationsMax != nil {
		params.Set("constellations_max", fmt.Sprintf("%v", *p.ConstellationsMax))
	}
	if p.GenshinAchievementMin != nil {
		params.Set("genshin_achievement_min", fmt.Sprintf("%v", *p.GenshinAchievementMin))
	}
	if p.GenshinAchievementMax != nil {
		params.Set("genshin_achievement_max", fmt.Sprintf("%v", *p.GenshinAchievementMax))
	}
	if p.GenshinCurrencyMin != nil {
		params.Set("genshin_currency_min", fmt.Sprintf("%v", *p.GenshinCurrencyMin))
	}
	if p.GenshinCurrencyMax != nil {
		params.Set("genshin_currency_max", fmt.Sprintf("%v", *p.GenshinCurrencyMax))
	}
	if p.HonkaiCharacter != nil {
		params.Set("honkai_character[]", fmt.Sprintf("%v", *p.HonkaiCharacter))
	}
	if p.HonkaiCharacterEidolons != nil {
		params.Set("honkai_character_eidolons", fmt.Sprintf("%v", *p.HonkaiCharacterEidolons))
	}
	if p.HonkaiCharacterEidolonsMax != nil {
		params.Set("honkai_character_eidolons_max", fmt.Sprintf("%v", *p.HonkaiCharacterEidolonsMax))
	}
	if p.HonkaiWeapon != nil {
		params.Set("honkai_weapon[]", fmt.Sprintf("%v", *p.HonkaiWeapon))
	}
	if p.HonkaiCharMin != nil {
		params.Set("honkai_char_min", fmt.Sprintf("%v", *p.HonkaiCharMin))
	}
	if p.HonkaiCharMax != nil {
		params.Set("honkai_char_max", fmt.Sprintf("%v", *p.HonkaiCharMax))
	}
	if p.HonkaiLegendaryMin != nil {
		params.Set("honkai_legendary_min", fmt.Sprintf("%v", *p.HonkaiLegendaryMin))
	}
	if p.HonkaiLegendaryMax != nil {
		params.Set("honkai_legendary_max", fmt.Sprintf("%v", *p.HonkaiLegendaryMax))
	}
	if p.HonkaiLevelMin != nil {
		params.Set("honkai_level_min", fmt.Sprintf("%v", *p.HonkaiLevelMin))
	}
	if p.HonkaiLevelMax != nil {
		params.Set("honkai_level_max", fmt.Sprintf("%v", *p.HonkaiLevelMax))
	}
	if p.HonkaiLegendaryWeaponMin != nil {
		params.Set("honkai_legendary_weapon_min", fmt.Sprintf("%v", *p.HonkaiLegendaryWeaponMin))
	}
	if p.HonkaiLegendaryWeaponMax != nil {
		params.Set("honkai_legendary_weapon_max", fmt.Sprintf("%v", *p.HonkaiLegendaryWeaponMax))
	}
	if p.EidolonsMin != nil {
		params.Set("eidolons_min", fmt.Sprintf("%v", *p.EidolonsMin))
	}
	if p.EidolonsMax != nil {
		params.Set("eidolons_max", fmt.Sprintf("%v", *p.EidolonsMax))
	}
	if p.HonkaiAchievementMin != nil {
		params.Set("honkai_achievement_min", fmt.Sprintf("%v", *p.HonkaiAchievementMin))
	}
	if p.HonkaiAchievementMax != nil {
		params.Set("honkai_achievement_max", fmt.Sprintf("%v", *p.HonkaiAchievementMax))
	}
	if p.HonkaiCurrencyMin != nil {
		params.Set("honkai_currency_min", fmt.Sprintf("%v", *p.HonkaiCurrencyMin))
	}
	if p.HonkaiCurrencyMax != nil {
		params.Set("honkai_currency_max", fmt.Sprintf("%v", *p.HonkaiCurrencyMax))
	}
	if p.ZenlessCharacter != nil {
		params.Set("zenless_character[]", fmt.Sprintf("%v", *p.ZenlessCharacter))
	}
	if p.ZenlessCharacterCinemas != nil {
		params.Set("zenless_character_cinemas", fmt.Sprintf("%v", *p.ZenlessCharacterCinemas))
	}
	if p.ZenlessCharacterCinemasMax != nil {
		params.Set("zenless_character_cinemas_max", fmt.Sprintf("%v", *p.ZenlessCharacterCinemasMax))
	}
	if p.ZenlessWeapon != nil {
		params.Set("zenless_weapon[]", fmt.Sprintf("%v", *p.ZenlessWeapon))
	}
	if p.ZenlessLegendaryMin != nil {
		params.Set("zenless_legendary_min", fmt.Sprintf("%v", *p.ZenlessLegendaryMin))
	}
	if p.ZenlessLegendaryMax != nil {
		params.Set("zenless_legendary_max", fmt.Sprintf("%v", *p.ZenlessLegendaryMax))
	}
	if p.CinemasMin != nil {
		params.Set("cinemas_min", fmt.Sprintf("%v", *p.CinemasMin))
	}
	if p.CinemasMax != nil {
		params.Set("cinemas_max", fmt.Sprintf("%v", *p.CinemasMax))
	}
	if p.ZenlessLegendaryWeaponMin != nil {
		params.Set("zenless_legendary_weapon_min", fmt.Sprintf("%v", *p.ZenlessLegendaryWeaponMin))
	}
	if p.ZenlessLegendaryWeaponMax != nil {
		params.Set("zenless_legendary_weapon_max", fmt.Sprintf("%v", *p.ZenlessLegendaryWeaponMax))
	}
	if p.ZenlessCharMin != nil {
		params.Set("zenless_char_min", fmt.Sprintf("%v", *p.ZenlessCharMin))
	}
	if p.ZenlessCharMax != nil {
		params.Set("zenless_char_max", fmt.Sprintf("%v", *p.ZenlessCharMax))
	}
	if p.ZenlessLevelMin != nil {
		params.Set("zenless_level_min", fmt.Sprintf("%v", *p.ZenlessLevelMin))
	}
	if p.ZenlessLevelMax != nil {
		params.Set("zenless_level_max", fmt.Sprintf("%v", *p.ZenlessLevelMax))
	}
	if p.ZenlessAchievementMin != nil {
		params.Set("zenless_achievement_min", fmt.Sprintf("%v", *p.ZenlessAchievementMin))
	}
	if p.ZenlessAchievementMax != nil {
		params.Set("zenless_achievement_max", fmt.Sprintf("%v", *p.ZenlessAchievementMax))
	}
	if p.ZenlessCurrencyMin != nil {
		params.Set("zenless_currency_min", fmt.Sprintf("%v", *p.ZenlessCurrencyMin))
	}
	if p.ZenlessCurrencyMax != nil {
		params.Set("zenless_currency_max", fmt.Sprintf("%v", *p.ZenlessCurrencyMax))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryMihoyoResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryMihoyoParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	EmailType *[]string
	ParseSameItemIds *bool
	ItemDomain *string
	Email *string
	Ea *string
	Region *[]string
	NotRegion *[]string
	GenshinCharacter *[]int64
	GenshinCharacterConstellations *map[string]any
	GenshinCharacterConstellationsMax *map[string]any
	GenshinWeapon *[]int64
	GenshinCharMin *int64
	GenshinCharMax *int64
	GenshinLegendaryMin *int64
	GenshinLegendaryMax *int64
	GenshinLevelMin *int64
	GenshinLevelMax *int64
	GenshinLegendaryWeaponMin *int64
	GenshinLegendaryWeaponMax *int64
	ConstellationsMin *int64
	ConstellationsMax *int64
	GenshinAchievementMin *int64
	GenshinAchievementMax *int64
	GenshinCurrencyMin *int64
	GenshinCurrencyMax *int64
	HonkaiCharacter *[]int64
	HonkaiCharacterEidolons *map[string]any
	HonkaiCharacterEidolonsMax *map[string]any
	HonkaiWeapon *[]int64
	HonkaiCharMin *int64
	HonkaiCharMax *int64
	HonkaiLegendaryMin *int64
	HonkaiLegendaryMax *int64
	HonkaiLevelMin *int64
	HonkaiLevelMax *int64
	HonkaiLegendaryWeaponMin *int64
	HonkaiLegendaryWeaponMax *int64
	EidolonsMin *int64
	EidolonsMax *int64
	HonkaiAchievementMin *int64
	HonkaiAchievementMax *int64
	HonkaiCurrencyMin *int64
	HonkaiCurrencyMax *int64
	ZenlessCharacter *[]int64
	ZenlessCharacterCinemas *map[string]any
	ZenlessCharacterCinemasMax *map[string]any
	ZenlessWeapon *[]int64
	ZenlessLegendaryMin *int64
	ZenlessLegendaryMax *int64
	CinemasMin *int64
	CinemasMax *int64
	ZenlessLegendaryWeaponMin *int64
	ZenlessLegendaryWeaponMax *int64
	ZenlessCharMin *int64
	ZenlessCharMax *int64
	ZenlessLevelMin *int64
	ZenlessLevelMax *int64
	ZenlessAchievementMin *int64
	ZenlessAchievementMax *int64
	ZenlessCurrencyMin *int64
	ZenlessCurrencyMax *int64
	Daybreak *int64
}

func (c *Client) ManagingEdit(itemId int64, opts ...func(*ManagingEditParams)) (map[string]any, error) {
	p := &ManagingEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/edit"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Currency != nil {
		bodyMap["currency"] = *p.Currency
	}
	if p.ItemOrigin != nil {
		bodyMap["item_origin"] = *p.ItemOrigin
	}
	if p.EmailLoginData != nil {
		bodyMap["email_login_data"] = *p.EmailLoginData
	}
	if p.EmailType != nil {
		bodyMap["email_type"] = *p.EmailType
	}
	if p.Description != nil {
		bodyMap["description"] = *p.Description
	}
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.Price != nil {
		bodyMap["price"] = *p.Price
	}
	if p.AllowAskDiscount != nil {
		bodyMap["allow_ask_discount"] = *p.AllowAskDiscount
	}
	if p.ProxyId != nil {
		bodyMap["proxy_id"] = *p.ProxyId
	}
	if p.Information != nil {
		bodyMap["information"] = *p.Information
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

type ManagingEditParams struct {
	Currency *string
	ItemOrigin *string
	EmailLoginData *string
	EmailType *string
	Description *string
	Title *string
	TitleEn *string
	Price *int64
	AllowAskDiscount *bool
	ProxyId *int64
	Information *string
}

func (c *Client) PaymentsCurrency(opts ...func(*PaymentsCurrencyParams)) (*PaymentsCurrencyResponse, error) {
	p := &PaymentsCurrencyParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/currency"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result PaymentsCurrencyResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsCurrencyParams struct {
}

func (c *Client) CategorySupercell(opts ...func(*CategorySupercellParams)) (*CategorySupercellResponse, error) {
	p := &CategorySupercellParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/supercell"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Eg != nil {
		params.Set("eg", fmt.Sprintf("%v", *p.Eg))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.BrawlLevelMin != nil {
		params.Set("brawl_level_min", fmt.Sprintf("%v", *p.BrawlLevelMin))
	}
	if p.BrawlLevelMax != nil {
		params.Set("brawl_level_max", fmt.Sprintf("%v", *p.BrawlLevelMax))
	}
	if p.BrawlCupMin != nil {
		params.Set("brawl_cup_min", fmt.Sprintf("%v", *p.BrawlCupMin))
	}
	if p.BrawlCupMax != nil {
		params.Set("brawl_cup_max", fmt.Sprintf("%v", *p.BrawlCupMax))
	}
	if p.BrawlWinsMin != nil {
		params.Set("brawl_wins_min", fmt.Sprintf("%v", *p.BrawlWinsMin))
	}
	if p.BrawlWinsMax != nil {
		params.Set("brawl_wins_max", fmt.Sprintf("%v", *p.BrawlWinsMax))
	}
	if p.BrawlPass != nil {
		params.Set("brawl_pass", fmt.Sprintf("%v", *p.BrawlPass))
	}
	if p.Brawler != nil {
		params.Set("brawler[]", fmt.Sprintf("%v", *p.Brawler))
	}
	if p.BrawlersMin != nil {
		params.Set("brawlers_min", fmt.Sprintf("%v", *p.BrawlersMin))
	}
	if p.BrawlersMax != nil {
		params.Set("brawlers_max", fmt.Sprintf("%v", *p.BrawlersMax))
	}
	if p.LegendaryBrawlersMin != nil {
		params.Set("legendary_brawlers_min", fmt.Sprintf("%v", *p.LegendaryBrawlersMin))
	}
	if p.LegendaryBrawlersMax != nil {
		params.Set("legendary_brawlers_max", fmt.Sprintf("%v", *p.LegendaryBrawlersMax))
	}
	if p.RoyaleLevelMin != nil {
		params.Set("royale_level_min", fmt.Sprintf("%v", *p.RoyaleLevelMin))
	}
	if p.RoyaleLevelMax != nil {
		params.Set("royale_level_max", fmt.Sprintf("%v", *p.RoyaleLevelMax))
	}
	if p.RoyaleCupMin != nil {
		params.Set("royale_cup_min", fmt.Sprintf("%v", *p.RoyaleCupMin))
	}
	if p.RoyaleCupMax != nil {
		params.Set("royale_cup_max", fmt.Sprintf("%v", *p.RoyaleCupMax))
	}
	if p.RoyaleWinsMin != nil {
		params.Set("royale_wins_min", fmt.Sprintf("%v", *p.RoyaleWinsMin))
	}
	if p.RoyaleWinsMax != nil {
		params.Set("royale_wins_max", fmt.Sprintf("%v", *p.RoyaleWinsMax))
	}
	if p.KingLevelMin != nil {
		params.Set("king_level_min", fmt.Sprintf("%v", *p.KingLevelMin))
	}
	if p.KingLevelMax != nil {
		params.Set("king_level_max", fmt.Sprintf("%v", *p.KingLevelMax))
	}
	if p.RoyalePass != nil {
		params.Set("royale_pass", fmt.Sprintf("%v", *p.RoyalePass))
	}
	if p.ClashLevelMin != nil {
		params.Set("clash_level_min", fmt.Sprintf("%v", *p.ClashLevelMin))
	}
	if p.ClashLevelMax != nil {
		params.Set("clash_level_max", fmt.Sprintf("%v", *p.ClashLevelMax))
	}
	if p.ClashCupMin != nil {
		params.Set("clash_cup_min", fmt.Sprintf("%v", *p.ClashCupMin))
	}
	if p.ClashCupMax != nil {
		params.Set("clash_cup_max", fmt.Sprintf("%v", *p.ClashCupMax))
	}
	if p.ClashWinsMin != nil {
		params.Set("clash_wins_min", fmt.Sprintf("%v", *p.ClashWinsMin))
	}
	if p.ClashWinsMax != nil {
		params.Set("clash_wins_max", fmt.Sprintf("%v", *p.ClashWinsMax))
	}
	if p.ClashPass != nil {
		params.Set("clash_pass", fmt.Sprintf("%v", *p.ClashPass))
	}
	if p.TotalHeroesLevelMin != nil {
		params.Set("total_heroes_level_min", fmt.Sprintf("%v", *p.TotalHeroesLevelMin))
	}
	if p.TotalHeroesLevelMax != nil {
		params.Set("total_heroes_level_max", fmt.Sprintf("%v", *p.TotalHeroesLevelMax))
	}
	if p.TotalTroopsLevelMin != nil {
		params.Set("total_troops_level_min", fmt.Sprintf("%v", *p.TotalTroopsLevelMin))
	}
	if p.TotalTroopsLevelMax != nil {
		params.Set("total_troops_level_max", fmt.Sprintf("%v", *p.TotalTroopsLevelMax))
	}
	if p.TotalSpellsLevelMin != nil {
		params.Set("total_spells_level_min", fmt.Sprintf("%v", *p.TotalSpellsLevelMin))
	}
	if p.TotalSpellsLevelMax != nil {
		params.Set("total_spells_level_max", fmt.Sprintf("%v", *p.TotalSpellsLevelMax))
	}
	if p.TotalBuilderHeroesLevelMin != nil {
		params.Set("total_builder_heroes_level_min", fmt.Sprintf("%v", *p.TotalBuilderHeroesLevelMin))
	}
	if p.TotalBuilderHeroesLevelMax != nil {
		params.Set("total_builder_heroes_level_max", fmt.Sprintf("%v", *p.TotalBuilderHeroesLevelMax))
	}
	if p.TotalBuilderTroopsLevelMin != nil {
		params.Set("total_builder_troops_level_min", fmt.Sprintf("%v", *p.TotalBuilderTroopsLevelMin))
	}
	if p.TotalBuilderTroopsLevelMax != nil {
		params.Set("total_builder_troops_level_max", fmt.Sprintf("%v", *p.TotalBuilderTroopsLevelMax))
	}
	if p.TownHallLevelMin != nil {
		params.Set("town_hall_level_min", fmt.Sprintf("%v", *p.TownHallLevelMin))
	}
	if p.TownHallLevelMax != nil {
		params.Set("town_hall_level_max", fmt.Sprintf("%v", *p.TownHallLevelMax))
	}
	if p.BuilderHallLevelMin != nil {
		params.Set("builder_hall_level_min", fmt.Sprintf("%v", *p.BuilderHallLevelMin))
	}
	if p.BuilderHallLevelMax != nil {
		params.Set("builder_hall_level_max", fmt.Sprintf("%v", *p.BuilderHallLevelMax))
	}
	if p.BuilderHallCupMin != nil {
		params.Set("builder_hall_cup_min", fmt.Sprintf("%v", *p.BuilderHallCupMin))
	}
	if p.BuilderHallCupMax != nil {
		params.Set("builder_hall_cup_max", fmt.Sprintf("%v", *p.BuilderHallCupMax))
	}
	if p.CreationYearMin != nil {
		params.Set("creation_year_min", fmt.Sprintf("%v", *p.CreationYearMin))
	}
	if p.CreationYearMax != nil {
		params.Set("creation_year_max", fmt.Sprintf("%v", *p.CreationYearMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategorySupercellResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategorySupercellParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Daybreak *int64
	EmailType *[]string
	ItemDomain *string
	Eg *int64
	Tel *string
	BrawlLevelMin *int64
	BrawlLevelMax *int64
	BrawlCupMin *int64
	BrawlCupMax *int64
	BrawlWinsMin *int64
	BrawlWinsMax *int64
	BrawlPass *string
	Brawler *[]string
	BrawlersMin *int64
	BrawlersMax *int64
	LegendaryBrawlersMin *int64
	LegendaryBrawlersMax *int64
	RoyaleLevelMin *int64
	RoyaleLevelMax *int64
	RoyaleCupMin *int64
	RoyaleCupMax *int64
	RoyaleWinsMin *int64
	RoyaleWinsMax *int64
	KingLevelMin *int64
	KingLevelMax *int64
	RoyalePass *string
	ClashLevelMin *int64
	ClashLevelMax *int64
	ClashCupMin *int64
	ClashCupMax *int64
	ClashWinsMin *int64
	ClashWinsMax *int64
	ClashPass *string
	TotalHeroesLevelMin *int64
	TotalHeroesLevelMax *int64
	TotalTroopsLevelMin *int64
	TotalTroopsLevelMax *int64
	TotalSpellsLevelMin *int64
	TotalSpellsLevelMax *int64
	TotalBuilderHeroesLevelMin *int64
	TotalBuilderHeroesLevelMax *int64
	TotalBuilderTroopsLevelMin *int64
	TotalBuilderTroopsLevelMax *int64
	TownHallLevelMin *int64
	TownHallLevelMax *int64
	BuilderHallLevelMin *int64
	BuilderHallLevelMax *int64
	BuilderHallCupMin *int64
	BuilderHallCupMax *int64
	CreationYearMin *int64
	CreationYearMax *int64
}

func (c *Client) ManagingGet(itemId int64, opts ...func(*ManagingGetParams)) (*ManagingGetResponse, error) {
	p := &ManagingGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)

	params := url.Values{}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingGetParams struct {
	ParseSameItemIds *bool
}

func (c *Client) MangingDelete(itemId int64, reason string, opts ...func(*MangingDeleteParams)) (map[string]any, error) {
	p := &MangingDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["reason"] = reason
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

type MangingDeleteParams struct {
}

func (c *Client) ManagingNote(itemId int64, opts ...func(*ManagingNoteParams)) (map[string]any, error) {
	p := &ManagingNoteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/note-save"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Text != nil {
		bodyMap["text"] = *p.Text
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

type ManagingNoteParams struct {
	Text *string
}

func (c *Client) ManagingBump(itemId int64, opts ...func(*ManagingBumpParams)) (map[string]any, error) {
	p := &ManagingBumpParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/bump"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingBumpParams struct {
}

func (c *Client) PurchasingFastBuy(itemId int64, opts ...func(*PurchasingFastBuyParams)) (map[string]any, error) {
	p := &PurchasingFastBuyParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/fast-buy"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Price != nil {
		bodyMap["price"] = *p.Price
	}
	if p.BalanceId != nil {
		bodyMap["balance_id"] = *p.BalanceId
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

type PurchasingFastBuyParams struct {
	Price *float64
	BalanceId *int64
}

func (c *Client) ManagingRefuseGuarantee(itemId int64, opts ...func(*ManagingRefuseGuaranteeParams)) (map[string]any, error) {
	p := &ManagingRefuseGuaranteeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/refuse-guarantee"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingRefuseGuaranteeParams struct {
}

func (c *Client) ManagingSteamInventoryValue(itemId int64, opts ...func(*ManagingSteamInventoryValueParams)) (*ManagingSteamInventoryValueResponse, error) {
	p := &ManagingSteamInventoryValueParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/inventory-value"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)

	params := url.Values{}
	if p.AppId != nil {
		params.Set("app_id", fmt.Sprintf("%v", *p.AppId))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.IgnoreCache != nil {
		params.Set("ignore_cache", fmt.Sprintf("%v", *p.IgnoreCache))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingSteamInventoryValueResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingSteamInventoryValueParams struct {
	AppId *int64
	Currency *string
	IgnoreCache *bool
}

func (c *Client) ManagingAutoBumpDisable(itemId int64, opts ...func(*ManagingAutoBumpDisableParams)) (map[string]any, error) {
	p := &ManagingAutoBumpDisableParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/auto-bump"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingAutoBumpDisableParams struct {
}

func (c *Client) ManagingAutoBump(itemId int64, hour int64, opts ...func(*ManagingAutoBumpParams)) (map[string]any, error) {
	p := &ManagingAutoBumpParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/auto-bump"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["hour"] = hour
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

type ManagingAutoBumpParams struct {
}

func (c *Client) ManagingImage(itemId int64, typeValue string, opts ...func(*ManagingImageParams)) (*ManagingImageResponse, error) {
	p := &ManagingImageParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/image"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)

	params := url.Values{}
	params.Set("type", fmt.Sprintf("%v", typeValue))

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingImageResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingImageParams struct {
}

func (c *Client) PaymentsInvoiceList(opts ...func(*PaymentsInvoiceListParams)) (*PaymentsInvoiceListResponse, error) {
	p := &PaymentsInvoiceListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/invoice/list"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.Status != nil {
		params.Set("status", fmt.Sprintf("%v", *p.Status))
	}
	if p.Amount != nil {
		params.Set("amount", fmt.Sprintf("%v", *p.Amount))
	}
	if p.MerchantId != nil {
		params.Set("merchant_id", fmt.Sprintf("%v", *p.MerchantId))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result PaymentsInvoiceListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsInvoiceListParams struct {
	Page *int64
	Currency *string
	Status *string
	Amount *float64
	MerchantId *int64
}

func (c *Client) CategorySteam(opts ...func(*CategorySteamParams)) (*CategorySteamResponse, error) {
	p := &CategorySteamParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/steam"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Game != nil {
		params.Set("game[]", fmt.Sprintf("%v", *p.Game))
	}
	if p.HoursPlayed != nil {
		params.Set("hours_played", fmt.Sprintf("%v", *p.HoursPlayed))
	}
	if p.HoursPlayedMax != nil {
		params.Set("hours_played_max", fmt.Sprintf("%v", *p.HoursPlayedMax))
	}
	if p.Eg != nil {
		params.Set("eg", fmt.Sprintf("%v", *p.Eg))
	}
	if p.Vac != nil {
		params.Set("vac[]", fmt.Sprintf("%v", *p.Vac))
	}
	if p.VacSkipGameCheck != nil {
		params.Set("vac_skip_game_check", fmt.Sprintf("%v", *p.VacSkipGameCheck))
	}
	if p.Rt != nil {
		params.Set("rt", fmt.Sprintf("%v", *p.Rt))
	}
	if p.TradeBan != nil {
		params.Set("trade_ban", fmt.Sprintf("%v", *p.TradeBan))
	}
	if p.TradeLimit != nil {
		params.Set("trade_limit", fmt.Sprintf("%v", *p.TradeLimit))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%v", *p.Limit))
	}
	if p.Mafile != nil {
		params.Set("mafile", fmt.Sprintf("%v", *p.Mafile))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.Lmin != nil {
		params.Set("lmin", fmt.Sprintf("%v", *p.Lmin))
	}
	if p.Lmax != nil {
		params.Set("lmax", fmt.Sprintf("%v", *p.Lmax))
	}
	if p.Rmin != nil {
		params.Set("rmin", fmt.Sprintf("%v", *p.Rmin))
	}
	if p.Rmax != nil {
		params.Set("rmax", fmt.Sprintf("%v", *p.Rmax))
	}
	if p.WingmanRmin != nil {
		params.Set("wingman_rmin", fmt.Sprintf("%v", *p.WingmanRmin))
	}
	if p.WingmanRmax != nil {
		params.Set("wingman_rmax", fmt.Sprintf("%v", *p.WingmanRmax))
	}
	if p.NoVac != nil {
		params.Set("no_vac", fmt.Sprintf("%v", *p.NoVac))
	}
	if p.MmBan != nil {
		params.Set("mm_ban", fmt.Sprintf("%v", *p.MmBan))
	}
	if p.BalanceMin != nil {
		params.Set("balance_min", fmt.Sprintf("%v", *p.BalanceMin))
	}
	if p.BalanceMax != nil {
		params.Set("balance_max", fmt.Sprintf("%v", *p.BalanceMax))
	}
	if p.InvGame != nil {
		params.Set("inv_game", fmt.Sprintf("%v", *p.InvGame))
	}
	if p.InvMin != nil {
		params.Set("inv_min", fmt.Sprintf("%v", *p.InvMin))
	}
	if p.InvMax != nil {
		params.Set("inv_max", fmt.Sprintf("%v", *p.InvMax))
	}
	if p.FriendsMin != nil {
		params.Set("friends_min", fmt.Sprintf("%v", *p.FriendsMin))
	}
	if p.FriendsMax != nil {
		params.Set("friends_max", fmt.Sprintf("%v", *p.FriendsMax))
	}
	if p.Gmin != nil {
		params.Set("gmin", fmt.Sprintf("%v", *p.Gmin))
	}
	if p.Gmax != nil {
		params.Set("gmax", fmt.Sprintf("%v", *p.Gmax))
	}
	if p.WinCountMin != nil {
		params.Set("win_count_min", fmt.Sprintf("%v", *p.WinCountMin))
	}
	if p.WinCountMax != nil {
		params.Set("win_count_max", fmt.Sprintf("%v", *p.WinCountMax))
	}
	if p.MedalId != nil {
		params.Set("medal_id[]", fmt.Sprintf("%v", *p.MedalId))
	}
	if p.MedalOperatorOr != nil {
		params.Set("medal_operator_or", fmt.Sprintf("%v", *p.MedalOperatorOr))
	}
	if p.MedalMin != nil {
		params.Set("medal_min", fmt.Sprintf("%v", *p.MedalMin))
	}
	if p.MedalMax != nil {
		params.Set("medal_max", fmt.Sprintf("%v", *p.MedalMax))
	}
	if p.Gift != nil {
		params.Set("gift[]", fmt.Sprintf("%v", *p.Gift))
	}
	if p.GiftMin != nil {
		params.Set("gift_min", fmt.Sprintf("%v", *p.GiftMin))
	}
	if p.GiftMax != nil {
		params.Set("gift_max", fmt.Sprintf("%v", *p.GiftMax))
	}
	if p.RecentlyHoursMin != nil {
		params.Set("recently_hours_min", fmt.Sprintf("%v", *p.RecentlyHoursMin))
	}
	if p.RecentlyHoursMax != nil {
		params.Set("recently_hours_max", fmt.Sprintf("%v", *p.RecentlyHoursMax))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Cs2ProfileRankMin != nil {
		params.Set("cs2_profile_rank_min", fmt.Sprintf("%v", *p.Cs2ProfileRankMin))
	}
	if p.Cs2ProfileRankMax != nil {
		params.Set("cs2_profile_rank_max", fmt.Sprintf("%v", *p.Cs2ProfileRankMax))
	}
	if p.SolommrMin != nil {
		params.Set("solommr_min", fmt.Sprintf("%v", *p.SolommrMin))
	}
	if p.SolommrMax != nil {
		params.Set("solommr_max", fmt.Sprintf("%v", *p.SolommrMax))
	}
	if p.D2GameCountMin != nil {
		params.Set("d2_game_count_min", fmt.Sprintf("%v", *p.D2GameCountMin))
	}
	if p.D2GameCountMax != nil {
		params.Set("d2_game_count_max", fmt.Sprintf("%v", *p.D2GameCountMax))
	}
	if p.D2WinCountMin != nil {
		params.Set("d2_win_count_min", fmt.Sprintf("%v", *p.D2WinCountMin))
	}
	if p.D2WinCountMax != nil {
		params.Set("d2_win_count_max", fmt.Sprintf("%v", *p.D2WinCountMax))
	}
	if p.D2BehaviorMin != nil {
		params.Set("d2_behavior_min", fmt.Sprintf("%v", *p.D2BehaviorMin))
	}
	if p.D2BehaviorMax != nil {
		params.Set("d2_behavior_max", fmt.Sprintf("%v", *p.D2BehaviorMax))
	}
	if p.FaceitLvlMin != nil {
		params.Set("faceit_lvl_min", fmt.Sprintf("%v", *p.FaceitLvlMin))
	}
	if p.FaceitLvlMax != nil {
		params.Set("faceit_lvl_max", fmt.Sprintf("%v", *p.FaceitLvlMax))
	}
	if p.PointsMin != nil {
		params.Set("points_min", fmt.Sprintf("%v", *p.PointsMin))
	}
	if p.PointsMax != nil {
		params.Set("points_max", fmt.Sprintf("%v", *p.PointsMax))
	}
	if p.RelevantGmin != nil {
		params.Set("relevant_gmin", fmt.Sprintf("%v", *p.RelevantGmin))
	}
	if p.RelevantGmax != nil {
		params.Set("relevant_gmax", fmt.Sprintf("%v", *p.RelevantGmax))
	}
	if p.LastTransDate != nil {
		params.Set("last_trans_date", fmt.Sprintf("%v", *p.LastTransDate))
	}
	if p.LastTransDatePeriod != nil {
		params.Set("last_trans_date_period", fmt.Sprintf("%v", *p.LastTransDatePeriod))
	}
	if p.LastTransDateLater != nil {
		params.Set("last_trans_date_later", fmt.Sprintf("%v", *p.LastTransDateLater))
	}
	if p.LastTransDatePeriodLater != nil {
		params.Set("last_trans_date_period_later", fmt.Sprintf("%v", *p.LastTransDatePeriodLater))
	}
	if p.NoTrans != nil {
		params.Set("no_trans", fmt.Sprintf("%v", *p.NoTrans))
	}
	if p.Trans != nil {
		params.Set("trans", fmt.Sprintf("%v", *p.Trans))
	}
	if p.GiftsPurchaseMin != nil {
		params.Set("gifts_purchase_min", fmt.Sprintf("%v", *p.GiftsPurchaseMin))
	}
	if p.GiftsPurchaseMax != nil {
		params.Set("gifts_purchase_max", fmt.Sprintf("%v", *p.GiftsPurchaseMax))
	}
	if p.RefundsPurchaseMin != nil {
		params.Set("refunds_purchase_min", fmt.Sprintf("%v", *p.RefundsPurchaseMin))
	}
	if p.RefundsPurchaseMax != nil {
		params.Set("refunds_purchase_max", fmt.Sprintf("%v", *p.RefundsPurchaseMax))
	}
	if p.IngamePurchaseMin != nil {
		params.Set("ingame_purchase_min", fmt.Sprintf("%v", *p.IngamePurchaseMin))
	}
	if p.IngamePurchaseMax != nil {
		params.Set("ingame_purchase_max", fmt.Sprintf("%v", *p.IngamePurchaseMax))
	}
	if p.GamesPurchaseMin != nil {
		params.Set("games_purchase_min", fmt.Sprintf("%v", *p.GamesPurchaseMin))
	}
	if p.GamesPurchaseMax != nil {
		params.Set("games_purchase_max", fmt.Sprintf("%v", *p.GamesPurchaseMax))
	}
	if p.PurchaseMin != nil {
		params.Set("purchase_min", fmt.Sprintf("%v", *p.PurchaseMin))
	}
	if p.PurchaseMax != nil {
		params.Set("purchase_max", fmt.Sprintf("%v", *p.PurchaseMax))
	}
	if p.HasActivatedKeys != nil {
		params.Set("has_activated_keys", fmt.Sprintf("%v", *p.HasActivatedKeys))
	}
	if p.EloMin != nil {
		params.Set("elo_min", fmt.Sprintf("%v", *p.EloMin))
	}
	if p.EloMax != nil {
		params.Set("elo_max", fmt.Sprintf("%v", *p.EloMax))
	}
	if p.Cs2MapRank != nil {
		params.Set("cs2_map_rank", fmt.Sprintf("%v", *p.Cs2MapRank))
	}
	if p.Cs2MapRmin != nil {
		params.Set("cs2_map_rmin", fmt.Sprintf("%v", *p.Cs2MapRmin))
	}
	if p.Cs2MapRmax != nil {
		params.Set("cs2_map_rmax", fmt.Sprintf("%v", *p.Cs2MapRmax))
	}
	if p.HasFaceit != nil {
		params.Set("has_faceit", fmt.Sprintf("%v", *p.HasFaceit))
	}
	if p.FaceitCsgoLvlMin != nil {
		params.Set("faceit_csgo_lvl_min", fmt.Sprintf("%v", *p.FaceitCsgoLvlMin))
	}
	if p.FaceitCsgoLvlMax != nil {
		params.Set("faceit_csgo_lvl_max", fmt.Sprintf("%v", *p.FaceitCsgoLvlMax))
	}
	if p.RustDeathsMin != nil {
		params.Set("rust_deaths_min", fmt.Sprintf("%v", *p.RustDeathsMin))
	}
	if p.RustDeathsMax != nil {
		params.Set("rust_deaths_max", fmt.Sprintf("%v", *p.RustDeathsMax))
	}
	if p.RustKillsMin != nil {
		params.Set("rust_kills_min", fmt.Sprintf("%v", *p.RustKillsMin))
	}
	if p.RustKillsMax != nil {
		params.Set("rust_kills_max", fmt.Sprintf("%v", *p.RustKillsMax))
	}
	if p.D2LastMatchDate != nil {
		params.Set("d2_last_match_date", fmt.Sprintf("%v", *p.D2LastMatchDate))
	}
	if p.D2LastMatchDatePeriod != nil {
		params.Set("d2_last_match_date_period", fmt.Sprintf("%v", *p.D2LastMatchDatePeriod))
	}
	if p.CardsMin != nil {
		params.Set("cards_min", fmt.Sprintf("%v", *p.CardsMin))
	}
	if p.CardsMax != nil {
		params.Set("cards_max", fmt.Sprintf("%v", *p.CardsMax))
	}
	if p.CardsGamesMin != nil {
		params.Set("cards_games_min", fmt.Sprintf("%v", *p.CardsGamesMin))
	}
	if p.CardsGamesMax != nil {
		params.Set("cards_games_max", fmt.Sprintf("%v", *p.CardsGamesMax))
	}
	if p.SkipVacInv != nil {
		params.Set("skip_vac_inv", fmt.Sprintf("%v", *p.SkipVacInv))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategorySteamResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategorySteamParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailType *[]string
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	ItemDomain *string
	Game *[]int64
	HoursPlayed *map[string]any
	HoursPlayedMax *map[string]any
	Eg *int64
	Vac *[]int64
	VacSkipGameCheck *bool
	Rt *string
	TradeBan *string
	TradeLimit *string
	Daybreak *int64
	Limit *string
	Mafile *string
	Reg *int64
	RegPeriod *string
	Lmin *int64
	Lmax *int64
	Rmin *int64
	Rmax *int64
	WingmanRmin *int64
	WingmanRmax *int64
	NoVac *bool
	MmBan *string
	BalanceMin *int64
	BalanceMax *int64
	InvGame *int64
	InvMin *float64
	InvMax *float64
	FriendsMin *int64
	FriendsMax *int64
	Gmin *int64
	Gmax *int64
	WinCountMin *int64
	WinCountMax *int64
	MedalId *[]int64
	MedalOperatorOr *bool
	MedalMin *int64
	MedalMax *int64
	Gift *[]string
	GiftMin *int64
	GiftMax *int64
	RecentlyHoursMin *int64
	RecentlyHoursMax *int64
	Country *[]string
	NotCountry *[]string
	Cs2ProfileRankMin *int64
	Cs2ProfileRankMax *int64
	SolommrMin *int64
	SolommrMax *int64
	D2GameCountMin *int64
	D2GameCountMax *int64
	D2WinCountMin *int64
	D2WinCountMax *int64
	D2BehaviorMin *int64
	D2BehaviorMax *int64
	FaceitLvlMin *int64
	FaceitLvlMax *int64
	PointsMin *int64
	PointsMax *int64
	RelevantGmin *int64
	RelevantGmax *int64
	LastTransDate *int64
	LastTransDatePeriod *string
	LastTransDateLater *int64
	LastTransDatePeriodLater *string
	NoTrans *bool
	Trans *bool
	GiftsPurchaseMin *float64
	GiftsPurchaseMax *float64
	RefundsPurchaseMin *float64
	RefundsPurchaseMax *float64
	IngamePurchaseMin *float64
	IngamePurchaseMax *float64
	GamesPurchaseMin *float64
	GamesPurchaseMax *float64
	PurchaseMin *float64
	PurchaseMax *float64
	HasActivatedKeys *string
	EloMin *int64
	EloMax *int64
	Cs2MapRank *int64
	Cs2MapRmin *int64
	Cs2MapRmax *int64
	HasFaceit *string
	FaceitCsgoLvlMin *int64
	FaceitCsgoLvlMax *int64
	RustDeathsMin *int64
	RustDeathsMax *int64
	RustKillsMin *int64
	RustKillsMax *int64
	D2LastMatchDate *int64
	D2LastMatchDatePeriod *string
	CardsMin *int64
	CardsMax *int64
	CardsGamesMin *int64
	CardsGamesMax *int64
	SkipVacInv *bool
}

func (c *Client) CategoryEA(opts ...func(*CategoryEAParams)) (*CategoryEAResponse, error) {
	p := &CategoryEAParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/ea"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Game != nil {
		params.Set("game[]", fmt.Sprintf("%v", *p.Game))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Gmin != nil {
		params.Set("gmin", fmt.Sprintf("%v", *p.Gmin))
	}
	if p.Gmax != nil {
		params.Set("gmax", fmt.Sprintf("%v", *p.Gmax))
	}
	if p.AlRankMin != nil {
		params.Set("al_rank_min", fmt.Sprintf("%v", *p.AlRankMin))
	}
	if p.AlRankMax != nil {
		params.Set("al_rank_max", fmt.Sprintf("%v", *p.AlRankMax))
	}
	if p.AlLevelMin != nil {
		params.Set("al_level_min", fmt.Sprintf("%v", *p.AlLevelMin))
	}
	if p.AlLevelMax != nil {
		params.Set("al_level_max", fmt.Sprintf("%v", *p.AlLevelMax))
	}
	if p.HasBan != nil {
		params.Set("has_ban", fmt.Sprintf("%v", *p.HasBan))
	}
	if p.XboxConnected != nil {
		params.Set("xbox_connected", fmt.Sprintf("%v", *p.XboxConnected))
	}
	if p.SteamConnected != nil {
		params.Set("steam_connected", fmt.Sprintf("%v", *p.SteamConnected))
	}
	if p.PsnConnected != nil {
		params.Set("psn_connected", fmt.Sprintf("%v", *p.PsnConnected))
	}
	if p.Subscription != nil {
		params.Set("subscription", fmt.Sprintf("%v", *p.Subscription))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}
	if p.HoursPlayed != nil {
		params.Set("hours_played", fmt.Sprintf("%v", *p.HoursPlayed))
	}
	if p.HoursPlayedMax != nil {
		params.Set("hours_played_max", fmt.Sprintf("%v", *p.HoursPlayedMax))
	}
	if p.Transactions != nil {
		params.Set("transactions", fmt.Sprintf("%v", *p.Transactions))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryEAResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryEAParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Game *[]string
	Country *[]string
	NotCountry *[]string
	Gmin *int64
	Gmax *int64
	AlRankMin *int64
	AlRankMax *int64
	AlLevelMin *int64
	AlLevelMax *int64
	HasBan *string
	XboxConnected *string
	SteamConnected *string
	PsnConnected *string
	Subscription *string
	SubscriptionLength *int64
	SubscriptionPeriod *string
	HoursPlayed *map[string]any
	HoursPlayedMax *map[string]any
	Transactions *string
}

func (c *Client) CategoryEpicGames(opts ...func(*CategoryEpicGamesParams)) (*CategoryEpicGamesResponse, error) {
	p := &CategoryEpicGamesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/epicgames"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Eg != nil {
		params.Set("eg", fmt.Sprintf("%v", *p.Eg))
	}
	if p.Game != nil {
		params.Set("game[]", fmt.Sprintf("%v", *p.Game))
	}
	if p.ChangeEmail != nil {
		params.Set("change_email", fmt.Sprintf("%v", *p.ChangeEmail))
	}
	if p.RlPurchases != nil {
		params.Set("rl_purchases", fmt.Sprintf("%v", *p.RlPurchases))
	}
	if p.BalanceMin != nil {
		params.Set("balance_min", fmt.Sprintf("%v", *p.BalanceMin))
	}
	if p.BalanceMax != nil {
		params.Set("balance_max", fmt.Sprintf("%v", *p.BalanceMax))
	}
	if p.RewardsBalanceMin != nil {
		params.Set("rewards_balance_min", fmt.Sprintf("%v", *p.RewardsBalanceMin))
	}
	if p.RewardsBalanceMax != nil {
		params.Set("rewards_balance_max", fmt.Sprintf("%v", *p.RewardsBalanceMax))
	}
	if p.Gmin != nil {
		params.Set("gmin", fmt.Sprintf("%v", *p.Gmin))
	}
	if p.Gmax != nil {
		params.Set("gmax", fmt.Sprintf("%v", *p.Gmax))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.HoursPlayed != nil {
		params.Set("hours_played", fmt.Sprintf("%v", *p.HoursPlayed))
	}
	if p.HoursPlayedMax != nil {
		params.Set("hours_played_max", fmt.Sprintf("%v", *p.HoursPlayedMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryEpicGamesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryEpicGamesParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Eg *int64
	Game *[]string
	ChangeEmail *string
	RlPurchases *bool
	BalanceMin *float64
	BalanceMax *float64
	RewardsBalanceMin *float64
	RewardsBalanceMax *float64
	Gmin *int64
	Gmax *int64
	Country *[]string
	NotCountry *[]string
	Daybreak *int64
	HoursPlayed *map[string]any
	HoursPlayedMax *map[string]any
}

func (c *Client) CategoryParams(categoryName string, opts ...func(*CategoryParamsParams)) (*CategoryParamsResponse, error) {
	p := &CategoryParamsParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{categoryName}/params"
	path = strings.Replace(path, "{categoryName}", fmt.Sprintf("%v", categoryName), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryParamsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryParamsParams struct {
}

func (c *Client) ListViewed(opts ...func(*ListViewedParams)) (*ItemListModel, error) {
	p := &ListViewedParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/viewed"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Show != nil {
		params.Set("show", fmt.Sprintf("%v", *p.Show))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ListViewedParams struct {
	Page *int64
	Show *string
	Title *string
	Pmin *int64
	Pmax *int64
	Origin *[]string
	NotOrigin *[]string
	OrderBy *string
	Sb *bool
	SbByMe *bool
	Nsb *bool
	NsbByMe *bool
}

func (c *Client) ManagingAutoBuyPrice(itemId int64, opts ...func(*ManagingAutoBuyPriceParams)) (*ManagingAutoBuyPriceResponse, error) {
	p := &ManagingAutoBuyPriceParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/auto-buy-price"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingAutoBuyPriceResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingAutoBuyPriceParams struct {
}

func (c *Client) CartGet(opts ...func(*CartGetParams)) (map[string]any, error) {
	p := &CartGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/cart"

	params := url.Values{}
	if p.CategoryId != nil {
		params.Set("category_id", fmt.Sprintf("%v", *p.CategoryId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
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

type CartGetParams struct {
	CategoryId *int64
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
}

func (c *Client) CartAdd(itemId int64, opts ...func(*CartAddParams)) (*CartAddResponse, error) {
	p := &CartAddParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/cart"


	bodyMap := map[string]any{}
	bodyMap["item_id"] = itemId
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result CartAddResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CartAddParams struct {
}

func (c *Client) CartDelete(opts ...func(*CartDeleteParams)) (*CartDeleteResponse, error) {
	p := &CartDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/cart"


	bodyMap := map[string]any{}
	if p.ItemId != nil {
		bodyMap["item_id"] = *p.ItemId
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("DELETE", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result CartDeleteResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CartDeleteParams struct {
	ItemId *int64
}

func (c *Client) ManagingTempEmailPassword(itemId int64, opts ...func(*ManagingTempEmailPasswordParams)) (*ManagingTempEmailPasswordResponse, error) {
	p := &ManagingTempEmailPasswordParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/temp-email-password"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingTempEmailPasswordResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingTempEmailPasswordParams struct {
}

func (c *Client) CategoryDiscord(opts ...func(*CategoryDiscordParams)) (*CategoryDiscordResponse, error) {
	p := &CategoryDiscordParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/discord"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Nitro != nil {
		params.Set("nitro", fmt.Sprintf("%v", *p.Nitro))
	}
	if p.NitroType != nil {
		params.Set("nitro_type[]", fmt.Sprintf("%v", *p.NitroType))
	}
	if p.NitroLength != nil {
		params.Set("nitro_length", fmt.Sprintf("%v", *p.NitroLength))
	}
	if p.NitroPeriod != nil {
		params.Set("nitro_period", fmt.Sprintf("%v", *p.NitroPeriod))
	}
	if p.BoostsMin != nil {
		params.Set("boosts_min", fmt.Sprintf("%v", *p.BoostsMin))
	}
	if p.BoostsMax != nil {
		params.Set("boosts_max", fmt.Sprintf("%v", *p.BoostsMax))
	}
	if p.Billing != nil {
		params.Set("billing", fmt.Sprintf("%v", *p.Billing))
	}
	if p.Gifts != nil {
		params.Set("gifts", fmt.Sprintf("%v", *p.Gifts))
	}
	if p.Transactions != nil {
		params.Set("transactions", fmt.Sprintf("%v", *p.Transactions))
	}
	if p.Badge != nil {
		params.Set("badge[]", fmt.Sprintf("%v", *p.Badge))
	}
	if p.Condition != nil {
		params.Set("condition[]", fmt.Sprintf("%v", *p.Condition))
	}
	if p.ChatMin != nil {
		params.Set("chat_min", fmt.Sprintf("%v", *p.ChatMin))
	}
	if p.ChatMax != nil {
		params.Set("chat_max", fmt.Sprintf("%v", *p.ChatMax))
	}
	if p.MinAdminMembers != nil {
		params.Set("min_admin_members", fmt.Sprintf("%v", *p.MinAdminMembers))
	}
	if p.MaxAdminMembers != nil {
		params.Set("max_admin_members", fmt.Sprintf("%v", *p.MaxAdminMembers))
	}
	if p.MinAdmin != nil {
		params.Set("min_admin", fmt.Sprintf("%v", *p.MinAdmin))
	}
	if p.MaxAdmin != nil {
		params.Set("max_admin", fmt.Sprintf("%v", *p.MaxAdmin))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.Language != nil {
		params.Set("language[]", fmt.Sprintf("%v", *p.Language))
	}
	if p.NotLanguage != nil {
		params.Set("not_language[]", fmt.Sprintf("%v", *p.NotLanguage))
	}
	if p.Clans != nil {
		params.Set("clans", fmt.Sprintf("%v", *p.Clans))
	}
	if p.MinAdminClans != nil {
		params.Set("min_admin_clans", fmt.Sprintf("%v", *p.MinAdminClans))
	}
	if p.MaxAdminClans != nil {
		params.Set("max_admin_clans", fmt.Sprintf("%v", *p.MaxAdminClans))
	}
	if p.MinOwnerClans != nil {
		params.Set("min_owner_clans", fmt.Sprintf("%v", *p.MinOwnerClans))
	}
	if p.MaxOwnerClans != nil {
		params.Set("max_owner_clans", fmt.Sprintf("%v", *p.MaxOwnerClans))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.MinServers != nil {
		params.Set("min_servers", fmt.Sprintf("%v", *p.MinServers))
	}
	if p.MaxServers != nil {
		params.Set("max_servers", fmt.Sprintf("%v", *p.MaxServers))
	}
	if p.X2fa != nil {
		params.Set("2fa", fmt.Sprintf("%v", *p.X2fa))
	}
	if p.MinFullCredits != nil {
		params.Set("min_full_credits", fmt.Sprintf("%v", *p.MinFullCredits))
	}
	if p.MaxFullCredits != nil {
		params.Set("max_full_credits", fmt.Sprintf("%v", *p.MaxFullCredits))
	}
	if p.MinBasicCredits != nil {
		params.Set("min_basic_credits", fmt.Sprintf("%v", *p.MinBasicCredits))
	}
	if p.MaxBasicCredits != nil {
		params.Set("max_basic_credits", fmt.Sprintf("%v", *p.MaxBasicCredits))
	}
	if p.MinOrbs != nil {
		params.Set("min_orbs", fmt.Sprintf("%v", *p.MinOrbs))
	}
	if p.MaxOrbs != nil {
		params.Set("max_orbs", fmt.Sprintf("%v", *p.MaxOrbs))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryDiscordResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryDiscordParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Tel *string
	Nitro *string
	NitroType *[]string
	NitroLength *int64
	NitroPeriod *string
	BoostsMin *int64
	BoostsMax *int64
	Billing *string
	Gifts *string
	Transactions *string
	Badge *[]string
	Condition *[]string
	ChatMin *int64
	ChatMax *int64
	MinAdminMembers *int64
	MaxAdminMembers *int64
	MinAdmin *int64
	MaxAdmin *int64
	Reg *int64
	RegPeriod *string
	Language *[]string
	NotLanguage *[]string
	Clans *string
	MinAdminClans *int64
	MaxAdminClans *int64
	MinOwnerClans *int64
	MaxOwnerClans *int64
	Country *[]string
	NotCountry *[]string
	MinServers *int64
	MaxServers *int64
	X2fa *string
	MinFullCredits *int64
	MaxFullCredits *int64
	MinBasicCredits *int64
	MaxBasicCredits *int64
	MinOrbs *int64
	MaxOrbs *int64
}

func (c *Client) CategoryWarface(opts ...func(*CategoryWarfaceParams)) (*CategoryWarfaceResponse, error) {
	p := &CategoryWarfaceParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/warface"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.RankMin != nil {
		params.Set("rank_min", fmt.Sprintf("%v", *p.RankMin))
	}
	if p.RankMax != nil {
		params.Set("rank_max", fmt.Sprintf("%v", *p.RankMax))
	}
	if p.BonusRankMin != nil {
		params.Set("bonus_rank_min", fmt.Sprintf("%v", *p.BonusRankMin))
	}
	if p.BonusRankMax != nil {
		params.Set("bonus_rank_max", fmt.Sprintf("%v", *p.BonusRankMax))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.KreditsMin != nil {
		params.Set("kredits_min", fmt.Sprintf("%v", *p.KreditsMin))
	}
	if p.KreditsMax != nil {
		params.Set("kredits_max", fmt.Sprintf("%v", *p.KreditsMax))
	}
	if p.TotalKreditsMin != nil {
		params.Set("total_kredits_min", fmt.Sprintf("%v", *p.TotalKreditsMin))
	}
	if p.TotalKreditsMax != nil {
		params.Set("total_kredits_max", fmt.Sprintf("%v", *p.TotalKreditsMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryWarfaceResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryWarfaceParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	RankMin *int64
	RankMax *int64
	BonusRankMin *int64
	BonusRankMax *int64
	Tel *string
	Daybreak *int64
	KreditsMin *int64
	KreditsMax *int64
	TotalKreditsMin *int64
	TotalKreditsMax *int64
}

func (c *Client) CategoryList(opts ...func(*CategoryListParams)) (*CategoryListResponse, error) {
	p := &CategoryListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/category"

	params := url.Values{}
	if p.TopQueries != nil {
		params.Set("top_queries", fmt.Sprintf("%v", *p.TopQueries))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryListParams struct {
	TopQueries *bool
}

func (c *Client) ManagingBulkGet(opts ...func(*ManagingBulkGetParams)) (*ManagingBulkGetResponse, error) {
	p := &ManagingBulkGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/bulk/items"


	bodyMap := map[string]any{}
	if p.ItemId != nil {
		bodyMap["item_id"] = *p.ItemId
	}
	if p.ParseSameItemIds != nil {
		bodyMap["parse_same_item_ids"] = *p.ParseSameItemIds
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ManagingBulkGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingBulkGetParams struct {
	ItemId *[]int64
	ParseSameItemIds *bool
}

func (c *Client) PublishingAdd(categoryId int64, currency string, itemOrigin string, price float64, opts ...func(*PublishingAddParams)) (*PublishingAddResponse, error) {
	p := &PublishingAddParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/item/add"


	bodyMap := map[string]any{}
	bodyMap["category_id"] = categoryId
	bodyMap["currency"] = currency
	if p.ForceTempEmail != nil {
		bodyMap["forceTempEmail"] = *p.ForceTempEmail
	}
	if p.HasEmailLoginData != nil {
		bodyMap["has_email_login_data"] = *p.HasEmailLoginData
	}
	if p.EmailLoginData != nil {
		bodyMap["email_login_data"] = *p.EmailLoginData
	}
	if p.ProxyId != nil {
		bodyMap["proxy_id"] = *p.ProxyId
	}
	if p.ExtendedGuarantee != nil {
		bodyMap["extended_guarantee"] = *p.ExtendedGuarantee
	}
	if p.Information != nil {
		bodyMap["information"] = *p.Information
	}
	bodyMap["item_origin"] = itemOrigin
	if p.ResellItemId != nil {
		bodyMap["resell_item_id"] = *p.ResellItemId
	}
	if p.AllowAskDiscount != nil {
		bodyMap["allow_ask_discount"] = *p.AllowAskDiscount
	}
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.Description != nil {
		bodyMap["description"] = *p.Description
	}
	if p.EmailType != nil {
		bodyMap["email_type"] = *p.EmailType
	}
	if p.RandomProxy != nil {
		bodyMap["random_proxy"] = *p.RandomProxy
	}
	bodyMap["price"] = price
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PublishingAddResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PublishingAddParams struct {
	ForceTempEmail *bool
	HasEmailLoginData *bool
	EmailLoginData *string
	ProxyId *int64
	ExtendedGuarantee *int64
	Information *string
	ResellItemId *int64
	AllowAskDiscount *bool
	Title *string
	TitleEn *string
	Description *string
	EmailType *string
	RandomProxy *any
}

func (c *Client) ManagingStick(itemId int64, opts ...func(*ManagingStickParams)) (map[string]any, error) {
	p := &ManagingStickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/stick"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingStickParams struct {
}

func (c *Client) ManagingUnstick(itemId int64, opts ...func(*ManagingUnstickParams)) (map[string]any, error) {
	p := &ManagingUnstickParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/stick"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingUnstickParams struct {
}

func (c *Client) ProfileGet(opts ...func(*ProfileGetParams)) (*ProfileGetResponse, error) {
	p := &ProfileGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/me"

	params := url.Values{}
	if p.FieldsInclude != nil {
		params.Set("fields_include", fmt.Sprintf("%v", *p.FieldsInclude))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ProfileGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ProfileGetParams struct {
	FieldsInclude *[]string
}

func (c *Client) ProfileEdit(opts ...func(*ProfileEditParams)) (map[string]any, error) {
	p := &ProfileEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/me"


	bodyMap := map[string]any{}
	if p.Option != nil {
		bodyMap["option"] = *p.Option
	}
	if p.TelegramApiId != nil {
		bodyMap["telegram_api_id"] = *p.TelegramApiId
	}
	if p.TelegramDeviceModel != nil {
		bodyMap["telegram_device_model"] = *p.TelegramDeviceModel
	}
	if p.TelegramAppVersion != nil {
		bodyMap["telegram_app_version"] = *p.TelegramAppVersion
	}
	if p.TelegramLangPack != nil {
		bodyMap["telegram_lang_pack"] = *p.TelegramLangPack
	}
	if p.User != nil {
		bodyMap["user"] = *p.User
	}
	if p.AllowAcceptAccounts != nil {
		bodyMap["allow_accept_accounts"] = *p.AllowAcceptAccounts
	}
	if p.TelegramApiHash != nil {
		bodyMap["telegram_api_hash"] = *p.TelegramApiHash
	}
	if p.TelegramSystemVersion != nil {
		bodyMap["telegram_system_version"] = *p.TelegramSystemVersion
	}
	if p.TelegramLangCode != nil {
		bodyMap["telegram_lang_code"] = *p.TelegramLangCode
	}
	if p.TelegramSystemLangCode != nil {
		bodyMap["telegram_system_lang_code"] = *p.TelegramSystemLangCode
	}
	if p.ClearTelegramClient != nil {
		bodyMap["clear_telegram_client"] = *p.ClearTelegramClient
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

type ProfileEditParams struct {
	Option *map[string]any
	TelegramApiId *string
	TelegramDeviceModel *string
	TelegramAppVersion *string
	TelegramLangPack *string
	User *map[string]any
	AllowAcceptAccounts *[]string
	TelegramApiHash *string
	TelegramSystemVersion *string
	TelegramLangCode *string
	TelegramSystemLangCode *string
	ClearTelegramClient *bool
}

func (c *Client) CategoryTelegram(opts ...func(*CategoryTelegramParams)) (*CategoryTelegramResponse, error) {
	p := &CategoryTelegramParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/telegram"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Spam != nil {
		params.Set("spam", fmt.Sprintf("%v", *p.Spam))
	}
	if p.Password != nil {
		params.Set("password", fmt.Sprintf("%v", *p.Password))
	}
	if p.Premium != nil {
		params.Set("premium", fmt.Sprintf("%v", *p.Premium))
	}
	if p.PremiumExpiration != nil {
		params.Set("premium_expiration", fmt.Sprintf("%v", *p.PremiumExpiration))
	}
	if p.PremiumExpirationPeriod != nil {
		params.Set("premium_expiration_period", fmt.Sprintf("%v", *p.PremiumExpirationPeriod))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.MinChannels != nil {
		params.Set("min_channels", fmt.Sprintf("%v", *p.MinChannels))
	}
	if p.MaxChannels != nil {
		params.Set("max_channels", fmt.Sprintf("%v", *p.MaxChannels))
	}
	if p.MinChats != nil {
		params.Set("min_chats", fmt.Sprintf("%v", *p.MinChats))
	}
	if p.MaxChats != nil {
		params.Set("max_chats", fmt.Sprintf("%v", *p.MaxChats))
	}
	if p.MinConversations != nil {
		params.Set("min_conversations", fmt.Sprintf("%v", *p.MinConversations))
	}
	if p.MaxConversations != nil {
		params.Set("max_conversations", fmt.Sprintf("%v", *p.MaxConversations))
	}
	if p.MinAdmin != nil {
		params.Set("min_admin", fmt.Sprintf("%v", *p.MinAdmin))
	}
	if p.MaxAdmin != nil {
		params.Set("max_admin", fmt.Sprintf("%v", *p.MaxAdmin))
	}
	if p.MinAdminSub != nil {
		params.Set("min_admin_sub", fmt.Sprintf("%v", *p.MinAdminSub))
	}
	if p.MaxAdminSub != nil {
		params.Set("max_admin_sub", fmt.Sprintf("%v", *p.MaxAdminSub))
	}
	if p.DigMin != nil {
		params.Set("dig_min", fmt.Sprintf("%v", *p.DigMin))
	}
	if p.DigMax != nil {
		params.Set("dig_max", fmt.Sprintf("%v", *p.DigMax))
	}
	if p.MinContacts != nil {
		params.Set("min_contacts", fmt.Sprintf("%v", *p.MinContacts))
	}
	if p.MaxContacts != nil {
		params.Set("max_contacts", fmt.Sprintf("%v", *p.MaxContacts))
	}
	if p.MinStars != nil {
		params.Set("min_stars", fmt.Sprintf("%v", *p.MinStars))
	}
	if p.MaxStars != nil {
		params.Set("max_stars", fmt.Sprintf("%v", *p.MaxStars))
	}
	if p.Birthday != nil {
		params.Set("birthday", fmt.Sprintf("%v", *p.Birthday))
	}
	if p.BirthdayPeriod != nil {
		params.Set("birthday_period", fmt.Sprintf("%v", *p.BirthdayPeriod))
	}
	if p.BirthdayAfter != nil {
		params.Set("birthday_after", fmt.Sprintf("%v", *p.BirthdayAfter))
	}
	if p.BirthdayAfterPeriod != nil {
		params.Set("birthday_after_period", fmt.Sprintf("%v", *p.BirthdayAfterPeriod))
	}
	if p.MinId != nil {
		params.Set("min_id", fmt.Sprintf("%v", *p.MinId))
	}
	if p.MaxId != nil {
		params.Set("max_id", fmt.Sprintf("%v", *p.MaxId))
	}
	if p.AllowGeoSpamblock != nil {
		params.Set("allow_geo_spamblock", fmt.Sprintf("%v", *p.AllowGeoSpamblock))
	}
	if p.MinGifts != nil {
		params.Set("min_gifts", fmt.Sprintf("%v", *p.MinGifts))
	}
	if p.MaxGifts != nil {
		params.Set("max_gifts", fmt.Sprintf("%v", *p.MaxGifts))
	}
	if p.MinNftGifts != nil {
		params.Set("min_nft_gifts", fmt.Sprintf("%v", *p.MinNftGifts))
	}
	if p.MaxNftGifts != nil {
		params.Set("max_nft_gifts", fmt.Sprintf("%v", *p.MaxNftGifts))
	}
	if p.MinGiftsStars != nil {
		params.Set("min_gifts_stars", fmt.Sprintf("%v", *p.MinGiftsStars))
	}
	if p.MaxGiftsStars != nil {
		params.Set("max_gifts_stars", fmt.Sprintf("%v", *p.MaxGiftsStars))
	}
	if p.MinGiftsConvertStars != nil {
		params.Set("min_gifts_convert_stars", fmt.Sprintf("%v", *p.MinGiftsConvertStars))
	}
	if p.MaxGiftsConvertStars != nil {
		params.Set("max_gifts_convert_stars", fmt.Sprintf("%v", *p.MaxGiftsConvertStars))
	}
	if p.DcId != nil {
		params.Set("dc_id[]", fmt.Sprintf("%v", *p.DcId))
	}
	if p.NotDcId != nil {
		params.Set("not_dc_id[]", fmt.Sprintf("%v", *p.NotDcId))
	}
	if p.Email != nil {
		params.Set("email", fmt.Sprintf("%v", *p.Email))
	}
	if p.MinBots != nil {
		params.Set("min_bots", fmt.Sprintf("%v", *p.MinBots))
	}
	if p.MaxBots != nil {
		params.Set("max_bots", fmt.Sprintf("%v", *p.MaxBots))
	}
	if p.MinBotActiveUsers != nil {
		params.Set("min_bot_active_users", fmt.Sprintf("%v", *p.MinBotActiveUsers))
	}
	if p.MaxBotActiveUsers != nil {
		params.Set("max_bot_active_users", fmt.Sprintf("%v", *p.MaxBotActiveUsers))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryTelegramResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryTelegramParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Spam *string
	Password *string
	Premium *string
	PremiumExpiration *int64
	PremiumExpirationPeriod *string
	Country *[]string
	NotCountry *[]string
	Daybreak *int64
	MinChannels *int64
	MaxChannels *int64
	MinChats *int64
	MaxChats *int64
	MinConversations *int64
	MaxConversations *int64
	MinAdmin *int64
	MaxAdmin *int64
	MinAdminSub *int64
	MaxAdminSub *int64
	DigMin *int64
	DigMax *int64
	MinContacts *int64
	MaxContacts *int64
	MinStars *int64
	MaxStars *int64
	Birthday *int64
	BirthdayPeriod *string
	BirthdayAfter *int64
	BirthdayAfterPeriod *string
	MinId *int64
	MaxId *int64
	AllowGeoSpamblock *bool
	MinGifts *int64
	MaxGifts *int64
	MinNftGifts *int64
	MaxNftGifts *int64
	MinGiftsStars *int64
	MaxGiftsStars *int64
	MinGiftsConvertStars *int64
	MaxGiftsConvertStars *int64
	DcId *[]int64
	NotDcId *[]int64
	Email *string
	MinBots *int64
	MaxBots *int64
	MinBotActiveUsers *int64
	MaxBotActiveUsers *int64
}

func (c *Client) ListDownload(typeValue string, opts ...func(*ListDownloadParams)) (map[string]any, error) {
	p := &ListDownloadParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/user/{type}/download"
	path = strings.Replace(path, "{type}", fmt.Sprintf("%v", typeValue), 1)

	params := url.Values{}
	if p.Format != nil {
		params.Set("format", fmt.Sprintf("%v", *p.Format))
	}
	if p.CustomFormat != nil {
		params.Set("custom_format", fmt.Sprintf("%v", *p.CustomFormat))
	}
	if p.CategoryId != nil {
		params.Set("category_id", fmt.Sprintf("%v", *p.CategoryId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Show != nil {
		params.Set("show", fmt.Sprintf("%v", *p.Show))
	}
	if p.DeleteReason != nil {
		params.Set("delete_reason", fmt.Sprintf("%v", *p.DeleteReason))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.Username != nil {
		params.Set("username", fmt.Sprintf("%v", *p.Username))
	}
	if p.PublishedStartDate != nil {
		params.Set("published_startDate", fmt.Sprintf("%v", *p.PublishedStartDate))
	}
	if p.PublishedEndDate != nil {
		params.Set("published_endDate", fmt.Sprintf("%v", *p.PublishedEndDate))
	}
	if p.FilterByPublishedDate != nil {
		params.Set("filter_by_published_date", fmt.Sprintf("%v", *p.FilterByPublishedDate))
	}
	if p.PaidStartDate != nil {
		params.Set("paid_startDate", fmt.Sprintf("%v", *p.PaidStartDate))
	}
	if p.PaidEndDate != nil {
		params.Set("paid_endDate", fmt.Sprintf("%v", *p.PaidEndDate))
	}
	if p.FilterByBuyerOperationDate != nil {
		params.Set("filter_by_buyer_operation_date", fmt.Sprintf("%v", *p.FilterByBuyerOperationDate))
	}
	if p.DeleteStartDate != nil {
		params.Set("delete_startDate", fmt.Sprintf("%v", *p.DeleteStartDate))
	}
	if p.DeleteEndDate != nil {
		params.Set("delete_endDate", fmt.Sprintf("%v", *p.DeleteEndDate))
	}
	if p.FilterByDeleteDate != nil {
		params.Set("filter_by_delete_date", fmt.Sprintf("%v", *p.FilterByDeleteDate))
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

type ListDownloadParams struct {
	Format *string
	CustomFormat *string
	CategoryId *int64
	Page *int64
	Show *string
	DeleteReason *string
	Title *string
	Pmin *int64
	Pmax *int64
	Origin *[]string
	NotOrigin *[]string
	OrderBy *string
	Sb *bool
	SbByMe *bool
	Nsb *bool
	NsbByMe *bool
	Username *string
	PublishedStartDate *string
	PublishedEndDate *string
	FilterByPublishedDate *bool
	PaidStartDate *string
	PaidEndDate *string
	FilterByBuyerOperationDate *bool
	DeleteStartDate *string
	DeleteEndDate *string
	FilterByDeleteDate *bool
}

func (c *Client) CustomDiscountsEdit(discountId int64, opts ...func(*CustomDiscountsEditParams)) (*CustomDiscountsEditResponse, error) {
	p := &CustomDiscountsEditParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/custom-discounts"


	bodyMap := map[string]any{}
	bodyMap["discount_id"] = discountId
	if p.DiscountPercent != nil {
		bodyMap["discount_percent"] = *p.DiscountPercent
	}
	if p.MinPrice != nil {
		bodyMap["min_price"] = *p.MinPrice
	}
	if p.MaxPrice != nil {
		bodyMap["max_price"] = *p.MaxPrice
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("PUT", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result CustomDiscountsEditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CustomDiscountsEditParams struct {
	DiscountPercent *float64
	MinPrice *float64
	MaxPrice *float64
}

func (c *Client) CustomDiscountsDelete(discountId int64, opts ...func(*CustomDiscountsDeleteParams)) (map[string]any, error) {
	p := &CustomDiscountsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/custom-discounts"


	bodyMap := map[string]any{}
	bodyMap["discount_id"] = discountId
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

type CustomDiscountsDeleteParams struct {
}

func (c *Client) CustomDiscountsGet(opts ...func(*CustomDiscountsGetParams)) (*CustomDiscountsGetResponse, error) {
	p := &CustomDiscountsGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/custom-discounts"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result CustomDiscountsGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CustomDiscountsGetParams struct {
}

func (c *Client) CustomDiscountsCreate(userId int64, categoryId int64, discountPercent float64, minPrice float64, opts ...func(*CustomDiscountsCreateParams)) (*CustomDiscountsCreateResponse, error) {
	p := &CustomDiscountsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/custom-discounts"


	bodyMap := map[string]any{}
	if p.Currency != nil {
		bodyMap["currency"] = *p.Currency
	}
	bodyMap["user_id"] = userId
	bodyMap["category_id"] = categoryId
	bodyMap["discount_percent"] = discountPercent
	bodyMap["min_price"] = minPrice
	if p.MaxPrice != nil {
		bodyMap["max_price"] = *p.MaxPrice
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result CustomDiscountsCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CustomDiscountsCreateParams struct {
	Currency *string
	MaxPrice *float64
}

func (c *Client) PublishingFastSell(categoryId int64, price float64, itemOrigin string, currency string, opts ...func(*PublishingFastSellParams)) (*PublishingFastSellResponse, error) {
	p := &PublishingFastSellParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/item/fast-sell"


	bodyMap := map[string]any{}
	bodyMap["category_id"] = categoryId
	if p.ProxyId != nil {
		bodyMap["proxy_id"] = *p.ProxyId
	}
	if p.Login != nil {
		bodyMap["login"] = *p.Login
	}
	bodyMap["price"] = price
	bodyMap["item_origin"] = itemOrigin
	if p.ExtendedGuarantee != nil {
		bodyMap["extended_guarantee"] = *p.ExtendedGuarantee
	}
	if p.RandomProxy != nil {
		bodyMap["random_proxy"] = *p.RandomProxy
	}
	if p.Extra != nil {
		bodyMap["extra"] = *p.Extra
	}
	if p.Title != nil {
		bodyMap["title"] = *p.Title
	}
	if p.TitleEn != nil {
		bodyMap["title_en"] = *p.TitleEn
	}
	if p.AllowAskDiscount != nil {
		bodyMap["allow_ask_discount"] = *p.AllowAskDiscount
	}
	if p.Information != nil {
		bodyMap["information"] = *p.Information
	}
	if p.LoginPassword != nil {
		bodyMap["login_password"] = *p.LoginPassword
	}
	if p.EmailType != nil {
		bodyMap["email_type"] = *p.EmailType
	}
	bodyMap["currency"] = currency
	if p.Description != nil {
		bodyMap["description"] = *p.Description
	}
	if p.Password != nil {
		bodyMap["password"] = *p.Password
	}
	if p.HasEmailLoginData != nil {
		bodyMap["has_email_login_data"] = *p.HasEmailLoginData
	}
	if p.EmailLoginData != nil {
		bodyMap["email_login_data"] = *p.EmailLoginData
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PublishingFastSellResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PublishingFastSellParams struct {
	ProxyId *int64
	Login *string
	ExtendedGuarantee *int64
	RandomProxy *any
	Extra *ExtraModel
	Title *string
	TitleEn *string
	AllowAskDiscount *bool
	Information *string
	LoginPassword *string
	EmailType *string
	Description *string
	Password *string
	HasEmailLoginData *bool
	EmailLoginData *string
}

func (c *Client) PublishingExternal(itemId int64, typeValue string, opts ...func(*PublishingExternalParams)) (map[string]any, error) {
	p := &PublishingExternalParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/external-account"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["type"] = typeValue
	if p.Login != nil {
		bodyMap["login"] = *p.Login
	}
	if p.EmailLoginData != nil {
		bodyMap["email_login_data"] = *p.EmailLoginData
	}
	if p.Cookies != nil {
		bodyMap["cookies"] = *p.Cookies
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

type PublishingExternalParams struct {
	Login *string
	EmailLoginData *string
	Cookies *string
}

func (c *Client) ManagingPublicTag(itemId int64, tagId int64, opts ...func(*ManagingPublicTagParams)) (map[string]any, error) {
	p := &ManagingPublicTagParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/public-tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["tag_id"] = tagId
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

type ManagingPublicTagParams struct {
}

func (c *Client) ManagingPublicUntag(itemId int64, tagId int64, opts ...func(*ManagingPublicUntagParams)) (map[string]any, error) {
	p := &ManagingPublicUntagParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/public-tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["tag_id"] = tagId
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

type ManagingPublicUntagParams struct {
}

func (c *Client) ManagingUnfavorite(itemId int64, opts ...func(*ManagingUnfavoriteParams)) (map[string]any, error) {
	p := &ManagingUnfavoriteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/star"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingUnfavoriteParams struct {
}

func (c *Client) ManagingFavorite(itemId int64, opts ...func(*ManagingFavoriteParams)) (map[string]any, error) {
	p := &ManagingFavoriteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/star"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingFavoriteParams struct {
}

func (c *Client) AutoPaymentsList(opts ...func(*AutoPaymentsListParams)) (*AutoPaymentsListResponse, error) {
	p := &AutoPaymentsListParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/auto-payments"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result AutoPaymentsListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type AutoPaymentsListParams struct {
}

func (c *Client) CategoryEscapeFromTarkov(opts ...func(*CategoryEscapeFromTarkovParams)) (*CategoryEscapeFromTarkovResponse, error) {
	p := &CategoryEscapeFromTarkovParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/escape-from-tarkov"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Region != nil {
		params.Set("region", fmt.Sprintf("%v", *p.Region))
	}
	if p.Version != nil {
		params.Set("version[]", fmt.Sprintf("%v", *p.Version))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.LevelMin != nil {
		params.Set("level_min", fmt.Sprintf("%v", *p.LevelMin))
	}
	if p.LevelMax != nil {
		params.Set("level_max", fmt.Sprintf("%v", *p.LevelMax))
	}
	if p.Pve != nil {
		params.Set("pve", fmt.Sprintf("%v", *p.Pve))
	}
	if p.Side != nil {
		params.Set("side", fmt.Sprintf("%v", *p.Side))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryEscapeFromTarkovResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryEscapeFromTarkovParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Region *string
	Version *[]string
	Reg *int64
	RegPeriod *string
	LevelMin *int64
	LevelMax *int64
	Pve *string
	Side *string
}

func (c *Client) ManagingSteamUpdateValue(itemId int64, opts ...func(*ManagingSteamUpdateValueParams)) (*ManagingSteamUpdateValueResponse, error) {
	p := &ManagingSteamUpdateValueParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/update-inventory"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.All != nil {
		bodyMap["all"] = *p.All
	}
	if p.AppId != nil {
		bodyMap["app_id"] = *p.AppId
	}
	if p.Authorize != nil {
		bodyMap["authorize"] = *p.Authorize
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result ManagingSteamUpdateValueResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingSteamUpdateValueParams struct {
	All *bool
	AppId *int64
	Authorize *bool
}

func (c *Client) PurchasingConfirm(itemId int64, opts ...func(*PurchasingConfirmParams)) (*PurchasingConfirmResponse, error) {
	p := &PurchasingConfirmParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/confirm-buy"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Price != nil {
		bodyMap["price"] = *p.Price
	}
	if p.BalanceId != nil {
		bodyMap["balance_id"] = *p.BalanceId
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PurchasingConfirmResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PurchasingConfirmParams struct {
	Price *int64
	BalanceId *int64
}

func (c *Client) ManagingTelegramCode(itemId int64, opts ...func(*ManagingTelegramCodeParams)) (*ManagingTelegramCodeResponse, error) {
	p := &ManagingTelegramCodeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/telegram-login-code"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingTelegramCodeResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingTelegramCodeParams struct {
}

func (c *Client) ManagingClose(itemId int64, opts ...func(*ManagingCloseParams)) (map[string]any, error) {
	p := &ManagingCloseParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/close"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingCloseParams struct {
}

func (c *Client) PurchasingDiscountRequest(itemId int64, discountPrice float64, opts ...func(*PurchasingDiscountRequestParams)) (map[string]any, error) {
	p := &PurchasingDiscountRequestParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/discount"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	if p.Message != nil {
		bodyMap["message"] = *p.Message
	}
	bodyMap["discount_price"] = discountPrice
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

type PurchasingDiscountRequestParams struct {
	Message *string
}

func (c *Client) PurchasingDiscountCancel(itemId int64, opts ...func(*PurchasingDiscountCancelParams)) (map[string]any, error) {
	p := &PurchasingDiscountCancelParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/discount"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type PurchasingDiscountCancelParams struct {
}

func (c *Client) PaymentsFee(opts ...func(*PaymentsFeeParams)) (*PaymentsFeeResponse, error) {
	p := &PaymentsFeeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/transfer/fee"

	params := url.Values{}
	if p.Amount != nil {
		params.Set("amount", fmt.Sprintf("%v", *p.Amount))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result PaymentsFeeResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsFeeParams struct {
	Amount *float64
}

func (c *Client) CategoryAll(opts ...func(*CategoryAllParams)) (map[string]any, error) {
	p := &CategoryAllParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
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

type CategoryAllParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
}

func (c *Client) CategoryWot(opts ...func(*CategoryWotParams)) (*CategoryWotResponse, error) {
	p := &CategoryWotParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/world-of-tanks"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Tel != nil {
		params.Set("tel", fmt.Sprintf("%v", *p.Tel))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.BattlesMin != nil {
		params.Set("battles_min", fmt.Sprintf("%v", *p.BattlesMin))
	}
	if p.BattlesMax != nil {
		params.Set("battles_max", fmt.Sprintf("%v", *p.BattlesMax))
	}
	if p.GoldMin != nil {
		params.Set("gold_min", fmt.Sprintf("%v", *p.GoldMin))
	}
	if p.GoldMax != nil {
		params.Set("gold_max", fmt.Sprintf("%v", *p.GoldMax))
	}
	if p.SilverMin != nil {
		params.Set("silver_min", fmt.Sprintf("%v", *p.SilverMin))
	}
	if p.SilverMax != nil {
		params.Set("silver_max", fmt.Sprintf("%v", *p.SilverMax))
	}
	if p.TopMin != nil {
		params.Set("top_min", fmt.Sprintf("%v", *p.TopMin))
	}
	if p.TopMax != nil {
		params.Set("top_max", fmt.Sprintf("%v", *p.TopMax))
	}
	if p.PremMin != nil {
		params.Set("prem_min", fmt.Sprintf("%v", *p.PremMin))
	}
	if p.PremMax != nil {
		params.Set("prem_max", fmt.Sprintf("%v", *p.PremMax))
	}
	if p.TopPremMin != nil {
		params.Set("top_prem_min", fmt.Sprintf("%v", *p.TopPremMin))
	}
	if p.TopPremMax != nil {
		params.Set("top_prem_max", fmt.Sprintf("%v", *p.TopPremMax))
	}
	if p.WinPmin != nil {
		params.Set("win_pmin", fmt.Sprintf("%v", *p.WinPmin))
	}
	if p.WinPmax != nil {
		params.Set("win_pmax", fmt.Sprintf("%v", *p.WinPmax))
	}
	if p.Tank != nil {
		params.Set("tank[]", fmt.Sprintf("%v", *p.Tank))
	}
	if p.Region != nil {
		params.Set("region[]", fmt.Sprintf("%v", *p.Region))
	}
	if p.NotRegion != nil {
		params.Set("not_region[]", fmt.Sprintf("%v", *p.NotRegion))
	}
	if p.Premium != nil {
		params.Set("premium", fmt.Sprintf("%v", *p.Premium))
	}
	if p.PremiumExpiration != nil {
		params.Set("premium_expiration", fmt.Sprintf("%v", *p.PremiumExpiration))
	}
	if p.PremiumExpirationPeriod != nil {
		params.Set("premium_expiration_period", fmt.Sprintf("%v", *p.PremiumExpirationPeriod))
	}
	if p.Clan != nil {
		params.Set("clan", fmt.Sprintf("%v", *p.Clan))
	}
	if p.ClanRole != nil {
		params.Set("clan_role[]", fmt.Sprintf("%v", *p.ClanRole))
	}
	if p.NotClanRole != nil {
		params.Set("not_clan_role[]", fmt.Sprintf("%v", *p.NotClanRole))
	}
	if p.ClanGoldMin != nil {
		params.Set("clan_gold_min", fmt.Sprintf("%v", *p.ClanGoldMin))
	}
	if p.ClanGoldMax != nil {
		params.Set("clan_gold_max", fmt.Sprintf("%v", *p.ClanGoldMax))
	}
	if p.ClanCreditsMin != nil {
		params.Set("clan_credits_min", fmt.Sprintf("%v", *p.ClanCreditsMin))
	}
	if p.ClanCreditsMax != nil {
		params.Set("clan_credits_max", fmt.Sprintf("%v", *p.ClanCreditsMax))
	}
	if p.ClanCrystalMin != nil {
		params.Set("clan_crystal_min", fmt.Sprintf("%v", *p.ClanCrystalMin))
	}
	if p.ClanCrystalMax != nil {
		params.Set("clan_crystal_max", fmt.Sprintf("%v", *p.ClanCrystalMax))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryWotResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryWotParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	EmailType *[]string
	ItemDomain *string
	Tel *string
	Daybreak *int64
	BattlesMin *int64
	BattlesMax *int64
	GoldMin *int64
	GoldMax *int64
	SilverMin *int64
	SilverMax *int64
	TopMin *int64
	TopMax *int64
	PremMin *int64
	PremMax *int64
	TopPremMin *int64
	TopPremMax *int64
	WinPmin *int64
	WinPmax *int64
	Tank *[]int64
	Region *[]string
	NotRegion *[]string
	Premium *string
	PremiumExpiration *int64
	PremiumExpirationPeriod *string
	Clan *string
	ClanRole *[]string
	NotClanRole *[]string
	ClanGoldMin *int64
	ClanGoldMax *int64
	ClanCreditsMin *int64
	ClanCreditsMax *int64
	ClanCrystalMin *int64
	ClanCrystalMax *int64
	Country *[]string
	NotCountry *[]string
}

func (c *Client) CategoryMinecraft(opts ...func(*CategoryMinecraftParams)) (*CategoryMinecraftResponse, error) {
	p := &CategoryMinecraftParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/minecraft"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Subscription != nil {
		params.Set("subscription", fmt.Sprintf("%v", *p.Subscription))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}
	if p.Autorenewal != nil {
		params.Set("autorenewal", fmt.Sprintf("%v", *p.Autorenewal))
	}
	if p.Java != nil {
		params.Set("java", fmt.Sprintf("%v", *p.Java))
	}
	if p.Bedrock != nil {
		params.Set("bedrock", fmt.Sprintf("%v", *p.Bedrock))
	}
	if p.Dungeons != nil {
		params.Set("dungeons", fmt.Sprintf("%v", *p.Dungeons))
	}
	if p.Legends != nil {
		params.Set("legends", fmt.Sprintf("%v", *p.Legends))
	}
	if p.ChangeNickname != nil {
		params.Set("change_nickname", fmt.Sprintf("%v", *p.ChangeNickname))
	}
	if p.Capes != nil {
		params.Set("capes[]", fmt.Sprintf("%v", *p.Capes))
	}
	if p.CapesMin != nil {
		params.Set("capes_min", fmt.Sprintf("%v", *p.CapesMin))
	}
	if p.CapesMax != nil {
		params.Set("capes_max", fmt.Sprintf("%v", *p.CapesMax))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.HypixelBan != nil {
		params.Set("hypixel_ban", fmt.Sprintf("%v", *p.HypixelBan))
	}
	if p.HypixelSkyblockApiEnabled != nil {
		params.Set("hypixel_skyblock_api_enabled", fmt.Sprintf("%v", *p.HypixelSkyblockApiEnabled))
	}
	if p.RankHypixel != nil {
		params.Set("rank_hypixel[]", fmt.Sprintf("%v", *p.RankHypixel))
	}
	if p.LevelHypixelMin != nil {
		params.Set("level_hypixel_min", fmt.Sprintf("%v", *p.LevelHypixelMin))
	}
	if p.LevelHypixelMax != nil {
		params.Set("level_hypixel_max", fmt.Sprintf("%v", *p.LevelHypixelMax))
	}
	if p.AchievementHypixelMin != nil {
		params.Set("achievement_hypixel_min", fmt.Sprintf("%v", *p.AchievementHypixelMin))
	}
	if p.AchievementHypixelMax != nil {
		params.Set("achievement_hypixel_max", fmt.Sprintf("%v", *p.AchievementHypixelMax))
	}
	if p.LevelHypixelSkyblockMin != nil {
		params.Set("level_hypixel_skyblock_min", fmt.Sprintf("%v", *p.LevelHypixelSkyblockMin))
	}
	if p.LevelHypixelSkyblockMax != nil {
		params.Set("level_hypixel_skyblock_max", fmt.Sprintf("%v", *p.LevelHypixelSkyblockMax))
	}
	if p.NetWorthHypixelSkyblockMin != nil {
		params.Set("net_worth_hypixel_skyblock_min", fmt.Sprintf("%v", *p.NetWorthHypixelSkyblockMin))
	}
	if p.NetWorthHypixelSkyblockMax != nil {
		params.Set("net_worth_hypixel_skyblock_max", fmt.Sprintf("%v", *p.NetWorthHypixelSkyblockMax))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.LastLoginHypixel != nil {
		params.Set("last_login_hypixel", fmt.Sprintf("%v", *p.LastLoginHypixel))
	}
	if p.LastLoginHypixelPeriod != nil {
		params.Set("last_login_hypixel_period", fmt.Sprintf("%v", *p.LastLoginHypixelPeriod))
	}
	if p.CanChangeDetails != nil {
		params.Set("can_change_details", fmt.Sprintf("%v", *p.CanChangeDetails))
	}
	if p.NicknameLengthMin != nil {
		params.Set("nickname_length_min", fmt.Sprintf("%v", *p.NicknameLengthMin))
	}
	if p.NicknameLengthMax != nil {
		params.Set("nickname_length_max", fmt.Sprintf("%v", *p.NicknameLengthMax))
	}
	if p.HypixelBanParsed != nil {
		params.Set("hypixel_ban_parsed", fmt.Sprintf("%v", *p.HypixelBanParsed))
	}
	if p.MinecoinsMin != nil {
		params.Set("minecoins_min", fmt.Sprintf("%v", *p.MinecoinsMin))
	}
	if p.MinecoinsMax != nil {
		params.Set("minecoins_max", fmt.Sprintf("%v", *p.MinecoinsMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryMinecraftResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryMinecraftParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Subscription *string
	SubscriptionLength *int64
	SubscriptionPeriod *string
	Autorenewal *string
	Java *string
	Bedrock *string
	Dungeons *string
	Legends *string
	ChangeNickname *string
	Capes *[]string
	CapesMin *int64
	CapesMax *int64
	Country *[]string
	NotCountry *[]string
	HypixelBan *string
	HypixelSkyblockApiEnabled *string
	RankHypixel *[]string
	LevelHypixelMin *int64
	LevelHypixelMax *int64
	AchievementHypixelMin *int64
	AchievementHypixelMax *int64
	LevelHypixelSkyblockMin *int64
	LevelHypixelSkyblockMax *int64
	NetWorthHypixelSkyblockMin *int64
	NetWorthHypixelSkyblockMax *int64
	Reg *int64
	RegPeriod *string
	LastLoginHypixel *int64
	LastLoginHypixelPeriod *string
	CanChangeDetails *string
	NicknameLengthMin *int64
	NicknameLengthMax *int64
	HypixelBanParsed *string
	MinecoinsMin *int64
	MinecoinsMax *int64
}

func (c *Client) ManagingTag(itemId int64, tagId int64, opts ...func(*ManagingTagParams)) (map[string]any, error) {
	p := &ManagingTagParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["tag_id"] = tagId
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

type ManagingTagParams struct {
}

func (c *Client) ManagingUntag(itemId int64, tagId int64, opts ...func(*ManagingUntagParams)) (map[string]any, error) {
	p := &ManagingUntagParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	bodyMap := map[string]any{}
	bodyMap["tag_id"] = tagId
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

type ManagingUntagParams struct {
}

func (c *Client) Batch(opts ...func(*BatchParams)) (*BatchResponse, error) {
	p := &BatchParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/batch"


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result BatchResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type BatchParams struct {
}

func (c *Client) CategoryRoblox(opts ...func(*CategoryRobloxParams)) (*CategoryRobloxResponse, error) {
	p := &CategoryRobloxParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/roblox"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Email != nil {
		params.Set("email", fmt.Sprintf("%v", *p.Email))
	}
	if p.RobuxMin != nil {
		params.Set("robux_min", fmt.Sprintf("%v", *p.RobuxMin))
	}
	if p.RobuxMax != nil {
		params.Set("robux_max", fmt.Sprintf("%v", *p.RobuxMax))
	}
	if p.FriendsMin != nil {
		params.Set("friends_min", fmt.Sprintf("%v", *p.FriendsMin))
	}
	if p.FriendsMax != nil {
		params.Set("friends_max", fmt.Sprintf("%v", *p.FriendsMax))
	}
	if p.FollowersMin != nil {
		params.Set("followers_min", fmt.Sprintf("%v", *p.FollowersMin))
	}
	if p.FollowersMax != nil {
		params.Set("followers_max", fmt.Sprintf("%v", *p.FollowersMax))
	}
	if p.Country != nil {
		params.Set("country", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country", fmt.Sprintf("%v", *p.NotCountry))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.Subscription != nil {
		params.Set("subscription", fmt.Sprintf("%v", *p.Subscription))
	}
	if p.SubscriptionLength != nil {
		params.Set("subscription_length", fmt.Sprintf("%v", *p.SubscriptionLength))
	}
	if p.SubscriptionPeriod != nil {
		params.Set("subscription_period", fmt.Sprintf("%v", *p.SubscriptionPeriod))
	}
	if p.Autorenewal != nil {
		params.Set("autorenewal", fmt.Sprintf("%v", *p.Autorenewal))
	}
	if p.XboxConnected != nil {
		params.Set("xbox_connected", fmt.Sprintf("%v", *p.XboxConnected))
	}
	if p.PsnConnected != nil {
		params.Set("psn_connected", fmt.Sprintf("%v", *p.PsnConnected))
	}
	if p.Verified != nil {
		params.Set("verified", fmt.Sprintf("%v", *p.Verified))
	}
	if p.AgeVerified != nil {
		params.Set("age_verified", fmt.Sprintf("%v", *p.AgeVerified))
	}
	if p.IncomingRobuxTotalMin != nil {
		params.Set("incoming_robux_total_min", fmt.Sprintf("%v", *p.IncomingRobuxTotalMin))
	}
	if p.IncomingRobuxTotalMax != nil {
		params.Set("incoming_robux_total_max", fmt.Sprintf("%v", *p.IncomingRobuxTotalMax))
	}
	if p.LimitedPriceMin != nil {
		params.Set("limited_price_min", fmt.Sprintf("%v", *p.LimitedPriceMin))
	}
	if p.LimitedPriceMax != nil {
		params.Set("limited_price_max", fmt.Sprintf("%v", *p.LimitedPriceMax))
	}
	if p.GamepassMin != nil {
		params.Set("gamepass_min", fmt.Sprintf("%v", *p.GamepassMin))
	}
	if p.GamepassMax != nil {
		params.Set("gamepass_max", fmt.Sprintf("%v", *p.GamepassMax))
	}
	if p.GameDonations != nil {
		params.Set("game_donations", fmt.Sprintf("%v", *p.GameDonations))
	}
	if p.InvMin != nil {
		params.Set("inv_min", fmt.Sprintf("%v", *p.InvMin))
	}
	if p.InvMax != nil {
		params.Set("inv_max", fmt.Sprintf("%v", *p.InvMax))
	}
	if p.UgcLimitedPriceMin != nil {
		params.Set("ugc_limited_price_min", fmt.Sprintf("%v", *p.UgcLimitedPriceMin))
	}
	if p.UgcLimitedPriceMax != nil {
		params.Set("ugc_limited_price_max", fmt.Sprintf("%v", *p.UgcLimitedPriceMax))
	}
	if p.CreditBalanceMin != nil {
		params.Set("credit_balance_min", fmt.Sprintf("%v", *p.CreditBalanceMin))
	}
	if p.CreditBalanceMax != nil {
		params.Set("credit_balance_max", fmt.Sprintf("%v", *p.CreditBalanceMax))
	}
	if p.OffsaleMin != nil {
		params.Set("offsale_min", fmt.Sprintf("%v", *p.OffsaleMin))
	}
	if p.OffsaleMax != nil {
		params.Set("offsale_max", fmt.Sprintf("%v", *p.OffsaleMax))
	}
	if p.Voice != nil {
		params.Set("voice", fmt.Sprintf("%v", *p.Voice))
	}
	if p.AgeGroup != nil {
		params.Set("age_group[]", fmt.Sprintf("%v", *p.AgeGroup))
	}
	if p.NotAgeGroup != nil {
		params.Set("not_age_group[]", fmt.Sprintf("%v", *p.NotAgeGroup))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryRobloxResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryRobloxParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Email *string
	RobuxMin *int64
	RobuxMax *int64
	FriendsMin *int64
	FriendsMax *int64
	FollowersMin *int64
	FollowersMax *int64
	Country *[]string
	NotCountry *[]string
	Reg *int64
	RegPeriod *string
	Subscription *string
	SubscriptionLength *int64
	SubscriptionPeriod *string
	Autorenewal *string
	XboxConnected *string
	PsnConnected *string
	Verified *string
	AgeVerified *string
	IncomingRobuxTotalMin *int64
	IncomingRobuxTotalMax *int64
	LimitedPriceMin *int64
	LimitedPriceMax *int64
	GamepassMin *int64
	GamepassMax *int64
	GameDonations *string
	InvMin *int64
	InvMax *int64
	UgcLimitedPriceMin *int64
	UgcLimitedPriceMax *int64
	CreditBalanceMin *int64
	CreditBalanceMax *int64
	OffsaleMin *int64
	OffsaleMax *int64
	Voice *string
	AgeGroup *[]string
	NotAgeGroup *[]string
}

func (c *Client) ManagingSteamMafileCode(itemId int64, opts ...func(*ManagingSteamMafileCodeParams)) (*ConfirmationCodeModel, error) {
	p := &ManagingSteamMafileCodeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/guard-code"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ConfirmationCodeModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingSteamMafileCodeParams struct {
}

func (c *Client) ManagingTelegramResetAuth(itemId int64, opts ...func(*ManagingTelegramResetAuthParams)) (map[string]any, error) {
	p := &ManagingTelegramResetAuthParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/telegram-reset-authorizations"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingTelegramResetAuthParams struct {
}

func (c *Client) PaymentsCancel(paymentId int64, opts ...func(*PaymentsCancelParams)) (map[string]any, error) {
	p := &PaymentsCancelParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/transfer/cancel"


	bodyMap := map[string]any{}
	bodyMap["payment_id"] = paymentId
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

type PaymentsCancelParams struct {
}

func (c *Client) PaymentsHistory(opts ...func(*PaymentsHistoryParams)) (*PaymentsHistoryResponse, error) {
	p := &PaymentsHistoryParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/user/payments"

	params := url.Values{}
	if p.TypeValue != nil {
		params.Set("type", fmt.Sprintf("%v", *p.TypeValue))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.OperationIdLt != nil {
		params.Set("operation_id_lt", fmt.Sprintf("%v", *p.OperationIdLt))
	}
	if p.Receiver != nil {
		params.Set("receiver", fmt.Sprintf("%v", *p.Receiver))
	}
	if p.Sender != nil {
		params.Set("sender", fmt.Sprintf("%v", *p.Sender))
	}
	if p.IsApi != nil {
		params.Set("is_api", fmt.Sprintf("%v", *p.IsApi))
	}
	if p.StartDate != nil {
		params.Set("startDate", fmt.Sprintf("%v", *p.StartDate))
	}
	if p.EndDate != nil {
		params.Set("endDate", fmt.Sprintf("%v", *p.EndDate))
	}
	if p.Wallet != nil {
		params.Set("wallet", fmt.Sprintf("%v", *p.Wallet))
	}
	if p.Comment != nil {
		params.Set("comment", fmt.Sprintf("%v", *p.Comment))
	}
	if p.IsHold != nil {
		params.Set("is_hold", fmt.Sprintf("%v", *p.IsHold))
	}
	if p.ShowPaymentStats != nil {
		params.Set("show_payment_stats", fmt.Sprintf("%v", *p.ShowPaymentStats))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result PaymentsHistoryResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsHistoryParams struct {
	TypeValue *string
	Pmin *int64
	Pmax *int64
	Currency *string
	Page *int64
	OperationIdLt *int64
	Receiver *string
	Sender *string
	IsApi *bool
	StartDate *string
	EndDate *string
	Wallet *string
	Comment *string
	IsHold *bool
	ShowPaymentStats *bool
}

func (c *Client) ListOrders(opts ...func(*ListOrdersParams)) (*ItemListModel, error) {
	p := &ListOrdersParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/user/orders"

	params := url.Values{}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.CategoryId != nil {
		params.Set("category_id", fmt.Sprintf("%v", *p.CategoryId))
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Show != nil {
		params.Set("show", fmt.Sprintf("%v", *p.Show))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Login != nil {
		params.Set("login", fmt.Sprintf("%v", *p.Login))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ListOrdersParams struct {
	UserId *int64
	CategoryId *int64
	Page *int64
	Show *string
	Title *string
	Pmin *int64
	Pmax *int64
	Login *string
	Origin *[]string
	NotOrigin *[]string
	OrderBy *string
	Sb *bool
	SbByMe *bool
	Nsb *bool
	NsbByMe *bool
}

func (c *Client) ManagingOpen(itemId int64, opts ...func(*ManagingOpenParams)) (map[string]any, error) {
	p := &ManagingOpenParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/open"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


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

type ManagingOpenParams struct {
}

func (c *Client) PaymentsTransfer(amount int64, currency string, opts ...func(*PaymentsTransferParams)) (map[string]any, error) {
	p := &PaymentsTransferParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/transfer"


	bodyMap := map[string]any{}
	if p.UserId != nil {
		bodyMap["user_id"] = *p.UserId
	}
	bodyMap["amount"] = amount
	if p.Comment != nil {
		bodyMap["comment"] = *p.Comment
	}
	if p.TelegramDeal != nil {
		bodyMap["telegram_deal"] = *p.TelegramDeal
	}
	if p.TransferHold != nil {
		bodyMap["transfer_hold"] = *p.TransferHold
	}
	if p.HoldLengthValue != nil {
		bodyMap["hold_length_value"] = *p.HoldLengthValue
	}
	if p.Username != nil {
		bodyMap["username"] = *p.Username
	}
	bodyMap["currency"] = currency
	if p.TelegramUsername != nil {
		bodyMap["telegram_username"] = *p.TelegramUsername
	}
	if p.HoldLengthOption != nil {
		bodyMap["hold_length_option"] = *p.HoldLengthOption
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

type PaymentsTransferParams struct {
	UserId *int64
	Comment *string
	TelegramDeal *bool
	TransferHold *bool
	HoldLengthValue *int64
	Username *string
	TelegramUsername *string
	HoldLengthOption *string
}

func (c *Client) AutoPaymentsCreate(amount float64, usernameReceiver string, day int64, opts ...func(*AutoPaymentsCreateParams)) (*AutoPaymentsCreateResponse, error) {
	p := &AutoPaymentsCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/auto-payment"


	bodyMap := map[string]any{}
	bodyMap["amount"] = amount
	if p.Currency != nil {
		bodyMap["currency"] = *p.Currency
	}
	if p.Description != nil {
		bodyMap["description"] = *p.Description
	}
	if p.SecretAnswer != nil {
		bodyMap["secret_answer"] = *p.SecretAnswer
	}
	bodyMap["username_receiver"] = usernameReceiver
	bodyMap["day"] = day
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result AutoPaymentsCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type AutoPaymentsCreateParams struct {
	Currency *string
	Description *string
	SecretAnswer *string
}

func (c *Client) AutoPaymentsDelete(autoPaymentId int64, opts ...func(*AutoPaymentsDeleteParams)) (map[string]any, error) {
	p := &AutoPaymentsDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/auto-payment"


	bodyMap := map[string]any{}
	bodyMap["auto_payment_id"] = autoPaymentId
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

type AutoPaymentsDeleteParams struct {
}

func (c *Client) PaymentsPayout(paymentSystem string, wallet string, amount float64, currency string, opts ...func(*PaymentsPayoutParams)) (map[string]any, error) {
	p := &PaymentsPayoutParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/payout"


	bodyMap := map[string]any{}
	bodyMap["payment_system"] = paymentSystem
	bodyMap["wallet"] = wallet
	bodyMap["amount"] = amount
	bodyMap["currency"] = currency
	if p.IncludeFee != nil {
		bodyMap["include_fee"] = *p.IncludeFee
	}
	if p.Extra != nil {
		bodyMap["extra"] = *p.Extra
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

type PaymentsPayoutParams struct {
	IncludeFee *bool
	Extra *map[string]any
}

func (c *Client) CategoryFortnite(opts ...func(*CategoryFortniteParams)) (*CategoryFortniteResponse, error) {
	p := &CategoryFortniteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/fortnite"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.EmailType != nil {
		params.Set("email_type[]", fmt.Sprintf("%v", *p.EmailType))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.TempEmail != nil {
		params.Set("temp_email", fmt.Sprintf("%v", *p.TempEmail))
	}
	if p.ItemDomain != nil {
		params.Set("item_domain", fmt.Sprintf("%v", *p.ItemDomain))
	}
	if p.Eg != nil {
		params.Set("eg", fmt.Sprintf("%v", *p.Eg))
	}
	if p.Smin != nil {
		params.Set("smin", fmt.Sprintf("%v", *p.Smin))
	}
	if p.Smax != nil {
		params.Set("smax", fmt.Sprintf("%v", *p.Smax))
	}
	if p.Vbmin != nil {
		params.Set("vbmin", fmt.Sprintf("%v", *p.Vbmin))
	}
	if p.Vbmax != nil {
		params.Set("vbmax", fmt.Sprintf("%v", *p.Vbmax))
	}
	if p.Skin != nil {
		params.Set("skin[]", fmt.Sprintf("%v", *p.Skin))
	}
	if p.Pickaxe != nil {
		params.Set("pickaxe[]", fmt.Sprintf("%v", *p.Pickaxe))
	}
	if p.Glider != nil {
		params.Set("glider[]", fmt.Sprintf("%v", *p.Glider))
	}
	if p.Dance != nil {
		params.Set("dance[]", fmt.Sprintf("%v", *p.Dance))
	}
	if p.ChangeEmail != nil {
		params.Set("change_email", fmt.Sprintf("%v", *p.ChangeEmail))
	}
	if p.Platform != nil {
		params.Set("platform[]", fmt.Sprintf("%v", *p.Platform))
	}
	if p.SkinsShopMin != nil {
		params.Set("skins_shop_min", fmt.Sprintf("%v", *p.SkinsShopMin))
	}
	if p.SkinsShopMax != nil {
		params.Set("skins_shop_max", fmt.Sprintf("%v", *p.SkinsShopMax))
	}
	if p.PickaxesShopMin != nil {
		params.Set("pickaxes_shop_min", fmt.Sprintf("%v", *p.PickaxesShopMin))
	}
	if p.PickaxesShopMax != nil {
		params.Set("pickaxes_shop_max", fmt.Sprintf("%v", *p.PickaxesShopMax))
	}
	if p.DancesShopMin != nil {
		params.Set("dances_shop_min", fmt.Sprintf("%v", *p.DancesShopMin))
	}
	if p.DancesShopMax != nil {
		params.Set("dances_shop_max", fmt.Sprintf("%v", *p.DancesShopMax))
	}
	if p.GlidersShopMin != nil {
		params.Set("gliders_shop_min", fmt.Sprintf("%v", *p.GlidersShopMin))
	}
	if p.GlidersShopMax != nil {
		params.Set("gliders_shop_max", fmt.Sprintf("%v", *p.GlidersShopMax))
	}
	if p.SkinsShopVbmin != nil {
		params.Set("skins_shop_vbmin", fmt.Sprintf("%v", *p.SkinsShopVbmin))
	}
	if p.SkinsShopVbmax != nil {
		params.Set("skins_shop_vbmax", fmt.Sprintf("%v", *p.SkinsShopVbmax))
	}
	if p.PickaxesShopVbmin != nil {
		params.Set("pickaxes_shop_vbmin", fmt.Sprintf("%v", *p.PickaxesShopVbmin))
	}
	if p.PickaxesShopVbmax != nil {
		params.Set("pickaxes_shop_vbmax", fmt.Sprintf("%v", *p.PickaxesShopVbmax))
	}
	if p.DancesShopVbmin != nil {
		params.Set("dances_shop_vbmin", fmt.Sprintf("%v", *p.DancesShopVbmin))
	}
	if p.DancesShopVbmax != nil {
		params.Set("dances_shop_vbmax", fmt.Sprintf("%v", *p.DancesShopVbmax))
	}
	if p.GlidersShopVbmin != nil {
		params.Set("gliders_shop_vbmin", fmt.Sprintf("%v", *p.GlidersShopVbmin))
	}
	if p.GlidersShopVbmax != nil {
		params.Set("gliders_shop_vbmax", fmt.Sprintf("%v", *p.GlidersShopVbmax))
	}
	if p.Bp != nil {
		params.Set("bp", fmt.Sprintf("%v", *p.Bp))
	}
	if p.Lmin != nil {
		params.Set("lmin", fmt.Sprintf("%v", *p.Lmin))
	}
	if p.Lmax != nil {
		params.Set("lmax", fmt.Sprintf("%v", *p.Lmax))
	}
	if p.BpLmin != nil {
		params.Set("bp_lmin", fmt.Sprintf("%v", *p.BpLmin))
	}
	if p.BpLmax != nil {
		params.Set("bp_lmax", fmt.Sprintf("%v", *p.BpLmax))
	}
	if p.LastTransDate != nil {
		params.Set("last_trans_date", fmt.Sprintf("%v", *p.LastTransDate))
	}
	if p.LastTransDatePeriod != nil {
		params.Set("last_trans_date_period", fmt.Sprintf("%v", *p.LastTransDatePeriod))
	}
	if p.NoTrans != nil {
		params.Set("no_trans", fmt.Sprintf("%v", *p.NoTrans))
	}
	if p.XboxLinkable != nil {
		params.Set("xbox_linkable", fmt.Sprintf("%v", *p.XboxLinkable))
	}
	if p.PsnLinkable != nil {
		params.Set("psn_linkable", fmt.Sprintf("%v", *p.PsnLinkable))
	}
	if p.Daybreak != nil {
		params.Set("daybreak", fmt.Sprintf("%v", *p.Daybreak))
	}
	if p.RlPurchases != nil {
		params.Set("rl_purchases", fmt.Sprintf("%v", *p.RlPurchases))
	}
	if p.Reg != nil {
		params.Set("reg", fmt.Sprintf("%v", *p.Reg))
	}
	if p.RegPeriod != nil {
		params.Set("reg_period", fmt.Sprintf("%v", *p.RegPeriod))
	}
	if p.RefundCreditsMin != nil {
		params.Set("refund_credits_min", fmt.Sprintf("%v", *p.RefundCreditsMin))
	}
	if p.RefundCreditsMax != nil {
		params.Set("refund_credits_max", fmt.Sprintf("%v", *p.RefundCreditsMax))
	}
	if p.PickaxeMin != nil {
		params.Set("pickaxe_min", fmt.Sprintf("%v", *p.PickaxeMin))
	}
	if p.PickaxeMax != nil {
		params.Set("pickaxe_max", fmt.Sprintf("%v", *p.PickaxeMax))
	}
	if p.Dmin != nil {
		params.Set("dmin", fmt.Sprintf("%v", *p.Dmin))
	}
	if p.Dmax != nil {
		params.Set("dmax", fmt.Sprintf("%v", *p.Dmax))
	}
	if p.Gmin != nil {
		params.Set("gmin", fmt.Sprintf("%v", *p.Gmin))
	}
	if p.Gmax != nil {
		params.Set("gmax", fmt.Sprintf("%v", *p.Gmax))
	}
	if p.Country != nil {
		params.Set("country[]", fmt.Sprintf("%v", *p.Country))
	}
	if p.NotCountry != nil {
		params.Set("not_country[]", fmt.Sprintf("%v", *p.NotCountry))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryFortniteResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryFortniteParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	EmailType *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	TempEmail *string
	ItemDomain *string
	Eg *int64
	Smin *int64
	Smax *int64
	Vbmin *int64
	Vbmax *int64
	Skin *[]string
	Pickaxe *[]string
	Glider *[]string
	Dance *[]string
	ChangeEmail *string
	Platform *[]string
	SkinsShopMin *int64
	SkinsShopMax *int64
	PickaxesShopMin *int64
	PickaxesShopMax *int64
	DancesShopMin *int64
	DancesShopMax *int64
	GlidersShopMin *int64
	GlidersShopMax *int64
	SkinsShopVbmin *int64
	SkinsShopVbmax *int64
	PickaxesShopVbmin *int64
	PickaxesShopVbmax *int64
	DancesShopVbmin *int64
	DancesShopVbmax *int64
	GlidersShopVbmin *int64
	GlidersShopVbmax *int64
	Bp *string
	Lmin *int64
	Lmax *int64
	BpLmin *int64
	BpLmax *int64
	LastTransDate *int64
	LastTransDatePeriod *string
	NoTrans *bool
	XboxLinkable *string
	PsnLinkable *string
	Daybreak *int64
	RlPurchases *bool
	Reg *int64
	RegPeriod *string
	RefundCreditsMin *int64
	RefundCreditsMax *int64
	PickaxeMin *int64
	PickaxeMax *int64
	Dmin *int64
	Dmax *int64
	Gmin *int64
	Gmax *int64
	Country *[]string
	NotCountry *[]string
}

func (c *Client) CategoryHytale(opts ...func(*CategoryHytaleParams)) (*CategoryHytaleResponse, error) {
	p := &CategoryHytaleParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/hytale"

	params := url.Values{}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%v", *p.Page))
	}
	if p.Pmin != nil {
		params.Set("pmin", fmt.Sprintf("%v", *p.Pmin))
	}
	if p.Pmax != nil {
		params.Set("pmax", fmt.Sprintf("%v", *p.Pmax))
	}
	if p.Title != nil {
		params.Set("title", fmt.Sprintf("%v", *p.Title))
	}
	if p.OrderBy != nil {
		params.Set("order_by", fmt.Sprintf("%v", *p.OrderBy))
	}
	if p.TagId != nil {
		params.Set("tag_id[]", fmt.Sprintf("%v", *p.TagId))
	}
	if p.NotTagId != nil {
		params.Set("not_tag_id[]", fmt.Sprintf("%v", *p.NotTagId))
	}
	if p.PublicTagId != nil {
		params.Set("public_tag_id[]", fmt.Sprintf("%v", *p.PublicTagId))
	}
	if p.NotPublicTagId != nil {
		params.Set("not_public_tag_id[]", fmt.Sprintf("%v", *p.NotPublicTagId))
	}
	if p.Origin != nil {
		params.Set("origin[]", fmt.Sprintf("%v", *p.Origin))
	}
	if p.NotOrigin != nil {
		params.Set("not_origin[]", fmt.Sprintf("%v", *p.NotOrigin))
	}
	if p.UserId != nil {
		params.Set("user_id", fmt.Sprintf("%v", *p.UserId))
	}
	if p.Nsb != nil {
		params.Set("nsb", fmt.Sprintf("%v", *p.Nsb))
	}
	if p.Sb != nil {
		params.Set("sb", fmt.Sprintf("%v", *p.Sb))
	}
	if p.NsbByMe != nil {
		params.Set("nsb_by_me", fmt.Sprintf("%v", *p.NsbByMe))
	}
	if p.SbByMe != nil {
		params.Set("sb_by_me", fmt.Sprintf("%v", *p.SbByMe))
	}
	if p.Currency != nil {
		params.Set("currency", fmt.Sprintf("%v", *p.Currency))
	}
	if p.EmailLoginData != nil {
		params.Set("email_login_data", fmt.Sprintf("%v", *p.EmailLoginData))
	}
	if p.EmailProvider != nil {
		params.Set("email_provider[]", fmt.Sprintf("%v", *p.EmailProvider))
	}
	if p.NotEmailProvider != nil {
		params.Set("not_email_provider[]", fmt.Sprintf("%v", *p.NotEmailProvider))
	}
	if p.ParseSameItemIds != nil {
		params.Set("parse_same_item_ids", fmt.Sprintf("%v", *p.ParseSameItemIds))
	}
	if p.Edition != nil {
		params.Set("edition[]", fmt.Sprintf("%v", *p.Edition))
	}
	if p.ProfilesMin != nil {
		params.Set("profiles_min", fmt.Sprintf("%v", *p.ProfilesMin))
	}
	if p.ProfilesMax != nil {
		params.Set("profiles_max", fmt.Sprintf("%v", *p.ProfilesMax))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result CategoryHytaleResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CategoryHytaleParams struct {
	Page *int64
	Pmin *int64
	Pmax *int64
	Title *string
	OrderBy *string
	TagId *[]int64
	NotTagId *[]int64
	PublicTagId *[]int64
	NotPublicTagId *[]int64
	Origin *[]string
	NotOrigin *[]string
	UserId *int64
	Nsb *bool
	Sb *bool
	NsbByMe *bool
	SbByMe *bool
	Currency *string
	EmailLoginData *bool
	EmailProvider *[]string
	NotEmailProvider *string
	ParseSameItemIds *bool
	Edition *[]string
	ProfilesMin *int64
	ProfilesMax *int64
}

func (c *Client) ManagingCheckGuarantee(itemId int64, opts ...func(*ManagingCheckGuaranteeParams)) (*ManagingCheckGuaranteeResponse, error) {
	p := &ManagingCheckGuaranteeParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/{item_id}/check-guarantee"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%v", itemId), 1)


	raw, err := c.DoRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ManagingCheckGuaranteeResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ManagingCheckGuaranteeParams struct {
}

func (c *Client) PaymentsInvoiceGet(opts ...func(*PaymentsInvoiceGetParams)) (*PaymentsInvoiceGetResponse, error) {
	p := &PaymentsInvoiceGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/invoice"

	params := url.Values{}
	if p.InvoiceId != nil {
		params.Set("invoice_id", fmt.Sprintf("%v", *p.InvoiceId))
	}
	if p.PaymentId != nil {
		params.Set("payment_id", fmt.Sprintf("%v", *p.PaymentId))
	}

	raw, err := c.DoRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var result PaymentsInvoiceGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsInvoiceGetParams struct {
	InvoiceId *int64
	PaymentId *string
}

func (c *Client) PaymentsInvoiceCreate(paymentId string, comment string, merchantId int64, currency string, urlSuccess string, amount float64, opts ...func(*PaymentsInvoiceCreateParams)) (*PaymentsInvoiceCreateResponse, error) {
	p := &PaymentsInvoiceCreateParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/invoice"


	bodyMap := map[string]any{}
	bodyMap["payment_id"] = paymentId
	bodyMap["comment"] = comment
	if p.UrlCallback != nil {
		bodyMap["url_callback"] = *p.UrlCallback
	}
	bodyMap["merchant_id"] = merchantId
	if p.Lifetime != nil {
		bodyMap["lifetime"] = *p.Lifetime
	}
	if p.AdditionalData != nil {
		bodyMap["additional_data"] = *p.AdditionalData
	}
	if p.IsTest != nil {
		bodyMap["is_test"] = *p.IsTest
	}
	bodyMap["currency"] = currency
	bodyMap["url_success"] = urlSuccess
	if p.RequiredTelegramId != nil {
		bodyMap["required_telegram_id"] = *p.RequiredTelegramId
	}
	if p.RequiredTelegramUsername != nil {
		bodyMap["required_telegram_username"] = *p.RequiredTelegramUsername
	}
	bodyMap["amount"] = amount
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	raw, err := c.DoRequest("POST", path, nil, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}

	var result PaymentsInvoiceCreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsInvoiceCreateParams struct {
	UrlCallback *string
	Lifetime *float64
	AdditionalData *string
	IsTest *bool
	RequiredTelegramId *int64
	RequiredTelegramUsername *string
}

func (c *Client) PaymentsPayoutServices(opts ...func(*PaymentsPayoutServicesParams)) (*PaymentsPayoutServicesResponse, error) {
	p := &PaymentsPayoutServicesParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/balance/payout/services"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result PaymentsPayoutServicesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PaymentsPayoutServicesParams struct {
}

func (c *Client) ProxyGet(opts ...func(*ProxyGetParams)) (*ProxyGetResponse, error) {
	p := &ProxyGetParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/proxy"


	raw, err := c.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result ProxyGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ProxyGetParams struct {
}

func (c *Client) ProxyAdd(opts ...func(*ProxyAddParams)) (map[string]any, error) {
	p := &ProxyAddParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/proxy"


	bodyMap := map[string]any{}
	if p.ProxyRow != nil {
		bodyMap["proxy_row"] = *p.ProxyRow
	}
	if p.ProxyIp != nil {
		bodyMap["proxy_ip"] = *p.ProxyIp
	}
	if p.ProxyPort != nil {
		bodyMap["proxy_port"] = *p.ProxyPort
	}
	if p.ProxyUser != nil {
		bodyMap["proxy_user"] = *p.ProxyUser
	}
	if p.ProxyPass != nil {
		bodyMap["proxy_pass"] = *p.ProxyPass
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

type ProxyAddParams struct {
	ProxyRow *string
	ProxyIp *string
	ProxyPort *int64
	ProxyUser *string
	ProxyPass *string
}

func (c *Client) ProxyDelete(opts ...func(*ProxyDeleteParams)) (map[string]any, error) {
	p := &ProxyDeleteParams{}
	for _, opt := range opts {
		opt(p)
	}

	path := "/proxy"


	bodyMap := map[string]any{}
	if p.DeleteAll != nil {
		bodyMap["delete_all"] = *p.DeleteAll
	}
	if p.ProxyId != nil {
		bodyMap["proxy_id"] = *p.ProxyId
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

type ProxyDeleteParams struct {
	DeleteAll *bool
	ProxyId *int64
}
