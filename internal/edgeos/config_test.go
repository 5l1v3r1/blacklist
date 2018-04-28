package edgeos

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/britannic/blacklist/internal/tdata"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddInc(t *testing.T) {
	Convey("Testing addInc()", t, func() {
		var (
			c   = NewConfig()
			err = c.ReadCfg(&CFGstatic{Cfg: tdata.Cfg})
		)

		So(err, ShouldBeNil)

		tests := []struct {
			name string
			exp  *source
			node string
		}{
			{
				name: "rootNode",
				node: rootNode,
				exp: &source{
					Env: &Env{
						Wildcard: Wildcard{
							Node: "",
							Name: "",
						},
						ctr:   make(ctr),
						API:   "",
						Arch:  "",
						Bash:  "",
						Cores: 0,
						Dbug:  false,
						Dex: list{
							RWMutex: &sync.RWMutex{},
							entry:   entry{},
						},
						Dir:    "",
						DNSsvc: "",
						Exc: list{
							RWMutex: &sync.RWMutex{},
							entry:   entry{},
						},
						Ext:      "",
						File:     "",
						FnFmt:    "",
						InCLI:    "",
						ioWriter: nil,
						Level:    "",
						Method:   "",
						Pfx:      dnsPfx{domain: "", host: ""},
						Test:     false,
						Timeout:  time.Duration(0),
						Verb:     false},
					desc:     "pre-configured global blacklisted domains",
					disabled: false,
					err:      nil,
					exc:      nil,
					file:     "",
					inc:      []string{},
					ip:       "0.0.0.0",
					ltype:    "global-blacklisted-domains",
					name:     "global-blacklisted-domains",
					nType:    ntype(8),
					Objects: Objects{
						Env: nil,
						src: nil,
					},
					prefix: "",
					r:      nil,
					url:    "",
				},
			},
			{
				name: "domains",
				node: domains,
				exp: &source{
					Env: &Env{
						Wildcard: Wildcard{
							Node: "",
							Name: "",
						},
						ctr:   make(ctr),
						API:   "",
						Arch:  "",
						Bash:  "",
						Cores: 0,
						Dbug:  false,
						Dex: list{
							RWMutex: &sync.RWMutex{},
							entry:   entry{},
						},
						Dir:    "",
						DNSsvc: "",
						Exc: list{
							RWMutex: &sync.RWMutex{},
							entry:   entry{},
						},
						Ext:      "",
						File:     "",
						FnFmt:    "",
						InCLI:    "",
						ioWriter: nil,
						Level:    "",
						Method:   "",
						Pfx:      dnsPfx{domain: "", host: ""},
						Test:     false,
						Timeout:  time.Duration(0),
						Verb:     false},
					desc:     "pre-configured blacklisted subdomains",
					disabled: false,
					err:      nil,
					exc:      nil,
					file:     "",
					inc:      []string{"adsrvr.org", "adtechus.net", "advertising.com", "centade.com", "doubleclick.net", "free-counter.co.uk", "intellitxt.com", "kiosked.com", "patoghee.in"},
					ip:       "192.168.100.1",
					ltype:    "blacklisted-subdomains",
					name:     "blacklisted-subdomains",
					nType:    ntype(6),
					Objects: Objects{
						Env: nil,
						src: nil,
					},
					prefix: "",
					r:      nil,
					url:    "",
				},
			},
			{name: "hosts",
				node: hosts,
				exp: &source{
					Env: &Env{
						Wildcard: Wildcard{
							Node: "",
							Name: "",
						},
						ctr:   make(ctr),
						API:   "",
						Arch:  "",
						Bash:  "",
						Cores: 0,
						Dbug:  false,
						Dex: list{
							RWMutex: &sync.RWMutex{},
							entry:   entry{},
						},
						Dir:    "",
						DNSsvc: "",
						Exc: list{
							RWMutex: &sync.RWMutex{},
							entry:   entry{},
						},
						Ext:      "",
						File:     "",
						FnFmt:    "",
						InCLI:    "",
						ioWriter: nil,
						Level:    "",
						Method:   "",
						Pfx:      dnsPfx{domain: "", host: ""},
						Test:     false,
						Timeout:  time.Duration(0),
						Verb:     false},
					desc:     "pre-configured blacklisted servers",
					disabled: false,
					err:      nil,
					exc:      nil,
					file:     "",
					inc:      []string{"beap.gemini.yahoo.com"},
					ip:       "0.0.0.0",
					ltype:    "blacklisted-servers",
					name:     "blacklisted-servers",
					nType:    ntype(7),
					Objects: Objects{
						Env: nil,
						src: nil,
					},
					prefix: "",
					r:      nil,
					url:    "",
				},
			},
		}

		for _, tt := range tests {
			Convey("Testing "+tt.name, func() {
				inc := c.addInc(tt.node)

				So(inc, ShouldResemble, tt.exp)
			})
		}
	})
}

