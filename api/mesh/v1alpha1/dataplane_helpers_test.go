package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	gomega_types "github.com/onsi/gomega/types"

	. "github.com/Kong/kuma/api/mesh/v1alpha1"
)

var _ = Describe("ServiceTagValue", func() {

	Describe("HasPort()", func() {
		type testCase struct {
			value    string
			expected bool
		}

		DescribeTable("should determine correctly whether a service tag includes service port",
			func(given testCase) {
				Expect(ServiceTagValue(given.value).HasPort()).To(Equal(given.expected))
			},
			Entry("name only", testCase{
				value:    "web",
				expected: false,
			}),
			Entry("name and port", testCase{
				value:    "web.default.svc:80",
				expected: true,
			}),
			Entry("incomplete value", testCase{
				value:    "web:",
				expected: true,
			}),
		)
	})

	Describe("HostAndPort()", func() {
		type testCase struct {
			value        string
			expectedHost string
			expectedPort uint32
			expectedErr  string
		}

		DescribeTable("should parse `service` tag value into host and port",
			func(given testCase) {
				// when
				host, port, err := ServiceTagValue(given.value).HostAndPort()

				if given.expectedErr != "" {
					Expect(err).To(MatchError(given.expectedErr))
				} else {
					Expect(err).ToNot(HaveOccurred())
					Expect(host).To(Equal(given.expectedHost))
					Expect(port).To(Equal(given.expectedPort))
				}
			},
			Entry("name and port", testCase{
				value:        "web.default.svc:80",
				expectedHost: "web.default.svc",
				expectedPort: 80,
				expectedErr:  "",
			}),
			Entry("incomplete value", testCase{
				value:        "web:",
				expectedHost: "",
				expectedPort: 0,
				expectedErr:  `strconv.ParseUint: parsing "": invalid syntax`,
			}),
			Entry("name only", testCase{
				value:        "web",
				expectedHost: "",
				expectedPort: 0,
				expectedErr:  "address web: missing port in address",
			}),
		)
	})
})

var _ = Describe("InboundInterface", func() {

	Describe("String()", func() {
		type testCase struct {
			iface    InboundInterface
			expected string
		}

		DescribeTable("should format properly",
			func(given testCase) {
				Expect(given.iface.String()).To(Equal(given.expected))
			},
			Entry("all fields set", testCase{
				iface: InboundInterface{
					DataplaneIP:   "1.2.3.4",
					DataplanePort: 80,
					WorkloadPort:  8080,
				},
				expected: "1.2.3.4:80:8080",
			}),
		)
	})
})

var _ = Describe("ParseInboundInterface(..)", func() {

	Context("valid input values", func() {
		type testCase struct {
			input    string
			expected InboundInterface
		}

		DescribeTable("should parse valid input values",
			func(given testCase) {
				// when
				iface, err := ParseInboundInterface(given.input)
				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(iface).To(Equal(given.expected))
			},
			Entry("all fields set", testCase{
				input: "1.2.3.4:80:8080",
				expected: InboundInterface{
					DataplaneIP:   "1.2.3.4",
					DataplanePort: 80,
					WorkloadPort:  8080,
				},
			}),
		)
	})

	Context("invalid input values", func() {
		type testCase struct {
			input       string
			expectedErr gomega_types.GomegaMatcher
		}

		DescribeTable("should fail on invalid input values",
			func(given testCase) {
				// when
				iface, err := ParseInboundInterface(given.input)
				// then
				Expect(err.Error()).To(given.expectedErr)
				// and
				Expect(iface).To(BeZero())
			},
			Entry("dataplane IP address is missing", testCase{
				input:       ":80:8080",
				expectedErr: MatchRegexp(`invalid format: expected .*, got ":80:8080"`),
			}),
			Entry("dataplane IP address is not valid", testCase{
				input:       "localhost:80:65536",
				expectedErr: MatchRegexp(`invalid format: expected .*, got "localhost:80:65536"`),
			}),
			Entry("service port is missing", testCase{
				input:       "1.2.3.4::8080",
				expectedErr: MatchRegexp(`invalid format: expected .*, got "1.2.3.4::8080"`),
			}),
			Entry("service port is out of range", testCase{
				input:       "1.2.3.4:0:8080",
				expectedErr: Equal(`invalid <DATAPLANE_PORT> in "1.2.3.4:0:8080": port number must be in the range [1, 65535] but got 0`),
			}),
			Entry("application port is missing", testCase{
				input:       "1.2.3.4:80:",
				expectedErr: MatchRegexp(`invalid format: expected .*, got "1.2.3.4:80:"`),
			}),
			Entry("application port is out of range", testCase{
				input:       "1.2.3.4:80:65536",
				expectedErr: Equal(`invalid <WORKLOAD_PORT> in "1.2.3.4:80:65536": port number must be in the range [1, 65535] but got 65536`),
			}),
		)
	})
})

