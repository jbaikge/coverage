package body

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

// http://www.biosciencetechnology.com/news/2013/03/brain-adds-specified-cells-during-puberty
// <div>Content</div><div>&nbsp;</div>
func TestBioscienceTechnology(t *testing.T) {
	//compareBodies(t, "BioscienceTechnology")
}

// http://blogs.cio.com/security/17837/scientists-create-telepathic-rats-and-robosparrows
// Missing titles from article list at end of article - need to add \n after ul's li's
func TestCIOBlogs(t *testing.T) {
	//compareBodies(t, "CIOBlogs")
}

// http://blog.cppionline.org/2012/10/supercharging-us-economy-where-both.html
// <div>Content</div><div><br></div>
func TestCPPIOnline(t *testing.T) {
	//compareBodies(t, "CPPIOnline")
}

// 221797 - Data.gov Expands Federal Communities, Global Impact
// http://gov.aol.com/2012/09/26/data-gov-expands-federal-communities-global-impact/
func TestAOLGov(t *testing.T) {
	compareBodies(t, "AOLGov")
}

// 340032 - Big Mac returns to SoCal as Dodgers hitting coach
// http://www.usatoday.com/story/sports/mlb/dodgers/2012/11/07/mark-mcgwire-hired-as-dodgers-hitting-coach/1690053/
func TestUSAToday(t *testing.T) {
	//compareBodies(t, "USAToday")
}

// 340038 - Greek Lawmakers Pass Austerity Deal
// http://online.wsj.com/article/SB10001424127887323894704578104832247833100.html
func TestWallStreetJournal(t *testing.T) {
	//compareBodies(t, "WallStreetJournal")
}

// 340046 - Concur FY Q4 Revs Light; Profits Beat; Shrs Edge Higher
// http://www.forbes.com/sites/ericsavitz/2012/11/07/concur-fy-q4-revs-light-profits-beat-shrs-edge-higher/
func TestForbes(t *testing.T) {
	//compareBodies(t, "Forbes")
}

// 340066 - Echo360 buys LectureTools in first acquisition since funding
// http://www.bizjournals.com/washington/news/2012/11/07/in-first-acquisition-since-funding.html?s=article_search
func TestBizJournals(t *testing.T) {
	//compareBodies(t, "BizJournals")
}

// 340140 - Waze Launches In-App Advertising Platform
// http://www.pcmag.com/article2/0,2817,2411868,00.asp
func TestPCMag(t *testing.T) {
	//compareBodies(t, "PCMag")
}

// 393887 - 2012 Federal Employee Viewpoint Survey Results
// http://gov.aol.com/2012/11/29/2012-federal-employee-viewpoint-survey-results/
func TestAOLGovSurveyResults(t *testing.T) {
	//compareBodies(t, "AOLGovSurveyResults")
}

func compareBodies(t *testing.T, basename string) {
	htmlIn, err := os.Open("../samples/" + basename + ".html")
	if err != nil {
		t.Error(err)
		return
	}
	defer htmlIn.Close()
	html, err := ioutil.ReadAll(htmlIn)
	if err != nil {
		t.Error(err)
		return
	}
	b, err := GetBody(html)
	if err != nil {
		t.Error(err)
		return
	}
	bodyIn, err := os.Open("../samples/" + basename + ".body")
	if err != nil {
		t.Error(err)
		return
	}
	defer bodyIn.Close()
	expect, err := ioutil.ReadAll(bodyIn)
	if err != nil {
		t.Error(err)
		return
	}
	expect = bytes.TrimSpace(expect)
	if string(b.Text) != string(expect) {
		t.Error("Bodies do not match")
		t.Logf("Expect:\n%s", expect)
		t.Logf("Got:\n%s", b.Text)
	}
}
