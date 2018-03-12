package v1ext

import (
	"strconv"

	kcodec "github.com/dbenque/kcodec/pkg/api/kcodec"
	"github.com/golang/glog"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

//Convert_kcodec_Item_To_v1ext_Item generator can't handle that convertion.
func Convert_kcodec_Item_To_v1ext_Item(in *kcodec.Item, out *Item, s conversion.Scope) (err error) {
	if err = autoConvert_kcodec_Item_To_v1ext_Item(in, out, s); err != nil {
		return err
	}

	if out.Annotations == nil {
		out.Annotations = map[string]string{}
	}
	out.Annotations["specvalue"] = string(in.Spec.Value)
	return nil
}

func autoConvert_v1ext_Item_To_kcodec_Item(in *Item, out *kcodec.Item, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta

	if out.ObjectMeta.Annotations != nil {
		if v, ok := out.ObjectMeta.Annotations["specvalue"]; ok {
			var err error
			if out.Spec.Value, err = strconv.Atoi(v); err == nil {
				delete(out.ObjectMeta.Annotations, "specvalue")
			} else {
				glog.Warningf("can't convert specvalue to int")
			}
		}
	}

	return nil
}
