package v1ext

import (
	kcodec "github.com/dbenque/kcodec/pkg/api/kcodec"
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
