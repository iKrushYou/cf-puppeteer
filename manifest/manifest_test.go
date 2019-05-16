package manifest_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/happytobi/cf-puppeteer/manifest"
)

func TestManifestPartser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manifest Testsuite")
}

var _ = Describe("Parse Manifest", func() {
	It("parses complete manifest", func() {
		manifest, err := Parse("../fixtures/manifest.yml")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(manifest.ApplicationManifests[0].Name).Should(Equal("myApp"))
		Expect(manifest.ApplicationManifests[0].Buildpacks[0]).Should(Equal("java_buildpack"))
		Expect(manifest.ApplicationManifests[0].Buildpacks[1]).Should(Equal("go_buildpack"))
	})

	It("parses complete manifest with services", func() {
		manifest, err := Parse("../fixtures/manifest.yml")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(manifest.ApplicationManifests[0].Name).Should(Equal("myApp"))
		Expect(manifest.ApplicationManifests[0].Services[0]).Should(Equal("service1"))
		Expect(manifest.ApplicationManifests[0].Services[1]).Should(Equal("service2"))
	})
})

var _ = Describe("Parse multi Application Manifest", func() {
	It("parses complete manifest", func() {
		manifest, err := Parse("../fixtures/multiManifest.yml")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(manifest.ApplicationManifests[0].Name).Should(Equal("myApp"))
		Expect(manifest.ApplicationManifests[1].Name).Should(Equal("myApp2"))
	})
})

var _ = Describe("Parse invalid Application Manifest", func() {
	It("parses invalid manifest", func() {
		manifest, err := Parse("../fixtures/invalidManifest.yml")
		Expect(err).ShouldNot(BeNil())
		Expect(manifest.ApplicationManifests).Should(BeNil())
	})
})
