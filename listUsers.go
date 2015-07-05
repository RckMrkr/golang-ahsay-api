package ahsay

import (
	"encoding/xml"
)

type UserList struct {
	Users []User `xml:",attr"`
}

type Contact struct {
	Name  string `xml:",attr"`
	Email string `xml:",attr"`
}

type User struct {
	Username                     string     `xml:"LoginName,attr"`
	Owner                        string     `xml:",attr"`
	UserType                     UserType   `xml:",attr"`
	ClientType                   ClientType `xml:",attr"`
	Quota                        ByteSize   `xml:",attr"`
	Timezone                     string     `xml:",attr"`
	Language                     string     `xml:",attr"`
	DataFile                     int        `xml:",attr"`
	DataSize                     ByteSize   `xml:",attr"`
	RetainFile                   int        `xml:",attr"`
	RetainSize                   ByteSize   `xml:",attr"`
	EnableMSSQL                  Boolean    `xml:",attr"`
	EnableMSExchange             Boolean    `xml:",attr"`
	EnableOracle                 Boolean    `xml:",attr"`
	EnableLotusNotes             Boolean    `xml:",attr"`
	EnableLotusDomino            Boolean    `xml:",attr"`
	EnableMySQL                  Boolean    `xml:",attr"`
	EnableInFileDelta            Boolean    `xml:",attr"`
	EnableShadowCopy             Boolean    `xml:",attr"`
	EnableExchangeMailbox        Boolean    `xml:",attr"`
	ExchangeMailboxQuota         Boolean    `xml:",attr"`
	EnableNASClient              Boolean    `xml:",attr"`
	EnableDeltaMerge             Boolean    `xml:",attr"`
	EnableMsVm                   Boolean    `xml:",attr"`
	MsVmQuota                    ByteSize   `xml:",attr"`
	EnableVMware                 Boolean    `xml:",attr"`
	VmWareQuota                  ByteSize   `xml:",attr"`
	Bandwidth                    string     `xml:",attr"` // not sure of the format. using string to be safe
	Notes                        string     `xml:",attr"`
	Status                       Status     `xml:",attr"`
	RegistrationDate             Timestamp  `xml:",attr"`
	SuspendPaidUser              Boolean    `xml:",attr"`
	SuspendPaidUserDate          string     `xml:",attr"` // SOMETHING HAS TO BE FIXED HERE
	LastBackupDate               Timestamp  `xml:",attr"`
	EnableCDP                    Boolean    `xml:",attr"`
	EnableShadowProtectBareMetal Boolean    `xml:",attr"`
	EnableWinServer2008BareMetal Boolean    `xml:",attr"`
	Hostname                     string     `xml:",attr"`
	FileSizeLimit                ByteSize   `xml:",attr"`
	ExcludeNetworkShare          Boolean    `xml:",attr"`
	Contact                      Contact
}

func ListUsers(s Server, args map[string]string) <-chan UserList {
	returnChan = make(chan UserList)
	go func(){
		body := <-request(s, data, ep)
		object := UserList{}
		_ := xml.
	}()
}
