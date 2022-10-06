package kata

func Race(v1, v2, g int) [3]int {
  
  if v1 >= v2 {
    return [3]int{-1,-1,-1}
  }
  
  time := 3600*g/(v2-v1)
  return [3]int{time/3600,(time%3600)/60,time%60}
}