func TestExcludes(t *testing.T) {
	Convey("Testing excludes()", t, func() {
		c := NewConfig(
			Dir("/tmp"),
			Ext("blacklist.conf"),
		)

		So(c.ReadCfg(&CFGstatic{Cfg: tdata.Cfg}), ShouldBeNil)

		excludes := list{
			entry: entry{
				"sstatic.net":             0,
				"yimg.com":                0,
				"ytimg.com":               0,
				"google.com":              0,
				"images-amazon.com":       0,
				"msdn.com":                0,
				"schema.org":              0,
				"skype.com":               0,
				"avast.com":               0,
				"bitdefender.com":         0,
				"cdn.visiblemeasures.com": 0,
				"cloudfront.net":          0,
				"microsoft.com":           0,
				"akamaihd.net":            0,
				"amazon.com":              0,
				"apple.com":               0,
				"shopify.com":             0,
				"storage.googleapis.com":  0,
				"msecnd.net":              0,
				"ssl-on9.com":             0,
				"windows.net":             0,
				"1e100.net":               0,
				"akamai.net":              0,
				"coremetrics.com":         0,
				"gstatic.com":             0,
				"gvt1.com":                0,
				"freedns.afraid.org":      0,
				"hb.disney.go.com":        0,
				"hp.com":                  0,
				"live.com":                0,
				"rackcdn.com":             0,
				"edgesuite.net":           0,
				"googleapis.com":          0,
				"smacargo.com":            0,
				"static.chartbeat.com":    0,
				"gvt1.net":                0,
				"hulu.com":                0,
				"paypal.com":              0,
				"amazonaws.com":           0,
				"ask.com":                 0,
				"github.com":              0,
				"githubusercontent.com":   0,
				"googletagmanager.com":    0,
				"sourceforge.net":         0,
				"xboxlive.com":            0,
				"2o7.net":                 0,
				"adobedtm.com":            0,
				"googleadservices.com":    0,
				"googleusercontent.com":   0,
				"ssl-on9.net":             0,
			},
		}
		tests := []struct {
			get  list
			list list
			name string
			node string
		}{
			{name: "c.excludes(rootNode)", get: c.excludes(rootNode), list: excludes, node: rootNode},
			{name: "c.excludes()", get: c.excludes(), list: excludes},
			{name: "c.excludes(domains)", get: c.excludes(domains), list: list{RWMutex: (*sync.RWMutex)(nil), entry: entry{}}, node: domains},
			{name: "c.excludes(hosts)", get: c.excludes(hosts), list: list{RWMutex: (*sync.RWMutex)(nil), entry: entry{}}, node: hosts},
		}

		for _, tt := range tests {
			Convey("Testing "+tt.name, func() {
				switch tt.node {
				case "":
					So(c.excludes(), ShouldResemble, tt.list)
				default:
					So(c.excludes(tt.node), ShouldResemble, tt.list)
				}
			})
		}
	})
}

func TestGetIP(t *testing.T) {
	b := tree{}
	Convey("Testing getIP(badnode)", t, func() {
		So(b.getIP("badnode"), ShouldEqual, "0.0.0.0")
	})
	b = tree{
		rootNode: &source{
			ip: "192.168.1.50",
		},
		domains: &source{
			ip: "192.168.1.20",
		},
		hosts: &source{
			ip: "192.168.1.30",
		},
	}
	Convey("Testing getIP("+rootNode+")", t, func() {
		So(b.getIP(rootNode), ShouldEqual, "192.168.1.50")
	})
	Convey("Testing getIP("+domains+")", t, func() {
		So(b.getIP(domains), ShouldEqual, "192.168.1.20")
	})
	Convey("Testing getIP("+hosts+")", t, func() {
		So(b.getIP(hosts), ShouldEqual, "192.168.1.30")
	})
}

