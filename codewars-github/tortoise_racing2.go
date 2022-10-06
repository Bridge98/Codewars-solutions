package kata

func Race(v1, v2, g int) [3]int {
  
  if v1 >= v2 {
    return [3]int{-1,-1,-1}
  }

  time := 3600*g/(v2-v1)
  m    := (time%3600)/60
  s    := time%60
  h    := time/3600
  
  return [3]int{h,m,s}
}