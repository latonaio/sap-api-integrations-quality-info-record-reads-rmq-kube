package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			Material                      string      `json:"Material"`
			QltyInProcmtIntID             string      `json:"QltyInProcmtIntID"`
			Supplier                      string      `json:"Supplier"`
			Plant                         string      `json:"Plant"`
			MaterialRevisionLevel         string      `json:"MaterialRevisionLevel"`
			QltyInProcmtReleaseValidTo    string      `json:"QltyInProcmtReleaseValidTo"`
			BaseUnit                      string      `json:"BaseUnit"`
			ReleasedQuantity              string      `json:"ReleasedQuantity"`
			BlockReason                   string      `json:"BlockReason"`
			ProcurementBlock              string      `json:"ProcurementBlock"`
			QltyInProcmtRelQtyIsActive    string      `json:"QltyInProcmtRelQtyIsActive"`
			QltyInProcmtInspControl       string      `json:"QltyInProcmtInspControl"`
			ProdnPieceApprovalIsRequired  string      `json:"ProdnPieceApprovalIsRequired"`
			ProductionPieceApproval       string      `json:"ProductionPieceApproval"`
			ProductionPieceApprovalLevel  string      `json:"ProductionPieceApprovalLevel"`
			ProductionPieceApprovalStatus string      `json:"ProductionPieceApprovalStatus"`
			IsDeleted                     string      `json:"IsDeleted"`
			QualityManagementSystem       string      `json:"QualityManagementSystem"`
			QltyInProcmtCertfnValidTo     string      `json:"QltyInProcmtCertfnValidTo"`
			InspectionLotType             string      `json:"InspectionLotType"`
			QltyInProcmtLotCrtnLeadTime   int         `json:"QltyInProcmtLotCrtnLeadTime"`
			QltyInProcmtCertificateCtrl   string      `json:"QltyInProcmtCertificateCtrl"`
			QltyInProcmtLongText          string      `json:"QltyInProcmtLongText"`
		} `json:"results"`
	} `json:"d"`
}  
