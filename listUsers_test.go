package ahsay

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrivateListUsers(t *testing.T) {
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
					<Contact Name="" Email="www@qqq.com"/>
				</User>
			</users>`
		fmt.Fprint(w, s)
	}))

	defer ts.Close()
	server := servertest(true)
	data := make(map[string]string)
	objChan, errChan := listUsers(server, data, ts.URL)
	users := UserList{
		Users: []User{
			User{
				LoginName:                    "test1",
				Owner:                        "",
				UserType:                     Paid,
				ClientType:                   Obm,
				Quota:                        ByteSize(10737418240),
				Timezone:                     "GMT+08:00 (CST)",
				Language:                     "en",
				DataFile:                     1,
				DataSize:                     ByteSize(1536),
				RetainFile:                   0,
				RetainSize:                   ByteSize(0),
				EnableMSSQL:                  BooleanTrue,
				EnableMSExchange:             BooleanTrue,
				EnableOracle:                 BooleanTrue,
				EnableLotusNotes:             BooleanTrue,
				EnableLotusDomino:            BooleanTrue,
				EnableMySQL:                  BooleanTrue,
				EnableInFileDelta:            BooleanTrue,
				EnableShadowCopy:             BooleanTrue,
				EnableExchangeMailbox:        BooleanFalse,
				ExchangeMailboxQuota:         ByteSize(0),
				EnableNASClient:              BooleanTrue,
				EnableDeltaMerge:             BooleanTrue,
				EnableMsVM:                   BooleanFalse,
				MsVMQuota:                    ByteSize(0),
				EnableVMware:                 BooleanFalse,
				VMWareQuota:                  ByteSize(0),
				Bandwidth:                    "0",
				Notes:                        "",
				Status:                       StatusEnabled,
				RegistrationDate:             Timestamp(time.Unix(1302687743, 242000000)),
				SuspendPaidUser:              BooleanFalse,
				SuspendPaidUserDate:          "20140503",
				LastBackupDate:               Timestamp(time.Unix(1302699594, 652000000)),
				EnableCDP:                    BooleanTrue,
				EnableShadowProtectBareMetal: BooleanTrue,
				EnableWinServer2008BareMetal: BooleanTrue,
				Hostname:                     "123.abc.com",
				FileSizeLimit:                ByteSize(52428800),
				ExcludeNetworkShare:          BooleanTrue,
				Contact:                      []Contact{Contact{Name: "", Email: "www@qqq.com"}},
			},
		},
	}

	assert.Nil(<-errChan)
	assert.Equal(users, <-objChan)
}

func TestPrivateListUsersInvalid(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example taken from the api documentation
		s := `Invalid`
		fmt.Fprint(w, s)
	}))

	defer ts.Close()
	server := servertest(true)
	data := make(map[string]string)
	_, errChan := listUsers(server, data, ts.URL)

	assert.NotNil(<-errChan)
	_, errChan = listUsers(server, data, "http://this.must.be.invalid.xaz")

	assert.NotNil(<-errChan)
}
