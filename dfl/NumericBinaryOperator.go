package dfl

type NumericBinaryOperator struct {
	*BinaryOperator
}

func (nbo NumericBinaryOperator) EvaluateAndCompare(ctx map[string]interface{}, funcs FunctionMap) (int, error) {

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
