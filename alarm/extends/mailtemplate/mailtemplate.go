package mailtemplate

import (
	"bytes"
	"github.com/open-falcon/alarm/g"
	"github.com/open-falcon/common/model"
	"github.com/open-falcon/common/utils"
	"log"
	"text/template"
)

var tmpl, err = template.ParseFiles("mail.tpl")

type MailTmpl struct {
	Status        string
	Priority      int
	Endpoint      string
	IP            string
	Metric        string
	Note          string
	PushedTags    string
	Func          string
	LeftValue     string
	Operator      string
	RightValue    string
	CurrentStep   int
	MaxStep       int
	FormattedTime string
	Link          string
}

func initMailTmpl(event *model.Event, mailinfo *MailTmpl) {
	mailinfo.Status = event.Status
	mailinfo.Priority = event.Priority()
	mailinfo.Endpoint = event.Endpoint
	mailinfo.IP = ""
	mailinfo.Metric = event.Metric()
	mailinfo.Note = event.Note()
	mailinfo.PushedTags = utils.SortedTags(event.PushedTags)
	mailinfo.Func = event.Func()
	mailinfo.LeftValue = utils.ReadableFloat(event.LeftValue)
	mailinfo.Operator = event.Operator()
	mailinfo.RightValue = utils.ReadableFloat(event.RightValue())
	mailinfo.CurrentStep = event.CurrentStep
	mailinfo.MaxStep = event.MaxStep()
	mailinfo.FormattedTime = event.FormattedTime()
	mailinfo.Link = g.Link(event)
}

func BuildCommonMailContent(event *model.Event) string {
	var mailinfo *MailTmpl = new(MailTmpl)
	initMailTmpl(event, mailinfo)

	if err != nil {
		panic(err)
	}
	buff := bytes.NewBufferString("")
	err = tmpl.Execute(buff, mailinfo)
	if err != nil {
		panic(err)
	}
	str := buff.String()
	log.Println("gen alarm mail content", str)
	return str
}
