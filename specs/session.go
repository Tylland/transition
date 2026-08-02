package specs

import (
	"log"
	"os"
	"path/filepath"
)

var S *Session = &Session{}

type Session struct {
	Name string
	ReportWriter
	ReporterList
	WorkingDir string
	OutputDir  string
}

func (s *Session) Start(name string, outputPath string) *Session {
	s.Name = name

	currentDir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	s.WorkingDir = currentDir
	s.OutputDir = filepath.Join(s.WorkingDir, outputPath)

	return s
}

func (s *Session) WithReportWriter(writer ReportWriter) *Session {
	writer.Attach(s)
	s.ReportWriter = writer

	return s
}

func (s *Session) WithReporter(rep Reporter) *Session {
	rep.Attach(s)
	s.Add(rep)

	return s
}

func (s *Session) Finish() {
	reportModel := s.ReportWriter.Finish()

	s.ReporterList.Report(reportModel)
}