func TestFiles(t *testing.T) {
	Convey("Testing c.GetAll().Files()", t, func() {
		r := &CFGstatic{Cfg: tdata.Cfg}
		c := NewConfig(
			Dir("/tmp"),
			Ext("blacklist.conf"),
		)

		So(c.ReadCfg(r), ShouldBeNil)

		exp := `/tmp/domains.blacklisted-subdomains.blacklist.conf
/tmp/domains.malc0de.blacklist.conf
/tmp/domains.malwaredomains.com.blacklist.conf
/tmp/domains.simple_tracking.blacklist.conf
/tmp/domains.zeus.blacklist.conf
/tmp/hosts.blacklisted-servers.blacklist.conf
/tmp/hosts.openphish.blacklist.conf
/tmp/hosts.raw.github.com.blacklist.conf
/tmp/hosts.sysctl.org.blacklist.conf
/tmp/hosts.tasty.blacklist.conf
/tmp/hosts.volkerschatz.blacklist.conf
/tmp/hosts.yoyo.blacklist.conf
/tmp/roots.global-blacklisted-domains.blacklist.conf`

		act := c.GetAll().Files().String()
		So(act, ShouldEqual, exp)
	})
}

func TestInSession(t *testing.T) {
	Convey("Testing InSession()", t, func() {
		c := NewConfig()
		So(c.InSession(), ShouldBeFalse)

		So(os.Setenv("_OFR_CONFIGURE", "ok"), ShouldBeNil)
		So(c.InSession(), ShouldBeTrue)

		So(os.Unsetenv("_OFR_CONFIGURE"), ShouldBeNil)
		So(c.InSession(), ShouldBeFalse)
	})
}

func TestIsSource(t *testing.T) {
	Convey("Testing TestIsSource()", t, func() {
		var node []string
		So(isntSource(node), ShouldBeTrue)
	})
}
func TestNodeExists(t *testing.T) {
	Convey("Testing TestNodeExists()", t, func() {
		var (
			c   = NewConfig()
			err = c.ReadCfg(&CFGstatic{Cfg: tdata.Cfg})
		)
		So(err, ShouldBeNil)
		So(c.nodeExists("broken"), ShouldBeFalse)
	})
}

func TestReadCfg(t *testing.T) {
	Convey("Testing ReadCfg()", t, func() {
		var (
			err error
			b   []byte
			f   = "../testdata/config.erx.boot"
			r   io.Reader
		)

		if r, err = GetFile(f); err != nil {
			Printf("cannot open configuration file %s!", f)
		}

		b, _ = ioutil.ReadAll(r)

		Convey("Testing with a configuration loaded from a file", func() {
			act := NewConfig().ReadCfg(&CFGstatic{Cfg: string(b)})
			So(act, ShouldBeEmpty)
		})

		Convey("Testing with an empty configuration", func() {
			exp := errors.New("no blacklist configuration has been detected")
			act := NewConfig().ReadCfg(&CFGstatic{Cfg: ""})
			So(act, ShouldResemble, exp)
		})
		Convey("Testing with a disabled configuration", func() {
			act := NewConfig().ReadCfg(&CFGstatic{Cfg: tdata.DisabledCfg})
			So(act, ShouldBeEmpty)
		})

		Convey("Testing with a single source configuration", func() {
			act := NewConfig().ReadCfg(&CFGstatic{Cfg: tdata.SingleSource})
			So(act, ShouldBeEmpty)
		})

		Convey("Testing with an active configuration", func() {
			c := NewConfig()
			So(c.ReadCfg(&CFGstatic{Cfg: tdata.Cfg}), ShouldBeNil)
			So(c.Nodes(), ShouldResemble, []string{"blacklist", "domains", "hosts"})
		})
	})
}

func TestReadUnconfiguredCfg(t *testing.T) {
	Convey("Testing ReadCfg()", t, func() {
		exp := errors.New("no blacklist configuration has been detected")
		act := NewConfig().ReadCfg(&CFGstatic{Cfg: tdata.NoBlacklist})
		So(act, ShouldResemble, exp)
	})
}

func TestReloadDNS(t *testing.T) {
	Convey("Testing ReloadDNS()", t, func() {
		act, err := NewConfig(Bash("/bin/bash"), DNSsvc("true")).ReloadDNS()
		So(err, ShouldBeNil)
		So(string(act), ShouldEqual, "")
	})
}

