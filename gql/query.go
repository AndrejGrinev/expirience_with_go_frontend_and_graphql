package gql

import (
	"context"
	"tpro/stage1g/utils"
)

func GetAdsessionBySID(sid string, userip string) (*utils.Adsession_result, error) {
	var face_id, face_type, company_id int
        var face_name string
	err := utils.Db.QueryRow(context.Background(), "SELECT face_id, face_type, face_name, COALESCE(company_id, 0) AS company_id FROM app.sid_info($1::ws.d_sid, $2::inet)", sid, userip).Scan(&face_id, &face_type, &face_name, &company_id)
	if err != nil {
		return nil, err
	}
	return &utils.Adsession_result{
		Sid:        sid,
		Face_id:    face_id,
		Face_type:  face_type,
		Face_name:  face_name,
		Company_id: company_id,
	}, nil
}

func GetModerInviteById(id string) (*utils.ModerInvite_result, error) {
	var tp, state, org_companyid, org_testing, company_country_id, company_region_id, company_area_id, tender_country_id, tender_region_id, tender_area_id, count_repeat_all, company_lgt_invite, count_lgt_invite int
        var title, company_country_name, company_region_name, company_area_name, tender_country_name, tender_region_name, tender_area_name, dt string
        var is_auction, tender_invite_admin, access_org_to_marketing bool
	err := utils.Db.QueryRow(context.Background(), "SELECT type, is_auction, state, title, org_companyId, org_testing::INT, company_country_id, company_region_id, company_area_id, COALESCE(company_country_name, ''), COALESCE(company_region_name, ''), COALESCE(company_area_name, ''), tender_country_id, tender_region_id, tender_area_id, tender_country_name, tender_region_name, tender_area_name, dt, count_repeat_all, company_lgt_invite, count_lgt_invite, tender_invite_admin, access_org_to_marketing FROM tnd_v2.moder_invite_data($1)", id).Scan(&tp, &is_auction, &state, &title, &org_companyid, &org_testing, &company_country_id, &company_region_id, &company_area_id, &company_country_name, &company_region_name, &company_area_name, &tender_country_id, &tender_region_id, &tender_area_id, &tender_country_name, &tender_region_name, &tender_area_name, &dt, &count_repeat_all, &company_lgt_invite, &count_lgt_invite, &tender_invite_admin, &access_org_to_marketing)
	if err != nil {
		return nil, err
	}
	return &utils.ModerInvite_result{
		Id:                   id,
		Tp:                   tp,
		IsAuction:            is_auction,
		State:                state,
		Title:                title,
		OrgCompanyId:         org_companyid,
		OrgTesting:           org_testing,
		CompanyCountryId:     company_country_id,
		CompanyRegionId:      company_region_id,
		CompanyAreaId:        company_area_id,
		CompanyCountryName:   company_country_name,
		CompanyRegionName:    company_region_name,
		CompanyAreaName:      company_area_name,
		TenderCountryId:      tender_country_id,
		TenderRegionId:       tender_region_id,
		TenderAreaId:         tender_area_id,
		TenderCountryName:    tender_country_name,
		TenderRegionName:     tender_region_name,
		TenderAreaName:       tender_area_name,
		Dt:                   dt,
		CountRepeatAll:       count_repeat_all,
		CompanyLgtInvite:     company_lgt_invite,
		CountLgtInvite:       count_lgt_invite,
		TenderInviteAdmin:    tender_invite_admin,
		AccessOrgToMarketing: access_org_to_marketing,
	}, nil
}
