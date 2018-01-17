def exercise_loops_and_functions(x)
  e = 0.00000000001 # 精度
  number1 = 1.0 # 给定一个起始值(任意浮点数)
  number2 = 0.0 # 给定另一个起始值(任意浮点数)

  while(true)
	return number1 if number1 - number2 <= e && number1 - number2 >= -e
	number2 = number1; number1 = number1-(number1*number1 - x)/(2*number1)
  end
end

p exercise_loops_and_functions(2)