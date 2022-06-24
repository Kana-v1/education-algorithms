/*
 * @lc app=leetcode id=9 lang=rust
 *
 * [9] Palindrome Number
 */

// @lc code=start
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        let mut res = x >= 0;
        let mut nums = Vec::<i32>::new();
        let mut variable = x;

        while variable.abs() >= 10 {
            let n = (variable % 10).abs();
            variable = variable / 10;
            nums.push(n);
        }

        nums.push(variable);

        for i in 0..nums.len() / 2 {
            if nums[i] != nums[nums.len() - i - 1] {
                return false;
            }

            res = true;
        }

        res
    }
}
// @lc code=end
