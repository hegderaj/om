package commands_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
)

var _ = Describe("CreateVMExtension", func() {
	var (
		fakeVMExtensionService *fakes.VMExtensionCreator
		fakeLogger             *fakes.Logger
		command                commands.CreateVMExtension
	)

	BeforeEach(func() {
		fakeVMExtensionService = &fakes.VMExtensionCreator{}
		fakeLogger = &fakes.Logger{}
		command = commands.NewCreateVMExtension(fakeVMExtensionService, fakeLogger)
	})

	Describe("Execute", func() {
		It("makes a request to the OpsMan to create a VM extension", func() {
			err := command.Execute([]string{
				"--name", "some-vm-extension",
				"--cloud-properties", "{ \"iam_instance_profile\": \"some-iam-profile\", \"elbs\": [\"some-elb\"] }",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeVMExtensionService.CreateArgsForCall(0)).To(Equal(api.VMExtensionInput{
				Name:            "some-vm-extension",
				CloudProperties: "{ \"iam_instance_profile\": \"some-iam-profile\", \"elbs\": [\"some-elb\"] }",
			}))

			Expect(fakeLogger.PrintfCallCount()).To(Equal(1))
			format, content := fakeLogger.PrintfArgsForCall(0)
			Expect(fmt.Sprintf(format, content...)).To(Equal("VM Extension 'some-vm-extension' created\n"))
		})

		Context("failure cases", func() {
			Context("when the service fails to create a VM extension", func() {
				It("returns an error", func() {
					fakeVMExtensionService.CreateReturns(errors.New("failed to create VM extension"))

					err := command.Execute([]string{
						"--name", "some-vm-extension",
						"--cloud-properties", "{ \"iam_instance_profile\": \"some-iam-profile\", \"elbs\": [\"some-elb\"] }",
					})

					Expect(err).To(MatchError("failed to create VM extension"))
				})
			})

			Context("when an unknown flag is provided", func() {
				It("returns an error", func() {
					err := command.Execute([]string{"--badflag"})
					Expect(err).To(MatchError("could not parse create-vm-extension flags: flag provided but not defined: -badflag"))
				})
			})
		})
	})
})
