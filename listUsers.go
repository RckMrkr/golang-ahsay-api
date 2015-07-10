package ahsay

import "encoding/xml"

// UserList is a slice of list user
type UserList struct {
	Users []User `xml:"User"`
}

// Contact is the contact info for a user
type Contact struct {
	Name  string `xml:",attr"`
	Email string `xml:",attr"`
}

// User is a struct with properties for a specific user
type User struct {
	LoginName                    string     `xml:",attr"`
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
	ExchangeMailboxQuota         ByteSize   `xml:",attr"`
	EnableNASClient              Boolean    `xml:",attr"`
	EnableDeltaMerge             Boolean    `xml:",attr"`
	EnableMsVM                   Boolean    `xml:"EnableMsVm,attr"`
	MsVMQuota                    ByteSize   `xml:"MsVmQuota,attr"`
	EnableVMware                 Boolean    `xml:",attr"`
	VMWareQuota                  ByteSize   `xml:"VmWareQuota,attr"`
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
	Contact                      []Contact
}

// ListUsers calls the endpoint "ListUsers.do" on server s with argurments args and returns a channel for the response
func ListUsers(s Server, args map[string]string) (<-chan UserList, <-chan error) {
	url := createURL(s, "ListUsers.do")
	return listUsers(s, args, url)
}

func listUsers(s Server, args map[string]string, url string) (<-chan UserList, <-chan error) {
	c := request(s, args, url)
	objChan := make(chan UserList)
	errChan := make(chan error)
	go func() {
		r := <-c
		if r.Err != nil {
			errChan <- r.Err
		}
		obj := *new(UserList)
		err := xml.Unmarshal(r.Body, &obj)
		if err != nil {
			errChan <- err
		}
		close(errChan)
		objChan <- UserList(obj)
	}()

	return objChan, errChan
}
