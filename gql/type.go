package gql

import (
	"github.com/graphql-go/graphql"
	"tpro/stage1g/utils"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"adsession": &graphql.Field{
			Type: AdsessionType,
			Args: graphql.FieldConfigArgument{
				"sid": &graphql.ArgumentConfig{
					Description: "Session ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"userip": &graphql.ArgumentConfig{
					Description: "Session User IP",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sid := p.Args["sid"].(string)
				userip := p.Args["userip"].(string)
				return GetAdsessionBySID(sid, userip)
			},
		},
		"moderInvite": &graphql.Field{
			Type: ModerInviteType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Tender ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)
				return GetModerInviteById(id)
			},
		},
	},
})

var AdsessionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Adsession",
	Fields: graphql.Fields{
		"sid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if adsession, ok := p.Source.(*utils.Adsession_result); ok == true {
					return adsession.Sid, nil
				}
				return nil, nil
			},
		},
		"face_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if adsession, ok := p.Source.(*utils.Adsession_result); ok == true {
					return adsession.Face_id, nil
				}
				return nil, nil
			},
		},
		"face_type": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if adsession, ok := p.Source.(*utils.Adsession_result); ok == true {
					return adsession.Face_type, nil
				}
				return nil, nil
			},
		},
		"face_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if adsession, ok := p.Source.(*utils.Adsession_result); ok == true {
					return adsession.Face_name, nil
				}
				return nil, nil
			},
		},
		"company_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if adsession, ok := p.Source.(*utils.Adsession_result); ok == true {
					return adsession.Company_id, nil
				}
				return nil, nil
			},
		},
	},
})

var ModerInviteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ModerInvite",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.Id, nil
				}
				return nil, nil
			},
		},
		"type": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.Tp, nil
				}
				return nil, nil
			},
		},
		"isAuction": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.IsAuction, nil
				}
				return nil, nil
			},
		},
		"state": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.State, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.Title, nil
				}
				return nil, nil
			},
		},
		"org_companyid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.OrgCompanyId, nil
				}
				return nil, nil
			},
		},
		"org_testing": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.OrgTesting, nil
				}
				return nil, nil
			},
		},
		"company_country_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyCountryId, nil
				}
				return nil, nil
			},
		},
		"company_region_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyRegionId, nil
				}
				return nil, nil
			},
		},
		"company_area_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyAreaId, nil
				}
				return nil, nil
			},
		},
		"company_country_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyCountryName, nil
				}
				return nil, nil
			},
		},
		"company_region_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyRegionName, nil
				}
				return nil, nil
			},
		},
		"company_area_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyAreaName, nil
				}
				return nil, nil
			},
		},
		"tender_country_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderCountryId, nil
				}
				return nil, nil
			},
		},
		"tender_region_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderRegionId, nil
				}
				return nil, nil
			},
		},
		"tender_area_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderAreaId, nil
				}
				return nil, nil
			},
		},
		"tender_country_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderCountryName, nil
				}
				return nil, nil
			},
		},
		"tender_region_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderRegionName, nil
				}
				return nil, nil
			},
		},
		"tender_area_name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderAreaName, nil
				}
				return nil, nil
			},
		},
		"dt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.Dt, nil
				}
				return nil, nil
			},
		},
		"count_repeat_all": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CountRepeatAll, nil
				}
				return nil, nil
			},
		},
		"company_lgt_invite": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CompanyLgtInvite, nil
				}
				return nil, nil
			},
		},
		"count_lgt_invite": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.CountLgtInvite, nil
				}
				return nil, nil
			},
		},
		"tender_invite_admin": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.TenderInviteAdmin, nil
				}
				return nil, nil
			},
		},
		"access_org_to_marketing": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if moderInvite, ok := p.Source.(*utils.ModerInvite_result); ok == true {
					return moderInvite.AccessOrgToMarketing, nil
				}
				return nil, nil
			},
		},
	},
})
