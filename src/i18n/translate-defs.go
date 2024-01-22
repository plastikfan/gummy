package i18n

// TODO: Should be updated to use url of the implementing project,
// so should not be left as arcadia.
const GummySourceID = "github.com/plastikfan/gummy"

type gummyTemplData struct{}

func (td gummyTemplData) SourceID() string {
	return GummySourceID
}
