/*
 * @lc app=leetcode id=35 lang=rust
 *
 * [35] Search Insert Position
 */

// @lc code=start
impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
    
        let mut max = (nums.len() - 1) as f32;
        let mut min = 0_f32;
    
        let mut middle = ((max - min) / 2_f32).floor();
    
        loop {
            if target == nums[middle as usize] {
                return middle as i32;
            }
    
            if min == max {
                if target > nums[middle as usize] as i32 {
                    return max as i32 + 1;
                }
    
                return max as i32;
            }
    
            if target > nums[middle as usize] as i32 {
                min = middle;
                middle += ((max - min) / 2_f32).ceil();
            } else if target < nums[middle as usize] as i32 {
                if middle == 0.0 {
                    return 0;
                }
                max = middle - 1_f32;
                middle = ((max - min) / 2_f32).ceil();
            }
            
    
        }
    }
}
// @lc code=end

