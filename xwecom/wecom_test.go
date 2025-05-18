package xwecom

import (
	"fmt"
	"testing"
	"time"
)

func TestWecom(t *testing.T) {
	wecom := &Wecom{
		AgentID:    1000007,
		CorpID:     "wweb8736907463b6f1",
		CorpSecret: "Whx_pkmyHhawpktZZxEGFGwEQz9Y68fPpPFnX5Emq6c",
		ToParty:    "",
		ToUser:     "gtlions",
	}
	wecom.ToUser = "gtlions"
	// wecom.SendText([]string{"hello world"})

	msg := fmt.Sprintf("**日期** `%s`\n", time.Now())
	msg = msg + fmt.Sprintf("**[数据合计]** `%d`\n", time.Now().Unix())
	wecom.SendMD([]string{msg})
}
