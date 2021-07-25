package handlers

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"tpro/stage1g/utils"
	"github.com/graphql-go/graphql"
	"log"
	"encoding/json"
	"tpro/stage1g/gql"
)
// Обработчик страницы moder_tender_company_invite
func ModerInviteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var errStaff, on, countLgtInvite, sesType, companyAddressDiffersFromTenderAddress int
	var msgErrStaff, mode, countLgtInviteStr string
	var reportByTender, reportFrom, reportWhere template.HTML
	var marketingInviteDisplay bool
	// получим аргументы страницы и ip
	paramSid, _ := strconv.Atoi(r.URL.Query().Get("sid"))
	paramTenderid, _ := strconv.Atoi(r.URL.Query().Get("tenderid"))
	paramPlaceCode := r.URL.Query().Get("place_code")
	paramAccMarketing, _ := strconv.Atoi(r.URL.Query().Get("acc_marketing"))
	paramUserip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		return
	}
	paramSidStr := strconv.Itoa(paramSid)
	paramTenderidStr := strconv.Itoa(paramTenderid)
	// если сид не указан, то ошибка
	if len(paramSidStr) == 0 {
		fmt.Println("Отсутствует сид:", paramUserip)
		http.Error(w, "Отсутствует сид.", http.StatusFound)
		return
	}
	// получаем пул соединений
	utils.Db, err = utils.SetupDb()
	if err != nil {
		fmt.Println("Ошибка соединения к базе:", err)
		http.Error(w, "Ошибка соединения к базе.", http.StatusFound)
		return
	}
	// NewSchema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: gql.QueryType})
	if err != nil {
		log.Fatal(err)
	}
	// Query, запрашиваем данные сессии
	query := `{adsession(sid: ` + paramSidStr + `, userip:"` + paramUserip + `") {sid face_id face_type face_name company_id}}`
	params := graphql.Params{Schema: schema, RequestString: query}
	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		http.Error(w, "Ошибка запроса adsession.", http.StatusFound)
	}
	rJSON, _ := json.Marshal(result)
	adsession_result := utils.Graphql_Data_Adsession_result{}
	json.Unmarshal([]byte(rJSON), &adsession_result)
	// проверяем доступ sid к стаффу
	errStaff, msgErrStaff = utils.StaffAcl(adsession_result)
	if errStaff > 0 {
		if errStaff != 2 || paramAccMarketing != 1 {
			fmt.Println("Ошибка доступа к стаффу:", msgErrStaff)
			http.Error(w, msgErrStaff, http.StatusFound)
			return
		}
	}
	// Query, запрашиваем данные moder_tender_company_invite
	query2 := `{moderInvite(id: ` + paramTenderidStr + `) {type state title org_companyid org_testing company_country_id company_region_id company_area_id company_country_name company_region_name company_area_name tender_country_id tender_region_id tender_area_id tender_country_name tender_region_name tender_area_name dt count_repeat_all company_lgt_invite count_lgt_invite tender_invite_admin access_org_to_marketing}}`
	params2 := graphql.Params{Schema: schema, RequestString: query2}
	result2 := graphql.Do(params2)
	if len(result2.Errors) > 0 {
		http.Error(w, "Ошибка запроса moderInvite.", http.StatusFound)
	}
	rJSON2, _ := json.Marshal(result2)
	moderInvite_result := utils.Graphql_Data_ModerInvite_result{}
	json.Unmarshal([]byte(rJSON2), &moderInvite_result)
	defer utils.Db.Close()
	if paramAccMarketing == 1 {
		if moderInvite_result.Data.ModerInvite.AccessOrgToMarketing && moderInvite_result.Data.ModerInvite.OrgCompanyId == adsession_result.Data.Adsession.Company_id {
			sesType = 336
		} else {
			fmt.Println("Доступ запрещён:", paramSid)
			http.Error(w, "Доступ запрещён.", http.StatusFound)
			return
		}
	} else {
		sesType = adsession_result.Data.Adsession.Face_type
	}
	// определяем недостающие переменные
	if (moderInvite_result.Data.ModerInvite.Tp == 2) || (moderInvite_result.Data.ModerInvite.Tp == 8) {
		mode = "sell"
	} else {
		mode = "buy"
	}
	if !utils.ValidPlaceCode.MatchString(paramPlaceCode) {
		paramPlaceCode = ""
	}
	if len(paramPlaceCode) > 0 {
		reportByTender = template.HTML(`stat.get_marketing_report_by_tender(` + paramTenderidStr + `, '` + paramPlaceCode + `')`)
		marketingInviteDisplay = true
	} else {
		reportByTender = template.HTML(`stat.get_marketing_report_by_tender(` + paramTenderidStr + `)`)
		marketingInviteDisplay = false
	}
	reportFrom = template.HTML(reportByTender + ` mr, company c, mv.company mc, an_city_qualifier acq, an_city ac, company_bt cbt`)

	reportWhere = template.HTML(`c.id = mr.company_id AND mc.id = c.id AND acq.id = c.cityid AND ac.id = acq.city and c.blocked = 0 AND c.id <> ` + strconv.Itoa(moderInvite_result.Data.ModerInvite.OrgCompanyId) + ` AND cbt.id = c.id AND ((` + strconv.Itoa(moderInvite_result.Data.ModerInvite.OrgTesting) + ` = 1 AND cbt.status_id = 5) OR (` + strconv.Itoa(moderInvite_result.Data.ModerInvite.OrgTesting) + ` = 0 AND cbt.status_id IN (2, 3, 4)))`)

	if (moderInvite_result.Data.ModerInvite.CompanyCountryId != moderInvite_result.Data.ModerInvite.TenderCountryId) || (moderInvite_result.Data.ModerInvite.CompanyRegionId != moderInvite_result.Data.ModerInvite.TenderRegionId) || (moderInvite_result.Data.ModerInvite.CompanyAreaId != moderInvite_result.Data.ModerInvite.TenderAreaId) {
		companyAddressDiffersFromTenderAddress = 1
	} else {
		companyAddressDiffersFromTenderAddress = 0
	}

	if moderInvite_result.Data.ModerInvite.TenderInviteAdmin {
		on = 1
		countLgtInvite = moderInvite_result.Data.ModerInvite.CompanyLgtInvite - moderInvite_result.Data.ModerInvite.CountLgtInvite
		if countLgtInvite < 0 {
			countLgtInvite = 0
		}
		countLgtInviteStr = strconv.Itoa(countLgtInvite)
		if moderInvite_result.Data.ModerInvite.CompanyLgtInvite == 99999 {
			countLgtInviteStr = "Без ограничений"
		}
	} else {
		on = 0
	}
	// кладем их в структуру
	data := utils.ModerInviteVars{
		// сначала данные для футера и хидера
		Sid:          paramSid,
		Tenderid:     paramTenderid,
		TenderTitle:  moderInvite_result.Data.ModerInvite.Title,
		PlaceCode:    paramPlaceCode,
		AccMarketing: paramAccMarketing,

		Ctype:         utils.Cfg.CType,
		Enc:           utils.Cfg.Enc,
		Lang:          utils.Cfg.Lang,
		UriApi:        utils.Cfg.UriApi,
		UriStatic:     utils.Cfg.UriStatic,
		HostContracts: utils.Cfg.HostContracts,

		SessionName:     adsession_result.Data.Adsession.Face_name,
		SessionFaceId:   adsession_result.Data.Adsession.Face_id,
		SessionFaceType: sesType,

		Title:  "Основная страница маркетинга",
		PageId: 1,
		Year:   time.Now().Year(),
		// теперь данные для страницы
		State:                                  moderInvite_result.Data.ModerInvite.State,
		Mode:                                   mode,
		RandomStr:                              utils.RandSeq(7),
		CountRepeatAll:                         moderInvite_result.Data.ModerInvite.CountRepeatAll,
		On:                                     on,
		UsedLgtInvite:                          moderInvite_result.Data.ModerInvite.CountLgtInvite,
		CountLgtInvite:                         countLgtInviteStr,
		PromoKod:                               strconv.Itoa(adsession_result.Data.Adsession.Face_id) + "-30-" + moderInvite_result.Data.ModerInvite.Dt + "-" + paramTenderidStr,
		MarketingInviteDisplay:                 marketingInviteDisplay,
		ReportFrom:                             reportFrom,
		ReportWhere:                            reportWhere,
		OrgCompanyid:                           moderInvite_result.Data.ModerInvite.OrgCompanyId,
		CompanyAddressDiffersFromTenderAddress: companyAddressDiffersFromTenderAddress,
		HrefPageSidTenderid:                    "/staff/moder_tender_company_invite/sid=" + paramSidStr + "&tenderid=" + paramTenderidStr,

		TenderCountryId:   strconv.Itoa(moderInvite_result.Data.ModerInvite.TenderCountryId),
		TenderCountryName: moderInvite_result.Data.ModerInvite.TenderCountryName,

		CompanyCountryId:   strconv.Itoa(moderInvite_result.Data.ModerInvite.CompanyCountryId),
		CompanyCountryName: moderInvite_result.Data.ModerInvite.CompanyCountryName,

		TenderRegionId:             moderInvite_result.Data.ModerInvite.TenderRegionId,
		TenderRegionName:           moderInvite_result.Data.ModerInvite.TenderRegionName,
		PlaceCodeEqRTenderRegionId: paramPlaceCode == "R"+strconv.Itoa(moderInvite_result.Data.ModerInvite.TenderRegionId),

		TenderAreaId:             moderInvite_result.Data.ModerInvite.TenderAreaId,
		TenderAreaName:           moderInvite_result.Data.ModerInvite.TenderAreaName,
		PlaceCodeEqATenderAreaId: paramPlaceCode == "A"+strconv.Itoa(moderInvite_result.Data.ModerInvite.TenderAreaId),

		CompanyRegionId:             moderInvite_result.Data.ModerInvite.CompanyRegionId,
		CompanyRegionName:           moderInvite_result.Data.ModerInvite.CompanyRegionName,
		PlaceCodeEqRCompanyRegionId: paramPlaceCode == "R"+strconv.Itoa(moderInvite_result.Data.ModerInvite.CompanyRegionId),

		CompanyAreaId:             moderInvite_result.Data.ModerInvite.CompanyAreaId,
		CompanyAreaName:           moderInvite_result.Data.ModerInvite.CompanyAreaName,
		PlaceCodeEqACompanyAreaId: paramPlaceCode == "A"+strconv.Itoa(moderInvite_result.Data.ModerInvite.CompanyAreaId),
	}
	// парсим шаблоны и загоняем данные
	fmap := template.FuncMap{
		"ToHtml":    utils.ToHtml,
		"ConcatStr": utils.ConcatStr,
		"ToStr":     utils.ToStr,
	}
	t := template.Must(template.New("moder_tender_company_invite.gohtml").Funcs(fmap).ParseFiles(filepath.Join(utils.Cfg.StaffPath, "moder_tender_company_invite.gohtml"), filepath.Join(utils.Cfg.IncludesPath, "header.gohtml"), filepath.Join(utils.Cfg.IncludesPath, "footer.gohtml"), filepath.Join(utils.Cfg.IncludesPath, "page_path.gohtml")))
	errExec := t.Execute(w, data)
	if errExec != nil {
		fmt.Println("Ошибка генерации шаблона:", errExec.Error())
		http.Error(w, errExec.Error(), http.StatusInternalServerError)
	}
}
