package moduledoc

import "fmt"

func (s docHandler) BuildTOC() {
	para := s.doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Heading1")
	run.AddText("List of Endpoint")
	para = s.doc.AddParagraph()

	nd := s.doc.Numbering.AddDefinition()
	for path := range s.swagger.Paths {
		pathItems := s.swagger.Paths.Find(path)

		if pathItems.Get != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodGet, path, pathItems.Get.Summary))
		}

		if pathItems.Post != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodPost, path, pathItems.Post.Summary))
		}

		if pathItems.Put != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodPut, path, pathItems.Put.Summary))
		}

		if pathItems.Patch != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodPatch, path, pathItems.Patch.Summary))
		}

		if pathItems.Delete != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodDelete, path, pathItems.Delete.Summary))
		}

		if pathItems.Head != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodHead, path, pathItems.Head.Summary))
		}

		if pathItems.Options != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodOptions, path, pathItems.Options.Summary))
		}

		if pathItems.Trace != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s - %s", MethodTrace, path, pathItems.Trace.Summary))
		}
	}

	para.AddRun().AddPageBreak()
}
