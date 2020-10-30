package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_chkCity(t *testing.T) {
	item := &models.NewsSetting{
		CitySel:  1,
		CityList: "北京",
	}
	allow := chkCity("北京", item)
	assert.True(t, allow)

	//正选，非设置城市
	item = &models.NewsSetting{
		CitySel:  1,
		CityList: "上海",
	}
	allow = chkCity("北京", item)
	assert.False(t, allow)

	//反选，未命中屏蔽城市
	item = &models.NewsSetting{
		CitySel:  2,
		CityList: "北京",
	}
	allow = chkCity("济南", item)
	assert.True(t, allow)

	//城市过滤关闭，全部通过
	item = &models.NewsSetting{
		CitySel:  0,
		CityList: "北京",
	}
	allow = chkCity("北京", item)
	assert.True(t, allow)

	//正选全部，全部通过
	item = &models.NewsSetting{
		CitySel:  1,
		CityList: "all",
	}
	allow = chkCity("北京", item)
	assert.True(t, allow)

	//反选，全部拦截
	item = &models.NewsSetting{
		CitySel:  2,
		CityList: "all",
	}
	allow = chkCity("北京", item)
	assert.False(t, allow)
}

func Test_chkChannel(t *testing.T) {
	//正选,命中
	item := &models.NewsSetting{
		ChnSel:  1,
		ChnList: "",
	}
	allow := chkChannel("BZ", item)
	assert.True(t, allow)

	//正选,未命中
	item = &models.NewsSetting{
		ChnSel:  1,
		ChnList: "BZ",
	}
	allow = chkChannel("qq", item)
	assert.False(t, allow)

	//反选，未命中
	item = &models.NewsSetting{
		ChnSel:  2,
		ChnList: "BZ",
	}
	allow = chkChannel("QQ", item)
	assert.True(t, allow)

	//过滤关闭，全部通过
	item = &models.NewsSetting{
		ChnSel:  0,
		ChnList: "all",
	}
	allow = chkChannel("BZ", item)
	assert.True(t, allow)

	//反选，全部拦截
	item = &models.NewsSetting{
		ChnSel:  2,
		ChnList: "all",
	}
	allow = chkChannel("BZ", item)
	assert.False(t, allow)
}

func Test_chkHID(t *testing.T) {
	//hid正选，全部通过
	hid := "111-aaa"
	inHids := "111-aaa,111-bbb,111-cccc"
	notHids := "222-aaa,222-bbb,222-cccc"
	item := &models.NewsSetting{
		HidSel:  1,
		HidList: "",
	}
	allow := chkHID(hid, item)
	assert.True(t, allow)

	item = &models.NewsSetting{
		HidSel:  1,
		HidList: "all",
	}
	allow = chkHID(hid, item)
	assert.True(t, allow)
	//正选，在列表中
	item = &models.NewsSetting{
		HidSel:  1,
		HidList: inHids,
	}
	allow = chkHID(hid, item)
	assert.True(t, allow)

	//正选，未在列表中
	item = &models.NewsSetting{
		HidSel:  1,
		HidList: notHids,
	}
	allow = chkHID(hid, item)
	assert.False(t, allow)

	//反选，hid在列表中
	item = &models.NewsSetting{
		HidSel:  2,
		HidList: inHids,
	}
	allow = chkHID(hid, item)
	assert.False(t, allow)

	//反选，hid未在列表中
	item = &models.NewsSetting{
		HidSel:  2,
		HidList: notHids,
	}
	allow = chkHID(hid, item)
	assert.True(t, allow)
}

func Test_compareVersion(t *testing.T) {
	result, err := compareVersion("1.0.0.1", "1.0.0.0")
	assert.Nil(t, err)
	assert.Equal(t, 1, result)

	result, err = compareVersion("1.0.0.1", "1.0.0.1")
	assert.Nil(t, err)
	assert.Equal(t, 0, result)

	result, err = compareVersion("1.0.0.1", "1.0.0.2")
	assert.Nil(t, err)
	assert.Equal(t, -1, result)

}
