public class Kata {
	// Equal Side function:
	public static int findEvenIndex(int[] arr) {
		
		int sumR = 0, sumL = 0;

		for (int i = 0; i < arr.length; i++) {
			sumR += arr[i];
		}

		for (int i = 0; i < arr.length; i++) {
			
			sumR -= arr[i];

			if (sumL == sumR) {
				return i;
			}

			sumL += arr[i];
		}

		return -1;
	}
}