type tree struct {
  value int
  left, right *tree
}

//sort Sorts values in place
func Sort(values []int){
  var root *tree
  for _, v := range values {
    root = add(root, v)
  }
  appendValues(values[:0], root)
}

//appendValues append the element of t to values in order 
//and returns the resulting slice

func appendValues( values []int, t *tree ) []int {
  if t != nil {
    values = appendValues(values, t.left)
    values = appendValues(values, t)
    values = appendValues(values, t.right)
  }
  return values
}

func add(t *tree, value int) *tree {
  if t == nil {
    //Eqiuvalent to return  &tree {value: value}.
    t = new(tree)
    t.value = value
    return t
  }
  if value < t.value {
    t.left = add(t.left, value)
  } else {
    t.right = add(t.right, value)
  }
  return t
}
