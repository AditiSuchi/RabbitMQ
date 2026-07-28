package controllers

import( "github.com/gin-gonic/gin"

)

// // ReportController handles report export HTTP requests
// type ReportController struct {
// 	// Add dependencies here (e.g., ReportService)
// }

// // NewReportController creates a new instance of ReportController
// func NewReportController() *ReportController {
// 	return &ReportController{}
// }

func ExportReport(c *gin.Context) {

	var req ExportRequest

	c.BindJSON(&req)

	Publish(req)

	c.JSON(200, gin.H{
		"message": "Report generation started",
	})
}
