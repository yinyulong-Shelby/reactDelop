package vi

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	vietnamese "github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {
	vie := vietnamese.New()
	uni := ut.New(vie, vie)
	trans, _ := uni.GetTranslator("vi")

	validate := validator.New()

	err := RegisterDefaultTranslations(validate, trans)
	Equal(t, err, nil)

	type Inner struct {
		EqCSFieldString  string
		NeCSFieldString  string
		GtCSFieldString  string
		GteCSFieldString string
		LtCSFieldString  string
		LteCSFieldString string
	}

	type Test struct {
		Inner             Inner
		RequiredString    string            `validate:"required"`
		RequiredNumber    int               `validate:"required"`
		RequiredMultiple  []string          `validate:"required"`
		LenString         string            `validate:"len=1"`
		LenNumber         float64           `validate:"len=1113.00"`
		LenMultiple       []string          `validate:"len=7"`
		MinString         string            `validate:"min=1"`
		MinNumber         float64           `validate:"min=1113.00"`
		MinMultiple       []string          `validate:"min=7"`
		MaxString         string            `validate:"max=3"`
		MaxNumber         float64           `validate:"max=1113.00"`
		MaxMultiple       []string          `validate:"max=7"`
		EqString          string            `validate:"eq=3"`
		EqNumber          float64           `validate:"eq=2.33"`
		EqMultiple        []string          `validate:"eq=7"`
		NeString          string            `validate:"ne="`
		NeNumber          float64           `validate:"ne=0.00"`
		NeMultiple        []string          `validate:"ne=0"`
		LtString          string            `validate:"lt=3"`
		LtNumber          float64           `validate:"lt=5.56"`
		LtMultiple        []string          `validate:"lt=2"`
		LtTime            time.Time         `validate:"lt"`
		LteString         string            `validate:"lte=3"`
		LteNumber         float64           `validate:"lte=5.56"`
		LteMultiple       []string          `validate:"lte=2"`
		LteTime           time.Time         `validate:"lte"`
		GtString          string            `validate:"gt=3"`
		GtNumber          float64           `validate:"gt=5.56"`
		GtMultiple        []string          `validate:"gt=2"`
		GtTime            time.Time         `validate:"gt"`
		GteString         string            `validate:"gte=3"`
		GteNumber         float64           `validate:"gte=5.56"`
		GteMultiple       []string          `validate:"gte=2"`
		GteTime           time.Time         `validate:"gte"`
		EqFieldString     string            `validate:"eqfield=MaxString"`
		EqCSFieldString   string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string            `validate:"nefield=EqFieldString"`
		GtFieldString     string            `validate:"gtfield=MaxString"`
		GteFieldString    string            `validate:"gtefield=MaxString"`
		LtFieldString     string            `validate:"ltfield=MaxString"`
		LteFieldString    string            `validate:"ltefield=MaxString"`
		AlphaString       string            `validate:"alpha"`
		AlphanumString    string            `validate:"alphanum"`
		NumericString     string            `validate:"numeric"`
		NumberString      string            `validate:"number"`
		HexadecimalString string            `validate:"hexadecimal"`
		HexColorString    string            `validate:"hexcolor"`
		RGBColorString    string            `validate:"rgb"`
		RGBAColorString   string            `validate:"rgba"`
		HSLColorString    string            `validate:"hsl"`
		HSLAColorString   string            `validate:"hsla"`
		Email             string            `validate:"email"`
		URL               string            `validate:"url"`
		URI               string            `validate:"uri"`
		Base64            string            `validate:"base64"`
		Contains          string            `validate:"contains=purpose"`
		ContainsAny       string            `validate:"containsany=!@#$"`
		Excludes          string            `validate:"excludes=text"`
		ExcludesAll       string            `validate:"excludesall=!@#$"`
		ExcludesRune      string            `validate:"excludesrune=???"`
		ISBN              string            `validate:"isbn"`
		ISBN10            string            `validate:"isbn10"`
		ISBN13            string            `validate:"isbn13"`
		UUID              string            `validate:"uuid"`
		UUID3             string            `validate:"uuid3"`
		UUID4             string            `validate:"uuid4"`
		UUID5             string            `validate:"uuid5"`
		ASCII             string            `validate:"ascii"`
		PrintableASCII    string            `validate:"printascii"`
		MultiByte         string            `validate:"multibyte"`
		DataURI           string            `validate:"datauri"`
		Latitude          string            `validate:"latitude"`
		Longitude         string            `validate:"longitude"`
		SSN               string            `validate:"ssn"`
		IP                string            `validate:"ip"`
		IPv4              string            `validate:"ipv4"`
		IPv6              string            `validate:"ipv6"`
		CIDR              string            `validate:"cidr"`
		CIDRv4            string            `validate:"cidrv4"`
		CIDRv6            string            `validate:"cidrv6"`
		TCPAddr           string            `validate:"tcp_addr"`
		TCPAddrv4         string            `validate:"tcp4_addr"`
		TCPAddrv6         string            `validate:"tcp6_addr"`
		UDPAddr           string            `validate:"udp_addr"`
		UDPAddrv4         string            `validate:"udp4_addr"`
		UDPAddrv6         string            `validate:"udp6_addr"`
		IPAddr            string            `validate:"ip_addr"`
		IPAddrv4          string            `validate:"ip4_addr"`
		IPAddrv6          string            `validate:"ip6_addr"`
		UinxAddr          string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string            `validate:"mac"`
		IsColor           string            `validate:"iscolor"`
		StrPtrMinLen      *string           `validate:"min=10"`
		StrPtrMaxLen      *string           `validate:"max=1"`
		StrPtrLen         *string           `validate:"len=2"`
		StrPtrLt          *string           `validate:"lt=1"`
		StrPtrLte         *string           `validate:"lte=1"`
		StrPtrGt          *string           `validate:"gt=10"`
		StrPtrGte         *string           `validate:"gte=10"`
		OneOfString       string            `validate:"oneof=red green"`
		OneOfInt          int               `validate:"oneof=5 63"`
		UniqueSlice       []string          `validate:"unique"`
		UniqueArray       [3]string         `validate:"unique"`
		UniqueMap         map[string]string `validate:"unique"`
		JSONString        string            `validate:"json"`
		JWTString         string            `validate:"jwt"`
		LowercaseString   string            `validate:"lowercase"`
		UppercaseString   string            `validate:"uppercase"`
		Datetime          string            `validate:"datetime=2006-01-02"`
		PostCode          string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry   string
		PostCodeByField   string `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"

	test.MaxString = "1234"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)

	test.LtString = "1234"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteTime = time.Now().Add(time.Hour * 24)

	test.LtFieldString = "12345"
	test.LteFieldString = "12345"

	test.LtCSFieldString = "1234"
	test.LteCSFieldString = "1234"

	test.AlphaString = "abc3"
	test.AlphanumString = "abc3!"
	test.NumericString = "12E.00"
	test.NumberString = "12E"

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ???"

	test.ASCII = "????????????"
	test.PrintableASCII = "????????????"

	test.MultiByte = "1234feerf"

	test.LowercaseString = "ABCDEFG"
	test.UppercaseString = "abcdefg"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}
	test.Datetime = "2008-Feb-01"

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(validator.ValidationErrors)
	Equal(t, ok, true)

	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "Test.IsColor",
			expected: "IsColor ph???i l?? m??u s???c h???p l???",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC ch??? ???????c ch???a ?????a ch??? MAC",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr ph???i l?? ?????a ch??? IP c?? th??? ph??n gi???i",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 ph???i l?? ?????a ch??? IPv4 c?? th??? ph??n gi???i",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 ph???i l?? ?????a ch??? IPv6 c?? th??? ph??n gi???i",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr ph???i l?? ?????a ch??? UDP",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 ph???i l?? ?????a ch??? IPv4 UDP",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 ph???i l?? ?????a ch??? IPv6 UDP",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr ph???i l?? ?????a ch??? TCP",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 ph???i l?? ?????a ch??? IPv4 TCP",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 ph???i l?? ?????a ch??? IPv6 TCP",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR ch??? ???????c ch???a CIDR notation",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 ch??? ???????c ch???a CIDR notation c???a m???t ?????a ch??? IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 ch??? ???????c ch???a CIDR notation c???a m???t ?????a ch??? IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN ph???i l?? SSN number",
		},
		{
			ns:       "Test.IP",
			expected: "IP ph???i l?? ?????a ch??? IP",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 ph???i l?? ?????a ch??? IPv4",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 ph???i l?? ?????a ch??? IPv6",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI ch??? ???????c ch???a Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude ch??? ???????c ch???a latitude (v??? ?????)",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude ch??? ???????c ch???a longitude (kinh ?????)",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte ch??? ???????c ch???a k?? t??? multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII ch??? ???????c ch???a k?? t??? ASCII",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII ch??? ???????c ch???a k?? t??? ASCII c?? th??? in ???n",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID ph???i l?? gi?? tr??? UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 ph???i l?? gi?? tr??? UUID phi??n b???n 3",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 ph???i l?? gi?? tr??? UUID phi??n b???n 4",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 ph???i l?? gi?? tr??? UUID phi??n b???n 5",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN ph???i l?? s??? ISBN",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 ph???i l?? s??? ISBN-10",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 ph???i l?? s??? ISBN-13",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes kh??ng ???????c ch???a chu???i 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll kh??ng ???????c ch???a b???t k??? k?? t??? n??o trong nh??m k?? t??? '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune kh??ng ???????c ch???a '???'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny ph???i ch???a ??t nh???t 1 trong c??ch k?? t??? sau '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains ph???i ch???a chu???i 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 ph???i l?? gi?? tr??? chu???i Base64",
		},
		{
			ns:       "Test.Email",
			expected: "Email ph???i l?? gi?? tr??? email address",
		},
		{
			ns:       "Test.URL",
			expected: "URL ph???i l?? gi?? tr??? URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI ph???i l?? gi?? tr??? URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString ph???i l?? gi?? tr??? RGB color",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString ph???i l?? gi?? tr??? RGBA color",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString ph???i l?? gi?? tr??? HSL color",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString ph???i l?? gi?? tr??? HSLA color",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString ph???i l?? gi?? tr??? hexadecimal",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString ph???i l?? gi?? tr??? HEX color",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString ch??? ???????c ch???a gi?? tr??? s???",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString ch??? ???????c ch???a gi?? tr??? s??? ho???c s??? d?????i d???ng ch???",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString ch??? ???????c ch???a k?? t??? d???ng alphanumeric",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString ch??? ???????c ch???a k?? t??? d???ng alphabetic",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString ch??? ???????c nh??? h??n MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString ch??? ???????c nh??? h??n ho???c b???ng MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString ph???i l???n h??n MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString ph???i l???n h??n ho???c b???ng MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString kh??ng ???????c ph??p b???ng EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString ch??? ???????c nh??? h??n Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString ch??? ???????c nh??? h??n ho???c b???ng Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString ph???i l???n h??n Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString ph???i l???n h??n ho???c b???ng Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString kh??ng ???????c ph??p b???ng Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString ph???i b???ng Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString ph???i b???ng MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString ph???i c?? ????? d??i ??t nh???t 3 k?? t???",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber ph???i l?? 5,56 ho???c l???n h??n",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple ph???i ch???a ??t nh???t 2 ph???n t???",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime ph???i l???n h??n ho???c b???ng Ng??y & Gi??? hi???n t???i",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString ph???i c?? ????? d??i l???n h??n 3 k?? t???",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber ph???i l???n h??n 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple ph???i ch???a nhi???u h??n 2 ph???n t???",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime ph???i l???n h??n Ng??y & Gi??? hi???n t???i",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString ch??? ???????c c?? ????? d??i t???i ??a l?? 3 k?? t???",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber ph???i l?? 5,56 ho???c nh??? h??n",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple ch??? ???????c ch???a nhi???u nh???t 2 ph???n t???",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime ch??? ???????c nh??? h??n ho???c b???ng Ng??y & Gi??? hi???n t???i",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString ph???i c?? ????? d??i nh??? h??n 3 k?? t???",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber ph???i nh??? h??n 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple ch??? ???????c ch???a ??t h??n 2 ph???n t???",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime ph???i nh??? h??n Ng??y & Gi??? hi???n t???i",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString kh??ng ???????c b???ng ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber kh??ng ???????c b???ng 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple kh??ng ???????c b???ng 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString kh??ng b???ng 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber kh??ng b???ng 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple kh??ng b???ng 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString ch??? ???????c ch???a t???i ??a 3 k?? t???",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber ph???i l?? 1.113,00 ho???c nh??? h??n",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple ch??? ???????c ch???a t???i ??a 7 ph???n t???",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString ph???i ch???a ??t nh???t 1 k?? t???",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber ph???i b???ng 1.113,00 ho???c l???n h??n",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple ph???i ch???a ??t nh???t 7 ph???n t???",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString ph???i c?? ????? d??i l?? 1 k?? t???",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber ph???i b???ng 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple ph???i ch???a 7 ph???n t???",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString kh??ng ???????c b??? tr???ng",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber kh??ng ???????c b??? tr???ng",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple kh??ng ???????c b??? tr???ng",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen ph???i ch???a ??t nh???t 10 k?? t???",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen ch??? ???????c ch???a t???i ??a 1 k?? t???",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen ph???i c?? ????? d??i l?? 2 k?? t???",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt ph???i c?? ????? d??i nh??? h??n 1 k?? t???",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte ch??? ???????c c?? ????? d??i t???i ??a l?? 1 k?? t???",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt ph???i c?? ????? d??i l???n h??n 10 k?? t???",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte ph???i c?? ????? d??i ??t nh???t 10 k?? t???",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString ph???i l?? trong nh???ng gi?? tr??? [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt ph???i l?? trong nh???ng gi?? tr??? [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice ch??? ???????c ch???a nh???ng gi?? tr??? kh??ng tr??ng l???p",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray ch??? ???????c ch???a nh???ng gi?? tr??? kh??ng tr??ng l???p",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap ch??? ???????c ch???a nh???ng gi?? tr??? kh??ng tr??ng l???p",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString ph???i l?? m???t chu???i json h???p l???",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString ph???i l?? m???t chu???i jwt h???p l???",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString ph???i ???????c vi???t th?????ng",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString ph???i ???????c vi???t hoa",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime kh??ng tr??ng ?????nh d???ng ng??y th??ng 2006-01-02",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode sai ?????nh d???ng postcode c???a qu???c gia SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField sai ?????nh d???ng postcode c???a qu???c gia t????ng ???ng thu???c tr?????ng PostCodeCountry",
		},
	}

	for _, tt := range tests {

		var fe validator.FieldError

		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fe = e
				break
			}
		}

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}
}