func TestRemove(t *testing.T) {
	Convey("Testing c.GetAll().Files().Remove()", t, func() {
		dir, _ := ioutil.TempDir("/tmp", "testBlacklist")
		defer os.RemoveAll(dir)

		c := NewConfig(
			Dir(dir),
			Ext("blacklist.conf"),
			FileNameFmt("%v/%v.%v.%v"),
			WCard(Wildcard{Node: "*s", Name: "*"}),
		)

		So(c.ReadCfg(&CFGstatic{Cfg: tdata.CfgMimimal}), ShouldBeNil)

		Convey("Creating special case file", func() {
			f, err := os.Create(fmt.Sprintf("%v/hosts.raw.github.com.blacklist.conf", dir))
			So(err, ShouldBeNil)
			f.Close()
		})

		for _, node := range c.sortKeys() {
			for i := range Iter(10) {
				fname := fmt.Sprintf("%v/%v.%v.%v", dir, node, i, c.Ext)
				f, err := os.Create(fname)
				So(err, ShouldBeNil)
				f.Close()
			}
		}

		for _, fname := range c.GetAll().Files().Strings() {
			f, err := os.Create(fname)
			So(err, ShouldBeNil)
			f.Close()
		}

		c.GetAll().Files().Remove()

		cf := &CFile{Env: c.Env}
		pattern := fmt.Sprintf(c.FnFmt, c.Dir, "*s", "*", c.Env.Ext)
		act, err := cf.readDir(pattern)

		So(err, ShouldBeNil)
		So(act, ShouldResemble, c.GetAll().Files().Strings())

		prev := c.SetOpt(WCard(Wildcard{Node: "[]a]", Name: "]"}))

		So(cf.Remove(), ShouldNotBeNil)
		c.SetOpt(prev)
	})
}

func TestBooltoString(t *testing.T) {
	Convey("Testing booltoString()", t, func() {
		So(booltoStr(true), ShouldEqual, True)
		So(booltoStr(false), ShouldEqual, False)
	})
}

func TestToBool(t *testing.T) {
	Convey("Testing strToBool()", t, func() {
		So(strToBool(True), ShouldBeTrue)
		So(strToBool(False), ShouldBeFalse)
	})
}

func TestGetAll(t *testing.T) {
	Convey("Testing GetAll() sources", t, func() {
		c := NewConfig(
			Dir("/tmp"),
			Ext(".blacklist.conf"),
		)

		So(c.ReadCfg(&CFGstatic{Cfg: tdata.Cfg}), ShouldBeNil)

		tests := []struct {
			exp   string
			ltype string
			name  string
		}{
			{name: "GetAll()", ltype: "", exp: expGetAll},
			{name: "GetAll(url)", ltype: urls, exp: expURLS},
			{name: "GetAll(files)", ltype: files, exp: expFiles},
			{name: "GetAll(PreDomns, PreHosts)", ltype: PreDomns, exp: expPre},
			{name: "Get(all).String()", ltype: all, exp: c.Get(all).String()},
			{name: "c.Get(hosts)", ltype: hosts, exp: expHostObj},
			{name: "c.Get(domains)", ltype: domains, exp: expDomainObj},
		}

		for _, tt := range tests {
			Convey("Testing "+tt.name, func() {
				switch tt.ltype {
				case "":
					So(c.GetAll().String(), ShouldEqual, tt.exp)
				case all:
					So(c.GetAll().String(), ShouldEqual, tt.exp)
				case domains:
					So(c.Get(domains).String(), ShouldEqual, tt.exp)
				case hosts:
					So(c.Get(hosts).String(), ShouldEqual, tt.exp)
				case PreDomns:
					So(c.GetAll(PreDomns, PreHosts).String(), ShouldEqual, tt.exp)
				default:
					So(c.GetAll(tt.ltype).String(), ShouldEqual, tt.exp)
				}
			})
		}
	})
}

func TestValidate(t *testing.T) {
	Convey("Testing validate() sources", t, func() {
		b := make(tree)
		So(b.validate("borked").String(), ShouldEqual, "")
	})
}

