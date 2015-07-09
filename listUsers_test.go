package ahsay

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example taken from the api documentation
		s := `
		<users>
			<User
				LoginName="test1" Owner="" Alias="" UserType="PAID" ClientType="OBM"
				Quota="10737418240" Timezone="GMT+08:00 (CST)" Language="en"
				DataFile="1" DataSize="1536" RetainFile="0" RetainSize="0"
				EnableMSSQL="Y" EnableMSExchange="Y" EnableOracle="Y"
				EnableLotusNotes="Y" EnableLotusDomino="Y" EnableMySQL="Y"
				EnableInFileDelta="Y" EnableShadowCopy="Y" EnableExchangeMailbox="N"
				ExchangeMailboxQuota="0" EnableNASClient="Y" EnableDeltaMerge="Y"
				EnableMsVm="N" MsVmQuota="0" EnableVMware="N" VMwareQuota="0"
				Bandwidth="0" Notes="" Status="ENABLE" RegistrationDate="1302687743242"
				SuspendPaidUser="N" SuspendPaidUserDate="20140503"
				LastBackupDate="1302699594652" EnableCDP="Y"
				EnableShadowProtectBareMetal="Y" EnableWinServer2008BareMetal="Y"
				Hostname="123.abc.com" FileSizeLimit="52428800" ExcludeNetworkShare="Y">
					<Contact Name=""Email="www@qqq.com"/>
				</user>
			</users>`
		fmt.Fprint(w, s)
	}))

	defer ts.Close()
	data := make(map[string]string)
	server := servertest(true)

	response := <-request(server, data, ts.URL, UserList{})
	fmt.Printf("%v", response.Object)
	users := UserList{
		Users: []User{
			User{
				LoginName: "Test",
			},
		},
	}
	assert.Nil(response.Err)
	assert.Equal(users, response.Object)
}
