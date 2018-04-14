package dfl

type NumericBinaryOperator struct {
	*BinaryOperator
}

func (nbo NumericBinaryOperator) EvaluateAndCompare(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (int, error) {

	lv, rv, err := nbo.EvaluateLeftAndRight(ctx, funcs)
	if err != nil {
		return 0, err
	}

	v, err := CompareNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err

}
