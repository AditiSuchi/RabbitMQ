package services

// ReportService coordinates the generation and export of reports
type ReportService struct {
	// Add dependencies here (e.g., S3Service, ExcelService, EmailService)
}

// NewReportService creates a new instance of ReportService
func NewReportService() *ReportService {
	return &ReportService{}
}
