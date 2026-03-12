package jsonld

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-qat"
	"github.com/reiver/go-rtxt"
	"github.com/reiver/go-schemaorg"
)

// MarshalFAQPage converts a slice of QA pairs into schema.org FAQPage JSON-LD bytes.
//
// For example, a QAT file such as this:
//
//	Q: What is Go?
//	
//	A: A programming-language.
//	
//	Q: Is Go have any alternative names?
//	
//	A: Yes, an alternative name for Go is: golang.
//
// Would produce a []qat.QA similar to this:
//
//	qas := []qat.QA{
//		{
//			Question: "What is Go?",
//			Answer: "A programming-language.",
//		},
//		{
//			Question: "Is Go have any alternative names?",
//			Answer: "Yes, an alternative name for Go is: golang.",
//		},
//	}
//
// Which when passed to MarshalFAQPage:
//
//	bytes, err := jsonld.MarshalFAQPage(qas)
//
// Would create JSON-LD similar to the following:
//
//	{
//		"@context": "https://schema.org",
//		"@type": "FAQPage",
//		"mainEntity": [
//			{
//				"@type": "Question",
//				"name": "What is Go?",
//				"acceptedAnswer": {
//					"@type": "Answer",
//					"text": "A programming-language."
//				}
//			},
//			{
//				"@type": "Question",
//				"name": "Is Go have any alternative names?",
//				"acceptedAnswer": {
//					"@type": "Answer",
//					"text": "Yes, an alternative name for Go is: golang."
//				}
//			}
//		]
//	}
func MarshalFAQPage(qas []qat.QA) ([]byte, error) {

	var questions []schemaorg.ProtoThing

	for _, qa := range qas {
		question := schemaorg.Question{
			Name: opt.Something(qa.Question),
			AcceptedAnswer: opt.Something(schemaorg.Answer{
				Text: opt.Something(rtxt.ToHTML(qa.Answer)),
			}),
		}
		questions = append(questions, question)
	}

	faqPage := schemaorg.FAQPage{
		MainEntity: questions,
	}

	return jsonld.Marshal(faqPage)
}
