package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-quality-info-record-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			Material:                      data.Material,
			QltyInProcmtIntID:             data.QltyInProcmtIntID,
			Supplier:                      data.Supplier,
			Plant:                         data.Plant,
			MaterialRevisionLevel:         data.MaterialRevisionLevel,
			QltyInProcmtReleaseValidTo:    data.QltyInProcmtReleaseValidTo,
			BaseUnit:                      data.BaseUnit,
			ReleasedQuantity:              data.ReleasedQuantity,
			BlockReason:                   data.BlockReason,
			ProcurementBlock:              data.ProcurementBlock,
			QltyInProcmtRelQtyIsActive:    data.QltyInProcmtRelQtyIsActive,
			QltyInProcmtInspControl:       data.QltyInProcmtInspControl,
			ProdnPieceApprovalIsRequired:  data.ProdnPieceApprovalIsRequired,
			ProductionPieceApproval:       data.ProductionPieceApproval,
			ProductionPieceApprovalLevel:  data.ProductionPieceApprovalLevel,
			ProductionPieceApprovalStatus: data.ProductionPieceApprovalStatus,
			IsDeleted:                     data.IsDeleted,
			QualityManagementSystem:       data.QualityManagementSystem,
			QltyInProcmtCertfnValidTo:     data.QltyInProcmtCertfnValidTo,
			InspectionLotType:             data.InspectionLotType,
			QltyInProcmtLotCrtnLeadTime:   data.QltyInProcmtLotCrtnLeadTime,
			QltyInProcmtCertificateCtrl:   data.QltyInProcmtCertificateCtrl,
			QltyInProcmtLongText:          data.QltyInProcmtLongText,
		})
	}

	return header, nil
}
