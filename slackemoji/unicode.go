package slackemoji

import "fmt"

func UnicodeURL(colonSyntax string) string {
	unicode, ok := colonSyntaxToUnicode[colonSyntax]
	if !ok {
		return ""
	}

	const (
		size   = "large"
		cdnURL = "https://a.slack-edge.com/production-standard-emoji-assets/14.0"
	)

	return fmt.Sprintf("%s/apple-%s/%s.png", cdnURL, size, unicode)
}

var colonSyntaxToUnicode = map[string]string{
	// 0023-FE0F-20E3 hash
	":hash:": "0023-fe0f-20e3",
	// 002A-FE0F-20E3 keycap_star
	":keycap_star:": "002a-fe0f-20e3",
	// 0030-FE0F-20E3 zero
	":zero:": "0030-fe0f-20e3",
	// 0031-FE0F-20E3 one
	":one:": "0031-fe0f-20e3",
	// 0032-FE0F-20E3 two
	":two:": "0032-fe0f-20e3",
	// 0033-FE0F-20E3 three
	":three:": "0033-fe0f-20e3",
	// 0034-FE0F-20E3 four
	":four:": "0034-fe0f-20e3",
	// 0035-FE0F-20E3 five
	":five:": "0035-fe0f-20e3",
	// 0036-FE0F-20E3 six
	":six:": "0036-fe0f-20e3",
	// 0037-FE0F-20E3 seven
	":seven:": "0037-fe0f-20e3",
	// 0038-FE0F-20E3 eight
	":eight:": "0038-fe0f-20e3",
	// 0039-FE0F-20E3 nine
	":nine:": "0039-fe0f-20e3",
	// 00A9-FE0F copyright
	":copyright:": "00a9-fe0f",
	// 00AE-FE0F registered
	":registered:": "00ae-fe0f",
	// 1F004 mahjong
	":mahjong:": "1f004",
	// 1F0CF black_joker
	":black_joker:": "1f0cf",
	// 1F170-FE0F a
	":a:": "1f170-fe0f",
	// 1F171-FE0F b
	":b:": "1f171-fe0f",
	// 1F17E-FE0F o2
	":o2:": "1f17e-fe0f",
	// 1F17F-FE0F parking
	":parking:": "1f17f-fe0f",
	// 1F18E ab
	":ab:": "1f18e",
	// 1F191 cl
	":cl:": "1f191",
	// 1F192 cool
	":cool:": "1f192",
	// 1F193 free
	":free:": "1f193",
	// 1F194 id
	":id:": "1f194",
	// 1F195 new
	":new:": "1f195",
	// 1F196 ng
	":ng:": "1f196",
	// 1F197 ok
	":ok:": "1f197",
	// 1F198 sos
	":sos:": "1f198",
	// 1F199 up
	":up:": "1f199",
	// 1F19A vs
	":vs:": "1f19a",
	// 1F1E6-1F1E8 flag-ac
	":flag-ac:": "1f1e6-1f1e8",
	// 1F1E6-1F1E9 flag-ad
	":flag-ad:": "1f1e6-1f1e9",
	// 1F1E6-1F1EA flag-ae
	":flag-ae:": "1f1e6-1f1ea",
	// 1F1E6-1F1EB flag-af
	":flag-af:": "1f1e6-1f1eb",
	// 1F1E6-1F1EC flag-ag
	":flag-ag:": "1f1e6-1f1ec",
	// 1F1E6-1F1EE flag-ai
	":flag-ai:": "1f1e6-1f1ee",
	// 1F1E6-1F1F1 flag-al
	":flag-al:": "1f1e6-1f1f1",
	// 1F1E6-1F1F2 flag-am
	":flag-am:": "1f1e6-1f1f2",
	// 1F1E6-1F1F4 flag-ao
	":flag-ao:": "1f1e6-1f1f4",
	// 1F1E6-1F1F6 flag-aq
	":flag-aq:": "1f1e6-1f1f6",
	// 1F1E6-1F1F7 flag-ar
	":flag-ar:": "1f1e6-1f1f7",
	// 1F1E6-1F1F8 flag-as
	":flag-as:": "1f1e6-1f1f8",
	// 1F1E6-1F1F9 flag-at
	":flag-at:": "1f1e6-1f1f9",
	// 1F1E6-1F1FA flag-au
	":flag-au:": "1f1e6-1f1fa",
	// 1F1E6-1F1FC flag-aw
	":flag-aw:": "1f1e6-1f1fc",
	// 1F1E6-1F1FD flag-ax
	":flag-ax:": "1f1e6-1f1fd",
	// 1F1E6-1F1FF flag-az
	":flag-az:": "1f1e6-1f1ff",
	// 1F1E7-1F1E6 flag-ba
	":flag-ba:": "1f1e7-1f1e6",
	// 1F1E7-1F1E7 flag-bb
	":flag-bb:": "1f1e7-1f1e7",
	// 1F1E7-1F1E9 flag-bd
	":flag-bd:": "1f1e7-1f1e9",
	// 1F1E7-1F1EA flag-be
	":flag-be:": "1f1e7-1f1ea",
	// 1F1E7-1F1EB flag-bf
	":flag-bf:": "1f1e7-1f1eb",
	// 1F1E7-1F1EC flag-bg
	":flag-bg:": "1f1e7-1f1ec",
	// 1F1E7-1F1ED flag-bh
	":flag-bh:": "1f1e7-1f1ed",
	// 1F1E7-1F1EE flag-bi
	":flag-bi:": "1f1e7-1f1ee",
	// 1F1E7-1F1EF flag-bj
	":flag-bj:": "1f1e7-1f1ef",
	// 1F1E7-1F1F1 flag-bl
	":flag-bl:": "1f1e7-1f1f1",
	// 1F1E7-1F1F2 flag-bm
	":flag-bm:": "1f1e7-1f1f2",
	// 1F1E7-1F1F3 flag-bn
	":flag-bn:": "1f1e7-1f1f3",
	// 1F1E7-1F1F4 flag-bo
	":flag-bo:": "1f1e7-1f1f4",
	// 1F1E7-1F1F6 flag-bq
	":flag-bq:": "1f1e7-1f1f6",
	// 1F1E7-1F1F7 flag-br
	":flag-br:": "1f1e7-1f1f7",
	// 1F1E7-1F1F8 flag-bs
	":flag-bs:": "1f1e7-1f1f8",
	// 1F1E7-1F1F9 flag-bt
	":flag-bt:": "1f1e7-1f1f9",
	// 1F1E7-1F1FB flag-bv
	":flag-bv:": "1f1e7-1f1fb",
	// 1F1E7-1F1FC flag-bw
	":flag-bw:": "1f1e7-1f1fc",
	// 1F1E7-1F1FE flag-by
	":flag-by:": "1f1e7-1f1fe",
	// 1F1E7-1F1FF flag-bz
	":flag-bz:": "1f1e7-1f1ff",
	// 1F1E8-1F1E6 flag-ca
	":flag-ca:": "1f1e8-1f1e6",
	// 1F1E8-1F1E8 flag-cc
	":flag-cc:": "1f1e8-1f1e8",
	// 1F1E8-1F1E9 flag-cd
	":flag-cd:": "1f1e8-1f1e9",
	// 1F1E8-1F1EB flag-cf
	":flag-cf:": "1f1e8-1f1eb",
	// 1F1E8-1F1EC flag-cg
	":flag-cg:": "1f1e8-1f1ec",
	// 1F1E8-1F1ED flag-ch
	":flag-ch:": "1f1e8-1f1ed",
	// 1F1E8-1F1EE flag-ci
	":flag-ci:": "1f1e8-1f1ee",
	// 1F1E8-1F1F0 flag-ck
	":flag-ck:": "1f1e8-1f1f0",
	// 1F1E8-1F1F1 flag-cl
	":flag-cl:": "1f1e8-1f1f1",
	// 1F1E8-1F1F2 flag-cm
	":flag-cm:": "1f1e8-1f1f2",
	// 1F1E8-1F1F3 cn
	":cn:": "1f1e8-1f1f3",
	// 1F1E8-1F1F3 cn
	":flag-cn:": "1f1e8-1f1f3",
	// 1F1E8-1F1F4 flag-co
	":flag-co:": "1f1e8-1f1f4",
	// 1F1E8-1F1F5 flag-cp
	":flag-cp:": "1f1e8-1f1f5",
	// 1F1E8-1F1F7 flag-cr
	":flag-cr:": "1f1e8-1f1f7",
	// 1F1E8-1F1FA flag-cu
	":flag-cu:": "1f1e8-1f1fa",
	// 1F1E8-1F1FB flag-cv
	":flag-cv:": "1f1e8-1f1fb",
	// 1F1E8-1F1FC flag-cw
	":flag-cw:": "1f1e8-1f1fc",
	// 1F1E8-1F1FD flag-cx
	":flag-cx:": "1f1e8-1f1fd",
	// 1F1E8-1F1FE flag-cy
	":flag-cy:": "1f1e8-1f1fe",
	// 1F1E8-1F1FF flag-cz
	":flag-cz:": "1f1e8-1f1ff",
	// 1F1E9-1F1EA de
	":de:": "1f1e9-1f1ea",
	// 1F1E9-1F1EA de
	":flag-de:": "1f1e9-1f1ea",
	// 1F1E9-1F1EC flag-dg
	":flag-dg:": "1f1e9-1f1ec",
	// 1F1E9-1F1EF flag-dj
	":flag-dj:": "1f1e9-1f1ef",
	// 1F1E9-1F1F0 flag-dk
	":flag-dk:": "1f1e9-1f1f0",
	// 1F1E9-1F1F2 flag-dm
	":flag-dm:": "1f1e9-1f1f2",
	// 1F1E9-1F1F4 flag-do
	":flag-do:": "1f1e9-1f1f4",
	// 1F1E9-1F1FF flag-dz
	":flag-dz:": "1f1e9-1f1ff",
	// 1F1EA-1F1E6 flag-ea
	":flag-ea:": "1f1ea-1f1e6",
	// 1F1EA-1F1E8 flag-ec
	":flag-ec:": "1f1ea-1f1e8",
	// 1F1EA-1F1EA flag-ee
	":flag-ee:": "1f1ea-1f1ea",
	// 1F1EA-1F1EC flag-eg
	":flag-eg:": "1f1ea-1f1ec",
	// 1F1EA-1F1ED flag-eh
	":flag-eh:": "1f1ea-1f1ed",
	// 1F1EA-1F1F7 flag-er
	":flag-er:": "1f1ea-1f1f7",
	// 1F1EA-1F1F8 es
	":es:": "1f1ea-1f1f8",
	// 1F1EA-1F1F8 es
	":flag-es:": "1f1ea-1f1f8",
	// 1F1EA-1F1F9 flag-et
	":flag-et:": "1f1ea-1f1f9",
	// 1F1EA-1F1FA flag-eu
	":flag-eu:": "1f1ea-1f1fa",
	// 1F1EB-1F1EE flag-fi
	":flag-fi:": "1f1eb-1f1ee",
	// 1F1EB-1F1EF flag-fj
	":flag-fj:": "1f1eb-1f1ef",
	// 1F1EB-1F1F0 flag-fk
	":flag-fk:": "1f1eb-1f1f0",
	// 1F1EB-1F1F2 flag-fm
	":flag-fm:": "1f1eb-1f1f2",
	// 1F1EB-1F1F4 flag-fo
	":flag-fo:": "1f1eb-1f1f4",
	// 1F1EB-1F1F7 fr
	":fr:": "1f1eb-1f1f7",
	// 1F1EB-1F1F7 fr
	":flag-fr:": "1f1eb-1f1f7",
	// 1F1EC-1F1E6 flag-ga
	":flag-ga:": "1f1ec-1f1e6",
	// 1F1EC-1F1E7 gb
	":gb:": "1f1ec-1f1e7",
	// 1F1EC-1F1E7 gb
	":uk:": "1f1ec-1f1e7",
	// 1F1EC-1F1E7 gb
	":flag-gb:": "1f1ec-1f1e7",
	// 1F1EC-1F1E9 flag-gd
	":flag-gd:": "1f1ec-1f1e9",
	// 1F1EC-1F1EA flag-ge
	":flag-ge:": "1f1ec-1f1ea",
	// 1F1EC-1F1EB flag-gf
	":flag-gf:": "1f1ec-1f1eb",
	// 1F1EC-1F1EC flag-gg
	":flag-gg:": "1f1ec-1f1ec",
	// 1F1EC-1F1ED flag-gh
	":flag-gh:": "1f1ec-1f1ed",
	// 1F1EC-1F1EE flag-gi
	":flag-gi:": "1f1ec-1f1ee",
	// 1F1EC-1F1F1 flag-gl
	":flag-gl:": "1f1ec-1f1f1",
	// 1F1EC-1F1F2 flag-gm
	":flag-gm:": "1f1ec-1f1f2",
	// 1F1EC-1F1F3 flag-gn
	":flag-gn:": "1f1ec-1f1f3",
	// 1F1EC-1F1F5 flag-gp
	":flag-gp:": "1f1ec-1f1f5",
	// 1F1EC-1F1F6 flag-gq
	":flag-gq:": "1f1ec-1f1f6",
	// 1F1EC-1F1F7 flag-gr
	":flag-gr:": "1f1ec-1f1f7",
	// 1F1EC-1F1F8 flag-gs
	":flag-gs:": "1f1ec-1f1f8",
	// 1F1EC-1F1F9 flag-gt
	":flag-gt:": "1f1ec-1f1f9",
	// 1F1EC-1F1FA flag-gu
	":flag-gu:": "1f1ec-1f1fa",
	// 1F1EC-1F1FC flag-gw
	":flag-gw:": "1f1ec-1f1fc",
	// 1F1EC-1F1FE flag-gy
	":flag-gy:": "1f1ec-1f1fe",
	// 1F1ED-1F1F0 flag-hk
	":flag-hk:": "1f1ed-1f1f0",
	// 1F1ED-1F1F2 flag-hm
	":flag-hm:": "1f1ed-1f1f2",
	// 1F1ED-1F1F3 flag-hn
	":flag-hn:": "1f1ed-1f1f3",
	// 1F1ED-1F1F7 flag-hr
	":flag-hr:": "1f1ed-1f1f7",
	// 1F1ED-1F1F9 flag-ht
	":flag-ht:": "1f1ed-1f1f9",
	// 1F1ED-1F1FA flag-hu
	":flag-hu:": "1f1ed-1f1fa",
	// 1F1EE-1F1E8 flag-ic
	":flag-ic:": "1f1ee-1f1e8",
	// 1F1EE-1F1E9 flag-id
	":flag-id:": "1f1ee-1f1e9",
	// 1F1EE-1F1EA flag-ie
	":flag-ie:": "1f1ee-1f1ea",
	// 1F1EE-1F1F1 flag-il
	":flag-il:": "1f1ee-1f1f1",
	// 1F1EE-1F1F2 flag-im
	":flag-im:": "1f1ee-1f1f2",
	// 1F1EE-1F1F3 flag-in
	":flag-in:": "1f1ee-1f1f3",
	// 1F1EE-1F1F4 flag-io
	":flag-io:": "1f1ee-1f1f4",
	// 1F1EE-1F1F6 flag-iq
	":flag-iq:": "1f1ee-1f1f6",
	// 1F1EE-1F1F7 flag-ir
	":flag-ir:": "1f1ee-1f1f7",
	// 1F1EE-1F1F8 flag-is
	":flag-is:": "1f1ee-1f1f8",
	// 1F1EE-1F1F9 it
	":it:": "1f1ee-1f1f9",
	// 1F1EE-1F1F9 it
	":flag-it:": "1f1ee-1f1f9",
	// 1F1EF-1F1EA flag-je
	":flag-je:": "1f1ef-1f1ea",
	// 1F1EF-1F1F2 flag-jm
	":flag-jm:": "1f1ef-1f1f2",
	// 1F1EF-1F1F4 flag-jo
	":flag-jo:": "1f1ef-1f1f4",
	// 1F1EF-1F1F5 jp
	":jp:": "1f1ef-1f1f5",
	// 1F1EF-1F1F5 jp
	":flag-jp:": "1f1ef-1f1f5",
	// 1F1F0-1F1EA flag-ke
	":flag-ke:": "1f1f0-1f1ea",
	// 1F1F0-1F1EC flag-kg
	":flag-kg:": "1f1f0-1f1ec",
	// 1F1F0-1F1ED flag-kh
	":flag-kh:": "1f1f0-1f1ed",
	// 1F1F0-1F1EE flag-ki
	":flag-ki:": "1f1f0-1f1ee",
	// 1F1F0-1F1F2 flag-km
	":flag-km:": "1f1f0-1f1f2",
	// 1F1F0-1F1F3 flag-kn
	":flag-kn:": "1f1f0-1f1f3",
	// 1F1F0-1F1F5 flag-kp
	":flag-kp:": "1f1f0-1f1f5",
	// 1F1F0-1F1F7 kr
	":kr:": "1f1f0-1f1f7",
	// 1F1F0-1F1F7 kr
	":flag-kr:": "1f1f0-1f1f7",
	// 1F1F0-1F1FC flag-kw
	":flag-kw:": "1f1f0-1f1fc",
	// 1F1F0-1F1FE flag-ky
	":flag-ky:": "1f1f0-1f1fe",
	// 1F1F0-1F1FF flag-kz
	":flag-kz:": "1f1f0-1f1ff",
	// 1F1F1-1F1E6 flag-la
	":flag-la:": "1f1f1-1f1e6",
	// 1F1F1-1F1E7 flag-lb
	":flag-lb:": "1f1f1-1f1e7",
	// 1F1F1-1F1E8 flag-lc
	":flag-lc:": "1f1f1-1f1e8",
	// 1F1F1-1F1EE flag-li
	":flag-li:": "1f1f1-1f1ee",
	// 1F1F1-1F1F0 flag-lk
	":flag-lk:": "1f1f1-1f1f0",
	// 1F1F1-1F1F7 flag-lr
	":flag-lr:": "1f1f1-1f1f7",
	// 1F1F1-1F1F8 flag-ls
	":flag-ls:": "1f1f1-1f1f8",
	// 1F1F1-1F1F9 flag-lt
	":flag-lt:": "1f1f1-1f1f9",
	// 1F1F1-1F1FA flag-lu
	":flag-lu:": "1f1f1-1f1fa",
	// 1F1F1-1F1FB flag-lv
	":flag-lv:": "1f1f1-1f1fb",
	// 1F1F1-1F1FE flag-ly
	":flag-ly:": "1f1f1-1f1fe",
	// 1F1F2-1F1E6 flag-ma
	":flag-ma:": "1f1f2-1f1e6",
	// 1F1F2-1F1E8 flag-mc
	":flag-mc:": "1f1f2-1f1e8",
	// 1F1F2-1F1E9 flag-md
	":flag-md:": "1f1f2-1f1e9",
	// 1F1F2-1F1EA flag-me
	":flag-me:": "1f1f2-1f1ea",
	// 1F1F2-1F1EB flag-mf
	":flag-mf:": "1f1f2-1f1eb",
	// 1F1F2-1F1EC flag-mg
	":flag-mg:": "1f1f2-1f1ec",
	// 1F1F2-1F1ED flag-mh
	":flag-mh:": "1f1f2-1f1ed",
	// 1F1F2-1F1F0 flag-mk
	":flag-mk:": "1f1f2-1f1f0",
	// 1F1F2-1F1F1 flag-ml
	":flag-ml:": "1f1f2-1f1f1",
	// 1F1F2-1F1F2 flag-mm
	":flag-mm:": "1f1f2-1f1f2",
	// 1F1F2-1F1F3 flag-mn
	":flag-mn:": "1f1f2-1f1f3",
	// 1F1F2-1F1F4 flag-mo
	":flag-mo:": "1f1f2-1f1f4",
	// 1F1F2-1F1F5 flag-mp
	":flag-mp:": "1f1f2-1f1f5",
	// 1F1F2-1F1F6 flag-mq
	":flag-mq:": "1f1f2-1f1f6",
	// 1F1F2-1F1F7 flag-mr
	":flag-mr:": "1f1f2-1f1f7",
	// 1F1F2-1F1F8 flag-ms
	":flag-ms:": "1f1f2-1f1f8",
	// 1F1F2-1F1F9 flag-mt
	":flag-mt:": "1f1f2-1f1f9",
	// 1F1F2-1F1FA flag-mu
	":flag-mu:": "1f1f2-1f1fa",
	// 1F1F2-1F1FB flag-mv
	":flag-mv:": "1f1f2-1f1fb",
	// 1F1F2-1F1FC flag-mw
	":flag-mw:": "1f1f2-1f1fc",
	// 1F1F2-1F1FD flag-mx
	":flag-mx:": "1f1f2-1f1fd",
	// 1F1F2-1F1FE flag-my
	":flag-my:": "1f1f2-1f1fe",
	// 1F1F2-1F1FF flag-mz
	":flag-mz:": "1f1f2-1f1ff",
	// 1F1F3-1F1E6 flag-na
	":flag-na:": "1f1f3-1f1e6",
	// 1F1F3-1F1E8 flag-nc
	":flag-nc:": "1f1f3-1f1e8",
	// 1F1F3-1F1EA flag-ne
	":flag-ne:": "1f1f3-1f1ea",
	// 1F1F3-1F1EB flag-nf
	":flag-nf:": "1f1f3-1f1eb",
	// 1F1F3-1F1EC flag-ng
	":flag-ng:": "1f1f3-1f1ec",
	// 1F1F3-1F1EE flag-ni
	":flag-ni:": "1f1f3-1f1ee",
	// 1F1F3-1F1F1 flag-nl
	":flag-nl:": "1f1f3-1f1f1",
	// 1F1F3-1F1F4 flag-no
	":flag-no:": "1f1f3-1f1f4",
	// 1F1F3-1F1F5 flag-np
	":flag-np:": "1f1f3-1f1f5",
	// 1F1F3-1F1F7 flag-nr
	":flag-nr:": "1f1f3-1f1f7",
	// 1F1F3-1F1FA flag-nu
	":flag-nu:": "1f1f3-1f1fa",
	// 1F1F3-1F1FF flag-nz
	":flag-nz:": "1f1f3-1f1ff",
	// 1F1F4-1F1F2 flag-om
	":flag-om:": "1f1f4-1f1f2",
	// 1F1F5-1F1E6 flag-pa
	":flag-pa:": "1f1f5-1f1e6",
	// 1F1F5-1F1EA flag-pe
	":flag-pe:": "1f1f5-1f1ea",
	// 1F1F5-1F1EB flag-pf
	":flag-pf:": "1f1f5-1f1eb",
	// 1F1F5-1F1EC flag-pg
	":flag-pg:": "1f1f5-1f1ec",
	// 1F1F5-1F1ED flag-ph
	":flag-ph:": "1f1f5-1f1ed",
	// 1F1F5-1F1F0 flag-pk
	":flag-pk:": "1f1f5-1f1f0",
	// 1F1F5-1F1F1 flag-pl
	":flag-pl:": "1f1f5-1f1f1",
	// 1F1F5-1F1F2 flag-pm
	":flag-pm:": "1f1f5-1f1f2",
	// 1F1F5-1F1F3 flag-pn
	":flag-pn:": "1f1f5-1f1f3",
	// 1F1F5-1F1F7 flag-pr
	":flag-pr:": "1f1f5-1f1f7",
	// 1F1F5-1F1F8 flag-ps
	":flag-ps:": "1f1f5-1f1f8",
	// 1F1F5-1F1F9 flag-pt
	":flag-pt:": "1f1f5-1f1f9",
	// 1F1F5-1F1FC flag-pw
	":flag-pw:": "1f1f5-1f1fc",
	// 1F1F5-1F1FE flag-py
	":flag-py:": "1f1f5-1f1fe",
	// 1F1F6-1F1E6 flag-qa
	":flag-qa:": "1f1f6-1f1e6",
	// 1F1F7-1F1EA flag-re
	":flag-re:": "1f1f7-1f1ea",
	// 1F1F7-1F1F4 flag-ro
	":flag-ro:": "1f1f7-1f1f4",
	// 1F1F7-1F1F8 flag-rs
	":flag-rs:": "1f1f7-1f1f8",
	// 1F1F7-1F1FA ru
	":ru:": "1f1f7-1f1fa",
	// 1F1F7-1F1FA ru
	":flag-ru:": "1f1f7-1f1fa",
	// 1F1F7-1F1FC flag-rw
	":flag-rw:": "1f1f7-1f1fc",
	// 1F1F8-1F1E6 flag-sa
	":flag-sa:": "1f1f8-1f1e6",
	// 1F1F8-1F1E7 flag-sb
	":flag-sb:": "1f1f8-1f1e7",
	// 1F1F8-1F1E8 flag-sc
	":flag-sc:": "1f1f8-1f1e8",
	// 1F1F8-1F1E9 flag-sd
	":flag-sd:": "1f1f8-1f1e9",
	// 1F1F8-1F1EA flag-se
	":flag-se:": "1f1f8-1f1ea",
	// 1F1F8-1F1EC flag-sg
	":flag-sg:": "1f1f8-1f1ec",
	// 1F1F8-1F1ED flag-sh
	":flag-sh:": "1f1f8-1f1ed",
	// 1F1F8-1F1EE flag-si
	":flag-si:": "1f1f8-1f1ee",
	// 1F1F8-1F1EF flag-sj
	":flag-sj:": "1f1f8-1f1ef",
	// 1F1F8-1F1F0 flag-sk
	":flag-sk:": "1f1f8-1f1f0",
	// 1F1F8-1F1F1 flag-sl
	":flag-sl:": "1f1f8-1f1f1",
	// 1F1F8-1F1F2 flag-sm
	":flag-sm:": "1f1f8-1f1f2",
	// 1F1F8-1F1F3 flag-sn
	":flag-sn:": "1f1f8-1f1f3",
	// 1F1F8-1F1F4 flag-so
	":flag-so:": "1f1f8-1f1f4",
	// 1F1F8-1F1F7 flag-sr
	":flag-sr:": "1f1f8-1f1f7",
	// 1F1F8-1F1F8 flag-ss
	":flag-ss:": "1f1f8-1f1f8",
	// 1F1F8-1F1F9 flag-st
	":flag-st:": "1f1f8-1f1f9",
	// 1F1F8-1F1FB flag-sv
	":flag-sv:": "1f1f8-1f1fb",
	// 1F1F8-1F1FD flag-sx
	":flag-sx:": "1f1f8-1f1fd",
	// 1F1F8-1F1FE flag-sy
	":flag-sy:": "1f1f8-1f1fe",
	// 1F1F8-1F1FF flag-sz
	":flag-sz:": "1f1f8-1f1ff",
	// 1F1F9-1F1E6 flag-ta
	":flag-ta:": "1f1f9-1f1e6",
	// 1F1F9-1F1E8 flag-tc
	":flag-tc:": "1f1f9-1f1e8",
	// 1F1F9-1F1E9 flag-td
	":flag-td:": "1f1f9-1f1e9",
	// 1F1F9-1F1EB flag-tf
	":flag-tf:": "1f1f9-1f1eb",
	// 1F1F9-1F1EC flag-tg
	":flag-tg:": "1f1f9-1f1ec",
	// 1F1F9-1F1ED flag-th
	":flag-th:": "1f1f9-1f1ed",
	// 1F1F9-1F1EF flag-tj
	":flag-tj:": "1f1f9-1f1ef",
	// 1F1F9-1F1F0 flag-tk
	":flag-tk:": "1f1f9-1f1f0",
	// 1F1F9-1F1F1 flag-tl
	":flag-tl:": "1f1f9-1f1f1",
	// 1F1F9-1F1F2 flag-tm
	":flag-tm:": "1f1f9-1f1f2",
	// 1F1F9-1F1F3 flag-tn
	":flag-tn:": "1f1f9-1f1f3",
	// 1F1F9-1F1F4 flag-to
	":flag-to:": "1f1f9-1f1f4",
	// 1F1F9-1F1F7 flag-tr
	":flag-tr:": "1f1f9-1f1f7",
	// 1F1F9-1F1F9 flag-tt
	":flag-tt:": "1f1f9-1f1f9",
	// 1F1F9-1F1FB flag-tv
	":flag-tv:": "1f1f9-1f1fb",
	// 1F1F9-1F1FC flag-tw
	":flag-tw:": "1f1f9-1f1fc",
	// 1F1F9-1F1FF flag-tz
	":flag-tz:": "1f1f9-1f1ff",
	// 1F1FA-1F1E6 flag-ua
	":flag-ua:": "1f1fa-1f1e6",
	// 1F1FA-1F1EC flag-ug
	":flag-ug:": "1f1fa-1f1ec",
	// 1F1FA-1F1F2 flag-um
	":flag-um:": "1f1fa-1f1f2",
	// 1F1FA-1F1F3 flag-un
	":flag-un:": "1f1fa-1f1f3",
	// 1F1FA-1F1F8 us
	":us:": "1f1fa-1f1f8",
	// 1F1FA-1F1F8 us
	":flag-us:": "1f1fa-1f1f8",
	// 1F1FA-1F1FE flag-uy
	":flag-uy:": "1f1fa-1f1fe",
	// 1F1FA-1F1FF flag-uz
	":flag-uz:": "1f1fa-1f1ff",
	// 1F1FB-1F1E6 flag-va
	":flag-va:": "1f1fb-1f1e6",
	// 1F1FB-1F1E8 flag-vc
	":flag-vc:": "1f1fb-1f1e8",
	// 1F1FB-1F1EA flag-ve
	":flag-ve:": "1f1fb-1f1ea",
	// 1F1FB-1F1EC flag-vg
	":flag-vg:": "1f1fb-1f1ec",
	// 1F1FB-1F1EE flag-vi
	":flag-vi:": "1f1fb-1f1ee",
	// 1F1FB-1F1F3 flag-vn
	":flag-vn:": "1f1fb-1f1f3",
	// 1F1FB-1F1FA flag-vu
	":flag-vu:": "1f1fb-1f1fa",
	// 1F1FC-1F1EB flag-wf
	":flag-wf:": "1f1fc-1f1eb",
	// 1F1FC-1F1F8 flag-ws
	":flag-ws:": "1f1fc-1f1f8",
	// 1F1FD-1F1F0 flag-xk
	":flag-xk:": "1f1fd-1f1f0",
	// 1F1FE-1F1EA flag-ye
	":flag-ye:": "1f1fe-1f1ea",
	// 1F1FE-1F1F9 flag-yt
	":flag-yt:": "1f1fe-1f1f9",
	// 1F1FF-1F1E6 flag-za
	":flag-za:": "1f1ff-1f1e6",
	// 1F1FF-1F1F2 flag-zm
	":flag-zm:": "1f1ff-1f1f2",
	// 1F1FF-1F1FC flag-zw
	":flag-zw:": "1f1ff-1f1fc",
	// 1F201 koko
	":koko:": "1f201",
	// 1F202-FE0F sa
	":sa:": "1f202-fe0f",
	// 1F21A u7121
	":u7121:": "1f21a",
	// 1F22F u6307
	":u6307:": "1f22f",
	// 1F232 u7981
	":u7981:": "1f232",
	// 1F233 u7a7a
	":u7a7a:": "1f233",
	// 1F234 u5408
	":u5408:": "1f234",
	// 1F235 u6e80
	":u6e80:": "1f235",
	// 1F236 u6709
	":u6709:": "1f236",
	// 1F237-FE0F u6708
	":u6708:": "1f237-fe0f",
	// 1F238 u7533
	":u7533:": "1f238",
	// 1F239 u5272
	":u5272:": "1f239",
	// 1F23A u55b6
	":u55b6:": "1f23a",
	// 1F250 ideograph_advantage
	":ideograph_advantage:": "1f250",
	// 1F251 accept
	":accept:": "1f251",
	// 1F300 cyclone
	":cyclone:": "1f300",
	// 1F301 foggy
	":foggy:": "1f301",
	// 1F302 closed_umbrella
	":closed_umbrella:": "1f302",
	// 1F303 night_with_stars
	":night_with_stars:": "1f303",
	// 1F304 sunrise_over_mountains
	":sunrise_over_mountains:": "1f304",
	// 1F305 sunrise
	":sunrise:": "1f305",
	// 1F306 city_sunset
	":city_sunset:": "1f306",
	// 1F307 city_sunrise
	":city_sunrise:": "1f307",
	// 1F308 rainbow
	":rainbow:": "1f308",
	// 1F309 bridge_at_night
	":bridge_at_night:": "1f309",
	// 1F30A ocean
	":ocean:": "1f30a",
	// 1F30B volcano
	":volcano:": "1f30b",
	// 1F30C milky_way
	":milky_way:": "1f30c",
	// 1F30D earth_africa
	":earth_africa:": "1f30d",
	// 1F30E earth_americas
	":earth_americas:": "1f30e",
	// 1F30F earth_asia
	":earth_asia:": "1f30f",
	// 1F310 globe_with_meridians
	":globe_with_meridians:": "1f310",
	// 1F311 new_moon
	":new_moon:": "1f311",
	// 1F312 waxing_crescent_moon
	":waxing_crescent_moon:": "1f312",
	// 1F313 first_quarter_moon
	":first_quarter_moon:": "1f313",
	// 1F314 moon
	":moon:": "1f314",
	// 1F314 moon
	":waxing_gibbous_moon:": "1f314",
	// 1F315 full_moon
	":full_moon:": "1f315",
	// 1F316 waning_gibbous_moon
	":waning_gibbous_moon:": "1f316",
	// 1F317 last_quarter_moon
	":last_quarter_moon:": "1f317",
	// 1F318 waning_crescent_moon
	":waning_crescent_moon:": "1f318",
	// 1F319 crescent_moon
	":crescent_moon:": "1f319",
	// 1F31A new_moon_with_face
	":new_moon_with_face:": "1f31a",
	// 1F31B first_quarter_moon_with_face
	":first_quarter_moon_with_face:": "1f31b",
	// 1F31C last_quarter_moon_with_face
	":last_quarter_moon_with_face:": "1f31c",
	// 1F31D full_moon_with_face
	":full_moon_with_face:": "1f31d",
	// 1F31E sun_with_face
	":sun_with_face:": "1f31e",
	// 1F31F star2
	":star2:": "1f31f",
	// 1F320 stars
	":stars:": "1f320",
	// 1F321-FE0F thermometer
	":thermometer:": "1f321-fe0f",
	// 1F324-FE0F mostly_sunny
	":mostly_sunny:": "1f324-fe0f",
	// 1F324-FE0F mostly_sunny
	":sun_small_cloud:": "1f324-fe0f",
	// 1F325-FE0F barely_sunny
	":barely_sunny:": "1f325-fe0f",
	// 1F325-FE0F barely_sunny
	":sun_behind_cloud:": "1f325-fe0f",
	// 1F326-FE0F partly_sunny_rain
	":partly_sunny_rain:": "1f326-fe0f",
	// 1F326-FE0F partly_sunny_rain
	":sun_behind_rain_cloud:": "1f326-fe0f",
	// 1F327-FE0F rain_cloud
	":rain_cloud:": "1f327-fe0f",
	// 1F328-FE0F snow_cloud
	":snow_cloud:": "1f328-fe0f",
	// 1F329-FE0F lightning
	":lightning:": "1f329-fe0f",
	// 1F329-FE0F lightning
	":lightning_cloud:": "1f329-fe0f",
	// 1F32A-FE0F tornado
	":tornado:": "1f32a-fe0f",
	// 1F32A-FE0F tornado
	":tornado_cloud:": "1f32a-fe0f",
	// 1F32B-FE0F fog
	":fog:": "1f32b-fe0f",
	// 1F32C-FE0F wind_blowing_face
	":wind_blowing_face:": "1f32c-fe0f",
	// 1F32D hotdog
	":hotdog:": "1f32d",
	// 1F32E taco
	":taco:": "1f32e",
	// 1F32F burrito
	":burrito:": "1f32f",
	// 1F330 chestnut
	":chestnut:": "1f330",
	// 1F331 seedling
	":seedling:": "1f331",
	// 1F332 evergreen_tree
	":evergreen_tree:": "1f332",
	// 1F333 deciduous_tree
	":deciduous_tree:": "1f333",
	// 1F334 palm_tree
	":palm_tree:": "1f334",
	// 1F335 cactus
	":cactus:": "1f335",
	// 1F336-FE0F hot_pepper
	":hot_pepper:": "1f336-fe0f",
	// 1F337 tulip
	":tulip:": "1f337",
	// 1F338 cherry_blossom
	":cherry_blossom:": "1f338",
	// 1F339 rose
	":rose:": "1f339",
	// 1F33A hibiscus
	":hibiscus:": "1f33a",
	// 1F33B sunflower
	":sunflower:": "1f33b",
	// 1F33C blossom
	":blossom:": "1f33c",
	// 1F33D corn
	":corn:": "1f33d",
	// 1F33E ear_of_rice
	":ear_of_rice:": "1f33e",
	// 1F33F herb
	":herb:": "1f33f",
	// 1F340 four_leaf_clover
	":four_leaf_clover:": "1f340",
	// 1F341 maple_leaf
	":maple_leaf:": "1f341",
	// 1F342 fallen_leaf
	":fallen_leaf:": "1f342",
	// 1F343 leaves
	":leaves:": "1f343",
	// 1F344 mushroom
	":mushroom:": "1f344",
	// 1F345 tomato
	":tomato:": "1f345",
	// 1F346 eggplant
	":eggplant:": "1f346",
	// 1F347 grapes
	":grapes:": "1f347",
	// 1F348 melon
	":melon:": "1f348",
	// 1F349 watermelon
	":watermelon:": "1f349",
	// 1F34A tangerine
	":tangerine:": "1f34a",
	// 1F34B lemon
	":lemon:": "1f34b",
	// 1F34C banana
	":banana:": "1f34c",
	// 1F34D pineapple
	":pineapple:": "1f34d",
	// 1F34E apple
	":apple:": "1f34e",
	// 1F34F green_apple
	":green_apple:": "1f34f",
	// 1F350 pear
	":pear:": "1f350",
	// 1F351 peach
	":peach:": "1f351",
	// 1F352 cherries
	":cherries:": "1f352",
	// 1F353 strawberry
	":strawberry:": "1f353",
	// 1F354 hamburger
	":hamburger:": "1f354",
	// 1F355 pizza
	":pizza:": "1f355",
	// 1F356 meat_on_bone
	":meat_on_bone:": "1f356",
	// 1F357 poultry_leg
	":poultry_leg:": "1f357",
	// 1F358 rice_cracker
	":rice_cracker:": "1f358",
	// 1F359 rice_ball
	":rice_ball:": "1f359",
	// 1F35A rice
	":rice:": "1f35a",
	// 1F35B curry
	":curry:": "1f35b",
	// 1F35C ramen
	":ramen:": "1f35c",
	// 1F35D spaghetti
	":spaghetti:": "1f35d",
	// 1F35E bread
	":bread:": "1f35e",
	// 1F35F fries
	":fries:": "1f35f",
	// 1F360 sweet_potato
	":sweet_potato:": "1f360",
	// 1F361 dango
	":dango:": "1f361",
	// 1F362 oden
	":oden:": "1f362",
	// 1F363 sushi
	":sushi:": "1f363",
	// 1F364 fried_shrimp
	":fried_shrimp:": "1f364",
	// 1F365 fish_cake
	":fish_cake:": "1f365",
	// 1F366 icecream
	":icecream:": "1f366",
	// 1F367 shaved_ice
	":shaved_ice:": "1f367",
	// 1F368 ice_cream
	":ice_cream:": "1f368",
	// 1F369 doughnut
	":doughnut:": "1f369",
	// 1F36A cookie
	":cookie:": "1f36a",
	// 1F36B chocolate_bar
	":chocolate_bar:": "1f36b",
	// 1F36C candy
	":candy:": "1f36c",
	// 1F36D lollipop
	":lollipop:": "1f36d",
	// 1F36E custard
	":custard:": "1f36e",
	// 1F36F honey_pot
	":honey_pot:": "1f36f",
	// 1F370 cake
	":cake:": "1f370",
	// 1F371 bento
	":bento:": "1f371",
	// 1F372 stew
	":stew:": "1f372",
	// 1F373 fried_egg
	":fried_egg:": "1f373",
	// 1F373 fried_egg
	":cooking:": "1f373",
	// 1F374 fork_and_knife
	":fork_and_knife:": "1f374",
	// 1F375 tea
	":tea:": "1f375",
	// 1F376 sake
	":sake:": "1f376",
	// 1F377 wine_glass
	":wine_glass:": "1f377",
	// 1F378 cocktail
	":cocktail:": "1f378",
	// 1F379 tropical_drink
	":tropical_drink:": "1f379",
	// 1F37A beer
	":beer:": "1f37a",
	// 1F37B beers
	":beers:": "1f37b",
	// 1F37C baby_bottle
	":baby_bottle:": "1f37c",
	// 1F37D-FE0F knife_fork_plate
	":knife_fork_plate:": "1f37d-fe0f",
	// 1F37E champagne
	":champagne:": "1f37e",
	// 1F37F popcorn
	":popcorn:": "1f37f",
	// 1F380 ribbon
	":ribbon:": "1f380",
	// 1F381 gift
	":gift:": "1f381",
	// 1F382 birthday
	":birthday:": "1f382",
	// 1F383 jack_o_lantern
	":jack_o_lantern:": "1f383",
	// 1F384 christmas_tree
	":christmas_tree:": "1f384",
	// 1F385 santa
	":santa:":              "1f385",
	":santa::skin-tone-2:": "1f385-1f3fb",
	":santa::skin-tone-3:": "1f385-1f3fc",
	":santa::skin-tone-4:": "1f385-1f3fd",
	":santa::skin-tone-5:": "1f385-1f3fe",
	":santa::skin-tone-6:": "1f385-1f3ff",
	// 1F386 fireworks
	":fireworks:": "1f386",
	// 1F387 sparkler
	":sparkler:": "1f387",
	// 1F388 balloon
	":balloon:": "1f388",
	// 1F389 tada
	":tada:": "1f389",
	// 1F38A confetti_ball
	":confetti_ball:": "1f38a",
	// 1F38B tanabata_tree
	":tanabata_tree:": "1f38b",
	// 1F38C crossed_flags
	":crossed_flags:": "1f38c",
	// 1F38D bamboo
	":bamboo:": "1f38d",
	// 1F38E dolls
	":dolls:": "1f38e",
	// 1F38F flags
	":flags:": "1f38f",
	// 1F390 wind_chime
	":wind_chime:": "1f390",
	// 1F391 rice_scene
	":rice_scene:": "1f391",
	// 1F392 school_satchel
	":school_satchel:": "1f392",
	// 1F393 mortar_board
	":mortar_board:": "1f393",
	// 1F396-FE0F medal
	":medal:": "1f396-fe0f",
	// 1F397-FE0F reminder_ribbon
	":reminder_ribbon:": "1f397-fe0f",
	// 1F399-FE0F studio_microphone
	":studio_microphone:": "1f399-fe0f",
	// 1F39A-FE0F level_slider
	":level_slider:": "1f39a-fe0f",
	// 1F39B-FE0F control_knobs
	":control_knobs:": "1f39b-fe0f",
	// 1F39E-FE0F film_frames
	":film_frames:": "1f39e-fe0f",
	// 1F39F-FE0F admission_tickets
	":admission_tickets:": "1f39f-fe0f",
	// 1F3A0 carousel_horse
	":carousel_horse:": "1f3a0",
	// 1F3A1 ferris_wheel
	":ferris_wheel:": "1f3a1",
	// 1F3A2 roller_coaster
	":roller_coaster:": "1f3a2",
	// 1F3A3 fishing_pole_and_fish
	":fishing_pole_and_fish:": "1f3a3",
	// 1F3A4 microphone
	":microphone:": "1f3a4",
	// 1F3A5 movie_camera
	":movie_camera:": "1f3a5",
	// 1F3A6 cinema
	":cinema:": "1f3a6",
	// 1F3A7 headphones
	":headphones:": "1f3a7",
	// 1F3A8 art
	":art:": "1f3a8",
	// 1F3A9 tophat
	":tophat:": "1f3a9",
	// 1F3AA circus_tent
	":circus_tent:": "1f3aa",
	// 1F3AB ticket
	":ticket:": "1f3ab",
	// 1F3AC clapper
	":clapper:": "1f3ac",
	// 1F3AD performing_arts
	":performing_arts:": "1f3ad",
	// 1F3AE video_game
	":video_game:": "1f3ae",
	// 1F3AF dart
	":dart:": "1f3af",
	// 1F3B0 slot_machine
	":slot_machine:": "1f3b0",
	// 1F3B1 8ball
	":8ball:": "1f3b1",
	// 1F3B2 game_die
	":game_die:": "1f3b2",
	// 1F3B3 bowling
	":bowling:": "1f3b3",
	// 1F3B4 flower_playing_cards
	":flower_playing_cards:": "1f3b4",
	// 1F3B5 musical_note
	":musical_note:": "1f3b5",
	// 1F3B6 notes
	":notes:": "1f3b6",
	// 1F3B7 saxophone
	":saxophone:": "1f3b7",
	// 1F3B8 guitar
	":guitar:": "1f3b8",
	// 1F3B9 musical_keyboard
	":musical_keyboard:": "1f3b9",
	// 1F3BA trumpet
	":trumpet:": "1f3ba",
	// 1F3BB violin
	":violin:": "1f3bb",
	// 1F3BC musical_score
	":musical_score:": "1f3bc",
	// 1F3BD running_shirt_with_sash
	":running_shirt_with_sash:": "1f3bd",
	// 1F3BE tennis
	":tennis:": "1f3be",
	// 1F3BF ski
	":ski:": "1f3bf",
	// 1F3C0 basketball
	":basketball:": "1f3c0",
	// 1F3C1 checkered_flag
	":checkered_flag:": "1f3c1",
	// 1F3C2 snowboarder
	":snowboarder:":              "1f3c2",
	":snowboarder::skin-tone-2:": "1f3c2-1f3fb",
	":snowboarder::skin-tone-3:": "1f3c2-1f3fc",
	":snowboarder::skin-tone-4:": "1f3c2-1f3fd",
	":snowboarder::skin-tone-5:": "1f3c2-1f3fe",
	":snowboarder::skin-tone-6:": "1f3c2-1f3ff",
	// 1F3C3-200D-2640-FE0F woman-running
	":woman-running:":              "1f3c3-200d-2640-fe0f",
	":woman-running::skin-tone-2:": "1f3c3-1f3fb-200d-2640-fe0f",
	":woman-running::skin-tone-3:": "1f3c3-1f3fc-200d-2640-fe0f",
	":woman-running::skin-tone-4:": "1f3c3-1f3fd-200d-2640-fe0f",
	":woman-running::skin-tone-5:": "1f3c3-1f3fe-200d-2640-fe0f",
	":woman-running::skin-tone-6:": "1f3c3-1f3ff-200d-2640-fe0f",
	// 1F3C3-200D-2642-FE0F man-running
	":man-running:":              "1f3c3-200d-2642-fe0f",
	":man-running::skin-tone-2:": "1f3c3-1f3fb-200d-2642-fe0f",
	":man-running::skin-tone-3:": "1f3c3-1f3fc-200d-2642-fe0f",
	":man-running::skin-tone-4:": "1f3c3-1f3fd-200d-2642-fe0f",
	":man-running::skin-tone-5:": "1f3c3-1f3fe-200d-2642-fe0f",
	":man-running::skin-tone-6:": "1f3c3-1f3ff-200d-2642-fe0f",
	// 1F3C3 runner
	":runner:": "1f3c3",
	// 1F3C3 runner
	":running:":              "1f3c3",
	":runner::skin-tone-2:":  "1f3c3-1f3fb",
	":running::skin-tone-2:": "1f3c3-1f3fb",
	":runner::skin-tone-3:":  "1f3c3-1f3fc",
	":running::skin-tone-3:": "1f3c3-1f3fc",
	":runner::skin-tone-4:":  "1f3c3-1f3fd",
	":running::skin-tone-4:": "1f3c3-1f3fd",
	":runner::skin-tone-5:":  "1f3c3-1f3fe",
	":running::skin-tone-5:": "1f3c3-1f3fe",
	":runner::skin-tone-6:":  "1f3c3-1f3ff",
	":running::skin-tone-6:": "1f3c3-1f3ff",
	// 1F3C4-200D-2640-FE0F woman-surfing
	":woman-surfing:":              "1f3c4-200d-2640-fe0f",
	":woman-surfing::skin-tone-2:": "1f3c4-1f3fb-200d-2640-fe0f",
	":woman-surfing::skin-tone-3:": "1f3c4-1f3fc-200d-2640-fe0f",
	":woman-surfing::skin-tone-4:": "1f3c4-1f3fd-200d-2640-fe0f",
	":woman-surfing::skin-tone-5:": "1f3c4-1f3fe-200d-2640-fe0f",
	":woman-surfing::skin-tone-6:": "1f3c4-1f3ff-200d-2640-fe0f",
	// 1F3C4-200D-2642-FE0F man-surfing
	":man-surfing:":              "1f3c4-200d-2642-fe0f",
	":man-surfing::skin-tone-2:": "1f3c4-1f3fb-200d-2642-fe0f",
	":man-surfing::skin-tone-3:": "1f3c4-1f3fc-200d-2642-fe0f",
	":man-surfing::skin-tone-4:": "1f3c4-1f3fd-200d-2642-fe0f",
	":man-surfing::skin-tone-5:": "1f3c4-1f3fe-200d-2642-fe0f",
	":man-surfing::skin-tone-6:": "1f3c4-1f3ff-200d-2642-fe0f",
	// 1F3C4 surfer
	":surfer:":              "1f3c4",
	":surfer::skin-tone-2:": "1f3c4-1f3fb",
	":surfer::skin-tone-3:": "1f3c4-1f3fc",
	":surfer::skin-tone-4:": "1f3c4-1f3fd",
	":surfer::skin-tone-5:": "1f3c4-1f3fe",
	":surfer::skin-tone-6:": "1f3c4-1f3ff",
	// 1F3C5 sports_medal
	":sports_medal:": "1f3c5",
	// 1F3C6 trophy
	":trophy:": "1f3c6",
	// 1F3C7 horse_racing
	":horse_racing:":              "1f3c7",
	":horse_racing::skin-tone-2:": "1f3c7-1f3fb",
	":horse_racing::skin-tone-3:": "1f3c7-1f3fc",
	":horse_racing::skin-tone-4:": "1f3c7-1f3fd",
	":horse_racing::skin-tone-5:": "1f3c7-1f3fe",
	":horse_racing::skin-tone-6:": "1f3c7-1f3ff",
	// 1F3C8 football
	":football:": "1f3c8",
	// 1F3C9 rugby_football
	":rugby_football:": "1f3c9",
	// 1F3CA-200D-2640-FE0F woman-swimming
	":woman-swimming:":              "1f3ca-200d-2640-fe0f",
	":woman-swimming::skin-tone-2:": "1f3ca-1f3fb-200d-2640-fe0f",
	":woman-swimming::skin-tone-3:": "1f3ca-1f3fc-200d-2640-fe0f",
	":woman-swimming::skin-tone-4:": "1f3ca-1f3fd-200d-2640-fe0f",
	":woman-swimming::skin-tone-5:": "1f3ca-1f3fe-200d-2640-fe0f",
	":woman-swimming::skin-tone-6:": "1f3ca-1f3ff-200d-2640-fe0f",
	// 1F3CA-200D-2642-FE0F man-swimming
	":man-swimming:":              "1f3ca-200d-2642-fe0f",
	":man-swimming::skin-tone-2:": "1f3ca-1f3fb-200d-2642-fe0f",
	":man-swimming::skin-tone-3:": "1f3ca-1f3fc-200d-2642-fe0f",
	":man-swimming::skin-tone-4:": "1f3ca-1f3fd-200d-2642-fe0f",
	":man-swimming::skin-tone-5:": "1f3ca-1f3fe-200d-2642-fe0f",
	":man-swimming::skin-tone-6:": "1f3ca-1f3ff-200d-2642-fe0f",
	// 1F3CA swimmer
	":swimmer:":              "1f3ca",
	":swimmer::skin-tone-2:": "1f3ca-1f3fb",
	":swimmer::skin-tone-3:": "1f3ca-1f3fc",
	":swimmer::skin-tone-4:": "1f3ca-1f3fd",
	":swimmer::skin-tone-5:": "1f3ca-1f3fe",
	":swimmer::skin-tone-6:": "1f3ca-1f3ff",
	// 1F3CB-FE0F-200D-2640-FE0F woman-lifting-weights
	":woman-lifting-weights:":              "1f3cb-fe0f-200d-2640-fe0f",
	":woman-lifting-weights::skin-tone-2:": "1f3cb-1f3fb-200d-2640-fe0f",
	":woman-lifting-weights::skin-tone-3:": "1f3cb-1f3fc-200d-2640-fe0f",
	":woman-lifting-weights::skin-tone-4:": "1f3cb-1f3fd-200d-2640-fe0f",
	":woman-lifting-weights::skin-tone-5:": "1f3cb-1f3fe-200d-2640-fe0f",
	":woman-lifting-weights::skin-tone-6:": "1f3cb-1f3ff-200d-2640-fe0f",
	// 1F3CB-FE0F-200D-2642-FE0F man-lifting-weights
	":man-lifting-weights:":              "1f3cb-fe0f-200d-2642-fe0f",
	":man-lifting-weights::skin-tone-2:": "1f3cb-1f3fb-200d-2642-fe0f",
	":man-lifting-weights::skin-tone-3:": "1f3cb-1f3fc-200d-2642-fe0f",
	":man-lifting-weights::skin-tone-4:": "1f3cb-1f3fd-200d-2642-fe0f",
	":man-lifting-weights::skin-tone-5:": "1f3cb-1f3fe-200d-2642-fe0f",
	":man-lifting-weights::skin-tone-6:": "1f3cb-1f3ff-200d-2642-fe0f",
	// 1F3CB-FE0F weight_lifter
	":weight_lifter:":              "1f3cb-fe0f",
	":weight_lifter::skin-tone-2:": "1f3cb-1f3fb",
	":weight_lifter::skin-tone-3:": "1f3cb-1f3fc",
	":weight_lifter::skin-tone-4:": "1f3cb-1f3fd",
	":weight_lifter::skin-tone-5:": "1f3cb-1f3fe",
	":weight_lifter::skin-tone-6:": "1f3cb-1f3ff",
	// 1F3CC-FE0F-200D-2640-FE0F woman-golfing
	":woman-golfing:":              "1f3cc-fe0f-200d-2640-fe0f",
	":woman-golfing::skin-tone-2:": "1f3cc-1f3fb-200d-2640-fe0f",
	":woman-golfing::skin-tone-3:": "1f3cc-1f3fc-200d-2640-fe0f",
	":woman-golfing::skin-tone-4:": "1f3cc-1f3fd-200d-2640-fe0f",
	":woman-golfing::skin-tone-5:": "1f3cc-1f3fe-200d-2640-fe0f",
	":woman-golfing::skin-tone-6:": "1f3cc-1f3ff-200d-2640-fe0f",
	// 1F3CC-FE0F-200D-2642-FE0F man-golfing
	":man-golfing:":              "1f3cc-fe0f-200d-2642-fe0f",
	":man-golfing::skin-tone-2:": "1f3cc-1f3fb-200d-2642-fe0f",
	":man-golfing::skin-tone-3:": "1f3cc-1f3fc-200d-2642-fe0f",
	":man-golfing::skin-tone-4:": "1f3cc-1f3fd-200d-2642-fe0f",
	":man-golfing::skin-tone-5:": "1f3cc-1f3fe-200d-2642-fe0f",
	":man-golfing::skin-tone-6:": "1f3cc-1f3ff-200d-2642-fe0f",
	// 1F3CC-FE0F golfer
	":golfer:":              "1f3cc-fe0f",
	":golfer::skin-tone-2:": "1f3cc-1f3fb",
	":golfer::skin-tone-3:": "1f3cc-1f3fc",
	":golfer::skin-tone-4:": "1f3cc-1f3fd",
	":golfer::skin-tone-5:": "1f3cc-1f3fe",
	":golfer::skin-tone-6:": "1f3cc-1f3ff",
	// 1F3CD-FE0F racing_motorcycle
	":racing_motorcycle:": "1f3cd-fe0f",
	// 1F3CE-FE0F racing_car
	":racing_car:": "1f3ce-fe0f",
	// 1F3CF cricket_bat_and_ball
	":cricket_bat_and_ball:": "1f3cf",
	// 1F3D0 volleyball
	":volleyball:": "1f3d0",
	// 1F3D1 field_hockey_stick_and_ball
	":field_hockey_stick_and_ball:": "1f3d1",
	// 1F3D2 ice_hockey_stick_and_puck
	":ice_hockey_stick_and_puck:": "1f3d2",
	// 1F3D3 table_tennis_paddle_and_ball
	":table_tennis_paddle_and_ball:": "1f3d3",
	// 1F3D4-FE0F snow_capped_mountain
	":snow_capped_mountain:": "1f3d4-fe0f",
	// 1F3D5-FE0F camping
	":camping:": "1f3d5-fe0f",
	// 1F3D6-FE0F beach_with_umbrella
	":beach_with_umbrella:": "1f3d6-fe0f",
	// 1F3D7-FE0F building_construction
	":building_construction:": "1f3d7-fe0f",
	// 1F3D8-FE0F house_buildings
	":house_buildings:": "1f3d8-fe0f",
	// 1F3D9-FE0F cityscape
	":cityscape:": "1f3d9-fe0f",
	// 1F3DA-FE0F derelict_house_building
	":derelict_house_building:": "1f3da-fe0f",
	// 1F3DB-FE0F classical_building
	":classical_building:": "1f3db-fe0f",
	// 1F3DC-FE0F desert
	":desert:": "1f3dc-fe0f",
	// 1F3DD-FE0F desert_island
	":desert_island:": "1f3dd-fe0f",
	// 1F3DE-FE0F national_park
	":national_park:": "1f3de-fe0f",
	// 1F3DF-FE0F stadium
	":stadium:": "1f3df-fe0f",
	// 1F3E0 house
	":house:": "1f3e0",
	// 1F3E1 house_with_garden
	":house_with_garden:": "1f3e1",
	// 1F3E2 office
	":office:": "1f3e2",
	// 1F3E3 post_office
	":post_office:": "1f3e3",
	// 1F3E4 european_post_office
	":european_post_office:": "1f3e4",
	// 1F3E5 hospital
	":hospital:": "1f3e5",
	// 1F3E6 bank
	":bank:": "1f3e6",
	// 1F3E7 atm
	":atm:": "1f3e7",
	// 1F3E8 hotel
	":hotel:": "1f3e8",
	// 1F3E9 love_hotel
	":love_hotel:": "1f3e9",
	// 1F3EA convenience_store
	":convenience_store:": "1f3ea",
	// 1F3EB school
	":school:": "1f3eb",
	// 1F3EC department_store
	":department_store:": "1f3ec",
	// 1F3ED factory
	":factory:": "1f3ed",
	// 1F3EE izakaya_lantern
	":izakaya_lantern:": "1f3ee",
	// 1F3EE izakaya_lantern
	":lantern:": "1f3ee",
	// 1F3EF japanese_castle
	":japanese_castle:": "1f3ef",
	// 1F3F0 european_castle
	":european_castle:": "1f3f0",
	// 1F3F3-FE0F-200D-1F308 rainbow-flag
	":rainbow-flag:": "1f3f3-fe0f-200d-1f308",
	// 1F3F3-FE0F-200D-26A7-FE0F transgender_flag
	":transgender_flag:": "1f3f3-fe0f-200d-26a7-fe0f",
	// 1F3F3-FE0F waving_white_flag
	":waving_white_flag:": "1f3f3-fe0f",
	// 1F3F4-200D-2620-FE0F pirate_flag
	":pirate_flag:": "1f3f4-200d-2620-fe0f",
	// 1F3F4-E0067-E0062-E0065-E006E-E0067-E007F flag-england
	":flag-england:": "1f3f4-e0067-e0062-e0065-e006e-e0067-e007f",
	// 1F3F4-E0067-E0062-E0073-E0063-E0074-E007F flag-scotland
	":flag-scotland:": "1f3f4-e0067-e0062-e0073-e0063-e0074-e007f",
	// 1F3F4-E0067-E0062-E0077-E006C-E0073-E007F flag-wales
	":flag-wales:": "1f3f4-e0067-e0062-e0077-e006c-e0073-e007f",
	// 1F3F4 waving_black_flag
	":waving_black_flag:": "1f3f4",
	// 1F3F5-FE0F rosette
	":rosette:": "1f3f5-fe0f",
	// 1F3F7-FE0F label
	":label:": "1f3f7-fe0f",
	// 1F3F8 badminton_racquet_and_shuttlecock
	":badminton_racquet_and_shuttlecock:": "1f3f8",
	// 1F3F9 bow_and_arrow
	":bow_and_arrow:": "1f3f9",
	// 1F3FA amphora
	":amphora:": "1f3fa",
	// 1F3FB skin-tone-2
	":skin-tone-2:": "1f3fb",
	// 1F3FC skin-tone-3
	":skin-tone-3:": "1f3fc",
	// 1F3FD skin-tone-4
	":skin-tone-4:": "1f3fd",
	// 1F3FE skin-tone-5
	":skin-tone-5:": "1f3fe",
	// 1F3FF skin-tone-6
	":skin-tone-6:": "1f3ff",
	// 1F400 rat
	":rat:": "1f400",
	// 1F401 mouse2
	":mouse2:": "1f401",
	// 1F402 ox
	":ox:": "1f402",
	// 1F403 water_buffalo
	":water_buffalo:": "1f403",
	// 1F404 cow2
	":cow2:": "1f404",
	// 1F405 tiger2
	":tiger2:": "1f405",
	// 1F406 leopard
	":leopard:": "1f406",
	// 1F407 rabbit2
	":rabbit2:": "1f407",
	// 1F408-200D-2B1B black_cat
	":black_cat:": "1f408-200d-2b1b",
	// 1F408 cat2
	":cat2:": "1f408",
	// 1F409 dragon
	":dragon:": "1f409",
	// 1F40A crocodile
	":crocodile:": "1f40a",
	// 1F40B whale2
	":whale2:": "1f40b",
	// 1F40C snail
	":snail:": "1f40c",
	// 1F40D snake
	":snake:": "1f40d",
	// 1F40E racehorse
	":racehorse:": "1f40e",
	// 1F40F ram
	":ram:": "1f40f",
	// 1F410 goat
	":goat:": "1f410",
	// 1F411 sheep
	":sheep:": "1f411",
	// 1F412 monkey
	":monkey:": "1f412",
	// 1F413 rooster
	":rooster:": "1f413",
	// 1F414 chicken
	":chicken:": "1f414",
	// 1F415-200D-1F9BA service_dog
	":service_dog:": "1f415-200d-1f9ba",
	// 1F415 dog2
	":dog2:": "1f415",
	// 1F416 pig2
	":pig2:": "1f416",
	// 1F417 boar
	":boar:": "1f417",
	// 1F418 elephant
	":elephant:": "1f418",
	// 1F419 octopus
	":octopus:": "1f419",
	// 1F41A shell
	":shell:": "1f41a",
	// 1F41B bug
	":bug:": "1f41b",
	// 1F41C ant
	":ant:": "1f41c",
	// 1F41D bee
	":bee:": "1f41d",
	// 1F41D bee
	":honeybee:": "1f41d",
	// 1F41E ladybug
	":ladybug:": "1f41e",
	// 1F41E ladybug
	":lady_beetle:": "1f41e",
	// 1F41F fish
	":fish:": "1f41f",
	// 1F420 tropical_fish
	":tropical_fish:": "1f420",
	// 1F421 blowfish
	":blowfish:": "1f421",
	// 1F422 turtle
	":turtle:": "1f422",
	// 1F423 hatching_chick
	":hatching_chick:": "1f423",
	// 1F424 baby_chick
	":baby_chick:": "1f424",
	// 1F425 hatched_chick
	":hatched_chick:": "1f425",
	// 1F426 bird
	":bird:": "1f426",
	// 1F427 penguin
	":penguin:": "1f427",
	// 1F428 koala
	":koala:": "1f428",
	// 1F429 poodle
	":poodle:": "1f429",
	// 1F42A dromedary_camel
	":dromedary_camel:": "1f42a",
	// 1F42B camel
	":camel:": "1f42b",
	// 1F42C dolphin
	":dolphin:": "1f42c",
	// 1F42C dolphin
	":flipper:": "1f42c",
	// 1F42D mouse
	":mouse:": "1f42d",
	// 1F42E cow
	":cow:": "1f42e",
	// 1F42F tiger
	":tiger:": "1f42f",
	// 1F430 rabbit
	":rabbit:": "1f430",
	// 1F431 cat
	":cat:": "1f431",
	// 1F432 dragon_face
	":dragon_face:": "1f432",
	// 1F433 whale
	":whale:": "1f433",
	// 1F434 horse
	":horse:": "1f434",
	// 1F435 monkey_face
	":monkey_face:": "1f435",
	// 1F436 dog
	":dog:": "1f436",
	// 1F437 pig
	":pig:": "1f437",
	// 1F438 frog
	":frog:": "1f438",
	// 1F439 hamster
	":hamster:": "1f439",
	// 1F43A wolf
	":wolf:": "1f43a",
	// 1F43B-200D-2744-FE0F polar_bear
	":polar_bear:": "1f43b-200d-2744-fe0f",
	// 1F43B bear
	":bear:": "1f43b",
	// 1F43C panda_face
	":panda_face:": "1f43c",
	// 1F43D pig_nose
	":pig_nose:": "1f43d",
	// 1F43E feet
	":feet:": "1f43e",
	// 1F43E feet
	":paw_prints:": "1f43e",
	// 1F43F-FE0F chipmunk
	":chipmunk:": "1f43f-fe0f",
	// 1F440 eyes
	":eyes:": "1f440",
	// 1F441-FE0F-200D-1F5E8-FE0F eye-in-speech-bubble
	":eye-in-speech-bubble:": "1f441-fe0f-200d-1f5e8-fe0f",
	// 1F441-FE0F eye
	":eye:": "1f441-fe0f",
	// 1F442 ear
	":ear:":              "1f442",
	":ear::skin-tone-2:": "1f442-1f3fb",
	":ear::skin-tone-3:": "1f442-1f3fc",
	":ear::skin-tone-4:": "1f442-1f3fd",
	":ear::skin-tone-5:": "1f442-1f3fe",
	":ear::skin-tone-6:": "1f442-1f3ff",
	// 1F443 nose
	":nose:":              "1f443",
	":nose::skin-tone-2:": "1f443-1f3fb",
	":nose::skin-tone-3:": "1f443-1f3fc",
	":nose::skin-tone-4:": "1f443-1f3fd",
	":nose::skin-tone-5:": "1f443-1f3fe",
	":nose::skin-tone-6:": "1f443-1f3ff",
	// 1F444 lips
	":lips:": "1f444",
	// 1F445 tongue
	":tongue:": "1f445",
	// 1F446 point_up_2
	":point_up_2:":              "1f446",
	":point_up_2::skin-tone-2:": "1f446-1f3fb",
	":point_up_2::skin-tone-3:": "1f446-1f3fc",
	":point_up_2::skin-tone-4:": "1f446-1f3fd",
	":point_up_2::skin-tone-5:": "1f446-1f3fe",
	":point_up_2::skin-tone-6:": "1f446-1f3ff",
	// 1F447 point_down
	":point_down:":              "1f447",
	":point_down::skin-tone-2:": "1f447-1f3fb",
	":point_down::skin-tone-3:": "1f447-1f3fc",
	":point_down::skin-tone-4:": "1f447-1f3fd",
	":point_down::skin-tone-5:": "1f447-1f3fe",
	":point_down::skin-tone-6:": "1f447-1f3ff",
	// 1F448 point_left
	":point_left:":              "1f448",
	":point_left::skin-tone-2:": "1f448-1f3fb",
	":point_left::skin-tone-3:": "1f448-1f3fc",
	":point_left::skin-tone-4:": "1f448-1f3fd",
	":point_left::skin-tone-5:": "1f448-1f3fe",
	":point_left::skin-tone-6:": "1f448-1f3ff",
	// 1F449 point_right
	":point_right:":              "1f449",
	":point_right::skin-tone-2:": "1f449-1f3fb",
	":point_right::skin-tone-3:": "1f449-1f3fc",
	":point_right::skin-tone-4:": "1f449-1f3fd",
	":point_right::skin-tone-5:": "1f449-1f3fe",
	":point_right::skin-tone-6:": "1f449-1f3ff",
	// 1F44A facepunch
	":facepunch:": "1f44a",
	// 1F44A facepunch
	":punch:":                  "1f44a",
	":facepunch::skin-tone-2:": "1f44a-1f3fb",
	":punch::skin-tone-2:":     "1f44a-1f3fb",
	":facepunch::skin-tone-3:": "1f44a-1f3fc",
	":punch::skin-tone-3:":     "1f44a-1f3fc",
	":facepunch::skin-tone-4:": "1f44a-1f3fd",
	":punch::skin-tone-4:":     "1f44a-1f3fd",
	":facepunch::skin-tone-5:": "1f44a-1f3fe",
	":punch::skin-tone-5:":     "1f44a-1f3fe",
	":facepunch::skin-tone-6:": "1f44a-1f3ff",
	":punch::skin-tone-6:":     "1f44a-1f3ff",
	// 1F44B wave
	":wave:":              "1f44b",
	":wave::skin-tone-2:": "1f44b-1f3fb",
	":wave::skin-tone-3:": "1f44b-1f3fc",
	":wave::skin-tone-4:": "1f44b-1f3fd",
	":wave::skin-tone-5:": "1f44b-1f3fe",
	":wave::skin-tone-6:": "1f44b-1f3ff",
	// 1F44C ok_hand
	":ok_hand:":              "1f44c",
	":ok_hand::skin-tone-2:": "1f44c-1f3fb",
	":ok_hand::skin-tone-3:": "1f44c-1f3fc",
	":ok_hand::skin-tone-4:": "1f44c-1f3fd",
	":ok_hand::skin-tone-5:": "1f44c-1f3fe",
	":ok_hand::skin-tone-6:": "1f44c-1f3ff",
	// 1F44D +1
	":+1:": "1f44d",
	// 1F44D +1
	":thumbsup:":              "1f44d",
	":+1::skin-tone-2:":       "1f44d-1f3fb",
	":thumbsup::skin-tone-2:": "1f44d-1f3fb",
	":+1::skin-tone-3:":       "1f44d-1f3fc",
	":thumbsup::skin-tone-3:": "1f44d-1f3fc",
	":+1::skin-tone-4:":       "1f44d-1f3fd",
	":thumbsup::skin-tone-4:": "1f44d-1f3fd",
	":+1::skin-tone-5:":       "1f44d-1f3fe",
	":thumbsup::skin-tone-5:": "1f44d-1f3fe",
	":+1::skin-tone-6:":       "1f44d-1f3ff",
	":thumbsup::skin-tone-6:": "1f44d-1f3ff",
	// 1F44E -1
	":-1:": "1f44e",
	// 1F44E -1
	":thumbsdown:":              "1f44e",
	":-1::skin-tone-2:":         "1f44e-1f3fb",
	":thumbsdown::skin-tone-2:": "1f44e-1f3fb",
	":-1::skin-tone-3:":         "1f44e-1f3fc",
	":thumbsdown::skin-tone-3:": "1f44e-1f3fc",
	":-1::skin-tone-4:":         "1f44e-1f3fd",
	":thumbsdown::skin-tone-4:": "1f44e-1f3fd",
	":-1::skin-tone-5:":         "1f44e-1f3fe",
	":thumbsdown::skin-tone-5:": "1f44e-1f3fe",
	":-1::skin-tone-6:":         "1f44e-1f3ff",
	":thumbsdown::skin-tone-6:": "1f44e-1f3ff",
	// 1F44F clap
	":clap:":              "1f44f",
	":clap::skin-tone-2:": "1f44f-1f3fb",
	":clap::skin-tone-3:": "1f44f-1f3fc",
	":clap::skin-tone-4:": "1f44f-1f3fd",
	":clap::skin-tone-5:": "1f44f-1f3fe",
	":clap::skin-tone-6:": "1f44f-1f3ff",
	// 1F450 open_hands
	":open_hands:":              "1f450",
	":open_hands::skin-tone-2:": "1f450-1f3fb",
	":open_hands::skin-tone-3:": "1f450-1f3fc",
	":open_hands::skin-tone-4:": "1f450-1f3fd",
	":open_hands::skin-tone-5:": "1f450-1f3fe",
	":open_hands::skin-tone-6:": "1f450-1f3ff",
	// 1F451 crown
	":crown:": "1f451",
	// 1F452 womans_hat
	":womans_hat:": "1f452",
	// 1F453 eyeglasses
	":eyeglasses:": "1f453",
	// 1F454 necktie
	":necktie:": "1f454",
	// 1F455 shirt
	":shirt:": "1f455",
	// 1F455 shirt
	":tshirt:": "1f455",
	// 1F456 jeans
	":jeans:": "1f456",
	// 1F457 dress
	":dress:": "1f457",
	// 1F458 kimono
	":kimono:": "1f458",
	// 1F459 bikini
	":bikini:": "1f459",
	// 1F45A womans_clothes
	":womans_clothes:": "1f45a",
	// 1F45B purse
	":purse:": "1f45b",
	// 1F45C handbag
	":handbag:": "1f45c",
	// 1F45D pouch
	":pouch:": "1f45d",
	// 1F45E mans_shoe
	":mans_shoe:": "1f45e",
	// 1F45E mans_shoe
	":shoe:": "1f45e",
	// 1F45F athletic_shoe
	":athletic_shoe:": "1f45f",
	// 1F460 high_heel
	":high_heel:": "1f460",
	// 1F461 sandal
	":sandal:": "1f461",
	// 1F462 boot
	":boot:": "1f462",
	// 1F463 footprints
	":footprints:": "1f463",
	// 1F464 bust_in_silhouette
	":bust_in_silhouette:": "1f464",
	// 1F465 busts_in_silhouette
	":busts_in_silhouette:": "1f465",
	// 1F466 boy
	":boy:":              "1f466",
	":boy::skin-tone-2:": "1f466-1f3fb",
	":boy::skin-tone-3:": "1f466-1f3fc",
	":boy::skin-tone-4:": "1f466-1f3fd",
	":boy::skin-tone-5:": "1f466-1f3fe",
	":boy::skin-tone-6:": "1f466-1f3ff",
	// 1F467 girl
	":girl:":              "1f467",
	":girl::skin-tone-2:": "1f467-1f3fb",
	":girl::skin-tone-3:": "1f467-1f3fc",
	":girl::skin-tone-4:": "1f467-1f3fd",
	":girl::skin-tone-5:": "1f467-1f3fe",
	":girl::skin-tone-6:": "1f467-1f3ff",
	// 1F468-200D-1F33E male-farmer
	":male-farmer:":              "1f468-200d-1f33e",
	":male-farmer::skin-tone-2:": "1f468-1f3fb-200d-1f33e",
	":male-farmer::skin-tone-3:": "1f468-1f3fc-200d-1f33e",
	":male-farmer::skin-tone-4:": "1f468-1f3fd-200d-1f33e",
	":male-farmer::skin-tone-5:": "1f468-1f3fe-200d-1f33e",
	":male-farmer::skin-tone-6:": "1f468-1f3ff-200d-1f33e",
	// 1F468-200D-1F373 male-cook
	":male-cook:":              "1f468-200d-1f373",
	":male-cook::skin-tone-2:": "1f468-1f3fb-200d-1f373",
	":male-cook::skin-tone-3:": "1f468-1f3fc-200d-1f373",
	":male-cook::skin-tone-4:": "1f468-1f3fd-200d-1f373",
	":male-cook::skin-tone-5:": "1f468-1f3fe-200d-1f373",
	":male-cook::skin-tone-6:": "1f468-1f3ff-200d-1f373",
	// 1F468-200D-1F37C man_feeding_baby
	":man_feeding_baby:":              "1f468-200d-1f37c",
	":man_feeding_baby::skin-tone-2:": "1f468-1f3fb-200d-1f37c",
	":man_feeding_baby::skin-tone-3:": "1f468-1f3fc-200d-1f37c",
	":man_feeding_baby::skin-tone-4:": "1f468-1f3fd-200d-1f37c",
	":man_feeding_baby::skin-tone-5:": "1f468-1f3fe-200d-1f37c",
	":man_feeding_baby::skin-tone-6:": "1f468-1f3ff-200d-1f37c",
	// 1F468-200D-1F393 male-student
	":male-student:":              "1f468-200d-1f393",
	":male-student::skin-tone-2:": "1f468-1f3fb-200d-1f393",
	":male-student::skin-tone-3:": "1f468-1f3fc-200d-1f393",
	":male-student::skin-tone-4:": "1f468-1f3fd-200d-1f393",
	":male-student::skin-tone-5:": "1f468-1f3fe-200d-1f393",
	":male-student::skin-tone-6:": "1f468-1f3ff-200d-1f393",
	// 1F468-200D-1F3A4 male-singer
	":male-singer:":              "1f468-200d-1f3a4",
	":male-singer::skin-tone-2:": "1f468-1f3fb-200d-1f3a4",
	":male-singer::skin-tone-3:": "1f468-1f3fc-200d-1f3a4",
	":male-singer::skin-tone-4:": "1f468-1f3fd-200d-1f3a4",
	":male-singer::skin-tone-5:": "1f468-1f3fe-200d-1f3a4",
	":male-singer::skin-tone-6:": "1f468-1f3ff-200d-1f3a4",
	// 1F468-200D-1F3A8 male-artist
	":male-artist:":              "1f468-200d-1f3a8",
	":male-artist::skin-tone-2:": "1f468-1f3fb-200d-1f3a8",
	":male-artist::skin-tone-3:": "1f468-1f3fc-200d-1f3a8",
	":male-artist::skin-tone-4:": "1f468-1f3fd-200d-1f3a8",
	":male-artist::skin-tone-5:": "1f468-1f3fe-200d-1f3a8",
	":male-artist::skin-tone-6:": "1f468-1f3ff-200d-1f3a8",
	// 1F468-200D-1F3EB male-teacher
	":male-teacher:":              "1f468-200d-1f3eb",
	":male-teacher::skin-tone-2:": "1f468-1f3fb-200d-1f3eb",
	":male-teacher::skin-tone-3:": "1f468-1f3fc-200d-1f3eb",
	":male-teacher::skin-tone-4:": "1f468-1f3fd-200d-1f3eb",
	":male-teacher::skin-tone-5:": "1f468-1f3fe-200d-1f3eb",
	":male-teacher::skin-tone-6:": "1f468-1f3ff-200d-1f3eb",
	// 1F468-200D-1F3ED male-factory-worker
	":male-factory-worker:":              "1f468-200d-1f3ed",
	":male-factory-worker::skin-tone-2:": "1f468-1f3fb-200d-1f3ed",
	":male-factory-worker::skin-tone-3:": "1f468-1f3fc-200d-1f3ed",
	":male-factory-worker::skin-tone-4:": "1f468-1f3fd-200d-1f3ed",
	":male-factory-worker::skin-tone-5:": "1f468-1f3fe-200d-1f3ed",
	":male-factory-worker::skin-tone-6:": "1f468-1f3ff-200d-1f3ed",
	// 1F468-200D-1F466-200D-1F466 man-boy-boy
	":man-boy-boy:": "1f468-200d-1f466-200d-1f466",
	// 1F468-200D-1F466 man-boy
	":man-boy:": "1f468-200d-1f466",
	// 1F468-200D-1F467-200D-1F466 man-girl-boy
	":man-girl-boy:": "1f468-200d-1f467-200d-1f466",
	// 1F468-200D-1F467-200D-1F467 man-girl-girl
	":man-girl-girl:": "1f468-200d-1f467-200d-1f467",
	// 1F468-200D-1F467 man-girl
	":man-girl:": "1f468-200d-1f467",
	// 1F468-200D-1F468-200D-1F466 man-man-boy
	":man-man-boy:": "1f468-200d-1f468-200d-1f466",
	// 1F468-200D-1F468-200D-1F466-200D-1F466 man-man-boy-boy
	":man-man-boy-boy:": "1f468-200d-1f468-200d-1f466-200d-1f466",
	// 1F468-200D-1F468-200D-1F467 man-man-girl
	":man-man-girl:": "1f468-200d-1f468-200d-1f467",
	// 1F468-200D-1F468-200D-1F467-200D-1F466 man-man-girl-boy
	":man-man-girl-boy:": "1f468-200d-1f468-200d-1f467-200d-1f466",
	// 1F468-200D-1F468-200D-1F467-200D-1F467 man-man-girl-girl
	":man-man-girl-girl:": "1f468-200d-1f468-200d-1f467-200d-1f467",
	// 1F468-200D-1F469-200D-1F466 man-woman-boy
	":man-woman-boy:": "1f468-200d-1f469-200d-1f466",
	// 1F468-200D-1F469-200D-1F466-200D-1F466 man-woman-boy-boy
	":man-woman-boy-boy:": "1f468-200d-1f469-200d-1f466-200d-1f466",
	// 1F468-200D-1F469-200D-1F467 man-woman-girl
	":man-woman-girl:": "1f468-200d-1f469-200d-1f467",
	// 1F468-200D-1F469-200D-1F467-200D-1F466 man-woman-girl-boy
	":man-woman-girl-boy:": "1f468-200d-1f469-200d-1f467-200d-1f466",
	// 1F468-200D-1F469-200D-1F467-200D-1F467 man-woman-girl-girl
	":man-woman-girl-girl:": "1f468-200d-1f469-200d-1f467-200d-1f467",
	// 1F468-200D-1F4BB male-technologist
	":male-technologist:":              "1f468-200d-1f4bb",
	":male-technologist::skin-tone-2:": "1f468-1f3fb-200d-1f4bb",
	":male-technologist::skin-tone-3:": "1f468-1f3fc-200d-1f4bb",
	":male-technologist::skin-tone-4:": "1f468-1f3fd-200d-1f4bb",
	":male-technologist::skin-tone-5:": "1f468-1f3fe-200d-1f4bb",
	":male-technologist::skin-tone-6:": "1f468-1f3ff-200d-1f4bb",
	// 1F468-200D-1F4BC male-office-worker
	":male-office-worker:":              "1f468-200d-1f4bc",
	":male-office-worker::skin-tone-2:": "1f468-1f3fb-200d-1f4bc",
	":male-office-worker::skin-tone-3:": "1f468-1f3fc-200d-1f4bc",
	":male-office-worker::skin-tone-4:": "1f468-1f3fd-200d-1f4bc",
	":male-office-worker::skin-tone-5:": "1f468-1f3fe-200d-1f4bc",
	":male-office-worker::skin-tone-6:": "1f468-1f3ff-200d-1f4bc",
	// 1F468-200D-1F527 male-mechanic
	":male-mechanic:":              "1f468-200d-1f527",
	":male-mechanic::skin-tone-2:": "1f468-1f3fb-200d-1f527",
	":male-mechanic::skin-tone-3:": "1f468-1f3fc-200d-1f527",
	":male-mechanic::skin-tone-4:": "1f468-1f3fd-200d-1f527",
	":male-mechanic::skin-tone-5:": "1f468-1f3fe-200d-1f527",
	":male-mechanic::skin-tone-6:": "1f468-1f3ff-200d-1f527",
	// 1F468-200D-1F52C male-scientist
	":male-scientist:":              "1f468-200d-1f52c",
	":male-scientist::skin-tone-2:": "1f468-1f3fb-200d-1f52c",
	":male-scientist::skin-tone-3:": "1f468-1f3fc-200d-1f52c",
	":male-scientist::skin-tone-4:": "1f468-1f3fd-200d-1f52c",
	":male-scientist::skin-tone-5:": "1f468-1f3fe-200d-1f52c",
	":male-scientist::skin-tone-6:": "1f468-1f3ff-200d-1f52c",
	// 1F468-200D-1F680 male-astronaut
	":male-astronaut:":              "1f468-200d-1f680",
	":male-astronaut::skin-tone-2:": "1f468-1f3fb-200d-1f680",
	":male-astronaut::skin-tone-3:": "1f468-1f3fc-200d-1f680",
	":male-astronaut::skin-tone-4:": "1f468-1f3fd-200d-1f680",
	":male-astronaut::skin-tone-5:": "1f468-1f3fe-200d-1f680",
	":male-astronaut::skin-tone-6:": "1f468-1f3ff-200d-1f680",
	// 1F468-200D-1F692 male-firefighter
	":male-firefighter:":              "1f468-200d-1f692",
	":male-firefighter::skin-tone-2:": "1f468-1f3fb-200d-1f692",
	":male-firefighter::skin-tone-3:": "1f468-1f3fc-200d-1f692",
	":male-firefighter::skin-tone-4:": "1f468-1f3fd-200d-1f692",
	":male-firefighter::skin-tone-5:": "1f468-1f3fe-200d-1f692",
	":male-firefighter::skin-tone-6:": "1f468-1f3ff-200d-1f692",
	// 1F468-200D-1F9AF man_with_probing_cane
	":man_with_probing_cane:":              "1f468-200d-1f9af",
	":man_with_probing_cane::skin-tone-2:": "1f468-1f3fb-200d-1f9af",
	":man_with_probing_cane::skin-tone-3:": "1f468-1f3fc-200d-1f9af",
	":man_with_probing_cane::skin-tone-4:": "1f468-1f3fd-200d-1f9af",
	":man_with_probing_cane::skin-tone-5:": "1f468-1f3fe-200d-1f9af",
	":man_with_probing_cane::skin-tone-6:": "1f468-1f3ff-200d-1f9af",
	// 1F468-200D-1F9B0 red_haired_man
	":red_haired_man:":              "1f468-200d-1f9b0",
	":red_haired_man::skin-tone-2:": "1f468-1f3fb-200d-1f9b0",
	":red_haired_man::skin-tone-3:": "1f468-1f3fc-200d-1f9b0",
	":red_haired_man::skin-tone-4:": "1f468-1f3fd-200d-1f9b0",
	":red_haired_man::skin-tone-5:": "1f468-1f3fe-200d-1f9b0",
	":red_haired_man::skin-tone-6:": "1f468-1f3ff-200d-1f9b0",
	// 1F468-200D-1F9B1 curly_haired_man
	":curly_haired_man:":              "1f468-200d-1f9b1",
	":curly_haired_man::skin-tone-2:": "1f468-1f3fb-200d-1f9b1",
	":curly_haired_man::skin-tone-3:": "1f468-1f3fc-200d-1f9b1",
	":curly_haired_man::skin-tone-4:": "1f468-1f3fd-200d-1f9b1",
	":curly_haired_man::skin-tone-5:": "1f468-1f3fe-200d-1f9b1",
	":curly_haired_man::skin-tone-6:": "1f468-1f3ff-200d-1f9b1",
	// 1F468-200D-1F9B2 bald_man
	":bald_man:":              "1f468-200d-1f9b2",
	":bald_man::skin-tone-2:": "1f468-1f3fb-200d-1f9b2",
	":bald_man::skin-tone-3:": "1f468-1f3fc-200d-1f9b2",
	":bald_man::skin-tone-4:": "1f468-1f3fd-200d-1f9b2",
	":bald_man::skin-tone-5:": "1f468-1f3fe-200d-1f9b2",
	":bald_man::skin-tone-6:": "1f468-1f3ff-200d-1f9b2",
	// 1F468-200D-1F9B3 white_haired_man
	":white_haired_man:":              "1f468-200d-1f9b3",
	":white_haired_man::skin-tone-2:": "1f468-1f3fb-200d-1f9b3",
	":white_haired_man::skin-tone-3:": "1f468-1f3fc-200d-1f9b3",
	":white_haired_man::skin-tone-4:": "1f468-1f3fd-200d-1f9b3",
	":white_haired_man::skin-tone-5:": "1f468-1f3fe-200d-1f9b3",
	":white_haired_man::skin-tone-6:": "1f468-1f3ff-200d-1f9b3",
	// 1F468-200D-1F9BC man_in_motorized_wheelchair
	":man_in_motorized_wheelchair:":              "1f468-200d-1f9bc",
	":man_in_motorized_wheelchair::skin-tone-2:": "1f468-1f3fb-200d-1f9bc",
	":man_in_motorized_wheelchair::skin-tone-3:": "1f468-1f3fc-200d-1f9bc",
	":man_in_motorized_wheelchair::skin-tone-4:": "1f468-1f3fd-200d-1f9bc",
	":man_in_motorized_wheelchair::skin-tone-5:": "1f468-1f3fe-200d-1f9bc",
	":man_in_motorized_wheelchair::skin-tone-6:": "1f468-1f3ff-200d-1f9bc",
	// 1F468-200D-1F9BD man_in_manual_wheelchair
	":man_in_manual_wheelchair:":              "1f468-200d-1f9bd",
	":man_in_manual_wheelchair::skin-tone-2:": "1f468-1f3fb-200d-1f9bd",
	":man_in_manual_wheelchair::skin-tone-3:": "1f468-1f3fc-200d-1f9bd",
	":man_in_manual_wheelchair::skin-tone-4:": "1f468-1f3fd-200d-1f9bd",
	":man_in_manual_wheelchair::skin-tone-5:": "1f468-1f3fe-200d-1f9bd",
	":man_in_manual_wheelchair::skin-tone-6:": "1f468-1f3ff-200d-1f9bd",
	// 1F468-200D-2695-FE0F male-doctor
	":male-doctor:":              "1f468-200d-2695-fe0f",
	":male-doctor::skin-tone-2:": "1f468-1f3fb-200d-2695-fe0f",
	":male-doctor::skin-tone-3:": "1f468-1f3fc-200d-2695-fe0f",
	":male-doctor::skin-tone-4:": "1f468-1f3fd-200d-2695-fe0f",
	":male-doctor::skin-tone-5:": "1f468-1f3fe-200d-2695-fe0f",
	":male-doctor::skin-tone-6:": "1f468-1f3ff-200d-2695-fe0f",
	// 1F468-200D-2696-FE0F male-judge
	":male-judge:":              "1f468-200d-2696-fe0f",
	":male-judge::skin-tone-2:": "1f468-1f3fb-200d-2696-fe0f",
	":male-judge::skin-tone-3:": "1f468-1f3fc-200d-2696-fe0f",
	":male-judge::skin-tone-4:": "1f468-1f3fd-200d-2696-fe0f",
	":male-judge::skin-tone-5:": "1f468-1f3fe-200d-2696-fe0f",
	":male-judge::skin-tone-6:": "1f468-1f3ff-200d-2696-fe0f",
	// 1F468-200D-2708-FE0F male-pilot
	":male-pilot:":              "1f468-200d-2708-fe0f",
	":male-pilot::skin-tone-2:": "1f468-1f3fb-200d-2708-fe0f",
	":male-pilot::skin-tone-3:": "1f468-1f3fc-200d-2708-fe0f",
	":male-pilot::skin-tone-4:": "1f468-1f3fd-200d-2708-fe0f",
	":male-pilot::skin-tone-5:": "1f468-1f3fe-200d-2708-fe0f",
	":male-pilot::skin-tone-6:": "1f468-1f3ff-200d-2708-fe0f",
	// 1F468-200D-2764-FE0F-200D-1F468 man-heart-man
	":man-heart-man:":                "1f468-200d-2764-fe0f-200d-1f468",
	":man-heart-man::skin-tone-2:":   "1f468-1f3fb-200d-2764-fe0f-200d-1f468-1f3fb",
	":man-heart-man::skin-tone-2-3:": "1f468-1f3fb-200d-2764-fe0f-200d-1f468-1f3fc",
	":man-heart-man::skin-tone-2-4:": "1f468-1f3fb-200d-2764-fe0f-200d-1f468-1f3fd",
	":man-heart-man::skin-tone-2-5:": "1f468-1f3fb-200d-2764-fe0f-200d-1f468-1f3fe",
	":man-heart-man::skin-tone-2-6:": "1f468-1f3fb-200d-2764-fe0f-200d-1f468-1f3ff",
	":man-heart-man::skin-tone-3-2:": "1f468-1f3fc-200d-2764-fe0f-200d-1f468-1f3fb",
	":man-heart-man::skin-tone-3:":   "1f468-1f3fc-200d-2764-fe0f-200d-1f468-1f3fc",
	":man-heart-man::skin-tone-3-4:": "1f468-1f3fc-200d-2764-fe0f-200d-1f468-1f3fd",
	":man-heart-man::skin-tone-3-5:": "1f468-1f3fc-200d-2764-fe0f-200d-1f468-1f3fe",
	":man-heart-man::skin-tone-3-6:": "1f468-1f3fc-200d-2764-fe0f-200d-1f468-1f3ff",
	":man-heart-man::skin-tone-4-2:": "1f468-1f3fd-200d-2764-fe0f-200d-1f468-1f3fb",
	":man-heart-man::skin-tone-4-3:": "1f468-1f3fd-200d-2764-fe0f-200d-1f468-1f3fc",
	":man-heart-man::skin-tone-4:":   "1f468-1f3fd-200d-2764-fe0f-200d-1f468-1f3fd",
	":man-heart-man::skin-tone-4-5:": "1f468-1f3fd-200d-2764-fe0f-200d-1f468-1f3fe",
	":man-heart-man::skin-tone-4-6:": "1f468-1f3fd-200d-2764-fe0f-200d-1f468-1f3ff",
	":man-heart-man::skin-tone-5-2:": "1f468-1f3fe-200d-2764-fe0f-200d-1f468-1f3fb",
	":man-heart-man::skin-tone-5-3:": "1f468-1f3fe-200d-2764-fe0f-200d-1f468-1f3fc",
	":man-heart-man::skin-tone-5-4:": "1f468-1f3fe-200d-2764-fe0f-200d-1f468-1f3fd",
	":man-heart-man::skin-tone-5:":   "1f468-1f3fe-200d-2764-fe0f-200d-1f468-1f3fe",
	":man-heart-man::skin-tone-5-6:": "1f468-1f3fe-200d-2764-fe0f-200d-1f468-1f3ff",
	":man-heart-man::skin-tone-6-2:": "1f468-1f3ff-200d-2764-fe0f-200d-1f468-1f3fb",
	":man-heart-man::skin-tone-6-3:": "1f468-1f3ff-200d-2764-fe0f-200d-1f468-1f3fc",
	":man-heart-man::skin-tone-6-4:": "1f468-1f3ff-200d-2764-fe0f-200d-1f468-1f3fd",
	":man-heart-man::skin-tone-6-5:": "1f468-1f3ff-200d-2764-fe0f-200d-1f468-1f3fe",
	":man-heart-man::skin-tone-6:":   "1f468-1f3ff-200d-2764-fe0f-200d-1f468-1f3ff",
	// 1F468-200D-2764-FE0F-200D-1F48B-200D-1F468 man-kiss-man
	":man-kiss-man:":                "1f468-200d-2764-fe0f-200d-1f48b-200d-1f468",
	":man-kiss-man::skin-tone-2:":   "1f468-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":man-kiss-man::skin-tone-2-3:": "1f468-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":man-kiss-man::skin-tone-2-4:": "1f468-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":man-kiss-man::skin-tone-2-5:": "1f468-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":man-kiss-man::skin-tone-2-6:": "1f468-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":man-kiss-man::skin-tone-3-2:": "1f468-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":man-kiss-man::skin-tone-3:":   "1f468-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":man-kiss-man::skin-tone-3-4:": "1f468-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":man-kiss-man::skin-tone-3-5:": "1f468-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":man-kiss-man::skin-tone-3-6:": "1f468-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":man-kiss-man::skin-tone-4-2:": "1f468-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":man-kiss-man::skin-tone-4-3:": "1f468-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":man-kiss-man::skin-tone-4:":   "1f468-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":man-kiss-man::skin-tone-4-5:": "1f468-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":man-kiss-man::skin-tone-4-6:": "1f468-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":man-kiss-man::skin-tone-5-2:": "1f468-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":man-kiss-man::skin-tone-5-3:": "1f468-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":man-kiss-man::skin-tone-5-4:": "1f468-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":man-kiss-man::skin-tone-5:":   "1f468-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":man-kiss-man::skin-tone-5-6:": "1f468-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":man-kiss-man::skin-tone-6-2:": "1f468-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":man-kiss-man::skin-tone-6-3:": "1f468-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":man-kiss-man::skin-tone-6-4:": "1f468-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":man-kiss-man::skin-tone-6-5:": "1f468-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":man-kiss-man::skin-tone-6:":   "1f468-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	// 1F468 man
	":man:":              "1f468",
	":man::skin-tone-2:": "1f468-1f3fb",
	":man::skin-tone-3:": "1f468-1f3fc",
	":man::skin-tone-4:": "1f468-1f3fd",
	":man::skin-tone-5:": "1f468-1f3fe",
	":man::skin-tone-6:": "1f468-1f3ff",
	// 1F469-200D-1F33E female-farmer
	":female-farmer:":              "1f469-200d-1f33e",
	":female-farmer::skin-tone-2:": "1f469-1f3fb-200d-1f33e",
	":female-farmer::skin-tone-3:": "1f469-1f3fc-200d-1f33e",
	":female-farmer::skin-tone-4:": "1f469-1f3fd-200d-1f33e",
	":female-farmer::skin-tone-5:": "1f469-1f3fe-200d-1f33e",
	":female-farmer::skin-tone-6:": "1f469-1f3ff-200d-1f33e",
	// 1F469-200D-1F373 female-cook
	":female-cook:":              "1f469-200d-1f373",
	":female-cook::skin-tone-2:": "1f469-1f3fb-200d-1f373",
	":female-cook::skin-tone-3:": "1f469-1f3fc-200d-1f373",
	":female-cook::skin-tone-4:": "1f469-1f3fd-200d-1f373",
	":female-cook::skin-tone-5:": "1f469-1f3fe-200d-1f373",
	":female-cook::skin-tone-6:": "1f469-1f3ff-200d-1f373",
	// 1F469-200D-1F37C woman_feeding_baby
	":woman_feeding_baby:":              "1f469-200d-1f37c",
	":woman_feeding_baby::skin-tone-2:": "1f469-1f3fb-200d-1f37c",
	":woman_feeding_baby::skin-tone-3:": "1f469-1f3fc-200d-1f37c",
	":woman_feeding_baby::skin-tone-4:": "1f469-1f3fd-200d-1f37c",
	":woman_feeding_baby::skin-tone-5:": "1f469-1f3fe-200d-1f37c",
	":woman_feeding_baby::skin-tone-6:": "1f469-1f3ff-200d-1f37c",
	// 1F469-200D-1F393 female-student
	":female-student:":              "1f469-200d-1f393",
	":female-student::skin-tone-2:": "1f469-1f3fb-200d-1f393",
	":female-student::skin-tone-3:": "1f469-1f3fc-200d-1f393",
	":female-student::skin-tone-4:": "1f469-1f3fd-200d-1f393",
	":female-student::skin-tone-5:": "1f469-1f3fe-200d-1f393",
	":female-student::skin-tone-6:": "1f469-1f3ff-200d-1f393",
	// 1F469-200D-1F3A4 female-singer
	":female-singer:":              "1f469-200d-1f3a4",
	":female-singer::skin-tone-2:": "1f469-1f3fb-200d-1f3a4",
	":female-singer::skin-tone-3:": "1f469-1f3fc-200d-1f3a4",
	":female-singer::skin-tone-4:": "1f469-1f3fd-200d-1f3a4",
	":female-singer::skin-tone-5:": "1f469-1f3fe-200d-1f3a4",
	":female-singer::skin-tone-6:": "1f469-1f3ff-200d-1f3a4",
	// 1F469-200D-1F3A8 female-artist
	":female-artist:":              "1f469-200d-1f3a8",
	":female-artist::skin-tone-2:": "1f469-1f3fb-200d-1f3a8",
	":female-artist::skin-tone-3:": "1f469-1f3fc-200d-1f3a8",
	":female-artist::skin-tone-4:": "1f469-1f3fd-200d-1f3a8",
	":female-artist::skin-tone-5:": "1f469-1f3fe-200d-1f3a8",
	":female-artist::skin-tone-6:": "1f469-1f3ff-200d-1f3a8",
	// 1F469-200D-1F3EB female-teacher
	":female-teacher:":              "1f469-200d-1f3eb",
	":female-teacher::skin-tone-2:": "1f469-1f3fb-200d-1f3eb",
	":female-teacher::skin-tone-3:": "1f469-1f3fc-200d-1f3eb",
	":female-teacher::skin-tone-4:": "1f469-1f3fd-200d-1f3eb",
	":female-teacher::skin-tone-5:": "1f469-1f3fe-200d-1f3eb",
	":female-teacher::skin-tone-6:": "1f469-1f3ff-200d-1f3eb",
	// 1F469-200D-1F3ED female-factory-worker
	":female-factory-worker:":              "1f469-200d-1f3ed",
	":female-factory-worker::skin-tone-2:": "1f469-1f3fb-200d-1f3ed",
	":female-factory-worker::skin-tone-3:": "1f469-1f3fc-200d-1f3ed",
	":female-factory-worker::skin-tone-4:": "1f469-1f3fd-200d-1f3ed",
	":female-factory-worker::skin-tone-5:": "1f469-1f3fe-200d-1f3ed",
	":female-factory-worker::skin-tone-6:": "1f469-1f3ff-200d-1f3ed",
	// 1F469-200D-1F466-200D-1F466 woman-boy-boy
	":woman-boy-boy:": "1f469-200d-1f466-200d-1f466",
	// 1F469-200D-1F466 woman-boy
	":woman-boy:": "1f469-200d-1f466",
	// 1F469-200D-1F467-200D-1F466 woman-girl-boy
	":woman-girl-boy:": "1f469-200d-1f467-200d-1f466",
	// 1F469-200D-1F467-200D-1F467 woman-girl-girl
	":woman-girl-girl:": "1f469-200d-1f467-200d-1f467",
	// 1F469-200D-1F467 woman-girl
	":woman-girl:": "1f469-200d-1f467",
	// 1F469-200D-1F469-200D-1F466 woman-woman-boy
	":woman-woman-boy:": "1f469-200d-1f469-200d-1f466",
	// 1F469-200D-1F469-200D-1F466-200D-1F466 woman-woman-boy-boy
	":woman-woman-boy-boy:": "1f469-200d-1f469-200d-1f466-200d-1f466",
	// 1F469-200D-1F469-200D-1F467 woman-woman-girl
	":woman-woman-girl:": "1f469-200d-1f469-200d-1f467",
	// 1F469-200D-1F469-200D-1F467-200D-1F466 woman-woman-girl-boy
	":woman-woman-girl-boy:": "1f469-200d-1f469-200d-1f467-200d-1f466",
	// 1F469-200D-1F469-200D-1F467-200D-1F467 woman-woman-girl-girl
	":woman-woman-girl-girl:": "1f469-200d-1f469-200d-1f467-200d-1f467",
	// 1F469-200D-1F4BB female-technologist
	":female-technologist:":              "1f469-200d-1f4bb",
	":female-technologist::skin-tone-2:": "1f469-1f3fb-200d-1f4bb",
	":female-technologist::skin-tone-3:": "1f469-1f3fc-200d-1f4bb",
	":female-technologist::skin-tone-4:": "1f469-1f3fd-200d-1f4bb",
	":female-technologist::skin-tone-5:": "1f469-1f3fe-200d-1f4bb",
	":female-technologist::skin-tone-6:": "1f469-1f3ff-200d-1f4bb",
	// 1F469-200D-1F4BC female-office-worker
	":female-office-worker:":              "1f469-200d-1f4bc",
	":female-office-worker::skin-tone-2:": "1f469-1f3fb-200d-1f4bc",
	":female-office-worker::skin-tone-3:": "1f469-1f3fc-200d-1f4bc",
	":female-office-worker::skin-tone-4:": "1f469-1f3fd-200d-1f4bc",
	":female-office-worker::skin-tone-5:": "1f469-1f3fe-200d-1f4bc",
	":female-office-worker::skin-tone-6:": "1f469-1f3ff-200d-1f4bc",
	// 1F469-200D-1F527 female-mechanic
	":female-mechanic:":              "1f469-200d-1f527",
	":female-mechanic::skin-tone-2:": "1f469-1f3fb-200d-1f527",
	":female-mechanic::skin-tone-3:": "1f469-1f3fc-200d-1f527",
	":female-mechanic::skin-tone-4:": "1f469-1f3fd-200d-1f527",
	":female-mechanic::skin-tone-5:": "1f469-1f3fe-200d-1f527",
	":female-mechanic::skin-tone-6:": "1f469-1f3ff-200d-1f527",
	// 1F469-200D-1F52C female-scientist
	":female-scientist:":              "1f469-200d-1f52c",
	":female-scientist::skin-tone-2:": "1f469-1f3fb-200d-1f52c",
	":female-scientist::skin-tone-3:": "1f469-1f3fc-200d-1f52c",
	":female-scientist::skin-tone-4:": "1f469-1f3fd-200d-1f52c",
	":female-scientist::skin-tone-5:": "1f469-1f3fe-200d-1f52c",
	":female-scientist::skin-tone-6:": "1f469-1f3ff-200d-1f52c",
	// 1F469-200D-1F680 female-astronaut
	":female-astronaut:":              "1f469-200d-1f680",
	":female-astronaut::skin-tone-2:": "1f469-1f3fb-200d-1f680",
	":female-astronaut::skin-tone-3:": "1f469-1f3fc-200d-1f680",
	":female-astronaut::skin-tone-4:": "1f469-1f3fd-200d-1f680",
	":female-astronaut::skin-tone-5:": "1f469-1f3fe-200d-1f680",
	":female-astronaut::skin-tone-6:": "1f469-1f3ff-200d-1f680",
	// 1F469-200D-1F692 female-firefighter
	":female-firefighter:":              "1f469-200d-1f692",
	":female-firefighter::skin-tone-2:": "1f469-1f3fb-200d-1f692",
	":female-firefighter::skin-tone-3:": "1f469-1f3fc-200d-1f692",
	":female-firefighter::skin-tone-4:": "1f469-1f3fd-200d-1f692",
	":female-firefighter::skin-tone-5:": "1f469-1f3fe-200d-1f692",
	":female-firefighter::skin-tone-6:": "1f469-1f3ff-200d-1f692",
	// 1F469-200D-1F9AF woman_with_probing_cane
	":woman_with_probing_cane:":              "1f469-200d-1f9af",
	":woman_with_probing_cane::skin-tone-2:": "1f469-1f3fb-200d-1f9af",
	":woman_with_probing_cane::skin-tone-3:": "1f469-1f3fc-200d-1f9af",
	":woman_with_probing_cane::skin-tone-4:": "1f469-1f3fd-200d-1f9af",
	":woman_with_probing_cane::skin-tone-5:": "1f469-1f3fe-200d-1f9af",
	":woman_with_probing_cane::skin-tone-6:": "1f469-1f3ff-200d-1f9af",
	// 1F469-200D-1F9B0 red_haired_woman
	":red_haired_woman:":              "1f469-200d-1f9b0",
	":red_haired_woman::skin-tone-2:": "1f469-1f3fb-200d-1f9b0",
	":red_haired_woman::skin-tone-3:": "1f469-1f3fc-200d-1f9b0",
	":red_haired_woman::skin-tone-4:": "1f469-1f3fd-200d-1f9b0",
	":red_haired_woman::skin-tone-5:": "1f469-1f3fe-200d-1f9b0",
	":red_haired_woman::skin-tone-6:": "1f469-1f3ff-200d-1f9b0",
	// 1F469-200D-1F9B1 curly_haired_woman
	":curly_haired_woman:":              "1f469-200d-1f9b1",
	":curly_haired_woman::skin-tone-2:": "1f469-1f3fb-200d-1f9b1",
	":curly_haired_woman::skin-tone-3:": "1f469-1f3fc-200d-1f9b1",
	":curly_haired_woman::skin-tone-4:": "1f469-1f3fd-200d-1f9b1",
	":curly_haired_woman::skin-tone-5:": "1f469-1f3fe-200d-1f9b1",
	":curly_haired_woman::skin-tone-6:": "1f469-1f3ff-200d-1f9b1",
	// 1F469-200D-1F9B2 bald_woman
	":bald_woman:":              "1f469-200d-1f9b2",
	":bald_woman::skin-tone-2:": "1f469-1f3fb-200d-1f9b2",
	":bald_woman::skin-tone-3:": "1f469-1f3fc-200d-1f9b2",
	":bald_woman::skin-tone-4:": "1f469-1f3fd-200d-1f9b2",
	":bald_woman::skin-tone-5:": "1f469-1f3fe-200d-1f9b2",
	":bald_woman::skin-tone-6:": "1f469-1f3ff-200d-1f9b2",
	// 1F469-200D-1F9B3 white_haired_woman
	":white_haired_woman:":              "1f469-200d-1f9b3",
	":white_haired_woman::skin-tone-2:": "1f469-1f3fb-200d-1f9b3",
	":white_haired_woman::skin-tone-3:": "1f469-1f3fc-200d-1f9b3",
	":white_haired_woman::skin-tone-4:": "1f469-1f3fd-200d-1f9b3",
	":white_haired_woman::skin-tone-5:": "1f469-1f3fe-200d-1f9b3",
	":white_haired_woman::skin-tone-6:": "1f469-1f3ff-200d-1f9b3",
	// 1F469-200D-1F9BC woman_in_motorized_wheelchair
	":woman_in_motorized_wheelchair:":              "1f469-200d-1f9bc",
	":woman_in_motorized_wheelchair::skin-tone-2:": "1f469-1f3fb-200d-1f9bc",
	":woman_in_motorized_wheelchair::skin-tone-3:": "1f469-1f3fc-200d-1f9bc",
	":woman_in_motorized_wheelchair::skin-tone-4:": "1f469-1f3fd-200d-1f9bc",
	":woman_in_motorized_wheelchair::skin-tone-5:": "1f469-1f3fe-200d-1f9bc",
	":woman_in_motorized_wheelchair::skin-tone-6:": "1f469-1f3ff-200d-1f9bc",
	// 1F469-200D-1F9BD woman_in_manual_wheelchair
	":woman_in_manual_wheelchair:":              "1f469-200d-1f9bd",
	":woman_in_manual_wheelchair::skin-tone-2:": "1f469-1f3fb-200d-1f9bd",
	":woman_in_manual_wheelchair::skin-tone-3:": "1f469-1f3fc-200d-1f9bd",
	":woman_in_manual_wheelchair::skin-tone-4:": "1f469-1f3fd-200d-1f9bd",
	":woman_in_manual_wheelchair::skin-tone-5:": "1f469-1f3fe-200d-1f9bd",
	":woman_in_manual_wheelchair::skin-tone-6:": "1f469-1f3ff-200d-1f9bd",
	// 1F469-200D-2695-FE0F female-doctor
	":female-doctor:":              "1f469-200d-2695-fe0f",
	":female-doctor::skin-tone-2:": "1f469-1f3fb-200d-2695-fe0f",
	":female-doctor::skin-tone-3:": "1f469-1f3fc-200d-2695-fe0f",
	":female-doctor::skin-tone-4:": "1f469-1f3fd-200d-2695-fe0f",
	":female-doctor::skin-tone-5:": "1f469-1f3fe-200d-2695-fe0f",
	":female-doctor::skin-tone-6:": "1f469-1f3ff-200d-2695-fe0f",
	// 1F469-200D-2696-FE0F female-judge
	":female-judge:":              "1f469-200d-2696-fe0f",
	":female-judge::skin-tone-2:": "1f469-1f3fb-200d-2696-fe0f",
	":female-judge::skin-tone-3:": "1f469-1f3fc-200d-2696-fe0f",
	":female-judge::skin-tone-4:": "1f469-1f3fd-200d-2696-fe0f",
	":female-judge::skin-tone-5:": "1f469-1f3fe-200d-2696-fe0f",
	":female-judge::skin-tone-6:": "1f469-1f3ff-200d-2696-fe0f",
	// 1F469-200D-2708-FE0F female-pilot
	":female-pilot:":              "1f469-200d-2708-fe0f",
	":female-pilot::skin-tone-2:": "1f469-1f3fb-200d-2708-fe0f",
	":female-pilot::skin-tone-3:": "1f469-1f3fc-200d-2708-fe0f",
	":female-pilot::skin-tone-4:": "1f469-1f3fd-200d-2708-fe0f",
	":female-pilot::skin-tone-5:": "1f469-1f3fe-200d-2708-fe0f",
	":female-pilot::skin-tone-6:": "1f469-1f3ff-200d-2708-fe0f",
	// 1F469-200D-2764-FE0F-200D-1F468 woman-heart-man
	":woman-heart-man:":                "1f469-200d-2764-fe0f-200d-1f468",
	":woman-heart-man::skin-tone-2:":   "1f469-1f3fb-200d-2764-fe0f-200d-1f468-1f3fb",
	":woman-heart-man::skin-tone-2-3:": "1f469-1f3fb-200d-2764-fe0f-200d-1f468-1f3fc",
	":woman-heart-man::skin-tone-2-4:": "1f469-1f3fb-200d-2764-fe0f-200d-1f468-1f3fd",
	":woman-heart-man::skin-tone-2-5:": "1f469-1f3fb-200d-2764-fe0f-200d-1f468-1f3fe",
	":woman-heart-man::skin-tone-2-6:": "1f469-1f3fb-200d-2764-fe0f-200d-1f468-1f3ff",
	":woman-heart-man::skin-tone-3-2:": "1f469-1f3fc-200d-2764-fe0f-200d-1f468-1f3fb",
	":woman-heart-man::skin-tone-3:":   "1f469-1f3fc-200d-2764-fe0f-200d-1f468-1f3fc",
	":woman-heart-man::skin-tone-3-4:": "1f469-1f3fc-200d-2764-fe0f-200d-1f468-1f3fd",
	":woman-heart-man::skin-tone-3-5:": "1f469-1f3fc-200d-2764-fe0f-200d-1f468-1f3fe",
	":woman-heart-man::skin-tone-3-6:": "1f469-1f3fc-200d-2764-fe0f-200d-1f468-1f3ff",
	":woman-heart-man::skin-tone-4-2:": "1f469-1f3fd-200d-2764-fe0f-200d-1f468-1f3fb",
	":woman-heart-man::skin-tone-4-3:": "1f469-1f3fd-200d-2764-fe0f-200d-1f468-1f3fc",
	":woman-heart-man::skin-tone-4:":   "1f469-1f3fd-200d-2764-fe0f-200d-1f468-1f3fd",
	":woman-heart-man::skin-tone-4-5:": "1f469-1f3fd-200d-2764-fe0f-200d-1f468-1f3fe",
	":woman-heart-man::skin-tone-4-6:": "1f469-1f3fd-200d-2764-fe0f-200d-1f468-1f3ff",
	":woman-heart-man::skin-tone-5-2:": "1f469-1f3fe-200d-2764-fe0f-200d-1f468-1f3fb",
	":woman-heart-man::skin-tone-5-3:": "1f469-1f3fe-200d-2764-fe0f-200d-1f468-1f3fc",
	":woman-heart-man::skin-tone-5-4:": "1f469-1f3fe-200d-2764-fe0f-200d-1f468-1f3fd",
	":woman-heart-man::skin-tone-5:":   "1f469-1f3fe-200d-2764-fe0f-200d-1f468-1f3fe",
	":woman-heart-man::skin-tone-5-6:": "1f469-1f3fe-200d-2764-fe0f-200d-1f468-1f3ff",
	":woman-heart-man::skin-tone-6-2:": "1f469-1f3ff-200d-2764-fe0f-200d-1f468-1f3fb",
	":woman-heart-man::skin-tone-6-3:": "1f469-1f3ff-200d-2764-fe0f-200d-1f468-1f3fc",
	":woman-heart-man::skin-tone-6-4:": "1f469-1f3ff-200d-2764-fe0f-200d-1f468-1f3fd",
	":woman-heart-man::skin-tone-6-5:": "1f469-1f3ff-200d-2764-fe0f-200d-1f468-1f3fe",
	":woman-heart-man::skin-tone-6:":   "1f469-1f3ff-200d-2764-fe0f-200d-1f468-1f3ff",
	// 1F469-200D-2764-FE0F-200D-1F469 woman-heart-woman
	":woman-heart-woman:":                "1f469-200d-2764-fe0f-200d-1f469",
	":woman-heart-woman::skin-tone-2:":   "1f469-1f3fb-200d-2764-fe0f-200d-1f469-1f3fb",
	":woman-heart-woman::skin-tone-2-3:": "1f469-1f3fb-200d-2764-fe0f-200d-1f469-1f3fc",
	":woman-heart-woman::skin-tone-2-4:": "1f469-1f3fb-200d-2764-fe0f-200d-1f469-1f3fd",
	":woman-heart-woman::skin-tone-2-5:": "1f469-1f3fb-200d-2764-fe0f-200d-1f469-1f3fe",
	":woman-heart-woman::skin-tone-2-6:": "1f469-1f3fb-200d-2764-fe0f-200d-1f469-1f3ff",
	":woman-heart-woman::skin-tone-3-2:": "1f469-1f3fc-200d-2764-fe0f-200d-1f469-1f3fb",
	":woman-heart-woman::skin-tone-3:":   "1f469-1f3fc-200d-2764-fe0f-200d-1f469-1f3fc",
	":woman-heart-woman::skin-tone-3-4:": "1f469-1f3fc-200d-2764-fe0f-200d-1f469-1f3fd",
	":woman-heart-woman::skin-tone-3-5:": "1f469-1f3fc-200d-2764-fe0f-200d-1f469-1f3fe",
	":woman-heart-woman::skin-tone-3-6:": "1f469-1f3fc-200d-2764-fe0f-200d-1f469-1f3ff",
	":woman-heart-woman::skin-tone-4-2:": "1f469-1f3fd-200d-2764-fe0f-200d-1f469-1f3fb",
	":woman-heart-woman::skin-tone-4-3:": "1f469-1f3fd-200d-2764-fe0f-200d-1f469-1f3fc",
	":woman-heart-woman::skin-tone-4:":   "1f469-1f3fd-200d-2764-fe0f-200d-1f469-1f3fd",
	":woman-heart-woman::skin-tone-4-5:": "1f469-1f3fd-200d-2764-fe0f-200d-1f469-1f3fe",
	":woman-heart-woman::skin-tone-4-6:": "1f469-1f3fd-200d-2764-fe0f-200d-1f469-1f3ff",
	":woman-heart-woman::skin-tone-5-2:": "1f469-1f3fe-200d-2764-fe0f-200d-1f469-1f3fb",
	":woman-heart-woman::skin-tone-5-3:": "1f469-1f3fe-200d-2764-fe0f-200d-1f469-1f3fc",
	":woman-heart-woman::skin-tone-5-4:": "1f469-1f3fe-200d-2764-fe0f-200d-1f469-1f3fd",
	":woman-heart-woman::skin-tone-5:":   "1f469-1f3fe-200d-2764-fe0f-200d-1f469-1f3fe",
	":woman-heart-woman::skin-tone-5-6:": "1f469-1f3fe-200d-2764-fe0f-200d-1f469-1f3ff",
	":woman-heart-woman::skin-tone-6-2:": "1f469-1f3ff-200d-2764-fe0f-200d-1f469-1f3fb",
	":woman-heart-woman::skin-tone-6-3:": "1f469-1f3ff-200d-2764-fe0f-200d-1f469-1f3fc",
	":woman-heart-woman::skin-tone-6-4:": "1f469-1f3ff-200d-2764-fe0f-200d-1f469-1f3fd",
	":woman-heart-woman::skin-tone-6-5:": "1f469-1f3ff-200d-2764-fe0f-200d-1f469-1f3fe",
	":woman-heart-woman::skin-tone-6:":   "1f469-1f3ff-200d-2764-fe0f-200d-1f469-1f3ff",
	// 1F469-200D-2764-FE0F-200D-1F48B-200D-1F468 woman-kiss-man
	":woman-kiss-man:":                "1f469-200d-2764-fe0f-200d-1f48b-200d-1f468",
	":woman-kiss-man::skin-tone-2:":   "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":woman-kiss-man::skin-tone-2-3:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":woman-kiss-man::skin-tone-2-4:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":woman-kiss-man::skin-tone-2-5:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":woman-kiss-man::skin-tone-2-6:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":woman-kiss-man::skin-tone-3-2:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":woman-kiss-man::skin-tone-3:":   "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":woman-kiss-man::skin-tone-3-4:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":woman-kiss-man::skin-tone-3-5:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":woman-kiss-man::skin-tone-3-6:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":woman-kiss-man::skin-tone-4-2:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":woman-kiss-man::skin-tone-4-3:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":woman-kiss-man::skin-tone-4:":   "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":woman-kiss-man::skin-tone-4-5:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":woman-kiss-man::skin-tone-4-6:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":woman-kiss-man::skin-tone-5-2:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":woman-kiss-man::skin-tone-5-3:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":woman-kiss-man::skin-tone-5-4:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":woman-kiss-man::skin-tone-5:":   "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":woman-kiss-man::skin-tone-5-6:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	":woman-kiss-man::skin-tone-6-2:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fb",
	":woman-kiss-man::skin-tone-6-3:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fc",
	":woman-kiss-man::skin-tone-6-4:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fd",
	":woman-kiss-man::skin-tone-6-5:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3fe",
	":woman-kiss-man::skin-tone-6:":   "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f468-1f3ff",
	// 1F469-200D-2764-FE0F-200D-1F48B-200D-1F469 woman-kiss-woman
	":woman-kiss-woman:":                "1f469-200d-2764-fe0f-200d-1f48b-200d-1f469",
	":woman-kiss-woman::skin-tone-2:":   "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fb",
	":woman-kiss-woman::skin-tone-2-3:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fc",
	":woman-kiss-woman::skin-tone-2-4:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fd",
	":woman-kiss-woman::skin-tone-2-5:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fe",
	":woman-kiss-woman::skin-tone-2-6:": "1f469-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3ff",
	":woman-kiss-woman::skin-tone-3-2:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fb",
	":woman-kiss-woman::skin-tone-3:":   "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fc",
	":woman-kiss-woman::skin-tone-3-4:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fd",
	":woman-kiss-woman::skin-tone-3-5:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fe",
	":woman-kiss-woman::skin-tone-3-6:": "1f469-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3ff",
	":woman-kiss-woman::skin-tone-4-2:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fb",
	":woman-kiss-woman::skin-tone-4-3:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fc",
	":woman-kiss-woman::skin-tone-4:":   "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fd",
	":woman-kiss-woman::skin-tone-4-5:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fe",
	":woman-kiss-woman::skin-tone-4-6:": "1f469-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3ff",
	":woman-kiss-woman::skin-tone-5-2:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fb",
	":woman-kiss-woman::skin-tone-5-3:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fc",
	":woman-kiss-woman::skin-tone-5-4:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fd",
	":woman-kiss-woman::skin-tone-5:":   "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fe",
	":woman-kiss-woman::skin-tone-5-6:": "1f469-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3ff",
	":woman-kiss-woman::skin-tone-6-2:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fb",
	":woman-kiss-woman::skin-tone-6-3:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fc",
	":woman-kiss-woman::skin-tone-6-4:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fd",
	":woman-kiss-woman::skin-tone-6-5:": "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3fe",
	":woman-kiss-woman::skin-tone-6:":   "1f469-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f469-1f3ff",
	// 1F469 woman
	":woman:":              "1f469",
	":woman::skin-tone-2:": "1f469-1f3fb",
	":woman::skin-tone-3:": "1f469-1f3fc",
	":woman::skin-tone-4:": "1f469-1f3fd",
	":woman::skin-tone-5:": "1f469-1f3fe",
	":woman::skin-tone-6:": "1f469-1f3ff",
	// 1F46A family
	":family:": "1f46a",
	// 1F46B man_and_woman_holding_hands
	":man_and_woman_holding_hands:": "1f46b",
	// 1F46B man_and_woman_holding_hands
	":woman_and_man_holding_hands:": "1f46b",
	// 1F46B man_and_woman_holding_hands
	":couple:": "1f46b",
	":man_and_woman_holding_hands::skin-tone-2:":   "1f46b-1f3fb",
	":woman_and_man_holding_hands::skin-tone-2:":   "1f46b-1f3fb",
	":couple::skin-tone-2:":                        "1f46b-1f3fb",
	":man_and_woman_holding_hands::skin-tone-3:":   "1f46b-1f3fc",
	":woman_and_man_holding_hands::skin-tone-3:":   "1f46b-1f3fc",
	":couple::skin-tone-3:":                        "1f46b-1f3fc",
	":man_and_woman_holding_hands::skin-tone-4:":   "1f46b-1f3fd",
	":woman_and_man_holding_hands::skin-tone-4:":   "1f46b-1f3fd",
	":couple::skin-tone-4:":                        "1f46b-1f3fd",
	":man_and_woman_holding_hands::skin-tone-5:":   "1f46b-1f3fe",
	":woman_and_man_holding_hands::skin-tone-5:":   "1f46b-1f3fe",
	":couple::skin-tone-5:":                        "1f46b-1f3fe",
	":man_and_woman_holding_hands::skin-tone-6:":   "1f46b-1f3ff",
	":woman_and_man_holding_hands::skin-tone-6:":   "1f46b-1f3ff",
	":couple::skin-tone-6:":                        "1f46b-1f3ff",
	":man_and_woman_holding_hands::skin-tone-2-3:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fc",
	":woman_and_man_holding_hands::skin-tone-2-3:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fc",
	":couple::skin-tone-2-3:":                      "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fc",
	":man_and_woman_holding_hands::skin-tone-2-4:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fd",
	":woman_and_man_holding_hands::skin-tone-2-4:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fd",
	":couple::skin-tone-2-4:":                      "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fd",
	":man_and_woman_holding_hands::skin-tone-2-5:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fe",
	":woman_and_man_holding_hands::skin-tone-2-5:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fe",
	":couple::skin-tone-2-5:":                      "1f469-1f3fb-200d-1f91d-200d-1f468-1f3fe",
	":man_and_woman_holding_hands::skin-tone-2-6:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3ff",
	":woman_and_man_holding_hands::skin-tone-2-6:": "1f469-1f3fb-200d-1f91d-200d-1f468-1f3ff",
	":couple::skin-tone-2-6:":                      "1f469-1f3fb-200d-1f91d-200d-1f468-1f3ff",
	":man_and_woman_holding_hands::skin-tone-3-2:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fb",
	":woman_and_man_holding_hands::skin-tone-3-2:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fb",
	":couple::skin-tone-3-2:":                      "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fb",
	":man_and_woman_holding_hands::skin-tone-3-4:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fd",
	":woman_and_man_holding_hands::skin-tone-3-4:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fd",
	":couple::skin-tone-3-4:":                      "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fd",
	":man_and_woman_holding_hands::skin-tone-3-5:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fe",
	":woman_and_man_holding_hands::skin-tone-3-5:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fe",
	":couple::skin-tone-3-5:":                      "1f469-1f3fc-200d-1f91d-200d-1f468-1f3fe",
	":man_and_woman_holding_hands::skin-tone-3-6:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3ff",
	":woman_and_man_holding_hands::skin-tone-3-6:": "1f469-1f3fc-200d-1f91d-200d-1f468-1f3ff",
	":couple::skin-tone-3-6:":                      "1f469-1f3fc-200d-1f91d-200d-1f468-1f3ff",
	":man_and_woman_holding_hands::skin-tone-4-2:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fb",
	":woman_and_man_holding_hands::skin-tone-4-2:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fb",
	":couple::skin-tone-4-2:":                      "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fb",
	":man_and_woman_holding_hands::skin-tone-4-3:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fc",
	":woman_and_man_holding_hands::skin-tone-4-3:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fc",
	":couple::skin-tone-4-3:":                      "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fc",
	":man_and_woman_holding_hands::skin-tone-4-5:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fe",
	":woman_and_man_holding_hands::skin-tone-4-5:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fe",
	":couple::skin-tone-4-5:":                      "1f469-1f3fd-200d-1f91d-200d-1f468-1f3fe",
	":man_and_woman_holding_hands::skin-tone-4-6:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3ff",
	":woman_and_man_holding_hands::skin-tone-4-6:": "1f469-1f3fd-200d-1f91d-200d-1f468-1f3ff",
	":couple::skin-tone-4-6:":                      "1f469-1f3fd-200d-1f91d-200d-1f468-1f3ff",
	":man_and_woman_holding_hands::skin-tone-5-2:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fb",
	":woman_and_man_holding_hands::skin-tone-5-2:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fb",
	":couple::skin-tone-5-2:":                      "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fb",
	":man_and_woman_holding_hands::skin-tone-5-3:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fc",
	":woman_and_man_holding_hands::skin-tone-5-3:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fc",
	":couple::skin-tone-5-3:":                      "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fc",
	":man_and_woman_holding_hands::skin-tone-5-4:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fd",
	":woman_and_man_holding_hands::skin-tone-5-4:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fd",
	":couple::skin-tone-5-4:":                      "1f469-1f3fe-200d-1f91d-200d-1f468-1f3fd",
	":man_and_woman_holding_hands::skin-tone-5-6:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3ff",
	":woman_and_man_holding_hands::skin-tone-5-6:": "1f469-1f3fe-200d-1f91d-200d-1f468-1f3ff",
	":couple::skin-tone-5-6:":                      "1f469-1f3fe-200d-1f91d-200d-1f468-1f3ff",
	":man_and_woman_holding_hands::skin-tone-6-2:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fb",
	":woman_and_man_holding_hands::skin-tone-6-2:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fb",
	":couple::skin-tone-6-2:":                      "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fb",
	":man_and_woman_holding_hands::skin-tone-6-3:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fc",
	":woman_and_man_holding_hands::skin-tone-6-3:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fc",
	":couple::skin-tone-6-3:":                      "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fc",
	":man_and_woman_holding_hands::skin-tone-6-4:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fd",
	":woman_and_man_holding_hands::skin-tone-6-4:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fd",
	":couple::skin-tone-6-4:":                      "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fd",
	":man_and_woman_holding_hands::skin-tone-6-5:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fe",
	":woman_and_man_holding_hands::skin-tone-6-5:": "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fe",
	":couple::skin-tone-6-5:":                      "1f469-1f3ff-200d-1f91d-200d-1f468-1f3fe",
	// 1F46C two_men_holding_hands
	":two_men_holding_hands:": "1f46c",
	// 1F46C two_men_holding_hands
	":men_holding_hands:":                    "1f46c",
	":two_men_holding_hands::skin-tone-2:":   "1f46c-1f3fb",
	":men_holding_hands::skin-tone-2:":       "1f46c-1f3fb",
	":two_men_holding_hands::skin-tone-3:":   "1f46c-1f3fc",
	":men_holding_hands::skin-tone-3:":       "1f46c-1f3fc",
	":two_men_holding_hands::skin-tone-4:":   "1f46c-1f3fd",
	":men_holding_hands::skin-tone-4:":       "1f46c-1f3fd",
	":two_men_holding_hands::skin-tone-5:":   "1f46c-1f3fe",
	":men_holding_hands::skin-tone-5:":       "1f46c-1f3fe",
	":two_men_holding_hands::skin-tone-6:":   "1f46c-1f3ff",
	":men_holding_hands::skin-tone-6:":       "1f46c-1f3ff",
	":two_men_holding_hands::skin-tone-2-3:": "1f468-1f3fb-200d-1f91d-200d-1f468-1f3fc",
	":men_holding_hands::skin-tone-2-3:":     "1f468-1f3fb-200d-1f91d-200d-1f468-1f3fc",
	":two_men_holding_hands::skin-tone-2-4:": "1f468-1f3fb-200d-1f91d-200d-1f468-1f3fd",
	":men_holding_hands::skin-tone-2-4:":     "1f468-1f3fb-200d-1f91d-200d-1f468-1f3fd",
	":two_men_holding_hands::skin-tone-2-5:": "1f468-1f3fb-200d-1f91d-200d-1f468-1f3fe",
	":men_holding_hands::skin-tone-2-5:":     "1f468-1f3fb-200d-1f91d-200d-1f468-1f3fe",
	":two_men_holding_hands::skin-tone-2-6:": "1f468-1f3fb-200d-1f91d-200d-1f468-1f3ff",
	":men_holding_hands::skin-tone-2-6:":     "1f468-1f3fb-200d-1f91d-200d-1f468-1f3ff",
	":two_men_holding_hands::skin-tone-3-2:": "1f468-1f3fc-200d-1f91d-200d-1f468-1f3fb",
	":men_holding_hands::skin-tone-3-2:":     "1f468-1f3fc-200d-1f91d-200d-1f468-1f3fb",
	":two_men_holding_hands::skin-tone-3-4:": "1f468-1f3fc-200d-1f91d-200d-1f468-1f3fd",
	":men_holding_hands::skin-tone-3-4:":     "1f468-1f3fc-200d-1f91d-200d-1f468-1f3fd",
	":two_men_holding_hands::skin-tone-3-5:": "1f468-1f3fc-200d-1f91d-200d-1f468-1f3fe",
	":men_holding_hands::skin-tone-3-5:":     "1f468-1f3fc-200d-1f91d-200d-1f468-1f3fe",
	":two_men_holding_hands::skin-tone-3-6:": "1f468-1f3fc-200d-1f91d-200d-1f468-1f3ff",
	":men_holding_hands::skin-tone-3-6:":     "1f468-1f3fc-200d-1f91d-200d-1f468-1f3ff",
	":two_men_holding_hands::skin-tone-4-2:": "1f468-1f3fd-200d-1f91d-200d-1f468-1f3fb",
	":men_holding_hands::skin-tone-4-2:":     "1f468-1f3fd-200d-1f91d-200d-1f468-1f3fb",
	":two_men_holding_hands::skin-tone-4-3:": "1f468-1f3fd-200d-1f91d-200d-1f468-1f3fc",
	":men_holding_hands::skin-tone-4-3:":     "1f468-1f3fd-200d-1f91d-200d-1f468-1f3fc",
	":two_men_holding_hands::skin-tone-4-5:": "1f468-1f3fd-200d-1f91d-200d-1f468-1f3fe",
	":men_holding_hands::skin-tone-4-5:":     "1f468-1f3fd-200d-1f91d-200d-1f468-1f3fe",
	":two_men_holding_hands::skin-tone-4-6:": "1f468-1f3fd-200d-1f91d-200d-1f468-1f3ff",
	":men_holding_hands::skin-tone-4-6:":     "1f468-1f3fd-200d-1f91d-200d-1f468-1f3ff",
	":two_men_holding_hands::skin-tone-5-2:": "1f468-1f3fe-200d-1f91d-200d-1f468-1f3fb",
	":men_holding_hands::skin-tone-5-2:":     "1f468-1f3fe-200d-1f91d-200d-1f468-1f3fb",
	":two_men_holding_hands::skin-tone-5-3:": "1f468-1f3fe-200d-1f91d-200d-1f468-1f3fc",
	":men_holding_hands::skin-tone-5-3:":     "1f468-1f3fe-200d-1f91d-200d-1f468-1f3fc",
	":two_men_holding_hands::skin-tone-5-4:": "1f468-1f3fe-200d-1f91d-200d-1f468-1f3fd",
	":men_holding_hands::skin-tone-5-4:":     "1f468-1f3fe-200d-1f91d-200d-1f468-1f3fd",
	":two_men_holding_hands::skin-tone-5-6:": "1f468-1f3fe-200d-1f91d-200d-1f468-1f3ff",
	":men_holding_hands::skin-tone-5-6:":     "1f468-1f3fe-200d-1f91d-200d-1f468-1f3ff",
	":two_men_holding_hands::skin-tone-6-2:": "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fb",
	":men_holding_hands::skin-tone-6-2:":     "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fb",
	":two_men_holding_hands::skin-tone-6-3:": "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fc",
	":men_holding_hands::skin-tone-6-3:":     "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fc",
	":two_men_holding_hands::skin-tone-6-4:": "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fd",
	":men_holding_hands::skin-tone-6-4:":     "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fd",
	":two_men_holding_hands::skin-tone-6-5:": "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fe",
	":men_holding_hands::skin-tone-6-5:":     "1f468-1f3ff-200d-1f91d-200d-1f468-1f3fe",
	// 1F46D two_women_holding_hands
	":two_women_holding_hands:": "1f46d",
	// 1F46D two_women_holding_hands
	":women_holding_hands:":                    "1f46d",
	":two_women_holding_hands::skin-tone-2:":   "1f46d-1f3fb",
	":women_holding_hands::skin-tone-2:":       "1f46d-1f3fb",
	":two_women_holding_hands::skin-tone-3:":   "1f46d-1f3fc",
	":women_holding_hands::skin-tone-3:":       "1f46d-1f3fc",
	":two_women_holding_hands::skin-tone-4:":   "1f46d-1f3fd",
	":women_holding_hands::skin-tone-4:":       "1f46d-1f3fd",
	":two_women_holding_hands::skin-tone-5:":   "1f46d-1f3fe",
	":women_holding_hands::skin-tone-5:":       "1f46d-1f3fe",
	":two_women_holding_hands::skin-tone-6:":   "1f46d-1f3ff",
	":women_holding_hands::skin-tone-6:":       "1f46d-1f3ff",
	":two_women_holding_hands::skin-tone-2-3:": "1f469-1f3fb-200d-1f91d-200d-1f469-1f3fc",
	":women_holding_hands::skin-tone-2-3:":     "1f469-1f3fb-200d-1f91d-200d-1f469-1f3fc",
	":two_women_holding_hands::skin-tone-2-4:": "1f469-1f3fb-200d-1f91d-200d-1f469-1f3fd",
	":women_holding_hands::skin-tone-2-4:":     "1f469-1f3fb-200d-1f91d-200d-1f469-1f3fd",
	":two_women_holding_hands::skin-tone-2-5:": "1f469-1f3fb-200d-1f91d-200d-1f469-1f3fe",
	":women_holding_hands::skin-tone-2-5:":     "1f469-1f3fb-200d-1f91d-200d-1f469-1f3fe",
	":two_women_holding_hands::skin-tone-2-6:": "1f469-1f3fb-200d-1f91d-200d-1f469-1f3ff",
	":women_holding_hands::skin-tone-2-6:":     "1f469-1f3fb-200d-1f91d-200d-1f469-1f3ff",
	":two_women_holding_hands::skin-tone-3-2:": "1f469-1f3fc-200d-1f91d-200d-1f469-1f3fb",
	":women_holding_hands::skin-tone-3-2:":     "1f469-1f3fc-200d-1f91d-200d-1f469-1f3fb",
	":two_women_holding_hands::skin-tone-3-4:": "1f469-1f3fc-200d-1f91d-200d-1f469-1f3fd",
	":women_holding_hands::skin-tone-3-4:":     "1f469-1f3fc-200d-1f91d-200d-1f469-1f3fd",
	":two_women_holding_hands::skin-tone-3-5:": "1f469-1f3fc-200d-1f91d-200d-1f469-1f3fe",
	":women_holding_hands::skin-tone-3-5:":     "1f469-1f3fc-200d-1f91d-200d-1f469-1f3fe",
	":two_women_holding_hands::skin-tone-3-6:": "1f469-1f3fc-200d-1f91d-200d-1f469-1f3ff",
	":women_holding_hands::skin-tone-3-6:":     "1f469-1f3fc-200d-1f91d-200d-1f469-1f3ff",
	":two_women_holding_hands::skin-tone-4-2:": "1f469-1f3fd-200d-1f91d-200d-1f469-1f3fb",
	":women_holding_hands::skin-tone-4-2:":     "1f469-1f3fd-200d-1f91d-200d-1f469-1f3fb",
	":two_women_holding_hands::skin-tone-4-3:": "1f469-1f3fd-200d-1f91d-200d-1f469-1f3fc",
	":women_holding_hands::skin-tone-4-3:":     "1f469-1f3fd-200d-1f91d-200d-1f469-1f3fc",
	":two_women_holding_hands::skin-tone-4-5:": "1f469-1f3fd-200d-1f91d-200d-1f469-1f3fe",
	":women_holding_hands::skin-tone-4-5:":     "1f469-1f3fd-200d-1f91d-200d-1f469-1f3fe",
	":two_women_holding_hands::skin-tone-4-6:": "1f469-1f3fd-200d-1f91d-200d-1f469-1f3ff",
	":women_holding_hands::skin-tone-4-6:":     "1f469-1f3fd-200d-1f91d-200d-1f469-1f3ff",
	":two_women_holding_hands::skin-tone-5-2:": "1f469-1f3fe-200d-1f91d-200d-1f469-1f3fb",
	":women_holding_hands::skin-tone-5-2:":     "1f469-1f3fe-200d-1f91d-200d-1f469-1f3fb",
	":two_women_holding_hands::skin-tone-5-3:": "1f469-1f3fe-200d-1f91d-200d-1f469-1f3fc",
	":women_holding_hands::skin-tone-5-3:":     "1f469-1f3fe-200d-1f91d-200d-1f469-1f3fc",
	":two_women_holding_hands::skin-tone-5-4:": "1f469-1f3fe-200d-1f91d-200d-1f469-1f3fd",
	":women_holding_hands::skin-tone-5-4:":     "1f469-1f3fe-200d-1f91d-200d-1f469-1f3fd",
	":two_women_holding_hands::skin-tone-5-6:": "1f469-1f3fe-200d-1f91d-200d-1f469-1f3ff",
	":women_holding_hands::skin-tone-5-6:":     "1f469-1f3fe-200d-1f91d-200d-1f469-1f3ff",
	":two_women_holding_hands::skin-tone-6-2:": "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fb",
	":women_holding_hands::skin-tone-6-2:":     "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fb",
	":two_women_holding_hands::skin-tone-6-3:": "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fc",
	":women_holding_hands::skin-tone-6-3:":     "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fc",
	":two_women_holding_hands::skin-tone-6-4:": "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fd",
	":women_holding_hands::skin-tone-6-4:":     "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fd",
	":two_women_holding_hands::skin-tone-6-5:": "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fe",
	":women_holding_hands::skin-tone-6-5:":     "1f469-1f3ff-200d-1f91d-200d-1f469-1f3fe",
	// 1F46E-200D-2640-FE0F female-police-officer
	":female-police-officer:":              "1f46e-200d-2640-fe0f",
	":female-police-officer::skin-tone-2:": "1f46e-1f3fb-200d-2640-fe0f",
	":female-police-officer::skin-tone-3:": "1f46e-1f3fc-200d-2640-fe0f",
	":female-police-officer::skin-tone-4:": "1f46e-1f3fd-200d-2640-fe0f",
	":female-police-officer::skin-tone-5:": "1f46e-1f3fe-200d-2640-fe0f",
	":female-police-officer::skin-tone-6:": "1f46e-1f3ff-200d-2640-fe0f",
	// 1F46E-200D-2642-FE0F male-police-officer
	":male-police-officer:":              "1f46e-200d-2642-fe0f",
	":male-police-officer::skin-tone-2:": "1f46e-1f3fb-200d-2642-fe0f",
	":male-police-officer::skin-tone-3:": "1f46e-1f3fc-200d-2642-fe0f",
	":male-police-officer::skin-tone-4:": "1f46e-1f3fd-200d-2642-fe0f",
	":male-police-officer::skin-tone-5:": "1f46e-1f3fe-200d-2642-fe0f",
	":male-police-officer::skin-tone-6:": "1f46e-1f3ff-200d-2642-fe0f",
	// 1F46E cop
	":cop:":              "1f46e",
	":cop::skin-tone-2:": "1f46e-1f3fb",
	":cop::skin-tone-3:": "1f46e-1f3fc",
	":cop::skin-tone-4:": "1f46e-1f3fd",
	":cop::skin-tone-5:": "1f46e-1f3fe",
	":cop::skin-tone-6:": "1f46e-1f3ff",
	// 1F46F-200D-2640-FE0F women-with-bunny-ears-partying
	":women-with-bunny-ears-partying:": "1f46f-200d-2640-fe0f",
	// 1F46F-200D-2640-FE0F women-with-bunny-ears-partying
	":woman-with-bunny-ears-partying:": "1f46f-200d-2640-fe0f",
	// 1F46F-200D-2642-FE0F men-with-bunny-ears-partying
	":men-with-bunny-ears-partying:": "1f46f-200d-2642-fe0f",
	// 1F46F-200D-2642-FE0F men-with-bunny-ears-partying
	":man-with-bunny-ears-partying:": "1f46f-200d-2642-fe0f",
	// 1F46F dancers
	":dancers:": "1f46f",
	// 1F470-200D-2640-FE0F woman_with_veil
	":woman_with_veil:":              "1f470-200d-2640-fe0f",
	":woman_with_veil::skin-tone-2:": "1f470-1f3fb-200d-2640-fe0f",
	":woman_with_veil::skin-tone-3:": "1f470-1f3fc-200d-2640-fe0f",
	":woman_with_veil::skin-tone-4:": "1f470-1f3fd-200d-2640-fe0f",
	":woman_with_veil::skin-tone-5:": "1f470-1f3fe-200d-2640-fe0f",
	":woman_with_veil::skin-tone-6:": "1f470-1f3ff-200d-2640-fe0f",
	// 1F470-200D-2642-FE0F man_with_veil
	":man_with_veil:":              "1f470-200d-2642-fe0f",
	":man_with_veil::skin-tone-2:": "1f470-1f3fb-200d-2642-fe0f",
	":man_with_veil::skin-tone-3:": "1f470-1f3fc-200d-2642-fe0f",
	":man_with_veil::skin-tone-4:": "1f470-1f3fd-200d-2642-fe0f",
	":man_with_veil::skin-tone-5:": "1f470-1f3fe-200d-2642-fe0f",
	":man_with_veil::skin-tone-6:": "1f470-1f3ff-200d-2642-fe0f",
	// 1F470 bride_with_veil
	":bride_with_veil:":              "1f470",
	":bride_with_veil::skin-tone-2:": "1f470-1f3fb",
	":bride_with_veil::skin-tone-3:": "1f470-1f3fc",
	":bride_with_veil::skin-tone-4:": "1f470-1f3fd",
	":bride_with_veil::skin-tone-5:": "1f470-1f3fe",
	":bride_with_veil::skin-tone-6:": "1f470-1f3ff",
	// 1F471-200D-2640-FE0F blond-haired-woman
	":blond-haired-woman:":              "1f471-200d-2640-fe0f",
	":blond-haired-woman::skin-tone-2:": "1f471-1f3fb-200d-2640-fe0f",
	":blond-haired-woman::skin-tone-3:": "1f471-1f3fc-200d-2640-fe0f",
	":blond-haired-woman::skin-tone-4:": "1f471-1f3fd-200d-2640-fe0f",
	":blond-haired-woman::skin-tone-5:": "1f471-1f3fe-200d-2640-fe0f",
	":blond-haired-woman::skin-tone-6:": "1f471-1f3ff-200d-2640-fe0f",
	// 1F471-200D-2642-FE0F blond-haired-man
	":blond-haired-man:":              "1f471-200d-2642-fe0f",
	":blond-haired-man::skin-tone-2:": "1f471-1f3fb-200d-2642-fe0f",
	":blond-haired-man::skin-tone-3:": "1f471-1f3fc-200d-2642-fe0f",
	":blond-haired-man::skin-tone-4:": "1f471-1f3fd-200d-2642-fe0f",
	":blond-haired-man::skin-tone-5:": "1f471-1f3fe-200d-2642-fe0f",
	":blond-haired-man::skin-tone-6:": "1f471-1f3ff-200d-2642-fe0f",
	// 1F471 person_with_blond_hair
	":person_with_blond_hair:":              "1f471",
	":person_with_blond_hair::skin-tone-2:": "1f471-1f3fb",
	":person_with_blond_hair::skin-tone-3:": "1f471-1f3fc",
	":person_with_blond_hair::skin-tone-4:": "1f471-1f3fd",
	":person_with_blond_hair::skin-tone-5:": "1f471-1f3fe",
	":person_with_blond_hair::skin-tone-6:": "1f471-1f3ff",
	// 1F472 man_with_gua_pi_mao
	":man_with_gua_pi_mao:":              "1f472",
	":man_with_gua_pi_mao::skin-tone-2:": "1f472-1f3fb",
	":man_with_gua_pi_mao::skin-tone-3:": "1f472-1f3fc",
	":man_with_gua_pi_mao::skin-tone-4:": "1f472-1f3fd",
	":man_with_gua_pi_mao::skin-tone-5:": "1f472-1f3fe",
	":man_with_gua_pi_mao::skin-tone-6:": "1f472-1f3ff",
	// 1F473-200D-2640-FE0F woman-wearing-turban
	":woman-wearing-turban:":              "1f473-200d-2640-fe0f",
	":woman-wearing-turban::skin-tone-2:": "1f473-1f3fb-200d-2640-fe0f",
	":woman-wearing-turban::skin-tone-3:": "1f473-1f3fc-200d-2640-fe0f",
	":woman-wearing-turban::skin-tone-4:": "1f473-1f3fd-200d-2640-fe0f",
	":woman-wearing-turban::skin-tone-5:": "1f473-1f3fe-200d-2640-fe0f",
	":woman-wearing-turban::skin-tone-6:": "1f473-1f3ff-200d-2640-fe0f",
	// 1F473-200D-2642-FE0F man-wearing-turban
	":man-wearing-turban:":              "1f473-200d-2642-fe0f",
	":man-wearing-turban::skin-tone-2:": "1f473-1f3fb-200d-2642-fe0f",
	":man-wearing-turban::skin-tone-3:": "1f473-1f3fc-200d-2642-fe0f",
	":man-wearing-turban::skin-tone-4:": "1f473-1f3fd-200d-2642-fe0f",
	":man-wearing-turban::skin-tone-5:": "1f473-1f3fe-200d-2642-fe0f",
	":man-wearing-turban::skin-tone-6:": "1f473-1f3ff-200d-2642-fe0f",
	// 1F473 man_with_turban
	":man_with_turban:":              "1f473",
	":man_with_turban::skin-tone-2:": "1f473-1f3fb",
	":man_with_turban::skin-tone-3:": "1f473-1f3fc",
	":man_with_turban::skin-tone-4:": "1f473-1f3fd",
	":man_with_turban::skin-tone-5:": "1f473-1f3fe",
	":man_with_turban::skin-tone-6:": "1f473-1f3ff",
	// 1F474 older_man
	":older_man:":              "1f474",
	":older_man::skin-tone-2:": "1f474-1f3fb",
	":older_man::skin-tone-3:": "1f474-1f3fc",
	":older_man::skin-tone-4:": "1f474-1f3fd",
	":older_man::skin-tone-5:": "1f474-1f3fe",
	":older_man::skin-tone-6:": "1f474-1f3ff",
	// 1F475 older_woman
	":older_woman:":              "1f475",
	":older_woman::skin-tone-2:": "1f475-1f3fb",
	":older_woman::skin-tone-3:": "1f475-1f3fc",
	":older_woman::skin-tone-4:": "1f475-1f3fd",
	":older_woman::skin-tone-5:": "1f475-1f3fe",
	":older_woman::skin-tone-6:": "1f475-1f3ff",
	// 1F476 baby
	":baby:":              "1f476",
	":baby::skin-tone-2:": "1f476-1f3fb",
	":baby::skin-tone-3:": "1f476-1f3fc",
	":baby::skin-tone-4:": "1f476-1f3fd",
	":baby::skin-tone-5:": "1f476-1f3fe",
	":baby::skin-tone-6:": "1f476-1f3ff",
	// 1F477-200D-2640-FE0F female-construction-worker
	":female-construction-worker:":              "1f477-200d-2640-fe0f",
	":female-construction-worker::skin-tone-2:": "1f477-1f3fb-200d-2640-fe0f",
	":female-construction-worker::skin-tone-3:": "1f477-1f3fc-200d-2640-fe0f",
	":female-construction-worker::skin-tone-4:": "1f477-1f3fd-200d-2640-fe0f",
	":female-construction-worker::skin-tone-5:": "1f477-1f3fe-200d-2640-fe0f",
	":female-construction-worker::skin-tone-6:": "1f477-1f3ff-200d-2640-fe0f",
	// 1F477-200D-2642-FE0F male-construction-worker
	":male-construction-worker:":              "1f477-200d-2642-fe0f",
	":male-construction-worker::skin-tone-2:": "1f477-1f3fb-200d-2642-fe0f",
	":male-construction-worker::skin-tone-3:": "1f477-1f3fc-200d-2642-fe0f",
	":male-construction-worker::skin-tone-4:": "1f477-1f3fd-200d-2642-fe0f",
	":male-construction-worker::skin-tone-5:": "1f477-1f3fe-200d-2642-fe0f",
	":male-construction-worker::skin-tone-6:": "1f477-1f3ff-200d-2642-fe0f",
	// 1F477 construction_worker
	":construction_worker:":              "1f477",
	":construction_worker::skin-tone-2:": "1f477-1f3fb",
	":construction_worker::skin-tone-3:": "1f477-1f3fc",
	":construction_worker::skin-tone-4:": "1f477-1f3fd",
	":construction_worker::skin-tone-5:": "1f477-1f3fe",
	":construction_worker::skin-tone-6:": "1f477-1f3ff",
	// 1F478 princess
	":princess:":              "1f478",
	":princess::skin-tone-2:": "1f478-1f3fb",
	":princess::skin-tone-3:": "1f478-1f3fc",
	":princess::skin-tone-4:": "1f478-1f3fd",
	":princess::skin-tone-5:": "1f478-1f3fe",
	":princess::skin-tone-6:": "1f478-1f3ff",
	// 1F479 japanese_ogre
	":japanese_ogre:": "1f479",
	// 1F47A japanese_goblin
	":japanese_goblin:": "1f47a",
	// 1F47B ghost
	":ghost:": "1f47b",
	// 1F47C angel
	":angel:":              "1f47c",
	":angel::skin-tone-2:": "1f47c-1f3fb",
	":angel::skin-tone-3:": "1f47c-1f3fc",
	":angel::skin-tone-4:": "1f47c-1f3fd",
	":angel::skin-tone-5:": "1f47c-1f3fe",
	":angel::skin-tone-6:": "1f47c-1f3ff",
	// 1F47D alien
	":alien:": "1f47d",
	// 1F47E space_invader
	":space_invader:": "1f47e",
	// 1F47F imp
	":imp:": "1f47f",
	// 1F480 skull
	":skull:": "1f480",
	// 1F481-200D-2640-FE0F woman-tipping-hand
	":woman-tipping-hand:":              "1f481-200d-2640-fe0f",
	":woman-tipping-hand::skin-tone-2:": "1f481-1f3fb-200d-2640-fe0f",
	":woman-tipping-hand::skin-tone-3:": "1f481-1f3fc-200d-2640-fe0f",
	":woman-tipping-hand::skin-tone-4:": "1f481-1f3fd-200d-2640-fe0f",
	":woman-tipping-hand::skin-tone-5:": "1f481-1f3fe-200d-2640-fe0f",
	":woman-tipping-hand::skin-tone-6:": "1f481-1f3ff-200d-2640-fe0f",
	// 1F481-200D-2642-FE0F man-tipping-hand
	":man-tipping-hand:":              "1f481-200d-2642-fe0f",
	":man-tipping-hand::skin-tone-2:": "1f481-1f3fb-200d-2642-fe0f",
	":man-tipping-hand::skin-tone-3:": "1f481-1f3fc-200d-2642-fe0f",
	":man-tipping-hand::skin-tone-4:": "1f481-1f3fd-200d-2642-fe0f",
	":man-tipping-hand::skin-tone-5:": "1f481-1f3fe-200d-2642-fe0f",
	":man-tipping-hand::skin-tone-6:": "1f481-1f3ff-200d-2642-fe0f",
	// 1F481 information_desk_person
	":information_desk_person:":              "1f481",
	":information_desk_person::skin-tone-2:": "1f481-1f3fb",
	":information_desk_person::skin-tone-3:": "1f481-1f3fc",
	":information_desk_person::skin-tone-4:": "1f481-1f3fd",
	":information_desk_person::skin-tone-5:": "1f481-1f3fe",
	":information_desk_person::skin-tone-6:": "1f481-1f3ff",
	// 1F482-200D-2640-FE0F female-guard
	":female-guard:":              "1f482-200d-2640-fe0f",
	":female-guard::skin-tone-2:": "1f482-1f3fb-200d-2640-fe0f",
	":female-guard::skin-tone-3:": "1f482-1f3fc-200d-2640-fe0f",
	":female-guard::skin-tone-4:": "1f482-1f3fd-200d-2640-fe0f",
	":female-guard::skin-tone-5:": "1f482-1f3fe-200d-2640-fe0f",
	":female-guard::skin-tone-6:": "1f482-1f3ff-200d-2640-fe0f",
	// 1F482-200D-2642-FE0F male-guard
	":male-guard:":              "1f482-200d-2642-fe0f",
	":male-guard::skin-tone-2:": "1f482-1f3fb-200d-2642-fe0f",
	":male-guard::skin-tone-3:": "1f482-1f3fc-200d-2642-fe0f",
	":male-guard::skin-tone-4:": "1f482-1f3fd-200d-2642-fe0f",
	":male-guard::skin-tone-5:": "1f482-1f3fe-200d-2642-fe0f",
	":male-guard::skin-tone-6:": "1f482-1f3ff-200d-2642-fe0f",
	// 1F482 guardsman
	":guardsman:":              "1f482",
	":guardsman::skin-tone-2:": "1f482-1f3fb",
	":guardsman::skin-tone-3:": "1f482-1f3fc",
	":guardsman::skin-tone-4:": "1f482-1f3fd",
	":guardsman::skin-tone-5:": "1f482-1f3fe",
	":guardsman::skin-tone-6:": "1f482-1f3ff",
	// 1F483 dancer
	":dancer:":              "1f483",
	":dancer::skin-tone-2:": "1f483-1f3fb",
	":dancer::skin-tone-3:": "1f483-1f3fc",
	":dancer::skin-tone-4:": "1f483-1f3fd",
	":dancer::skin-tone-5:": "1f483-1f3fe",
	":dancer::skin-tone-6:": "1f483-1f3ff",
	// 1F484 lipstick
	":lipstick:": "1f484",
	// 1F485 nail_care
	":nail_care:":              "1f485",
	":nail_care::skin-tone-2:": "1f485-1f3fb",
	":nail_care::skin-tone-3:": "1f485-1f3fc",
	":nail_care::skin-tone-4:": "1f485-1f3fd",
	":nail_care::skin-tone-5:": "1f485-1f3fe",
	":nail_care::skin-tone-6:": "1f485-1f3ff",
	// 1F486-200D-2640-FE0F woman-getting-massage
	":woman-getting-massage:":              "1f486-200d-2640-fe0f",
	":woman-getting-massage::skin-tone-2:": "1f486-1f3fb-200d-2640-fe0f",
	":woman-getting-massage::skin-tone-3:": "1f486-1f3fc-200d-2640-fe0f",
	":woman-getting-massage::skin-tone-4:": "1f486-1f3fd-200d-2640-fe0f",
	":woman-getting-massage::skin-tone-5:": "1f486-1f3fe-200d-2640-fe0f",
	":woman-getting-massage::skin-tone-6:": "1f486-1f3ff-200d-2640-fe0f",
	// 1F486-200D-2642-FE0F man-getting-massage
	":man-getting-massage:":              "1f486-200d-2642-fe0f",
	":man-getting-massage::skin-tone-2:": "1f486-1f3fb-200d-2642-fe0f",
	":man-getting-massage::skin-tone-3:": "1f486-1f3fc-200d-2642-fe0f",
	":man-getting-massage::skin-tone-4:": "1f486-1f3fd-200d-2642-fe0f",
	":man-getting-massage::skin-tone-5:": "1f486-1f3fe-200d-2642-fe0f",
	":man-getting-massage::skin-tone-6:": "1f486-1f3ff-200d-2642-fe0f",
	// 1F486 massage
	":massage:":              "1f486",
	":massage::skin-tone-2:": "1f486-1f3fb",
	":massage::skin-tone-3:": "1f486-1f3fc",
	":massage::skin-tone-4:": "1f486-1f3fd",
	":massage::skin-tone-5:": "1f486-1f3fe",
	":massage::skin-tone-6:": "1f486-1f3ff",
	// 1F487-200D-2640-FE0F woman-getting-haircut
	":woman-getting-haircut:":              "1f487-200d-2640-fe0f",
	":woman-getting-haircut::skin-tone-2:": "1f487-1f3fb-200d-2640-fe0f",
	":woman-getting-haircut::skin-tone-3:": "1f487-1f3fc-200d-2640-fe0f",
	":woman-getting-haircut::skin-tone-4:": "1f487-1f3fd-200d-2640-fe0f",
	":woman-getting-haircut::skin-tone-5:": "1f487-1f3fe-200d-2640-fe0f",
	":woman-getting-haircut::skin-tone-6:": "1f487-1f3ff-200d-2640-fe0f",
	// 1F487-200D-2642-FE0F man-getting-haircut
	":man-getting-haircut:":              "1f487-200d-2642-fe0f",
	":man-getting-haircut::skin-tone-2:": "1f487-1f3fb-200d-2642-fe0f",
	":man-getting-haircut::skin-tone-3:": "1f487-1f3fc-200d-2642-fe0f",
	":man-getting-haircut::skin-tone-4:": "1f487-1f3fd-200d-2642-fe0f",
	":man-getting-haircut::skin-tone-5:": "1f487-1f3fe-200d-2642-fe0f",
	":man-getting-haircut::skin-tone-6:": "1f487-1f3ff-200d-2642-fe0f",
	// 1F487 haircut
	":haircut:":              "1f487",
	":haircut::skin-tone-2:": "1f487-1f3fb",
	":haircut::skin-tone-3:": "1f487-1f3fc",
	":haircut::skin-tone-4:": "1f487-1f3fd",
	":haircut::skin-tone-5:": "1f487-1f3fe",
	":haircut::skin-tone-6:": "1f487-1f3ff",
	// 1F488 barber
	":barber:": "1f488",
	// 1F489 syringe
	":syringe:": "1f489",
	// 1F48A pill
	":pill:": "1f48a",
	// 1F48B kiss
	":kiss:": "1f48b",
	// 1F48C love_letter
	":love_letter:": "1f48c",
	// 1F48D ring
	":ring:": "1f48d",
	// 1F48E gem
	":gem:": "1f48e",
	// 1F48F couplekiss
	":couplekiss:":                "1f48f",
	":couplekiss::skin-tone-2:":   "1f48f-1f3fb",
	":couplekiss::skin-tone-3:":   "1f48f-1f3fc",
	":couplekiss::skin-tone-4:":   "1f48f-1f3fd",
	":couplekiss::skin-tone-5:":   "1f48f-1f3fe",
	":couplekiss::skin-tone-6:":   "1f48f-1f3ff",
	":couplekiss::skin-tone-2-3:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fc",
	":couplekiss::skin-tone-2-4:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fd",
	":couplekiss::skin-tone-2-5:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fe",
	":couplekiss::skin-tone-2-6:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3ff",
	":couplekiss::skin-tone-3-2:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fb",
	":couplekiss::skin-tone-3-4:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fd",
	":couplekiss::skin-tone-3-5:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fe",
	":couplekiss::skin-tone-3-6:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3ff",
	":couplekiss::skin-tone-4-2:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fb",
	":couplekiss::skin-tone-4-3:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fc",
	":couplekiss::skin-tone-4-5:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fe",
	":couplekiss::skin-tone-4-6:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3ff",
	":couplekiss::skin-tone-5-2:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fb",
	":couplekiss::skin-tone-5-3:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fc",
	":couplekiss::skin-tone-5-4:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fd",
	":couplekiss::skin-tone-5-6:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3ff",
	":couplekiss::skin-tone-6-2:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fb",
	":couplekiss::skin-tone-6-3:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fc",
	":couplekiss::skin-tone-6-4:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fd",
	":couplekiss::skin-tone-6-5:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f48b-200d-1f9d1-1f3fe",
	// 1F490 bouquet
	":bouquet:": "1f490",
	// 1F491 couple_with_heart
	":couple_with_heart:":                "1f491",
	":couple_with_heart::skin-tone-2:":   "1f491-1f3fb",
	":couple_with_heart::skin-tone-3:":   "1f491-1f3fc",
	":couple_with_heart::skin-tone-4:":   "1f491-1f3fd",
	":couple_with_heart::skin-tone-5:":   "1f491-1f3fe",
	":couple_with_heart::skin-tone-6:":   "1f491-1f3ff",
	":couple_with_heart::skin-tone-2-3:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f9d1-1f3fc",
	":couple_with_heart::skin-tone-2-4:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f9d1-1f3fd",
	":couple_with_heart::skin-tone-2-5:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f9d1-1f3fe",
	":couple_with_heart::skin-tone-2-6:": "1f9d1-1f3fb-200d-2764-fe0f-200d-1f9d1-1f3ff",
	":couple_with_heart::skin-tone-3-2:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f9d1-1f3fb",
	":couple_with_heart::skin-tone-3-4:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f9d1-1f3fd",
	":couple_with_heart::skin-tone-3-5:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f9d1-1f3fe",
	":couple_with_heart::skin-tone-3-6:": "1f9d1-1f3fc-200d-2764-fe0f-200d-1f9d1-1f3ff",
	":couple_with_heart::skin-tone-4-2:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f9d1-1f3fb",
	":couple_with_heart::skin-tone-4-3:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f9d1-1f3fc",
	":couple_with_heart::skin-tone-4-5:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f9d1-1f3fe",
	":couple_with_heart::skin-tone-4-6:": "1f9d1-1f3fd-200d-2764-fe0f-200d-1f9d1-1f3ff",
	":couple_with_heart::skin-tone-5-2:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f9d1-1f3fb",
	":couple_with_heart::skin-tone-5-3:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f9d1-1f3fc",
	":couple_with_heart::skin-tone-5-4:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f9d1-1f3fd",
	":couple_with_heart::skin-tone-5-6:": "1f9d1-1f3fe-200d-2764-fe0f-200d-1f9d1-1f3ff",
	":couple_with_heart::skin-tone-6-2:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f9d1-1f3fb",
	":couple_with_heart::skin-tone-6-3:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f9d1-1f3fc",
	":couple_with_heart::skin-tone-6-4:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f9d1-1f3fd",
	":couple_with_heart::skin-tone-6-5:": "1f9d1-1f3ff-200d-2764-fe0f-200d-1f9d1-1f3fe",
	// 1F492 wedding
	":wedding:": "1f492",
	// 1F493 heartbeat
	":heartbeat:": "1f493",
	// 1F494 broken_heart
	":broken_heart:": "1f494",
	// 1F495 two_hearts
	":two_hearts:": "1f495",
	// 1F496 sparkling_heart
	":sparkling_heart:": "1f496",
	// 1F497 heartpulse
	":heartpulse:": "1f497",
	// 1F498 cupid
	":cupid:": "1f498",
	// 1F499 blue_heart
	":blue_heart:": "1f499",
	// 1F49A green_heart
	":green_heart:": "1f49a",
	// 1F49B yellow_heart
	":yellow_heart:": "1f49b",
	// 1F49C purple_heart
	":purple_heart:": "1f49c",
	// 1F49D gift_heart
	":gift_heart:": "1f49d",
	// 1F49E revolving_hearts
	":revolving_hearts:": "1f49e",
	// 1F49F heart_decoration
	":heart_decoration:": "1f49f",
	// 1F4A0 diamond_shape_with_a_dot_inside
	":diamond_shape_with_a_dot_inside:": "1f4a0",
	// 1F4A1 bulb
	":bulb:": "1f4a1",
	// 1F4A2 anger
	":anger:": "1f4a2",
	// 1F4A3 bomb
	":bomb:": "1f4a3",
	// 1F4A4 zzz
	":zzz:": "1f4a4",
	// 1F4A5 boom
	":boom:": "1f4a5",
	// 1F4A5 boom
	":collision:": "1f4a5",
	// 1F4A6 sweat_drops
	":sweat_drops:": "1f4a6",
	// 1F4A7 droplet
	":droplet:": "1f4a7",
	// 1F4A8 dash
	":dash:": "1f4a8",
	// 1F4A9 hankey
	":hankey:": "1f4a9",
	// 1F4A9 hankey
	":poop:": "1f4a9",
	// 1F4A9 hankey
	":shit:": "1f4a9",
	// 1F4AA muscle
	":muscle:":              "1f4aa",
	":muscle::skin-tone-2:": "1f4aa-1f3fb",
	":muscle::skin-tone-3:": "1f4aa-1f3fc",
	":muscle::skin-tone-4:": "1f4aa-1f3fd",
	":muscle::skin-tone-5:": "1f4aa-1f3fe",
	":muscle::skin-tone-6:": "1f4aa-1f3ff",
	// 1F4AB dizzy
	":dizzy:": "1f4ab",
	// 1F4AC speech_balloon
	":speech_balloon:": "1f4ac",
	// 1F4AD thought_balloon
	":thought_balloon:": "1f4ad",
	// 1F4AE white_flower
	":white_flower:": "1f4ae",
	// 1F4AF 100
	":100:": "1f4af",
	// 1F4B0 moneybag
	":moneybag:": "1f4b0",
	// 1F4B1 currency_exchange
	":currency_exchange:": "1f4b1",
	// 1F4B2 heavy_dollar_sign
	":heavy_dollar_sign:": "1f4b2",
	// 1F4B3 credit_card
	":credit_card:": "1f4b3",
	// 1F4B4 yen
	":yen:": "1f4b4",
	// 1F4B5 dollar
	":dollar:": "1f4b5",
	// 1F4B6 euro
	":euro:": "1f4b6",
	// 1F4B7 pound
	":pound:": "1f4b7",
	// 1F4B8 money_with_wings
	":money_with_wings:": "1f4b8",
	// 1F4B9 chart
	":chart:": "1f4b9",
	// 1F4BA seat
	":seat:": "1f4ba",
	// 1F4BB computer
	":computer:": "1f4bb",
	// 1F4BC briefcase
	":briefcase:": "1f4bc",
	// 1F4BD minidisc
	":minidisc:": "1f4bd",
	// 1F4BE floppy_disk
	":floppy_disk:": "1f4be",
	// 1F4BF cd
	":cd:": "1f4bf",
	// 1F4C0 dvd
	":dvd:": "1f4c0",
	// 1F4C1 file_folder
	":file_folder:": "1f4c1",
	// 1F4C2 open_file_folder
	":open_file_folder:": "1f4c2",
	// 1F4C3 page_with_curl
	":page_with_curl:": "1f4c3",
	// 1F4C4 page_facing_up
	":page_facing_up:": "1f4c4",
	// 1F4C5 date
	":date:": "1f4c5",
	// 1F4C6 calendar
	":calendar:": "1f4c6",
	// 1F4C7 card_index
	":card_index:": "1f4c7",
	// 1F4C8 chart_with_upwards_trend
	":chart_with_upwards_trend:": "1f4c8",
	// 1F4C9 chart_with_downwards_trend
	":chart_with_downwards_trend:": "1f4c9",
	// 1F4CA bar_chart
	":bar_chart:": "1f4ca",
	// 1F4CB clipboard
	":clipboard:": "1f4cb",
	// 1F4CC pushpin
	":pushpin:": "1f4cc",
	// 1F4CD round_pushpin
	":round_pushpin:": "1f4cd",
	// 1F4CE paperclip
	":paperclip:": "1f4ce",
	// 1F4CF straight_ruler
	":straight_ruler:": "1f4cf",
	// 1F4D0 triangular_ruler
	":triangular_ruler:": "1f4d0",
	// 1F4D1 bookmark_tabs
	":bookmark_tabs:": "1f4d1",
	// 1F4D2 ledger
	":ledger:": "1f4d2",
	// 1F4D3 notebook
	":notebook:": "1f4d3",
	// 1F4D4 notebook_with_decorative_cover
	":notebook_with_decorative_cover:": "1f4d4",
	// 1F4D5 closed_book
	":closed_book:": "1f4d5",
	// 1F4D6 book
	":book:": "1f4d6",
	// 1F4D6 book
	":open_book:": "1f4d6",
	// 1F4D7 green_book
	":green_book:": "1f4d7",
	// 1F4D8 blue_book
	":blue_book:": "1f4d8",
	// 1F4D9 orange_book
	":orange_book:": "1f4d9",
	// 1F4DA books
	":books:": "1f4da",
	// 1F4DB name_badge
	":name_badge:": "1f4db",
	// 1F4DC scroll
	":scroll:": "1f4dc",
	// 1F4DD memo
	":memo:": "1f4dd",
	// 1F4DD memo
	":pencil:": "1f4dd",
	// 1F4DE telephone_receiver
	":telephone_receiver:": "1f4de",
	// 1F4DF pager
	":pager:": "1f4df",
	// 1F4E0 fax
	":fax:": "1f4e0",
	// 1F4E1 satellite_antenna
	":satellite_antenna:": "1f4e1",
	// 1F4E2 loudspeaker
	":loudspeaker:": "1f4e2",
	// 1F4E3 mega
	":mega:": "1f4e3",
	// 1F4E4 outbox_tray
	":outbox_tray:": "1f4e4",
	// 1F4E5 inbox_tray
	":inbox_tray:": "1f4e5",
	// 1F4E6 package
	":package:": "1f4e6",
	// 1F4E7 e-mail
	":e-mail:": "1f4e7",
	// 1F4E8 incoming_envelope
	":incoming_envelope:": "1f4e8",
	// 1F4E9 envelope_with_arrow
	":envelope_with_arrow:": "1f4e9",
	// 1F4EA mailbox_closed
	":mailbox_closed:": "1f4ea",
	// 1F4EB mailbox
	":mailbox:": "1f4eb",
	// 1F4EC mailbox_with_mail
	":mailbox_with_mail:": "1f4ec",
	// 1F4ED mailbox_with_no_mail
	":mailbox_with_no_mail:": "1f4ed",
	// 1F4EE postbox
	":postbox:": "1f4ee",
	// 1F4EF postal_horn
	":postal_horn:": "1f4ef",
	// 1F4F0 newspaper
	":newspaper:": "1f4f0",
	// 1F4F1 iphone
	":iphone:": "1f4f1",
	// 1F4F2 calling
	":calling:": "1f4f2",
	// 1F4F3 vibration_mode
	":vibration_mode:": "1f4f3",
	// 1F4F4 mobile_phone_off
	":mobile_phone_off:": "1f4f4",
	// 1F4F5 no_mobile_phones
	":no_mobile_phones:": "1f4f5",
	// 1F4F6 signal_strength
	":signal_strength:": "1f4f6",
	// 1F4F7 camera
	":camera:": "1f4f7",
	// 1F4F8 camera_with_flash
	":camera_with_flash:": "1f4f8",
	// 1F4F9 video_camera
	":video_camera:": "1f4f9",
	// 1F4FA tv
	":tv:": "1f4fa",
	// 1F4FB radio
	":radio:": "1f4fb",
	// 1F4FC vhs
	":vhs:": "1f4fc",
	// 1F4FD-FE0F film_projector
	":film_projector:": "1f4fd-fe0f",
	// 1F4FF prayer_beads
	":prayer_beads:": "1f4ff",
	// 1F500 twisted_rightwards_arrows
	":twisted_rightwards_arrows:": "1f500",
	// 1F501 repeat
	":repeat:": "1f501",
	// 1F502 repeat_one
	":repeat_one:": "1f502",
	// 1F503 arrows_clockwise
	":arrows_clockwise:": "1f503",
	// 1F504 arrows_counterclockwise
	":arrows_counterclockwise:": "1f504",
	// 1F505 low_brightness
	":low_brightness:": "1f505",
	// 1F506 high_brightness
	":high_brightness:": "1f506",
	// 1F507 mute
	":mute:": "1f507",
	// 1F508 speaker
	":speaker:": "1f508",
	// 1F509 sound
	":sound:": "1f509",
	// 1F50A loud_sound
	":loud_sound:": "1f50a",
	// 1F50B battery
	":battery:": "1f50b",
	// 1F50C electric_plug
	":electric_plug:": "1f50c",
	// 1F50D mag
	":mag:": "1f50d",
	// 1F50E mag_right
	":mag_right:": "1f50e",
	// 1F50F lock_with_ink_pen
	":lock_with_ink_pen:": "1f50f",
	// 1F510 closed_lock_with_key
	":closed_lock_with_key:": "1f510",
	// 1F511 key
	":key:": "1f511",
	// 1F512 lock
	":lock:": "1f512",
	// 1F513 unlock
	":unlock:": "1f513",
	// 1F514 bell
	":bell:": "1f514",
	// 1F515 no_bell
	":no_bell:": "1f515",
	// 1F516 bookmark
	":bookmark:": "1f516",
	// 1F517 link
	":link:": "1f517",
	// 1F518 radio_button
	":radio_button:": "1f518",
	// 1F519 back
	":back:": "1f519",
	// 1F51A end
	":end:": "1f51a",
	// 1F51B on
	":on:": "1f51b",
	// 1F51C soon
	":soon:": "1f51c",
	// 1F51D top
	":top:": "1f51d",
	// 1F51E underage
	":underage:": "1f51e",
	// 1F51F keycap_ten
	":keycap_ten:": "1f51f",
	// 1F520 capital_abcd
	":capital_abcd:": "1f520",
	// 1F521 abcd
	":abcd:": "1f521",
	// 1F522 1234
	":1234:": "1f522",
	// 1F523 symbols
	":symbols:": "1f523",
	// 1F524 abc
	":abc:": "1f524",
	// 1F525 fire
	":fire:": "1f525",
	// 1F526 flashlight
	":flashlight:": "1f526",
	// 1F527 wrench
	":wrench:": "1f527",
	// 1F528 hammer
	":hammer:": "1f528",
	// 1F529 nut_and_bolt
	":nut_and_bolt:": "1f529",
	// 1F52A hocho
	":hocho:": "1f52a",
	// 1F52A hocho
	":knife:": "1f52a",
	// 1F52B gun
	":gun:": "1f52b",
	// 1F52C microscope
	":microscope:": "1f52c",
	// 1F52D telescope
	":telescope:": "1f52d",
	// 1F52E crystal_ball
	":crystal_ball:": "1f52e",
	// 1F52F six_pointed_star
	":six_pointed_star:": "1f52f",
	// 1F530 beginner
	":beginner:": "1f530",
	// 1F531 trident
	":trident:": "1f531",
	// 1F532 black_square_button
	":black_square_button:": "1f532",
	// 1F533 white_square_button
	":white_square_button:": "1f533",
	// 1F534 red_circle
	":red_circle:": "1f534",
	// 1F535 large_blue_circle
	":large_blue_circle:": "1f535",
	// 1F536 large_orange_diamond
	":large_orange_diamond:": "1f536",
	// 1F537 large_blue_diamond
	":large_blue_diamond:": "1f537",
	// 1F538 small_orange_diamond
	":small_orange_diamond:": "1f538",
	// 1F539 small_blue_diamond
	":small_blue_diamond:": "1f539",
	// 1F53A small_red_triangle
	":small_red_triangle:": "1f53a",
	// 1F53B small_red_triangle_down
	":small_red_triangle_down:": "1f53b",
	// 1F53C arrow_up_small
	":arrow_up_small:": "1f53c",
	// 1F53D arrow_down_small
	":arrow_down_small:": "1f53d",
	// 1F549-FE0F om_symbol
	":om_symbol:": "1f549-fe0f",
	// 1F54A-FE0F dove_of_peace
	":dove_of_peace:": "1f54a-fe0f",
	// 1F54B kaaba
	":kaaba:": "1f54b",
	// 1F54C mosque
	":mosque:": "1f54c",
	// 1F54D synagogue
	":synagogue:": "1f54d",
	// 1F54E menorah_with_nine_branches
	":menorah_with_nine_branches:": "1f54e",
	// 1F550 clock1
	":clock1:": "1f550",
	// 1F551 clock2
	":clock2:": "1f551",
	// 1F552 clock3
	":clock3:": "1f552",
	// 1F553 clock4
	":clock4:": "1f553",
	// 1F554 clock5
	":clock5:": "1f554",
	// 1F555 clock6
	":clock6:": "1f555",
	// 1F556 clock7
	":clock7:": "1f556",
	// 1F557 clock8
	":clock8:": "1f557",
	// 1F558 clock9
	":clock9:": "1f558",
	// 1F559 clock10
	":clock10:": "1f559",
	// 1F55A clock11
	":clock11:": "1f55a",
	// 1F55B clock12
	":clock12:": "1f55b",
	// 1F55C clock130
	":clock130:": "1f55c",
	// 1F55D clock230
	":clock230:": "1f55d",
	// 1F55E clock330
	":clock330:": "1f55e",
	// 1F55F clock430
	":clock430:": "1f55f",
	// 1F560 clock530
	":clock530:": "1f560",
	// 1F561 clock630
	":clock630:": "1f561",
	// 1F562 clock730
	":clock730:": "1f562",
	// 1F563 clock830
	":clock830:": "1f563",
	// 1F564 clock930
	":clock930:": "1f564",
	// 1F565 clock1030
	":clock1030:": "1f565",
	// 1F566 clock1130
	":clock1130:": "1f566",
	// 1F567 clock1230
	":clock1230:": "1f567",
	// 1F56F-FE0F candle
	":candle:": "1f56f-fe0f",
	// 1F570-FE0F mantelpiece_clock
	":mantelpiece_clock:": "1f570-fe0f",
	// 1F573-FE0F hole
	":hole:": "1f573-fe0f",
	// 1F574-FE0F man_in_business_suit_levitating
	":man_in_business_suit_levitating:":              "1f574-fe0f",
	":man_in_business_suit_levitating::skin-tone-2:": "1f574-1f3fb",
	":man_in_business_suit_levitating::skin-tone-3:": "1f574-1f3fc",
	":man_in_business_suit_levitating::skin-tone-4:": "1f574-1f3fd",
	":man_in_business_suit_levitating::skin-tone-5:": "1f574-1f3fe",
	":man_in_business_suit_levitating::skin-tone-6:": "1f574-1f3ff",
	// 1F575-FE0F-200D-2640-FE0F female-detective
	":female-detective:":              "1f575-fe0f-200d-2640-fe0f",
	":female-detective::skin-tone-2:": "1f575-1f3fb-200d-2640-fe0f",
	":female-detective::skin-tone-3:": "1f575-1f3fc-200d-2640-fe0f",
	":female-detective::skin-tone-4:": "1f575-1f3fd-200d-2640-fe0f",
	":female-detective::skin-tone-5:": "1f575-1f3fe-200d-2640-fe0f",
	":female-detective::skin-tone-6:": "1f575-1f3ff-200d-2640-fe0f",
	// 1F575-FE0F-200D-2642-FE0F male-detective
	":male-detective:":              "1f575-fe0f-200d-2642-fe0f",
	":male-detective::skin-tone-2:": "1f575-1f3fb-200d-2642-fe0f",
	":male-detective::skin-tone-3:": "1f575-1f3fc-200d-2642-fe0f",
	":male-detective::skin-tone-4:": "1f575-1f3fd-200d-2642-fe0f",
	":male-detective::skin-tone-5:": "1f575-1f3fe-200d-2642-fe0f",
	":male-detective::skin-tone-6:": "1f575-1f3ff-200d-2642-fe0f",
	// 1F575-FE0F sleuth_or_spy
	":sleuth_or_spy:":              "1f575-fe0f",
	":sleuth_or_spy::skin-tone-2:": "1f575-1f3fb",
	":sleuth_or_spy::skin-tone-3:": "1f575-1f3fc",
	":sleuth_or_spy::skin-tone-4:": "1f575-1f3fd",
	":sleuth_or_spy::skin-tone-5:": "1f575-1f3fe",
	":sleuth_or_spy::skin-tone-6:": "1f575-1f3ff",
	// 1F576-FE0F dark_sunglasses
	":dark_sunglasses:": "1f576-fe0f",
	// 1F577-FE0F spider
	":spider:": "1f577-fe0f",
	// 1F578-FE0F spider_web
	":spider_web:": "1f578-fe0f",
	// 1F579-FE0F joystick
	":joystick:": "1f579-fe0f",
	// 1F57A man_dancing
	":man_dancing:":              "1f57a",
	":man_dancing::skin-tone-2:": "1f57a-1f3fb",
	":man_dancing::skin-tone-3:": "1f57a-1f3fc",
	":man_dancing::skin-tone-4:": "1f57a-1f3fd",
	":man_dancing::skin-tone-5:": "1f57a-1f3fe",
	":man_dancing::skin-tone-6:": "1f57a-1f3ff",
	// 1F587-FE0F linked_paperclips
	":linked_paperclips:": "1f587-fe0f",
	// 1F58A-FE0F lower_left_ballpoint_pen
	":lower_left_ballpoint_pen:": "1f58a-fe0f",
	// 1F58B-FE0F lower_left_fountain_pen
	":lower_left_fountain_pen:": "1f58b-fe0f",
	// 1F58C-FE0F lower_left_paintbrush
	":lower_left_paintbrush:": "1f58c-fe0f",
	// 1F58D-FE0F lower_left_crayon
	":lower_left_crayon:": "1f58d-fe0f",
	// 1F590-FE0F raised_hand_with_fingers_splayed
	":raised_hand_with_fingers_splayed:":              "1f590-fe0f",
	":raised_hand_with_fingers_splayed::skin-tone-2:": "1f590-1f3fb",
	":raised_hand_with_fingers_splayed::skin-tone-3:": "1f590-1f3fc",
	":raised_hand_with_fingers_splayed::skin-tone-4:": "1f590-1f3fd",
	":raised_hand_with_fingers_splayed::skin-tone-5:": "1f590-1f3fe",
	":raised_hand_with_fingers_splayed::skin-tone-6:": "1f590-1f3ff",
	// 1F595 middle_finger
	":middle_finger:": "1f595",
	// 1F595 middle_finger
	":reversed_hand_with_middle_finger_extended:":              "1f595",
	":middle_finger::skin-tone-2:":                             "1f595-1f3fb",
	":reversed_hand_with_middle_finger_extended::skin-tone-2:": "1f595-1f3fb",
	":middle_finger::skin-tone-3:":                             "1f595-1f3fc",
	":reversed_hand_with_middle_finger_extended::skin-tone-3:": "1f595-1f3fc",
	":middle_finger::skin-tone-4:":                             "1f595-1f3fd",
	":reversed_hand_with_middle_finger_extended::skin-tone-4:": "1f595-1f3fd",
	":middle_finger::skin-tone-5:":                             "1f595-1f3fe",
	":reversed_hand_with_middle_finger_extended::skin-tone-5:": "1f595-1f3fe",
	":middle_finger::skin-tone-6:":                             "1f595-1f3ff",
	":reversed_hand_with_middle_finger_extended::skin-tone-6:": "1f595-1f3ff",
	// 1F596 spock-hand
	":spock-hand:":              "1f596",
	":spock-hand::skin-tone-2:": "1f596-1f3fb",
	":spock-hand::skin-tone-3:": "1f596-1f3fc",
	":spock-hand::skin-tone-4:": "1f596-1f3fd",
	":spock-hand::skin-tone-5:": "1f596-1f3fe",
	":spock-hand::skin-tone-6:": "1f596-1f3ff",
	// 1F5A4 black_heart
	":black_heart:": "1f5a4",
	// 1F5A5-FE0F desktop_computer
	":desktop_computer:": "1f5a5-fe0f",
	// 1F5A8-FE0F printer
	":printer:": "1f5a8-fe0f",
	// 1F5B1-FE0F three_button_mouse
	":three_button_mouse:": "1f5b1-fe0f",
	// 1F5B2-FE0F trackball
	":trackball:": "1f5b2-fe0f",
	// 1F5BC-FE0F frame_with_picture
	":frame_with_picture:": "1f5bc-fe0f",
	// 1F5C2-FE0F card_index_dividers
	":card_index_dividers:": "1f5c2-fe0f",
	// 1F5C3-FE0F card_file_box
	":card_file_box:": "1f5c3-fe0f",
	// 1F5C4-FE0F file_cabinet
	":file_cabinet:": "1f5c4-fe0f",
	// 1F5D1-FE0F wastebasket
	":wastebasket:": "1f5d1-fe0f",
	// 1F5D2-FE0F spiral_note_pad
	":spiral_note_pad:": "1f5d2-fe0f",
	// 1F5D3-FE0F spiral_calendar_pad
	":spiral_calendar_pad:": "1f5d3-fe0f",
	// 1F5DC-FE0F compression
	":compression:": "1f5dc-fe0f",
	// 1F5DD-FE0F old_key
	":old_key:": "1f5dd-fe0f",
	// 1F5DE-FE0F rolled_up_newspaper
	":rolled_up_newspaper:": "1f5de-fe0f",
	// 1F5E1-FE0F dagger_knife
	":dagger_knife:": "1f5e1-fe0f",
	// 1F5E3-FE0F speaking_head_in_silhouette
	":speaking_head_in_silhouette:": "1f5e3-fe0f",
	// 1F5E8-FE0F left_speech_bubble
	":left_speech_bubble:": "1f5e8-fe0f",
	// 1F5EF-FE0F right_anger_bubble
	":right_anger_bubble:": "1f5ef-fe0f",
	// 1F5F3-FE0F ballot_box_with_ballot
	":ballot_box_with_ballot:": "1f5f3-fe0f",
	// 1F5FA-FE0F world_map
	":world_map:": "1f5fa-fe0f",
	// 1F5FB mount_fuji
	":mount_fuji:": "1f5fb",
	// 1F5FC tokyo_tower
	":tokyo_tower:": "1f5fc",
	// 1F5FD statue_of_liberty
	":statue_of_liberty:": "1f5fd",
	// 1F5FE japan
	":japan:": "1f5fe",
	// 1F5FF moyai
	":moyai:": "1f5ff",
	// 1F600 grinning
	":grinning:": "1f600",
	// 1F601 grin
	":grin:": "1f601",
	// 1F602 joy
	":joy:": "1f602",
	// 1F603 smiley
	":smiley:": "1f603",
	// 1F604 smile
	":smile:": "1f604",
	// 1F605 sweat_smile
	":sweat_smile:": "1f605",
	// 1F606 laughing
	":laughing:": "1f606",
	// 1F606 laughing
	":satisfied:": "1f606",
	// 1F607 innocent
	":innocent:": "1f607",
	// 1F608 smiling_imp
	":smiling_imp:": "1f608",
	// 1F609 wink
	":wink:": "1f609",
	// 1F60A blush
	":blush:": "1f60a",
	// 1F60B yum
	":yum:": "1f60b",
	// 1F60C relieved
	":relieved:": "1f60c",
	// 1F60D heart_eyes
	":heart_eyes:": "1f60d",
	// 1F60E sunglasses
	":sunglasses:": "1f60e",
	// 1F60F smirk
	":smirk:": "1f60f",
	// 1F610 neutral_face
	":neutral_face:": "1f610",
	// 1F611 expressionless
	":expressionless:": "1f611",
	// 1F612 unamused
	":unamused:": "1f612",
	// 1F613 sweat
	":sweat:": "1f613",
	// 1F614 pensive
	":pensive:": "1f614",
	// 1F615 confused
	":confused:": "1f615",
	// 1F616 confounded
	":confounded:": "1f616",
	// 1F617 kissing
	":kissing:": "1f617",
	// 1F618 kissing_heart
	":kissing_heart:": "1f618",
	// 1F619 kissing_smiling_eyes
	":kissing_smiling_eyes:": "1f619",
	// 1F61A kissing_closed_eyes
	":kissing_closed_eyes:": "1f61a",
	// 1F61B stuck_out_tongue
	":stuck_out_tongue:": "1f61b",
	// 1F61C stuck_out_tongue_winking_eye
	":stuck_out_tongue_winking_eye:": "1f61c",
	// 1F61D stuck_out_tongue_closed_eyes
	":stuck_out_tongue_closed_eyes:": "1f61d",
	// 1F61E disappointed
	":disappointed:": "1f61e",
	// 1F61F worried
	":worried:": "1f61f",
	// 1F620 angry
	":angry:": "1f620",
	// 1F621 rage
	":rage:": "1f621",
	// 1F622 cry
	":cry:": "1f622",
	// 1F623 persevere
	":persevere:": "1f623",
	// 1F624 triumph
	":triumph:": "1f624",
	// 1F625 disappointed_relieved
	":disappointed_relieved:": "1f625",
	// 1F626 frowning
	":frowning:": "1f626",
	// 1F627 anguished
	":anguished:": "1f627",
	// 1F628 fearful
	":fearful:": "1f628",
	// 1F629 weary
	":weary:": "1f629",
	// 1F62A sleepy
	":sleepy:": "1f62a",
	// 1F62B tired_face
	":tired_face:": "1f62b",
	// 1F62C grimacing
	":grimacing:": "1f62c",
	// 1F62D sob
	":sob:": "1f62d",
	// 1F62E-200D-1F4A8 face_exhaling
	":face_exhaling:": "1f62e-200d-1f4a8",
	// 1F62E open_mouth
	":open_mouth:": "1f62e",
	// 1F62F hushed
	":hushed:": "1f62f",
	// 1F630 cold_sweat
	":cold_sweat:": "1f630",
	// 1F631 scream
	":scream:": "1f631",
	// 1F632 astonished
	":astonished:": "1f632",
	// 1F633 flushed
	":flushed:": "1f633",
	// 1F634 sleeping
	":sleeping:": "1f634",
	// 1F635-200D-1F4AB face_with_spiral_eyes
	":face_with_spiral_eyes:": "1f635-200d-1f4ab",
	// 1F635 dizzy_face
	":dizzy_face:": "1f635",
	// 1F636-200D-1F32B-FE0F face_in_clouds
	":face_in_clouds:": "1f636-200d-1f32b-fe0f",
	// 1F636 no_mouth
	":no_mouth:": "1f636",
	// 1F637 mask
	":mask:": "1f637",
	// 1F638 smile_cat
	":smile_cat:": "1f638",
	// 1F639 joy_cat
	":joy_cat:": "1f639",
	// 1F63A smiley_cat
	":smiley_cat:": "1f63a",
	// 1F63B heart_eyes_cat
	":heart_eyes_cat:": "1f63b",
	// 1F63C smirk_cat
	":smirk_cat:": "1f63c",
	// 1F63D kissing_cat
	":kissing_cat:": "1f63d",
	// 1F63E pouting_cat
	":pouting_cat:": "1f63e",
	// 1F63F crying_cat_face
	":crying_cat_face:": "1f63f",
	// 1F640 scream_cat
	":scream_cat:": "1f640",
	// 1F641 slightly_frowning_face
	":slightly_frowning_face:": "1f641",
	// 1F642 slightly_smiling_face
	":slightly_smiling_face:": "1f642",
	// 1F643 upside_down_face
	":upside_down_face:": "1f643",
	// 1F644 face_with_rolling_eyes
	":face_with_rolling_eyes:": "1f644",
	// 1F645-200D-2640-FE0F woman-gesturing-no
	":woman-gesturing-no:":              "1f645-200d-2640-fe0f",
	":woman-gesturing-no::skin-tone-2:": "1f645-1f3fb-200d-2640-fe0f",
	":woman-gesturing-no::skin-tone-3:": "1f645-1f3fc-200d-2640-fe0f",
	":woman-gesturing-no::skin-tone-4:": "1f645-1f3fd-200d-2640-fe0f",
	":woman-gesturing-no::skin-tone-5:": "1f645-1f3fe-200d-2640-fe0f",
	":woman-gesturing-no::skin-tone-6:": "1f645-1f3ff-200d-2640-fe0f",
	// 1F645-200D-2642-FE0F man-gesturing-no
	":man-gesturing-no:":              "1f645-200d-2642-fe0f",
	":man-gesturing-no::skin-tone-2:": "1f645-1f3fb-200d-2642-fe0f",
	":man-gesturing-no::skin-tone-3:": "1f645-1f3fc-200d-2642-fe0f",
	":man-gesturing-no::skin-tone-4:": "1f645-1f3fd-200d-2642-fe0f",
	":man-gesturing-no::skin-tone-5:": "1f645-1f3fe-200d-2642-fe0f",
	":man-gesturing-no::skin-tone-6:": "1f645-1f3ff-200d-2642-fe0f",
	// 1F645 no_good
	":no_good:":              "1f645",
	":no_good::skin-tone-2:": "1f645-1f3fb",
	":no_good::skin-tone-3:": "1f645-1f3fc",
	":no_good::skin-tone-4:": "1f645-1f3fd",
	":no_good::skin-tone-5:": "1f645-1f3fe",
	":no_good::skin-tone-6:": "1f645-1f3ff",
	// 1F646-200D-2640-FE0F woman-gesturing-ok
	":woman-gesturing-ok:":              "1f646-200d-2640-fe0f",
	":woman-gesturing-ok::skin-tone-2:": "1f646-1f3fb-200d-2640-fe0f",
	":woman-gesturing-ok::skin-tone-3:": "1f646-1f3fc-200d-2640-fe0f",
	":woman-gesturing-ok::skin-tone-4:": "1f646-1f3fd-200d-2640-fe0f",
	":woman-gesturing-ok::skin-tone-5:": "1f646-1f3fe-200d-2640-fe0f",
	":woman-gesturing-ok::skin-tone-6:": "1f646-1f3ff-200d-2640-fe0f",
	// 1F646-200D-2642-FE0F man-gesturing-ok
	":man-gesturing-ok:":              "1f646-200d-2642-fe0f",
	":man-gesturing-ok::skin-tone-2:": "1f646-1f3fb-200d-2642-fe0f",
	":man-gesturing-ok::skin-tone-3:": "1f646-1f3fc-200d-2642-fe0f",
	":man-gesturing-ok::skin-tone-4:": "1f646-1f3fd-200d-2642-fe0f",
	":man-gesturing-ok::skin-tone-5:": "1f646-1f3fe-200d-2642-fe0f",
	":man-gesturing-ok::skin-tone-6:": "1f646-1f3ff-200d-2642-fe0f",
	// 1F646 ok_woman
	":ok_woman:":              "1f646",
	":ok_woman::skin-tone-2:": "1f646-1f3fb",
	":ok_woman::skin-tone-3:": "1f646-1f3fc",
	":ok_woman::skin-tone-4:": "1f646-1f3fd",
	":ok_woman::skin-tone-5:": "1f646-1f3fe",
	":ok_woman::skin-tone-6:": "1f646-1f3ff",
	// 1F647-200D-2640-FE0F woman-bowing
	":woman-bowing:":              "1f647-200d-2640-fe0f",
	":woman-bowing::skin-tone-2:": "1f647-1f3fb-200d-2640-fe0f",
	":woman-bowing::skin-tone-3:": "1f647-1f3fc-200d-2640-fe0f",
	":woman-bowing::skin-tone-4:": "1f647-1f3fd-200d-2640-fe0f",
	":woman-bowing::skin-tone-5:": "1f647-1f3fe-200d-2640-fe0f",
	":woman-bowing::skin-tone-6:": "1f647-1f3ff-200d-2640-fe0f",
	// 1F647-200D-2642-FE0F man-bowing
	":man-bowing:":              "1f647-200d-2642-fe0f",
	":man-bowing::skin-tone-2:": "1f647-1f3fb-200d-2642-fe0f",
	":man-bowing::skin-tone-3:": "1f647-1f3fc-200d-2642-fe0f",
	":man-bowing::skin-tone-4:": "1f647-1f3fd-200d-2642-fe0f",
	":man-bowing::skin-tone-5:": "1f647-1f3fe-200d-2642-fe0f",
	":man-bowing::skin-tone-6:": "1f647-1f3ff-200d-2642-fe0f",
	// 1F647 bow
	":bow:":              "1f647",
	":bow::skin-tone-2:": "1f647-1f3fb",
	":bow::skin-tone-3:": "1f647-1f3fc",
	":bow::skin-tone-4:": "1f647-1f3fd",
	":bow::skin-tone-5:": "1f647-1f3fe",
	":bow::skin-tone-6:": "1f647-1f3ff",
	// 1F648 see_no_evil
	":see_no_evil:": "1f648",
	// 1F649 hear_no_evil
	":hear_no_evil:": "1f649",
	// 1F64A speak_no_evil
	":speak_no_evil:": "1f64a",
	// 1F64B-200D-2640-FE0F woman-raising-hand
	":woman-raising-hand:":              "1f64b-200d-2640-fe0f",
	":woman-raising-hand::skin-tone-2:": "1f64b-1f3fb-200d-2640-fe0f",
	":woman-raising-hand::skin-tone-3:": "1f64b-1f3fc-200d-2640-fe0f",
	":woman-raising-hand::skin-tone-4:": "1f64b-1f3fd-200d-2640-fe0f",
	":woman-raising-hand::skin-tone-5:": "1f64b-1f3fe-200d-2640-fe0f",
	":woman-raising-hand::skin-tone-6:": "1f64b-1f3ff-200d-2640-fe0f",
	// 1F64B-200D-2642-FE0F man-raising-hand
	":man-raising-hand:":              "1f64b-200d-2642-fe0f",
	":man-raising-hand::skin-tone-2:": "1f64b-1f3fb-200d-2642-fe0f",
	":man-raising-hand::skin-tone-3:": "1f64b-1f3fc-200d-2642-fe0f",
	":man-raising-hand::skin-tone-4:": "1f64b-1f3fd-200d-2642-fe0f",
	":man-raising-hand::skin-tone-5:": "1f64b-1f3fe-200d-2642-fe0f",
	":man-raising-hand::skin-tone-6:": "1f64b-1f3ff-200d-2642-fe0f",
	// 1F64B raising_hand
	":raising_hand:":              "1f64b",
	":raising_hand::skin-tone-2:": "1f64b-1f3fb",
	":raising_hand::skin-tone-3:": "1f64b-1f3fc",
	":raising_hand::skin-tone-4:": "1f64b-1f3fd",
	":raising_hand::skin-tone-5:": "1f64b-1f3fe",
	":raising_hand::skin-tone-6:": "1f64b-1f3ff",
	// 1F64C raised_hands
	":raised_hands:":              "1f64c",
	":raised_hands::skin-tone-2:": "1f64c-1f3fb",
	":raised_hands::skin-tone-3:": "1f64c-1f3fc",
	":raised_hands::skin-tone-4:": "1f64c-1f3fd",
	":raised_hands::skin-tone-5:": "1f64c-1f3fe",
	":raised_hands::skin-tone-6:": "1f64c-1f3ff",
	// 1F64D-200D-2640-FE0F woman-frowning
	":woman-frowning:":              "1f64d-200d-2640-fe0f",
	":woman-frowning::skin-tone-2:": "1f64d-1f3fb-200d-2640-fe0f",
	":woman-frowning::skin-tone-3:": "1f64d-1f3fc-200d-2640-fe0f",
	":woman-frowning::skin-tone-4:": "1f64d-1f3fd-200d-2640-fe0f",
	":woman-frowning::skin-tone-5:": "1f64d-1f3fe-200d-2640-fe0f",
	":woman-frowning::skin-tone-6:": "1f64d-1f3ff-200d-2640-fe0f",
	// 1F64D-200D-2642-FE0F man-frowning
	":man-frowning:":              "1f64d-200d-2642-fe0f",
	":man-frowning::skin-tone-2:": "1f64d-1f3fb-200d-2642-fe0f",
	":man-frowning::skin-tone-3:": "1f64d-1f3fc-200d-2642-fe0f",
	":man-frowning::skin-tone-4:": "1f64d-1f3fd-200d-2642-fe0f",
	":man-frowning::skin-tone-5:": "1f64d-1f3fe-200d-2642-fe0f",
	":man-frowning::skin-tone-6:": "1f64d-1f3ff-200d-2642-fe0f",
	// 1F64D person_frowning
	":person_frowning:":              "1f64d",
	":person_frowning::skin-tone-2:": "1f64d-1f3fb",
	":person_frowning::skin-tone-3:": "1f64d-1f3fc",
	":person_frowning::skin-tone-4:": "1f64d-1f3fd",
	":person_frowning::skin-tone-5:": "1f64d-1f3fe",
	":person_frowning::skin-tone-6:": "1f64d-1f3ff",
	// 1F64E-200D-2640-FE0F woman-pouting
	":woman-pouting:":              "1f64e-200d-2640-fe0f",
	":woman-pouting::skin-tone-2:": "1f64e-1f3fb-200d-2640-fe0f",
	":woman-pouting::skin-tone-3:": "1f64e-1f3fc-200d-2640-fe0f",
	":woman-pouting::skin-tone-4:": "1f64e-1f3fd-200d-2640-fe0f",
	":woman-pouting::skin-tone-5:": "1f64e-1f3fe-200d-2640-fe0f",
	":woman-pouting::skin-tone-6:": "1f64e-1f3ff-200d-2640-fe0f",
	// 1F64E-200D-2642-FE0F man-pouting
	":man-pouting:":              "1f64e-200d-2642-fe0f",
	":man-pouting::skin-tone-2:": "1f64e-1f3fb-200d-2642-fe0f",
	":man-pouting::skin-tone-3:": "1f64e-1f3fc-200d-2642-fe0f",
	":man-pouting::skin-tone-4:": "1f64e-1f3fd-200d-2642-fe0f",
	":man-pouting::skin-tone-5:": "1f64e-1f3fe-200d-2642-fe0f",
	":man-pouting::skin-tone-6:": "1f64e-1f3ff-200d-2642-fe0f",
	// 1F64E person_with_pouting_face
	":person_with_pouting_face:":              "1f64e",
	":person_with_pouting_face::skin-tone-2:": "1f64e-1f3fb",
	":person_with_pouting_face::skin-tone-3:": "1f64e-1f3fc",
	":person_with_pouting_face::skin-tone-4:": "1f64e-1f3fd",
	":person_with_pouting_face::skin-tone-5:": "1f64e-1f3fe",
	":person_with_pouting_face::skin-tone-6:": "1f64e-1f3ff",
	// 1F64F pray
	":pray:":              "1f64f",
	":pray::skin-tone-2:": "1f64f-1f3fb",
	":pray::skin-tone-3:": "1f64f-1f3fc",
	":pray::skin-tone-4:": "1f64f-1f3fd",
	":pray::skin-tone-5:": "1f64f-1f3fe",
	":pray::skin-tone-6:": "1f64f-1f3ff",
	// 1F680 rocket
	":rocket:": "1f680",
	// 1F681 helicopter
	":helicopter:": "1f681",
	// 1F682 steam_locomotive
	":steam_locomotive:": "1f682",
	// 1F683 railway_car
	":railway_car:": "1f683",
	// 1F684 bullettrain_side
	":bullettrain_side:": "1f684",
	// 1F685 bullettrain_front
	":bullettrain_front:": "1f685",
	// 1F686 train2
	":train2:": "1f686",
	// 1F687 metro
	":metro:": "1f687",
	// 1F688 light_rail
	":light_rail:": "1f688",
	// 1F689 station
	":station:": "1f689",
	// 1F68A tram
	":tram:": "1f68a",
	// 1F68B train
	":train:": "1f68b",
	// 1F68C bus
	":bus:": "1f68c",
	// 1F68D oncoming_bus
	":oncoming_bus:": "1f68d",
	// 1F68E trolleybus
	":trolleybus:": "1f68e",
	// 1F68F busstop
	":busstop:": "1f68f",
	// 1F690 minibus
	":minibus:": "1f690",
	// 1F691 ambulance
	":ambulance:": "1f691",
	// 1F692 fire_engine
	":fire_engine:": "1f692",
	// 1F693 police_car
	":police_car:": "1f693",
	// 1F694 oncoming_police_car
	":oncoming_police_car:": "1f694",
	// 1F695 taxi
	":taxi:": "1f695",
	// 1F696 oncoming_taxi
	":oncoming_taxi:": "1f696",
	// 1F697 car
	":car:": "1f697",
	// 1F697 car
	":red_car:": "1f697",
	// 1F698 oncoming_automobile
	":oncoming_automobile:": "1f698",
	// 1F699 blue_car
	":blue_car:": "1f699",
	// 1F69A truck
	":truck:": "1f69a",
	// 1F69B articulated_lorry
	":articulated_lorry:": "1f69b",
	// 1F69C tractor
	":tractor:": "1f69c",
	// 1F69D monorail
	":monorail:": "1f69d",
	// 1F69E mountain_railway
	":mountain_railway:": "1f69e",
	// 1F69F suspension_railway
	":suspension_railway:": "1f69f",
	// 1F6A0 mountain_cableway
	":mountain_cableway:": "1f6a0",
	// 1F6A1 aerial_tramway
	":aerial_tramway:": "1f6a1",
	// 1F6A2 ship
	":ship:": "1f6a2",
	// 1F6A3-200D-2640-FE0F woman-rowing-boat
	":woman-rowing-boat:":              "1f6a3-200d-2640-fe0f",
	":woman-rowing-boat::skin-tone-2:": "1f6a3-1f3fb-200d-2640-fe0f",
	":woman-rowing-boat::skin-tone-3:": "1f6a3-1f3fc-200d-2640-fe0f",
	":woman-rowing-boat::skin-tone-4:": "1f6a3-1f3fd-200d-2640-fe0f",
	":woman-rowing-boat::skin-tone-5:": "1f6a3-1f3fe-200d-2640-fe0f",
	":woman-rowing-boat::skin-tone-6:": "1f6a3-1f3ff-200d-2640-fe0f",
	// 1F6A3-200D-2642-FE0F man-rowing-boat
	":man-rowing-boat:":              "1f6a3-200d-2642-fe0f",
	":man-rowing-boat::skin-tone-2:": "1f6a3-1f3fb-200d-2642-fe0f",
	":man-rowing-boat::skin-tone-3:": "1f6a3-1f3fc-200d-2642-fe0f",
	":man-rowing-boat::skin-tone-4:": "1f6a3-1f3fd-200d-2642-fe0f",
	":man-rowing-boat::skin-tone-5:": "1f6a3-1f3fe-200d-2642-fe0f",
	":man-rowing-boat::skin-tone-6:": "1f6a3-1f3ff-200d-2642-fe0f",
	// 1F6A3 rowboat
	":rowboat:":              "1f6a3",
	":rowboat::skin-tone-2:": "1f6a3-1f3fb",
	":rowboat::skin-tone-3:": "1f6a3-1f3fc",
	":rowboat::skin-tone-4:": "1f6a3-1f3fd",
	":rowboat::skin-tone-5:": "1f6a3-1f3fe",
	":rowboat::skin-tone-6:": "1f6a3-1f3ff",
	// 1F6A4 speedboat
	":speedboat:": "1f6a4",
	// 1F6A5 traffic_light
	":traffic_light:": "1f6a5",
	// 1F6A6 vertical_traffic_light
	":vertical_traffic_light:": "1f6a6",
	// 1F6A7 construction
	":construction:": "1f6a7",
	// 1F6A8 rotating_light
	":rotating_light:": "1f6a8",
	// 1F6A9 triangular_flag_on_post
	":triangular_flag_on_post:": "1f6a9",
	// 1F6AA door
	":door:": "1f6aa",
	// 1F6AB no_entry_sign
	":no_entry_sign:": "1f6ab",
	// 1F6AC smoking
	":smoking:": "1f6ac",
	// 1F6AD no_smoking
	":no_smoking:": "1f6ad",
	// 1F6AE put_litter_in_its_place
	":put_litter_in_its_place:": "1f6ae",
	// 1F6AF do_not_litter
	":do_not_litter:": "1f6af",
	// 1F6B0 potable_water
	":potable_water:": "1f6b0",
	// 1F6B1 non-potable_water
	":non-potable_water:": "1f6b1",
	// 1F6B2 bike
	":bike:": "1f6b2",
	// 1F6B3 no_bicycles
	":no_bicycles:": "1f6b3",
	// 1F6B4-200D-2640-FE0F woman-biking
	":woman-biking:":              "1f6b4-200d-2640-fe0f",
	":woman-biking::skin-tone-2:": "1f6b4-1f3fb-200d-2640-fe0f",
	":woman-biking::skin-tone-3:": "1f6b4-1f3fc-200d-2640-fe0f",
	":woman-biking::skin-tone-4:": "1f6b4-1f3fd-200d-2640-fe0f",
	":woman-biking::skin-tone-5:": "1f6b4-1f3fe-200d-2640-fe0f",
	":woman-biking::skin-tone-6:": "1f6b4-1f3ff-200d-2640-fe0f",
	// 1F6B4-200D-2642-FE0F man-biking
	":man-biking:":              "1f6b4-200d-2642-fe0f",
	":man-biking::skin-tone-2:": "1f6b4-1f3fb-200d-2642-fe0f",
	":man-biking::skin-tone-3:": "1f6b4-1f3fc-200d-2642-fe0f",
	":man-biking::skin-tone-4:": "1f6b4-1f3fd-200d-2642-fe0f",
	":man-biking::skin-tone-5:": "1f6b4-1f3fe-200d-2642-fe0f",
	":man-biking::skin-tone-6:": "1f6b4-1f3ff-200d-2642-fe0f",
	// 1F6B4 bicyclist
	":bicyclist:":              "1f6b4",
	":bicyclist::skin-tone-2:": "1f6b4-1f3fb",
	":bicyclist::skin-tone-3:": "1f6b4-1f3fc",
	":bicyclist::skin-tone-4:": "1f6b4-1f3fd",
	":bicyclist::skin-tone-5:": "1f6b4-1f3fe",
	":bicyclist::skin-tone-6:": "1f6b4-1f3ff",
	// 1F6B5-200D-2640-FE0F woman-mountain-biking
	":woman-mountain-biking:":              "1f6b5-200d-2640-fe0f",
	":woman-mountain-biking::skin-tone-2:": "1f6b5-1f3fb-200d-2640-fe0f",
	":woman-mountain-biking::skin-tone-3:": "1f6b5-1f3fc-200d-2640-fe0f",
	":woman-mountain-biking::skin-tone-4:": "1f6b5-1f3fd-200d-2640-fe0f",
	":woman-mountain-biking::skin-tone-5:": "1f6b5-1f3fe-200d-2640-fe0f",
	":woman-mountain-biking::skin-tone-6:": "1f6b5-1f3ff-200d-2640-fe0f",
	// 1F6B5-200D-2642-FE0F man-mountain-biking
	":man-mountain-biking:":              "1f6b5-200d-2642-fe0f",
	":man-mountain-biking::skin-tone-2:": "1f6b5-1f3fb-200d-2642-fe0f",
	":man-mountain-biking::skin-tone-3:": "1f6b5-1f3fc-200d-2642-fe0f",
	":man-mountain-biking::skin-tone-4:": "1f6b5-1f3fd-200d-2642-fe0f",
	":man-mountain-biking::skin-tone-5:": "1f6b5-1f3fe-200d-2642-fe0f",
	":man-mountain-biking::skin-tone-6:": "1f6b5-1f3ff-200d-2642-fe0f",
	// 1F6B5 mountain_bicyclist
	":mountain_bicyclist:":              "1f6b5",
	":mountain_bicyclist::skin-tone-2:": "1f6b5-1f3fb",
	":mountain_bicyclist::skin-tone-3:": "1f6b5-1f3fc",
	":mountain_bicyclist::skin-tone-4:": "1f6b5-1f3fd",
	":mountain_bicyclist::skin-tone-5:": "1f6b5-1f3fe",
	":mountain_bicyclist::skin-tone-6:": "1f6b5-1f3ff",
	// 1F6B6-200D-2640-FE0F woman-walking
	":woman-walking:":              "1f6b6-200d-2640-fe0f",
	":woman-walking::skin-tone-2:": "1f6b6-1f3fb-200d-2640-fe0f",
	":woman-walking::skin-tone-3:": "1f6b6-1f3fc-200d-2640-fe0f",
	":woman-walking::skin-tone-4:": "1f6b6-1f3fd-200d-2640-fe0f",
	":woman-walking::skin-tone-5:": "1f6b6-1f3fe-200d-2640-fe0f",
	":woman-walking::skin-tone-6:": "1f6b6-1f3ff-200d-2640-fe0f",
	// 1F6B6-200D-2642-FE0F man-walking
	":man-walking:":              "1f6b6-200d-2642-fe0f",
	":man-walking::skin-tone-2:": "1f6b6-1f3fb-200d-2642-fe0f",
	":man-walking::skin-tone-3:": "1f6b6-1f3fc-200d-2642-fe0f",
	":man-walking::skin-tone-4:": "1f6b6-1f3fd-200d-2642-fe0f",
	":man-walking::skin-tone-5:": "1f6b6-1f3fe-200d-2642-fe0f",
	":man-walking::skin-tone-6:": "1f6b6-1f3ff-200d-2642-fe0f",
	// 1F6B6 walking
	":walking:":              "1f6b6",
	":walking::skin-tone-2:": "1f6b6-1f3fb",
	":walking::skin-tone-3:": "1f6b6-1f3fc",
	":walking::skin-tone-4:": "1f6b6-1f3fd",
	":walking::skin-tone-5:": "1f6b6-1f3fe",
	":walking::skin-tone-6:": "1f6b6-1f3ff",
	// 1F6B7 no_pedestrians
	":no_pedestrians:": "1f6b7",
	// 1F6B8 children_crossing
	":children_crossing:": "1f6b8",
	// 1F6B9 mens
	":mens:": "1f6b9",
	// 1F6BA womens
	":womens:": "1f6ba",
	// 1F6BB restroom
	":restroom:": "1f6bb",
	// 1F6BC baby_symbol
	":baby_symbol:": "1f6bc",
	// 1F6BD toilet
	":toilet:": "1f6bd",
	// 1F6BE wc
	":wc:": "1f6be",
	// 1F6BF shower
	":shower:": "1f6bf",
	// 1F6C0 bath
	":bath:":              "1f6c0",
	":bath::skin-tone-2:": "1f6c0-1f3fb",
	":bath::skin-tone-3:": "1f6c0-1f3fc",
	":bath::skin-tone-4:": "1f6c0-1f3fd",
	":bath::skin-tone-5:": "1f6c0-1f3fe",
	":bath::skin-tone-6:": "1f6c0-1f3ff",
	// 1F6C1 bathtub
	":bathtub:": "1f6c1",
	// 1F6C2 passport_control
	":passport_control:": "1f6c2",
	// 1F6C3 customs
	":customs:": "1f6c3",
	// 1F6C4 baggage_claim
	":baggage_claim:": "1f6c4",
	// 1F6C5 left_luggage
	":left_luggage:": "1f6c5",
	// 1F6CB-FE0F couch_and_lamp
	":couch_and_lamp:": "1f6cb-fe0f",
	// 1F6CC sleeping_accommodation
	":sleeping_accommodation:":              "1f6cc",
	":sleeping_accommodation::skin-tone-2:": "1f6cc-1f3fb",
	":sleeping_accommodation::skin-tone-3:": "1f6cc-1f3fc",
	":sleeping_accommodation::skin-tone-4:": "1f6cc-1f3fd",
	":sleeping_accommodation::skin-tone-5:": "1f6cc-1f3fe",
	":sleeping_accommodation::skin-tone-6:": "1f6cc-1f3ff",
	// 1F6CD-FE0F shopping_bags
	":shopping_bags:": "1f6cd-fe0f",
	// 1F6CE-FE0F bellhop_bell
	":bellhop_bell:": "1f6ce-fe0f",
	// 1F6CF-FE0F bed
	":bed:": "1f6cf-fe0f",
	// 1F6D0 place_of_worship
	":place_of_worship:": "1f6d0",
	// 1F6D1 octagonal_sign
	":octagonal_sign:": "1f6d1",
	// 1F6D2 shopping_trolley
	":shopping_trolley:": "1f6d2",
	// 1F6D5 hindu_temple
	":hindu_temple:": "1f6d5",
	// 1F6D6 hut
	":hut:": "1f6d6",
	// 1F6D7 elevator
	":elevator:": "1f6d7",
	// 1F6DD playground_slide
	":playground_slide:": "1f6dd",
	// 1F6DE wheel
	":wheel:": "1f6de",
	// 1F6DF ring_buoy
	":ring_buoy:": "1f6df",
	// 1F6E0-FE0F hammer_and_wrench
	":hammer_and_wrench:": "1f6e0-fe0f",
	// 1F6E1-FE0F shield
	":shield:": "1f6e1-fe0f",
	// 1F6E2-FE0F oil_drum
	":oil_drum:": "1f6e2-fe0f",
	// 1F6E3-FE0F motorway
	":motorway:": "1f6e3-fe0f",
	// 1F6E4-FE0F railway_track
	":railway_track:": "1f6e4-fe0f",
	// 1F6E5-FE0F motor_boat
	":motor_boat:": "1f6e5-fe0f",
	// 1F6E9-FE0F small_airplane
	":small_airplane:": "1f6e9-fe0f",
	// 1F6EB airplane_departure
	":airplane_departure:": "1f6eb",
	// 1F6EC airplane_arriving
	":airplane_arriving:": "1f6ec",
	// 1F6F0-FE0F satellite
	":satellite:": "1f6f0-fe0f",
	// 1F6F3-FE0F passenger_ship
	":passenger_ship:": "1f6f3-fe0f",
	// 1F6F4 scooter
	":scooter:": "1f6f4",
	// 1F6F5 motor_scooter
	":motor_scooter:": "1f6f5",
	// 1F6F6 canoe
	":canoe:": "1f6f6",
	// 1F6F7 sled
	":sled:": "1f6f7",
	// 1F6F8 flying_saucer
	":flying_saucer:": "1f6f8",
	// 1F6F9 skateboard
	":skateboard:": "1f6f9",
	// 1F6FA auto_rickshaw
	":auto_rickshaw:": "1f6fa",
	// 1F6FB pickup_truck
	":pickup_truck:": "1f6fb",
	// 1F6FC roller_skate
	":roller_skate:": "1f6fc",
	// 1F7E0 large_orange_circle
	":large_orange_circle:": "1f7e0",
	// 1F7E1 large_yellow_circle
	":large_yellow_circle:": "1f7e1",
	// 1F7E2 large_green_circle
	":large_green_circle:": "1f7e2",
	// 1F7E3 large_purple_circle
	":large_purple_circle:": "1f7e3",
	// 1F7E4 large_brown_circle
	":large_brown_circle:": "1f7e4",
	// 1F7E5 large_red_square
	":large_red_square:": "1f7e5",
	// 1F7E6 large_blue_square
	":large_blue_square:": "1f7e6",
	// 1F7E7 large_orange_square
	":large_orange_square:": "1f7e7",
	// 1F7E8 large_yellow_square
	":large_yellow_square:": "1f7e8",
	// 1F7E9 large_green_square
	":large_green_square:": "1f7e9",
	// 1F7EA large_purple_square
	":large_purple_square:": "1f7ea",
	// 1F7EB large_brown_square
	":large_brown_square:": "1f7eb",
	// 1F7F0 heavy_equals_sign
	":heavy_equals_sign:": "1f7f0",
	// 1F90C pinched_fingers
	":pinched_fingers:":              "1f90c",
	":pinched_fingers::skin-tone-2:": "1f90c-1f3fb",
	":pinched_fingers::skin-tone-3:": "1f90c-1f3fc",
	":pinched_fingers::skin-tone-4:": "1f90c-1f3fd",
	":pinched_fingers::skin-tone-5:": "1f90c-1f3fe",
	":pinched_fingers::skin-tone-6:": "1f90c-1f3ff",
	// 1F90D white_heart
	":white_heart:": "1f90d",
	// 1F90E brown_heart
	":brown_heart:": "1f90e",
	// 1F90F pinching_hand
	":pinching_hand:":              "1f90f",
	":pinching_hand::skin-tone-2:": "1f90f-1f3fb",
	":pinching_hand::skin-tone-3:": "1f90f-1f3fc",
	":pinching_hand::skin-tone-4:": "1f90f-1f3fd",
	":pinching_hand::skin-tone-5:": "1f90f-1f3fe",
	":pinching_hand::skin-tone-6:": "1f90f-1f3ff",
	// 1F910 zipper_mouth_face
	":zipper_mouth_face:": "1f910",
	// 1F911 money_mouth_face
	":money_mouth_face:": "1f911",
	// 1F912 face_with_thermometer
	":face_with_thermometer:": "1f912",
	// 1F913 nerd_face
	":nerd_face:": "1f913",
	// 1F914 thinking_face
	":thinking_face:": "1f914",
	// 1F915 face_with_head_bandage
	":face_with_head_bandage:": "1f915",
	// 1F916 robot_face
	":robot_face:": "1f916",
	// 1F917 hugging_face
	":hugging_face:": "1f917",
	// 1F918 the_horns
	":the_horns:": "1f918",
	// 1F918 the_horns
	":sign_of_the_horns:":              "1f918",
	":the_horns::skin-tone-2:":         "1f918-1f3fb",
	":sign_of_the_horns::skin-tone-2:": "1f918-1f3fb",
	":the_horns::skin-tone-3:":         "1f918-1f3fc",
	":sign_of_the_horns::skin-tone-3:": "1f918-1f3fc",
	":the_horns::skin-tone-4:":         "1f918-1f3fd",
	":sign_of_the_horns::skin-tone-4:": "1f918-1f3fd",
	":the_horns::skin-tone-5:":         "1f918-1f3fe",
	":sign_of_the_horns::skin-tone-5:": "1f918-1f3fe",
	":the_horns::skin-tone-6:":         "1f918-1f3ff",
	":sign_of_the_horns::skin-tone-6:": "1f918-1f3ff",
	// 1F919 call_me_hand
	":call_me_hand:":              "1f919",
	":call_me_hand::skin-tone-2:": "1f919-1f3fb",
	":call_me_hand::skin-tone-3:": "1f919-1f3fc",
	":call_me_hand::skin-tone-4:": "1f919-1f3fd",
	":call_me_hand::skin-tone-5:": "1f919-1f3fe",
	":call_me_hand::skin-tone-6:": "1f919-1f3ff",
	// 1F91A raised_back_of_hand
	":raised_back_of_hand:":              "1f91a",
	":raised_back_of_hand::skin-tone-2:": "1f91a-1f3fb",
	":raised_back_of_hand::skin-tone-3:": "1f91a-1f3fc",
	":raised_back_of_hand::skin-tone-4:": "1f91a-1f3fd",
	":raised_back_of_hand::skin-tone-5:": "1f91a-1f3fe",
	":raised_back_of_hand::skin-tone-6:": "1f91a-1f3ff",
	// 1F91B left-facing_fist
	":left-facing_fist:":              "1f91b",
	":left-facing_fist::skin-tone-2:": "1f91b-1f3fb",
	":left-facing_fist::skin-tone-3:": "1f91b-1f3fc",
	":left-facing_fist::skin-tone-4:": "1f91b-1f3fd",
	":left-facing_fist::skin-tone-5:": "1f91b-1f3fe",
	":left-facing_fist::skin-tone-6:": "1f91b-1f3ff",
	// 1F91C right-facing_fist
	":right-facing_fist:":              "1f91c",
	":right-facing_fist::skin-tone-2:": "1f91c-1f3fb",
	":right-facing_fist::skin-tone-3:": "1f91c-1f3fc",
	":right-facing_fist::skin-tone-4:": "1f91c-1f3fd",
	":right-facing_fist::skin-tone-5:": "1f91c-1f3fe",
	":right-facing_fist::skin-tone-6:": "1f91c-1f3ff",
	// 1F91D handshake
	":handshake:":                "1f91d",
	":handshake::skin-tone-2:":   "1f91d-1f3fb",
	":handshake::skin-tone-3:":   "1f91d-1f3fc",
	":handshake::skin-tone-4:":   "1f91d-1f3fd",
	":handshake::skin-tone-5:":   "1f91d-1f3fe",
	":handshake::skin-tone-6:":   "1f91d-1f3ff",
	":handshake::skin-tone-2-3:": "1faf1-1f3fb-200d-1faf2-1f3fc",
	":handshake::skin-tone-2-4:": "1faf1-1f3fb-200d-1faf2-1f3fd",
	":handshake::skin-tone-2-5:": "1faf1-1f3fb-200d-1faf2-1f3fe",
	":handshake::skin-tone-2-6:": "1faf1-1f3fb-200d-1faf2-1f3ff",
	":handshake::skin-tone-3-2:": "1faf1-1f3fc-200d-1faf2-1f3fb",
	":handshake::skin-tone-3-4:": "1faf1-1f3fc-200d-1faf2-1f3fd",
	":handshake::skin-tone-3-5:": "1faf1-1f3fc-200d-1faf2-1f3fe",
	":handshake::skin-tone-3-6:": "1faf1-1f3fc-200d-1faf2-1f3ff",
	":handshake::skin-tone-4-2:": "1faf1-1f3fd-200d-1faf2-1f3fb",
	":handshake::skin-tone-4-3:": "1faf1-1f3fd-200d-1faf2-1f3fc",
	":handshake::skin-tone-4-5:": "1faf1-1f3fd-200d-1faf2-1f3fe",
	":handshake::skin-tone-4-6:": "1faf1-1f3fd-200d-1faf2-1f3ff",
	":handshake::skin-tone-5-2:": "1faf1-1f3fe-200d-1faf2-1f3fb",
	":handshake::skin-tone-5-3:": "1faf1-1f3fe-200d-1faf2-1f3fc",
	":handshake::skin-tone-5-4:": "1faf1-1f3fe-200d-1faf2-1f3fd",
	":handshake::skin-tone-5-6:": "1faf1-1f3fe-200d-1faf2-1f3ff",
	":handshake::skin-tone-6-2:": "1faf1-1f3ff-200d-1faf2-1f3fb",
	":handshake::skin-tone-6-3:": "1faf1-1f3ff-200d-1faf2-1f3fc",
	":handshake::skin-tone-6-4:": "1faf1-1f3ff-200d-1faf2-1f3fd",
	":handshake::skin-tone-6-5:": "1faf1-1f3ff-200d-1faf2-1f3fe",
	// 1F91E crossed_fingers
	":crossed_fingers:": "1f91e",
	// 1F91E crossed_fingers
	":hand_with_index_and_middle_fingers_crossed:":              "1f91e",
	":crossed_fingers::skin-tone-2:":                            "1f91e-1f3fb",
	":hand_with_index_and_middle_fingers_crossed::skin-tone-2:": "1f91e-1f3fb",
	":crossed_fingers::skin-tone-3:":                            "1f91e-1f3fc",
	":hand_with_index_and_middle_fingers_crossed::skin-tone-3:": "1f91e-1f3fc",
	":crossed_fingers::skin-tone-4:":                            "1f91e-1f3fd",
	":hand_with_index_and_middle_fingers_crossed::skin-tone-4:": "1f91e-1f3fd",
	":crossed_fingers::skin-tone-5:":                            "1f91e-1f3fe",
	":hand_with_index_and_middle_fingers_crossed::skin-tone-5:": "1f91e-1f3fe",
	":crossed_fingers::skin-tone-6:":                            "1f91e-1f3ff",
	":hand_with_index_and_middle_fingers_crossed::skin-tone-6:": "1f91e-1f3ff",
	// 1F91F i_love_you_hand_sign
	":i_love_you_hand_sign:":              "1f91f",
	":i_love_you_hand_sign::skin-tone-2:": "1f91f-1f3fb",
	":i_love_you_hand_sign::skin-tone-3:": "1f91f-1f3fc",
	":i_love_you_hand_sign::skin-tone-4:": "1f91f-1f3fd",
	":i_love_you_hand_sign::skin-tone-5:": "1f91f-1f3fe",
	":i_love_you_hand_sign::skin-tone-6:": "1f91f-1f3ff",
	// 1F920 face_with_cowboy_hat
	":face_with_cowboy_hat:": "1f920",
	// 1F921 clown_face
	":clown_face:": "1f921",
	// 1F922 nauseated_face
	":nauseated_face:": "1f922",
	// 1F923 rolling_on_the_floor_laughing
	":rolling_on_the_floor_laughing:": "1f923",
	// 1F924 drooling_face
	":drooling_face:": "1f924",
	// 1F925 lying_face
	":lying_face:": "1f925",
	// 1F926-200D-2640-FE0F woman-facepalming
	":woman-facepalming:":              "1f926-200d-2640-fe0f",
	":woman-facepalming::skin-tone-2:": "1f926-1f3fb-200d-2640-fe0f",
	":woman-facepalming::skin-tone-3:": "1f926-1f3fc-200d-2640-fe0f",
	":woman-facepalming::skin-tone-4:": "1f926-1f3fd-200d-2640-fe0f",
	":woman-facepalming::skin-tone-5:": "1f926-1f3fe-200d-2640-fe0f",
	":woman-facepalming::skin-tone-6:": "1f926-1f3ff-200d-2640-fe0f",
	// 1F926-200D-2642-FE0F man-facepalming
	":man-facepalming:":              "1f926-200d-2642-fe0f",
	":man-facepalming::skin-tone-2:": "1f926-1f3fb-200d-2642-fe0f",
	":man-facepalming::skin-tone-3:": "1f926-1f3fc-200d-2642-fe0f",
	":man-facepalming::skin-tone-4:": "1f926-1f3fd-200d-2642-fe0f",
	":man-facepalming::skin-tone-5:": "1f926-1f3fe-200d-2642-fe0f",
	":man-facepalming::skin-tone-6:": "1f926-1f3ff-200d-2642-fe0f",
	// 1F926 face_palm
	":face_palm:":              "1f926",
	":face_palm::skin-tone-2:": "1f926-1f3fb",
	":face_palm::skin-tone-3:": "1f926-1f3fc",
	":face_palm::skin-tone-4:": "1f926-1f3fd",
	":face_palm::skin-tone-5:": "1f926-1f3fe",
	":face_palm::skin-tone-6:": "1f926-1f3ff",
	// 1F927 sneezing_face
	":sneezing_face:": "1f927",
	// 1F928 face_with_raised_eyebrow
	":face_with_raised_eyebrow:": "1f928",
	// 1F928 face_with_raised_eyebrow
	":face_with_one_eyebrow_raised:": "1f928",
	// 1F929 star-struck
	":star-struck:": "1f929",
	// 1F929 star-struck
	":grinning_face_with_star_eyes:": "1f929",
	// 1F92A zany_face
	":zany_face:": "1f92a",
	// 1F92A zany_face
	":grinning_face_with_one_large_and_one_small_eye:": "1f92a",
	// 1F92B shushing_face
	":shushing_face:": "1f92b",
	// 1F92B shushing_face
	":face_with_finger_covering_closed_lips:": "1f92b",
	// 1F92C face_with_symbols_on_mouth
	":face_with_symbols_on_mouth:": "1f92c",
	// 1F92C face_with_symbols_on_mouth
	":serious_face_with_symbols_covering_mouth:": "1f92c",
	// 1F92D face_with_hand_over_mouth
	":face_with_hand_over_mouth:": "1f92d",
	// 1F92D face_with_hand_over_mouth
	":smiling_face_with_smiling_eyes_and_hand_covering_mouth:": "1f92d",
	// 1F92E face_vomiting
	":face_vomiting:": "1f92e",
	// 1F92E face_vomiting
	":face_with_open_mouth_vomiting:": "1f92e",
	// 1F92F exploding_head
	":exploding_head:": "1f92f",
	// 1F92F exploding_head
	":shocked_face_with_exploding_head:": "1f92f",
	// 1F930 pregnant_woman
	":pregnant_woman:":              "1f930",
	":pregnant_woman::skin-tone-2:": "1f930-1f3fb",
	":pregnant_woman::skin-tone-3:": "1f930-1f3fc",
	":pregnant_woman::skin-tone-4:": "1f930-1f3fd",
	":pregnant_woman::skin-tone-5:": "1f930-1f3fe",
	":pregnant_woman::skin-tone-6:": "1f930-1f3ff",
	// 1F931 breast-feeding
	":breast-feeding:":              "1f931",
	":breast-feeding::skin-tone-2:": "1f931-1f3fb",
	":breast-feeding::skin-tone-3:": "1f931-1f3fc",
	":breast-feeding::skin-tone-4:": "1f931-1f3fd",
	":breast-feeding::skin-tone-5:": "1f931-1f3fe",
	":breast-feeding::skin-tone-6:": "1f931-1f3ff",
	// 1F932 palms_up_together
	":palms_up_together:":              "1f932",
	":palms_up_together::skin-tone-2:": "1f932-1f3fb",
	":palms_up_together::skin-tone-3:": "1f932-1f3fc",
	":palms_up_together::skin-tone-4:": "1f932-1f3fd",
	":palms_up_together::skin-tone-5:": "1f932-1f3fe",
	":palms_up_together::skin-tone-6:": "1f932-1f3ff",
	// 1F933 selfie
	":selfie:":              "1f933",
	":selfie::skin-tone-2:": "1f933-1f3fb",
	":selfie::skin-tone-3:": "1f933-1f3fc",
	":selfie::skin-tone-4:": "1f933-1f3fd",
	":selfie::skin-tone-5:": "1f933-1f3fe",
	":selfie::skin-tone-6:": "1f933-1f3ff",
	// 1F934 prince
	":prince:":              "1f934",
	":prince::skin-tone-2:": "1f934-1f3fb",
	":prince::skin-tone-3:": "1f934-1f3fc",
	":prince::skin-tone-4:": "1f934-1f3fd",
	":prince::skin-tone-5:": "1f934-1f3fe",
	":prince::skin-tone-6:": "1f934-1f3ff",
	// 1F935-200D-2640-FE0F woman_in_tuxedo
	":woman_in_tuxedo:":              "1f935-200d-2640-fe0f",
	":woman_in_tuxedo::skin-tone-2:": "1f935-1f3fb-200d-2640-fe0f",
	":woman_in_tuxedo::skin-tone-3:": "1f935-1f3fc-200d-2640-fe0f",
	":woman_in_tuxedo::skin-tone-4:": "1f935-1f3fd-200d-2640-fe0f",
	":woman_in_tuxedo::skin-tone-5:": "1f935-1f3fe-200d-2640-fe0f",
	":woman_in_tuxedo::skin-tone-6:": "1f935-1f3ff-200d-2640-fe0f",
	// 1F935-200D-2642-FE0F man_in_tuxedo
	":man_in_tuxedo:":              "1f935-200d-2642-fe0f",
	":man_in_tuxedo::skin-tone-2:": "1f935-1f3fb-200d-2642-fe0f",
	":man_in_tuxedo::skin-tone-3:": "1f935-1f3fc-200d-2642-fe0f",
	":man_in_tuxedo::skin-tone-4:": "1f935-1f3fd-200d-2642-fe0f",
	":man_in_tuxedo::skin-tone-5:": "1f935-1f3fe-200d-2642-fe0f",
	":man_in_tuxedo::skin-tone-6:": "1f935-1f3ff-200d-2642-fe0f",
	// 1F935 person_in_tuxedo
	":person_in_tuxedo:":              "1f935",
	":person_in_tuxedo::skin-tone-2:": "1f935-1f3fb",
	":person_in_tuxedo::skin-tone-3:": "1f935-1f3fc",
	":person_in_tuxedo::skin-tone-4:": "1f935-1f3fd",
	":person_in_tuxedo::skin-tone-5:": "1f935-1f3fe",
	":person_in_tuxedo::skin-tone-6:": "1f935-1f3ff",
	// 1F936 mrs_claus
	":mrs_claus:": "1f936",
	// 1F936 mrs_claus
	":mother_christmas:":              "1f936",
	":mrs_claus::skin-tone-2:":        "1f936-1f3fb",
	":mother_christmas::skin-tone-2:": "1f936-1f3fb",
	":mrs_claus::skin-tone-3:":        "1f936-1f3fc",
	":mother_christmas::skin-tone-3:": "1f936-1f3fc",
	":mrs_claus::skin-tone-4:":        "1f936-1f3fd",
	":mother_christmas::skin-tone-4:": "1f936-1f3fd",
	":mrs_claus::skin-tone-5:":        "1f936-1f3fe",
	":mother_christmas::skin-tone-5:": "1f936-1f3fe",
	":mrs_claus::skin-tone-6:":        "1f936-1f3ff",
	":mother_christmas::skin-tone-6:": "1f936-1f3ff",
	// 1F937-200D-2640-FE0F woman-shrugging
	":woman-shrugging:":              "1f937-200d-2640-fe0f",
	":woman-shrugging::skin-tone-2:": "1f937-1f3fb-200d-2640-fe0f",
	":woman-shrugging::skin-tone-3:": "1f937-1f3fc-200d-2640-fe0f",
	":woman-shrugging::skin-tone-4:": "1f937-1f3fd-200d-2640-fe0f",
	":woman-shrugging::skin-tone-5:": "1f937-1f3fe-200d-2640-fe0f",
	":woman-shrugging::skin-tone-6:": "1f937-1f3ff-200d-2640-fe0f",
	// 1F937-200D-2642-FE0F man-shrugging
	":man-shrugging:":              "1f937-200d-2642-fe0f",
	":man-shrugging::skin-tone-2:": "1f937-1f3fb-200d-2642-fe0f",
	":man-shrugging::skin-tone-3:": "1f937-1f3fc-200d-2642-fe0f",
	":man-shrugging::skin-tone-4:": "1f937-1f3fd-200d-2642-fe0f",
	":man-shrugging::skin-tone-5:": "1f937-1f3fe-200d-2642-fe0f",
	":man-shrugging::skin-tone-6:": "1f937-1f3ff-200d-2642-fe0f",
	// 1F937 shrug
	":shrug:":              "1f937",
	":shrug::skin-tone-2:": "1f937-1f3fb",
	":shrug::skin-tone-3:": "1f937-1f3fc",
	":shrug::skin-tone-4:": "1f937-1f3fd",
	":shrug::skin-tone-5:": "1f937-1f3fe",
	":shrug::skin-tone-6:": "1f937-1f3ff",
	// 1F938-200D-2640-FE0F woman-cartwheeling
	":woman-cartwheeling:":              "1f938-200d-2640-fe0f",
	":woman-cartwheeling::skin-tone-2:": "1f938-1f3fb-200d-2640-fe0f",
	":woman-cartwheeling::skin-tone-3:": "1f938-1f3fc-200d-2640-fe0f",
	":woman-cartwheeling::skin-tone-4:": "1f938-1f3fd-200d-2640-fe0f",
	":woman-cartwheeling::skin-tone-5:": "1f938-1f3fe-200d-2640-fe0f",
	":woman-cartwheeling::skin-tone-6:": "1f938-1f3ff-200d-2640-fe0f",
	// 1F938-200D-2642-FE0F man-cartwheeling
	":man-cartwheeling:":              "1f938-200d-2642-fe0f",
	":man-cartwheeling::skin-tone-2:": "1f938-1f3fb-200d-2642-fe0f",
	":man-cartwheeling::skin-tone-3:": "1f938-1f3fc-200d-2642-fe0f",
	":man-cartwheeling::skin-tone-4:": "1f938-1f3fd-200d-2642-fe0f",
	":man-cartwheeling::skin-tone-5:": "1f938-1f3fe-200d-2642-fe0f",
	":man-cartwheeling::skin-tone-6:": "1f938-1f3ff-200d-2642-fe0f",
	// 1F938 person_doing_cartwheel
	":person_doing_cartwheel:":              "1f938",
	":person_doing_cartwheel::skin-tone-2:": "1f938-1f3fb",
	":person_doing_cartwheel::skin-tone-3:": "1f938-1f3fc",
	":person_doing_cartwheel::skin-tone-4:": "1f938-1f3fd",
	":person_doing_cartwheel::skin-tone-5:": "1f938-1f3fe",
	":person_doing_cartwheel::skin-tone-6:": "1f938-1f3ff",
	// 1F939-200D-2640-FE0F woman-juggling
	":woman-juggling:":              "1f939-200d-2640-fe0f",
	":woman-juggling::skin-tone-2:": "1f939-1f3fb-200d-2640-fe0f",
	":woman-juggling::skin-tone-3:": "1f939-1f3fc-200d-2640-fe0f",
	":woman-juggling::skin-tone-4:": "1f939-1f3fd-200d-2640-fe0f",
	":woman-juggling::skin-tone-5:": "1f939-1f3fe-200d-2640-fe0f",
	":woman-juggling::skin-tone-6:": "1f939-1f3ff-200d-2640-fe0f",
	// 1F939-200D-2642-FE0F man-juggling
	":man-juggling:":              "1f939-200d-2642-fe0f",
	":man-juggling::skin-tone-2:": "1f939-1f3fb-200d-2642-fe0f",
	":man-juggling::skin-tone-3:": "1f939-1f3fc-200d-2642-fe0f",
	":man-juggling::skin-tone-4:": "1f939-1f3fd-200d-2642-fe0f",
	":man-juggling::skin-tone-5:": "1f939-1f3fe-200d-2642-fe0f",
	":man-juggling::skin-tone-6:": "1f939-1f3ff-200d-2642-fe0f",
	// 1F939 juggling
	":juggling:":              "1f939",
	":juggling::skin-tone-2:": "1f939-1f3fb",
	":juggling::skin-tone-3:": "1f939-1f3fc",
	":juggling::skin-tone-4:": "1f939-1f3fd",
	":juggling::skin-tone-5:": "1f939-1f3fe",
	":juggling::skin-tone-6:": "1f939-1f3ff",
	// 1F93A fencer
	":fencer:": "1f93a",
	// 1F93C-200D-2640-FE0F woman-wrestling
	":woman-wrestling:": "1f93c-200d-2640-fe0f",
	// 1F93C-200D-2642-FE0F man-wrestling
	":man-wrestling:": "1f93c-200d-2642-fe0f",
	// 1F93C wrestlers
	":wrestlers:": "1f93c",
	// 1F93D-200D-2640-FE0F woman-playing-water-polo
	":woman-playing-water-polo:":              "1f93d-200d-2640-fe0f",
	":woman-playing-water-polo::skin-tone-2:": "1f93d-1f3fb-200d-2640-fe0f",
	":woman-playing-water-polo::skin-tone-3:": "1f93d-1f3fc-200d-2640-fe0f",
	":woman-playing-water-polo::skin-tone-4:": "1f93d-1f3fd-200d-2640-fe0f",
	":woman-playing-water-polo::skin-tone-5:": "1f93d-1f3fe-200d-2640-fe0f",
	":woman-playing-water-polo::skin-tone-6:": "1f93d-1f3ff-200d-2640-fe0f",
	// 1F93D-200D-2642-FE0F man-playing-water-polo
	":man-playing-water-polo:":              "1f93d-200d-2642-fe0f",
	":man-playing-water-polo::skin-tone-2:": "1f93d-1f3fb-200d-2642-fe0f",
	":man-playing-water-polo::skin-tone-3:": "1f93d-1f3fc-200d-2642-fe0f",
	":man-playing-water-polo::skin-tone-4:": "1f93d-1f3fd-200d-2642-fe0f",
	":man-playing-water-polo::skin-tone-5:": "1f93d-1f3fe-200d-2642-fe0f",
	":man-playing-water-polo::skin-tone-6:": "1f93d-1f3ff-200d-2642-fe0f",
	// 1F93D water_polo
	":water_polo:":              "1f93d",
	":water_polo::skin-tone-2:": "1f93d-1f3fb",
	":water_polo::skin-tone-3:": "1f93d-1f3fc",
	":water_polo::skin-tone-4:": "1f93d-1f3fd",
	":water_polo::skin-tone-5:": "1f93d-1f3fe",
	":water_polo::skin-tone-6:": "1f93d-1f3ff",
	// 1F93E-200D-2640-FE0F woman-playing-handball
	":woman-playing-handball:":              "1f93e-200d-2640-fe0f",
	":woman-playing-handball::skin-tone-2:": "1f93e-1f3fb-200d-2640-fe0f",
	":woman-playing-handball::skin-tone-3:": "1f93e-1f3fc-200d-2640-fe0f",
	":woman-playing-handball::skin-tone-4:": "1f93e-1f3fd-200d-2640-fe0f",
	":woman-playing-handball::skin-tone-5:": "1f93e-1f3fe-200d-2640-fe0f",
	":woman-playing-handball::skin-tone-6:": "1f93e-1f3ff-200d-2640-fe0f",
	// 1F93E-200D-2642-FE0F man-playing-handball
	":man-playing-handball:":              "1f93e-200d-2642-fe0f",
	":man-playing-handball::skin-tone-2:": "1f93e-1f3fb-200d-2642-fe0f",
	":man-playing-handball::skin-tone-3:": "1f93e-1f3fc-200d-2642-fe0f",
	":man-playing-handball::skin-tone-4:": "1f93e-1f3fd-200d-2642-fe0f",
	":man-playing-handball::skin-tone-5:": "1f93e-1f3fe-200d-2642-fe0f",
	":man-playing-handball::skin-tone-6:": "1f93e-1f3ff-200d-2642-fe0f",
	// 1F93E handball
	":handball:":              "1f93e",
	":handball::skin-tone-2:": "1f93e-1f3fb",
	":handball::skin-tone-3:": "1f93e-1f3fc",
	":handball::skin-tone-4:": "1f93e-1f3fd",
	":handball::skin-tone-5:": "1f93e-1f3fe",
	":handball::skin-tone-6:": "1f93e-1f3ff",
	// 1F93F diving_mask
	":diving_mask:": "1f93f",
	// 1F940 wilted_flower
	":wilted_flower:": "1f940",
	// 1F941 drum_with_drumsticks
	":drum_with_drumsticks:": "1f941",
	// 1F942 clinking_glasses
	":clinking_glasses:": "1f942",
	// 1F943 tumbler_glass
	":tumbler_glass:": "1f943",
	// 1F944 spoon
	":spoon:": "1f944",
	// 1F945 goal_net
	":goal_net:": "1f945",
	// 1F947 first_place_medal
	":first_place_medal:": "1f947",
	// 1F948 second_place_medal
	":second_place_medal:": "1f948",
	// 1F949 third_place_medal
	":third_place_medal:": "1f949",
	// 1F94A boxing_glove
	":boxing_glove:": "1f94a",
	// 1F94B martial_arts_uniform
	":martial_arts_uniform:": "1f94b",
	// 1F94C curling_stone
	":curling_stone:": "1f94c",
	// 1F94D lacrosse
	":lacrosse:": "1f94d",
	// 1F94E softball
	":softball:": "1f94e",
	// 1F94F flying_disc
	":flying_disc:": "1f94f",
	// 1F950 croissant
	":croissant:": "1f950",
	// 1F951 avocado
	":avocado:": "1f951",
	// 1F952 cucumber
	":cucumber:": "1f952",
	// 1F953 bacon
	":bacon:": "1f953",
	// 1F954 potato
	":potato:": "1f954",
	// 1F955 carrot
	":carrot:": "1f955",
	// 1F956 baguette_bread
	":baguette_bread:": "1f956",
	// 1F957 green_salad
	":green_salad:": "1f957",
	// 1F958 shallow_pan_of_food
	":shallow_pan_of_food:": "1f958",
	// 1F959 stuffed_flatbread
	":stuffed_flatbread:": "1f959",
	// 1F95A egg
	":egg:": "1f95a",
	// 1F95B glass_of_milk
	":glass_of_milk:": "1f95b",
	// 1F95C peanuts
	":peanuts:": "1f95c",
	// 1F95D kiwifruit
	":kiwifruit:": "1f95d",
	// 1F95E pancakes
	":pancakes:": "1f95e",
	// 1F95F dumpling
	":dumpling:": "1f95f",
	// 1F960 fortune_cookie
	":fortune_cookie:": "1f960",
	// 1F961 takeout_box
	":takeout_box:": "1f961",
	// 1F962 chopsticks
	":chopsticks:": "1f962",
	// 1F963 bowl_with_spoon
	":bowl_with_spoon:": "1f963",
	// 1F964 cup_with_straw
	":cup_with_straw:": "1f964",
	// 1F965 coconut
	":coconut:": "1f965",
	// 1F966 broccoli
	":broccoli:": "1f966",
	// 1F967 pie
	":pie:": "1f967",
	// 1F968 pretzel
	":pretzel:": "1f968",
	// 1F969 cut_of_meat
	":cut_of_meat:": "1f969",
	// 1F96A sandwich
	":sandwich:": "1f96a",
	// 1F96B canned_food
	":canned_food:": "1f96b",
	// 1F96C leafy_green
	":leafy_green:": "1f96c",
	// 1F96D mango
	":mango:": "1f96d",
	// 1F96E moon_cake
	":moon_cake:": "1f96e",
	// 1F96F bagel
	":bagel:": "1f96f",
	// 1F970 smiling_face_with_3_hearts
	":smiling_face_with_3_hearts:": "1f970",
	// 1F971 yawning_face
	":yawning_face:": "1f971",
	// 1F972 smiling_face_with_tear
	":smiling_face_with_tear:": "1f972",
	// 1F973 partying_face
	":partying_face:": "1f973",
	// 1F974 woozy_face
	":woozy_face:": "1f974",
	// 1F975 hot_face
	":hot_face:": "1f975",
	// 1F976 cold_face
	":cold_face:": "1f976",
	// 1F977 ninja
	":ninja:":              "1f977",
	":ninja::skin-tone-2:": "1f977-1f3fb",
	":ninja::skin-tone-3:": "1f977-1f3fc",
	":ninja::skin-tone-4:": "1f977-1f3fd",
	":ninja::skin-tone-5:": "1f977-1f3fe",
	":ninja::skin-tone-6:": "1f977-1f3ff",
	// 1F978 disguised_face
	":disguised_face:": "1f978",
	// 1F979 face_holding_back_tears
	":face_holding_back_tears:": "1f979",
	// 1F97A pleading_face
	":pleading_face:": "1f97a",
	// 1F97B sari
	":sari:": "1f97b",
	// 1F97C lab_coat
	":lab_coat:": "1f97c",
	// 1F97D goggles
	":goggles:": "1f97d",
	// 1F97E hiking_boot
	":hiking_boot:": "1f97e",
	// 1F97F womans_flat_shoe
	":womans_flat_shoe:": "1f97f",
	// 1F980 crab
	":crab:": "1f980",
	// 1F981 lion_face
	":lion_face:": "1f981",
	// 1F982 scorpion
	":scorpion:": "1f982",
	// 1F983 turkey
	":turkey:": "1f983",
	// 1F984 unicorn_face
	":unicorn_face:": "1f984",
	// 1F985 eagle
	":eagle:": "1f985",
	// 1F986 duck
	":duck:": "1f986",
	// 1F987 bat
	":bat:": "1f987",
	// 1F988 shark
	":shark:": "1f988",
	// 1F989 owl
	":owl:": "1f989",
	// 1F98A fox_face
	":fox_face:": "1f98a",
	// 1F98B butterfly
	":butterfly:": "1f98b",
	// 1F98C deer
	":deer:": "1f98c",
	// 1F98D gorilla
	":gorilla:": "1f98d",
	// 1F98E lizard
	":lizard:": "1f98e",
	// 1F98F rhinoceros
	":rhinoceros:": "1f98f",
	// 1F990 shrimp
	":shrimp:": "1f990",
	// 1F991 squid
	":squid:": "1f991",
	// 1F992 giraffe_face
	":giraffe_face:": "1f992",
	// 1F993 zebra_face
	":zebra_face:": "1f993",
	// 1F994 hedgehog
	":hedgehog:": "1f994",
	// 1F995 sauropod
	":sauropod:": "1f995",
	// 1F996 t-rex
	":t-rex:": "1f996",
	// 1F997 cricket
	":cricket:": "1f997",
	// 1F998 kangaroo
	":kangaroo:": "1f998",
	// 1F999 llama
	":llama:": "1f999",
	// 1F99A peacock
	":peacock:": "1f99a",
	// 1F99B hippopotamus
	":hippopotamus:": "1f99b",
	// 1F99C parrot
	":parrot:": "1f99c",
	// 1F99D raccoon
	":raccoon:": "1f99d",
	// 1F99E lobster
	":lobster:": "1f99e",
	// 1F99F mosquito
	":mosquito:": "1f99f",
	// 1F9A0 microbe
	":microbe:": "1f9a0",
	// 1F9A1 badger
	":badger:": "1f9a1",
	// 1F9A2 swan
	":swan:": "1f9a2",
	// 1F9A3 mammoth
	":mammoth:": "1f9a3",
	// 1F9A4 dodo
	":dodo:": "1f9a4",
	// 1F9A5 sloth
	":sloth:": "1f9a5",
	// 1F9A6 otter
	":otter:": "1f9a6",
	// 1F9A7 orangutan
	":orangutan:": "1f9a7",
	// 1F9A8 skunk
	":skunk:": "1f9a8",
	// 1F9A9 flamingo
	":flamingo:": "1f9a9",
	// 1F9AA oyster
	":oyster:": "1f9aa",
	// 1F9AB beaver
	":beaver:": "1f9ab",
	// 1F9AC bison
	":bison:": "1f9ac",
	// 1F9AD seal
	":seal:": "1f9ad",
	// 1F9AE guide_dog
	":guide_dog:": "1f9ae",
	// 1F9AF probing_cane
	":probing_cane:": "1f9af",
	// 1F9B4 bone
	":bone:": "1f9b4",
	// 1F9B5 leg
	":leg:":              "1f9b5",
	":leg::skin-tone-2:": "1f9b5-1f3fb",
	":leg::skin-tone-3:": "1f9b5-1f3fc",
	":leg::skin-tone-4:": "1f9b5-1f3fd",
	":leg::skin-tone-5:": "1f9b5-1f3fe",
	":leg::skin-tone-6:": "1f9b5-1f3ff",
	// 1F9B6 foot
	":foot:":              "1f9b6",
	":foot::skin-tone-2:": "1f9b6-1f3fb",
	":foot::skin-tone-3:": "1f9b6-1f3fc",
	":foot::skin-tone-4:": "1f9b6-1f3fd",
	":foot::skin-tone-5:": "1f9b6-1f3fe",
	":foot::skin-tone-6:": "1f9b6-1f3ff",
	// 1F9B7 tooth
	":tooth:": "1f9b7",
	// 1F9B8-200D-2640-FE0F female_superhero
	":female_superhero:":              "1f9b8-200d-2640-fe0f",
	":female_superhero::skin-tone-2:": "1f9b8-1f3fb-200d-2640-fe0f",
	":female_superhero::skin-tone-3:": "1f9b8-1f3fc-200d-2640-fe0f",
	":female_superhero::skin-tone-4:": "1f9b8-1f3fd-200d-2640-fe0f",
	":female_superhero::skin-tone-5:": "1f9b8-1f3fe-200d-2640-fe0f",
	":female_superhero::skin-tone-6:": "1f9b8-1f3ff-200d-2640-fe0f",
	// 1F9B8-200D-2642-FE0F male_superhero
	":male_superhero:":              "1f9b8-200d-2642-fe0f",
	":male_superhero::skin-tone-2:": "1f9b8-1f3fb-200d-2642-fe0f",
	":male_superhero::skin-tone-3:": "1f9b8-1f3fc-200d-2642-fe0f",
	":male_superhero::skin-tone-4:": "1f9b8-1f3fd-200d-2642-fe0f",
	":male_superhero::skin-tone-5:": "1f9b8-1f3fe-200d-2642-fe0f",
	":male_superhero::skin-tone-6:": "1f9b8-1f3ff-200d-2642-fe0f",
	// 1F9B8 superhero
	":superhero:":              "1f9b8",
	":superhero::skin-tone-2:": "1f9b8-1f3fb",
	":superhero::skin-tone-3:": "1f9b8-1f3fc",
	":superhero::skin-tone-4:": "1f9b8-1f3fd",
	":superhero::skin-tone-5:": "1f9b8-1f3fe",
	":superhero::skin-tone-6:": "1f9b8-1f3ff",
	// 1F9B9-200D-2640-FE0F female_supervillain
	":female_supervillain:":              "1f9b9-200d-2640-fe0f",
	":female_supervillain::skin-tone-2:": "1f9b9-1f3fb-200d-2640-fe0f",
	":female_supervillain::skin-tone-3:": "1f9b9-1f3fc-200d-2640-fe0f",
	":female_supervillain::skin-tone-4:": "1f9b9-1f3fd-200d-2640-fe0f",
	":female_supervillain::skin-tone-5:": "1f9b9-1f3fe-200d-2640-fe0f",
	":female_supervillain::skin-tone-6:": "1f9b9-1f3ff-200d-2640-fe0f",
	// 1F9B9-200D-2642-FE0F male_supervillain
	":male_supervillain:":              "1f9b9-200d-2642-fe0f",
	":male_supervillain::skin-tone-2:": "1f9b9-1f3fb-200d-2642-fe0f",
	":male_supervillain::skin-tone-3:": "1f9b9-1f3fc-200d-2642-fe0f",
	":male_supervillain::skin-tone-4:": "1f9b9-1f3fd-200d-2642-fe0f",
	":male_supervillain::skin-tone-5:": "1f9b9-1f3fe-200d-2642-fe0f",
	":male_supervillain::skin-tone-6:": "1f9b9-1f3ff-200d-2642-fe0f",
	// 1F9B9 supervillain
	":supervillain:":              "1f9b9",
	":supervillain::skin-tone-2:": "1f9b9-1f3fb",
	":supervillain::skin-tone-3:": "1f9b9-1f3fc",
	":supervillain::skin-tone-4:": "1f9b9-1f3fd",
	":supervillain::skin-tone-5:": "1f9b9-1f3fe",
	":supervillain::skin-tone-6:": "1f9b9-1f3ff",
	// 1F9BA safety_vest
	":safety_vest:": "1f9ba",
	// 1F9BB ear_with_hearing_aid
	":ear_with_hearing_aid:":              "1f9bb",
	":ear_with_hearing_aid::skin-tone-2:": "1f9bb-1f3fb",
	":ear_with_hearing_aid::skin-tone-3:": "1f9bb-1f3fc",
	":ear_with_hearing_aid::skin-tone-4:": "1f9bb-1f3fd",
	":ear_with_hearing_aid::skin-tone-5:": "1f9bb-1f3fe",
	":ear_with_hearing_aid::skin-tone-6:": "1f9bb-1f3ff",
	// 1F9BC motorized_wheelchair
	":motorized_wheelchair:": "1f9bc",
	// 1F9BD manual_wheelchair
	":manual_wheelchair:": "1f9bd",
	// 1F9BE mechanical_arm
	":mechanical_arm:": "1f9be",
	// 1F9BF mechanical_leg
	":mechanical_leg:": "1f9bf",
	// 1F9C0 cheese_wedge
	":cheese_wedge:": "1f9c0",
	// 1F9C1 cupcake
	":cupcake:": "1f9c1",
	// 1F9C2 salt
	":salt:": "1f9c2",
	// 1F9C3 beverage_box
	":beverage_box:": "1f9c3",
	// 1F9C4 garlic
	":garlic:": "1f9c4",
	// 1F9C5 onion
	":onion:": "1f9c5",
	// 1F9C6 falafel
	":falafel:": "1f9c6",
	// 1F9C7 waffle
	":waffle:": "1f9c7",
	// 1F9C8 butter
	":butter:": "1f9c8",
	// 1F9C9 mate_drink
	":mate_drink:": "1f9c9",
	// 1F9CA ice_cube
	":ice_cube:": "1f9ca",
	// 1F9CB bubble_tea
	":bubble_tea:": "1f9cb",
	// 1F9CC troll
	":troll:": "1f9cc",
	// 1F9CD-200D-2640-FE0F woman_standing
	":woman_standing:":              "1f9cd-200d-2640-fe0f",
	":woman_standing::skin-tone-2:": "1f9cd-1f3fb-200d-2640-fe0f",
	":woman_standing::skin-tone-3:": "1f9cd-1f3fc-200d-2640-fe0f",
	":woman_standing::skin-tone-4:": "1f9cd-1f3fd-200d-2640-fe0f",
	":woman_standing::skin-tone-5:": "1f9cd-1f3fe-200d-2640-fe0f",
	":woman_standing::skin-tone-6:": "1f9cd-1f3ff-200d-2640-fe0f",
	// 1F9CD-200D-2642-FE0F man_standing
	":man_standing:":              "1f9cd-200d-2642-fe0f",
	":man_standing::skin-tone-2:": "1f9cd-1f3fb-200d-2642-fe0f",
	":man_standing::skin-tone-3:": "1f9cd-1f3fc-200d-2642-fe0f",
	":man_standing::skin-tone-4:": "1f9cd-1f3fd-200d-2642-fe0f",
	":man_standing::skin-tone-5:": "1f9cd-1f3fe-200d-2642-fe0f",
	":man_standing::skin-tone-6:": "1f9cd-1f3ff-200d-2642-fe0f",
	// 1F9CD standing_person
	":standing_person:":              "1f9cd",
	":standing_person::skin-tone-2:": "1f9cd-1f3fb",
	":standing_person::skin-tone-3:": "1f9cd-1f3fc",
	":standing_person::skin-tone-4:": "1f9cd-1f3fd",
	":standing_person::skin-tone-5:": "1f9cd-1f3fe",
	":standing_person::skin-tone-6:": "1f9cd-1f3ff",
	// 1F9CE-200D-2640-FE0F woman_kneeling
	":woman_kneeling:":              "1f9ce-200d-2640-fe0f",
	":woman_kneeling::skin-tone-2:": "1f9ce-1f3fb-200d-2640-fe0f",
	":woman_kneeling::skin-tone-3:": "1f9ce-1f3fc-200d-2640-fe0f",
	":woman_kneeling::skin-tone-4:": "1f9ce-1f3fd-200d-2640-fe0f",
	":woman_kneeling::skin-tone-5:": "1f9ce-1f3fe-200d-2640-fe0f",
	":woman_kneeling::skin-tone-6:": "1f9ce-1f3ff-200d-2640-fe0f",
	// 1F9CE-200D-2642-FE0F man_kneeling
	":man_kneeling:":              "1f9ce-200d-2642-fe0f",
	":man_kneeling::skin-tone-2:": "1f9ce-1f3fb-200d-2642-fe0f",
	":man_kneeling::skin-tone-3:": "1f9ce-1f3fc-200d-2642-fe0f",
	":man_kneeling::skin-tone-4:": "1f9ce-1f3fd-200d-2642-fe0f",
	":man_kneeling::skin-tone-5:": "1f9ce-1f3fe-200d-2642-fe0f",
	":man_kneeling::skin-tone-6:": "1f9ce-1f3ff-200d-2642-fe0f",
	// 1F9CE kneeling_person
	":kneeling_person:":              "1f9ce",
	":kneeling_person::skin-tone-2:": "1f9ce-1f3fb",
	":kneeling_person::skin-tone-3:": "1f9ce-1f3fc",
	":kneeling_person::skin-tone-4:": "1f9ce-1f3fd",
	":kneeling_person::skin-tone-5:": "1f9ce-1f3fe",
	":kneeling_person::skin-tone-6:": "1f9ce-1f3ff",
	// 1F9CF-200D-2640-FE0F deaf_woman
	":deaf_woman:":              "1f9cf-200d-2640-fe0f",
	":deaf_woman::skin-tone-2:": "1f9cf-1f3fb-200d-2640-fe0f",
	":deaf_woman::skin-tone-3:": "1f9cf-1f3fc-200d-2640-fe0f",
	":deaf_woman::skin-tone-4:": "1f9cf-1f3fd-200d-2640-fe0f",
	":deaf_woman::skin-tone-5:": "1f9cf-1f3fe-200d-2640-fe0f",
	":deaf_woman::skin-tone-6:": "1f9cf-1f3ff-200d-2640-fe0f",
	// 1F9CF-200D-2642-FE0F deaf_man
	":deaf_man:":              "1f9cf-200d-2642-fe0f",
	":deaf_man::skin-tone-2:": "1f9cf-1f3fb-200d-2642-fe0f",
	":deaf_man::skin-tone-3:": "1f9cf-1f3fc-200d-2642-fe0f",
	":deaf_man::skin-tone-4:": "1f9cf-1f3fd-200d-2642-fe0f",
	":deaf_man::skin-tone-5:": "1f9cf-1f3fe-200d-2642-fe0f",
	":deaf_man::skin-tone-6:": "1f9cf-1f3ff-200d-2642-fe0f",
	// 1F9CF deaf_person
	":deaf_person:":              "1f9cf",
	":deaf_person::skin-tone-2:": "1f9cf-1f3fb",
	":deaf_person::skin-tone-3:": "1f9cf-1f3fc",
	":deaf_person::skin-tone-4:": "1f9cf-1f3fd",
	":deaf_person::skin-tone-5:": "1f9cf-1f3fe",
	":deaf_person::skin-tone-6:": "1f9cf-1f3ff",
	// 1F9D0 face_with_monocle
	":face_with_monocle:": "1f9d0",
	// 1F9D1-200D-1F33E farmer
	":farmer:":              "1f9d1-200d-1f33e",
	":farmer::skin-tone-2:": "1f9d1-1f3fb-200d-1f33e",
	":farmer::skin-tone-3:": "1f9d1-1f3fc-200d-1f33e",
	":farmer::skin-tone-4:": "1f9d1-1f3fd-200d-1f33e",
	":farmer::skin-tone-5:": "1f9d1-1f3fe-200d-1f33e",
	":farmer::skin-tone-6:": "1f9d1-1f3ff-200d-1f33e",
	// 1F9D1-200D-1F373 cook
	":cook:":              "1f9d1-200d-1f373",
	":cook::skin-tone-2:": "1f9d1-1f3fb-200d-1f373",
	":cook::skin-tone-3:": "1f9d1-1f3fc-200d-1f373",
	":cook::skin-tone-4:": "1f9d1-1f3fd-200d-1f373",
	":cook::skin-tone-5:": "1f9d1-1f3fe-200d-1f373",
	":cook::skin-tone-6:": "1f9d1-1f3ff-200d-1f373",
	// 1F9D1-200D-1F37C person_feeding_baby
	":person_feeding_baby:":              "1f9d1-200d-1f37c",
	":person_feeding_baby::skin-tone-2:": "1f9d1-1f3fb-200d-1f37c",
	":person_feeding_baby::skin-tone-3:": "1f9d1-1f3fc-200d-1f37c",
	":person_feeding_baby::skin-tone-4:": "1f9d1-1f3fd-200d-1f37c",
	":person_feeding_baby::skin-tone-5:": "1f9d1-1f3fe-200d-1f37c",
	":person_feeding_baby::skin-tone-6:": "1f9d1-1f3ff-200d-1f37c",
	// 1F9D1-200D-1F384 mx_claus
	":mx_claus:":              "1f9d1-200d-1f384",
	":mx_claus::skin-tone-2:": "1f9d1-1f3fb-200d-1f384",
	":mx_claus::skin-tone-3:": "1f9d1-1f3fc-200d-1f384",
	":mx_claus::skin-tone-4:": "1f9d1-1f3fd-200d-1f384",
	":mx_claus::skin-tone-5:": "1f9d1-1f3fe-200d-1f384",
	":mx_claus::skin-tone-6:": "1f9d1-1f3ff-200d-1f384",
	// 1F9D1-200D-1F393 student
	":student:":              "1f9d1-200d-1f393",
	":student::skin-tone-2:": "1f9d1-1f3fb-200d-1f393",
	":student::skin-tone-3:": "1f9d1-1f3fc-200d-1f393",
	":student::skin-tone-4:": "1f9d1-1f3fd-200d-1f393",
	":student::skin-tone-5:": "1f9d1-1f3fe-200d-1f393",
	":student::skin-tone-6:": "1f9d1-1f3ff-200d-1f393",
	// 1F9D1-200D-1F3A4 singer
	":singer:":              "1f9d1-200d-1f3a4",
	":singer::skin-tone-2:": "1f9d1-1f3fb-200d-1f3a4",
	":singer::skin-tone-3:": "1f9d1-1f3fc-200d-1f3a4",
	":singer::skin-tone-4:": "1f9d1-1f3fd-200d-1f3a4",
	":singer::skin-tone-5:": "1f9d1-1f3fe-200d-1f3a4",
	":singer::skin-tone-6:": "1f9d1-1f3ff-200d-1f3a4",
	// 1F9D1-200D-1F3A8 artist
	":artist:":              "1f9d1-200d-1f3a8",
	":artist::skin-tone-2:": "1f9d1-1f3fb-200d-1f3a8",
	":artist::skin-tone-3:": "1f9d1-1f3fc-200d-1f3a8",
	":artist::skin-tone-4:": "1f9d1-1f3fd-200d-1f3a8",
	":artist::skin-tone-5:": "1f9d1-1f3fe-200d-1f3a8",
	":artist::skin-tone-6:": "1f9d1-1f3ff-200d-1f3a8",
	// 1F9D1-200D-1F3EB teacher
	":teacher:":              "1f9d1-200d-1f3eb",
	":teacher::skin-tone-2:": "1f9d1-1f3fb-200d-1f3eb",
	":teacher::skin-tone-3:": "1f9d1-1f3fc-200d-1f3eb",
	":teacher::skin-tone-4:": "1f9d1-1f3fd-200d-1f3eb",
	":teacher::skin-tone-5:": "1f9d1-1f3fe-200d-1f3eb",
	":teacher::skin-tone-6:": "1f9d1-1f3ff-200d-1f3eb",
	// 1F9D1-200D-1F3ED factory_worker
	":factory_worker:":              "1f9d1-200d-1f3ed",
	":factory_worker::skin-tone-2:": "1f9d1-1f3fb-200d-1f3ed",
	":factory_worker::skin-tone-3:": "1f9d1-1f3fc-200d-1f3ed",
	":factory_worker::skin-tone-4:": "1f9d1-1f3fd-200d-1f3ed",
	":factory_worker::skin-tone-5:": "1f9d1-1f3fe-200d-1f3ed",
	":factory_worker::skin-tone-6:": "1f9d1-1f3ff-200d-1f3ed",
	// 1F9D1-200D-1F4BB technologist
	":technologist:":              "1f9d1-200d-1f4bb",
	":technologist::skin-tone-2:": "1f9d1-1f3fb-200d-1f4bb",
	":technologist::skin-tone-3:": "1f9d1-1f3fc-200d-1f4bb",
	":technologist::skin-tone-4:": "1f9d1-1f3fd-200d-1f4bb",
	":technologist::skin-tone-5:": "1f9d1-1f3fe-200d-1f4bb",
	":technologist::skin-tone-6:": "1f9d1-1f3ff-200d-1f4bb",
	// 1F9D1-200D-1F4BC office_worker
	":office_worker:":              "1f9d1-200d-1f4bc",
	":office_worker::skin-tone-2:": "1f9d1-1f3fb-200d-1f4bc",
	":office_worker::skin-tone-3:": "1f9d1-1f3fc-200d-1f4bc",
	":office_worker::skin-tone-4:": "1f9d1-1f3fd-200d-1f4bc",
	":office_worker::skin-tone-5:": "1f9d1-1f3fe-200d-1f4bc",
	":office_worker::skin-tone-6:": "1f9d1-1f3ff-200d-1f4bc",
	// 1F9D1-200D-1F527 mechanic
	":mechanic:":              "1f9d1-200d-1f527",
	":mechanic::skin-tone-2:": "1f9d1-1f3fb-200d-1f527",
	":mechanic::skin-tone-3:": "1f9d1-1f3fc-200d-1f527",
	":mechanic::skin-tone-4:": "1f9d1-1f3fd-200d-1f527",
	":mechanic::skin-tone-5:": "1f9d1-1f3fe-200d-1f527",
	":mechanic::skin-tone-6:": "1f9d1-1f3ff-200d-1f527",
	// 1F9D1-200D-1F52C scientist
	":scientist:":              "1f9d1-200d-1f52c",
	":scientist::skin-tone-2:": "1f9d1-1f3fb-200d-1f52c",
	":scientist::skin-tone-3:": "1f9d1-1f3fc-200d-1f52c",
	":scientist::skin-tone-4:": "1f9d1-1f3fd-200d-1f52c",
	":scientist::skin-tone-5:": "1f9d1-1f3fe-200d-1f52c",
	":scientist::skin-tone-6:": "1f9d1-1f3ff-200d-1f52c",
	// 1F9D1-200D-1F680 astronaut
	":astronaut:":              "1f9d1-200d-1f680",
	":astronaut::skin-tone-2:": "1f9d1-1f3fb-200d-1f680",
	":astronaut::skin-tone-3:": "1f9d1-1f3fc-200d-1f680",
	":astronaut::skin-tone-4:": "1f9d1-1f3fd-200d-1f680",
	":astronaut::skin-tone-5:": "1f9d1-1f3fe-200d-1f680",
	":astronaut::skin-tone-6:": "1f9d1-1f3ff-200d-1f680",
	// 1F9D1-200D-1F692 firefighter
	":firefighter:":              "1f9d1-200d-1f692",
	":firefighter::skin-tone-2:": "1f9d1-1f3fb-200d-1f692",
	":firefighter::skin-tone-3:": "1f9d1-1f3fc-200d-1f692",
	":firefighter::skin-tone-4:": "1f9d1-1f3fd-200d-1f692",
	":firefighter::skin-tone-5:": "1f9d1-1f3fe-200d-1f692",
	":firefighter::skin-tone-6:": "1f9d1-1f3ff-200d-1f692",
	// 1F9D1-200D-1F91D-200D-1F9D1 people_holding_hands
	":people_holding_hands:":                "1f9d1-200d-1f91d-200d-1f9d1",
	":people_holding_hands::skin-tone-2:":   "1f9d1-1f3fb-200d-1f91d-200d-1f9d1-1f3fb",
	":people_holding_hands::skin-tone-2-3:": "1f9d1-1f3fb-200d-1f91d-200d-1f9d1-1f3fc",
	":people_holding_hands::skin-tone-2-4:": "1f9d1-1f3fb-200d-1f91d-200d-1f9d1-1f3fd",
	":people_holding_hands::skin-tone-2-5:": "1f9d1-1f3fb-200d-1f91d-200d-1f9d1-1f3fe",
	":people_holding_hands::skin-tone-2-6:": "1f9d1-1f3fb-200d-1f91d-200d-1f9d1-1f3ff",
	":people_holding_hands::skin-tone-3-2:": "1f9d1-1f3fc-200d-1f91d-200d-1f9d1-1f3fb",
	":people_holding_hands::skin-tone-3:":   "1f9d1-1f3fc-200d-1f91d-200d-1f9d1-1f3fc",
	":people_holding_hands::skin-tone-3-4:": "1f9d1-1f3fc-200d-1f91d-200d-1f9d1-1f3fd",
	":people_holding_hands::skin-tone-3-5:": "1f9d1-1f3fc-200d-1f91d-200d-1f9d1-1f3fe",
	":people_holding_hands::skin-tone-3-6:": "1f9d1-1f3fc-200d-1f91d-200d-1f9d1-1f3ff",
	":people_holding_hands::skin-tone-4-2:": "1f9d1-1f3fd-200d-1f91d-200d-1f9d1-1f3fb",
	":people_holding_hands::skin-tone-4-3:": "1f9d1-1f3fd-200d-1f91d-200d-1f9d1-1f3fc",
	":people_holding_hands::skin-tone-4:":   "1f9d1-1f3fd-200d-1f91d-200d-1f9d1-1f3fd",
	":people_holding_hands::skin-tone-4-5:": "1f9d1-1f3fd-200d-1f91d-200d-1f9d1-1f3fe",
	":people_holding_hands::skin-tone-4-6:": "1f9d1-1f3fd-200d-1f91d-200d-1f9d1-1f3ff",
	":people_holding_hands::skin-tone-5-2:": "1f9d1-1f3fe-200d-1f91d-200d-1f9d1-1f3fb",
	":people_holding_hands::skin-tone-5-3:": "1f9d1-1f3fe-200d-1f91d-200d-1f9d1-1f3fc",
	":people_holding_hands::skin-tone-5-4:": "1f9d1-1f3fe-200d-1f91d-200d-1f9d1-1f3fd",
	":people_holding_hands::skin-tone-5:":   "1f9d1-1f3fe-200d-1f91d-200d-1f9d1-1f3fe",
	":people_holding_hands::skin-tone-5-6:": "1f9d1-1f3fe-200d-1f91d-200d-1f9d1-1f3ff",
	":people_holding_hands::skin-tone-6-2:": "1f9d1-1f3ff-200d-1f91d-200d-1f9d1-1f3fb",
	":people_holding_hands::skin-tone-6-3:": "1f9d1-1f3ff-200d-1f91d-200d-1f9d1-1f3fc",
	":people_holding_hands::skin-tone-6-4:": "1f9d1-1f3ff-200d-1f91d-200d-1f9d1-1f3fd",
	":people_holding_hands::skin-tone-6-5:": "1f9d1-1f3ff-200d-1f91d-200d-1f9d1-1f3fe",
	":people_holding_hands::skin-tone-6:":   "1f9d1-1f3ff-200d-1f91d-200d-1f9d1-1f3ff",
	// 1F9D1-200D-1F9AF person_with_probing_cane
	":person_with_probing_cane:":              "1f9d1-200d-1f9af",
	":person_with_probing_cane::skin-tone-2:": "1f9d1-1f3fb-200d-1f9af",
	":person_with_probing_cane::skin-tone-3:": "1f9d1-1f3fc-200d-1f9af",
	":person_with_probing_cane::skin-tone-4:": "1f9d1-1f3fd-200d-1f9af",
	":person_with_probing_cane::skin-tone-5:": "1f9d1-1f3fe-200d-1f9af",
	":person_with_probing_cane::skin-tone-6:": "1f9d1-1f3ff-200d-1f9af",
	// 1F9D1-200D-1F9B0 red_haired_person
	":red_haired_person:":              "1f9d1-200d-1f9b0",
	":red_haired_person::skin-tone-2:": "1f9d1-1f3fb-200d-1f9b0",
	":red_haired_person::skin-tone-3:": "1f9d1-1f3fc-200d-1f9b0",
	":red_haired_person::skin-tone-4:": "1f9d1-1f3fd-200d-1f9b0",
	":red_haired_person::skin-tone-5:": "1f9d1-1f3fe-200d-1f9b0",
	":red_haired_person::skin-tone-6:": "1f9d1-1f3ff-200d-1f9b0",
	// 1F9D1-200D-1F9B1 curly_haired_person
	":curly_haired_person:":              "1f9d1-200d-1f9b1",
	":curly_haired_person::skin-tone-2:": "1f9d1-1f3fb-200d-1f9b1",
	":curly_haired_person::skin-tone-3:": "1f9d1-1f3fc-200d-1f9b1",
	":curly_haired_person::skin-tone-4:": "1f9d1-1f3fd-200d-1f9b1",
	":curly_haired_person::skin-tone-5:": "1f9d1-1f3fe-200d-1f9b1",
	":curly_haired_person::skin-tone-6:": "1f9d1-1f3ff-200d-1f9b1",
	// 1F9D1-200D-1F9B2 bald_person
	":bald_person:":              "1f9d1-200d-1f9b2",
	":bald_person::skin-tone-2:": "1f9d1-1f3fb-200d-1f9b2",
	":bald_person::skin-tone-3:": "1f9d1-1f3fc-200d-1f9b2",
	":bald_person::skin-tone-4:": "1f9d1-1f3fd-200d-1f9b2",
	":bald_person::skin-tone-5:": "1f9d1-1f3fe-200d-1f9b2",
	":bald_person::skin-tone-6:": "1f9d1-1f3ff-200d-1f9b2",
	// 1F9D1-200D-1F9B3 white_haired_person
	":white_haired_person:":              "1f9d1-200d-1f9b3",
	":white_haired_person::skin-tone-2:": "1f9d1-1f3fb-200d-1f9b3",
	":white_haired_person::skin-tone-3:": "1f9d1-1f3fc-200d-1f9b3",
	":white_haired_person::skin-tone-4:": "1f9d1-1f3fd-200d-1f9b3",
	":white_haired_person::skin-tone-5:": "1f9d1-1f3fe-200d-1f9b3",
	":white_haired_person::skin-tone-6:": "1f9d1-1f3ff-200d-1f9b3",
	// 1F9D1-200D-1F9BC person_in_motorized_wheelchair
	":person_in_motorized_wheelchair:":              "1f9d1-200d-1f9bc",
	":person_in_motorized_wheelchair::skin-tone-2:": "1f9d1-1f3fb-200d-1f9bc",
	":person_in_motorized_wheelchair::skin-tone-3:": "1f9d1-1f3fc-200d-1f9bc",
	":person_in_motorized_wheelchair::skin-tone-4:": "1f9d1-1f3fd-200d-1f9bc",
	":person_in_motorized_wheelchair::skin-tone-5:": "1f9d1-1f3fe-200d-1f9bc",
	":person_in_motorized_wheelchair::skin-tone-6:": "1f9d1-1f3ff-200d-1f9bc",
	// 1F9D1-200D-1F9BD person_in_manual_wheelchair
	":person_in_manual_wheelchair:":              "1f9d1-200d-1f9bd",
	":person_in_manual_wheelchair::skin-tone-2:": "1f9d1-1f3fb-200d-1f9bd",
	":person_in_manual_wheelchair::skin-tone-3:": "1f9d1-1f3fc-200d-1f9bd",
	":person_in_manual_wheelchair::skin-tone-4:": "1f9d1-1f3fd-200d-1f9bd",
	":person_in_manual_wheelchair::skin-tone-5:": "1f9d1-1f3fe-200d-1f9bd",
	":person_in_manual_wheelchair::skin-tone-6:": "1f9d1-1f3ff-200d-1f9bd",
	// 1F9D1-200D-2695-FE0F health_worker
	":health_worker:":              "1f9d1-200d-2695-fe0f",
	":health_worker::skin-tone-2:": "1f9d1-1f3fb-200d-2695-fe0f",
	":health_worker::skin-tone-3:": "1f9d1-1f3fc-200d-2695-fe0f",
	":health_worker::skin-tone-4:": "1f9d1-1f3fd-200d-2695-fe0f",
	":health_worker::skin-tone-5:": "1f9d1-1f3fe-200d-2695-fe0f",
	":health_worker::skin-tone-6:": "1f9d1-1f3ff-200d-2695-fe0f",
	// 1F9D1-200D-2696-FE0F judge
	":judge:":              "1f9d1-200d-2696-fe0f",
	":judge::skin-tone-2:": "1f9d1-1f3fb-200d-2696-fe0f",
	":judge::skin-tone-3:": "1f9d1-1f3fc-200d-2696-fe0f",
	":judge::skin-tone-4:": "1f9d1-1f3fd-200d-2696-fe0f",
	":judge::skin-tone-5:": "1f9d1-1f3fe-200d-2696-fe0f",
	":judge::skin-tone-6:": "1f9d1-1f3ff-200d-2696-fe0f",
	// 1F9D1-200D-2708-FE0F pilot
	":pilot:":              "1f9d1-200d-2708-fe0f",
	":pilot::skin-tone-2:": "1f9d1-1f3fb-200d-2708-fe0f",
	":pilot::skin-tone-3:": "1f9d1-1f3fc-200d-2708-fe0f",
	":pilot::skin-tone-4:": "1f9d1-1f3fd-200d-2708-fe0f",
	":pilot::skin-tone-5:": "1f9d1-1f3fe-200d-2708-fe0f",
	":pilot::skin-tone-6:": "1f9d1-1f3ff-200d-2708-fe0f",
	// 1F9D1 adult
	":adult:":              "1f9d1",
	":adult::skin-tone-2:": "1f9d1-1f3fb",
	":adult::skin-tone-3:": "1f9d1-1f3fc",
	":adult::skin-tone-4:": "1f9d1-1f3fd",
	":adult::skin-tone-5:": "1f9d1-1f3fe",
	":adult::skin-tone-6:": "1f9d1-1f3ff",
	// 1F9D2 child
	":child:":              "1f9d2",
	":child::skin-tone-2:": "1f9d2-1f3fb",
	":child::skin-tone-3:": "1f9d2-1f3fc",
	":child::skin-tone-4:": "1f9d2-1f3fd",
	":child::skin-tone-5:": "1f9d2-1f3fe",
	":child::skin-tone-6:": "1f9d2-1f3ff",
	// 1F9D3 older_adult
	":older_adult:":              "1f9d3",
	":older_adult::skin-tone-2:": "1f9d3-1f3fb",
	":older_adult::skin-tone-3:": "1f9d3-1f3fc",
	":older_adult::skin-tone-4:": "1f9d3-1f3fd",
	":older_adult::skin-tone-5:": "1f9d3-1f3fe",
	":older_adult::skin-tone-6:": "1f9d3-1f3ff",
	// 1F9D4-200D-2640-FE0F woman_with_beard
	":woman_with_beard:":              "1f9d4-200d-2640-fe0f",
	":woman_with_beard::skin-tone-2:": "1f9d4-1f3fb-200d-2640-fe0f",
	":woman_with_beard::skin-tone-3:": "1f9d4-1f3fc-200d-2640-fe0f",
	":woman_with_beard::skin-tone-4:": "1f9d4-1f3fd-200d-2640-fe0f",
	":woman_with_beard::skin-tone-5:": "1f9d4-1f3fe-200d-2640-fe0f",
	":woman_with_beard::skin-tone-6:": "1f9d4-1f3ff-200d-2640-fe0f",
	// 1F9D4-200D-2642-FE0F man_with_beard
	":man_with_beard:":              "1f9d4-200d-2642-fe0f",
	":man_with_beard::skin-tone-2:": "1f9d4-1f3fb-200d-2642-fe0f",
	":man_with_beard::skin-tone-3:": "1f9d4-1f3fc-200d-2642-fe0f",
	":man_with_beard::skin-tone-4:": "1f9d4-1f3fd-200d-2642-fe0f",
	":man_with_beard::skin-tone-5:": "1f9d4-1f3fe-200d-2642-fe0f",
	":man_with_beard::skin-tone-6:": "1f9d4-1f3ff-200d-2642-fe0f",
	// 1F9D4 bearded_person
	":bearded_person:":              "1f9d4",
	":bearded_person::skin-tone-2:": "1f9d4-1f3fb",
	":bearded_person::skin-tone-3:": "1f9d4-1f3fc",
	":bearded_person::skin-tone-4:": "1f9d4-1f3fd",
	":bearded_person::skin-tone-5:": "1f9d4-1f3fe",
	":bearded_person::skin-tone-6:": "1f9d4-1f3ff",
	// 1F9D5 person_with_headscarf
	":person_with_headscarf:":              "1f9d5",
	":person_with_headscarf::skin-tone-2:": "1f9d5-1f3fb",
	":person_with_headscarf::skin-tone-3:": "1f9d5-1f3fc",
	":person_with_headscarf::skin-tone-4:": "1f9d5-1f3fd",
	":person_with_headscarf::skin-tone-5:": "1f9d5-1f3fe",
	":person_with_headscarf::skin-tone-6:": "1f9d5-1f3ff",
	// 1F9D6-200D-2640-FE0F woman_in_steamy_room
	":woman_in_steamy_room:":              "1f9d6-200d-2640-fe0f",
	":woman_in_steamy_room::skin-tone-2:": "1f9d6-1f3fb-200d-2640-fe0f",
	":woman_in_steamy_room::skin-tone-3:": "1f9d6-1f3fc-200d-2640-fe0f",
	":woman_in_steamy_room::skin-tone-4:": "1f9d6-1f3fd-200d-2640-fe0f",
	":woman_in_steamy_room::skin-tone-5:": "1f9d6-1f3fe-200d-2640-fe0f",
	":woman_in_steamy_room::skin-tone-6:": "1f9d6-1f3ff-200d-2640-fe0f",
	// 1F9D6-200D-2642-FE0F man_in_steamy_room
	":man_in_steamy_room:":              "1f9d6-200d-2642-fe0f",
	":man_in_steamy_room::skin-tone-2:": "1f9d6-1f3fb-200d-2642-fe0f",
	":man_in_steamy_room::skin-tone-3:": "1f9d6-1f3fc-200d-2642-fe0f",
	":man_in_steamy_room::skin-tone-4:": "1f9d6-1f3fd-200d-2642-fe0f",
	":man_in_steamy_room::skin-tone-5:": "1f9d6-1f3fe-200d-2642-fe0f",
	":man_in_steamy_room::skin-tone-6:": "1f9d6-1f3ff-200d-2642-fe0f",
	// 1F9D6 person_in_steamy_room
	":person_in_steamy_room:":              "1f9d6",
	":person_in_steamy_room::skin-tone-2:": "1f9d6-1f3fb",
	":person_in_steamy_room::skin-tone-3:": "1f9d6-1f3fc",
	":person_in_steamy_room::skin-tone-4:": "1f9d6-1f3fd",
	":person_in_steamy_room::skin-tone-5:": "1f9d6-1f3fe",
	":person_in_steamy_room::skin-tone-6:": "1f9d6-1f3ff",
	// 1F9D7-200D-2640-FE0F woman_climbing
	":woman_climbing:":              "1f9d7-200d-2640-fe0f",
	":woman_climbing::skin-tone-2:": "1f9d7-1f3fb-200d-2640-fe0f",
	":woman_climbing::skin-tone-3:": "1f9d7-1f3fc-200d-2640-fe0f",
	":woman_climbing::skin-tone-4:": "1f9d7-1f3fd-200d-2640-fe0f",
	":woman_climbing::skin-tone-5:": "1f9d7-1f3fe-200d-2640-fe0f",
	":woman_climbing::skin-tone-6:": "1f9d7-1f3ff-200d-2640-fe0f",
	// 1F9D7-200D-2642-FE0F man_climbing
	":man_climbing:":              "1f9d7-200d-2642-fe0f",
	":man_climbing::skin-tone-2:": "1f9d7-1f3fb-200d-2642-fe0f",
	":man_climbing::skin-tone-3:": "1f9d7-1f3fc-200d-2642-fe0f",
	":man_climbing::skin-tone-4:": "1f9d7-1f3fd-200d-2642-fe0f",
	":man_climbing::skin-tone-5:": "1f9d7-1f3fe-200d-2642-fe0f",
	":man_climbing::skin-tone-6:": "1f9d7-1f3ff-200d-2642-fe0f",
	// 1F9D7 person_climbing
	":person_climbing:":              "1f9d7",
	":person_climbing::skin-tone-2:": "1f9d7-1f3fb",
	":person_climbing::skin-tone-3:": "1f9d7-1f3fc",
	":person_climbing::skin-tone-4:": "1f9d7-1f3fd",
	":person_climbing::skin-tone-5:": "1f9d7-1f3fe",
	":person_climbing::skin-tone-6:": "1f9d7-1f3ff",
	// 1F9D8-200D-2640-FE0F woman_in_lotus_position
	":woman_in_lotus_position:":              "1f9d8-200d-2640-fe0f",
	":woman_in_lotus_position::skin-tone-2:": "1f9d8-1f3fb-200d-2640-fe0f",
	":woman_in_lotus_position::skin-tone-3:": "1f9d8-1f3fc-200d-2640-fe0f",
	":woman_in_lotus_position::skin-tone-4:": "1f9d8-1f3fd-200d-2640-fe0f",
	":woman_in_lotus_position::skin-tone-5:": "1f9d8-1f3fe-200d-2640-fe0f",
	":woman_in_lotus_position::skin-tone-6:": "1f9d8-1f3ff-200d-2640-fe0f",
	// 1F9D8-200D-2642-FE0F man_in_lotus_position
	":man_in_lotus_position:":              "1f9d8-200d-2642-fe0f",
	":man_in_lotus_position::skin-tone-2:": "1f9d8-1f3fb-200d-2642-fe0f",
	":man_in_lotus_position::skin-tone-3:": "1f9d8-1f3fc-200d-2642-fe0f",
	":man_in_lotus_position::skin-tone-4:": "1f9d8-1f3fd-200d-2642-fe0f",
	":man_in_lotus_position::skin-tone-5:": "1f9d8-1f3fe-200d-2642-fe0f",
	":man_in_lotus_position::skin-tone-6:": "1f9d8-1f3ff-200d-2642-fe0f",
	// 1F9D8 person_in_lotus_position
	":person_in_lotus_position:":              "1f9d8",
	":person_in_lotus_position::skin-tone-2:": "1f9d8-1f3fb",
	":person_in_lotus_position::skin-tone-3:": "1f9d8-1f3fc",
	":person_in_lotus_position::skin-tone-4:": "1f9d8-1f3fd",
	":person_in_lotus_position::skin-tone-5:": "1f9d8-1f3fe",
	":person_in_lotus_position::skin-tone-6:": "1f9d8-1f3ff",
	// 1F9D9-200D-2640-FE0F female_mage
	":female_mage:":              "1f9d9-200d-2640-fe0f",
	":female_mage::skin-tone-2:": "1f9d9-1f3fb-200d-2640-fe0f",
	":female_mage::skin-tone-3:": "1f9d9-1f3fc-200d-2640-fe0f",
	":female_mage::skin-tone-4:": "1f9d9-1f3fd-200d-2640-fe0f",
	":female_mage::skin-tone-5:": "1f9d9-1f3fe-200d-2640-fe0f",
	":female_mage::skin-tone-6:": "1f9d9-1f3ff-200d-2640-fe0f",
	// 1F9D9-200D-2642-FE0F male_mage
	":male_mage:":              "1f9d9-200d-2642-fe0f",
	":male_mage::skin-tone-2:": "1f9d9-1f3fb-200d-2642-fe0f",
	":male_mage::skin-tone-3:": "1f9d9-1f3fc-200d-2642-fe0f",
	":male_mage::skin-tone-4:": "1f9d9-1f3fd-200d-2642-fe0f",
	":male_mage::skin-tone-5:": "1f9d9-1f3fe-200d-2642-fe0f",
	":male_mage::skin-tone-6:": "1f9d9-1f3ff-200d-2642-fe0f",
	// 1F9D9 mage
	":mage:":              "1f9d9",
	":mage::skin-tone-2:": "1f9d9-1f3fb",
	":mage::skin-tone-3:": "1f9d9-1f3fc",
	":mage::skin-tone-4:": "1f9d9-1f3fd",
	":mage::skin-tone-5:": "1f9d9-1f3fe",
	":mage::skin-tone-6:": "1f9d9-1f3ff",
	// 1F9DA-200D-2640-FE0F female_fairy
	":female_fairy:":              "1f9da-200d-2640-fe0f",
	":female_fairy::skin-tone-2:": "1f9da-1f3fb-200d-2640-fe0f",
	":female_fairy::skin-tone-3:": "1f9da-1f3fc-200d-2640-fe0f",
	":female_fairy::skin-tone-4:": "1f9da-1f3fd-200d-2640-fe0f",
	":female_fairy::skin-tone-5:": "1f9da-1f3fe-200d-2640-fe0f",
	":female_fairy::skin-tone-6:": "1f9da-1f3ff-200d-2640-fe0f",
	// 1F9DA-200D-2642-FE0F male_fairy
	":male_fairy:":              "1f9da-200d-2642-fe0f",
	":male_fairy::skin-tone-2:": "1f9da-1f3fb-200d-2642-fe0f",
	":male_fairy::skin-tone-3:": "1f9da-1f3fc-200d-2642-fe0f",
	":male_fairy::skin-tone-4:": "1f9da-1f3fd-200d-2642-fe0f",
	":male_fairy::skin-tone-5:": "1f9da-1f3fe-200d-2642-fe0f",
	":male_fairy::skin-tone-6:": "1f9da-1f3ff-200d-2642-fe0f",
	// 1F9DA fairy
	":fairy:":              "1f9da",
	":fairy::skin-tone-2:": "1f9da-1f3fb",
	":fairy::skin-tone-3:": "1f9da-1f3fc",
	":fairy::skin-tone-4:": "1f9da-1f3fd",
	":fairy::skin-tone-5:": "1f9da-1f3fe",
	":fairy::skin-tone-6:": "1f9da-1f3ff",
	// 1F9DB-200D-2640-FE0F female_vampire
	":female_vampire:":              "1f9db-200d-2640-fe0f",
	":female_vampire::skin-tone-2:": "1f9db-1f3fb-200d-2640-fe0f",
	":female_vampire::skin-tone-3:": "1f9db-1f3fc-200d-2640-fe0f",
	":female_vampire::skin-tone-4:": "1f9db-1f3fd-200d-2640-fe0f",
	":female_vampire::skin-tone-5:": "1f9db-1f3fe-200d-2640-fe0f",
	":female_vampire::skin-tone-6:": "1f9db-1f3ff-200d-2640-fe0f",
	// 1F9DB-200D-2642-FE0F male_vampire
	":male_vampire:":              "1f9db-200d-2642-fe0f",
	":male_vampire::skin-tone-2:": "1f9db-1f3fb-200d-2642-fe0f",
	":male_vampire::skin-tone-3:": "1f9db-1f3fc-200d-2642-fe0f",
	":male_vampire::skin-tone-4:": "1f9db-1f3fd-200d-2642-fe0f",
	":male_vampire::skin-tone-5:": "1f9db-1f3fe-200d-2642-fe0f",
	":male_vampire::skin-tone-6:": "1f9db-1f3ff-200d-2642-fe0f",
	// 1F9DB vampire
	":vampire:":              "1f9db",
	":vampire::skin-tone-2:": "1f9db-1f3fb",
	":vampire::skin-tone-3:": "1f9db-1f3fc",
	":vampire::skin-tone-4:": "1f9db-1f3fd",
	":vampire::skin-tone-5:": "1f9db-1f3fe",
	":vampire::skin-tone-6:": "1f9db-1f3ff",
	// 1F9DC-200D-2640-FE0F mermaid
	":mermaid:":              "1f9dc-200d-2640-fe0f",
	":mermaid::skin-tone-2:": "1f9dc-1f3fb-200d-2640-fe0f",
	":mermaid::skin-tone-3:": "1f9dc-1f3fc-200d-2640-fe0f",
	":mermaid::skin-tone-4:": "1f9dc-1f3fd-200d-2640-fe0f",
	":mermaid::skin-tone-5:": "1f9dc-1f3fe-200d-2640-fe0f",
	":mermaid::skin-tone-6:": "1f9dc-1f3ff-200d-2640-fe0f",
	// 1F9DC-200D-2642-FE0F merman
	":merman:":              "1f9dc-200d-2642-fe0f",
	":merman::skin-tone-2:": "1f9dc-1f3fb-200d-2642-fe0f",
	":merman::skin-tone-3:": "1f9dc-1f3fc-200d-2642-fe0f",
	":merman::skin-tone-4:": "1f9dc-1f3fd-200d-2642-fe0f",
	":merman::skin-tone-5:": "1f9dc-1f3fe-200d-2642-fe0f",
	":merman::skin-tone-6:": "1f9dc-1f3ff-200d-2642-fe0f",
	// 1F9DC merperson
	":merperson:":              "1f9dc",
	":merperson::skin-tone-2:": "1f9dc-1f3fb",
	":merperson::skin-tone-3:": "1f9dc-1f3fc",
	":merperson::skin-tone-4:": "1f9dc-1f3fd",
	":merperson::skin-tone-5:": "1f9dc-1f3fe",
	":merperson::skin-tone-6:": "1f9dc-1f3ff",
	// 1F9DD-200D-2640-FE0F female_elf
	":female_elf:":              "1f9dd-200d-2640-fe0f",
	":female_elf::skin-tone-2:": "1f9dd-1f3fb-200d-2640-fe0f",
	":female_elf::skin-tone-3:": "1f9dd-1f3fc-200d-2640-fe0f",
	":female_elf::skin-tone-4:": "1f9dd-1f3fd-200d-2640-fe0f",
	":female_elf::skin-tone-5:": "1f9dd-1f3fe-200d-2640-fe0f",
	":female_elf::skin-tone-6:": "1f9dd-1f3ff-200d-2640-fe0f",
	// 1F9DD-200D-2642-FE0F male_elf
	":male_elf:":              "1f9dd-200d-2642-fe0f",
	":male_elf::skin-tone-2:": "1f9dd-1f3fb-200d-2642-fe0f",
	":male_elf::skin-tone-3:": "1f9dd-1f3fc-200d-2642-fe0f",
	":male_elf::skin-tone-4:": "1f9dd-1f3fd-200d-2642-fe0f",
	":male_elf::skin-tone-5:": "1f9dd-1f3fe-200d-2642-fe0f",
	":male_elf::skin-tone-6:": "1f9dd-1f3ff-200d-2642-fe0f",
	// 1F9DD elf
	":elf:":              "1f9dd",
	":elf::skin-tone-2:": "1f9dd-1f3fb",
	":elf::skin-tone-3:": "1f9dd-1f3fc",
	":elf::skin-tone-4:": "1f9dd-1f3fd",
	":elf::skin-tone-5:": "1f9dd-1f3fe",
	":elf::skin-tone-6:": "1f9dd-1f3ff",
	// 1F9DE-200D-2640-FE0F female_genie
	":female_genie:": "1f9de-200d-2640-fe0f",
	// 1F9DE-200D-2642-FE0F male_genie
	":male_genie:": "1f9de-200d-2642-fe0f",
	// 1F9DE genie
	":genie:": "1f9de",
	// 1F9DF-200D-2640-FE0F female_zombie
	":female_zombie:": "1f9df-200d-2640-fe0f",
	// 1F9DF-200D-2642-FE0F male_zombie
	":male_zombie:": "1f9df-200d-2642-fe0f",
	// 1F9DF zombie
	":zombie:": "1f9df",
	// 1F9E0 brain
	":brain:": "1f9e0",
	// 1F9E1 orange_heart
	":orange_heart:": "1f9e1",
	// 1F9E2 billed_cap
	":billed_cap:": "1f9e2",
	// 1F9E3 scarf
	":scarf:": "1f9e3",
	// 1F9E4 gloves
	":gloves:": "1f9e4",
	// 1F9E5 coat
	":coat:": "1f9e5",
	// 1F9E6 socks
	":socks:": "1f9e6",
	// 1F9E7 red_envelope
	":red_envelope:": "1f9e7",
	// 1F9E8 firecracker
	":firecracker:": "1f9e8",
	// 1F9E9 jigsaw
	":jigsaw:": "1f9e9",
	// 1F9EA test_tube
	":test_tube:": "1f9ea",
	// 1F9EB petri_dish
	":petri_dish:": "1f9eb",
	// 1F9EC dna
	":dna:": "1f9ec",
	// 1F9ED compass
	":compass:": "1f9ed",
	// 1F9EE abacus
	":abacus:": "1f9ee",
	// 1F9EF fire_extinguisher
	":fire_extinguisher:": "1f9ef",
	// 1F9F0 toolbox
	":toolbox:": "1f9f0",
	// 1F9F1 bricks
	":bricks:": "1f9f1",
	// 1F9F2 magnet
	":magnet:": "1f9f2",
	// 1F9F3 luggage
	":luggage:": "1f9f3",
	// 1F9F4 lotion_bottle
	":lotion_bottle:": "1f9f4",
	// 1F9F5 thread
	":thread:": "1f9f5",
	// 1F9F6 yarn
	":yarn:": "1f9f6",
	// 1F9F7 safety_pin
	":safety_pin:": "1f9f7",
	// 1F9F8 teddy_bear
	":teddy_bear:": "1f9f8",
	// 1F9F9 broom
	":broom:": "1f9f9",
	// 1F9FA basket
	":basket:": "1f9fa",
	// 1F9FB roll_of_paper
	":roll_of_paper:": "1f9fb",
	// 1F9FC soap
	":soap:": "1f9fc",
	// 1F9FD sponge
	":sponge:": "1f9fd",
	// 1F9FE receipt
	":receipt:": "1f9fe",
	// 1F9FF nazar_amulet
	":nazar_amulet:": "1f9ff",
	// 1FA70 ballet_shoes
	":ballet_shoes:": "1fa70",
	// 1FA71 one-piece_swimsuit
	":one-piece_swimsuit:": "1fa71",
	// 1FA72 briefs
	":briefs:": "1fa72",
	// 1FA73 shorts
	":shorts:": "1fa73",
	// 1FA74 thong_sandal
	":thong_sandal:": "1fa74",
	// 1FA78 drop_of_blood
	":drop_of_blood:": "1fa78",
	// 1FA79 adhesive_bandage
	":adhesive_bandage:": "1fa79",
	// 1FA7A stethoscope
	":stethoscope:": "1fa7a",
	// 1FA7B x-ray
	":x-ray:": "1fa7b",
	// 1FA7C crutch
	":crutch:": "1fa7c",
	// 1FA80 yo-yo
	":yo-yo:": "1fa80",
	// 1FA81 kite
	":kite:": "1fa81",
	// 1FA82 parachute
	":parachute:": "1fa82",
	// 1FA83 boomerang
	":boomerang:": "1fa83",
	// 1FA84 magic_wand
	":magic_wand:": "1fa84",
	// 1FA85 pinata
	":pinata:": "1fa85",
	// 1FA86 nesting_dolls
	":nesting_dolls:": "1fa86",
	// 1FA90 ringed_planet
	":ringed_planet:": "1fa90",
	// 1FA91 chair
	":chair:": "1fa91",
	// 1FA92 razor
	":razor:": "1fa92",
	// 1FA93 axe
	":axe:": "1fa93",
	// 1FA94 diya_lamp
	":diya_lamp:": "1fa94",
	// 1FA95 banjo
	":banjo:": "1fa95",
	// 1FA96 military_helmet
	":military_helmet:": "1fa96",
	// 1FA97 accordion
	":accordion:": "1fa97",
	// 1FA98 long_drum
	":long_drum:": "1fa98",
	// 1FA99 coin
	":coin:": "1fa99",
	// 1FA9A carpentry_saw
	":carpentry_saw:": "1fa9a",
	// 1FA9B screwdriver
	":screwdriver:": "1fa9b",
	// 1FA9C ladder
	":ladder:": "1fa9c",
	// 1FA9D hook
	":hook:": "1fa9d",
	// 1FA9E mirror
	":mirror:": "1fa9e",
	// 1FA9F window
	":window:": "1fa9f",
	// 1FAA0 plunger
	":plunger:": "1faa0",
	// 1FAA1 sewing_needle
	":sewing_needle:": "1faa1",
	// 1FAA2 knot
	":knot:": "1faa2",
	// 1FAA3 bucket
	":bucket:": "1faa3",
	// 1FAA4 mouse_trap
	":mouse_trap:": "1faa4",
	// 1FAA5 toothbrush
	":toothbrush:": "1faa5",
	// 1FAA6 headstone
	":headstone:": "1faa6",
	// 1FAA7 placard
	":placard:": "1faa7",
	// 1FAA8 rock
	":rock:": "1faa8",
	// 1FAA9 mirror_ball
	":mirror_ball:": "1faa9",
	// 1FAAA identification_card
	":identification_card:": "1faaa",
	// 1FAAB low_battery
	":low_battery:": "1faab",
	// 1FAAC hamsa
	":hamsa:": "1faac",
	// 1FAB0 fly
	":fly:": "1fab0",
	// 1FAB1 worm
	":worm:": "1fab1",
	// 1FAB2 beetle
	":beetle:": "1fab2",
	// 1FAB3 cockroach
	":cockroach:": "1fab3",
	// 1FAB4 potted_plant
	":potted_plant:": "1fab4",
	// 1FAB5 wood
	":wood:": "1fab5",
	// 1FAB6 feather
	":feather:": "1fab6",
	// 1FAB7 lotus
	":lotus:": "1fab7",
	// 1FAB8 coral
	":coral:": "1fab8",
	// 1FAB9 empty_nest
	":empty_nest:": "1fab9",
	// 1FABA nest_with_eggs
	":nest_with_eggs:": "1faba",
	// 1FAC0 anatomical_heart
	":anatomical_heart:": "1fac0",
	// 1FAC1 lungs
	":lungs:": "1fac1",
	// 1FAC2 people_hugging
	":people_hugging:": "1fac2",
	// 1FAC3 pregnant_man
	":pregnant_man:":              "1fac3",
	":pregnant_man::skin-tone-2:": "1fac3-1f3fb",
	":pregnant_man::skin-tone-3:": "1fac3-1f3fc",
	":pregnant_man::skin-tone-4:": "1fac3-1f3fd",
	":pregnant_man::skin-tone-5:": "1fac3-1f3fe",
	":pregnant_man::skin-tone-6:": "1fac3-1f3ff",
	// 1FAC4 pregnant_person
	":pregnant_person:":              "1fac4",
	":pregnant_person::skin-tone-2:": "1fac4-1f3fb",
	":pregnant_person::skin-tone-3:": "1fac4-1f3fc",
	":pregnant_person::skin-tone-4:": "1fac4-1f3fd",
	":pregnant_person::skin-tone-5:": "1fac4-1f3fe",
	":pregnant_person::skin-tone-6:": "1fac4-1f3ff",
	// 1FAC5 person_with_crown
	":person_with_crown:":              "1fac5",
	":person_with_crown::skin-tone-2:": "1fac5-1f3fb",
	":person_with_crown::skin-tone-3:": "1fac5-1f3fc",
	":person_with_crown::skin-tone-4:": "1fac5-1f3fd",
	":person_with_crown::skin-tone-5:": "1fac5-1f3fe",
	":person_with_crown::skin-tone-6:": "1fac5-1f3ff",
	// 1FAD0 blueberries
	":blueberries:": "1fad0",
	// 1FAD1 bell_pepper
	":bell_pepper:": "1fad1",
	// 1FAD2 olive
	":olive:": "1fad2",
	// 1FAD3 flatbread
	":flatbread:": "1fad3",
	// 1FAD4 tamale
	":tamale:": "1fad4",
	// 1FAD5 fondue
	":fondue:": "1fad5",
	// 1FAD6 teapot
	":teapot:": "1fad6",
	// 1FAD7 pouring_liquid
	":pouring_liquid:": "1fad7",
	// 1FAD8 beans
	":beans:": "1fad8",
	// 1FAD9 jar
	":jar:": "1fad9",
	// 1FAE0 melting_face
	":melting_face:": "1fae0",
	// 1FAE1 saluting_face
	":saluting_face:": "1fae1",
	// 1FAE2 face_with_open_eyes_and_hand_over_mouth
	":face_with_open_eyes_and_hand_over_mouth:": "1fae2",
	// 1FAE3 face_with_peeking_eye
	":face_with_peeking_eye:": "1fae3",
	// 1FAE4 face_with_diagonal_mouth
	":face_with_diagonal_mouth:": "1fae4",
	// 1FAE5 dotted_line_face
	":dotted_line_face:": "1fae5",
	// 1FAE6 biting_lip
	":biting_lip:": "1fae6",
	// 1FAE7 bubbles
	":bubbles:": "1fae7",
	// 1FAF0 hand_with_index_finger_and_thumb_crossed
	":hand_with_index_finger_and_thumb_crossed:":              "1faf0",
	":hand_with_index_finger_and_thumb_crossed::skin-tone-2:": "1faf0-1f3fb",
	":hand_with_index_finger_and_thumb_crossed::skin-tone-3:": "1faf0-1f3fc",
	":hand_with_index_finger_and_thumb_crossed::skin-tone-4:": "1faf0-1f3fd",
	":hand_with_index_finger_and_thumb_crossed::skin-tone-5:": "1faf0-1f3fe",
	":hand_with_index_finger_and_thumb_crossed::skin-tone-6:": "1faf0-1f3ff",
	// 1FAF1 rightwards_hand
	":rightwards_hand:":              "1faf1",
	":rightwards_hand::skin-tone-2:": "1faf1-1f3fb",
	":rightwards_hand::skin-tone-3:": "1faf1-1f3fc",
	":rightwards_hand::skin-tone-4:": "1faf1-1f3fd",
	":rightwards_hand::skin-tone-5:": "1faf1-1f3fe",
	":rightwards_hand::skin-tone-6:": "1faf1-1f3ff",
	// 1FAF2 leftwards_hand
	":leftwards_hand:":              "1faf2",
	":leftwards_hand::skin-tone-2:": "1faf2-1f3fb",
	":leftwards_hand::skin-tone-3:": "1faf2-1f3fc",
	":leftwards_hand::skin-tone-4:": "1faf2-1f3fd",
	":leftwards_hand::skin-tone-5:": "1faf2-1f3fe",
	":leftwards_hand::skin-tone-6:": "1faf2-1f3ff",
	// 1FAF3 palm_down_hand
	":palm_down_hand:":              "1faf3",
	":palm_down_hand::skin-tone-2:": "1faf3-1f3fb",
	":palm_down_hand::skin-tone-3:": "1faf3-1f3fc",
	":palm_down_hand::skin-tone-4:": "1faf3-1f3fd",
	":palm_down_hand::skin-tone-5:": "1faf3-1f3fe",
	":palm_down_hand::skin-tone-6:": "1faf3-1f3ff",
	// 1FAF4 palm_up_hand
	":palm_up_hand:":              "1faf4",
	":palm_up_hand::skin-tone-2:": "1faf4-1f3fb",
	":palm_up_hand::skin-tone-3:": "1faf4-1f3fc",
	":palm_up_hand::skin-tone-4:": "1faf4-1f3fd",
	":palm_up_hand::skin-tone-5:": "1faf4-1f3fe",
	":palm_up_hand::skin-tone-6:": "1faf4-1f3ff",
	// 1FAF5 index_pointing_at_the_viewer
	":index_pointing_at_the_viewer:":              "1faf5",
	":index_pointing_at_the_viewer::skin-tone-2:": "1faf5-1f3fb",
	":index_pointing_at_the_viewer::skin-tone-3:": "1faf5-1f3fc",
	":index_pointing_at_the_viewer::skin-tone-4:": "1faf5-1f3fd",
	":index_pointing_at_the_viewer::skin-tone-5:": "1faf5-1f3fe",
	":index_pointing_at_the_viewer::skin-tone-6:": "1faf5-1f3ff",
	// 1FAF6 heart_hands
	":heart_hands:":              "1faf6",
	":heart_hands::skin-tone-2:": "1faf6-1f3fb",
	":heart_hands::skin-tone-3:": "1faf6-1f3fc",
	":heart_hands::skin-tone-4:": "1faf6-1f3fd",
	":heart_hands::skin-tone-5:": "1faf6-1f3fe",
	":heart_hands::skin-tone-6:": "1faf6-1f3ff",
	// 203C-FE0F bangbang
	":bangbang:": "203c-fe0f",
	// 2049-FE0F interrobang
	":interrobang:": "2049-fe0f",
	// 2122-FE0F tm
	":tm:": "2122-fe0f",
	// 2139-FE0F information_source
	":information_source:": "2139-fe0f",
	// 2194-FE0F left_right_arrow
	":left_right_arrow:": "2194-fe0f",
	// 2195-FE0F arrow_up_down
	":arrow_up_down:": "2195-fe0f",
	// 2196-FE0F arrow_upper_left
	":arrow_upper_left:": "2196-fe0f",
	// 2197-FE0F arrow_upper_right
	":arrow_upper_right:": "2197-fe0f",
	// 2198-FE0F arrow_lower_right
	":arrow_lower_right:": "2198-fe0f",
	// 2199-FE0F arrow_lower_left
	":arrow_lower_left:": "2199-fe0f",
	// 21A9-FE0F leftwards_arrow_with_hook
	":leftwards_arrow_with_hook:": "21a9-fe0f",
	// 21AA-FE0F arrow_right_hook
	":arrow_right_hook:": "21aa-fe0f",
	// 231A watch
	":watch:": "231a",
	// 231B hourglass
	":hourglass:": "231b",
	// 2328-FE0F keyboard
	":keyboard:": "2328-fe0f",
	// 23CF-FE0F eject
	":eject:": "23cf-fe0f",
	// 23E9 fast_forward
	":fast_forward:": "23e9",
	// 23EA rewind
	":rewind:": "23ea",
	// 23EB arrow_double_up
	":arrow_double_up:": "23eb",
	// 23EC arrow_double_down
	":arrow_double_down:": "23ec",
	// 23ED-FE0F black_right_pointing_double_triangle_with_vertical_bar
	":black_right_pointing_double_triangle_with_vertical_bar:": "23ed-fe0f",
	// 23EE-FE0F black_left_pointing_double_triangle_with_vertical_bar
	":black_left_pointing_double_triangle_with_vertical_bar:": "23ee-fe0f",
	// 23EF-FE0F black_right_pointing_triangle_with_double_vertical_bar
	":black_right_pointing_triangle_with_double_vertical_bar:": "23ef-fe0f",
	// 23F0 alarm_clock
	":alarm_clock:": "23f0",
	// 23F1-FE0F stopwatch
	":stopwatch:": "23f1-fe0f",
	// 23F2-FE0F timer_clock
	":timer_clock:": "23f2-fe0f",
	// 23F3 hourglass_flowing_sand
	":hourglass_flowing_sand:": "23f3",
	// 23F8-FE0F double_vertical_bar
	":double_vertical_bar:": "23f8-fe0f",
	// 23F9-FE0F black_square_for_stop
	":black_square_for_stop:": "23f9-fe0f",
	// 23FA-FE0F black_circle_for_record
	":black_circle_for_record:": "23fa-fe0f",
	// 24C2-FE0F m
	":m:": "24c2-fe0f",
	// 25AA-FE0F black_small_square
	":black_small_square:": "25aa-fe0f",
	// 25AB-FE0F white_small_square
	":white_small_square:": "25ab-fe0f",
	// 25B6-FE0F arrow_forward
	":arrow_forward:": "25b6-fe0f",
	// 25C0-FE0F arrow_backward
	":arrow_backward:": "25c0-fe0f",
	// 25FB-FE0F white_medium_square
	":white_medium_square:": "25fb-fe0f",
	// 25FC-FE0F black_medium_square
	":black_medium_square:": "25fc-fe0f",
	// 25FD white_medium_small_square
	":white_medium_small_square:": "25fd",
	// 25FE black_medium_small_square
	":black_medium_small_square:": "25fe",
	// 2600-FE0F sunny
	":sunny:": "2600-fe0f",
	// 2601-FE0F cloud
	":cloud:": "2601-fe0f",
	// 2602-FE0F umbrella
	":umbrella:": "2602-fe0f",
	// 2603-FE0F snowman
	":snowman:": "2603-fe0f",
	// 2604-FE0F comet
	":comet:": "2604-fe0f",
	// 260E-FE0F phone
	":phone:": "260e-fe0f",
	// 260E-FE0F phone
	":telephone:": "260e-fe0f",
	// 2611-FE0F ballot_box_with_check
	":ballot_box_with_check:": "2611-fe0f",
	// 2614 umbrella_with_rain_drops
	":umbrella_with_rain_drops:": "2614",
	// 2615 coffee
	":coffee:": "2615",
	// 2618-FE0F shamrock
	":shamrock:": "2618-fe0f",
	// 261D-FE0F point_up
	":point_up:":              "261d-fe0f",
	":point_up::skin-tone-2:": "261d-1f3fb",
	":point_up::skin-tone-3:": "261d-1f3fc",
	":point_up::skin-tone-4:": "261d-1f3fd",
	":point_up::skin-tone-5:": "261d-1f3fe",
	":point_up::skin-tone-6:": "261d-1f3ff",
	// 2620-FE0F skull_and_crossbones
	":skull_and_crossbones:": "2620-fe0f",
	// 2622-FE0F radioactive_sign
	":radioactive_sign:": "2622-fe0f",
	// 2623-FE0F biohazard_sign
	":biohazard_sign:": "2623-fe0f",
	// 2626-FE0F orthodox_cross
	":orthodox_cross:": "2626-fe0f",
	// 262A-FE0F star_and_crescent
	":star_and_crescent:": "262a-fe0f",
	// 262E-FE0F peace_symbol
	":peace_symbol:": "262e-fe0f",
	// 262F-FE0F yin_yang
	":yin_yang:": "262f-fe0f",
	// 2638-FE0F wheel_of_dharma
	":wheel_of_dharma:": "2638-fe0f",
	// 2639-FE0F white_frowning_face
	":white_frowning_face:": "2639-fe0f",
	// 263A-FE0F relaxed
	":relaxed:": "263a-fe0f",
	// 2640-FE0F female_sign
	":female_sign:": "2640-fe0f",
	// 2642-FE0F male_sign
	":male_sign:": "2642-fe0f",
	// 2648 aries
	":aries:": "2648",
	// 2649 taurus
	":taurus:": "2649",
	// 264A gemini
	":gemini:": "264a",
	// 264B cancer
	":cancer:": "264b",
	// 264C leo
	":leo:": "264c",
	// 264D virgo
	":virgo:": "264d",
	// 264E libra
	":libra:": "264e",
	// 264F scorpius
	":scorpius:": "264f",
	// 2650 sagittarius
	":sagittarius:": "2650",
	// 2651 capricorn
	":capricorn:": "2651",
	// 2652 aquarius
	":aquarius:": "2652",
	// 2653 pisces
	":pisces:": "2653",
	// 265F-FE0F chess_pawn
	":chess_pawn:": "265f-fe0f",
	// 2660-FE0F spades
	":spades:": "2660-fe0f",
	// 2663-FE0F clubs
	":clubs:": "2663-fe0f",
	// 2665-FE0F hearts
	":hearts:": "2665-fe0f",
	// 2666-FE0F diamonds
	":diamonds:": "2666-fe0f",
	// 2668-FE0F hotsprings
	":hotsprings:": "2668-fe0f",
	// 267B-FE0F recycle
	":recycle:": "267b-fe0f",
	// 267E-FE0F infinity
	":infinity:": "267e-fe0f",
	// 267F wheelchair
	":wheelchair:": "267f",
	// 2692-FE0F hammer_and_pick
	":hammer_and_pick:": "2692-fe0f",
	// 2693 anchor
	":anchor:": "2693",
	// 2694-FE0F crossed_swords
	":crossed_swords:": "2694-fe0f",
	// 2695-FE0F medical_symbol
	":medical_symbol:": "2695-fe0f",
	// 2695-FE0F medical_symbol
	":staff_of_aesculapius:": "2695-fe0f",
	// 2696-FE0F scales
	":scales:": "2696-fe0f",
	// 2697-FE0F alembic
	":alembic:": "2697-fe0f",
	// 2699-FE0F gear
	":gear:": "2699-fe0f",
	// 269B-FE0F atom_symbol
	":atom_symbol:": "269b-fe0f",
	// 269C-FE0F fleur_de_lis
	":fleur_de_lis:": "269c-fe0f",
	// 26A0-FE0F warning
	":warning:": "26a0-fe0f",
	// 26A1 zap
	":zap:": "26a1",
	// 26A7-FE0F transgender_symbol
	":transgender_symbol:": "26a7-fe0f",
	// 26AA white_circle
	":white_circle:": "26aa",
	// 26AB black_circle
	":black_circle:": "26ab",
	// 26B0-FE0F coffin
	":coffin:": "26b0-fe0f",
	// 26B1-FE0F funeral_urn
	":funeral_urn:": "26b1-fe0f",
	// 26BD soccer
	":soccer:": "26bd",
	// 26BE baseball
	":baseball:": "26be",
	// 26C4 snowman_without_snow
	":snowman_without_snow:": "26c4",
	// 26C5 partly_sunny
	":partly_sunny:": "26c5",
	// 26C8-FE0F thunder_cloud_and_rain
	":thunder_cloud_and_rain:": "26c8-fe0f",
	// 26CE ophiuchus
	":ophiuchus:": "26ce",
	// 26CF-FE0F pick
	":pick:": "26cf-fe0f",
	// 26D1-FE0F helmet_with_white_cross
	":helmet_with_white_cross:": "26d1-fe0f",
	// 26D3-FE0F chains
	":chains:": "26d3-fe0f",
	// 26D4 no_entry
	":no_entry:": "26d4",
	// 26E9-FE0F shinto_shrine
	":shinto_shrine:": "26e9-fe0f",
	// 26EA church
	":church:": "26ea",
	// 26F0-FE0F mountain
	":mountain:": "26f0-fe0f",
	// 26F1-FE0F umbrella_on_ground
	":umbrella_on_ground:": "26f1-fe0f",
	// 26F2 fountain
	":fountain:": "26f2",
	// 26F3 golf
	":golf:": "26f3",
	// 26F4-FE0F ferry
	":ferry:": "26f4-fe0f",
	// 26F5 boat
	":boat:": "26f5",
	// 26F5 boat
	":sailboat:": "26f5",
	// 26F7-FE0F skier
	":skier:": "26f7-fe0f",
	// 26F8-FE0F ice_skate
	":ice_skate:": "26f8-fe0f",
	// 26F9-FE0F-200D-2640-FE0F woman-bouncing-ball
	":woman-bouncing-ball:":              "26f9-fe0f-200d-2640-fe0f",
	":woman-bouncing-ball::skin-tone-2:": "26f9-1f3fb-200d-2640-fe0f",
	":woman-bouncing-ball::skin-tone-3:": "26f9-1f3fc-200d-2640-fe0f",
	":woman-bouncing-ball::skin-tone-4:": "26f9-1f3fd-200d-2640-fe0f",
	":woman-bouncing-ball::skin-tone-5:": "26f9-1f3fe-200d-2640-fe0f",
	":woman-bouncing-ball::skin-tone-6:": "26f9-1f3ff-200d-2640-fe0f",
	// 26F9-FE0F-200D-2642-FE0F man-bouncing-ball
	":man-bouncing-ball:":              "26f9-fe0f-200d-2642-fe0f",
	":man-bouncing-ball::skin-tone-2:": "26f9-1f3fb-200d-2642-fe0f",
	":man-bouncing-ball::skin-tone-3:": "26f9-1f3fc-200d-2642-fe0f",
	":man-bouncing-ball::skin-tone-4:": "26f9-1f3fd-200d-2642-fe0f",
	":man-bouncing-ball::skin-tone-5:": "26f9-1f3fe-200d-2642-fe0f",
	":man-bouncing-ball::skin-tone-6:": "26f9-1f3ff-200d-2642-fe0f",
	// 26F9-FE0F person_with_ball
	":person_with_ball:":              "26f9-fe0f",
	":person_with_ball::skin-tone-2:": "26f9-1f3fb",
	":person_with_ball::skin-tone-3:": "26f9-1f3fc",
	":person_with_ball::skin-tone-4:": "26f9-1f3fd",
	":person_with_ball::skin-tone-5:": "26f9-1f3fe",
	":person_with_ball::skin-tone-6:": "26f9-1f3ff",
	// 26FA tent
	":tent:": "26fa",
	// 26FD fuelpump
	":fuelpump:": "26fd",
	// 2702-FE0F scissors
	":scissors:": "2702-fe0f",
	// 2705 white_check_mark
	":white_check_mark:": "2705",
	// 2708-FE0F airplane
	":airplane:": "2708-fe0f",
	// 2709-FE0F email
	":email:": "2709-fe0f",
	// 2709-FE0F email
	":envelope:": "2709-fe0f",
	// 270A fist
	":fist:":              "270a",
	":fist::skin-tone-2:": "270a-1f3fb",
	":fist::skin-tone-3:": "270a-1f3fc",
	":fist::skin-tone-4:": "270a-1f3fd",
	":fist::skin-tone-5:": "270a-1f3fe",
	":fist::skin-tone-6:": "270a-1f3ff",
	// 270B hand
	":hand:": "270b",
	// 270B hand
	":raised_hand:":              "270b",
	":hand::skin-tone-2:":        "270b-1f3fb",
	":raised_hand::skin-tone-2:": "270b-1f3fb",
	":hand::skin-tone-3:":        "270b-1f3fc",
	":raised_hand::skin-tone-3:": "270b-1f3fc",
	":hand::skin-tone-4:":        "270b-1f3fd",
	":raised_hand::skin-tone-4:": "270b-1f3fd",
	":hand::skin-tone-5:":        "270b-1f3fe",
	":raised_hand::skin-tone-5:": "270b-1f3fe",
	":hand::skin-tone-6:":        "270b-1f3ff",
	":raised_hand::skin-tone-6:": "270b-1f3ff",
	// 270C-FE0F v
	":v:":              "270c-fe0f",
	":v::skin-tone-2:": "270c-1f3fb",
	":v::skin-tone-3:": "270c-1f3fc",
	":v::skin-tone-4:": "270c-1f3fd",
	":v::skin-tone-5:": "270c-1f3fe",
	":v::skin-tone-6:": "270c-1f3ff",
	// 270D-FE0F writing_hand
	":writing_hand:":              "270d-fe0f",
	":writing_hand::skin-tone-2:": "270d-1f3fb",
	":writing_hand::skin-tone-3:": "270d-1f3fc",
	":writing_hand::skin-tone-4:": "270d-1f3fd",
	":writing_hand::skin-tone-5:": "270d-1f3fe",
	":writing_hand::skin-tone-6:": "270d-1f3ff",
	// 270F-FE0F pencil2
	":pencil2:": "270f-fe0f",
	// 2712-FE0F black_nib
	":black_nib:": "2712-fe0f",
	// 2714-FE0F heavy_check_mark
	":heavy_check_mark:": "2714-fe0f",
	// 2716-FE0F heavy_multiplication_x
	":heavy_multiplication_x:": "2716-fe0f",
	// 271D-FE0F latin_cross
	":latin_cross:": "271d-fe0f",
	// 2721-FE0F star_of_david
	":star_of_david:": "2721-fe0f",
	// 2728 sparkles
	":sparkles:": "2728",
	// 2733-FE0F eight_spoked_asterisk
	":eight_spoked_asterisk:": "2733-fe0f",
	// 2734-FE0F eight_pointed_black_star
	":eight_pointed_black_star:": "2734-fe0f",
	// 2744-FE0F snowflake
	":snowflake:": "2744-fe0f",
	// 2747-FE0F sparkle
	":sparkle:": "2747-fe0f",
	// 274C x
	":x:": "274c",
	// 274E negative_squared_cross_mark
	":negative_squared_cross_mark:": "274e",
	// 2753 question
	":question:": "2753",
	// 2754 grey_question
	":grey_question:": "2754",
	// 2755 grey_exclamation
	":grey_exclamation:": "2755",
	// 2757 exclamation
	":exclamation:": "2757",
	// 2757 exclamation
	":heavy_exclamation_mark:": "2757",
	// 2763-FE0F heavy_heart_exclamation_mark_ornament
	":heavy_heart_exclamation_mark_ornament:": "2763-fe0f",
	// 2764-FE0F-200D-1F525 heart_on_fire
	":heart_on_fire:": "2764-fe0f-200d-1f525",
	// 2764-FE0F-200D-1FA79 mending_heart
	":mending_heart:": "2764-fe0f-200d-1fa79",
	// 2764-FE0F heart
	":heart:": "2764-fe0f",
	// 2795 heavy_plus_sign
	":heavy_plus_sign:": "2795",
	// 2796 heavy_minus_sign
	":heavy_minus_sign:": "2796",
	// 2797 heavy_division_sign
	":heavy_division_sign:": "2797",
	// 27A1-FE0F arrow_right
	":arrow_right:": "27a1-fe0f",
	// 27B0 curly_loop
	":curly_loop:": "27b0",
	// 27BF loop
	":loop:": "27bf",
	// 2934-FE0F arrow_heading_up
	":arrow_heading_up:": "2934-fe0f",
	// 2935-FE0F arrow_heading_down
	":arrow_heading_down:": "2935-fe0f",
	// 2B05-FE0F arrow_left
	":arrow_left:": "2b05-fe0f",
	// 2B06-FE0F arrow_up
	":arrow_up:": "2b06-fe0f",
	// 2B07-FE0F arrow_down
	":arrow_down:": "2b07-fe0f",
	// 2B1B black_large_square
	":black_large_square:": "2b1b",
	// 2B1C white_large_square
	":white_large_square:": "2b1c",
	// 2B50 star
	":star:": "2b50",
	// 2B55 o
	":o:": "2b55",
	// 3030-FE0F wavy_dash
	":wavy_dash:": "3030-fe0f",
	// 303D-FE0F part_alternation_mark
	":part_alternation_mark:": "303d-fe0f",
	// 3297-FE0F congratulations
	":congratulations:": "3297-fe0f",
	// 3299-FE0F secret
	":secret:": "3299-fe0f",
}