var _ = Describe("ParseOutboundInterface(..)", func() {

	Context("valid input values", func() {
		type testCase struct {
			input    string
			expected OutboundInterface
		}

		DescribeTable("should parse valid input values",
			func(given testCase) {
				// when
				oface, err := ParseOutboundInterface(given.input)
				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(oface).To(Equal(given.expected))
			},
			Entry("all fields set", testCase{
				input: "127.0.0.2:18080",
				expected: OutboundInterface{
					DataplaneIP:   "127.0.0.2",
					DataplanePort: 18080,
				},
			}),
			Entry("dataplane IP address is missing", testCase{
				input: ":18080",
				expected: OutboundInterface{
					DataplaneIP:   "127.0.0.1",
					DataplanePort: 18080,
				},
			}),
		)
	})

	Context("invalid input values", func() {
		type testCase struct {
			input       string
			expectedErr gomega_types.GomegaMatcher
		}

		DescribeTable("should fail on invalid input values",
			func(given testCase) {
				// when
				iface, err := ParseInboundInterface(given.input)
				// then
				Expect(err.Error()).To(given.expectedErr)
				// and
				Expect(iface).To(BeZero())
			},
			Entry("dataplane IP address is not valid", testCase{
				input:       "localhost:65536",
				expectedErr: MatchRegexp(`invalid format: expected .*, got "localhost:65536"`),
			}),
			Entry("port without colon", testCase{
				input:       "18080",
				expectedErr: MatchRegexp(`invalid format: expected .*, got "18080"`),
			}),
		)
	})
})

var _ = Describe("Dataplane_Networking", func() {

	Describe("GetInboundInterfaces()", func() {

		Context("valid input values", func() {
			type testCase struct {
				input    *Dataplane_Networking
				expected []InboundInterface
			}

			DescribeTable("should parse valid input values",
				func(given testCase) {
					// when
					ifaces, err := given.input.GetInboundInterfaces()
					// then
					Expect(err).ToNot(HaveOccurred())
					// and
					Expect(ifaces).To(ConsistOf(given.expected))
				},
				Entry("nil", testCase{
					input:    nil,
					expected: nil,
				}),
				Entry("empty", testCase{
					input:    &Dataplane_Networking{},
					expected: []InboundInterface{},
				}),
				Entry("2 inbound interfaces", testCase{
					input: &Dataplane_Networking{
						Inbound: []*Dataplane_Networking_Inbound{
							{Interface: "192.168.0.1:80:8080"},
							{Interface: "192.168.0.1:443:8443"},
						},
					},
					expected: []InboundInterface{
						{DataplaneIP: "192.168.0.1", DataplanePort: 80, WorkloadPort: 8080},
						{DataplaneIP: "192.168.0.1", DataplanePort: 443, WorkloadPort: 8443},
					},
				}),
			)
		})

		Context("invalid input values", func() {
			type testCase struct {
				input       *Dataplane_Networking
				expectedErr gomega_types.GomegaMatcher
			}

			DescribeTable("should fail on invalid input values",
				func(given testCase) {
					// when
					ifaces, err := given.input.GetInboundInterfaces()
					// then
					Expect(ifaces).To(BeNil())
					// and
					Expect(err.Error()).To(given.expectedErr)
				},
				Entry("dataplane IP address is missing", testCase{
					input: &Dataplane_Networking{
						Inbound: []*Dataplane_Networking_Inbound{
							{Interface: "192.168.0.1:80:8080"},
							{Interface: ":443:8443"},
						},
					},
					expectedErr: MatchRegexp(`invalid format: expected .*, got ":443:8443"`),
				}),
			)
		})
	})
})

var _ = Describe("Dataplane_Networking_Outbound", func() {
	type testCase struct {
		serviceTag    string
		selector      TagSelector
		expectedMatch bool
	}
	DescribeTable("MatchTags()",
		func(given testCase) {
			//given
			outbound := Dataplane_Networking_Outbound{
				Interface: "sdf",
				Service:   given.serviceTag,
			}

			// when
			matched := outbound.MatchTags(given.selector)

			// then
			Expect(matched).To(Equal(given.expectedMatch))
		},
		Entry("it should match *", testCase{
			serviceTag: "backend",
			selector: map[string]string{
				"service": "*",
			},
			expectedMatch: true,
		}),
		Entry("it should match service", testCase{
			serviceTag: "backend",
			selector: map[string]string{
				"service": "backend",
			},
			expectedMatch: true,
		}),
		Entry("it shouldn't match tag other than service", testCase{
			serviceTag: "backend",
			selector: map[string]string{
				"version": "1.0",
			},
			expectedMatch: false,
		}),
	)
})

