package reporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

type Report struct {
	ExecutiveSummary    string
	TechnicalDetails    string
	ProofOfConcept      string
	ImpactAnalysis      string
	Remediation         string
	References          string
}

func ExportReport(report Report, format string) {
	switch format {
	case "json":
		exportJSON(report)
	case "markdown":
		exportMarkdown(report)
	case "pdf":
		exportPDF(report)
	case "html":
		exportHTML(report)
	case "csv":
		exportCSV(report)
	}
}

func exportJSON(report Report) {
	fmt.Println("Exporting report to JSON...")
	file, _ := json.MarshalIndent(report, "", " ")
	_ = os.WriteFile("report.json", file, 0644)
}

func exportMarkdown(report Report) {
	fmt.Println("Exporting report to Markdown...")
	// Placeholder for actual Markdown export logic
}

func exportPDF(report Report) {
	fmt.Println("Exporting report to PDF...")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Vulnerability Report")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, "Executive Summary:\n"+report.ExecutiveSummary, "", "", false)
	pdf.Ln(10)
	pdf.MultiCell(0, 10, "Technical Details:\n"+report.TechnicalDetails, "", "", false)
	pdf.Ln(10)
	pdf.MultiCell(0, 10, "Proof of Concept:\n"+report.ProofOfConcept, "", "", false)
	pdf.Ln(10)
	pdf.MultiCell(0, 10, "Impact Analysis:\n"+report.ImpactAnalysis, "", "", false)
	pdf.Ln(10)
	pdf.MultiCell(0, 10, "Remediation Recommendations:\n"+report.Remediation, "", "", false)
	pdf.Ln(10)
	pdf.MultiCell(0, 10, "References:\n"+report.References, "", "", false)
	err := pdf.OutputFileAndClose("report.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

func exportHTML(report Report) {
	fmt.Println("Exporting report to HTML...")
	// Placeholder for actual HTML export logic
}

func exportCSV(report Report) {
	fmt.Println("Exporting report to CSV...")
	file, _ := os.Create("report.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Executive Summary", "Technical Details", "Proof of Concept", "Impact Analysis", "Remediation", "References"}
	row := []string{report.ExecutiveSummary, report.TechnicalDetails, report.ProofOfConcept, report.ImpactAnalysis, report.Remediation, report.References}

	writer.Write(headers)
	writer.Write(row)
}
