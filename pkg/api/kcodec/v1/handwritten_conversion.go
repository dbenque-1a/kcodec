package v1

import (
	"strconv"

	kcodec "github.com/dbenque/kcodec/pkg/api/kcodec"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

//autoConvert_v1_ItemSpec_To_kcodec_ItemSpec replace the generated function that can't handle type conversion
func autoConvert_v1_ItemSpec_To_kcodec_ItemSpec(in *ItemSpec, out *kcodec.ItemSpec, s conversion.Scope) (err error) {
	out.Value, err = strconv.Atoi(in.Value)
	return err
}
