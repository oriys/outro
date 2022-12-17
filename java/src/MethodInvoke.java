package org.example;

public class MethodInvoke {
    public static void main(String[] args) {
        int[] nums = new int[]{2, 7, 11, 15};
        final int accumulate = accumulate(nums);
        System.out.println(accumulate);
        final int max = max(nums);
        System.out.println(max);
        final int len = lengthOfLIS(nums);
        System.out.println(len);
    }

    public static int accumulate(int[] nums) {
        int sum = 0;
        for (int num : nums) {
            sum += num;
        }
        return sum;
    }

    public static int max(int[] nums) {
        int max = Integer.MIN_VALUE;
        for (int num : nums) {
            if (num > max) {
                max = num;
            }
        }
        return max;
    }

    public static int lengthOfLIS(int[] nums) {
        int[] dp = new int[nums.length];
        int max = 0;
        for (int i = 0; i < nums.length; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
            max = Math.max(max, dp[i]);
        }
        return max;
    }

}