var _ = Describe("Dataplane", func() {
	d := Dataplane{
		Networking: &Dataplane_Networking{
			Inbound: []*Dataplane_Networking_Inbound{
				{
					Tags: map[string]string{
						"service": "backend",
						"version": "v1",
					},
				},
				{
					Tags: map[string]string{
						"service": "backend-metrics",
						"version": "v1",
						"role":    "metrics",
					},
				},
			},
		},
	}

	Describe("Tags()", func() {
		It("should provide combined tags", func() {
			// when
			tags := d.Tags()

			// then
			Expect(tags.Values("service")).To(Equal([]string{"backend", "backend-metrics"}))
			Expect(tags.Values("version")).To(Equal([]string{"v1"}))
			Expect(tags.Values("role")).To(Equal([]string{"metrics"}))
		})
	})

	Describe("MatchTags()", func() {
		It("should match any inbound", func() {
			// when
			selector := TagSelector{
				"service": "backend",
				"version": "v1",
			}

			// then
			Expect(d.MatchTags(selector)).To(BeTrue())
		})

		It("should not match if all inbounds did not match", func() {
			// when
			selector := TagSelector{
				"service": "unknown",
			}

			// then
			Expect(d.MatchTags(selector)).To(BeFalse())
		})
	})
})

var _ = Describe("TagSelector", func() {

	Describe("Matches()", func() {
		type testCase struct {
			tags  map[string]string
			match bool
		}
		DescribeTable("should Match tags",
			func(given testCase) {
				// given
				dpTags := map[string]string{
					"service": "mobile",
					"version": "v1",
				}

				// when
				match := TagSelector(given.tags).Matches(dpTags)

				//then
				Expect(match).To(Equal(given.match))
			},
			Entry("should match 0 tags", testCase{
				tags:  map[string]string{},
				match: true,
			}),
			Entry("should match 1 tag", testCase{
				tags:  map[string]string{"service": "mobile"},
				match: true,
			}),
			Entry("should match all tags", testCase{
				tags: map[string]string{
					"service": "mobile",
					"version": "v1",
				},
				match: true,
			}),
			Entry("should match * tag", testCase{
				tags:  map[string]string{"service": "*"},
				match: true,
			}),
			Entry("should not match on one mismatch", testCase{
				tags: map[string]string{
					"service": "backend",
					"version": "v1",
				},
				match: false,
			}),
		)
	})

	Describe("Equal()", func() {
		type testCase struct {
			one      TagSelector
			another  TagSelector
			expected bool
		}

		DescribeTable("should correctly determine if two selectors are equal",
			func(given testCase) {
				// expect
				Expect(given.one.Equal(given.another)).To(Equal(given.expected))
			},
			Entry("two nil selectors", testCase{
				one:      nil,
				another:  nil,
				expected: true,
			}),
			Entry("nil selector and empty selector", testCase{
				one:      nil,
				another:  TagSelector{},
				expected: true,
			}),
			Entry("empty selector and nil selector", testCase{
				one:      TagSelector{},
				another:  nil,
				expected: true,
			}),
			Entry("two empty selectors", testCase{
				one:      TagSelector{},
				another:  TagSelector{},
				expected: true,
			}),
			Entry("equal selectors of 1 tag", testCase{
				one:      TagSelector{"service": "backend"},
				another:  TagSelector{"service": "backend"},
				expected: true,
			}),
			Entry("equal selectors of 2 tag", testCase{
				one:      TagSelector{"service": "backend", "version": "v1"},
				another:  TagSelector{"service": "backend", "version": "v1"},
				expected: true,
			}),
			Entry("unequal selectors of 1 tag", testCase{
				one:      TagSelector{"service": "backend"},
				another:  TagSelector{"service": "redis"},
				expected: false,
			}),
			Entry("one 1 tag selector and one 2 tags selector", testCase{
				one:      TagSelector{"service": "backend"},
				another:  TagSelector{"service": "redis", "version": "v1"},
				expected: false,
			}),
		)
	})
})

var _ = Describe("Tags", func() {
	It("should print tags", func() {
		// given
		tags := map[string]map[string]bool{
			"service": map[string]bool{
				"backend-api":   true,
				"backend-admin": true,
			},
			"version": {
				"v1": true,
			},
		}

		// when
		result := MultiValueTagSet(tags).String()

		// then
		Expect(result).To(Equal("service=backend-admin,backend-api version=v1"))
	})
})

