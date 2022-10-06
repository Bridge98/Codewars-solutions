package kata

import (
  "strings"
  "strconv"
)


func Is_valid_ip(ip string) bool {

  parts := strings.Split(ip, ".")
  
  if len(parts) != 4 {
    return false
  }
  
  for _, value := range parts {
    if value[0] == '0' && len(value) > 1 {
      return false
    }
    
    num, err := strconv.Atoi(value)
    
    if err != nil {
      return false
    }

    if num > 255 || num < 0 {
      return false
    }

  }
  
  return true
}