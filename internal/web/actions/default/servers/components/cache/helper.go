package cache

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/components/cache/cacheutils"
	"github.com/iwind/TeaGo/actions"
	"net/http"
	"reflect"
)

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) {
	action := actionPtr.Object()
	if action.Request.Method != http.MethodGet {
		return
	}

	action.Data["mainTab"] = "component"
	action.Data["secondMenuItem"] = "cache"

	cachePolicyId := action.ParamInt64("cachePolicyId")
	action.Data["cachePolicyId"] = cachePolicyId

	parentActionValue := reflect.ValueOf(actionPtr).Elem().FieldByName("ParentAction")
	if parentActionValue.IsValid() {
		parentAction, isOk := parentActionValue.Interface().(actionutils.ParentAction)
		if isOk {
			action.Data["cachePolicyName"] = cacheutils.FindCachePolicyNameWithoutError(&parentAction, cachePolicyId)
		}
	}
}