var _ = Describe("TagSelectorRank", func() {

	Describe("CompareTo()", func() {
		type testCase struct {
			rank1    TagSelectorRank
			rank2    TagSelectorRank
			expected int
		}
		DescribeTable("should correctly compare two ranks",
			func(given testCase) {
				// expect
				Expect(given.rank1.CompareTo(given.rank2)).To(Equal(given.expected))
			},
			Entry("0 ranks are equal", testCase{
				rank1:    TagSelectorRank{},
				rank2:    TagSelectorRank{},
				expected: 0,
			}),
			Entry("matches by the same number of exact values (1) are equal", testCase{
				rank1:    TagSelectorRank{ExactMatches: 1},
				rank2:    TagSelectorRank{ExactMatches: 1},
				expected: 0,
			}),
			Entry("matches by the same number of wildcard values (2) are equal", testCase{
				rank1:    TagSelectorRank{WildcardMatches: 2},
				rank2:    TagSelectorRank{WildcardMatches: 2},
				expected: 0,
			}),
			Entry("equal ranks by non-0 ExactMatches and WildcardMatches", testCase{
				rank1:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 2},
				rank2:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 2},
				expected: 0,
			}),
			Entry("match by an exact value (1) is more specific than match by a wildcard", testCase{
				rank1:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 0},
				rank2:    TagSelectorRank{ExactMatches: 0, WildcardMatches: 1},
				expected: 1,
			}),
			Entry("match by a wildcard is less specific than match by an exact value (1)", testCase{
				rank1:    TagSelectorRank{ExactMatches: 0, WildcardMatches: 1},
				rank2:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 0},
				expected: -1,
			}),
			Entry("match by an exact value (2) is more specific than match by a wildcard", testCase{
				rank1:    TagSelectorRank{ExactMatches: 2, WildcardMatches: 0},
				rank2:    TagSelectorRank{ExactMatches: 0, WildcardMatches: 2},
				expected: 2,
			}),
			Entry("match by a wildcard is less specific than match by an exact value (2)", testCase{
				rank1:    TagSelectorRank{ExactMatches: 0, WildcardMatches: 2},
				rank2:    TagSelectorRank{ExactMatches: 2, WildcardMatches: 0},
				expected: -2,
			}),
			Entry("match by an exact value (3) is more specific than match by a wildcard", testCase{
				rank1:    TagSelectorRank{ExactMatches: 3, WildcardMatches: 0},
				rank2:    TagSelectorRank{ExactMatches: 0, WildcardMatches: 3},
				expected: 3,
			}),
			Entry("match by a wildcard is less specific than match by an exact value (3)", testCase{
				rank1:    TagSelectorRank{ExactMatches: 0, WildcardMatches: 3},
				rank2:    TagSelectorRank{ExactMatches: 3, WildcardMatches: 0},
				expected: -3,
			}),
			Entry("match by an exact value is more specific than match by a wildcard", testCase{
				rank1:    TagSelectorRank{ExactMatches: 2, WildcardMatches: 1},
				rank2:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 1},
				expected: 1,
			}),
			Entry("match by a wildcard is less specific than match by an exact value", testCase{
				rank1:    TagSelectorRank{ExactMatches: 2, WildcardMatches: 1},
				rank2:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 1},
				expected: 1,
			}),
		)
	})
	Describe("CombinedWith()", func() {
		type testCase struct {
			rank1    TagSelectorRank
			rank2    TagSelectorRank
			expected TagSelectorRank
		}
		DescribeTable("should correctly aggregate two ranks",
			func(given testCase) {
				// expect
				Expect(given.rank1.CombinedWith(given.rank2)).To(Equal(given.expected))
			},
			Entry("combination of two 0 ranks is zero rank", testCase{
				rank1:    TagSelectorRank{},
				rank2:    TagSelectorRank{},
				expected: TagSelectorRank{},
			}),
			Entry("cobination of a match by an exact value with a match by a wildcard", testCase{
				rank1:    TagSelectorRank{ExactMatches: 1},
				rank2:    TagSelectorRank{WildcardMatches: 2},
				expected: TagSelectorRank{ExactMatches: 1, WildcardMatches: 2},
			}),
			Entry("cobination of two mixed matches", testCase{
				rank1:    TagSelectorRank{ExactMatches: 1, WildcardMatches: 2},
				rank2:    TagSelectorRank{ExactMatches: 10, WildcardMatches: 20},
				expected: TagSelectorRank{ExactMatches: 11, WildcardMatches: 22},
			}),
		)
	})
})
