package verify

import (
	"strconv"

	"github.com/sjlleo/netflix-verify/util"
)

type IPv4Verifier struct {
	Config
	IP              string
	unblockStatus   int
	unblockTestChan chan UnblockTestResult
}

func (v *IPv4Verifier) Execute() *VerifyResponse {
	var err error
	var response VerifyResponse
	v.unblockStatus = AreaUnavailable
	response.Type = 1

	if v.IP, err = util.DnsResolver(4); err != nil {
		response.StatusCode = NetworkUnrachable
		return &response
	}

	v.unblockTestChan = make(chan UnblockTestResult)
	defer close(v.unblockTestChan)

	go v.UnblockTest(AreaAvailableID)
	go v.UnblockTest(SelfMadeAvailableID)
	go v.UnblockTest(NonSelfMadeAvailableID)

	for i := 0; i < 3; i++ {
		switch res := <-v.unblockTestChan; {

		case res.err != nil:
			response.StatusCode = NetworkUnrachable

		case res.CountryCode != "":
			switch res.movieID {
			case AreaAvailableID:
				v.upgradeStatus(AreaAvailable)
			case SelfMadeAvailableID:
				v.upgradeStatus(UnblockSelfMadeMovie)
			case NonSelfMadeAvailableID:
				v.upgradeStatus(UnblockNonSelfMadeMovie)
			}
			response.CountryCode = res.CountryCode
			response.CountryName = util.CountryCodeToCountryName(res.CountryCode)
		default:
		}
	}
	response.StatusCode = v.unblockStatus
	return &response
}

func (v *IPv4Verifier) upgradeStatus(status int) {
	if v.unblockStatus < status {
		v.unblockStatus = status
	}
}

func (v *IPv4Verifier) UnblockTest(MoiveID int) {

	testURL := util.Netflix + strconv.Itoa(MoiveID)
	if reCode, err := util.RequestIP(testURL, v.IP, v.LocalAddr); err != nil {
		if err.Error() == "Banned" {
			v.unblockTestChan <- UnblockTestResult{MoiveID, "", nil}
		} else {
			v.unblockTestChan <- UnblockTestResult{MoiveID, "", err}
		}
	} else {
		v.unblockTestChan <- UnblockTestResult{MoiveID, reCode, nil}
	}
}
