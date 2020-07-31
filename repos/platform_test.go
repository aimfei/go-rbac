package repos

import (
	"go-rbac/model"
	"testing"
)

func TestSave(t *testing.T) {
	plat := model.Platform{PlatformCode: "1223243", PlatformName: "133312312"}
	Save(&plat)
}