var (
	expDomainObj = `
Desc:         "pre-configured blacklisted subdomains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "blacklisted-subdomains"
Name:         "blacklisted-subdomains"
nType:        "preDomn"
Prefix:       "**Undefined**"
Type:         "blacklisted-subdomains"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "adsrvr.org"
              "adtechus.net"
              "advertising.com"
              "centade.com"
              "doubleclick.net"
              "free-counter.co.uk"
              "intellitxt.com"
              "kiosked.com"
              "patoghee.in"

Desc:         "List of zones serving malicious executables observed by malc0de.com/database/"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.168.1"
Ltype:        "url"
Name:         "malc0de"
nType:        "domn"
Prefix:       "zone "
Type:         "domains"
URL:          "http://malc0de.com/bl/ZONES"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Just domains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "10.0.0.1"
Ltype:        "url"
Name:         "malwaredomains.com"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "http://mirror1.malwaredomains.com/files/justdomains"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Basic tracking list by Disconnect"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "url"
Name:         "simple_tracking"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "abuse.ch ZeuS domain blocklist"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "url"
Name:         "zeus"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"
`

	expFiles = `
Desc:         "File source"
Disabled:     "false"
File:         "../internal/testdata/blist.hosts.src"
IP:           "10.10.10.10"
Ltype:        "file"
Name:         "tasty"
nType:        "host"
Prefix:       "**Undefined**"
Type:         "hosts"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"
`

	expGetAll = `
Desc:         "pre-configured global blacklisted domains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "global-blacklisted-domains"
Name:         "global-blacklisted-domains"
nType:        "preRoot"
Prefix:       "**Undefined**"
Type:         "global-blacklisted-domains"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "pre-configured blacklisted subdomains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "blacklisted-subdomains"
Name:         "blacklisted-subdomains"
nType:        "preDomn"
Prefix:       "**Undefined**"
Type:         "blacklisted-subdomains"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "adsrvr.org"
              "adtechus.net"
              "advertising.com"
              "centade.com"
              "doubleclick.net"
              "free-counter.co.uk"
              "intellitxt.com"
              "kiosked.com"
              "patoghee.in"

Desc:         "List of zones serving malicious executables observed by malc0de.com/database/"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.168.1"
Ltype:        "url"
Name:         "malc0de"
nType:        "domn"
Prefix:       "zone "
Type:         "domains"
URL:          "http://malc0de.com/bl/ZONES"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Just domains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "10.0.0.1"
Ltype:        "url"
Name:         "malwaredomains.com"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "http://mirror1.malwaredomains.com/files/justdomains"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Basic tracking list by Disconnect"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "url"
Name:         "simple_tracking"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "abuse.ch ZeuS domain blocklist"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "url"
Name:         "zeus"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "pre-configured blacklisted servers"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "blacklisted-servers"
Name:         "blacklisted-servers"
nType:        "preHost"
Prefix:       "**Undefined**"
Type:         "blacklisted-servers"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "beap.gemini.yahoo.com"

Desc:         "OpenPhish automatic phishing detection"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "openphish"
nType:        "host"
Prefix:       "http"
Type:         "hosts"
URL:          "https://openphish.com/feed.txt"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "This hosts file is a merged collection of hosts from reputable sources"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "raw.github.com"
nType:        "host"
Prefix:       "0.0.0.0 "
Type:         "hosts"
URL:          "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "This hosts file is a merged collection of hosts from cameleon"
Disabled:     "false"
File:         "**Undefined**"
IP:           "172.16.16.1"
Ltype:        "url"
Name:         "sysctl.org"
nType:        "host"
Prefix:       "127.0.0.1\t "
Type:         "hosts"
URL:          "http://sysctl.org/cameleon/hosts"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "File source"
Disabled:     "false"
File:         "../internal/testdata/blist.hosts.src"
IP:           "10.10.10.10"
Ltype:        "file"
Name:         "tasty"
nType:        "host"
Prefix:       "**Undefined**"
Type:         "hosts"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Ad server blacklists"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "volkerschatz"
nType:        "host"
Prefix:       "http"
Type:         "hosts"
URL:          "http://www.volkerschatz.com/net/adpaths"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Fully Qualified Domain Names only - no prefix to strip"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "yoyo"
nType:        "host"
Prefix:       "**Undefined**"
Type:         "hosts"
URL:          "https://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml&showintro=1&mimetype=plaintext"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"
`

	expHostObj = `
Desc:         "pre-configured blacklisted servers"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "blacklisted-servers"
Name:         "blacklisted-servers"
nType:        "preHost"
Prefix:       "**Undefined**"
Type:         "blacklisted-servers"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "beap.gemini.yahoo.com"

Desc:         "OpenPhish automatic phishing detection"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "openphish"
nType:        "host"
Prefix:       "http"
Type:         "hosts"
URL:          "https://openphish.com/feed.txt"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "This hosts file is a merged collection of hosts from reputable sources"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "raw.github.com"
nType:        "host"
Prefix:       "0.0.0.0 "
Type:         "hosts"
URL:          "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "This hosts file is a merged collection of hosts from cameleon"
Disabled:     "false"
File:         "**Undefined**"
IP:           "172.16.16.1"
Ltype:        "url"
Name:         "sysctl.org"
nType:        "host"
Prefix:       "127.0.0.1\t "
Type:         "hosts"
URL:          "http://sysctl.org/cameleon/hosts"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "File source"
Disabled:     "false"
File:         "../internal/testdata/blist.hosts.src"
IP:           "10.10.10.10"
Ltype:        "file"
Name:         "tasty"
nType:        "host"
Prefix:       "**Undefined**"
Type:         "hosts"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Ad server blacklists"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "volkerschatz"
nType:        "host"
Prefix:       "http"
Type:         "hosts"
URL:          "http://www.volkerschatz.com/net/adpaths"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Fully Qualified Domain Names only - no prefix to strip"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "yoyo"
nType:        "host"
Prefix:       "**Undefined**"
Type:         "hosts"
URL:          "https://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml&showintro=1&mimetype=plaintext"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"
`

	expPre = `
Desc:         "pre-configured blacklisted subdomains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "blacklisted-subdomains"
Name:         "blacklisted-subdomains"
nType:        "preDomn"
Prefix:       "**Undefined**"
Type:         "blacklisted-subdomains"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "adsrvr.org"
              "adtechus.net"
              "advertising.com"
              "centade.com"
              "doubleclick.net"
              "free-counter.co.uk"
              "intellitxt.com"
              "kiosked.com"
              "patoghee.in"

Desc:         "pre-configured blacklisted servers"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "blacklisted-servers"
Name:         "blacklisted-servers"
nType:        "preHost"
Prefix:       "**Undefined**"
Type:         "blacklisted-servers"
URL:          "**Undefined**"
Whitelist:
              "**No entries found**"
Blacklist:
              "beap.gemini.yahoo.com"
`

	expURLS = `
Desc:         "List of zones serving malicious executables observed by malc0de.com/database/"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.168.1"
Ltype:        "url"
Name:         "malc0de"
nType:        "domn"
Prefix:       "zone "
Type:         "domains"
URL:          "http://malc0de.com/bl/ZONES"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Just domains"
Disabled:     "false"
File:         "**Undefined**"
IP:           "10.0.0.1"
Ltype:        "url"
Name:         "malwaredomains.com"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "http://mirror1.malwaredomains.com/files/justdomains"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Basic tracking list by Disconnect"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "url"
Name:         "simple_tracking"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "abuse.ch ZeuS domain blocklist"
Disabled:     "false"
File:         "**Undefined**"
IP:           "192.168.100.1"
Ltype:        "url"
Name:         "zeus"
nType:        "domn"
Prefix:       "**Undefined**"
Type:         "domains"
URL:          "https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "OpenPhish automatic phishing detection"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "openphish"
nType:        "host"
Prefix:       "http"
Type:         "hosts"
URL:          "https://openphish.com/feed.txt"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "This hosts file is a merged collection of hosts from reputable sources"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "raw.github.com"
nType:        "host"
Prefix:       "0.0.0.0 "
Type:         "hosts"
URL:          "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "This hosts file is a merged collection of hosts from cameleon"
Disabled:     "false"
File:         "**Undefined**"
IP:           "172.16.16.1"
Ltype:        "url"
Name:         "sysctl.org"
nType:        "host"
Prefix:       "127.0.0.1\t "
Type:         "hosts"
URL:          "http://sysctl.org/cameleon/hosts"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Ad server blacklists"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "volkerschatz"
nType:        "host"
Prefix:       "http"
Type:         "hosts"
URL:          "http://www.volkerschatz.com/net/adpaths"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"

Desc:         "Fully Qualified Domain Names only - no prefix to strip"
Disabled:     "false"
File:         "**Undefined**"
IP:           "0.0.0.0"
Ltype:        "url"
Name:         "yoyo"
nType:        "host"
Prefix:       "**Undefined**"
Type:         "hosts"
URL:          "https://pgl.yoyo.org/as/serverlist.php?hostformat=nohtml&showintro=1&mimetype=plaintext"
Whitelist:
              "**No entries found**"
Blacklist:
              "**No entries found**"
`
)
