package command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/plastikfan/gummy/src/app/command"
	"github.com/plastikfan/gummy/src/internal/helpers"
	"github.com/snivilised/extendio/xfs/utils"
)

var _ = Describe("RootCmd", Ordered, func() {
	var (
		repo     string
		l10nPath string
	)

	BeforeAll(func() {
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		Expect(utils.FolderExists(l10nPath)).To(BeTrue())
	})

	It("🧪 should: execute", func() {
		bootstrap := command.Bootstrap{}
		tester := helpers.CommandTester{
			Args: []string{},
			Root: bootstrap.Root(func(co *command.ConfigureOptions) {
				co.Detector = &DetectorStub{}
				co.Config.Name = configName
				co.Config.ConfigPath = configPath
			}),
		}
		_, err := tester.Execute()
		Expect(err).Error().To(BeNil())
	})
})
