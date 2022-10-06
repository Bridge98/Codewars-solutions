package uniq

func Uniq(a []string) []string {

	/// check nil cases
	if a == nil {
		return nil
	}

	/**uniq algorithm on golang**/
	if len(a) == 1 {
		return []string{a[0]}
	} else {
		for i := 0; i < len(a); i++ {
			for j := i+1; j < len(a); j++ {
				if a[i] == a[j] {
					a = append(a[:i],a[j:]...)
					j--
				} else {
					break
				}
			}
		}
	}

	return a
}
