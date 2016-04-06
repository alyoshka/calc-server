enum Operation {
  ADD,
  SUB,
  DIVIDE,
  MULTIPLY
}

service Calculator {
  double calculate(1: Operation operation, 2: double arg1, 3: double arg2)
}
