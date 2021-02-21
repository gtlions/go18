package tenxunloc

import (
	"fmt"
	"testing"
)

func TestAddress2Geo(t *testing.T) {
	rsp, err := Address2Geo("福建省福州市台江区升龙汇金中心", "XXXXXX-MHSWZ-SACGG-D7TOS-V2BZ7", "iXXXXXXXXXXwbW25V3G9Al2vqaN")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(rsp)
	t.Logf("操作成功")
}

func TestGeoDistance(t *testing.T) {
	rsp, err := GeoDistance("driving", "39.984092,116.306934", "39.981987,116.362896", "XXXXXX-MHSWZ-SACGG-D7TOS-V2BZ7", "iXXXXXXXXXXwbW25V3G9Al2vqaN")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(rsp)
	t.Logf("操作成功")
}
