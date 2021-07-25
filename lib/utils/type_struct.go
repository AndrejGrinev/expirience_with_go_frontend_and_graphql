package utils

import (
	"html/template"
)
// Конфиг
type Config struct {
	Listen        string `long:"listen"         default:":8080"                description:"Addr and port which server listens at"`
	CType         string `long:"ctype"          default:"text/html"            description:"content type"`
	Enc           string `long:"enc"            default:"charset=UTF-8"        description:"кодировка"`
	Lang          string `long:"lang"           default:"ru"                   description:"язык"`
	UriApi        string `long:"uri_api"        default:"//www.local.test"     description:"Адрес api"`
	UriStatic     string `long:"uri_static"     default:"//www.local.test"     description:"Адрес статики"`
	HostContracts string `long:"host_contracts" default:"//dev.iac.tender.pro" description:"Ссылка на обновлённый интерфейс"`
	PgUser        string `long:"pguser"         default:"usetender"            description:"Postgres user"`
	PgHost        string `long:"pghost"         default:"localhost"            description:"Postgres host"`
	PgPort        string `long:"pgport"         default:"5499"                 description:"Postgres port"`
	PgPassword    string `long:"pgpassword"     default:""                     description:"Postgres password"`
	PgDbName      string `long:"pgdbname"       default:"usetender"            description:"Postgres database"`
	IncludesPath  string `long:"includes_path"  default:"./www/includes/"      description:"Путь к инклюдам"`
	StaffPath     string `long:"staff_path"     default:"./www/tmpl/staff/"    description:"Путь к шаблонам staff"`
}
// Cессия
type Adsession_result struct {
	Sid        string `json:"sid"`
	Face_id    int    `json:"face_id"`
	Face_type  int    `json:"face_type"`
	Face_name  string `json:"face_name"`
	Company_id int    `json:"company_id"`
//	Addon        map[string]interface{}
}

type Data_Adsession_result struct {
	Adsession Adsession_result `json:"adsession"`
}

type Graphql_Data_Adsession_result struct {
	Data Data_Adsession_result `json:"data"`
}
// Начальные данные для moder_tender_company_invite
type ModerInvite_result struct {
	Id                   string `json:"id"`
	Tp                   int    `json:"type"`
	IsAuction            bool   `json:"isAuction"`
	State                int    `json:"state"`
	Title                string `json:"title"`
	OrgCompanyId         int    `json:"org_companyid"`
	OrgTesting           int    `json:"org_testing"`
	CompanyCountryId     int    `json:"company_country_id"`
	CompanyRegionId      int    `json:"company_region_id"`
	CompanyAreaId        int    `json:"company_area_id"`
	CompanyCountryName   string `json:"company_country_name"`
	CompanyRegionName    string `json:"company_region_name"`
	CompanyAreaName      string `json:"company_area_name"`
	TenderCountryId      int    `json:"tender_country_id"`
	TenderRegionId       int    `json:"tender_region_id"`
	TenderAreaId         int    `json:"tender_area_id"`
	TenderCountryName    string `json:"tender_country_name"`
	TenderRegionName     string `json:"tender_region_name"`
	TenderAreaName       string `json:"tender_area_name"`
	Dt                   string `json:"dt"`
	CountRepeatAll       int    `json:"count_repeat_all"`
	CompanyLgtInvite     int    `json:"company_lgt_invite"`
	CountLgtInvite       int    `json:"count_lgt_invite"`
	TenderInviteAdmin    bool   `json:"tender_invite_admin"`
	AccessOrgToMarketing bool   `json:"access_org_to_marketing"`
}

type Data_ModerInvite_result struct {
	ModerInvite ModerInvite_result `json:"moderInvite"`
}

type Graphql_Data_ModerInvite_result struct {
	Data Data_ModerInvite_result `json:"data"`
}

// Переменные для хидера и футера & ModerInviteHandler
type ModerInviteVars struct {
	Sid          int
	Tenderid     int
	TenderTitle  string
	PlaceCode    string
	AccMarketing int

	Ctype         string
	Enc           string
	Lang          string
	UriApi        string
	UriStatic     string
	HostContracts string

	SessionName     string
	SessionFaceId   int
	SessionFaceType int

	Title  string
	PageId int
	Year   int

	State                                  int
	Mode                                   string
	RandomStr                              string
	CountRepeatAll                         int
	On                                     int
	UsedLgtInvite                          int
	CountLgtInvite                         string
	PromoKod                               string
	MarketingInviteDisplay                 bool
	ReportFrom                             template.HTML
	ReportWhere                            template.HTML
	OrgCompanyid                           int
	CompanyAddressDiffersFromTenderAddress int
	HrefPageSidTenderid                    string
	TenderCountryId                        string
	TenderCountryName                      string
	TenderRegionId                         int
	TenderRegionName                       string
	PlaceCodeEqRTenderRegionId             bool
	TenderAreaId                           int
	TenderAreaName                         string
	PlaceCodeEqATenderAreaId               bool
	CompanyCountryId                       string
	CompanyCountryName                     string
	CompanyRegionId                        int
	CompanyRegionName                      string
	PlaceCodeEqRCompanyRegionId            bool
	CompanyAreaId                          int
	CompanyAreaName                        string
	PlaceCodeEqACompanyAreaId              bool
}
