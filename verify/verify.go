package verify

var AreaAvailableID = 80018499
var SelfMadeAvailableID = 80197526
var NonSelfMadeAvailableID = 70143836

const (
	AreaUnavailable         = -1
	AreaAvailable           = 0
	UnblockSelfMadeMovie    = 1
	UnblockNonSelfMadeMovie = 2
)

type Verify interface {
	Execute() *VerifyResponse
}

type FinalResult struct {
	Res map[int]VerifyResponse
}

type VerifyResult struct {
	Config
	response chan *VerifyResponse
}

type Config struct {
	LocalAddr string
	Custom    string
}

type VerifyResponse struct {
	Type        int
	StatusCode  int
	CountryCode string
	CountryName string
}

type UnblockTestResult struct {
	movieID     int
	CountryCode string
	err         error
}

func NewVerify(c Config) *FinalResult {
	var finalResult FinalResult
	finalResult.Res = make(map[int]VerifyResponse)

	vr := VerifyResult{Config: c}
	ch := make(chan *VerifyResponse)
	vr.response = ch
	defer close(ch)

	go vr.IPv4Verifier()
	go vr.IPv6Verifier()

	for i := 0; i < 2; i++ {
		if res, err := <-ch; err {
			finalResult.Res[res.Type] = *res
		}
	}
	return &finalResult
}

func (vr *VerifyResult) IPv4Verifier() {
	verify := &IPv4Verifier{Config: vr.Config}
	vr.response <- verify.Execute()
}

func (vr *VerifyResult) IPv6Verifier() {
	verify := &IPv6Verifier{Config: vr.Config}
	vr.response <- verify.Execute()
}
